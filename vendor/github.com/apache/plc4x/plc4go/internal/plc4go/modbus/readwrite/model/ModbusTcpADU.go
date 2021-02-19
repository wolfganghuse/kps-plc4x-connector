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
    "strconv"
)

// Constant values.
const ModbusTcpADU_PROTOCOLIDENTIFIER uint16 = 0x0000

// The data-structure of this message
type ModbusTcpADU struct {
    TransactionIdentifier uint16
    UnitIdentifier uint8
    Pdu *ModbusPDU
    IModbusTcpADU
}

// The corresponding interface
type IModbusTcpADU interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

func NewModbusTcpADU(transactionIdentifier uint16, unitIdentifier uint8, pdu *ModbusPDU) *ModbusTcpADU {
    return &ModbusTcpADU{TransactionIdentifier: transactionIdentifier, UnitIdentifier: unitIdentifier, Pdu: pdu}
}

func CastModbusTcpADU(structType interface{}) *ModbusTcpADU {
    castFunc := func(typ interface{}) *ModbusTcpADU {
        if casted, ok := typ.(ModbusTcpADU); ok {
            return &casted
        }
        if casted, ok := typ.(*ModbusTcpADU); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ModbusTcpADU) GetTypeName() string {
    return "ModbusTcpADU"
}

func (m *ModbusTcpADU) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (transactionIdentifier)
    lengthInBits += 16

    // Const Field (protocolIdentifier)
    lengthInBits += 16

    // Implicit Field (length)
    lengthInBits += 16

    // Simple field (unitIdentifier)
    lengthInBits += 8

    // Simple field (pdu)
    lengthInBits += m.Pdu.LengthInBits()

    return lengthInBits
}

func (m *ModbusTcpADU) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusTcpADUParse(io *utils.ReadBuffer, response bool) (*ModbusTcpADU, error) {

    // Simple Field (transactionIdentifier)
    transactionIdentifier, _transactionIdentifierErr := io.ReadUint16(16)
    if _transactionIdentifierErr != nil {
        return nil, errors.New("Error parsing 'transactionIdentifier' field " + _transactionIdentifierErr.Error())
    }

    // Const Field (protocolIdentifier)
    protocolIdentifier, _protocolIdentifierErr := io.ReadUint16(16)
    if _protocolIdentifierErr != nil {
        return nil, errors.New("Error parsing 'protocolIdentifier' field " + _protocolIdentifierErr.Error())
    }
    if protocolIdentifier != ModbusTcpADU_PROTOCOLIDENTIFIER {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(ModbusTcpADU_PROTOCOLIDENTIFIER)) + " but got " + strconv.Itoa(int(protocolIdentifier)))
    }

    // Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    _, _lengthErr := io.ReadUint16(16)
    if _lengthErr != nil {
        return nil, errors.New("Error parsing 'length' field " + _lengthErr.Error())
    }

    // Simple Field (unitIdentifier)
    unitIdentifier, _unitIdentifierErr := io.ReadUint8(8)
    if _unitIdentifierErr != nil {
        return nil, errors.New("Error parsing 'unitIdentifier' field " + _unitIdentifierErr.Error())
    }

    // Simple Field (pdu)
    pdu, _pduErr := ModbusPDUParse(io, response)
    if _pduErr != nil {
        return nil, errors.New("Error parsing 'pdu' field " + _pduErr.Error())
    }

    // Create the instance
    return NewModbusTcpADU(transactionIdentifier, unitIdentifier, pdu), nil
}

func (m *ModbusTcpADU) Serialize(io utils.WriteBuffer) error {

    // Simple Field (transactionIdentifier)
    transactionIdentifier := uint16(m.TransactionIdentifier)
    _transactionIdentifierErr := io.WriteUint16(16, (transactionIdentifier))
    if _transactionIdentifierErr != nil {
        return errors.New("Error serializing 'transactionIdentifier' field " + _transactionIdentifierErr.Error())
    }

    // Const Field (protocolIdentifier)
    _protocolIdentifierErr := io.WriteUint16(16, 0x0000)
    if _protocolIdentifierErr != nil {
        return errors.New("Error serializing 'protocolIdentifier' field " + _protocolIdentifierErr.Error())
    }

    // Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    length := uint16(uint16(m.Pdu.LengthInBytes()) + uint16(uint16(1)))
    _lengthErr := io.WriteUint16(16, (length))
    if _lengthErr != nil {
        return errors.New("Error serializing 'length' field " + _lengthErr.Error())
    }

    // Simple Field (unitIdentifier)
    unitIdentifier := uint8(m.UnitIdentifier)
    _unitIdentifierErr := io.WriteUint8(8, (unitIdentifier))
    if _unitIdentifierErr != nil {
        return errors.New("Error serializing 'unitIdentifier' field " + _unitIdentifierErr.Error())
    }

    // Simple Field (pdu)
    _pduErr := m.Pdu.Serialize(io)
    if _pduErr != nil {
        return errors.New("Error serializing 'pdu' field " + _pduErr.Error())
    }

    return nil
}

func (m *ModbusTcpADU) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "transactionIdentifier":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.TransactionIdentifier = data
            case "unitIdentifier":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.UnitIdentifier = data
            case "pdu":
                var dt *ModbusPDU
                if err := d.DecodeElement(&dt, &tok); err != nil {
                    return err
                }
                m.Pdu = dt
            }
        }
    }
}

func (m *ModbusTcpADU) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := "org.apache.plc4x.java.modbus.readwrite.ModbusTcpADU"
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.TransactionIdentifier, xml.StartElement{Name: xml.Name{Local: "transactionIdentifier"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.UnitIdentifier, xml.StartElement{Name: xml.Name{Local: "unitIdentifier"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Pdu, xml.StartElement{Name: xml.Name{Local: "pdu"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

