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
type ApduControl struct {
    Child IApduControlChild
    IApduControl
    IApduControlParent
}

// The corresponding interface
type IApduControl interface {
    ControlType() uint8
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type IApduControlParent interface {
    SerializeParent(io utils.WriteBuffer, child IApduControl, serializeChildFunction func() error) error
    GetTypeName() string
}

type IApduControlChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *ApduControl)
    GetTypeName() string
    IApduControl
}

func NewApduControl() *ApduControl {
    return &ApduControl{}
}

func CastApduControl(structType interface{}) *ApduControl {
    castFunc := func(typ interface{}) *ApduControl {
        if casted, ok := typ.(ApduControl); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduControl); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduControl) GetTypeName() string {
    return "ApduControl"
}

func (m *ApduControl) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Discriminator Field (controlType)
    lengthInBits += 2

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *ApduControl) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduControlParse(io *utils.ReadBuffer) (*ApduControl, error) {

    // Discriminator Field (controlType) (Used as input to a switch field)
    controlType, _controlTypeErr := io.ReadUint8(2)
    if _controlTypeErr != nil {
        return nil, errors.New("Error parsing 'controlType' field " + _controlTypeErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *ApduControl
    var typeSwitchError error
    switch {
    case controlType == 0x0:
        _parent, typeSwitchError = ApduControlConnectParse(io)
    case controlType == 0x1:
        _parent, typeSwitchError = ApduControlDisconnectParse(io)
    case controlType == 0x2:
        _parent, typeSwitchError = ApduControlAckParse(io)
    case controlType == 0x3:
        _parent, typeSwitchError = ApduControlNackParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent)
    return _parent, nil
}

func (m *ApduControl) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *ApduControl) SerializeParent(io utils.WriteBuffer, child IApduControl, serializeChildFunction func() error) error {

    // Discriminator Field (controlType) (Used as input to a switch field)
    controlType := uint8(child.ControlType())
    _controlTypeErr := io.WriteUint8(2, (controlType))
    if _controlTypeErr != nil {
        return errors.New("Error serializing 'controlType' field " + _controlTypeErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *ApduControl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlConnect":
                        var dt *ApduControlConnect
                        if m.Child != nil {
                            dt = m.Child.(*ApduControlConnect)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlDisconnect":
                        var dt *ApduControlDisconnect
                        if m.Child != nil {
                            dt = m.Child.(*ApduControlDisconnect)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlAck":
                        var dt *ApduControlAck
                        if m.Child != nil {
                            dt = m.Child.(*ApduControlAck)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlNack":
                        var dt *ApduControlNack
                        if m.Child != nil {
                            dt = m.Child.(*ApduControlNack)
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

func (m *ApduControl) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

