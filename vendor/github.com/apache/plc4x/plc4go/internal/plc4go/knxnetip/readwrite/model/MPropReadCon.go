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
)

// The data-structure of this message
type MPropReadCon struct {
    InterfaceObjectType uint16
    ObjectInstance uint8
    PropertyId uint8
    NumberOfElements uint8
    StartIndex uint16
    Unknown uint16
    Parent *CEMI
    IMPropReadCon
}

// The corresponding interface
type IMPropReadCon interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *MPropReadCon) MessageCode() uint8 {
    return 0xFB
}


func (m *MPropReadCon) InitializeParent(parent *CEMI) {
}

func NewMPropReadCon(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16, unknown uint16, ) *CEMI {
    child := &MPropReadCon{
        InterfaceObjectType: interfaceObjectType,
        ObjectInstance: objectInstance,
        PropertyId: propertyId,
        NumberOfElements: numberOfElements,
        StartIndex: startIndex,
        Unknown: unknown,
        Parent: NewCEMI(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastMPropReadCon(structType interface{}) *MPropReadCon {
    castFunc := func(typ interface{}) *MPropReadCon {
        if casted, ok := typ.(MPropReadCon); ok {
            return &casted
        }
        if casted, ok := typ.(*MPropReadCon); ok {
            return casted
        }
        if casted, ok := typ.(CEMI); ok {
            return CastMPropReadCon(casted.Child)
        }
        if casted, ok := typ.(*CEMI); ok {
            return CastMPropReadCon(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *MPropReadCon) GetTypeName() string {
    return "MPropReadCon"
}

func (m *MPropReadCon) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (interfaceObjectType)
    lengthInBits += 16

    // Simple field (objectInstance)
    lengthInBits += 8

    // Simple field (propertyId)
    lengthInBits += 8

    // Simple field (numberOfElements)
    lengthInBits += 4

    // Simple field (startIndex)
    lengthInBits += 12

    // Simple field (unknown)
    lengthInBits += 16

    return lengthInBits
}

func (m *MPropReadCon) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func MPropReadConParse(io *utils.ReadBuffer) (*CEMI, error) {

    // Simple Field (interfaceObjectType)
    interfaceObjectType, _interfaceObjectTypeErr := io.ReadUint16(16)
    if _interfaceObjectTypeErr != nil {
        return nil, errors.New("Error parsing 'interfaceObjectType' field " + _interfaceObjectTypeErr.Error())
    }

    // Simple Field (objectInstance)
    objectInstance, _objectInstanceErr := io.ReadUint8(8)
    if _objectInstanceErr != nil {
        return nil, errors.New("Error parsing 'objectInstance' field " + _objectInstanceErr.Error())
    }

    // Simple Field (propertyId)
    propertyId, _propertyIdErr := io.ReadUint8(8)
    if _propertyIdErr != nil {
        return nil, errors.New("Error parsing 'propertyId' field " + _propertyIdErr.Error())
    }

    // Simple Field (numberOfElements)
    numberOfElements, _numberOfElementsErr := io.ReadUint8(4)
    if _numberOfElementsErr != nil {
        return nil, errors.New("Error parsing 'numberOfElements' field " + _numberOfElementsErr.Error())
    }

    // Simple Field (startIndex)
    startIndex, _startIndexErr := io.ReadUint16(12)
    if _startIndexErr != nil {
        return nil, errors.New("Error parsing 'startIndex' field " + _startIndexErr.Error())
    }

    // Simple Field (unknown)
    unknown, _unknownErr := io.ReadUint16(16)
    if _unknownErr != nil {
        return nil, errors.New("Error parsing 'unknown' field " + _unknownErr.Error())
    }

    // Create a partially initialized instance
    _child := &MPropReadCon{
        InterfaceObjectType: interfaceObjectType,
        ObjectInstance: objectInstance,
        PropertyId: propertyId,
        NumberOfElements: numberOfElements,
        StartIndex: startIndex,
        Unknown: unknown,
        Parent: &CEMI{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *MPropReadCon) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (interfaceObjectType)
    interfaceObjectType := uint16(m.InterfaceObjectType)
    _interfaceObjectTypeErr := io.WriteUint16(16, (interfaceObjectType))
    if _interfaceObjectTypeErr != nil {
        return errors.New("Error serializing 'interfaceObjectType' field " + _interfaceObjectTypeErr.Error())
    }

    // Simple Field (objectInstance)
    objectInstance := uint8(m.ObjectInstance)
    _objectInstanceErr := io.WriteUint8(8, (objectInstance))
    if _objectInstanceErr != nil {
        return errors.New("Error serializing 'objectInstance' field " + _objectInstanceErr.Error())
    }

    // Simple Field (propertyId)
    propertyId := uint8(m.PropertyId)
    _propertyIdErr := io.WriteUint8(8, (propertyId))
    if _propertyIdErr != nil {
        return errors.New("Error serializing 'propertyId' field " + _propertyIdErr.Error())
    }

    // Simple Field (numberOfElements)
    numberOfElements := uint8(m.NumberOfElements)
    _numberOfElementsErr := io.WriteUint8(4, (numberOfElements))
    if _numberOfElementsErr != nil {
        return errors.New("Error serializing 'numberOfElements' field " + _numberOfElementsErr.Error())
    }

    // Simple Field (startIndex)
    startIndex := uint16(m.StartIndex)
    _startIndexErr := io.WriteUint16(12, (startIndex))
    if _startIndexErr != nil {
        return errors.New("Error serializing 'startIndex' field " + _startIndexErr.Error())
    }

    // Simple Field (unknown)
    unknown := uint16(m.Unknown)
    _unknownErr := io.WriteUint16(16, (unknown))
    if _unknownErr != nil {
        return errors.New("Error serializing 'unknown' field " + _unknownErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *MPropReadCon) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "interfaceObjectType":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.InterfaceObjectType = data
            case "objectInstance":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ObjectInstance = data
            case "propertyId":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.PropertyId = data
            case "numberOfElements":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.NumberOfElements = data
            case "startIndex":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.StartIndex = data
            case "unknown":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Unknown = data
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

func (m *MPropReadCon) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.InterfaceObjectType, xml.StartElement{Name: xml.Name{Local: "interfaceObjectType"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ObjectInstance, xml.StartElement{Name: xml.Name{Local: "objectInstance"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.PropertyId, xml.StartElement{Name: xml.Name{Local: "propertyId"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.NumberOfElements, xml.StartElement{Name: xml.Name{Local: "numberOfElements"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.StartIndex, xml.StartElement{Name: xml.Name{Local: "startIndex"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Unknown, xml.StartElement{Name: xml.Name{Local: "unknown"}}); err != nil {
        return err
    }
    return nil
}

