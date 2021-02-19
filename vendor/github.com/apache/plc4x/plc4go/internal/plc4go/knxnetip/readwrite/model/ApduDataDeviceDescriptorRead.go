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
type ApduDataDeviceDescriptorRead struct {
    DescriptorType uint8
    Parent *ApduData
    IApduDataDeviceDescriptorRead
}

// The corresponding interface
type IApduDataDeviceDescriptorRead interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataDeviceDescriptorRead) ApciType() uint8 {
    return 0xC
}


func (m *ApduDataDeviceDescriptorRead) InitializeParent(parent *ApduData) {
}

func NewApduDataDeviceDescriptorRead(descriptorType uint8, ) *ApduData {
    child := &ApduDataDeviceDescriptorRead{
        DescriptorType: descriptorType,
        Parent: NewApduData(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataDeviceDescriptorRead(structType interface{}) *ApduDataDeviceDescriptorRead {
    castFunc := func(typ interface{}) *ApduDataDeviceDescriptorRead {
        if casted, ok := typ.(ApduDataDeviceDescriptorRead); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataDeviceDescriptorRead); ok {
            return casted
        }
        if casted, ok := typ.(ApduData); ok {
            return CastApduDataDeviceDescriptorRead(casted.Child)
        }
        if casted, ok := typ.(*ApduData); ok {
            return CastApduDataDeviceDescriptorRead(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataDeviceDescriptorRead) GetTypeName() string {
    return "ApduDataDeviceDescriptorRead"
}

func (m *ApduDataDeviceDescriptorRead) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (descriptorType)
    lengthInBits += 6

    return lengthInBits
}

func (m *ApduDataDeviceDescriptorRead) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataDeviceDescriptorReadParse(io *utils.ReadBuffer) (*ApduData, error) {

    // Simple Field (descriptorType)
    descriptorType, _descriptorTypeErr := io.ReadUint8(6)
    if _descriptorTypeErr != nil {
        return nil, errors.New("Error parsing 'descriptorType' field " + _descriptorTypeErr.Error())
    }

    // Create a partially initialized instance
    _child := &ApduDataDeviceDescriptorRead{
        DescriptorType: descriptorType,
        Parent: &ApduData{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataDeviceDescriptorRead) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (descriptorType)
    descriptorType := uint8(m.DescriptorType)
    _descriptorTypeErr := io.WriteUint8(6, (descriptorType))
    if _descriptorTypeErr != nil {
        return errors.New("Error serializing 'descriptorType' field " + _descriptorTypeErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataDeviceDescriptorRead) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "descriptorType":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.DescriptorType = data
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

func (m *ApduDataDeviceDescriptorRead) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.DescriptorType, xml.StartElement{Name: xml.Name{Local: "descriptorType"}}); err != nil {
        return err
    }
    return nil
}

