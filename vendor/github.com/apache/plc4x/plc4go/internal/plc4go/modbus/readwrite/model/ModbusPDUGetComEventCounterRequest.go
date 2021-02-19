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
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
)

// The data-structure of this message
type ModbusPDUGetComEventCounterRequest struct {
    Parent *ModbusPDU
    IModbusPDUGetComEventCounterRequest
}

// The corresponding interface
type IModbusPDUGetComEventCounterRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUGetComEventCounterRequest) ErrorFlag() bool {
    return false
}

func (m *ModbusPDUGetComEventCounterRequest) FunctionFlag() uint8 {
    return 0x0B
}

func (m *ModbusPDUGetComEventCounterRequest) Response() bool {
    return false
}


func (m *ModbusPDUGetComEventCounterRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUGetComEventCounterRequest() *ModbusPDU {
    child := &ModbusPDUGetComEventCounterRequest{
        Parent: NewModbusPDU(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastModbusPDUGetComEventCounterRequest(structType interface{}) *ModbusPDUGetComEventCounterRequest {
    castFunc := func(typ interface{}) *ModbusPDUGetComEventCounterRequest {
        if casted, ok := typ.(ModbusPDUGetComEventCounterRequest); ok {
            return &casted
        }
        if casted, ok := typ.(*ModbusPDUGetComEventCounterRequest); ok {
            return casted
        }
        if casted, ok := typ.(ModbusPDU); ok {
            return CastModbusPDUGetComEventCounterRequest(casted.Child)
        }
        if casted, ok := typ.(*ModbusPDU); ok {
            return CastModbusPDUGetComEventCounterRequest(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ModbusPDUGetComEventCounterRequest) GetTypeName() string {
    return "ModbusPDUGetComEventCounterRequest"
}

func (m *ModbusPDUGetComEventCounterRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *ModbusPDUGetComEventCounterRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUGetComEventCounterRequestParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

    // Create a partially initialized instance
    _child := &ModbusPDUGetComEventCounterRequest{
        Parent: &ModbusPDU{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ModbusPDUGetComEventCounterRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUGetComEventCounterRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
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

func (m *ModbusPDUGetComEventCounterRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

