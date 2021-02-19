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
type ChannelInformation struct {
    NumChannels uint8
    ChannelCode uint16
    IChannelInformation
}

// The corresponding interface
type IChannelInformation interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

func NewChannelInformation(numChannels uint8, channelCode uint16) *ChannelInformation {
    return &ChannelInformation{NumChannels: numChannels, ChannelCode: channelCode}
}

func CastChannelInformation(structType interface{}) *ChannelInformation {
    castFunc := func(typ interface{}) *ChannelInformation {
        if casted, ok := typ.(ChannelInformation); ok {
            return &casted
        }
        if casted, ok := typ.(*ChannelInformation); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ChannelInformation) GetTypeName() string {
    return "ChannelInformation"
}

func (m *ChannelInformation) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (numChannels)
    lengthInBits += 3

    // Simple field (channelCode)
    lengthInBits += 13

    return lengthInBits
}

func (m *ChannelInformation) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ChannelInformationParse(io *utils.ReadBuffer) (*ChannelInformation, error) {

    // Simple Field (numChannels)
    numChannels, _numChannelsErr := io.ReadUint8(3)
    if _numChannelsErr != nil {
        return nil, errors.New("Error parsing 'numChannels' field " + _numChannelsErr.Error())
    }

    // Simple Field (channelCode)
    channelCode, _channelCodeErr := io.ReadUint16(13)
    if _channelCodeErr != nil {
        return nil, errors.New("Error parsing 'channelCode' field " + _channelCodeErr.Error())
    }

    // Create the instance
    return NewChannelInformation(numChannels, channelCode), nil
}

func (m *ChannelInformation) Serialize(io utils.WriteBuffer) error {

    // Simple Field (numChannels)
    numChannels := uint8(m.NumChannels)
    _numChannelsErr := io.WriteUint8(3, (numChannels))
    if _numChannelsErr != nil {
        return errors.New("Error serializing 'numChannels' field " + _numChannelsErr.Error())
    }

    // Simple Field (channelCode)
    channelCode := uint16(m.ChannelCode)
    _channelCodeErr := io.WriteUint16(13, (channelCode))
    if _channelCodeErr != nil {
        return errors.New("Error serializing 'channelCode' field " + _channelCodeErr.Error())
    }

    return nil
}

func (m *ChannelInformation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "numChannels":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.NumChannels = data
            case "channelCode":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ChannelCode = data
            }
        }
    }
}

func (m *ChannelInformation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := "org.apache.plc4x.java.knxnetip.readwrite.ChannelInformation"
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.NumChannels, xml.StartElement{Name: xml.Name{Local: "numChannels"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ChannelCode, xml.StartElement{Name: xml.Name{Local: "channelCode"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

