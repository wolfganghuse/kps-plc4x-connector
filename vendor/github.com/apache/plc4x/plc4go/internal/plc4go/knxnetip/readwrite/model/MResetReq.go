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
type MResetReq struct {
    Parent *CEMI
    IMResetReq
}

// The corresponding interface
type IMResetReq interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *MResetReq) MessageCode() uint8 {
    return 0xF1
}


func (m *MResetReq) InitializeParent(parent *CEMI) {
}

func NewMResetReq() *CEMI {
    child := &MResetReq{
        Parent: NewCEMI(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastMResetReq(structType interface{}) *MResetReq {
    castFunc := func(typ interface{}) *MResetReq {
        if casted, ok := typ.(MResetReq); ok {
            return &casted
        }
        if casted, ok := typ.(*MResetReq); ok {
            return casted
        }
        if casted, ok := typ.(CEMI); ok {
            return CastMResetReq(casted.Child)
        }
        if casted, ok := typ.(*CEMI); ok {
            return CastMResetReq(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *MResetReq) GetTypeName() string {
    return "MResetReq"
}

func (m *MResetReq) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *MResetReq) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func MResetReqParse(io *utils.ReadBuffer) (*CEMI, error) {

    // Create a partially initialized instance
    _child := &MResetReq{
        Parent: &CEMI{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *MResetReq) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *MResetReq) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *MResetReq) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

