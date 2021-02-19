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
type ApduDataOther struct {
    ExtendedApdu *ApduDataExt
    Parent *ApduData
    IApduDataOther
}

// The corresponding interface
type IApduDataOther interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataOther) ApciType() uint8 {
    return 0xF
}


func (m *ApduDataOther) InitializeParent(parent *ApduData) {
}

func NewApduDataOther(extendedApdu *ApduDataExt, ) *ApduData {
    child := &ApduDataOther{
        ExtendedApdu: extendedApdu,
        Parent: NewApduData(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataOther(structType interface{}) *ApduDataOther {
    castFunc := func(typ interface{}) *ApduDataOther {
        if casted, ok := typ.(ApduDataOther); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataOther); ok {
            return casted
        }
        if casted, ok := typ.(ApduData); ok {
            return CastApduDataOther(casted.Child)
        }
        if casted, ok := typ.(*ApduData); ok {
            return CastApduDataOther(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataOther) GetTypeName() string {
    return "ApduDataOther"
}

func (m *ApduDataOther) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (extendedApdu)
    lengthInBits += m.ExtendedApdu.LengthInBits()

    return lengthInBits
}

func (m *ApduDataOther) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataOtherParse(io *utils.ReadBuffer, dataLength uint8) (*ApduData, error) {

    // Simple Field (extendedApdu)
    extendedApdu, _extendedApduErr := ApduDataExtParse(io, dataLength)
    if _extendedApduErr != nil {
        return nil, errors.New("Error parsing 'extendedApdu' field " + _extendedApduErr.Error())
    }

    // Create a partially initialized instance
    _child := &ApduDataOther{
        ExtendedApdu: extendedApdu,
        Parent: &ApduData{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataOther) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (extendedApdu)
    _extendedApduErr := m.ExtendedApdu.Serialize(io)
    if _extendedApduErr != nil {
        return errors.New("Error serializing 'extendedApdu' field " + _extendedApduErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataOther) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "extendedApdu":
                var dt *ApduDataExt
                if err := d.DecodeElement(&dt, &tok); err != nil {
                    return err
                }
                m.ExtendedApdu = dt
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

func (m *ApduDataOther) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.ExtendedApdu, xml.StartElement{Name: xml.Name{Local: "extendedApdu"}}); err != nil {
        return err
    }
    return nil
}

