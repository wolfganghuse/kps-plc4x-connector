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
type ApduDataExtMemoryBitWrite struct {
    Parent *ApduDataExt
    IApduDataExtMemoryBitWrite
}

// The corresponding interface
type IApduDataExtMemoryBitWrite interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtMemoryBitWrite) ExtApciType() uint8 {
    return 0x10
}


func (m *ApduDataExtMemoryBitWrite) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtMemoryBitWrite() *ApduDataExt {
    child := &ApduDataExtMemoryBitWrite{
        Parent: NewApduDataExt(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataExtMemoryBitWrite(structType interface{}) *ApduDataExtMemoryBitWrite {
    castFunc := func(typ interface{}) *ApduDataExtMemoryBitWrite {
        if casted, ok := typ.(ApduDataExtMemoryBitWrite); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataExtMemoryBitWrite); ok {
            return casted
        }
        if casted, ok := typ.(ApduDataExt); ok {
            return CastApduDataExtMemoryBitWrite(casted.Child)
        }
        if casted, ok := typ.(*ApduDataExt); ok {
            return CastApduDataExtMemoryBitWrite(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataExtMemoryBitWrite) GetTypeName() string {
    return "ApduDataExtMemoryBitWrite"
}

func (m *ApduDataExtMemoryBitWrite) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *ApduDataExtMemoryBitWrite) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataExtMemoryBitWriteParse(io *utils.ReadBuffer) (*ApduDataExt, error) {

    // Create a partially initialized instance
    _child := &ApduDataExtMemoryBitWrite{
        Parent: &ApduDataExt{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataExtMemoryBitWrite) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtMemoryBitWrite) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *ApduDataExtMemoryBitWrite) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

