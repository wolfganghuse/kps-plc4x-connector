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
type ApduDataMemoryResponse struct {
    Address uint16
    Data []uint8
    Parent *ApduData
    IApduDataMemoryResponse
}

// The corresponding interface
type IApduDataMemoryResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataMemoryResponse) ApciType() uint8 {
    return 0x9
}


func (m *ApduDataMemoryResponse) InitializeParent(parent *ApduData) {
}

func NewApduDataMemoryResponse(address uint16, data []uint8, ) *ApduData {
    child := &ApduDataMemoryResponse{
        Address: address,
        Data: data,
        Parent: NewApduData(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataMemoryResponse(structType interface{}) *ApduDataMemoryResponse {
    castFunc := func(typ interface{}) *ApduDataMemoryResponse {
        if casted, ok := typ.(ApduDataMemoryResponse); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataMemoryResponse); ok {
            return casted
        }
        if casted, ok := typ.(ApduData); ok {
            return CastApduDataMemoryResponse(casted.Child)
        }
        if casted, ok := typ.(*ApduData); ok {
            return CastApduDataMemoryResponse(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataMemoryResponse) GetTypeName() string {
    return "ApduDataMemoryResponse"
}

func (m *ApduDataMemoryResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Implicit Field (numBytes)
    lengthInBits += 6

    // Simple field (address)
    lengthInBits += 16

    // Array field
    if len(m.Data) > 0 {
        lengthInBits += 8 * uint16(len(m.Data))
    }

    return lengthInBits
}

func (m *ApduDataMemoryResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataMemoryResponseParse(io *utils.ReadBuffer) (*ApduData, error) {

    // Implicit Field (numBytes) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    numBytes, _numBytesErr := io.ReadUint8(6)
    if _numBytesErr != nil {
        return nil, errors.New("Error parsing 'numBytes' field " + _numBytesErr.Error())
    }

    // Simple Field (address)
    address, _addressErr := io.ReadUint16(16)
    if _addressErr != nil {
        return nil, errors.New("Error parsing 'address' field " + _addressErr.Error())
    }

    // Array field (data)
    // Count array
    data := make([]uint8, numBytes)
    for curItem := uint16(0); curItem < uint16(numBytes); curItem++ {
        _item, _err := io.ReadUint8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'data' field " + _err.Error())
        }
        data[curItem] = _item
    }

    // Create a partially initialized instance
    _child := &ApduDataMemoryResponse{
        Address: address,
        Data: data,
        Parent: &ApduData{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataMemoryResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Implicit Field (numBytes) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    numBytes := uint8(uint8(len(m.Data)))
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

    // Array Field (data)
    if m.Data != nil {
        for _, _element := range m.Data {
            _elementErr := io.WriteUint8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'data' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataMemoryResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "address":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Address = data
            case "data":
                var data []uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Data = data
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

func (m *ApduDataMemoryResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    return nil
}

