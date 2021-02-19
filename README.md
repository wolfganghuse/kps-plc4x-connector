# Karbon Platform Services (KPS) Connector Go Template
Welcome to the Golang template for building KPS Connectors. This template is an opinionated implementation of the
gRPC service contract defined in [KPS Connector IDL](https://github.com/nutanix/kps-connector-idl) built using the
[KPS Connector Go SDK](https://github.com/nutanix/kps-connector-idl).

The main goal of this template is to provide a basic fully functional shell connector that connector developers
can use to quickly build new connectors. The template implements the gRPC interface, handles stream and config updates,
and publishes basic status and alert events during the connector runtime. 

By using the template to build a custom connector, you only need to write the connection code to get data in from and send data out to a given data
service. The bundled `Makefile` and github actions workflow provide a fully prepared CI setup and ensure Golang code
standards are maintained. The bundled `Dockerfile` builds the code within the container and then copies the binary
to another container that is used to run the binary.

## How Do I Build and Run a Connector?
There are three main phases to build and run a custom KPS Connector.
- Develop a connector
- Package the connector
- Deploy the connector

### Develop a connector
#### Overview
- Start by cloning this repo
- Copy all the contents into another repo in which you intend to build the connector
- Fill in the `TODO` blocks in `connector/template.go` and the connector name in `connector/config.go`
- Replace the default values marked with `TODO` comments for connector name, docker repository URI, and docker image tag in `Makefile`
- Change the module name in `go.mod` file

#### Understanding and updating `connectors/template.go`
`connectors/template.go` has three important structs `streamMetadata`, `consumer` and `producer`
- `streamMetadata` is a go struct used to unmarshall the `streamParameterSchema` defined in
`samples/connector_class.json`
    - `mapToStreamMetadata` is a function that translates the `stream.Metadata` object from an untyped `map[string]interface`
    to a typed struct
- Consumer connects to the custom data service defined by ingress stream when the `subscribe` method is called and makes
the data available in each successive `nextMsg` method call
    - `newConsumer` is a constructor for a `consumer` object
    - `subscribe` method contains the logic for creating a connection to the custom data service defined by the ingress stream
    - `nextMsg` method follows the iterator pattern and provides the next message to be consumed from the connection created by
    the `subscribe` method above.
    - If there is no data to pass, `nextMsg` should block the call. Otherwise the `consumerLoop` in `connector/streams.go` will go in a tight loop
- Producer connects to the custom data service defined by the egress stream when the `connect` method is called. It then
produces the message it receives in the `subscribeMsgHandler` callback to the custom data service
    - `newProducer` is a constructor for the `producer` object
    - `connect` method contains the logic for creating a connection to the custom data service defined by the egress stream
    - `subscribeMsgHandler` is a callback function called each time a message is received and to be published on
    the connection established in the `connect` method above. This message stores the logic for publishing data to the
    corresponding connection.

### Package the connector
- Run `make`.  The bundled `Makefile` and `Dockerfile` helps ensure that the connector is compiled and a corresponding
docker image is created with the docker URI and tag defined in `Makefile`
- Ensure you have `docker push` access to the docker registry specified in the `Makefile`
- Run `make publish`. This step publishes the docker image created in the `make` step to the docker registry defined in `Makefile`

### Deploy the connector
To use the connector, follow these steps:

#### Create a class
Update the `samples/connector_class.json` with the following information.
- JSON schema for static parameters (property `staticParameterSchema`) used for templating. For this template, the
docker image tag is a templatized value that is provided during the instance creation.
- JSON schema for dynamic config (property `configParameterSchema`) used for runtime config change
- JSON schema for stream (property `streamParameterSchema`) used for conveying connection information

You can register the connector class with KPS by using the kps CLI tool.
```
kps create connectorclass -f samples/connector_class.json
```
#### Create the instance
Update the `samples/connector_instance.json` with the following information.
- UUID of connector class registered above (property `connectorClassID`)
- UUID of the project on which you want to create the instance (property `projectId`)
- Values for the template variables (property `staticParameters`) defined earlier in connector class definition as
`staticParameterSchema`

You can create the connector instance by using the kps CLI tool.
```
kps create connectorinstance -f samples/connector_instance.json
```
#### Create the in/out stream
Update the `samples/connector_stream_ingress.json` 
- UUID of the connector instance created earlier (property: `connectorInstanceID`)
- UUIDs of the service domains in the project on which the instance is to be created (property `serviceDomainIds`). 
If no service domains are present in the list, we assume that the user means "All" service domains in the project.
- Categories and corresponding category values (property: `labels`) used for matching the stream
- Values for the stream (property `stream`) defined earlier in connector class definition as `streamParameterSchema`

Update the `samples/connector_stream_egress.json` similar to how you update `samples/connector_stream_ingress.json`. The difference
between ingress and egress stream is the `direction` property (`INGRESS/EGRESS`)`. Also, egress stream does not need labels as
data pipeline can only output to one stream at a time.

You can create the connector streams by using the CLI tool
```
kps create connectorstream -f samples/connector_stream_ingress.json
kps create connectorstream -f samples/connector_stream_egress.json
```
#### Create the config
Update the `samples/connector_config.json` 
- UUID of the connector instance created earlier (property: `connectorInstanceID`)
- UUIDs of the service domains in the project on which the instance is to be created (property `serviceDomainIds`). 
If no service domains are present in the list, we assume that the user means "All" service domains in the project.
- Values for the dynamic config (property `config`) defined earlier in connector class definition as `configParameterSchema`

You can create the connector config by using the CLI tool
```
kps create connectorstream -f samples/connector_config.json
```
#### Create the data pipeline
Update the `samples/pipeline.yaml` with the following information.
- Name of the project (property `project`) on which the connector is created
- List of names of functions in the data pipeline (see `function.yaml` and `echo.py` for creating a function)
- Category selectors for matching the ingress connector streams (property `input.categorySelectors`)
- Output endpoint to match the egress stream (prpoerty `output.localEdge.endpointName`)

You can create the data pipeline by using the kps CLI tool
```
kps create datapipeline -f samples/pipeline.yaml
```

### Questions, issues or suggestions?
Reach us at karbon-platform-services-api@nutanix.com or file an issue on the Github repository.