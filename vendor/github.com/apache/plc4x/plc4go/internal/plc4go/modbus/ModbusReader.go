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
package modbus

import (
	"errors"
	modbusModel "github.com/apache/plc4x/plc4go/internal/plc4go/modbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi"
	plc4goModel "github.com/apache/plc4x/plc4go/internal/plc4go/spi/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/apache/plc4x/plc4go/pkg/plc4go/model"
	"github.com/apache/plc4x/plc4go/pkg/plc4go/values"
	"math"
	"sync/atomic"
	"time"
)

type ModbusReader struct {
	transactionIdentifier int32
	unitIdentifier        uint8
	messageCodec          spi.MessageCodec
	spi.PlcReader
}

func NewModbusReader(unitIdentifier uint8, messageCodec spi.MessageCodec) *ModbusReader {
	return &ModbusReader{
		transactionIdentifier: 0,
		unitIdentifier:        unitIdentifier,
		messageCodec:          messageCodec,
	}
}

func (m *ModbusReader) Read(readRequest model.PlcReadRequest) <-chan model.PlcReadRequestResult {
	result := make(chan model.PlcReadRequestResult)
	// If we are requesting only one field, use a
	if len(readRequest.GetFieldNames()) == 1 {
		fieldName := readRequest.GetFieldNames()[0]
		field := readRequest.GetField(fieldName)
		modbusField, err := CastToModbusFieldFromPlcField(field)
		if err != nil {
			result <- model.PlcReadRequestResult{
				Request:  readRequest,
				Response: nil,
				Err:      errors.New("invalid field item type"),
			}
			return result
		}
		numWords := uint16(math.Ceil(float64(modbusField.Quantity*uint16(modbusField.Datatype.DataTypeSize())) / float64(2)))
		var pdu *modbusModel.ModbusPDU = nil
		switch modbusField.FieldType {
		case MODBUS_FIELD_COIL:
			pdu = modbusModel.NewModbusPDUReadCoilsRequest(modbusField.Address, modbusField.Quantity)
		case MODBUS_FIELD_DISCRETE_INPUT:
			pdu = modbusModel.NewModbusPDUReadDiscreteInputsRequest(modbusField.Address, modbusField.Quantity)
		case MODBUS_FIELD_INPUT_REGISTER:
			pdu = modbusModel.NewModbusPDUReadInputRegistersRequest(modbusField.Address, numWords)
		case MODBUS_FIELD_HOLDING_REGISTER:
			pdu = modbusModel.NewModbusPDUReadHoldingRegistersRequest(modbusField.Address, numWords)
		case MODBUS_FIELD_EXTENDED_REGISTER:
			result <- model.PlcReadRequestResult{
				Request:  readRequest,
				Response: nil,
				Err:      errors.New("modbus currently doesn't support extended register requests"),
			}
			return result
		default:
			result <- model.PlcReadRequestResult{
				Request:  readRequest,
				Response: nil,
				Err:      errors.New("unsupported field type"),
			}
			return result
		}

		// Calculate a new transaction identifier
		transactionIdentifier := atomic.AddInt32(&m.transactionIdentifier, 1)
		if transactionIdentifier > math.MaxUint8 {
			transactionIdentifier = 1
			atomic.StoreInt32(&m.transactionIdentifier, 1)
		}

		// Assemble the finished ADU
		requestAdu := modbusModel.ModbusTcpADU{
			TransactionIdentifier: uint16(transactionIdentifier),
			UnitIdentifier:        m.unitIdentifier,
			Pdu:                   pdu,
		}

		// Send the ADU over the wire
		err = m.messageCodec.SendRequest(
			requestAdu,
			func(message interface{}) bool {
				responseAdu := modbusModel.CastModbusTcpADU(message)
				return responseAdu.TransactionIdentifier == uint16(transactionIdentifier) &&
					responseAdu.UnitIdentifier == requestAdu.UnitIdentifier
			},
			func(message interface{}) error {
				// Convert the response into an ADU
				responseAdu := modbusModel.CastModbusTcpADU(message)
				// Convert the modbus response into a PLC4X response
				readResponse, err := m.ToPlc4xReadResponse(*responseAdu, readRequest)

				if err != nil {
					result <- model.PlcReadRequestResult{
						Request: readRequest,
						Err:     errors.New("Error decoding response: " + err.Error()),
					}
				} else {
					result <- model.PlcReadRequestResult{
						Request:  readRequest,
						Response: readResponse,
					}
				}
				return nil
			},
			func(err error) error {
				result <- model.PlcReadRequestResult{
					Request: readRequest,
					Err:     errors.New("got timeout while waiting for response"),
				}
				return nil
			},
			time.Second*1)
		if err != nil {
			result <- model.PlcReadRequestResult{
				Request:  readRequest,
				Response: nil,
				Err:      errors.New("error sending message: " + err.Error()),
			}
		}
	} else {
		result <- model.PlcReadRequestResult{
			Request:  readRequest,
			Response: nil,
			Err:      errors.New("modbus only supports single-item requests"),
		}
	}
	return result
}

func (m *ModbusReader) ToPlc4xReadResponse(responseAdu modbusModel.ModbusTcpADU, readRequest model.PlcReadRequest) (model.PlcReadResponse, error) {
	var data []uint8
	switch responseAdu.Pdu.Child.(type) {
	case *modbusModel.ModbusPDUReadDiscreteInputsResponse:
		pdu := modbusModel.CastModbusPDUReadDiscreteInputsResponse(responseAdu.Pdu)
		data = utils.Int8ArrayToUint8Array(pdu.Value)
		// Pure Boolean ...
	case *modbusModel.ModbusPDUReadCoilsResponse:
		pdu := modbusModel.CastModbusPDUReadCoilsResponse(&responseAdu.Pdu)
		data = utils.Int8ArrayToUint8Array(pdu.Value)
		// Pure Boolean ...
	case *modbusModel.ModbusPDUReadInputRegistersResponse:
		pdu := modbusModel.CastModbusPDUReadInputRegistersResponse(responseAdu.Pdu)
		data = utils.Int8ArrayToUint8Array(pdu.Value)
		// DataIo ...
	case *modbusModel.ModbusPDUReadHoldingRegistersResponse:
		pdu := modbusModel.CastModbusPDUReadHoldingRegistersResponse(responseAdu.Pdu)
		data = utils.Int8ArrayToUint8Array(pdu.Value)
	case *modbusModel.ModbusPDUError:
		return nil, errors.New("got an error from remote")
	default:
		return nil, errors.New("unsupported response type")
	}

	// Get the field from the request
	fieldName := readRequest.GetFieldNames()[0]
	field, err := CastToModbusFieldFromPlcField(readRequest.GetField(fieldName))
	if err != nil {
		return nil, errors.New("error casting to modbus-field")
	}

	// Decode the data according to the information from the request
	rb := utils.NewReadBuffer(data)
	value, err := modbusModel.DataItemParse(rb, field.Datatype, field.Quantity)
	if err != nil {
		return nil, err
	}
	responseCodes := map[string]model.PlcResponseCode{}
	values := map[string]values.PlcValue{}
	values[fieldName] = value
	responseCodes[fieldName] = model.PlcResponseCode_OK

	// Return the response
	return plc4goModel.NewDefaultPlcReadResponse(readRequest, responseCodes, values), nil
}
