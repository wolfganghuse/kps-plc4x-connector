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
type ApduDataContainer struct {
    DataApdu *ApduData
    Parent *Apdu
    IApduDataContainer
}

// The corresponding interface
type IApduDataContainer interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataContainer) Control() uint8 {
    return 0
}


func (m *ApduDataContainer) InitializeParent(parent *Apdu, numbered bool, counter uint8) {
    m.Parent.Numbered = numbered
    m.Parent.Counter = counter
}

func NewApduDataContainer(dataApdu *ApduData, numbered bool, counter uint8) *Apdu {
    child := &ApduDataContainer{
        DataApdu: dataApdu,
        Parent: NewApdu(numbered, counter),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataContainer(structType interface{}) *ApduDataContainer {
    castFunc := func(typ interface{}) *ApduDataContainer {
        if casted, ok := typ.(ApduDataContainer); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataContainer); ok {
            return casted
        }
        if casted, ok := typ.(Apdu); ok {
            return CastApduDataContainer(casted.Child)
        }
        if casted, ok := typ.(*Apdu); ok {
            return CastApduDataContainer(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataContainer) GetTypeName() string {
    return "ApduDataContainer"
}

func (m *ApduDataContainer) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (dataApdu)
    lengthInBits += m.DataApdu.LengthInBits()

    return lengthInBits
}

func (m *ApduDataContainer) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataContainerParse(io *utils.ReadBuffer, dataLength uint8) (*Apdu, error) {

    // Simple Field (dataApdu)
    dataApdu, _dataApduErr := ApduDataParse(io, dataLength)
    if _dataApduErr != nil {
        return nil, errors.New("Error parsing 'dataApdu' field " + _dataApduErr.Error())
    }

    // Create a partially initialized instance
    _child := &ApduDataContainer{
        DataApdu: dataApdu,
        Parent: &Apdu{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataContainer) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (dataApdu)
    _dataApduErr := m.DataApdu.Serialize(io)
    if _dataApduErr != nil {
        return errors.New("Error serializing 'dataApdu' field " + _dataApduErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataContainer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "dataApdu":
                var dt *ApduData
                if err := d.DecodeElement(&dt, &tok); err != nil {
                    return err
                }
                m.DataApdu = dt
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

func (m *ApduDataContainer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.DataApdu, xml.StartElement{Name: xml.Name{Local: "dataApdu"}}); err != nil {
        return err
    }
    return nil
}

