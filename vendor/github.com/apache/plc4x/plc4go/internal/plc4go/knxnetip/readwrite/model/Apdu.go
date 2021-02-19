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
type Apdu struct {
    Numbered bool
    Counter uint8
    Child IApduChild
    IApdu
    IApduParent
}

// The corresponding interface
type IApdu interface {
    Control() uint8
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type IApduParent interface {
    SerializeParent(io utils.WriteBuffer, child IApdu, serializeChildFunction func() error) error
    GetTypeName() string
}

type IApduChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *Apdu, numbered bool, counter uint8)
    GetTypeName() string
    IApdu
}

func NewApdu(numbered bool, counter uint8) *Apdu {
    return &Apdu{Numbered: numbered, Counter: counter}
}

func CastApdu(structType interface{}) *Apdu {
    castFunc := func(typ interface{}) *Apdu {
        if casted, ok := typ.(Apdu); ok {
            return &casted
        }
        if casted, ok := typ.(*Apdu); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *Apdu) GetTypeName() string {
    return "Apdu"
}

func (m *Apdu) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Discriminator Field (control)
    lengthInBits += 1

    // Simple field (numbered)
    lengthInBits += 1

    // Simple field (counter)
    lengthInBits += 4

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *Apdu) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduParse(io *utils.ReadBuffer, dataLength uint8) (*Apdu, error) {

    // Discriminator Field (control) (Used as input to a switch field)
    control, _controlErr := io.ReadUint8(1)
    if _controlErr != nil {
        return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
    }

    // Simple Field (numbered)
    numbered, _numberedErr := io.ReadBit()
    if _numberedErr != nil {
        return nil, errors.New("Error parsing 'numbered' field " + _numberedErr.Error())
    }

    // Simple Field (counter)
    counter, _counterErr := io.ReadUint8(4)
    if _counterErr != nil {
        return nil, errors.New("Error parsing 'counter' field " + _counterErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *Apdu
    var typeSwitchError error
    switch {
    case control == 1:
        _parent, typeSwitchError = ApduControlContainerParse(io)
    case control == 0:
        _parent, typeSwitchError = ApduDataContainerParse(io, dataLength)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent, numbered, counter)
    return _parent, nil
}

func (m *Apdu) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *Apdu) SerializeParent(io utils.WriteBuffer, child IApdu, serializeChildFunction func() error) error {

    // Discriminator Field (control) (Used as input to a switch field)
    control := uint8(child.Control())
    _controlErr := io.WriteUint8(1, (control))
    if _controlErr != nil {
        return errors.New("Error serializing 'control' field " + _controlErr.Error())
    }

    // Simple Field (numbered)
    numbered := bool(m.Numbered)
    _numberedErr := io.WriteBit((numbered))
    if _numberedErr != nil {
        return errors.New("Error serializing 'numbered' field " + _numberedErr.Error())
    }

    // Simple Field (counter)
    counter := uint8(m.Counter)
    _counterErr := io.WriteUint8(4, (counter))
    if _counterErr != nil {
        return errors.New("Error serializing 'counter' field " + _counterErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *Apdu) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "numbered":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Numbered = data
            case "counter":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Counter = data
            default:
                switch start.Attr[0].Value {
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlContainer":
                        var dt *ApduControlContainer
                        if m.Child != nil {
                            dt = m.Child.(*ApduControlContainer)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataContainer":
                        var dt *ApduDataContainer
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataContainer)
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

func (m *Apdu) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := reflect.TypeOf(m.Child).String()
    className = "org.apache.plc4x.java.knxnetip.readwrite." + className[strings.LastIndex(className, ".") + 1:]
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Numbered, xml.StartElement{Name: xml.Name{Local: "numbered"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Counter, xml.StartElement{Name: xml.Name{Local: "counter"}}); err != nil {
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

