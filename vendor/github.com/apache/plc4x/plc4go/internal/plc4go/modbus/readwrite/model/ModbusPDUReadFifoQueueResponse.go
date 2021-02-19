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
type ModbusPDUReadFifoQueueResponse struct {
    FifoValue []uint16
    Parent *ModbusPDU
    IModbusPDUReadFifoQueueResponse
}

// The corresponding interface
type IModbusPDUReadFifoQueueResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadFifoQueueResponse) ErrorFlag() bool {
    return false
}

func (m *ModbusPDUReadFifoQueueResponse) FunctionFlag() uint8 {
    return 0x18
}

func (m *ModbusPDUReadFifoQueueResponse) Response() bool {
    return true
}


func (m *ModbusPDUReadFifoQueueResponse) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUReadFifoQueueResponse(fifoValue []uint16, ) *ModbusPDU {
    child := &ModbusPDUReadFifoQueueResponse{
        FifoValue: fifoValue,
        Parent: NewModbusPDU(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastModbusPDUReadFifoQueueResponse(structType interface{}) *ModbusPDUReadFifoQueueResponse {
    castFunc := func(typ interface{}) *ModbusPDUReadFifoQueueResponse {
        if casted, ok := typ.(ModbusPDUReadFifoQueueResponse); ok {
            return &casted
        }
        if casted, ok := typ.(*ModbusPDUReadFifoQueueResponse); ok {
            return casted
        }
        if casted, ok := typ.(ModbusPDU); ok {
            return CastModbusPDUReadFifoQueueResponse(casted.Child)
        }
        if casted, ok := typ.(*ModbusPDU); ok {
            return CastModbusPDUReadFifoQueueResponse(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ModbusPDUReadFifoQueueResponse) GetTypeName() string {
    return "ModbusPDUReadFifoQueueResponse"
}

func (m *ModbusPDUReadFifoQueueResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Implicit Field (byteCount)
    lengthInBits += 16

    // Implicit Field (fifoCount)
    lengthInBits += 16

    // Array field
    if len(m.FifoValue) > 0 {
        lengthInBits += 16 * uint16(len(m.FifoValue))
    }

    return lengthInBits
}

func (m *ModbusPDUReadFifoQueueResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUReadFifoQueueResponseParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

    // Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    _, _byteCountErr := io.ReadUint16(16)
    if _byteCountErr != nil {
        return nil, errors.New("Error parsing 'byteCount' field " + _byteCountErr.Error())
    }

    // Implicit Field (fifoCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    fifoCount, _fifoCountErr := io.ReadUint16(16)
    if _fifoCountErr != nil {
        return nil, errors.New("Error parsing 'fifoCount' field " + _fifoCountErr.Error())
    }

    // Array field (fifoValue)
    // Count array
    fifoValue := make([]uint16, fifoCount)
    for curItem := uint16(0); curItem < uint16(fifoCount); curItem++ {
        _item, _err := io.ReadUint16(16)
        if _err != nil {
            return nil, errors.New("Error parsing 'fifoValue' field " + _err.Error())
        }
        fifoValue[curItem] = _item
    }

    // Create a partially initialized instance
    _child := &ModbusPDUReadFifoQueueResponse{
        FifoValue: fifoValue,
        Parent: &ModbusPDU{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ModbusPDUReadFifoQueueResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    byteCount := uint16(uint16(uint16(uint16(uint16(len(m.FifoValue))) * uint16(uint16(2)))) + uint16(uint16(2)))
    _byteCountErr := io.WriteUint16(16, (byteCount))
    if _byteCountErr != nil {
        return errors.New("Error serializing 'byteCount' field " + _byteCountErr.Error())
    }

    // Implicit Field (fifoCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    fifoCount := uint16(uint16(uint16(uint16(uint16(len(m.FifoValue))) * uint16(uint16(2)))) / uint16(uint16(2)))
    _fifoCountErr := io.WriteUint16(16, (fifoCount))
    if _fifoCountErr != nil {
        return errors.New("Error serializing 'fifoCount' field " + _fifoCountErr.Error())
    }

    // Array Field (fifoValue)
    if m.FifoValue != nil {
        for _, _element := range m.FifoValue {
            _elementErr := io.WriteUint16(16, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'fifoValue' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUReadFifoQueueResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "fifoValue":
                var data []uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.FifoValue = data
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

func (m *ModbusPDUReadFifoQueueResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "fifoValue"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.FifoValue, xml.StartElement{Name: xml.Name{Local: "fifoValue"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "fifoValue"}}); err != nil {
        return err
    }
    return nil
}

