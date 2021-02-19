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
type RelativeTimestamp struct {
    Timestamp uint16
    IRelativeTimestamp
}

// The corresponding interface
type IRelativeTimestamp interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

func NewRelativeTimestamp(timestamp uint16) *RelativeTimestamp {
    return &RelativeTimestamp{Timestamp: timestamp}
}

func CastRelativeTimestamp(structType interface{}) *RelativeTimestamp {
    castFunc := func(typ interface{}) *RelativeTimestamp {
        if casted, ok := typ.(RelativeTimestamp); ok {
            return &casted
        }
        if casted, ok := typ.(*RelativeTimestamp); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *RelativeTimestamp) GetTypeName() string {
    return "RelativeTimestamp"
}

func (m *RelativeTimestamp) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (timestamp)
    lengthInBits += 16

    return lengthInBits
}

func (m *RelativeTimestamp) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func RelativeTimestampParse(io *utils.ReadBuffer) (*RelativeTimestamp, error) {

    // Simple Field (timestamp)
    timestamp, _timestampErr := io.ReadUint16(16)
    if _timestampErr != nil {
        return nil, errors.New("Error parsing 'timestamp' field " + _timestampErr.Error())
    }

    // Create the instance
    return NewRelativeTimestamp(timestamp), nil
}

func (m *RelativeTimestamp) Serialize(io utils.WriteBuffer) error {

    // Simple Field (timestamp)
    timestamp := uint16(m.Timestamp)
    _timestampErr := io.WriteUint16(16, (timestamp))
    if _timestampErr != nil {
        return errors.New("Error serializing 'timestamp' field " + _timestampErr.Error())
    }

    return nil
}

func (m *RelativeTimestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "timestamp":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Timestamp = data
            }
        }
    }
}

func (m *RelativeTimestamp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := "org.apache.plc4x.java.knxnetip.readwrite.RelativeTimestamp"
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Timestamp, xml.StartElement{Name: xml.Name{Local: "timestamp"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

