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
type ModbusPDUReadWriteMultipleHoldingRegistersRequest struct {
    ReadStartingAddress uint16
    ReadQuantity uint16
    WriteStartingAddress uint16
    WriteQuantity uint16
    Value []int8
    Parent *ModbusPDU
    IModbusPDUReadWriteMultipleHoldingRegistersRequest
}

// The corresponding interface
type IModbusPDUReadWriteMultipleHoldingRegistersRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) ErrorFlag() bool {
    return false
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) FunctionFlag() uint8 {
    return 0x17
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) Response() bool {
    return false
}


func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUReadWriteMultipleHoldingRegistersRequest(readStartingAddress uint16, readQuantity uint16, writeStartingAddress uint16, writeQuantity uint16, value []int8, ) *ModbusPDU {
    child := &ModbusPDUReadWriteMultipleHoldingRegistersRequest{
        ReadStartingAddress: readStartingAddress,
        ReadQuantity: readQuantity,
        WriteStartingAddress: writeStartingAddress,
        WriteQuantity: writeQuantity,
        Value: value,
        Parent: NewModbusPDU(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastModbusPDUReadWriteMultipleHoldingRegistersRequest(structType interface{}) *ModbusPDUReadWriteMultipleHoldingRegistersRequest {
    castFunc := func(typ interface{}) *ModbusPDUReadWriteMultipleHoldingRegistersRequest {
        if casted, ok := typ.(ModbusPDUReadWriteMultipleHoldingRegistersRequest); ok {
            return &casted
        }
        if casted, ok := typ.(*ModbusPDUReadWriteMultipleHoldingRegistersRequest); ok {
            return casted
        }
        if casted, ok := typ.(ModbusPDU); ok {
            return CastModbusPDUReadWriteMultipleHoldingRegistersRequest(casted.Child)
        }
        if casted, ok := typ.(*ModbusPDU); ok {
            return CastModbusPDUReadWriteMultipleHoldingRegistersRequest(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetTypeName() string {
    return "ModbusPDUReadWriteMultipleHoldingRegistersRequest"
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (readStartingAddress)
    lengthInBits += 16

    // Simple field (readQuantity)
    lengthInBits += 16

    // Simple field (writeStartingAddress)
    lengthInBits += 16

    // Simple field (writeQuantity)
    lengthInBits += 16

    // Implicit Field (byteCount)
    lengthInBits += 8

    // Array field
    if len(m.Value) > 0 {
        lengthInBits += 8 * uint16(len(m.Value))
    }

    return lengthInBits
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUReadWriteMultipleHoldingRegistersRequestParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

    // Simple Field (readStartingAddress)
    readStartingAddress, _readStartingAddressErr := io.ReadUint16(16)
    if _readStartingAddressErr != nil {
        return nil, errors.New("Error parsing 'readStartingAddress' field " + _readStartingAddressErr.Error())
    }

    // Simple Field (readQuantity)
    readQuantity, _readQuantityErr := io.ReadUint16(16)
    if _readQuantityErr != nil {
        return nil, errors.New("Error parsing 'readQuantity' field " + _readQuantityErr.Error())
    }

    // Simple Field (writeStartingAddress)
    writeStartingAddress, _writeStartingAddressErr := io.ReadUint16(16)
    if _writeStartingAddressErr != nil {
        return nil, errors.New("Error parsing 'writeStartingAddress' field " + _writeStartingAddressErr.Error())
    }

    // Simple Field (writeQuantity)
    writeQuantity, _writeQuantityErr := io.ReadUint16(16)
    if _writeQuantityErr != nil {
        return nil, errors.New("Error parsing 'writeQuantity' field " + _writeQuantityErr.Error())
    }

    // Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    byteCount, _byteCountErr := io.ReadUint8(8)
    if _byteCountErr != nil {
        return nil, errors.New("Error parsing 'byteCount' field " + _byteCountErr.Error())
    }

    // Array field (value)
    // Count array
    value := make([]int8, byteCount)
    for curItem := uint16(0); curItem < uint16(byteCount); curItem++ {
        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'value' field " + _err.Error())
        }
        value[curItem] = _item
    }

    // Create a partially initialized instance
    _child := &ModbusPDUReadWriteMultipleHoldingRegistersRequest{
        ReadStartingAddress: readStartingAddress,
        ReadQuantity: readQuantity,
        WriteStartingAddress: writeStartingAddress,
        WriteQuantity: writeQuantity,
        Value: value,
        Parent: &ModbusPDU{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (readStartingAddress)
    readStartingAddress := uint16(m.ReadStartingAddress)
    _readStartingAddressErr := io.WriteUint16(16, (readStartingAddress))
    if _readStartingAddressErr != nil {
        return errors.New("Error serializing 'readStartingAddress' field " + _readStartingAddressErr.Error())
    }

    // Simple Field (readQuantity)
    readQuantity := uint16(m.ReadQuantity)
    _readQuantityErr := io.WriteUint16(16, (readQuantity))
    if _readQuantityErr != nil {
        return errors.New("Error serializing 'readQuantity' field " + _readQuantityErr.Error())
    }

    // Simple Field (writeStartingAddress)
    writeStartingAddress := uint16(m.WriteStartingAddress)
    _writeStartingAddressErr := io.WriteUint16(16, (writeStartingAddress))
    if _writeStartingAddressErr != nil {
        return errors.New("Error serializing 'writeStartingAddress' field " + _writeStartingAddressErr.Error())
    }

    // Simple Field (writeQuantity)
    writeQuantity := uint16(m.WriteQuantity)
    _writeQuantityErr := io.WriteUint16(16, (writeQuantity))
    if _writeQuantityErr != nil {
        return errors.New("Error serializing 'writeQuantity' field " + _writeQuantityErr.Error())
    }

    // Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    byteCount := uint8(uint8(len(m.Value)))
    _byteCountErr := io.WriteUint8(8, (byteCount))
    if _byteCountErr != nil {
        return errors.New("Error serializing 'byteCount' field " + _byteCountErr.Error())
    }

    // Array Field (value)
    if m.Value != nil {
        for _, _element := range m.Value {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'value' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "readStartingAddress":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ReadStartingAddress = data
            case "readQuantity":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ReadQuantity = data
            case "writeStartingAddress":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.WriteStartingAddress = data
            case "writeQuantity":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.WriteQuantity = data
            case "value":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.Value = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.ReadStartingAddress, xml.StartElement{Name: xml.Name{Local: "readStartingAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ReadQuantity, xml.StartElement{Name: xml.Name{Local: "readQuantity"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.WriteStartingAddress, xml.StartElement{Name: xml.Name{Local: "writeStartingAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.WriteQuantity, xml.StartElement{Name: xml.Name{Local: "writeQuantity"}}); err != nil {
        return err
    }
    _encodedValue := make([]byte, base64.StdEncoding.EncodedLen(len(m.Value)))
    base64.StdEncoding.Encode(_encodedValue, utils.Int8ArrayToByteArray(m.Value))
    if err := e.EncodeElement(_encodedValue, xml.StartElement{Name: xml.Name{Local: "value"}}); err != nil {
        return err
    }
    return nil
}

