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
type TunnelingRequest struct {
    TunnelingRequestDataBlock *TunnelingRequestDataBlock
    Cemi *CEMI
    Parent *KnxNetIpMessage
    ITunnelingRequest
}

// The corresponding interface
type ITunnelingRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *TunnelingRequest) MsgType() uint16 {
    return 0x0420
}


func (m *TunnelingRequest) InitializeParent(parent *KnxNetIpMessage) {
}

func NewTunnelingRequest(tunnelingRequestDataBlock *TunnelingRequestDataBlock, cemi *CEMI, ) *KnxNetIpMessage {
    child := &TunnelingRequest{
        TunnelingRequestDataBlock: tunnelingRequestDataBlock,
        Cemi: cemi,
        Parent: NewKnxNetIpMessage(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastTunnelingRequest(structType interface{}) *TunnelingRequest {
    castFunc := func(typ interface{}) *TunnelingRequest {
        if casted, ok := typ.(TunnelingRequest); ok {
            return &casted
        }
        if casted, ok := typ.(*TunnelingRequest); ok {
            return casted
        }
        if casted, ok := typ.(KnxNetIpMessage); ok {
            return CastTunnelingRequest(casted.Child)
        }
        if casted, ok := typ.(*KnxNetIpMessage); ok {
            return CastTunnelingRequest(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *TunnelingRequest) GetTypeName() string {
    return "TunnelingRequest"
}

func (m *TunnelingRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (tunnelingRequestDataBlock)
    lengthInBits += m.TunnelingRequestDataBlock.LengthInBits()

    // Simple field (cemi)
    lengthInBits += m.Cemi.LengthInBits()

    return lengthInBits
}

func (m *TunnelingRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func TunnelingRequestParse(io *utils.ReadBuffer, totalLength uint16) (*KnxNetIpMessage, error) {

    // Simple Field (tunnelingRequestDataBlock)
    tunnelingRequestDataBlock, _tunnelingRequestDataBlockErr := TunnelingRequestDataBlockParse(io)
    if _tunnelingRequestDataBlockErr != nil {
        return nil, errors.New("Error parsing 'tunnelingRequestDataBlock' field " + _tunnelingRequestDataBlockErr.Error())
    }

    // Simple Field (cemi)
    cemi, _cemiErr := CEMIParse(io, uint8(totalLength) - uint8(uint8(uint8(uint8(6)) + uint8(tunnelingRequestDataBlock.LengthInBytes()))))
    if _cemiErr != nil {
        return nil, errors.New("Error parsing 'cemi' field " + _cemiErr.Error())
    }

    // Create a partially initialized instance
    _child := &TunnelingRequest{
        TunnelingRequestDataBlock: tunnelingRequestDataBlock,
        Cemi: cemi,
        Parent: &KnxNetIpMessage{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *TunnelingRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (tunnelingRequestDataBlock)
    _tunnelingRequestDataBlockErr := m.TunnelingRequestDataBlock.Serialize(io)
    if _tunnelingRequestDataBlockErr != nil {
        return errors.New("Error serializing 'tunnelingRequestDataBlock' field " + _tunnelingRequestDataBlockErr.Error())
    }

    // Simple Field (cemi)
    _cemiErr := m.Cemi.Serialize(io)
    if _cemiErr != nil {
        return errors.New("Error serializing 'cemi' field " + _cemiErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *TunnelingRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "tunnelingRequestDataBlock":
                var data *TunnelingRequestDataBlock
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.TunnelingRequestDataBlock = data
            case "cemi":
                var dt *CEMI
                if err := d.DecodeElement(&dt, &tok); err != nil {
                    return err
                }
                m.Cemi = dt
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

func (m *TunnelingRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.TunnelingRequestDataBlock, xml.StartElement{Name: xml.Name{Local: "tunnelingRequestDataBlock"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Cemi, xml.StartElement{Name: xml.Name{Local: "cemi"}}); err != nil {
        return err
    }
    return nil
}

