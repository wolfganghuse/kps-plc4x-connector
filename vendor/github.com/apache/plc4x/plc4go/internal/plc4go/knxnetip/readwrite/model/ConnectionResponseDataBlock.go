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
    "reflect"
    "strings"
)

// The data-structure of this message
type ConnectionResponseDataBlock struct {
    Child IConnectionResponseDataBlockChild
    IConnectionResponseDataBlock
    IConnectionResponseDataBlockParent
}

// The corresponding interface
type IConnectionResponseDataBlock interface {
    ConnectionType() uint8
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type IConnectionResponseDataBlockParent interface {
    SerializeParent(io utils.WriteBuffer, child IConnectionResponseDataBlock, serializeChildFunction func() error) error
    GetTypeName() string
}

type IConnectionResponseDataBlockChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *ConnectionResponseDataBlock)
    GetTypeName() string
    IConnectionResponseDataBlock
}

func NewConnectionResponseDataBlock() *ConnectionResponseDataBlock {
    return &ConnectionResponseDataBlock{}
}

func CastConnectionResponseDataBlock(structType interface{}) *ConnectionResponseDataBlock {
    castFunc := func(typ interface{}) *ConnectionResponseDataBlock {
        if casted, ok := typ.(ConnectionResponseDataBlock); ok {
            return &casted
        }
        if casted, ok := typ.(*ConnectionResponseDataBlock); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ConnectionResponseDataBlock) GetTypeName() string {
    return "ConnectionResponseDataBlock"
}

func (m *ConnectionResponseDataBlock) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Implicit Field (structureLength)
    lengthInBits += 8

    // Discriminator Field (connectionType)
    lengthInBits += 8

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *ConnectionResponseDataBlock) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ConnectionResponseDataBlockParse(io *utils.ReadBuffer) (*ConnectionResponseDataBlock, error) {

    // Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    _, _structureLengthErr := io.ReadUint8(8)
    if _structureLengthErr != nil {
        return nil, errors.New("Error parsing 'structureLength' field " + _structureLengthErr.Error())
    }

    // Discriminator Field (connectionType) (Used as input to a switch field)
    connectionType, _connectionTypeErr := io.ReadUint8(8)
    if _connectionTypeErr != nil {
        return nil, errors.New("Error parsing 'connectionType' field " + _connectionTypeErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *ConnectionResponseDataBlock
    var typeSwitchError error
    switch {
    case connectionType == 0x03:
        _parent, typeSwitchError = ConnectionResponseDataBlockDeviceManagementParse(io)
    case connectionType == 0x04:
        _parent, typeSwitchError = ConnectionResponseDataBlockTunnelConnectionParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent)
    return _parent, nil
}

func (m *ConnectionResponseDataBlock) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *ConnectionResponseDataBlock) SerializeParent(io utils.WriteBuffer, child IConnectionResponseDataBlock, serializeChildFunction func() error) error {

    // Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    structureLength := uint8(uint8(m.LengthInBytes()))
    _structureLengthErr := io.WriteUint8(8, (structureLength))
    if _structureLengthErr != nil {
        return errors.New("Error serializing 'structureLength' field " + _structureLengthErr.Error())
    }

    // Discriminator Field (connectionType) (Used as input to a switch field)
    connectionType := uint8(child.ConnectionType())
    _connectionTypeErr := io.WriteUint8(8, (connectionType))
    if _connectionTypeErr != nil {
        return errors.New("Error serializing 'connectionType' field " + _connectionTypeErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *ConnectionResponseDataBlock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            default:
                switch start.Attr[0].Value {
                    case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionResponseDataBlockDeviceManagement":
                        var dt *ConnectionResponseDataBlockDeviceManagement
                        if m.Child != nil {
                            dt = m.Child.(*ConnectionResponseDataBlockDeviceManagement)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionResponseDataBlockTunnelConnection":
                        var dt *ConnectionResponseDataBlockTunnelConnection
                        if m.Child != nil {
                            dt = m.Child.(*ConnectionResponseDataBlockTunnelConnection)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                }
            }
        }
    }
}

func (m *ConnectionResponseDataBlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := reflect.TypeOf(m.Child).String()
    className = "org.apache.plc4x.java.knxnetip.readwrite." + className[strings.LastIndex(className, ".") + 1:]
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    marshaller, ok := m.Child.(xml.Marshaler)
    if !ok {
        return errors.New("child is not castable to Marshaler")
    }
    marshaller.MarshalXML(e, start)
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

