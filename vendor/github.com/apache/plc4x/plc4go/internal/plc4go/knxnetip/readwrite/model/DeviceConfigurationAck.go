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
type DeviceConfigurationAck struct {
    DeviceConfigurationAckDataBlock *DeviceConfigurationAckDataBlock
    Parent *KnxNetIpMessage
    IDeviceConfigurationAck
}

// The corresponding interface
type IDeviceConfigurationAck interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *DeviceConfigurationAck) MsgType() uint16 {
    return 0x0311
}


func (m *DeviceConfigurationAck) InitializeParent(parent *KnxNetIpMessage) {
}

func NewDeviceConfigurationAck(deviceConfigurationAckDataBlock *DeviceConfigurationAckDataBlock, ) *KnxNetIpMessage {
    child := &DeviceConfigurationAck{
        DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
        Parent: NewKnxNetIpMessage(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastDeviceConfigurationAck(structType interface{}) *DeviceConfigurationAck {
    castFunc := func(typ interface{}) *DeviceConfigurationAck {
        if casted, ok := typ.(DeviceConfigurationAck); ok {
            return &casted
        }
        if casted, ok := typ.(*DeviceConfigurationAck); ok {
            return casted
        }
        if casted, ok := typ.(KnxNetIpMessage); ok {
            return CastDeviceConfigurationAck(casted.Child)
        }
        if casted, ok := typ.(*KnxNetIpMessage); ok {
            return CastDeviceConfigurationAck(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *DeviceConfigurationAck) GetTypeName() string {
    return "DeviceConfigurationAck"
}

func (m *DeviceConfigurationAck) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (deviceConfigurationAckDataBlock)
    lengthInBits += m.DeviceConfigurationAckDataBlock.LengthInBits()

    return lengthInBits
}

func (m *DeviceConfigurationAck) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func DeviceConfigurationAckParse(io *utils.ReadBuffer) (*KnxNetIpMessage, error) {

    // Simple Field (deviceConfigurationAckDataBlock)
    deviceConfigurationAckDataBlock, _deviceConfigurationAckDataBlockErr := DeviceConfigurationAckDataBlockParse(io)
    if _deviceConfigurationAckDataBlockErr != nil {
        return nil, errors.New("Error parsing 'deviceConfigurationAckDataBlock' field " + _deviceConfigurationAckDataBlockErr.Error())
    }

    // Create a partially initialized instance
    _child := &DeviceConfigurationAck{
        DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
        Parent: &KnxNetIpMessage{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *DeviceConfigurationAck) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (deviceConfigurationAckDataBlock)
    _deviceConfigurationAckDataBlockErr := m.DeviceConfigurationAckDataBlock.Serialize(io)
    if _deviceConfigurationAckDataBlockErr != nil {
        return errors.New("Error serializing 'deviceConfigurationAckDataBlock' field " + _deviceConfigurationAckDataBlockErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *DeviceConfigurationAck) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "deviceConfigurationAckDataBlock":
                var data *DeviceConfigurationAckDataBlock
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.DeviceConfigurationAckDataBlock = data
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

func (m *DeviceConfigurationAck) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.DeviceConfigurationAckDataBlock, xml.StartElement{Name: xml.Name{Local: "deviceConfigurationAckDataBlock"}}); err != nil {
        return err
    }
    return nil
}

