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
type KnxGroupAddress2Level struct {
    MainGroup uint8
    SubGroup uint16
    Parent *KnxGroupAddress
    IKnxGroupAddress2Level
}

// The corresponding interface
type IKnxGroupAddress2Level interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *KnxGroupAddress2Level) NumLevels() uint8 {
    return 2
}


func (m *KnxGroupAddress2Level) InitializeParent(parent *KnxGroupAddress) {
}

func NewKnxGroupAddress2Level(mainGroup uint8, subGroup uint16, ) *KnxGroupAddress {
    child := &KnxGroupAddress2Level{
        MainGroup: mainGroup,
        SubGroup: subGroup,
        Parent: NewKnxGroupAddress(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastKnxGroupAddress2Level(structType interface{}) *KnxGroupAddress2Level {
    castFunc := func(typ interface{}) *KnxGroupAddress2Level {
        if casted, ok := typ.(KnxGroupAddress2Level); ok {
            return &casted
        }
        if casted, ok := typ.(*KnxGroupAddress2Level); ok {
            return casted
        }
        if casted, ok := typ.(KnxGroupAddress); ok {
            return CastKnxGroupAddress2Level(casted.Child)
        }
        if casted, ok := typ.(*KnxGroupAddress); ok {
            return CastKnxGroupAddress2Level(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *KnxGroupAddress2Level) GetTypeName() string {
    return "KnxGroupAddress2Level"
}

func (m *KnxGroupAddress2Level) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (mainGroup)
    lengthInBits += 5

    // Simple field (subGroup)
    lengthInBits += 11

    return lengthInBits
}

func (m *KnxGroupAddress2Level) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func KnxGroupAddress2LevelParse(io *utils.ReadBuffer) (*KnxGroupAddress, error) {

    // Simple Field (mainGroup)
    mainGroup, _mainGroupErr := io.ReadUint8(5)
    if _mainGroupErr != nil {
        return nil, errors.New("Error parsing 'mainGroup' field " + _mainGroupErr.Error())
    }

    // Simple Field (subGroup)
    subGroup, _subGroupErr := io.ReadUint16(11)
    if _subGroupErr != nil {
        return nil, errors.New("Error parsing 'subGroup' field " + _subGroupErr.Error())
    }

    // Create a partially initialized instance
    _child := &KnxGroupAddress2Level{
        MainGroup: mainGroup,
        SubGroup: subGroup,
        Parent: &KnxGroupAddress{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *KnxGroupAddress2Level) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (mainGroup)
    mainGroup := uint8(m.MainGroup)
    _mainGroupErr := io.WriteUint8(5, (mainGroup))
    if _mainGroupErr != nil {
        return errors.New("Error serializing 'mainGroup' field " + _mainGroupErr.Error())
    }

    // Simple Field (subGroup)
    subGroup := uint16(m.SubGroup)
    _subGroupErr := io.WriteUint16(11, (subGroup))
    if _subGroupErr != nil {
        return errors.New("Error serializing 'subGroup' field " + _subGroupErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *KnxGroupAddress2Level) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "mainGroup":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.MainGroup = data
            case "subGroup":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.SubGroup = data
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

func (m *KnxGroupAddress2Level) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.MainGroup, xml.StartElement{Name: xml.Name{Local: "mainGroup"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.SubGroup, xml.StartElement{Name: xml.Name{Local: "subGroup"}}); err != nil {
        return err
    }
    return nil
}

