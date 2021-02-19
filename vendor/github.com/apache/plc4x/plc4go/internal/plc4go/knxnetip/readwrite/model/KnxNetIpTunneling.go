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
type KnxNetIpTunneling struct {
    Version uint8
    Parent *ServiceId
    IKnxNetIpTunneling
}

// The corresponding interface
type IKnxNetIpTunneling interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *KnxNetIpTunneling) ServiceType() uint8 {
    return 0x04
}


func (m *KnxNetIpTunneling) InitializeParent(parent *ServiceId) {
}

func NewKnxNetIpTunneling(version uint8, ) *ServiceId {
    child := &KnxNetIpTunneling{
        Version: version,
        Parent: NewServiceId(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastKnxNetIpTunneling(structType interface{}) *KnxNetIpTunneling {
    castFunc := func(typ interface{}) *KnxNetIpTunneling {
        if casted, ok := typ.(KnxNetIpTunneling); ok {
            return &casted
        }
        if casted, ok := typ.(*KnxNetIpTunneling); ok {
            return casted
        }
        if casted, ok := typ.(ServiceId); ok {
            return CastKnxNetIpTunneling(casted.Child)
        }
        if casted, ok := typ.(*ServiceId); ok {
            return CastKnxNetIpTunneling(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *KnxNetIpTunneling) GetTypeName() string {
    return "KnxNetIpTunneling"
}

func (m *KnxNetIpTunneling) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (version)
    lengthInBits += 8

    return lengthInBits
}

func (m *KnxNetIpTunneling) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func KnxNetIpTunnelingParse(io *utils.ReadBuffer) (*ServiceId, error) {

    // Simple Field (version)
    version, _versionErr := io.ReadUint8(8)
    if _versionErr != nil {
        return nil, errors.New("Error parsing 'version' field " + _versionErr.Error())
    }

    // Create a partially initialized instance
    _child := &KnxNetIpTunneling{
        Version: version,
        Parent: &ServiceId{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *KnxNetIpTunneling) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (version)
    version := uint8(m.Version)
    _versionErr := io.WriteUint8(8, (version))
    if _versionErr != nil {
        return errors.New("Error serializing 'version' field " + _versionErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *KnxNetIpTunneling) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "version":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Version = data
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

func (m *KnxNetIpTunneling) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.Version, xml.StartElement{Name: xml.Name{Local: "version"}}); err != nil {
        return err
    }
    return nil
}

