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
type TDataIndividualReq struct {
    Parent *CEMI
    ITDataIndividualReq
}

// The corresponding interface
type ITDataIndividualReq interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *TDataIndividualReq) MessageCode() uint8 {
    return 0x4A
}


func (m *TDataIndividualReq) InitializeParent(parent *CEMI) {
}

func NewTDataIndividualReq() *CEMI {
    child := &TDataIndividualReq{
        Parent: NewCEMI(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastTDataIndividualReq(structType interface{}) *TDataIndividualReq {
    castFunc := func(typ interface{}) *TDataIndividualReq {
        if casted, ok := typ.(TDataIndividualReq); ok {
            return &casted
        }
        if casted, ok := typ.(*TDataIndividualReq); ok {
            return casted
        }
        if casted, ok := typ.(CEMI); ok {
            return CastTDataIndividualReq(casted.Child)
        }
        if casted, ok := typ.(*CEMI); ok {
            return CastTDataIndividualReq(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *TDataIndividualReq) GetTypeName() string {
    return "TDataIndividualReq"
}

func (m *TDataIndividualReq) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *TDataIndividualReq) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func TDataIndividualReqParse(io *utils.ReadBuffer) (*CEMI, error) {

    // Create a partially initialized instance
    _child := &TDataIndividualReq{
        Parent: &CEMI{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *TDataIndividualReq) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *TDataIndividualReq) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *TDataIndividualReq) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

