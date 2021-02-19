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
type ApduDataMemoryRead struct {
    NumBytes uint8
    Address uint16
    Parent *ApduData
    IApduDataMemoryRead
}

// The corresponding interface
type IApduDataMemoryRead interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataMemoryRead) ApciType() uint8 {
    return 0x8
}


func (m *ApduDataMemoryRead) InitializeParent(parent *ApduData) {
}

func NewApduDataMemoryRead(numBytes uint8, address uint16, ) *ApduData {
    child := &ApduDataMemoryRead{
        NumBytes: numBytes,
        Address: address,
        Parent: NewApduData(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataMemoryRead(structType interface{}) *ApduDataMemoryRead {
    castFunc := func(typ interface{}) *ApduDataMemoryRead {
        if casted, ok := typ.(ApduDataMemoryRead); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataMemoryRead); ok {
            return casted
        }
        if casted, ok := typ.(ApduData); ok {
            return CastApduDataMemoryRead(casted.Child)
        }
        if casted, ok := typ.(*ApduData); ok {
            return CastApduDataMemoryRead(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataMemoryRead) GetTypeName() string {
    return "ApduDataMemoryRead"
}

func (m *ApduDataMemoryRead) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (numBytes)
    lengthInBits += 6

    // Simple field (address)
    lengthInBits += 16

    return lengthInBits
}

func (m *ApduDataMemoryRead) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataMemoryReadParse(io *utils.ReadBuffer) (*ApduData, error) {

    // Simple Field (numBytes)
    numBytes, _numBytesErr := io.ReadUint8(6)
    if _numBytesErr != nil {
        return nil, errors.New("Error parsing 'numBytes' field " + _numBytesErr.Error())
    }

    // Simple Field (address)
    address, _addressErr := io.ReadUint16(16)
    if _addressErr != nil {
        return nil, errors.New("Error parsing 'address' field " + _addressErr.Error())
    }

    // Create a partially initialized instance
    _child := &ApduDataMemoryRead{
        NumBytes: numBytes,
        Address: address,
        Parent: &ApduData{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataMemoryRead) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (numBytes)
    numBytes := uint8(m.NumBytes)
    _numBytesErr := io.WriteUint8(6, (numBytes))
    if _numBytesErr != nil {
        return errors.New("Error serializing 'numBytes' field " + _numBytesErr.Error())
    }

    // Simple Field (address)
    address := uint16(m.Address)
    _addressErr := io.WriteUint16(16, (address))
    if _addressErr != nil {
        return errors.New("Error serializing 'address' field " + _addressErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataMemoryRead) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "numBytes":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.NumBytes = data
            case "address":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Address = data
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

func (m *ApduDataMemoryRead) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.NumBytes, xml.StartElement{Name: xml.Name{Local: "numBytes"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); err != nil {
        return err
    }
    return nil
}

