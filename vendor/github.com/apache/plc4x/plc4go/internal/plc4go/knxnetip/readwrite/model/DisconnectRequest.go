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
    log "github.com/sirupsen/logrus"
    "io"
)

// The data-structure of this message
type DisconnectRequest struct {
    CommunicationChannelId uint8
    HpaiControlEndpoint *HPAIControlEndpoint
    Parent *KnxNetIpMessage
    IDisconnectRequest
}

// The corresponding interface
type IDisconnectRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *DisconnectRequest) MsgType() uint16 {
    return 0x0209
}


func (m *DisconnectRequest) InitializeParent(parent *KnxNetIpMessage) {
}

func NewDisconnectRequest(communicationChannelId uint8, hpaiControlEndpoint *HPAIControlEndpoint, ) *KnxNetIpMessage {
    child := &DisconnectRequest{
        CommunicationChannelId: communicationChannelId,
        HpaiControlEndpoint: hpaiControlEndpoint,
        Parent: NewKnxNetIpMessage(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastDisconnectRequest(structType interface{}) *DisconnectRequest {
    castFunc := func(typ interface{}) *DisconnectRequest {
        if casted, ok := typ.(DisconnectRequest); ok {
            return &casted
        }
        if casted, ok := typ.(*DisconnectRequest); ok {
            return casted
        }
        if casted, ok := typ.(KnxNetIpMessage); ok {
            return CastDisconnectRequest(casted.Child)
        }
        if casted, ok := typ.(*KnxNetIpMessage); ok {
            return CastDisconnectRequest(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *DisconnectRequest) GetTypeName() string {
    return "DisconnectRequest"
}

func (m *DisconnectRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (communicationChannelId)
    lengthInBits += 8

    // Reserved Field (reserved)
    lengthInBits += 8

    // Simple field (hpaiControlEndpoint)
    lengthInBits += m.HpaiControlEndpoint.LengthInBits()

    return lengthInBits
}

func (m *DisconnectRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func DisconnectRequestParse(io *utils.ReadBuffer) (*KnxNetIpMessage, error) {

    // Simple Field (communicationChannelId)
    communicationChannelId, _communicationChannelIdErr := io.ReadUint8(8)
    if _communicationChannelIdErr != nil {
        return nil, errors.New("Error parsing 'communicationChannelId' field " + _communicationChannelIdErr.Error())
    }

    // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
    {
        reserved, _err := io.ReadUint8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'reserved' field " + _err.Error())
        }
        if reserved != uint8(0x00) {
            log.WithFields(log.Fields{
                "expected value": uint8(0x00),
                "got value": reserved,
            }).Info("Got unexpected response.")
        }
    }

    // Simple Field (hpaiControlEndpoint)
    hpaiControlEndpoint, _hpaiControlEndpointErr := HPAIControlEndpointParse(io)
    if _hpaiControlEndpointErr != nil {
        return nil, errors.New("Error parsing 'hpaiControlEndpoint' field " + _hpaiControlEndpointErr.Error())
    }

    // Create a partially initialized instance
    _child := &DisconnectRequest{
        CommunicationChannelId: communicationChannelId,
        HpaiControlEndpoint: hpaiControlEndpoint,
        Parent: &KnxNetIpMessage{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *DisconnectRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (communicationChannelId)
    communicationChannelId := uint8(m.CommunicationChannelId)
    _communicationChannelIdErr := io.WriteUint8(8, (communicationChannelId))
    if _communicationChannelIdErr != nil {
        return errors.New("Error serializing 'communicationChannelId' field " + _communicationChannelIdErr.Error())
    }

    // Reserved Field (reserved)
    {
        _err := io.WriteUint8(8, uint8(0x00))
        if _err != nil {
            return errors.New("Error serializing 'reserved' field " + _err.Error())
        }
    }

    // Simple Field (hpaiControlEndpoint)
    _hpaiControlEndpointErr := m.HpaiControlEndpoint.Serialize(io)
    if _hpaiControlEndpointErr != nil {
        return errors.New("Error serializing 'hpaiControlEndpoint' field " + _hpaiControlEndpointErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *DisconnectRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "communicationChannelId":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.CommunicationChannelId = data
            case "hpaiControlEndpoint":
                var data *HPAIControlEndpoint
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.HpaiControlEndpoint = data
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

func (m *DisconnectRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.CommunicationChannelId, xml.StartElement{Name: xml.Name{Local: "communicationChannelId"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.HpaiControlEndpoint, xml.StartElement{Name: xml.Name{Local: "hpaiControlEndpoint"}}); err != nil {
        return err
    }
    return nil
}

