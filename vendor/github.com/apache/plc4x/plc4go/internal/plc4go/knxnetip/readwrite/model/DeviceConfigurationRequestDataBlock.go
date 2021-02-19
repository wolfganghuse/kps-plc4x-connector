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
    log "github.com/sirupsen/logrus"
    "io"
)

// The data-structure of this message
type DeviceConfigurationRequestDataBlock struct {
    CommunicationChannelId uint8
    SequenceCounter uint8
    IDeviceConfigurationRequestDataBlock
}

// The corresponding interface
type IDeviceConfigurationRequestDataBlock interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

func NewDeviceConfigurationRequestDataBlock(communicationChannelId uint8, sequenceCounter uint8) *DeviceConfigurationRequestDataBlock {
    return &DeviceConfigurationRequestDataBlock{CommunicationChannelId: communicationChannelId, SequenceCounter: sequenceCounter}
}

func CastDeviceConfigurationRequestDataBlock(structType interface{}) *DeviceConfigurationRequestDataBlock {
    castFunc := func(typ interface{}) *DeviceConfigurationRequestDataBlock {
        if casted, ok := typ.(DeviceConfigurationRequestDataBlock); ok {
            return &casted
        }
        if casted, ok := typ.(*DeviceConfigurationRequestDataBlock); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *DeviceConfigurationRequestDataBlock) GetTypeName() string {
    return "DeviceConfigurationRequestDataBlock"
}

func (m *DeviceConfigurationRequestDataBlock) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Implicit Field (structureLength)
    lengthInBits += 8

    // Simple field (communicationChannelId)
    lengthInBits += 8

    // Simple field (sequenceCounter)
    lengthInBits += 8

    // Reserved Field (reserved)
    lengthInBits += 8

    return lengthInBits
}

func (m *DeviceConfigurationRequestDataBlock) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func DeviceConfigurationRequestDataBlockParse(io *utils.ReadBuffer) (*DeviceConfigurationRequestDataBlock, error) {

    // Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    _, _structureLengthErr := io.ReadUint8(8)
    if _structureLengthErr != nil {
        return nil, errors.New("Error parsing 'structureLength' field " + _structureLengthErr.Error())
    }

    // Simple Field (communicationChannelId)
    communicationChannelId, _communicationChannelIdErr := io.ReadUint8(8)
    if _communicationChannelIdErr != nil {
        return nil, errors.New("Error parsing 'communicationChannelId' field " + _communicationChannelIdErr.Error())
    }

    // Simple Field (sequenceCounter)
    sequenceCounter, _sequenceCounterErr := io.ReadUint8(8)
    if _sequenceCounterErr != nil {
        return nil, errors.New("Error parsing 'sequenceCounter' field " + _sequenceCounterErr.Error())
    }

    // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
    {
        reserved, _err := io.ReadUint8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'reserved' field " + _err.Error())
        }
        if reserved != uint8(0x00) {
            log.WithFields(log.Fields{
                "expected value": uint8(0x00),
                "got value": reserved,
            }).Info("Got unexpected response.")
        }
    }

    // Create the instance
    return NewDeviceConfigurationRequestDataBlock(communicationChannelId, sequenceCounter), nil
}

func (m *DeviceConfigurationRequestDataBlock) Serialize(io utils.WriteBuffer) error {

    // Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    structureLength := uint8(uint8(m.LengthInBytes()))
    _structureLengthErr := io.WriteUint8(8, (structureLength))
    if _structureLengthErr != nil {
        return errors.New("Error serializing 'structureLength' field " + _structureLengthErr.Error())
    }

    // Simple Field (communicationChannelId)
    communicationChannelId := uint8(m.CommunicationChannelId)
    _communicationChannelIdErr := io.WriteUint8(8, (communicationChannelId))
    if _communicationChannelIdErr != nil {
        return errors.New("Error serializing 'communicationChannelId' field " + _communicationChannelIdErr.Error())
    }

    // Simple Field (sequenceCounter)
    sequenceCounter := uint8(m.SequenceCounter)
    _sequenceCounterErr := io.WriteUint8(8, (sequenceCounter))
    if _sequenceCounterErr != nil {
        return errors.New("Error serializing 'sequenceCounter' field " + _sequenceCounterErr.Error())
    }

    // Reserved Field (reserved)
    {
        _err := io.WriteUint8(8, uint8(0x00))
        if _err != nil {
            return errors.New("Error serializing 'reserved' field " + _err.Error())
        }
    }

    return nil
}

func (m *DeviceConfigurationRequestDataBlock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "communicationChannelId":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.CommunicationChannelId = data
            case "sequenceCounter":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.SequenceCounter = data
            }
        }
    }
}

func (m *DeviceConfigurationRequestDataBlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := "org.apache.plc4x.java.knxnetip.readwrite.DeviceConfigurationRequestDataBlock"
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.CommunicationChannelId, xml.StartElement{Name: xml.Name{Local: "communicationChannelId"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.SequenceCounter, xml.StartElement{Name: xml.Name{Local: "sequenceCounter"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

