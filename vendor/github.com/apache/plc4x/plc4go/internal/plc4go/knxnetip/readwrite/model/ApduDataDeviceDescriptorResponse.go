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
    "encoding/base64"
    "encoding/xml"
    "errors"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
)

// The data-structure of this message
type ApduDataDeviceDescriptorResponse struct {
    DescriptorType uint8
    Data []int8
    Parent *ApduData
    IApduDataDeviceDescriptorResponse
}

// The corresponding interface
type IApduDataDeviceDescriptorResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataDeviceDescriptorResponse) ApciType() uint8 {
    return 0xD
}


func (m *ApduDataDeviceDescriptorResponse) InitializeParent(parent *ApduData) {
}

func NewApduDataDeviceDescriptorResponse(descriptorType uint8, data []int8, ) *ApduData {
    child := &ApduDataDeviceDescriptorResponse{
        DescriptorType: descriptorType,
        Data: data,
        Parent: NewApduData(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataDeviceDescriptorResponse(structType interface{}) *ApduDataDeviceDescriptorResponse {
    castFunc := func(typ interface{}) *ApduDataDeviceDescriptorResponse {
        if casted, ok := typ.(ApduDataDeviceDescriptorResponse); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataDeviceDescriptorResponse); ok {
            return casted
        }
        if casted, ok := typ.(ApduData); ok {
            return CastApduDataDeviceDescriptorResponse(casted.Child)
        }
        if casted, ok := typ.(*ApduData); ok {
            return CastApduDataDeviceDescriptorResponse(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataDeviceDescriptorResponse) GetTypeName() string {
    return "ApduDataDeviceDescriptorResponse"
}

func (m *ApduDataDeviceDescriptorResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (descriptorType)
    lengthInBits += 6

    // Array field
    if len(m.Data) > 0 {
        lengthInBits += 8 * uint16(len(m.Data))
    }

    return lengthInBits
}

func (m *ApduDataDeviceDescriptorResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataDeviceDescriptorResponseParse(io *utils.ReadBuffer, dataLength uint8) (*ApduData, error) {

    // Simple Field (descriptorType)
    descriptorType, _descriptorTypeErr := io.ReadUint8(6)
    if _descriptorTypeErr != nil {
        return nil, errors.New("Error parsing 'descriptorType' field " + _descriptorTypeErr.Error())
    }

    // Array field (data)
    // Count array
    data := make([]int8, utils.InlineIf(bool(bool((dataLength) < ((1)))), uint16(uint16(0)), uint16(uint16(dataLength) - uint16(uint16(1)))))
    for curItem := uint16(0); curItem < uint16(utils.InlineIf(bool(bool((dataLength) < ((1)))), uint16(uint16(0)), uint16(uint16(dataLength) - uint16(uint16(1))))); curItem++ {
        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'data' field " + _err.Error())
        }
        data[curItem] = _item
    }

    // Create a partially initialized instance
    _child := &ApduDataDeviceDescriptorResponse{
        DescriptorType: descriptorType,
        Data: data,
        Parent: &ApduData{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataDeviceDescriptorResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (descriptorType)
    descriptorType := uint8(m.DescriptorType)
    _descriptorTypeErr := io.WriteUint8(6, (descriptorType))
    if _descriptorTypeErr != nil {
        return errors.New("Error serializing 'descriptorType' field " + _descriptorTypeErr.Error())
    }

    // Array Field (data)
    if m.Data != nil {
        for _, _element := range m.Data {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'data' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataDeviceDescriptorResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "descriptorType":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.DescriptorType = data
            case "data":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *ApduDataDeviceDescriptorResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.DescriptorType, xml.StartElement{Name: xml.Name{Local: "descriptorType"}}); err != nil {
        return err
    }
    _encodedData := make([]byte, base64.StdEncoding.EncodedLen(len(m.Data)))
    base64.StdEncoding.Encode(_encodedData, utils.Int8ArrayToByteArray(m.Data))
    if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    return nil
}

