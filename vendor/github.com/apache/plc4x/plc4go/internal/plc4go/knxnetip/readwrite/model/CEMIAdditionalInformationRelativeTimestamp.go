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
    "strconv"
)

// Constant values.
const CEMIAdditionalInformationRelativeTimestamp_LEN uint8 = 2

// The data-structure of this message
type CEMIAdditionalInformationRelativeTimestamp struct {
    RelativeTimestamp *RelativeTimestamp
    Parent *CEMIAdditionalInformation
    ICEMIAdditionalInformationRelativeTimestamp
}

// The corresponding interface
type ICEMIAdditionalInformationRelativeTimestamp interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *CEMIAdditionalInformationRelativeTimestamp) AdditionalInformationType() uint8 {
    return 0x04
}


func (m *CEMIAdditionalInformationRelativeTimestamp) InitializeParent(parent *CEMIAdditionalInformation) {
}

func NewCEMIAdditionalInformationRelativeTimestamp(relativeTimestamp *RelativeTimestamp, ) *CEMIAdditionalInformation {
    child := &CEMIAdditionalInformationRelativeTimestamp{
        RelativeTimestamp: relativeTimestamp,
        Parent: NewCEMIAdditionalInformation(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastCEMIAdditionalInformationRelativeTimestamp(structType interface{}) *CEMIAdditionalInformationRelativeTimestamp {
    castFunc := func(typ interface{}) *CEMIAdditionalInformationRelativeTimestamp {
        if casted, ok := typ.(CEMIAdditionalInformationRelativeTimestamp); ok {
            return &casted
        }
        if casted, ok := typ.(*CEMIAdditionalInformationRelativeTimestamp); ok {
            return casted
        }
        if casted, ok := typ.(CEMIAdditionalInformation); ok {
            return CastCEMIAdditionalInformationRelativeTimestamp(casted.Child)
        }
        if casted, ok := typ.(*CEMIAdditionalInformation); ok {
            return CastCEMIAdditionalInformationRelativeTimestamp(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *CEMIAdditionalInformationRelativeTimestamp) GetTypeName() string {
    return "CEMIAdditionalInformationRelativeTimestamp"
}

func (m *CEMIAdditionalInformationRelativeTimestamp) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Const Field (len)
    lengthInBits += 8

    // Simple field (relativeTimestamp)
    lengthInBits += m.RelativeTimestamp.LengthInBits()

    return lengthInBits
}

func (m *CEMIAdditionalInformationRelativeTimestamp) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func CEMIAdditionalInformationRelativeTimestampParse(io *utils.ReadBuffer) (*CEMIAdditionalInformation, error) {

    // Const Field (len)
    len, _lenErr := io.ReadUint8(8)
    if _lenErr != nil {
        return nil, errors.New("Error parsing 'len' field " + _lenErr.Error())
    }
    if len != CEMIAdditionalInformationRelativeTimestamp_LEN {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(CEMIAdditionalInformationRelativeTimestamp_LEN)) + " but got " + strconv.Itoa(int(len)))
    }

    // Simple Field (relativeTimestamp)
    relativeTimestamp, _relativeTimestampErr := RelativeTimestampParse(io)
    if _relativeTimestampErr != nil {
        return nil, errors.New("Error parsing 'relativeTimestamp' field " + _relativeTimestampErr.Error())
    }

    // Create a partially initialized instance
    _child := &CEMIAdditionalInformationRelativeTimestamp{
        RelativeTimestamp: relativeTimestamp,
        Parent: &CEMIAdditionalInformation{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *CEMIAdditionalInformationRelativeTimestamp) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Const Field (len)
    _lenErr := io.WriteUint8(8, 2)
    if _lenErr != nil {
        return errors.New("Error serializing 'len' field " + _lenErr.Error())
    }

    // Simple Field (relativeTimestamp)
    _relativeTimestampErr := m.RelativeTimestamp.Serialize(io)
    if _relativeTimestampErr != nil {
        return errors.New("Error serializing 'relativeTimestamp' field " + _relativeTimestampErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *CEMIAdditionalInformationRelativeTimestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "relativeTimestamp":
                var data *RelativeTimestamp
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.RelativeTimestamp = data
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

func (m *CEMIAdditionalInformationRelativeTimestamp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.RelativeTimestamp, xml.StartElement{Name: xml.Name{Local: "relativeTimestamp"}}); err != nil {
        return err
    }
    return nil
}

