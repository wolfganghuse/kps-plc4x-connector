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
type ApduDataExtDomainAddressSerialNumberWrite struct {
    Parent *ApduDataExt
    IApduDataExtDomainAddressSerialNumberWrite
}

// The corresponding interface
type IApduDataExtDomainAddressSerialNumberWrite interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtDomainAddressSerialNumberWrite) ExtApciType() uint8 {
    return 0x2E
}


func (m *ApduDataExtDomainAddressSerialNumberWrite) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtDomainAddressSerialNumberWrite() *ApduDataExt {
    child := &ApduDataExtDomainAddressSerialNumberWrite{
        Parent: NewApduDataExt(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataExtDomainAddressSerialNumberWrite(structType interface{}) *ApduDataExtDomainAddressSerialNumberWrite {
    castFunc := func(typ interface{}) *ApduDataExtDomainAddressSerialNumberWrite {
        if casted, ok := typ.(ApduDataExtDomainAddressSerialNumberWrite); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataExtDomainAddressSerialNumberWrite); ok {
            return casted
        }
        if casted, ok := typ.(ApduDataExt); ok {
            return CastApduDataExtDomainAddressSerialNumberWrite(casted.Child)
        }
        if casted, ok := typ.(*ApduDataExt); ok {
            return CastApduDataExtDomainAddressSerialNumberWrite(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataExtDomainAddressSerialNumberWrite) GetTypeName() string {
    return "ApduDataExtDomainAddressSerialNumberWrite"
}

func (m *ApduDataExtDomainAddressSerialNumberWrite) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *ApduDataExtDomainAddressSerialNumberWrite) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataExtDomainAddressSerialNumberWriteParse(io *utils.ReadBuffer) (*ApduDataExt, error) {

    // Create a partially initialized instance
    _child := &ApduDataExtDomainAddressSerialNumberWrite{
        Parent: &ApduDataExt{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataExtDomainAddressSerialNumberWrite) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtDomainAddressSerialNumberWrite) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *ApduDataExtDomainAddressSerialNumberWrite) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

