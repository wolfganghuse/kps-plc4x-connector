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
type ApduDataExtPropertyValueRead struct {
    ObjectIndex uint8
    PropertyId uint8
    Count uint8
    Index uint16
    Parent *ApduDataExt
    IApduDataExtPropertyValueRead
}

// The corresponding interface
type IApduDataExtPropertyValueRead interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtPropertyValueRead) ExtApciType() uint8 {
    return 0x15
}


func (m *ApduDataExtPropertyValueRead) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtPropertyValueRead(objectIndex uint8, propertyId uint8, count uint8, index uint16, ) *ApduDataExt {
    child := &ApduDataExtPropertyValueRead{
        ObjectIndex: objectIndex,
        PropertyId: propertyId,
        Count: count,
        Index: index,
        Parent: NewApduDataExt(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataExtPropertyValueRead(structType interface{}) *ApduDataExtPropertyValueRead {
    castFunc := func(typ interface{}) *ApduDataExtPropertyValueRead {
        if casted, ok := typ.(ApduDataExtPropertyValueRead); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataExtPropertyValueRead); ok {
            return casted
        }
        if casted, ok := typ.(ApduDataExt); ok {
            return CastApduDataExtPropertyValueRead(casted.Child)
        }
        if casted, ok := typ.(*ApduDataExt); ok {
            return CastApduDataExtPropertyValueRead(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataExtPropertyValueRead) GetTypeName() string {
    return "ApduDataExtPropertyValueRead"
}

func (m *ApduDataExtPropertyValueRead) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (objectIndex)
    lengthInBits += 8

    // Simple field (propertyId)
    lengthInBits += 8

    // Simple field (count)
    lengthInBits += 4

    // Simple field (index)
    lengthInBits += 12

    return lengthInBits
}

func (m *ApduDataExtPropertyValueRead) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataExtPropertyValueReadParse(io *utils.ReadBuffer) (*ApduDataExt, error) {

    // Simple Field (objectIndex)
    objectIndex, _objectIndexErr := io.ReadUint8(8)
    if _objectIndexErr != nil {
        return nil, errors.New("Error parsing 'objectIndex' field " + _objectIndexErr.Error())
    }

    // Simple Field (propertyId)
    propertyId, _propertyIdErr := io.ReadUint8(8)
    if _propertyIdErr != nil {
        return nil, errors.New("Error parsing 'propertyId' field " + _propertyIdErr.Error())
    }

    // Simple Field (count)
    count, _countErr := io.ReadUint8(4)
    if _countErr != nil {
        return nil, errors.New("Error parsing 'count' field " + _countErr.Error())
    }

    // Simple Field (index)
    index, _indexErr := io.ReadUint16(12)
    if _indexErr != nil {
        return nil, errors.New("Error parsing 'index' field " + _indexErr.Error())
    }

    // Create a partially initialized instance
    _child := &ApduDataExtPropertyValueRead{
        ObjectIndex: objectIndex,
        PropertyId: propertyId,
        Count: count,
        Index: index,
        Parent: &ApduDataExt{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataExtPropertyValueRead) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (objectIndex)
    objectIndex := uint8(m.ObjectIndex)
    _objectIndexErr := io.WriteUint8(8, (objectIndex))
    if _objectIndexErr != nil {
        return errors.New("Error serializing 'objectIndex' field " + _objectIndexErr.Error())
    }

    // Simple Field (propertyId)
    propertyId := uint8(m.PropertyId)
    _propertyIdErr := io.WriteUint8(8, (propertyId))
    if _propertyIdErr != nil {
        return errors.New("Error serializing 'propertyId' field " + _propertyIdErr.Error())
    }

    // Simple Field (count)
    count := uint8(m.Count)
    _countErr := io.WriteUint8(4, (count))
    if _countErr != nil {
        return errors.New("Error serializing 'count' field " + _countErr.Error())
    }

    // Simple Field (index)
    index := uint16(m.Index)
    _indexErr := io.WriteUint16(12, (index))
    if _indexErr != nil {
        return errors.New("Error serializing 'index' field " + _indexErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtPropertyValueRead) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "objectIndex":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ObjectIndex = data
            case "propertyId":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.PropertyId = data
            case "count":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Count = data
            case "index":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Index = data
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

func (m *ApduDataExtPropertyValueRead) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.ObjectIndex, xml.StartElement{Name: xml.Name{Local: "objectIndex"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.PropertyId, xml.StartElement{Name: xml.Name{Local: "propertyId"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Count, xml.StartElement{Name: xml.Name{Local: "count"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Index, xml.StartElement{Name: xml.Name{Local: "index"}}); err != nil {
        return err
    }
    return nil
}

