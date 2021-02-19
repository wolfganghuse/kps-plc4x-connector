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
type ConnectionRequestInformationTunnelConnection struct {
    KnxLayer KnxLayer
    Parent *ConnectionRequestInformation
    IConnectionRequestInformationTunnelConnection
}

// The corresponding interface
type IConnectionRequestInformationTunnelConnection interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ConnectionRequestInformationTunnelConnection) ConnectionType() uint8 {
    return 0x04
}


func (m *ConnectionRequestInformationTunnelConnection) InitializeParent(parent *ConnectionRequestInformation) {
}

func NewConnectionRequestInformationTunnelConnection(knxLayer KnxLayer, ) *ConnectionRequestInformation {
    child := &ConnectionRequestInformationTunnelConnection{
        KnxLayer: knxLayer,
        Parent: NewConnectionRequestInformation(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastConnectionRequestInformationTunnelConnection(structType interface{}) *ConnectionRequestInformationTunnelConnection {
    castFunc := func(typ interface{}) *ConnectionRequestInformationTunnelConnection {
        if casted, ok := typ.(ConnectionRequestInformationTunnelConnection); ok {
            return &casted
        }
        if casted, ok := typ.(*ConnectionRequestInformationTunnelConnection); ok {
            return casted
        }
        if casted, ok := typ.(ConnectionRequestInformation); ok {
            return CastConnectionRequestInformationTunnelConnection(casted.Child)
        }
        if casted, ok := typ.(*ConnectionRequestInformation); ok {
            return CastConnectionRequestInformationTunnelConnection(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ConnectionRequestInformationTunnelConnection) GetTypeName() string {
    return "ConnectionRequestInformationTunnelConnection"
}

func (m *ConnectionRequestInformationTunnelConnection) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (knxLayer)
    lengthInBits += 8

    // Reserved Field (reserved)
    lengthInBits += 8

    return lengthInBits
}

func (m *ConnectionRequestInformationTunnelConnection) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ConnectionRequestInformationTunnelConnectionParse(io *utils.ReadBuffer) (*ConnectionRequestInformation, error) {

    // Simple Field (knxLayer)
    knxLayer, _knxLayerErr := KnxLayerParse(io)
    if _knxLayerErr != nil {
        return nil, errors.New("Error parsing 'knxLayer' field " + _knxLayerErr.Error())
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

    // Create a partially initialized instance
    _child := &ConnectionRequestInformationTunnelConnection{
        KnxLayer: knxLayer,
        Parent: &ConnectionRequestInformation{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ConnectionRequestInformationTunnelConnection) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (knxLayer)
    _knxLayerErr := m.KnxLayer.Serialize(io)
    if _knxLayerErr != nil {
        return errors.New("Error serializing 'knxLayer' field " + _knxLayerErr.Error())
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
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ConnectionRequestInformationTunnelConnection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "knxLayer":
                var data KnxLayer
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.KnxLayer = data
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

func (m *ConnectionRequestInformationTunnelConnection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.KnxLayer, xml.StartElement{Name: xml.Name{Local: "knxLayer"}}); err != nil {
        return err
    }
    return nil
}

