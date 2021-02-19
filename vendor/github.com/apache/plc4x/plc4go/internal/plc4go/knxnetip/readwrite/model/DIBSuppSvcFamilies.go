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
type DIBSuppSvcFamilies struct {
    DescriptionType uint8
    ServiceIds []*ServiceId
    IDIBSuppSvcFamilies
}

// The corresponding interface
type IDIBSuppSvcFamilies interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

func NewDIBSuppSvcFamilies(descriptionType uint8, serviceIds []*ServiceId) *DIBSuppSvcFamilies {
    return &DIBSuppSvcFamilies{DescriptionType: descriptionType, ServiceIds: serviceIds}
}

func CastDIBSuppSvcFamilies(structType interface{}) *DIBSuppSvcFamilies {
    castFunc := func(typ interface{}) *DIBSuppSvcFamilies {
        if casted, ok := typ.(DIBSuppSvcFamilies); ok {
            return &casted
        }
        if casted, ok := typ.(*DIBSuppSvcFamilies); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *DIBSuppSvcFamilies) GetTypeName() string {
    return "DIBSuppSvcFamilies"
}

func (m *DIBSuppSvcFamilies) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Implicit Field (structureLength)
    lengthInBits += 8

    // Simple field (descriptionType)
    lengthInBits += 8

    // Array field
    if len(m.ServiceIds) > 0 {
        for _, element := range m.ServiceIds {
            lengthInBits += element.LengthInBits()
        }
    }

    return lengthInBits
}

func (m *DIBSuppSvcFamilies) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func DIBSuppSvcFamiliesParse(io *utils.ReadBuffer) (*DIBSuppSvcFamilies, error) {

    // Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    structureLength, _structureLengthErr := io.ReadUint8(8)
    if _structureLengthErr != nil {
        return nil, errors.New("Error parsing 'structureLength' field " + _structureLengthErr.Error())
    }

    // Simple Field (descriptionType)
    descriptionType, _descriptionTypeErr := io.ReadUint8(8)
    if _descriptionTypeErr != nil {
        return nil, errors.New("Error parsing 'descriptionType' field " + _descriptionTypeErr.Error())
    }

    // Array field (serviceIds)
    // Length array
    serviceIds := make([]*ServiceId, 0)
    _serviceIdsLength := uint16(structureLength) - uint16(uint16(2))
    _serviceIdsEndPos := io.GetPos() + uint16(_serviceIdsLength)
    for ;io.GetPos() < _serviceIdsEndPos; {
        _item, _err := ServiceIdParse(io)
        if _err != nil {
            return nil, errors.New("Error parsing 'serviceIds' field " + _err.Error())
        }
        serviceIds = append(serviceIds, _item)
    }

    // Create the instance
    return NewDIBSuppSvcFamilies(descriptionType, serviceIds), nil
}

func (m *DIBSuppSvcFamilies) Serialize(io utils.WriteBuffer) error {

    // Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    structureLength := uint8(uint8(m.LengthInBytes()))
    _structureLengthErr := io.WriteUint8(8, (structureLength))
    if _structureLengthErr != nil {
        return errors.New("Error serializing 'structureLength' field " + _structureLengthErr.Error())
    }

    // Simple Field (descriptionType)
    descriptionType := uint8(m.DescriptionType)
    _descriptionTypeErr := io.WriteUint8(8, (descriptionType))
    if _descriptionTypeErr != nil {
        return errors.New("Error serializing 'descriptionType' field " + _descriptionTypeErr.Error())
    }

    // Array Field (serviceIds)
    if m.ServiceIds != nil {
        for _, _element := range m.ServiceIds {
            _elementErr := _element.Serialize(io)
            if _elementErr != nil {
                return errors.New("Error serializing 'serviceIds' field " + _elementErr.Error())
            }
        }
    }

    return nil
}

func (m *DIBSuppSvcFamilies) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "descriptionType":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.DescriptionType = data
            case "serviceIds":
                var _values []*ServiceId
                var dt *ServiceId
                if err := d.DecodeElement(&dt, &tok); err != nil {
                    return err
                }
                _values = append(_values, dt)
                m.ServiceIds = _values
            }
        }
    }
}

func (m *DIBSuppSvcFamilies) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := "org.apache.plc4x.java.knxnetip.readwrite.DIBSuppSvcFamilies"
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.DescriptionType, xml.StartElement{Name: xml.Name{Local: "descriptionType"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "serviceIds"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ServiceIds, xml.StartElement{Name: xml.Name{Local: "serviceIds"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "serviceIds"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

