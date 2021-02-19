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
type ModbusPDU struct {
    Child IModbusPDUChild
    IModbusPDU
    IModbusPDUParent
}

// The corresponding interface
type IModbusPDU interface {
    ErrorFlag() bool
    FunctionFlag() uint8
    Response() bool
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type IModbusPDUParent interface {
    SerializeParent(io utils.WriteBuffer, child IModbusPDU, serializeChildFunction func() error) error
    GetTypeName() string
}

type IModbusPDUChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *ModbusPDU)
    GetTypeName() string
    IModbusPDU
}

func NewModbusPDU() *ModbusPDU {
    return &ModbusPDU{}
}

func CastModbusPDU(structType interface{}) *ModbusPDU {
    castFunc := func(typ interface{}) *ModbusPDU {
        if casted, ok := typ.(ModbusPDU); ok {
            return &casted
        }
        if casted, ok := typ.(*ModbusPDU); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ModbusPDU) GetTypeName() string {
    return "ModbusPDU"
}

func (m *ModbusPDU) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Discriminator Field (errorFlag)
    lengthInBits += 1

    // Discriminator Field (functionFlag)
    lengthInBits += 7

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *ModbusPDU) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUParse(io *utils.ReadBuffer, response bool) (*ModbusPDU, error) {

    // Discriminator Field (errorFlag) (Used as input to a switch field)
    errorFlag, _errorFlagErr := io.ReadBit()
    if _errorFlagErr != nil {
        return nil, errors.New("Error parsing 'errorFlag' field " + _errorFlagErr.Error())
    }

    // Discriminator Field (functionFlag) (Used as input to a switch field)
    functionFlag, _functionFlagErr := io.ReadUint8(7)
    if _functionFlagErr != nil {
        return nil, errors.New("Error parsing 'functionFlag' field " + _functionFlagErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *ModbusPDU
    var typeSwitchError error
    switch {
    case errorFlag == true:
        _parent, typeSwitchError = ModbusPDUErrorParse(io)
    case errorFlag == false && functionFlag == 0x02 && response == false:
        _parent, typeSwitchError = ModbusPDUReadDiscreteInputsRequestParse(io)
    case errorFlag == false && functionFlag == 0x02 && response == true:
        _parent, typeSwitchError = ModbusPDUReadDiscreteInputsResponseParse(io)
    case errorFlag == false && functionFlag == 0x01 && response == false:
        _parent, typeSwitchError = ModbusPDUReadCoilsRequestParse(io)
    case errorFlag == false && functionFlag == 0x01 && response == true:
        _parent, typeSwitchError = ModbusPDUReadCoilsResponseParse(io)
    case errorFlag == false && functionFlag == 0x05 && response == false:
        _parent, typeSwitchError = ModbusPDUWriteSingleCoilRequestParse(io)
    case errorFlag == false && functionFlag == 0x05 && response == true:
        _parent, typeSwitchError = ModbusPDUWriteSingleCoilResponseParse(io)
    case errorFlag == false && functionFlag == 0x0F && response == false:
        _parent, typeSwitchError = ModbusPDUWriteMultipleCoilsRequestParse(io)
    case errorFlag == false && functionFlag == 0x0F && response == true:
        _parent, typeSwitchError = ModbusPDUWriteMultipleCoilsResponseParse(io)
    case errorFlag == false && functionFlag == 0x04 && response == false:
        _parent, typeSwitchError = ModbusPDUReadInputRegistersRequestParse(io)
    case errorFlag == false && functionFlag == 0x04 && response == true:
        _parent, typeSwitchError = ModbusPDUReadInputRegistersResponseParse(io)
    case errorFlag == false && functionFlag == 0x03 && response == false:
        _parent, typeSwitchError = ModbusPDUReadHoldingRegistersRequestParse(io)
    case errorFlag == false && functionFlag == 0x03 && response == true:
        _parent, typeSwitchError = ModbusPDUReadHoldingRegistersResponseParse(io)
    case errorFlag == false && functionFlag == 0x06 && response == false:
        _parent, typeSwitchError = ModbusPDUWriteSingleRegisterRequestParse(io)
    case errorFlag == false && functionFlag == 0x06 && response == true:
        _parent, typeSwitchError = ModbusPDUWriteSingleRegisterResponseParse(io)
    case errorFlag == false && functionFlag == 0x10 && response == false:
        _parent, typeSwitchError = ModbusPDUWriteMultipleHoldingRegistersRequestParse(io)
    case errorFlag == false && functionFlag == 0x10 && response == true:
        _parent, typeSwitchError = ModbusPDUWriteMultipleHoldingRegistersResponseParse(io)
    case errorFlag == false && functionFlag == 0x17 && response == false:
        _parent, typeSwitchError = ModbusPDUReadWriteMultipleHoldingRegistersRequestParse(io)
    case errorFlag == false && functionFlag == 0x17 && response == true:
        _parent, typeSwitchError = ModbusPDUReadWriteMultipleHoldingRegistersResponseParse(io)
    case errorFlag == false && functionFlag == 0x16 && response == false:
        _parent, typeSwitchError = ModbusPDUMaskWriteHoldingRegisterRequestParse(io)
    case errorFlag == false && functionFlag == 0x16 && response == true:
        _parent, typeSwitchError = ModbusPDUMaskWriteHoldingRegisterResponseParse(io)
    case errorFlag == false && functionFlag == 0x18 && response == false:
        _parent, typeSwitchError = ModbusPDUReadFifoQueueRequestParse(io)
    case errorFlag == false && functionFlag == 0x18 && response == true:
        _parent, typeSwitchError = ModbusPDUReadFifoQueueResponseParse(io)
    case errorFlag == false && functionFlag == 0x14 && response == false:
        _parent, typeSwitchError = ModbusPDUReadFileRecordRequestParse(io)
    case errorFlag == false && functionFlag == 0x14 && response == true:
        _parent, typeSwitchError = ModbusPDUReadFileRecordResponseParse(io)
    case errorFlag == false && functionFlag == 0x15 && response == false:
        _parent, typeSwitchError = ModbusPDUWriteFileRecordRequestParse(io)
    case errorFlag == false && functionFlag == 0x15 && response == true:
        _parent, typeSwitchError = ModbusPDUWriteFileRecordResponseParse(io)
    case errorFlag == false && functionFlag == 0x07 && response == false:
        _parent, typeSwitchError = ModbusPDUReadExceptionStatusRequestParse(io)
    case errorFlag == false && functionFlag == 0x07 && response == true:
        _parent, typeSwitchError = ModbusPDUReadExceptionStatusResponseParse(io)
    case errorFlag == false && functionFlag == 0x08 && response == false:
        _parent, typeSwitchError = ModbusPDUDiagnosticRequestParse(io)
    case errorFlag == false && functionFlag == 0x08 && response == true:
        _parent, typeSwitchError = ModbusPDUDiagnosticResponseParse(io)
    case errorFlag == false && functionFlag == 0x0B && response == false:
        _parent, typeSwitchError = ModbusPDUGetComEventCounterRequestParse(io)
    case errorFlag == false && functionFlag == 0x0B && response == true:
        _parent, typeSwitchError = ModbusPDUGetComEventCounterResponseParse(io)
    case errorFlag == false && functionFlag == 0x0C && response == false:
        _parent, typeSwitchError = ModbusPDUGetComEventLogRequestParse(io)
    case errorFlag == false && functionFlag == 0x0C && response == true:
        _parent, typeSwitchError = ModbusPDUGetComEventLogResponseParse(io)
    case errorFlag == false && functionFlag == 0x11 && response == false:
        _parent, typeSwitchError = ModbusPDUReportServerIdRequestParse(io)
    case errorFlag == false && functionFlag == 0x11 && response == true:
        _parent, typeSwitchError = ModbusPDUReportServerIdResponseParse(io)
    case errorFlag == false && functionFlag == 0x2B && response == false:
        _parent, typeSwitchError = ModbusPDUReadDeviceIdentificationRequestParse(io)
    case errorFlag == false && functionFlag == 0x2B && response == true:
        _parent, typeSwitchError = ModbusPDUReadDeviceIdentificationResponseParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent)
    return _parent, nil
}

func (m *ModbusPDU) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *ModbusPDU) SerializeParent(io utils.WriteBuffer, child IModbusPDU, serializeChildFunction func() error) error {

    // Discriminator Field (errorFlag) (Used as input to a switch field)
    errorFlag := bool(child.ErrorFlag())
    _errorFlagErr := io.WriteBit((errorFlag))
    if _errorFlagErr != nil {
        return errors.New("Error serializing 'errorFlag' field " + _errorFlagErr.Error())
    }

    // Discriminator Field (functionFlag) (Used as input to a switch field)
    functionFlag := uint8(child.FunctionFlag())
    _functionFlagErr := io.WriteUint8(7, (functionFlag))
    if _functionFlagErr != nil {
        return errors.New("Error serializing 'functionFlag' field " + _functionFlagErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *ModbusPDU) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUError":
                        var dt *ModbusPDUError
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUError)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadDiscreteInputsRequest":
                        var dt *ModbusPDUReadDiscreteInputsRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadDiscreteInputsRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadDiscreteInputsResponse":
                        var dt *ModbusPDUReadDiscreteInputsResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadDiscreteInputsResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadCoilsRequest":
                        var dt *ModbusPDUReadCoilsRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadCoilsRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadCoilsResponse":
                        var dt *ModbusPDUReadCoilsResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadCoilsResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteSingleCoilRequest":
                        var dt *ModbusPDUWriteSingleCoilRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUWriteSingleCoilRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteSingleCoilResponse":
                        var dt *ModbusPDUWriteSingleCoilResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUWriteSingleCoilResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteMultipleCoilsRequest":
                        var dt *ModbusPDUWriteMultipleCoilsRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUWriteMultipleCoilsRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteMultipleCoilsResponse":
                        var dt *ModbusPDUWriteMultipleCoilsResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUWriteMultipleCoilsResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadInputRegistersRequest":
                        var dt *ModbusPDUReadInputRegistersRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadInputRegistersRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadInputRegistersResponse":
                        var dt *ModbusPDUReadInputRegistersResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadInputRegistersResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadHoldingRegistersRequest":
                        var dt *ModbusPDUReadHoldingRegistersRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadHoldingRegistersRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadHoldingRegistersResponse":
                        var dt *ModbusPDUReadHoldingRegistersResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadHoldingRegistersResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteSingleRegisterRequest":
                        var dt *ModbusPDUWriteSingleRegisterRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUWriteSingleRegisterRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteSingleRegisterResponse":
                        var dt *ModbusPDUWriteSingleRegisterResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUWriteSingleRegisterResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteMultipleHoldingRegistersRequest":
                        var dt *ModbusPDUWriteMultipleHoldingRegistersRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUWriteMultipleHoldingRegistersRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteMultipleHoldingRegistersResponse":
                        var dt *ModbusPDUWriteMultipleHoldingRegistersResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUWriteMultipleHoldingRegistersResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadWriteMultipleHoldingRegistersRequest":
                        var dt *ModbusPDUReadWriteMultipleHoldingRegistersRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadWriteMultipleHoldingRegistersRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadWriteMultipleHoldingRegistersResponse":
                        var dt *ModbusPDUReadWriteMultipleHoldingRegistersResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadWriteMultipleHoldingRegistersResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUMaskWriteHoldingRegisterRequest":
                        var dt *ModbusPDUMaskWriteHoldingRegisterRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUMaskWriteHoldingRegisterRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUMaskWriteHoldingRegisterResponse":
                        var dt *ModbusPDUMaskWriteHoldingRegisterResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUMaskWriteHoldingRegisterResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadFifoQueueRequest":
                        var dt *ModbusPDUReadFifoQueueRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadFifoQueueRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadFifoQueueResponse":
                        var dt *ModbusPDUReadFifoQueueResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadFifoQueueResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadFileRecordRequest":
                        var dt *ModbusPDUReadFileRecordRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadFileRecordRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadFileRecordResponse":
                        var dt *ModbusPDUReadFileRecordResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadFileRecordResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteFileRecordRequest":
                        var dt *ModbusPDUWriteFileRecordRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUWriteFileRecordRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUWriteFileRecordResponse":
                        var dt *ModbusPDUWriteFileRecordResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUWriteFileRecordResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadExceptionStatusRequest":
                        var dt *ModbusPDUReadExceptionStatusRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadExceptionStatusRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadExceptionStatusResponse":
                        var dt *ModbusPDUReadExceptionStatusResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadExceptionStatusResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUDiagnosticRequest":
                        var dt *ModbusPDUDiagnosticRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUDiagnosticRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUDiagnosticResponse":
                        var dt *ModbusPDUDiagnosticResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUDiagnosticResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUGetComEventCounterRequest":
                        var dt *ModbusPDUGetComEventCounterRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUGetComEventCounterRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUGetComEventCounterResponse":
                        var dt *ModbusPDUGetComEventCounterResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUGetComEventCounterResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUGetComEventLogRequest":
                        var dt *ModbusPDUGetComEventLogRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUGetComEventLogRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUGetComEventLogResponse":
                        var dt *ModbusPDUGetComEventLogResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUGetComEventLogResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReportServerIdRequest":
                        var dt *ModbusPDUReportServerIdRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReportServerIdRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReportServerIdResponse":
                        var dt *ModbusPDUReportServerIdResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReportServerIdResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadDeviceIdentificationRequest":
                        var dt *ModbusPDUReadDeviceIdentificationRequest
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadDeviceIdentificationRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadDeviceIdentificationResponse":
                        var dt *ModbusPDUReadDeviceIdentificationResponse
                        if m.Child != nil {
                            dt = m.Child.(*ModbusPDUReadDeviceIdentificationResponse)
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

func (m *ModbusPDU) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := reflect.TypeOf(m.Child).String()
    className = "org.apache.plc4x.java.modbus.readwrite." + className[strings.LastIndex(className, ".") + 1:]
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

