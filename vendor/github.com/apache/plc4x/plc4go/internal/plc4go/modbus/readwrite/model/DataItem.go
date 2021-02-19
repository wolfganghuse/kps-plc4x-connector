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
	"errors"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/values"
	api "github.com/apache/plc4x/plc4go/pkg/plc4go/values"
)

func DataItemParse(io *utils.ReadBuffer, dataType ModbusDataType, numberOfValues uint16) (api.PlcValue, error) {
	switch {
	case dataType == ModbusDataType_BOOL && numberOfValues == 1: // BOOL

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(7); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (value)
		value, _valueErr := io.ReadBit()
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcBOOL(value), nil
	case dataType == ModbusDataType_BOOL: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadBit()
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcBOOL(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_BYTE && numberOfValues == 1: // BitString

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcBitString(value), nil
	case dataType == ModbusDataType_BYTE: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadUint8(8)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcUSINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_WORD && numberOfValues == 1: // BitString

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcBitString(value), nil
	case dataType == ModbusDataType_WORD: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadUint16(16)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcUINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_DWORD && numberOfValues == 1: // BitString

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcBitString(value), nil
	case dataType == ModbusDataType_DWORD: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadUint32(32)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcUDINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_LWORD && numberOfValues == 1: // BitString

		// Simple Field (value)
		value, _valueErr := io.ReadUint64(64)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcBitString(value), nil
	case dataType == ModbusDataType_LWORD: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadUint64(64)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcULINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_SINT && numberOfValues == 1: // SINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcSINT(value), nil
	case dataType == ModbusDataType_SINT: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadInt8(8)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcSINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_INT && numberOfValues == 1: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case dataType == ModbusDataType_INT: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadInt16(16)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_DINT && numberOfValues == 1: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case dataType == ModbusDataType_DINT: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadInt32(32)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcDINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_LINT && numberOfValues == 1: // LINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt64(64)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcLINT(value), nil
	case dataType == ModbusDataType_LINT: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadInt64(64)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcLINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_USINT && numberOfValues == 1: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case dataType == ModbusDataType_USINT: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadUint8(8)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcUSINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_UINT && numberOfValues == 1: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case dataType == ModbusDataType_UINT: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadUint16(16)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcUINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_UDINT && numberOfValues == 1: // UDINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUDINT(value), nil
	case dataType == ModbusDataType_UDINT: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadUint32(32)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcUDINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_ULINT && numberOfValues == 1: // ULINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint64(64)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcULINT(value), nil
	case dataType == ModbusDataType_ULINT: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadUint64(64)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcULINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_REAL && numberOfValues == 1: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case dataType == ModbusDataType_REAL: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadFloat32(true, 8, 23)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcREAL(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_LREAL && numberOfValues == 1: // LREAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat64(true, 11, 52)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcLREAL(value), nil
	case dataType == ModbusDataType_LREAL: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadFloat64(true, 11, 52)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcLREAL(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_CHAR && numberOfValues == 1: // CHAR

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcCHAR(value), nil
	case dataType == ModbusDataType_CHAR: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadUint8(8)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcUSINT(_item))
		}
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_WCHAR && numberOfValues == 1: // WCHAR

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcWCHAR(value), nil
	case dataType == ModbusDataType_WCHAR: // List

		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := io.ReadUint16(16)
			if _itemErr != nil {
				return nil, errors.New("Error parsing 'value' field " + _itemErr.Error())
			}
			value = append(value, values.NewPlcUINT(_item))
		}
		return values.NewPlcList(value), nil
	}
	return nil, errors.New("unsupported type")
}

func DataItemSerialize(io *utils.WriteBuffer, value api.PlcValue, dataType ModbusDataType, numberOfValues uint16) error {
	switch {
	case dataType == ModbusDataType_BOOL && numberOfValues == 1: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_BOOL: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteBit(value.GetIndex(i).GetBool())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_BYTE && numberOfValues == 1: // BitString

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_BYTE: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteUint8(8, value.GetIndex(i).GetUint8())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_WORD && numberOfValues == 1: // BitString

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_WORD: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteUint16(16, value.GetIndex(i).GetUint16())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_DWORD && numberOfValues == 1: // BitString

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_DWORD: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteUint32(32, value.GetIndex(i).GetUint32())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_LWORD && numberOfValues == 1: // BitString

		// Simple Field (value)
		if _err := io.WriteUint64(64, value.GetUint64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_LWORD: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteUint64(64, value.GetIndex(i).GetUint64())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_SINT && numberOfValues == 1: // SINT

		// Simple Field (value)
		if _err := io.WriteInt8(8, value.GetInt8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_SINT: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteInt8(8, value.GetIndex(i).GetInt8())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_INT && numberOfValues == 1: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_INT: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteInt16(16, value.GetIndex(i).GetInt16())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_DINT && numberOfValues == 1: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_DINT: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteInt32(32, value.GetIndex(i).GetInt32())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_LINT && numberOfValues == 1: // LINT

		// Simple Field (value)
		if _err := io.WriteInt64(64, value.GetInt64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_LINT: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteInt64(64, value.GetIndex(i).GetInt64())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_USINT && numberOfValues == 1: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_USINT: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteUint8(8, value.GetIndex(i).GetUint8())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_UINT && numberOfValues == 1: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_UINT: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteUint16(16, value.GetIndex(i).GetUint16())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_UDINT && numberOfValues == 1: // UDINT

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_UDINT: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteUint32(32, value.GetIndex(i).GetUint32())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_ULINT && numberOfValues == 1: // ULINT

		// Simple Field (value)
		if _err := io.WriteUint64(64, value.GetUint64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_ULINT: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteUint64(64, value.GetIndex(i).GetUint64())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_REAL && numberOfValues == 1: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_REAL: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteFloat32(32, value.GetIndex(i).GetFloat32())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_LREAL && numberOfValues == 1: // LREAL

		// Simple Field (value)
		if _err := io.WriteFloat64(64, value.GetFloat64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_LREAL: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteFloat64(64, value.GetIndex(i).GetFloat64())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_CHAR && numberOfValues == 1: // CHAR

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_CHAR: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteUint8(8, value.GetIndex(i).GetUint8())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	case dataType == ModbusDataType_WCHAR && numberOfValues == 1: // WCHAR

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case dataType == ModbusDataType_WCHAR: // List

		// Array Field (value)
		for i := uint32(0); i < uint32(numberOfValues); i++ {
			_itemErr := io.WriteUint16(16, value.GetIndex(i).GetUint16())
			if _itemErr != nil {
				return errors.New("Error serializing 'value' field " + _itemErr.Error())
			}
		}
	default:

		return errors.New("unsupported type")
	}
	return nil
}
