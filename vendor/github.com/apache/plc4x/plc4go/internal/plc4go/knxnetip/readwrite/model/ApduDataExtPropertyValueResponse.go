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
type ApduDataExtPropertyValueResponse struct {
    ObjectIndex uint8
    PropertyId uint8
    Count uint8
    Index uint16
    Data []uint8
    Parent *ApduDataExt
    IApduDataExtPropertyValueResponse
}

// The corresponding interface
type IApduDataExtPropertyValueResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtPropertyValueResponse) ExtApciType() uint8 {
    return 0x16
}


func (m *ApduDataExtPropertyValueResponse) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtPropertyValueResponse(objectIndex uint8, propertyId uint8, count uint8, index uint16, data []uint8, ) *ApduDataExt {
    child := &ApduDataExtPropertyValueResponse{
        ObjectIndex: objectIndex,
        PropertyId: propertyId,
        Count: count,
        Index: index,
        Data: data,
        Parent: NewApduDataExt(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataExtPropertyValueResponse(structType interface{}) *ApduDataExtPropertyValueResponse {
    castFunc := func(typ interface{}) *ApduDataExtPropertyValueResponse {
        if casted, ok := typ.(ApduDataExtPropertyValueResponse); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataExtPropertyValueResponse); ok {
            return casted
        }
        if casted, ok := typ.(ApduDataExt); ok {
            return CastApduDataExtPropertyValueResponse(casted.Child)
        }
        if casted, ok := typ.(*ApduDataExt); ok {
            return CastApduDataExtPropertyValueResponse(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataExtPropertyValueResponse) GetTypeName() string {
    return "ApduDataExtPropertyValueResponse"
}

func (m *ApduDataExtPropertyValueResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (objectIndex)
    lengthInBits += 8

    // Simple field (propertyId)
    lengthInBits += 8

    // Simple field (count)
    lengthInBits += 4

    // Simple field (index)
    lengthInBits += 12

    // Array field
    if len(m.Data) > 0 {
        lengthInBits += 8 * uint16(len(m.Data))
    }

    return lengthInBits
}

func (m *ApduDataExtPropertyValueResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataExtPropertyValueResponseParse(io *utils.ReadBuffer, length uint8) (*ApduDataExt, error) {

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

    // Array field (data)
    // Count array
    data := make([]uint8, uint16(length) - uint16(uint16(5)))
    for curItem := uint16(0); curItem < uint16(uint16(length) - uint16(uint16(5))); curItem++ {
        _item, _err := io.ReadUint8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'data' field " + _err.Error())
        }
        data[curItem] = _item
    }

    // Create a partially initialized instance
    _child := &ApduDataExtPropertyValueResponse{
        ObjectIndex: objectIndex,
        PropertyId: propertyId,
        Count: count,
        Index: index,
        Data: data,
        Parent: &ApduDataExt{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataExtPropertyValueResponse) Serialize(io utils.WriteBuffer) error {
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

    // Array Field (data)
    if m.Data != nil {
        for _, _element := range m.Data {
            _elementErr := io.WriteUint8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'data' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtPropertyValueResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "data":
                var data []uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Data = data
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

func (m *ApduDataExtPropertyValueResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
    if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    return nil
}

