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
type ApduDataExtLinkResponse struct {
    Parent *ApduDataExt
    IApduDataExtLinkResponse
}

// The corresponding interface
type IApduDataExtLinkResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtLinkResponse) ExtApciType() uint8 {
    return 0x26
}


func (m *ApduDataExtLinkResponse) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtLinkResponse() *ApduDataExt {
    child := &ApduDataExtLinkResponse{
        Parent: NewApduDataExt(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataExtLinkResponse(structType interface{}) *ApduDataExtLinkResponse {
    castFunc := func(typ interface{}) *ApduDataExtLinkResponse {
        if casted, ok := typ.(ApduDataExtLinkResponse); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataExtLinkResponse); ok {
            return casted
        }
        if casted, ok := typ.(ApduDataExt); ok {
            return CastApduDataExtLinkResponse(casted.Child)
        }
        if casted, ok := typ.(*ApduDataExt); ok {
            return CastApduDataExtLinkResponse(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataExtLinkResponse) GetTypeName() string {
    return "ApduDataExtLinkResponse"
}

func (m *ApduDataExtLinkResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *ApduDataExtLinkResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataExtLinkResponseParse(io *utils.ReadBuffer) (*ApduDataExt, error) {

    // Create a partially initialized instance
    _child := &ApduDataExtLinkResponse{
        Parent: &ApduDataExt{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataExtLinkResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtLinkResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *ApduDataExtLinkResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

