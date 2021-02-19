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
type ModbusPDUMaskWriteHoldingRegisterRequest struct {
    ReferenceAddress uint16
    AndMask uint16
    OrMask uint16
    Parent *ModbusPDU
    IModbusPDUMaskWriteHoldingRegisterRequest
}

// The corresponding interface
type IModbusPDUMaskWriteHoldingRegisterRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUMaskWriteHoldingRegisterRequest) ErrorFlag() bool {
    return false
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) FunctionFlag() uint8 {
    return 0x16
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) Response() bool {
    return false
}


func (m *ModbusPDUMaskWriteHoldingRegisterRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUMaskWriteHoldingRegisterRequest(referenceAddress uint16, andMask uint16, orMask uint16, ) *ModbusPDU {
    child := &ModbusPDUMaskWriteHoldingRegisterRequest{
        ReferenceAddress: referenceAddress,
        AndMask: andMask,
        OrMask: orMask,
        Parent: NewModbusPDU(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastModbusPDUMaskWriteHoldingRegisterRequest(structType interface{}) *ModbusPDUMaskWriteHoldingRegisterRequest {
    castFunc := func(typ interface{}) *ModbusPDUMaskWriteHoldingRegisterRequest {
        if casted, ok := typ.(ModbusPDUMaskWriteHoldingRegisterRequest); ok {
            return &casted
        }
        if casted, ok := typ.(*ModbusPDUMaskWriteHoldingRegisterRequest); ok {
            return casted
        }
        if casted, ok := typ.(ModbusPDU); ok {
            return CastModbusPDUMaskWriteHoldingRegisterRequest(casted.Child)
        }
        if casted, ok := typ.(*ModbusPDU); ok {
            return CastModbusPDUMaskWriteHoldingRegisterRequest(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) GetTypeName() string {
    return "ModbusPDUMaskWriteHoldingRegisterRequest"
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (referenceAddress)
    lengthInBits += 16

    // Simple field (andMask)
    lengthInBits += 16

    // Simple field (orMask)
    lengthInBits += 16

    return lengthInBits
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUMaskWriteHoldingRegisterRequestParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

    // Simple Field (referenceAddress)
    referenceAddress, _referenceAddressErr := io.ReadUint16(16)
    if _referenceAddressErr != nil {
        return nil, errors.New("Error parsing 'referenceAddress' field " + _referenceAddressErr.Error())
    }

    // Simple Field (andMask)
    andMask, _andMaskErr := io.ReadUint16(16)
    if _andMaskErr != nil {
        return nil, errors.New("Error parsing 'andMask' field " + _andMaskErr.Error())
    }

    // Simple Field (orMask)
    orMask, _orMaskErr := io.ReadUint16(16)
    if _orMaskErr != nil {
        return nil, errors.New("Error parsing 'orMask' field " + _orMaskErr.Error())
    }

    // Create a partially initialized instance
    _child := &ModbusPDUMaskWriteHoldingRegisterRequest{
        ReferenceAddress: referenceAddress,
        AndMask: andMask,
        OrMask: orMask,
        Parent: &ModbusPDU{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (referenceAddress)
    referenceAddress := uint16(m.ReferenceAddress)
    _referenceAddressErr := io.WriteUint16(16, (referenceAddress))
    if _referenceAddressErr != nil {
        return errors.New("Error serializing 'referenceAddress' field " + _referenceAddressErr.Error())
    }

    // Simple Field (andMask)
    andMask := uint16(m.AndMask)
    _andMaskErr := io.WriteUint16(16, (andMask))
    if _andMaskErr != nil {
        return errors.New("Error serializing 'andMask' field " + _andMaskErr.Error())
    }

    // Simple Field (orMask)
    orMask := uint16(m.OrMask)
    _orMaskErr := io.WriteUint16(16, (orMask))
    if _orMaskErr != nil {
        return errors.New("Error serializing 'orMask' field " + _orMaskErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "referenceAddress":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ReferenceAddress = data
            case "andMask":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.AndMask = data
            case "orMask":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.OrMask = data
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

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.ReferenceAddress, xml.StartElement{Name: xml.Name{Local: "referenceAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.AndMask, xml.StartElement{Name: xml.Name{Local: "andMask"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.OrMask, xml.StartElement{Name: xml.Name{Local: "orMask"}}); err != nil {
        return err
    }
    return nil
}

