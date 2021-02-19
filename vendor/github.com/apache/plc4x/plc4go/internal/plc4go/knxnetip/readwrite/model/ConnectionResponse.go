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
type ConnectionResponse struct {
    CommunicationChannelId uint8
    Status Status
    HpaiDataEndpoint *HPAIDataEndpoint
    ConnectionResponseDataBlock *ConnectionResponseDataBlock
    Parent *KnxNetIpMessage
    IConnectionResponse
}

// The corresponding interface
type IConnectionResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ConnectionResponse) MsgType() uint16 {
    return 0x0206
}


func (m *ConnectionResponse) InitializeParent(parent *KnxNetIpMessage) {
}

func NewConnectionResponse(communicationChannelId uint8, status Status, hpaiDataEndpoint *HPAIDataEndpoint, connectionResponseDataBlock *ConnectionResponseDataBlock, ) *KnxNetIpMessage {
    child := &ConnectionResponse{
        CommunicationChannelId: communicationChannelId,
        Status: status,
        HpaiDataEndpoint: hpaiDataEndpoint,
        ConnectionResponseDataBlock: connectionResponseDataBlock,
        Parent: NewKnxNetIpMessage(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastConnectionResponse(structType interface{}) *ConnectionResponse {
    castFunc := func(typ interface{}) *ConnectionResponse {
        if casted, ok := typ.(ConnectionResponse); ok {
            return &casted
        }
        if casted, ok := typ.(*ConnectionResponse); ok {
            return casted
        }
        if casted, ok := typ.(KnxNetIpMessage); ok {
            return CastConnectionResponse(casted.Child)
        }
        if casted, ok := typ.(*KnxNetIpMessage); ok {
            return CastConnectionResponse(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ConnectionResponse) GetTypeName() string {
    return "ConnectionResponse"
}

func (m *ConnectionResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (communicationChannelId)
    lengthInBits += 8

    // Simple field (status)
    lengthInBits += 8

    // Optional Field (hpaiDataEndpoint)
    if m.HpaiDataEndpoint != nil {
        lengthInBits += (*m.HpaiDataEndpoint).LengthInBits()
    }

    // Optional Field (connectionResponseDataBlock)
    if m.ConnectionResponseDataBlock != nil {
        lengthInBits += (*m.ConnectionResponseDataBlock).LengthInBits()
    }

    return lengthInBits
}

func (m *ConnectionResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ConnectionResponseParse(io *utils.ReadBuffer) (*KnxNetIpMessage, error) {

    // Simple Field (communicationChannelId)
    communicationChannelId, _communicationChannelIdErr := io.ReadUint8(8)
    if _communicationChannelIdErr != nil {
        return nil, errors.New("Error parsing 'communicationChannelId' field " + _communicationChannelIdErr.Error())
    }

    // Simple Field (status)
    status, _statusErr := StatusParse(io)
    if _statusErr != nil {
        return nil, errors.New("Error parsing 'status' field " + _statusErr.Error())
    }

    // Optional Field (hpaiDataEndpoint) (Can be skipped, if a given expression evaluates to false)
    var hpaiDataEndpoint *HPAIDataEndpoint = nil
    if bool((status) == (Status_NO_ERROR)) {
        _val, _err := HPAIDataEndpointParse(io)
        if _err != nil {
            return nil, errors.New("Error parsing 'hpaiDataEndpoint' field " + _err.Error())
        }
        hpaiDataEndpoint = _val
    }

    // Optional Field (connectionResponseDataBlock) (Can be skipped, if a given expression evaluates to false)
    var connectionResponseDataBlock *ConnectionResponseDataBlock = nil
    if bool((status) == (Status_NO_ERROR)) {
        _val, _err := ConnectionResponseDataBlockParse(io)
        if _err != nil {
            return nil, errors.New("Error parsing 'connectionResponseDataBlock' field " + _err.Error())
        }
        connectionResponseDataBlock = _val
    }

    // Create a partially initialized instance
    _child := &ConnectionResponse{
        CommunicationChannelId: communicationChannelId,
        Status: status,
        HpaiDataEndpoint: hpaiDataEndpoint,
        ConnectionResponseDataBlock: connectionResponseDataBlock,
        Parent: &KnxNetIpMessage{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ConnectionResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (communicationChannelId)
    communicationChannelId := uint8(m.CommunicationChannelId)
    _communicationChannelIdErr := io.WriteUint8(8, (communicationChannelId))
    if _communicationChannelIdErr != nil {
        return errors.New("Error serializing 'communicationChannelId' field " + _communicationChannelIdErr.Error())
    }

    // Simple Field (status)
    _statusErr := m.Status.Serialize(io)
    if _statusErr != nil {
        return errors.New("Error serializing 'status' field " + _statusErr.Error())
    }

    // Optional Field (hpaiDataEndpoint) (Can be skipped, if the value is null)
    var hpaiDataEndpoint *HPAIDataEndpoint = nil
    if m.HpaiDataEndpoint != nil {
        hpaiDataEndpoint = m.HpaiDataEndpoint
        _hpaiDataEndpointErr := hpaiDataEndpoint.Serialize(io)
        if _hpaiDataEndpointErr != nil {
            return errors.New("Error serializing 'hpaiDataEndpoint' field " + _hpaiDataEndpointErr.Error())
        }
    }

    // Optional Field (connectionResponseDataBlock) (Can be skipped, if the value is null)
    var connectionResponseDataBlock *ConnectionResponseDataBlock = nil
    if m.ConnectionResponseDataBlock != nil {
        connectionResponseDataBlock = m.ConnectionResponseDataBlock
        _connectionResponseDataBlockErr := connectionResponseDataBlock.Serialize(io)
        if _connectionResponseDataBlockErr != nil {
            return errors.New("Error serializing 'connectionResponseDataBlock' field " + _connectionResponseDataBlockErr.Error())
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ConnectionResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "status":
                var data Status
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Status = data
            case "hpaiDataEndpoint":
                var data *HPAIDataEndpoint
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.HpaiDataEndpoint = data
            case "connectionResponseDataBlock":
                var dt *ConnectionResponseDataBlock
                if err := d.DecodeElement(&dt, &tok); err != nil {
                    return err
                }
                m.ConnectionResponseDataBlock = dt
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

func (m *ConnectionResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.CommunicationChannelId, xml.StartElement{Name: xml.Name{Local: "communicationChannelId"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Status, xml.StartElement{Name: xml.Name{Local: "status"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.HpaiDataEndpoint, xml.StartElement{Name: xml.Name{Local: "hpaiDataEndpoint"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ConnectionResponseDataBlock, xml.StartElement{Name: xml.Name{Local: "connectionResponseDataBlock"}}); err != nil {
        return err
    }
    return nil
}

