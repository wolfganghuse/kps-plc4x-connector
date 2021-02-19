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
type ModbusPDUWriteSingleRegisterRequest struct {
    Address uint16
    Value uint16
    Parent *ModbusPDU
    IModbusPDUWriteSingleRegisterRequest
}

// The corresponding interface
type IModbusPDUWriteSingleRegisterRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUWriteSingleRegisterRequest) ErrorFlag() bool {
    return false
}

func (m *ModbusPDUWriteSingleRegisterRequest) FunctionFlag() uint8 {
    return 0x06
}

func (m *ModbusPDUWriteSingleRegisterRequest) Response() bool {
    return false
}


func (m *ModbusPDUWriteSingleRegisterRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUWriteSingleRegisterRequest(address uint16, value uint16, ) *ModbusPDU {
    child := &ModbusPDUWriteSingleRegisterRequest{
        Address: address,
        Value: value,
        Parent: NewModbusPDU(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastModbusPDUWriteSingleRegisterRequest(structType interface{}) *ModbusPDUWriteSingleRegisterRequest {
    castFunc := func(typ interface{}) *ModbusPDUWriteSingleRegisterRequest {
        if casted, ok := typ.(ModbusPDUWriteSingleRegisterRequest); ok {
            return &casted
        }
        if casted, ok := typ.(*ModbusPDUWriteSingleRegisterRequest); ok {
            return casted
        }
        if casted, ok := typ.(ModbusPDU); ok {
            return CastModbusPDUWriteSingleRegisterRequest(casted.Child)
        }
        if casted, ok := typ.(*ModbusPDU); ok {
            return CastModbusPDUWriteSingleRegisterRequest(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ModbusPDUWriteSingleRegisterRequest) GetTypeName() string {
    return "ModbusPDUWriteSingleRegisterRequest"
}

func (m *ModbusPDUWriteSingleRegisterRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (address)
    lengthInBits += 16

    // Simple field (value)
    lengthInBits += 16

    return lengthInBits
}

func (m *ModbusPDUWriteSingleRegisterRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUWriteSingleRegisterRequestParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

    // Simple Field (address)
    address, _addressErr := io.ReadUint16(16)
    if _addressErr != nil {
        return nil, errors.New("Error parsing 'address' field " + _addressErr.Error())
    }

    // Simple Field (value)
    value, _valueErr := io.ReadUint16(16)
    if _valueErr != nil {
        return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
    }

    // Create a partially initialized instance
    _child := &ModbusPDUWriteSingleRegisterRequest{
        Address: address,
        Value: value,
        Parent: &ModbusPDU{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ModbusPDUWriteSingleRegisterRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (address)
    address := uint16(m.Address)
    _addressErr := io.WriteUint16(16, (address))
    if _addressErr != nil {
        return errors.New("Error serializing 'address' field " + _addressErr.Error())
    }

    // Simple Field (value)
    value := uint16(m.Value)
    _valueErr := io.WriteUint16(16, (value))
    if _valueErr != nil {
        return errors.New("Error serializing 'value' field " + _valueErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUWriteSingleRegisterRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "value":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Value = data
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

func (m *ModbusPDUWriteSingleRegisterRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Value, xml.StartElement{Name: xml.Name{Local: "value"}}); err != nil {
        return err
    }
    return nil
}

