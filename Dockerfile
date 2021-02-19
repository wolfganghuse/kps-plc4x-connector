# ---------- This container is the build container ---------- #

FROM golang:latest as build

ARG connector_name=kps-connector-template

WORKDIR /go/src/$connector_name
COPY . .
RUN go build -o $connector_name

# ---------- This container is the runtime container ---------- #

FROM golang:latest

ARG connector_name=kps-connector-template
RUN mkdir -p /connectors/bin

COPY --from=build /go/src/$connector_name/$connector_name /connectors/bin/$connector_name
RUN chmod +x /connectors/bin/$connector_name

ARG EXECUTABLE=$connector_name
ENV EXECUTABLE ${EXECUTABLE}

CMD /connectors/bin/${EXECUTABLE}
