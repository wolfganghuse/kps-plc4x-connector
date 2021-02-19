//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "encoding/xml"
    "errors"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
)

// The data-structure of this message
type ConnectionRequest struct {
    HpaiDiscoveryEndpoint *HPAIDiscoveryEndpoint
    HpaiDataEndpoint *HPAIDataEndpoint
    ConnectionRequestInformation *ConnectionRequestInformation
    Parent *KnxNetIpMessage
    IConnectionRequest
}

// The corresponding interface
type IConnectionRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ConnectionRequest) MsgType() uint16 {
    return 0x0205
}


func (m *ConnectionRequest) InitializeParent(parent *KnxNetIpMessage) {
}

func NewConnectionRequest(hpaiDiscoveryEndpoint *HPAIDiscoveryEndpoint, hpaiDataEndpoint *HPAIDataEndpoint, connectionRequestInformation *ConnectionRequestInformation, ) *KnxNetIpMessage {
    child := &ConnectionRequest{
        HpaiDiscoveryEndpoint: hpaiDiscoveryEndpoint,
        HpaiDataEndpoint: hpaiDataEndpoint,
        ConnectionRequestInformation: connectionRequestInformation,
        Parent: NewKnxNetIpMessage(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastConnectionRequest(structType interface{}) *ConnectionRequest {
    castFunc := func(typ interface{}) *ConnectionRequest {
        if casted, ok := typ.(ConnectionRequest); ok {
            return &casted
        }
        if casted, ok := typ.(*ConnectionRequest); ok {
            return casted
        }
        if casted, ok := typ.(KnxNetIpMessage); ok {
            return CastConnectionRequest(casted.Child)
        }
        if casted, ok := typ.(*KnxNetIpMessage); ok {
            return CastConnectionRequest(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ConnectionRequest) GetTypeName() string {
    return "ConnectionRequest"
}

func (m *ConnectionRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (hpaiDiscoveryEndpoint)
    lengthInBits += m.HpaiDiscoveryEndpoint.LengthInBits()

    // Simple field (hpaiDataEndpoint)
    lengthInBits += m.HpaiDataEndpoint.LengthInBits()

    // Simple field (connectionRequestInformation)
    lengthInBits += m.ConnectionRequestInformation.LengthInBits()

    return lengthInBits
}

func (m *ConnectionRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ConnectionRequestParse(io *utils.ReadBuffer) (*KnxNetIpMessage, error) {

    // Simple Field (hpaiDiscoveryEndpoint)
    hpaiDiscoveryEndpoint, _hpaiDiscoveryEndpointErr := HPAIDiscoveryEndpointParse(io)
    if _hpaiDiscoveryEndpointErr != nil {
        return nil, errors.New("Error parsing 'hpaiDiscoveryEndpoint' field " + _hpaiDiscoveryEndpointErr.Error())
    }

    // Simple Field (hpaiDataEndpoint)
    hpaiDataEndpoint, _hpaiDataEndpointErr := HPAIDataEndpointParse(io)
    if _hpaiDataEndpointErr != nil {
        return nil, errors.New("Error parsing 'hpaiDataEndpoint' field " + _hpaiDataEndpointErr.Error())
    }

    // Simple Field (connectionRequestInformation)
    connectionRequestInformation, _connectionRequestInformationErr := ConnectionRequestInformationParse(io)
    if _connectionRequestInformationErr != nil {
        return nil, errors.New("Error parsing 'connectionRequestInformation' field " + _connectionRequestInformationErr.Error())
    }

    // Create a partially initialized instance
    _child := &ConnectionRequest{
        HpaiDiscoveryEndpoint: hpaiDiscoveryEndpoint,
        HpaiDataEndpoint: hpaiDataEndpoint,
        ConnectionRequestInformation: connectionRequestInformation,
        Parent: &KnxNetIpMessage{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ConnectionRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (hpaiDiscoveryEndpoint)
    _hpaiDiscoveryEndpointErr := m.HpaiDiscoveryEndpoint.Serialize(io)
    if _hpaiDiscoveryEndpointErr != nil {
        return errors.New("Error serializing 'hpaiDiscoveryEndpoint' field " + _hpaiDiscoveryEndpointErr.Error())
    }

    // Simple Field (hpaiDataEndpoint)
    _hpaiDataEndpointErr := m.HpaiDataEndpoint.Serialize(io)
    if _hpaiDataEndpointErr != nil {
        return errors.New("Error serializing 'hpaiDataEndpoint' field " + _hpaiDataEndpointErr.Error())
    }

    // Simple Field (connectionRequestInformation)
    _connectionRequestInformationErr := m.ConnectionRequestInformation.Serialize(io)
    if _connectionRequestInformationErr != nil {
        return errors.New("Error serializing 'connectionRequestInformation' field " + _connectionRequestInformationErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ConnectionRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "hpaiDiscoveryEndpoint":
                var data *HPAIDiscoveryEndpoint
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.HpaiDiscoveryEndpoint = data
            case "hpaiDataEndpoint":
                var data *HPAIDataEndpoint
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.HpaiDataEndpoint = data
            case "connectionRequestInformation":
                var dt *ConnectionRequestInformation
                if err := d.DecodeElement(&dt, &tok); err != nil {
                    return err
                }
                m.ConnectionRequestInformation = dt
            }
        }
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
}

func (m *ConnectionRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.HpaiDiscoveryEndpoint, xml.StartElement{Name: xml.Name{Local: "hpaiDiscoveryEndpoint"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.HpaiDataEndpoint, xml.StartElement{Name: xml.Name{Local: "hpaiDataEndpoint"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ConnectionRequestInformation, xml.StartElement{Name: xml.Name{Local: "connectionRequestInformation"}}); err != nil {
        return err
    }
    return nil
}

