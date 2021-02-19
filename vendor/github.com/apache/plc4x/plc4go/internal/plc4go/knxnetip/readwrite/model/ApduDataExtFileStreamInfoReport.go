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
type ApduDataExtFileStreamInfoReport struct {
    Parent *ApduDataExt
    IApduDataExtFileStreamInfoReport
}

// The corresponding interface
type IApduDataExtFileStreamInfoReport interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtFileStreamInfoReport) ExtApciType() uint8 {
    return 0x30
}


func (m *ApduDataExtFileStreamInfoReport) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtFileStreamInfoReport() *ApduDataExt {
    child := &ApduDataExtFileStreamInfoReport{
        Parent: NewApduDataExt(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataExtFileStreamInfoReport(structType interface{}) *ApduDataExtFileStreamInfoReport {
    castFunc := func(typ interface{}) *ApduDataExtFileStreamInfoReport {
        if casted, ok := typ.(ApduDataExtFileStreamInfoReport); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataExtFileStreamInfoReport); ok {
            return casted
        }
        if casted, ok := typ.(ApduDataExt); ok {
            return CastApduDataExtFileStreamInfoReport(casted.Child)
        }
        if casted, ok := typ.(*ApduDataExt); ok {
            return CastApduDataExtFileStreamInfoReport(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataExtFileStreamInfoReport) GetTypeName() string {
    return "ApduDataExtFileStreamInfoReport"
}

func (m *ApduDataExtFileStreamInfoReport) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *ApduDataExtFileStreamInfoReport) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataExtFileStreamInfoReportParse(io *utils.ReadBuffer) (*ApduDataExt, error) {

    // Create a partially initialized instance
    _child := &ApduDataExtFileStreamInfoReport{
        Parent: &ApduDataExt{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataExtFileStreamInfoReport) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtFileStreamInfoReport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *ApduDataExtFileStreamInfoReport) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

