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
    "reflect"
    "strings"
)

// The data-structure of this message
type CEMIAdditionalInformation struct {
    Child ICEMIAdditionalInformationChild
    ICEMIAdditionalInformation
    ICEMIAdditionalInformationParent
}

// The corresponding interface
type ICEMIAdditionalInformation interface {
    AdditionalInformationType() uint8
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type ICEMIAdditionalInformationParent interface {
    SerializeParent(io utils.WriteBuffer, child ICEMIAdditionalInformation, serializeChildFunction func() error) error
    GetTypeName() string
}

type ICEMIAdditionalInformationChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *CEMIAdditionalInformation)
    GetTypeName() string
    ICEMIAdditionalInformation
}

func NewCEMIAdditionalInformation() *CEMIAdditionalInformation {
    return &CEMIAdditionalInformation{}
}

func CastCEMIAdditionalInformation(structType interface{}) *CEMIAdditionalInformation {
    castFunc := func(typ interface{}) *CEMIAdditionalInformation {
        if casted, ok := typ.(CEMIAdditionalInformation); ok {
            return &casted
        }
        if casted, ok := typ.(*CEMIAdditionalInformation); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *CEMIAdditionalInformation) GetTypeName() string {
    return "CEMIAdditionalInformation"
}

func (m *CEMIAdditionalInformation) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Discriminator Field (additionalInformationType)
    lengthInBits += 8

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *CEMIAdditionalInformation) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func CEMIAdditionalInformationParse(io *utils.ReadBuffer) (*CEMIAdditionalInformation, error) {

    // Discriminator Field (additionalInformationType) (Used as input to a switch field)
    additionalInformationType, _additionalInformationTypeErr := io.ReadUint8(8)
    if _additionalInformationTypeErr != nil {
        return nil, errors.New("Error parsing 'additionalInformationType' field " + _additionalInformationTypeErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *CEMIAdditionalInformation
    var typeSwitchError error
    switch {
    case additionalInformationType == 0x03:
        _parent, typeSwitchError = CEMIAdditionalInformationBusmonitorInfoParse(io)
    case additionalInformationType == 0x04:
        _parent, typeSwitchError = CEMIAdditionalInformationRelativeTimestampParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent)
    return _parent, nil
}

func (m *CEMIAdditionalInformation) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *CEMIAdditionalInformation) SerializeParent(io utils.WriteBuffer, child ICEMIAdditionalInformation, serializeChildFunction func() error) error {

    // Discriminator Field (additionalInformationType) (Used as input to a switch field)
    additionalInformationType := uint8(child.AdditionalInformationType())
    _additionalInformationTypeErr := io.WriteUint8(8, (additionalInformationType))
    if _additionalInformationTypeErr != nil {
        return errors.New("Error serializing 'additionalInformationType' field " + _additionalInformationTypeErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *CEMIAdditionalInformation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    for {
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            default:
                switch start.Attr[0].Value {
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIAdditionalInformationBusmonitorInfo":
                        var dt *CEMIAdditionalInformationBusmonitorInfo
                        if m.Child != nil {
                            dt = m.Child.(*CEMIAdditionalInformationBusmonitorInfo)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIAdditionalInformationRelativeTimestamp":
                        var dt *CEMIAdditionalInformationRelativeTimestamp
                        if m.Child != nil {
                            dt = m.Child.(*CEMIAdditionalInformationRelativeTimestamp)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                }
            }
        }
    }
}

func (m *CEMIAdditionalInformation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := reflect.TypeOf(m.Child).String()
    className = "org.apache.plc4x.java.knxnetip.readwrite." + className[strings.LastIndex(className, ".") + 1:]
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    marshaller, ok := m.Child.(xml.Marshaler)
    if !ok {
        return errors.New("child is not castable to Marshaler")
    }
    marshaller.MarshalXML(e, start)
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

