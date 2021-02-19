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
	"time"
)

func KnxDatapointParse(io *utils.ReadBuffer, datapointType KnxDatapointType) (api.PlcValue, error) {
	switch {
	case datapointType == KnxDatapointType_BOOL: // BOOL

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
	case datapointType == KnxDatapointType_BYTE: // BYTE

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcBYTE(value), nil
	case datapointType == KnxDatapointType_WORD: // WORD

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcWORD(value), nil
	case datapointType == KnxDatapointType_DWORD: // DWORD

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDWORD(value), nil
	case datapointType == KnxDatapointType_LWORD: // LWORD

		// Simple Field (value)
		value, _valueErr := io.ReadUint64(64)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcLWORD(value), nil
	case datapointType == KnxDatapointType_USINT: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_SINT: // SINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcSINT(value), nil
	case datapointType == KnxDatapointType_UINT: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_INT: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_UDINT: // UDINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUDINT(value), nil
	case datapointType == KnxDatapointType_DINT: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_ULINT: // ULINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint64(64)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcULINT(value), nil
	case datapointType == KnxDatapointType_LINT: // LINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt64(64)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcLINT(value), nil
	case datapointType == KnxDatapointType_REAL: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_LREAL: // LREAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat64(true, 11, 52)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcLREAL(value), nil
	case datapointType == KnxDatapointType_CHAR: // CHAR

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcCHAR(value), nil
	case datapointType == KnxDatapointType_WCHAR: // WCHAR

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcWCHAR(value), nil
	case datapointType == KnxDatapointType_TIME: // TIME

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcTIME(value), nil
	case datapointType == KnxDatapointType_LTIME: // LTIME

		// Simple Field (value)
		value, _valueErr := io.ReadUint64(64)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcLTIME(value), nil
	case datapointType == KnxDatapointType_DATE: // DATE

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDATE(value), nil
	case datapointType == KnxDatapointType_TIME_OF_DAY: // TIME_OF_DAY

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcTIME_OF_DAY(value), nil
	case datapointType == KnxDatapointType_TOD: // TIME_OF_DAY

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcTIME_OF_DAY(value), nil
	case datapointType == KnxDatapointType_DATE_AND_TIME: // DATE_AND_TIME

		// Simple Field (year)
		year, _yearErr := io.ReadUint16(16)
		if _yearErr != nil {
			return nil, errors.New("Error parsing 'year' field " + _yearErr.Error())
		}

		// Simple Field (month)
		month, _monthErr := io.ReadUint8(8)
		if _monthErr != nil {
			return nil, errors.New("Error parsing 'month' field " + _monthErr.Error())
		}

		// Simple Field (day)
		day, _dayErr := io.ReadUint8(8)
		if _dayErr != nil {
			return nil, errors.New("Error parsing 'day' field " + _dayErr.Error())
		}

		// Simple Field (dayOfWeek)
		_, _dayOfWeekErr := io.ReadUint8(8)
		if _dayOfWeekErr != nil {
			return nil, errors.New("Error parsing 'dayOfWeek' field " + _dayOfWeekErr.Error())
		}

		// Simple Field (hour)
		hour, _hourErr := io.ReadUint8(8)
		if _hourErr != nil {
			return nil, errors.New("Error parsing 'hour' field " + _hourErr.Error())
		}

		// Simple Field (minutes)
		minutes, _minutesErr := io.ReadUint8(8)
		if _minutesErr != nil {
			return nil, errors.New("Error parsing 'minutes' field " + _minutesErr.Error())
		}

		// Simple Field (seconds)
		seconds, _secondsErr := io.ReadUint8(8)
		if _secondsErr != nil {
			return nil, errors.New("Error parsing 'seconds' field " + _secondsErr.Error())
		}

		// Simple Field (nanos)
		_, _nanosErr := io.ReadUint32(32)
		if _nanosErr != nil {
			return nil, errors.New("Error parsing 'nanos' field " + _nanosErr.Error())
		}
		value := time.Date(int(year), time.Month(month), int(day), int(hour), int(minutes), int(seconds), 0, nil)
		return values.NewPlcDATE_AND_TIME(value), nil
	case datapointType == KnxDatapointType_DT: // DATE_AND_TIME

		// Simple Field (year)
		year, _yearErr := io.ReadUint16(16)
		if _yearErr != nil {
			return nil, errors.New("Error parsing 'year' field " + _yearErr.Error())
		}

		// Simple Field (month)
		month, _monthErr := io.ReadUint8(8)
		if _monthErr != nil {
			return nil, errors.New("Error parsing 'month' field " + _monthErr.Error())
		}

		// Simple Field (day)
		day, _dayErr := io.ReadUint8(8)
		if _dayErr != nil {
			return nil, errors.New("Error parsing 'day' field " + _dayErr.Error())
		}

		// Simple Field (dayOfWeek)
		_, _dayOfWeekErr := io.ReadUint8(8)
		if _dayOfWeekErr != nil {
			return nil, errors.New("Error parsing 'dayOfWeek' field " + _dayOfWeekErr.Error())
		}

		// Simple Field (hour)
		hour, _hourErr := io.ReadUint8(8)
		if _hourErr != nil {
			return nil, errors.New("Error parsing 'hour' field " + _hourErr.Error())
		}

		// Simple Field (minutes)
		minutes, _minutesErr := io.ReadUint8(8)
		if _minutesErr != nil {
			return nil, errors.New("Error parsing 'minutes' field " + _minutesErr.Error())
		}

		// Simple Field (seconds)
		seconds, _secondsErr := io.ReadUint8(8)
		if _secondsErr != nil {
			return nil, errors.New("Error parsing 'seconds' field " + _secondsErr.Error())
		}

		// Simple Field (nanos)
		_, _nanosErr := io.ReadUint32(32)
		if _nanosErr != nil {
			return nil, errors.New("Error parsing 'nanos' field " + _nanosErr.Error())
		}
		value := time.Date(int(year), time.Month(month), int(day), int(hour), int(minutes), int(seconds), 0, nil)
		return values.NewPlcDATE_AND_TIME(value), nil
	case datapointType == KnxDatapointType_DPT_Switch: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Bool: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Enable: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Ramp: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Alarm: // BOOL

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
	case datapointType == KnxDatapointType_DPT_BinaryValue: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Step: // BOOL

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
	case datapointType == KnxDatapointType_DPT_UpDown: // BOOL

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
	case datapointType == KnxDatapointType_DPT_OpenClose: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Start: // BOOL

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
	case datapointType == KnxDatapointType_DPT_State: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Invert: // BOOL

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
	case datapointType == KnxDatapointType_DPT_DimSendStyle: // BOOL

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
	case datapointType == KnxDatapointType_DPT_InputSource: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Reset: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Ack: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Trigger: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Occupancy: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Window_Door: // BOOL

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
	case datapointType == KnxDatapointType_DPT_LogicalFunction: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Scene_AB: // BOOL

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
	case datapointType == KnxDatapointType_DPT_ShutterBlinds_Mode: // BOOL

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
	case datapointType == KnxDatapointType_DPT_DayNight: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Heat_Cool: // BOOL

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
	case datapointType == KnxDatapointType_DPT_Switch_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (on)
		on, _onErr := io.ReadBit()
		if _onErr != nil {
			return nil, errors.New("Error parsing 'on' field " + _onErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(on)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Bool_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (valueTrue)
		valueTrue, _valueTrueErr := io.ReadBit()
		if _valueTrueErr != nil {
			return nil, errors.New("Error parsing 'valueTrue' field " + _valueTrueErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(valueTrue)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Enable_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (enable)
		enable, _enableErr := io.ReadBit()
		if _enableErr != nil {
			return nil, errors.New("Error parsing 'enable' field " + _enableErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(enable)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Ramp_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (ramp)
		ramp, _rampErr := io.ReadBit()
		if _rampErr != nil {
			return nil, errors.New("Error parsing 'ramp' field " + _rampErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(ramp)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Alarm_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (alarm)
		alarm, _alarmErr := io.ReadBit()
		if _alarmErr != nil {
			return nil, errors.New("Error parsing 'alarm' field " + _alarmErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(alarm)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_BinaryValue_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (high)
		high, _highErr := io.ReadBit()
		if _highErr != nil {
			return nil, errors.New("Error parsing 'high' field " + _highErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(high)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Step_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (increase)
		increase, _increaseErr := io.ReadBit()
		if _increaseErr != nil {
			return nil, errors.New("Error parsing 'increase' field " + _increaseErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(increase)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Direction1_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (down)
		down, _downErr := io.ReadBit()
		if _downErr != nil {
			return nil, errors.New("Error parsing 'down' field " + _downErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(down)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Direction2_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (close)
		close, _closeErr := io.ReadBit()
		if _closeErr != nil {
			return nil, errors.New("Error parsing 'close' field " + _closeErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(close)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Start_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (start)
		start, _startErr := io.ReadBit()
		if _startErr != nil {
			return nil, errors.New("Error parsing 'start' field " + _startErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(start)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_State_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (active)
		active, _activeErr := io.ReadBit()
		if _activeErr != nil {
			return nil, errors.New("Error parsing 'active' field " + _activeErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(active)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Invert_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (control)
		control, _controlErr := io.ReadBit()
		if _controlErr != nil {
			return nil, errors.New("Error parsing 'control' field " + _controlErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(control)

		// Simple Field (inverted)
		inverted, _invertedErr := io.ReadBit()
		if _invertedErr != nil {
			return nil, errors.New("Error parsing 'inverted' field " + _invertedErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(inverted)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Control_Dimming: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (increase)
		increase, _increaseErr := io.ReadBit()
		if _increaseErr != nil {
			return nil, errors.New("Error parsing 'increase' field " + _increaseErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(increase)

		// Simple Field (stepcode)
		stepcode, _stepcodeErr := io.ReadUint8(3)
		if _stepcodeErr != nil {
			return nil, errors.New("Error parsing 'stepcode' field " + _stepcodeErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepcode)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Control_Blinds: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (down)
		down, _downErr := io.ReadBit()
		if _downErr != nil {
			return nil, errors.New("Error parsing 'down' field " + _downErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(down)

		// Simple Field (stepcode)
		stepcode, _stepcodeErr := io.ReadUint8(3)
		if _stepcodeErr != nil {
			return nil, errors.New("Error parsing 'stepcode' field " + _stepcodeErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepcode)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Char_ASCII: // STRING

		// Simple Field (value)
		value, _valueErr := io.ReadString(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcSTRING(value), nil
	case datapointType == KnxDatapointType_DPT_Char_8859_1: // STRING

		// Simple Field (value)
		value, _valueErr := io.ReadString(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcSTRING(value), nil
	case datapointType == KnxDatapointType_DPT_Scaling: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Angle: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Percent_U8: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_DecimalFactor: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Tariff: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Value_1_Ucount: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_FanStage: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Percent_V8: // SINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Value_1_Count: // SINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Status_Mode3: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (statusA)
		statusA, _statusAErr := io.ReadBit()
		if _statusAErr != nil {
			return nil, errors.New("Error parsing 'statusA' field " + _statusAErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusA)

		// Simple Field (statusB)
		statusB, _statusBErr := io.ReadBit()
		if _statusBErr != nil {
			return nil, errors.New("Error parsing 'statusB' field " + _statusBErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusB)

		// Simple Field (statusC)
		statusC, _statusCErr := io.ReadBit()
		if _statusCErr != nil {
			return nil, errors.New("Error parsing 'statusC' field " + _statusCErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusC)

		// Simple Field (statusD)
		statusD, _statusDErr := io.ReadBit()
		if _statusDErr != nil {
			return nil, errors.New("Error parsing 'statusD' field " + _statusDErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusD)

		// Simple Field (statusE)
		statusE, _statusEErr := io.ReadBit()
		if _statusEErr != nil {
			return nil, errors.New("Error parsing 'statusE' field " + _statusEErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusE)

		// Simple Field (mode)
		mode, _modeErr := io.ReadUint8(3)
		if _modeErr != nil {
			return nil, errors.New("Error parsing 'mode' field " + _modeErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(mode)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Value_2_Ucount: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_TimePeriodMsec: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_TimePeriod10Msec: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_TimePeriod100Msec: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_TimePeriodSec: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_TimePeriodMin: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_TimePeriodHrs: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_PropDataType: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_Length_mm: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_UElCurrentmA: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_Brightness: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_Absolute_Colour_Temperature: // UINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUINT(value), nil
	case datapointType == KnxDatapointType_DPT_Value_2_Count: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_DPT_DeltaTimeMsec: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_DPT_DeltaTime10Msec: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_DPT_DeltaTime100Msec: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_DPT_DeltaTimeSec: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_DPT_DeltaTimeMin: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_DPT_DeltaTimeHrs: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_DPT_Percent_V16: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_DPT_Rotation_Angle: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_DPT_Length_m: // INT

		// Simple Field (value)
		value, _valueErr := io.ReadInt16(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcINT(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Temp: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Tempd: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Tempa: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Lux: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Wsp: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Pres: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Humidity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_AirQuality: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_AirFlow: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Time1: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Time2: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Volt: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Curr: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_PowerDensity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_KelvinPerPercent: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Power: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Volume_Flow: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Rain_Amount: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Temp_F: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Wsp_kmh: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Absolute_Humidity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Concentration_ygm3: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_TimeOfDay: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (day)
		day, _dayErr := io.ReadUint8(3)
		if _dayErr != nil {
			return nil, errors.New("Error parsing 'day' field " + _dayErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(day)

		// Simple Field (hour)
		hour, _hourErr := io.ReadUint8(5)
		if _hourErr != nil {
			return nil, errors.New("Error parsing 'hour' field " + _hourErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(hour)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(2); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (minutes)
		minutes, _minutesErr := io.ReadUint8(6)
		if _minutesErr != nil {
			return nil, errors.New("Error parsing 'minutes' field " + _minutesErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(minutes)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(2); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (seconds)
		seconds, _secondsErr := io.ReadUint8(6)
		if _secondsErr != nil {
			return nil, errors.New("Error parsing 'seconds' field " + _secondsErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(seconds)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Date: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(3); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (dayOfMonth)
		dayOfMonth, _dayOfMonthErr := io.ReadUint8(5)
		if _dayOfMonthErr != nil {
			return nil, errors.New("Error parsing 'dayOfMonth' field " + _dayOfMonthErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(dayOfMonth)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (month)
		month, _monthErr := io.ReadUint8(4)
		if _monthErr != nil {
			return nil, errors.New("Error parsing 'month' field " + _monthErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(month)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(1); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (year)
		year, _yearErr := io.ReadUint8(7)
		if _yearErr != nil {
			return nil, errors.New("Error parsing 'year' field " + _yearErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(year)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Value_4_Ucount: // UDINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUDINT(value), nil
	case datapointType == KnxDatapointType_DPT_LongTimePeriod_Sec: // UDINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUDINT(value), nil
	case datapointType == KnxDatapointType_DPT_LongTimePeriod_Min: // UDINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUDINT(value), nil
	case datapointType == KnxDatapointType_DPT_LongTimePeriod_Hrs: // UDINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUDINT(value), nil
	case datapointType == KnxDatapointType_DPT_VolumeLiquid_Litre: // UDINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUDINT(value), nil
	case datapointType == KnxDatapointType_DPT_Volume_m_3: // UDINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUDINT(value), nil
	case datapointType == KnxDatapointType_DPT_Value_4_Count: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_FlowRate_m3h: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_ActiveEnergy: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_ApparantEnergy: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_ReactiveEnergy: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_ActiveEnergy_kWh: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_ApparantEnergy_kVAh: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_ReactiveEnergy_kVARh: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_ActiveEnergy_MWh: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_LongDeltaTimeSec: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_DeltaVolumeLiquid_Litre: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_DeltaVolume_m_3: // DINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt32(32)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcDINT(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Acceleration: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Acceleration_Angular: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Activation_Energy: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Activity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Mol: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Amplitude: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_AngleRad: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_AngleDeg: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Angular_Momentum: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Angular_Velocity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Area: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Capacitance: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Charge_DensitySurface: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Charge_DensityVolume: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Compressibility: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Conductance: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electrical_Conductivity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Density: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_Charge: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_Current: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_CurrentDensity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_DipoleMoment: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_Displacement: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_FieldStrength: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_Flux: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_FluxDensity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_Polarization: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_Potential: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electric_PotentialDifference: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_ElectromagneticMoment: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Electromotive_Force: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Energy: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Force: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Frequency: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Angular_Frequency: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Heat_Capacity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Heat_FlowRate: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Heat_Quantity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Impedance: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Length: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Light_Quantity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Luminance: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Luminous_Flux: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Luminous_Intensity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Magnetic_FieldStrength: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Magnetic_Flux: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Magnetic_FluxDensity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Magnetic_Moment: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Magnetic_Polarization: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Magnetization: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_MagnetomotiveForce: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Mass: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_MassFlux: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Momentum: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Phase_AngleRad: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Phase_AngleDeg: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Power: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Power_Factor: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Pressure: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Reactance: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Resistance: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Resistivity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_SelfInductance: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_SolidAngle: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Sound_Intensity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Speed: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Stress: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Surface_Tension: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Common_Temperature: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Absolute_Temperature: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_TemperatureDifference: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Thermal_Capacity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Thermal_Conductivity: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_ThermoelectricPower: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Time: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Torque: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Volume: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Volume_Flux: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Weight: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Value_Work: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 8, 23)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Volume_Flux_Meter: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Volume_Flux_ls: // REAL

		// Simple Field (value)
		value, _valueErr := io.ReadFloat32(true, 4, 11)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcREAL(value), nil
	case datapointType == KnxDatapointType_DPT_Access_Data: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (hurz)
		hurz, _hurzErr := io.ReadUint8(4)
		if _hurzErr != nil {
			return nil, errors.New("Error parsing 'hurz' field " + _hurzErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(hurz)

		// Simple Field (value1)
		value1, _value1Err := io.ReadUint8(4)
		if _value1Err != nil {
			return nil, errors.New("Error parsing 'value1' field " + _value1Err.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(value1)

		// Simple Field (value2)
		value2, _value2Err := io.ReadUint8(4)
		if _value2Err != nil {
			return nil, errors.New("Error parsing 'value2' field " + _value2Err.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(value2)

		// Simple Field (value3)
		value3, _value3Err := io.ReadUint8(4)
		if _value3Err != nil {
			return nil, errors.New("Error parsing 'value3' field " + _value3Err.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(value3)

		// Simple Field (value4)
		value4, _value4Err := io.ReadUint8(4)
		if _value4Err != nil {
			return nil, errors.New("Error parsing 'value4' field " + _value4Err.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(value4)

		// Simple Field (value5)
		value5, _value5Err := io.ReadUint8(4)
		if _value5Err != nil {
			return nil, errors.New("Error parsing 'value5' field " + _value5Err.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(value5)

		// Simple Field (detectionError)
		detectionError, _detectionErrorErr := io.ReadBit()
		if _detectionErrorErr != nil {
			return nil, errors.New("Error parsing 'detectionError' field " + _detectionErrorErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(detectionError)

		// Simple Field (permission)
		permission, _permissionErr := io.ReadBit()
		if _permissionErr != nil {
			return nil, errors.New("Error parsing 'permission' field " + _permissionErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(permission)

		// Simple Field (readDirection)
		readDirection, _readDirectionErr := io.ReadBit()
		if _readDirectionErr != nil {
			return nil, errors.New("Error parsing 'readDirection' field " + _readDirectionErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(readDirection)

		// Simple Field (encryptionOfAccessInformation)
		encryptionOfAccessInformation, _encryptionOfAccessInformationErr := io.ReadBit()
		if _encryptionOfAccessInformationErr != nil {
			return nil, errors.New("Error parsing 'encryptionOfAccessInformation' field " + _encryptionOfAccessInformationErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(encryptionOfAccessInformation)

		// Simple Field (indexOfAccessIdentificationCode)
		indexOfAccessIdentificationCode, _indexOfAccessIdentificationCodeErr := io.ReadUint8(4)
		if _indexOfAccessIdentificationCodeErr != nil {
			return nil, errors.New("Error parsing 'indexOfAccessIdentificationCode' field " + _indexOfAccessIdentificationCodeErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(indexOfAccessIdentificationCode)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_String_ASCII: // STRING

		// Simple Field (value)
		value, _valueErr := io.ReadString(112)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcSTRING(value), nil
	case datapointType == KnxDatapointType_DPT_String_8859_1: // STRING

		// Simple Field (value)
		value, _valueErr := io.ReadString(112)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcSTRING(value), nil
	case datapointType == KnxDatapointType_DPT_SceneNumber: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(6)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_SceneControl: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (learnTheSceneCorrespondingToTheFieldSceneNumber)
		learnTheSceneCorrespondingToTheFieldSceneNumber, _learnTheSceneCorrespondingToTheFieldSceneNumberErr := io.ReadBit()
		if _learnTheSceneCorrespondingToTheFieldSceneNumberErr != nil {
			return nil, errors.New("Error parsing 'learnTheSceneCorrespondingToTheFieldSceneNumber' field " + _learnTheSceneCorrespondingToTheFieldSceneNumberErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(learnTheSceneCorrespondingToTheFieldSceneNumber)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(1); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (sceneNumber)
		sceneNumber, _sceneNumberErr := io.ReadUint8(6)
		if _sceneNumberErr != nil {
			return nil, errors.New("Error parsing 'sceneNumber' field " + _sceneNumberErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(sceneNumber)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_DateTime: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (year)
		year, _yearErr := io.ReadUint8(8)
		if _yearErr != nil {
			return nil, errors.New("Error parsing 'year' field " + _yearErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(year)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (month)
		month, _monthErr := io.ReadUint8(4)
		if _monthErr != nil {
			return nil, errors.New("Error parsing 'month' field " + _monthErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(month)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(3); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (dayofmonth)
		dayofmonth, _dayofmonthErr := io.ReadUint8(5)
		if _dayofmonthErr != nil {
			return nil, errors.New("Error parsing 'dayofmonth' field " + _dayofmonthErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(dayofmonth)

		// Simple Field (dayofweek)
		dayofweek, _dayofweekErr := io.ReadUint8(3)
		if _dayofweekErr != nil {
			return nil, errors.New("Error parsing 'dayofweek' field " + _dayofweekErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(dayofweek)

		// Simple Field (hourofday)
		hourofday, _hourofdayErr := io.ReadUint8(5)
		if _hourofdayErr != nil {
			return nil, errors.New("Error parsing 'hourofday' field " + _hourofdayErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(hourofday)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(2); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (minutes)
		minutes, _minutesErr := io.ReadUint8(6)
		if _minutesErr != nil {
			return nil, errors.New("Error parsing 'minutes' field " + _minutesErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(minutes)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(2); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (seconds)
		seconds, _secondsErr := io.ReadUint8(6)
		if _secondsErr != nil {
			return nil, errors.New("Error parsing 'seconds' field " + _secondsErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(seconds)

		// Simple Field (fault)
		fault, _faultErr := io.ReadBit()
		if _faultErr != nil {
			return nil, errors.New("Error parsing 'fault' field " + _faultErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(fault)

		// Simple Field (workingDay)
		workingDay, _workingDayErr := io.ReadBit()
		if _workingDayErr != nil {
			return nil, errors.New("Error parsing 'workingDay' field " + _workingDayErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(workingDay)

		// Simple Field (noWd)
		noWd, _noWdErr := io.ReadBit()
		if _noWdErr != nil {
			return nil, errors.New("Error parsing 'noWd' field " + _noWdErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(noWd)

		// Simple Field (noYear)
		noYear, _noYearErr := io.ReadBit()
		if _noYearErr != nil {
			return nil, errors.New("Error parsing 'noYear' field " + _noYearErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(noYear)

		// Simple Field (noDate)
		noDate, _noDateErr := io.ReadBit()
		if _noDateErr != nil {
			return nil, errors.New("Error parsing 'noDate' field " + _noDateErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(noDate)

		// Simple Field (noDayOfWeek)
		noDayOfWeek, _noDayOfWeekErr := io.ReadBit()
		if _noDayOfWeekErr != nil {
			return nil, errors.New("Error parsing 'noDayOfWeek' field " + _noDayOfWeekErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(noDayOfWeek)

		// Simple Field (noTime)
		noTime, _noTimeErr := io.ReadBit()
		if _noTimeErr != nil {
			return nil, errors.New("Error parsing 'noTime' field " + _noTimeErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(noTime)

		// Simple Field (standardSummerTime)
		standardSummerTime, _standardSummerTimeErr := io.ReadBit()
		if _standardSummerTimeErr != nil {
			return nil, errors.New("Error parsing 'standardSummerTime' field " + _standardSummerTimeErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(standardSummerTime)

		// Simple Field (qualityOfClock)
		qualityOfClock, _qualityOfClockErr := io.ReadBit()
		if _qualityOfClockErr != nil {
			return nil, errors.New("Error parsing 'qualityOfClock' field " + _qualityOfClockErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(qualityOfClock)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(7); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_SCLOMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_BuildingMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_OccMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Priority: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_LightApplicationMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_ApplicationArea: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_AlarmClassType: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_PSUMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_ErrorClass_System: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_ErrorClass_HVAC: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Time_Delay: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Beaufort_Wind_Force_Scale: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_SensorSelect: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_ActuatorConnectType: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Cloud_Cover: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_PowerReturnMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_FuelType: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_BurnerType: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_HVACMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_DHWMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_LoadPriority: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_HVACContrMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_HVACEmergMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_ChangeoverMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_ValveMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_DamperMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_HeaterMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_FanMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_MasterSlaveMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_StatusRoomSetp: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Metering_DeviceType: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_HumDehumMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_EnableHCStage: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_ADAType: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_BackupMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_StartSynchronization: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Behaviour_Lock_Unlock: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Behaviour_Bus_Power_Up_Down: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_DALI_Fade_Time: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_BlinkingMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_LightControlMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_SwitchPBModel: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_PBAction: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_DimmPBModel: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_SwitchOnMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_LoadTypeSet: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_LoadTypeDetected: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Converter_Test_Control: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_SABExcept_Behaviour: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_SABBehaviour_Lock_Unlock: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_SSSBMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_BlindsControlMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_CommMode: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_AddInfoTypes: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_RF_ModeSelect: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_RF_FilterSelect: // USINT

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(8)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_StatusGen: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(3); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (alarmStatusOfCorrespondingDatapointIsNotAcknowledged)
		alarmStatusOfCorrespondingDatapointIsNotAcknowledged, _alarmStatusOfCorrespondingDatapointIsNotAcknowledgedErr := io.ReadBit()
		if _alarmStatusOfCorrespondingDatapointIsNotAcknowledgedErr != nil {
			return nil, errors.New("Error parsing 'alarmStatusOfCorrespondingDatapointIsNotAcknowledged' field " + _alarmStatusOfCorrespondingDatapointIsNotAcknowledgedErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(alarmStatusOfCorrespondingDatapointIsNotAcknowledged)

		// Simple Field (correspondingDatapointIsInAlarm)
		correspondingDatapointIsInAlarm, _correspondingDatapointIsInAlarmErr := io.ReadBit()
		if _correspondingDatapointIsInAlarmErr != nil {
			return nil, errors.New("Error parsing 'correspondingDatapointIsInAlarm' field " + _correspondingDatapointIsInAlarmErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(correspondingDatapointIsInAlarm)

		// Simple Field (correspondingDatapointMainValueIsOverridden)
		correspondingDatapointMainValueIsOverridden, _correspondingDatapointMainValueIsOverriddenErr := io.ReadBit()
		if _correspondingDatapointMainValueIsOverriddenErr != nil {
			return nil, errors.New("Error parsing 'correspondingDatapointMainValueIsOverridden' field " + _correspondingDatapointMainValueIsOverriddenErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(correspondingDatapointMainValueIsOverridden)

		// Simple Field (correspondingDatapointMainValueIsCorruptedDueToFailure)
		correspondingDatapointMainValueIsCorruptedDueToFailure, _correspondingDatapointMainValueIsCorruptedDueToFailureErr := io.ReadBit()
		if _correspondingDatapointMainValueIsCorruptedDueToFailureErr != nil {
			return nil, errors.New("Error parsing 'correspondingDatapointMainValueIsCorruptedDueToFailure' field " + _correspondingDatapointMainValueIsCorruptedDueToFailureErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(correspondingDatapointMainValueIsCorruptedDueToFailure)

		// Simple Field (correspondingDatapointValueIsOutOfService)
		correspondingDatapointValueIsOutOfService, _correspondingDatapointValueIsOutOfServiceErr := io.ReadBit()
		if _correspondingDatapointValueIsOutOfServiceErr != nil {
			return nil, errors.New("Error parsing 'correspondingDatapointValueIsOutOfService' field " + _correspondingDatapointValueIsOutOfServiceErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(correspondingDatapointValueIsOutOfService)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Device_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(5); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (verifyModeIsOn)
		verifyModeIsOn, _verifyModeIsOnErr := io.ReadBit()
		if _verifyModeIsOnErr != nil {
			return nil, errors.New("Error parsing 'verifyModeIsOn' field " + _verifyModeIsOnErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(verifyModeIsOn)

		// Simple Field (aDatagramWithTheOwnIndividualAddressAsSourceAddressHasBeenReceived)
		aDatagramWithTheOwnIndividualAddressAsSourceAddressHasBeenReceived, _aDatagramWithTheOwnIndividualAddressAsSourceAddressHasBeenReceivedErr := io.ReadBit()
		if _aDatagramWithTheOwnIndividualAddressAsSourceAddressHasBeenReceivedErr != nil {
			return nil, errors.New("Error parsing 'aDatagramWithTheOwnIndividualAddressAsSourceAddressHasBeenReceived' field " + _aDatagramWithTheOwnIndividualAddressAsSourceAddressHasBeenReceivedErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(aDatagramWithTheOwnIndividualAddressAsSourceAddressHasBeenReceived)

		// Simple Field (theUserApplicationIsStopped)
		theUserApplicationIsStopped, _theUserApplicationIsStoppedErr := io.ReadBit()
		if _theUserApplicationIsStoppedErr != nil {
			return nil, errors.New("Error parsing 'theUserApplicationIsStopped' field " + _theUserApplicationIsStoppedErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(theUserApplicationIsStopped)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_ForceSign: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (roomhmax)
		roomhmax, _roomhmaxErr := io.ReadBit()
		if _roomhmaxErr != nil {
			return nil, errors.New("Error parsing 'roomhmax' field " + _roomhmaxErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(roomhmax)

		// Simple Field (roomhconf)
		roomhconf, _roomhconfErr := io.ReadBit()
		if _roomhconfErr != nil {
			return nil, errors.New("Error parsing 'roomhconf' field " + _roomhconfErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(roomhconf)

		// Simple Field (dhwlegio)
		dhwlegio, _dhwlegioErr := io.ReadBit()
		if _dhwlegioErr != nil {
			return nil, errors.New("Error parsing 'dhwlegio' field " + _dhwlegioErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(dhwlegio)

		// Simple Field (dhwnorm)
		dhwnorm, _dhwnormErr := io.ReadBit()
		if _dhwnormErr != nil {
			return nil, errors.New("Error parsing 'dhwnorm' field " + _dhwnormErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(dhwnorm)

		// Simple Field (overrun)
		overrun, _overrunErr := io.ReadBit()
		if _overrunErr != nil {
			return nil, errors.New("Error parsing 'overrun' field " + _overrunErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(overrun)

		// Simple Field (oversupply)
		oversupply, _oversupplyErr := io.ReadBit()
		if _oversupplyErr != nil {
			return nil, errors.New("Error parsing 'oversupply' field " + _oversupplyErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(oversupply)

		// Simple Field (protection)
		protection, _protectionErr := io.ReadBit()
		if _protectionErr != nil {
			return nil, errors.New("Error parsing 'protection' field " + _protectionErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(protection)

		// Simple Field (forcerequest)
		forcerequest, _forcerequestErr := io.ReadBit()
		if _forcerequestErr != nil {
			return nil, errors.New("Error parsing 'forcerequest' field " + _forcerequestErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(forcerequest)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_ForceSignCool: // BOOL

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
	case datapointType == KnxDatapointType_DPT_StatusRHC: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (summermode)
		summermode, _summermodeErr := io.ReadBit()
		if _summermodeErr != nil {
			return nil, errors.New("Error parsing 'summermode' field " + _summermodeErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(summermode)

		// Simple Field (statusstopoptim)
		statusstopoptim, _statusstopoptimErr := io.ReadBit()
		if _statusstopoptimErr != nil {
			return nil, errors.New("Error parsing 'statusstopoptim' field " + _statusstopoptimErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusstopoptim)

		// Simple Field (statusstartoptim)
		statusstartoptim, _statusstartoptimErr := io.ReadBit()
		if _statusstartoptimErr != nil {
			return nil, errors.New("Error parsing 'statusstartoptim' field " + _statusstartoptimErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusstartoptim)

		// Simple Field (statusmorningboost)
		statusmorningboost, _statusmorningboostErr := io.ReadBit()
		if _statusmorningboostErr != nil {
			return nil, errors.New("Error parsing 'statusmorningboost' field " + _statusmorningboostErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusmorningboost)

		// Simple Field (tempreturnlimit)
		tempreturnlimit, _tempreturnlimitErr := io.ReadBit()
		if _tempreturnlimitErr != nil {
			return nil, errors.New("Error parsing 'tempreturnlimit' field " + _tempreturnlimitErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(tempreturnlimit)

		// Simple Field (tempflowlimit)
		tempflowlimit, _tempflowlimitErr := io.ReadBit()
		if _tempflowlimitErr != nil {
			return nil, errors.New("Error parsing 'tempflowlimit' field " + _tempflowlimitErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(tempflowlimit)

		// Simple Field (satuseco)
		satuseco, _satusecoErr := io.ReadBit()
		if _satusecoErr != nil {
			return nil, errors.New("Error parsing 'satuseco' field " + _satusecoErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(satuseco)

		// Simple Field (fault)
		fault, _faultErr := io.ReadBit()
		if _faultErr != nil {
			return nil, errors.New("Error parsing 'fault' field " + _faultErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(fault)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_StatusSDHWC: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(5); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (solarloadsufficient)
		solarloadsufficient, _solarloadsufficientErr := io.ReadBit()
		if _solarloadsufficientErr != nil {
			return nil, errors.New("Error parsing 'solarloadsufficient' field " + _solarloadsufficientErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(solarloadsufficient)

		// Simple Field (sdhwloadactive)
		sdhwloadactive, _sdhwloadactiveErr := io.ReadBit()
		if _sdhwloadactiveErr != nil {
			return nil, errors.New("Error parsing 'sdhwloadactive' field " + _sdhwloadactiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(sdhwloadactive)

		// Simple Field (fault)
		fault, _faultErr := io.ReadBit()
		if _faultErr != nil {
			return nil, errors.New("Error parsing 'fault' field " + _faultErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(fault)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_FuelTypeSet: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(5); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (solidstate)
		solidstate, _solidstateErr := io.ReadBit()
		if _solidstateErr != nil {
			return nil, errors.New("Error parsing 'solidstate' field " + _solidstateErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(solidstate)

		// Simple Field (gas)
		gas, _gasErr := io.ReadBit()
		if _gasErr != nil {
			return nil, errors.New("Error parsing 'gas' field " + _gasErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(gas)

		// Simple Field (oil)
		oil, _oilErr := io.ReadBit()
		if _oilErr != nil {
			return nil, errors.New("Error parsing 'oil' field " + _oilErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(oil)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_StatusRCC: // BOOL

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
	case datapointType == KnxDatapointType_DPT_StatusAHU: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (cool)
		cool, _coolErr := io.ReadBit()
		if _coolErr != nil {
			return nil, errors.New("Error parsing 'cool' field " + _coolErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cool)

		// Simple Field (heat)
		heat, _heatErr := io.ReadBit()
		if _heatErr != nil {
			return nil, errors.New("Error parsing 'heat' field " + _heatErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(heat)

		// Simple Field (fanactive)
		fanactive, _fanactiveErr := io.ReadBit()
		if _fanactiveErr != nil {
			return nil, errors.New("Error parsing 'fanactive' field " + _fanactiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(fanactive)

		// Simple Field (fault)
		fault, _faultErr := io.ReadBit()
		if _faultErr != nil {
			return nil, errors.New("Error parsing 'fault' field " + _faultErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(fault)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_CombinedStatus_RTSM: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(3); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (statusOfHvacModeUser)
		statusOfHvacModeUser, _statusOfHvacModeUserErr := io.ReadBit()
		if _statusOfHvacModeUserErr != nil {
			return nil, errors.New("Error parsing 'statusOfHvacModeUser' field " + _statusOfHvacModeUserErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusOfHvacModeUser)

		// Simple Field (statusOfComfortProlongationUser)
		statusOfComfortProlongationUser, _statusOfComfortProlongationUserErr := io.ReadBit()
		if _statusOfComfortProlongationUserErr != nil {
			return nil, errors.New("Error parsing 'statusOfComfortProlongationUser' field " + _statusOfComfortProlongationUserErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusOfComfortProlongationUser)

		// Simple Field (effectiveValueOfTheComfortPushButton)
		effectiveValueOfTheComfortPushButton, _effectiveValueOfTheComfortPushButtonErr := io.ReadBit()
		if _effectiveValueOfTheComfortPushButtonErr != nil {
			return nil, errors.New("Error parsing 'effectiveValueOfTheComfortPushButton' field " + _effectiveValueOfTheComfortPushButtonErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(effectiveValueOfTheComfortPushButton)

		// Simple Field (effectiveValueOfThePresenceStatus)
		effectiveValueOfThePresenceStatus, _effectiveValueOfThePresenceStatusErr := io.ReadBit()
		if _effectiveValueOfThePresenceStatusErr != nil {
			return nil, errors.New("Error parsing 'effectiveValueOfThePresenceStatus' field " + _effectiveValueOfThePresenceStatusErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(effectiveValueOfThePresenceStatus)

		// Simple Field (effectiveValueOfTheWindowStatus)
		effectiveValueOfTheWindowStatus, _effectiveValueOfTheWindowStatusErr := io.ReadBit()
		if _effectiveValueOfTheWindowStatusErr != nil {
			return nil, errors.New("Error parsing 'effectiveValueOfTheWindowStatus' field " + _effectiveValueOfTheWindowStatusErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(effectiveValueOfTheWindowStatus)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_LightActuatorErrorInfo: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(1); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (overheat)
		overheat, _overheatErr := io.ReadBit()
		if _overheatErr != nil {
			return nil, errors.New("Error parsing 'overheat' field " + _overheatErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(overheat)

		// Simple Field (lampfailure)
		lampfailure, _lampfailureErr := io.ReadBit()
		if _lampfailureErr != nil {
			return nil, errors.New("Error parsing 'lampfailure' field " + _lampfailureErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(lampfailure)

		// Simple Field (defectiveload)
		defectiveload, _defectiveloadErr := io.ReadBit()
		if _defectiveloadErr != nil {
			return nil, errors.New("Error parsing 'defectiveload' field " + _defectiveloadErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(defectiveload)

		// Simple Field (underload)
		underload, _underloadErr := io.ReadBit()
		if _underloadErr != nil {
			return nil, errors.New("Error parsing 'underload' field " + _underloadErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(underload)

		// Simple Field (overcurrent)
		overcurrent, _overcurrentErr := io.ReadBit()
		if _overcurrentErr != nil {
			return nil, errors.New("Error parsing 'overcurrent' field " + _overcurrentErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(overcurrent)

		// Simple Field (undervoltage)
		undervoltage, _undervoltageErr := io.ReadBit()
		if _undervoltageErr != nil {
			return nil, errors.New("Error parsing 'undervoltage' field " + _undervoltageErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(undervoltage)

		// Simple Field (loaddetectionerror)
		loaddetectionerror, _loaddetectionerrorErr := io.ReadBit()
		if _loaddetectionerrorErr != nil {
			return nil, errors.New("Error parsing 'loaddetectionerror' field " + _loaddetectionerrorErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(loaddetectionerror)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_RF_ModeInfo: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(5); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (bibatSlave)
		bibatSlave, _bibatSlaveErr := io.ReadBit()
		if _bibatSlaveErr != nil {
			return nil, errors.New("Error parsing 'bibatSlave' field " + _bibatSlaveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(bibatSlave)

		// Simple Field (bibatMaster)
		bibatMaster, _bibatMasterErr := io.ReadBit()
		if _bibatMasterErr != nil {
			return nil, errors.New("Error parsing 'bibatMaster' field " + _bibatMasterErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(bibatMaster)

		// Simple Field (asynchronous)
		asynchronous, _asynchronousErr := io.ReadBit()
		if _asynchronousErr != nil {
			return nil, errors.New("Error parsing 'asynchronous' field " + _asynchronousErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(asynchronous)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_RF_FilterInfo: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(5); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (doa)
		doa, _doaErr := io.ReadBit()
		if _doaErr != nil {
			return nil, errors.New("Error parsing 'doa' field " + _doaErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(doa)

		// Simple Field (knxSn)
		knxSn, _knxSnErr := io.ReadBit()
		if _knxSnErr != nil {
			return nil, errors.New("Error parsing 'knxSn' field " + _knxSnErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(knxSn)

		// Simple Field (doaAndKnxSn)
		doaAndKnxSn, _doaAndKnxSnErr := io.ReadBit()
		if _doaAndKnxSnErr != nil {
			return nil, errors.New("Error parsing 'doaAndKnxSn' field " + _doaAndKnxSnErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(doaAndKnxSn)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Channel_Activation_8: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (activationStateOfChannel1)
		activationStateOfChannel1, _activationStateOfChannel1Err := io.ReadBit()
		if _activationStateOfChannel1Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel1' field " + _activationStateOfChannel1Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel1)

		// Simple Field (activationStateOfChannel2)
		activationStateOfChannel2, _activationStateOfChannel2Err := io.ReadBit()
		if _activationStateOfChannel2Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel2' field " + _activationStateOfChannel2Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel2)

		// Simple Field (activationStateOfChannel3)
		activationStateOfChannel3, _activationStateOfChannel3Err := io.ReadBit()
		if _activationStateOfChannel3Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel3' field " + _activationStateOfChannel3Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel3)

		// Simple Field (activationStateOfChannel4)
		activationStateOfChannel4, _activationStateOfChannel4Err := io.ReadBit()
		if _activationStateOfChannel4Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel4' field " + _activationStateOfChannel4Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel4)

		// Simple Field (activationStateOfChannel5)
		activationStateOfChannel5, _activationStateOfChannel5Err := io.ReadBit()
		if _activationStateOfChannel5Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel5' field " + _activationStateOfChannel5Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel5)

		// Simple Field (activationStateOfChannel6)
		activationStateOfChannel6, _activationStateOfChannel6Err := io.ReadBit()
		if _activationStateOfChannel6Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel6' field " + _activationStateOfChannel6Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel6)

		// Simple Field (activationStateOfChannel7)
		activationStateOfChannel7, _activationStateOfChannel7Err := io.ReadBit()
		if _activationStateOfChannel7Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel7' field " + _activationStateOfChannel7Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel7)

		// Simple Field (activationStateOfChannel8)
		activationStateOfChannel8, _activationStateOfChannel8Err := io.ReadBit()
		if _activationStateOfChannel8Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel8' field " + _activationStateOfChannel8Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel8)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_StatusDHWC: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(8); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (tempoptimshiftactive)
		tempoptimshiftactive, _tempoptimshiftactiveErr := io.ReadBit()
		if _tempoptimshiftactiveErr != nil {
			return nil, errors.New("Error parsing 'tempoptimshiftactive' field " + _tempoptimshiftactiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(tempoptimshiftactive)

		// Simple Field (solarenergysupport)
		solarenergysupport, _solarenergysupportErr := io.ReadBit()
		if _solarenergysupportErr != nil {
			return nil, errors.New("Error parsing 'solarenergysupport' field " + _solarenergysupportErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(solarenergysupport)

		// Simple Field (solarenergyonly)
		solarenergyonly, _solarenergyonlyErr := io.ReadBit()
		if _solarenergyonlyErr != nil {
			return nil, errors.New("Error parsing 'solarenergyonly' field " + _solarenergyonlyErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(solarenergyonly)

		// Simple Field (otherenergysourceactive)
		otherenergysourceactive, _otherenergysourceactiveErr := io.ReadBit()
		if _otherenergysourceactiveErr != nil {
			return nil, errors.New("Error parsing 'otherenergysourceactive' field " + _otherenergysourceactiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(otherenergysourceactive)

		// Simple Field (dhwpushactive)
		dhwpushactive, _dhwpushactiveErr := io.ReadBit()
		if _dhwpushactiveErr != nil {
			return nil, errors.New("Error parsing 'dhwpushactive' field " + _dhwpushactiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(dhwpushactive)

		// Simple Field (legioprotactive)
		legioprotactive, _legioprotactiveErr := io.ReadBit()
		if _legioprotactiveErr != nil {
			return nil, errors.New("Error parsing 'legioprotactive' field " + _legioprotactiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(legioprotactive)

		// Simple Field (dhwloadactive)
		dhwloadactive, _dhwloadactiveErr := io.ReadBit()
		if _dhwloadactiveErr != nil {
			return nil, errors.New("Error parsing 'dhwloadactive' field " + _dhwloadactiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(dhwloadactive)

		// Simple Field (fault)
		fault, _faultErr := io.ReadBit()
		if _faultErr != nil {
			return nil, errors.New("Error parsing 'fault' field " + _faultErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(fault)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_StatusRHCC: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(1); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (overheatalarm)
		overheatalarm, _overheatalarmErr := io.ReadBit()
		if _overheatalarmErr != nil {
			return nil, errors.New("Error parsing 'overheatalarm' field " + _overheatalarmErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(overheatalarm)

		// Simple Field (frostalarm)
		frostalarm, _frostalarmErr := io.ReadBit()
		if _frostalarmErr != nil {
			return nil, errors.New("Error parsing 'frostalarm' field " + _frostalarmErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(frostalarm)

		// Simple Field (dewpointstatus)
		dewpointstatus, _dewpointstatusErr := io.ReadBit()
		if _dewpointstatusErr != nil {
			return nil, errors.New("Error parsing 'dewpointstatus' field " + _dewpointstatusErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(dewpointstatus)

		// Simple Field (coolingdisabled)
		coolingdisabled, _coolingdisabledErr := io.ReadBit()
		if _coolingdisabledErr != nil {
			return nil, errors.New("Error parsing 'coolingdisabled' field " + _coolingdisabledErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(coolingdisabled)

		// Simple Field (statusprecool)
		statusprecool, _statusprecoolErr := io.ReadBit()
		if _statusprecoolErr != nil {
			return nil, errors.New("Error parsing 'statusprecool' field " + _statusprecoolErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusprecool)

		// Simple Field (statusecoc)
		statusecoc, _statusecocErr := io.ReadBit()
		if _statusecocErr != nil {
			return nil, errors.New("Error parsing 'statusecoc' field " + _statusecocErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusecoc)

		// Simple Field (heatcoolmode)
		heatcoolmode, _heatcoolmodeErr := io.ReadBit()
		if _heatcoolmodeErr != nil {
			return nil, errors.New("Error parsing 'heatcoolmode' field " + _heatcoolmodeErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(heatcoolmode)

		// Simple Field (heatingdiabled)
		heatingdiabled, _heatingdiabledErr := io.ReadBit()
		if _heatingdiabledErr != nil {
			return nil, errors.New("Error parsing 'heatingdiabled' field " + _heatingdiabledErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(heatingdiabled)

		// Simple Field (statusstopoptim)
		statusstopoptim, _statusstopoptimErr := io.ReadBit()
		if _statusstopoptimErr != nil {
			return nil, errors.New("Error parsing 'statusstopoptim' field " + _statusstopoptimErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusstopoptim)

		// Simple Field (statusstartoptim)
		statusstartoptim, _statusstartoptimErr := io.ReadBit()
		if _statusstartoptimErr != nil {
			return nil, errors.New("Error parsing 'statusstartoptim' field " + _statusstartoptimErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusstartoptim)

		// Simple Field (statusmorningboosth)
		statusmorningboosth, _statusmorningboosthErr := io.ReadBit()
		if _statusmorningboosthErr != nil {
			return nil, errors.New("Error parsing 'statusmorningboosth' field " + _statusmorningboosthErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusmorningboosth)

		// Simple Field (tempflowreturnlimit)
		tempflowreturnlimit, _tempflowreturnlimitErr := io.ReadBit()
		if _tempflowreturnlimitErr != nil {
			return nil, errors.New("Error parsing 'tempflowreturnlimit' field " + _tempflowreturnlimitErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(tempflowreturnlimit)

		// Simple Field (tempflowlimit)
		tempflowlimit, _tempflowlimitErr := io.ReadBit()
		if _tempflowlimitErr != nil {
			return nil, errors.New("Error parsing 'tempflowlimit' field " + _tempflowlimitErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(tempflowlimit)

		// Simple Field (statusecoh)
		statusecoh, _statusecohErr := io.ReadBit()
		if _statusecohErr != nil {
			return nil, errors.New("Error parsing 'statusecoh' field " + _statusecohErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(statusecoh)

		// Simple Field (fault)
		fault, _faultErr := io.ReadBit()
		if _faultErr != nil {
			return nil, errors.New("Error parsing 'fault' field " + _faultErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(fault)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_CombinedStatus_HVA: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(7); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (calibrationMode)
		calibrationMode, _calibrationModeErr := io.ReadBit()
		if _calibrationModeErr != nil {
			return nil, errors.New("Error parsing 'calibrationMode' field " + _calibrationModeErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(calibrationMode)

		// Simple Field (lockedPosition)
		lockedPosition, _lockedPositionErr := io.ReadBit()
		if _lockedPositionErr != nil {
			return nil, errors.New("Error parsing 'lockedPosition' field " + _lockedPositionErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(lockedPosition)

		// Simple Field (forcedPosition)
		forcedPosition, _forcedPositionErr := io.ReadBit()
		if _forcedPositionErr != nil {
			return nil, errors.New("Error parsing 'forcedPosition' field " + _forcedPositionErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(forcedPosition)

		// Simple Field (manuaOperationOverridden)
		manuaOperationOverridden, _manuaOperationOverriddenErr := io.ReadBit()
		if _manuaOperationOverriddenErr != nil {
			return nil, errors.New("Error parsing 'manuaOperationOverridden' field " + _manuaOperationOverriddenErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(manuaOperationOverridden)

		// Simple Field (serviceMode)
		serviceMode, _serviceModeErr := io.ReadBit()
		if _serviceModeErr != nil {
			return nil, errors.New("Error parsing 'serviceMode' field " + _serviceModeErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(serviceMode)

		// Simple Field (valveKick)
		valveKick, _valveKickErr := io.ReadBit()
		if _valveKickErr != nil {
			return nil, errors.New("Error parsing 'valveKick' field " + _valveKickErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(valveKick)

		// Simple Field (overload)
		overload, _overloadErr := io.ReadBit()
		if _overloadErr != nil {
			return nil, errors.New("Error parsing 'overload' field " + _overloadErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(overload)

		// Simple Field (shortCircuit)
		shortCircuit, _shortCircuitErr := io.ReadBit()
		if _shortCircuitErr != nil {
			return nil, errors.New("Error parsing 'shortCircuit' field " + _shortCircuitErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(shortCircuit)

		// Simple Field (currentValvePosition)
		currentValvePosition, _currentValvePositionErr := io.ReadBit()
		if _currentValvePositionErr != nil {
			return nil, errors.New("Error parsing 'currentValvePosition' field " + _currentValvePositionErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(currentValvePosition)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_CombinedStatus_RTC: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(7); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (coolingModeEnabled)
		coolingModeEnabled, _coolingModeEnabledErr := io.ReadBit()
		if _coolingModeEnabledErr != nil {
			return nil, errors.New("Error parsing 'coolingModeEnabled' field " + _coolingModeEnabledErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(coolingModeEnabled)

		// Simple Field (heatingModeEnabled)
		heatingModeEnabled, _heatingModeEnabledErr := io.ReadBit()
		if _heatingModeEnabledErr != nil {
			return nil, errors.New("Error parsing 'heatingModeEnabled' field " + _heatingModeEnabledErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(heatingModeEnabled)

		// Simple Field (additionalHeatingCoolingStage2Stage)
		additionalHeatingCoolingStage2Stage, _additionalHeatingCoolingStage2StageErr := io.ReadBit()
		if _additionalHeatingCoolingStage2StageErr != nil {
			return nil, errors.New("Error parsing 'additionalHeatingCoolingStage2Stage' field " + _additionalHeatingCoolingStage2StageErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(additionalHeatingCoolingStage2Stage)

		// Simple Field (controllerInactive)
		controllerInactive, _controllerInactiveErr := io.ReadBit()
		if _controllerInactiveErr != nil {
			return nil, errors.New("Error parsing 'controllerInactive' field " + _controllerInactiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(controllerInactive)

		// Simple Field (overheatAlarm)
		overheatAlarm, _overheatAlarmErr := io.ReadBit()
		if _overheatAlarmErr != nil {
			return nil, errors.New("Error parsing 'overheatAlarm' field " + _overheatAlarmErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(overheatAlarm)

		// Simple Field (frostAlarm)
		frostAlarm, _frostAlarmErr := io.ReadBit()
		if _frostAlarmErr != nil {
			return nil, errors.New("Error parsing 'frostAlarm' field " + _frostAlarmErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(frostAlarm)

		// Simple Field (dewPointStatus)
		dewPointStatus, _dewPointStatusErr := io.ReadBit()
		if _dewPointStatusErr != nil {
			return nil, errors.New("Error parsing 'dewPointStatus' field " + _dewPointStatusErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(dewPointStatus)

		// Simple Field (activeMode)
		activeMode, _activeModeErr := io.ReadBit()
		if _activeModeErr != nil {
			return nil, errors.New("Error parsing 'activeMode' field " + _activeModeErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activeMode)

		// Simple Field (generalFailureInformation)
		generalFailureInformation, _generalFailureInformationErr := io.ReadBit()
		if _generalFailureInformationErr != nil {
			return nil, errors.New("Error parsing 'generalFailureInformation' field " + _generalFailureInformationErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(generalFailureInformation)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Media: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint16(10); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (knxIp)
		knxIp, _knxIpErr := io.ReadBit()
		if _knxIpErr != nil {
			return nil, errors.New("Error parsing 'knxIp' field " + _knxIpErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(knxIp)

		// Simple Field (rf)
		rf, _rfErr := io.ReadBit()
		if _rfErr != nil {
			return nil, errors.New("Error parsing 'rf' field " + _rfErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(rf)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(1); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (pl110)
		pl110, _pl110Err := io.ReadBit()
		if _pl110Err != nil {
			return nil, errors.New("Error parsing 'pl110' field " + _pl110Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(pl110)

		// Simple Field (tp1)
		tp1, _tp1Err := io.ReadBit()
		if _tp1Err != nil {
			return nil, errors.New("Error parsing 'tp1' field " + _tp1Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(tp1)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(1); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Channel_Activation_16: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (activationStateOfChannel1)
		activationStateOfChannel1, _activationStateOfChannel1Err := io.ReadBit()
		if _activationStateOfChannel1Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel1' field " + _activationStateOfChannel1Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel1)

		// Simple Field (activationStateOfChannel2)
		activationStateOfChannel2, _activationStateOfChannel2Err := io.ReadBit()
		if _activationStateOfChannel2Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel2' field " + _activationStateOfChannel2Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel2)

		// Simple Field (activationStateOfChannel3)
		activationStateOfChannel3, _activationStateOfChannel3Err := io.ReadBit()
		if _activationStateOfChannel3Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel3' field " + _activationStateOfChannel3Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel3)

		// Simple Field (activationStateOfChannel4)
		activationStateOfChannel4, _activationStateOfChannel4Err := io.ReadBit()
		if _activationStateOfChannel4Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel4' field " + _activationStateOfChannel4Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel4)

		// Simple Field (activationStateOfChannel5)
		activationStateOfChannel5, _activationStateOfChannel5Err := io.ReadBit()
		if _activationStateOfChannel5Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel5' field " + _activationStateOfChannel5Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel5)

		// Simple Field (activationStateOfChannel6)
		activationStateOfChannel6, _activationStateOfChannel6Err := io.ReadBit()
		if _activationStateOfChannel6Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel6' field " + _activationStateOfChannel6Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel6)

		// Simple Field (activationStateOfChannel7)
		activationStateOfChannel7, _activationStateOfChannel7Err := io.ReadBit()
		if _activationStateOfChannel7Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel7' field " + _activationStateOfChannel7Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel7)

		// Simple Field (activationStateOfChannel8)
		activationStateOfChannel8, _activationStateOfChannel8Err := io.ReadBit()
		if _activationStateOfChannel8Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel8' field " + _activationStateOfChannel8Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel8)

		// Simple Field (activationStateOfChannel9)
		activationStateOfChannel9, _activationStateOfChannel9Err := io.ReadBit()
		if _activationStateOfChannel9Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel9' field " + _activationStateOfChannel9Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel9)

		// Simple Field (activationStateOfChannel10)
		activationStateOfChannel10, _activationStateOfChannel10Err := io.ReadBit()
		if _activationStateOfChannel10Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel10' field " + _activationStateOfChannel10Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel10)

		// Simple Field (activationStateOfChannel11)
		activationStateOfChannel11, _activationStateOfChannel11Err := io.ReadBit()
		if _activationStateOfChannel11Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel11' field " + _activationStateOfChannel11Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel11)

		// Simple Field (activationStateOfChannel12)
		activationStateOfChannel12, _activationStateOfChannel12Err := io.ReadBit()
		if _activationStateOfChannel12Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel12' field " + _activationStateOfChannel12Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel12)

		// Simple Field (activationStateOfChannel13)
		activationStateOfChannel13, _activationStateOfChannel13Err := io.ReadBit()
		if _activationStateOfChannel13Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel13' field " + _activationStateOfChannel13Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel13)

		// Simple Field (activationStateOfChannel14)
		activationStateOfChannel14, _activationStateOfChannel14Err := io.ReadBit()
		if _activationStateOfChannel14Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel14' field " + _activationStateOfChannel14Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel14)

		// Simple Field (activationStateOfChannel15)
		activationStateOfChannel15, _activationStateOfChannel15Err := io.ReadBit()
		if _activationStateOfChannel15Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel15' field " + _activationStateOfChannel15Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel15)

		// Simple Field (activationStateOfChannel16)
		activationStateOfChannel16, _activationStateOfChannel16Err := io.ReadBit()
		if _activationStateOfChannel16Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel16' field " + _activationStateOfChannel16Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel16)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_OnOffAction: // USINT

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(2)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_Alarm_Reaction: // USINT

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(2)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_UpDown_Action: // USINT

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(2)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_HVAC_PB_Action: // USINT

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (value)
		value, _valueErr := io.ReadUint8(2)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcUSINT(value), nil
	case datapointType == KnxDatapointType_DPT_DoubleNibble: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (busy)
		busy, _busyErr := io.ReadUint8(4)
		if _busyErr != nil {
			return nil, errors.New("Error parsing 'busy' field " + _busyErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(busy)

		// Simple Field (nak)
		nak, _nakErr := io.ReadUint8(4)
		if _nakErr != nil {
			return nil, errors.New("Error parsing 'nak' field " + _nakErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(nak)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_SceneInfo: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(1); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (sceneIsInactive)
		sceneIsInactive, _sceneIsInactiveErr := io.ReadBit()
		if _sceneIsInactiveErr != nil {
			return nil, errors.New("Error parsing 'sceneIsInactive' field " + _sceneIsInactiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(sceneIsInactive)

		// Simple Field (scenenumber)
		scenenumber, _scenenumberErr := io.ReadUint8(6)
		if _scenenumberErr != nil {
			return nil, errors.New("Error parsing 'scenenumber' field " + _scenenumberErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(scenenumber)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_CombinedInfoOnOff: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (maskBitInfoOnOffOutput16)
		maskBitInfoOnOffOutput16, _maskBitInfoOnOffOutput16Err := io.ReadBit()
		if _maskBitInfoOnOffOutput16Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput16' field " + _maskBitInfoOnOffOutput16Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput16)

		// Simple Field (maskBitInfoOnOffOutput15)
		maskBitInfoOnOffOutput15, _maskBitInfoOnOffOutput15Err := io.ReadBit()
		if _maskBitInfoOnOffOutput15Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput15' field " + _maskBitInfoOnOffOutput15Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput15)

		// Simple Field (maskBitInfoOnOffOutput14)
		maskBitInfoOnOffOutput14, _maskBitInfoOnOffOutput14Err := io.ReadBit()
		if _maskBitInfoOnOffOutput14Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput14' field " + _maskBitInfoOnOffOutput14Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput14)

		// Simple Field (maskBitInfoOnOffOutput13)
		maskBitInfoOnOffOutput13, _maskBitInfoOnOffOutput13Err := io.ReadBit()
		if _maskBitInfoOnOffOutput13Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput13' field " + _maskBitInfoOnOffOutput13Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput13)

		// Simple Field (maskBitInfoOnOffOutput12)
		maskBitInfoOnOffOutput12, _maskBitInfoOnOffOutput12Err := io.ReadBit()
		if _maskBitInfoOnOffOutput12Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput12' field " + _maskBitInfoOnOffOutput12Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput12)

		// Simple Field (maskBitInfoOnOffOutput11)
		maskBitInfoOnOffOutput11, _maskBitInfoOnOffOutput11Err := io.ReadBit()
		if _maskBitInfoOnOffOutput11Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput11' field " + _maskBitInfoOnOffOutput11Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput11)

		// Simple Field (maskBitInfoOnOffOutput10)
		maskBitInfoOnOffOutput10, _maskBitInfoOnOffOutput10Err := io.ReadBit()
		if _maskBitInfoOnOffOutput10Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput10' field " + _maskBitInfoOnOffOutput10Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput10)

		// Simple Field (maskBitInfoOnOffOutput9)
		maskBitInfoOnOffOutput9, _maskBitInfoOnOffOutput9Err := io.ReadBit()
		if _maskBitInfoOnOffOutput9Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput9' field " + _maskBitInfoOnOffOutput9Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput9)

		// Simple Field (maskBitInfoOnOffOutput8)
		maskBitInfoOnOffOutput8, _maskBitInfoOnOffOutput8Err := io.ReadBit()
		if _maskBitInfoOnOffOutput8Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput8' field " + _maskBitInfoOnOffOutput8Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput8)

		// Simple Field (maskBitInfoOnOffOutput7)
		maskBitInfoOnOffOutput7, _maskBitInfoOnOffOutput7Err := io.ReadBit()
		if _maskBitInfoOnOffOutput7Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput7' field " + _maskBitInfoOnOffOutput7Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput7)

		// Simple Field (maskBitInfoOnOffOutput6)
		maskBitInfoOnOffOutput6, _maskBitInfoOnOffOutput6Err := io.ReadBit()
		if _maskBitInfoOnOffOutput6Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput6' field " + _maskBitInfoOnOffOutput6Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput6)

		// Simple Field (maskBitInfoOnOffOutput5)
		maskBitInfoOnOffOutput5, _maskBitInfoOnOffOutput5Err := io.ReadBit()
		if _maskBitInfoOnOffOutput5Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput5' field " + _maskBitInfoOnOffOutput5Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput5)

		// Simple Field (maskBitInfoOnOffOutput4)
		maskBitInfoOnOffOutput4, _maskBitInfoOnOffOutput4Err := io.ReadBit()
		if _maskBitInfoOnOffOutput4Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput4' field " + _maskBitInfoOnOffOutput4Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput4)

		// Simple Field (maskBitInfoOnOffOutput3)
		maskBitInfoOnOffOutput3, _maskBitInfoOnOffOutput3Err := io.ReadBit()
		if _maskBitInfoOnOffOutput3Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput3' field " + _maskBitInfoOnOffOutput3Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput3)

		// Simple Field (maskBitInfoOnOffOutput2)
		maskBitInfoOnOffOutput2, _maskBitInfoOnOffOutput2Err := io.ReadBit()
		if _maskBitInfoOnOffOutput2Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput2' field " + _maskBitInfoOnOffOutput2Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput2)

		// Simple Field (maskBitInfoOnOffOutput1)
		maskBitInfoOnOffOutput1, _maskBitInfoOnOffOutput1Err := io.ReadBit()
		if _maskBitInfoOnOffOutput1Err != nil {
			return nil, errors.New("Error parsing 'maskBitInfoOnOffOutput1' field " + _maskBitInfoOnOffOutput1Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskBitInfoOnOffOutput1)

		// Simple Field (infoOnOffOutput16)
		infoOnOffOutput16, _infoOnOffOutput16Err := io.ReadBit()
		if _infoOnOffOutput16Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput16' field " + _infoOnOffOutput16Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput16)

		// Simple Field (infoOnOffOutput15)
		infoOnOffOutput15, _infoOnOffOutput15Err := io.ReadBit()
		if _infoOnOffOutput15Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput15' field " + _infoOnOffOutput15Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput15)

		// Simple Field (infoOnOffOutput14)
		infoOnOffOutput14, _infoOnOffOutput14Err := io.ReadBit()
		if _infoOnOffOutput14Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput14' field " + _infoOnOffOutput14Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput14)

		// Simple Field (infoOnOffOutput13)
		infoOnOffOutput13, _infoOnOffOutput13Err := io.ReadBit()
		if _infoOnOffOutput13Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput13' field " + _infoOnOffOutput13Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput13)

		// Simple Field (infoOnOffOutput12)
		infoOnOffOutput12, _infoOnOffOutput12Err := io.ReadBit()
		if _infoOnOffOutput12Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput12' field " + _infoOnOffOutput12Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput12)

		// Simple Field (infoOnOffOutput11)
		infoOnOffOutput11, _infoOnOffOutput11Err := io.ReadBit()
		if _infoOnOffOutput11Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput11' field " + _infoOnOffOutput11Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput11)

		// Simple Field (infoOnOffOutput10)
		infoOnOffOutput10, _infoOnOffOutput10Err := io.ReadBit()
		if _infoOnOffOutput10Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput10' field " + _infoOnOffOutput10Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput10)

		// Simple Field (infoOnOffOutput9)
		infoOnOffOutput9, _infoOnOffOutput9Err := io.ReadBit()
		if _infoOnOffOutput9Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput9' field " + _infoOnOffOutput9Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput9)

		// Simple Field (infoOnOffOutput8)
		infoOnOffOutput8, _infoOnOffOutput8Err := io.ReadBit()
		if _infoOnOffOutput8Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput8' field " + _infoOnOffOutput8Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput8)

		// Simple Field (infoOnOffOutput7)
		infoOnOffOutput7, _infoOnOffOutput7Err := io.ReadBit()
		if _infoOnOffOutput7Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput7' field " + _infoOnOffOutput7Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput7)

		// Simple Field (infoOnOffOutput6)
		infoOnOffOutput6, _infoOnOffOutput6Err := io.ReadBit()
		if _infoOnOffOutput6Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput6' field " + _infoOnOffOutput6Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput6)

		// Simple Field (infoOnOffOutput5)
		infoOnOffOutput5, _infoOnOffOutput5Err := io.ReadBit()
		if _infoOnOffOutput5Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput5' field " + _infoOnOffOutput5Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput5)

		// Simple Field (infoOnOffOutput4)
		infoOnOffOutput4, _infoOnOffOutput4Err := io.ReadBit()
		if _infoOnOffOutput4Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput4' field " + _infoOnOffOutput4Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput4)

		// Simple Field (infoOnOffOutput3)
		infoOnOffOutput3, _infoOnOffOutput3Err := io.ReadBit()
		if _infoOnOffOutput3Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput3' field " + _infoOnOffOutput3Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput3)

		// Simple Field (infoOnOffOutput2)
		infoOnOffOutput2, _infoOnOffOutput2Err := io.ReadBit()
		if _infoOnOffOutput2Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput2' field " + _infoOnOffOutput2Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput2)

		// Simple Field (infoOnOffOutput1)
		infoOnOffOutput1, _infoOnOffOutput1Err := io.ReadBit()
		if _infoOnOffOutput1Err != nil {
			return nil, errors.New("Error parsing 'infoOnOffOutput1' field " + _infoOnOffOutput1Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(infoOnOffOutput1)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_ActiveEnergy_V64: // LINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt64(64)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcLINT(value), nil
	case datapointType == KnxDatapointType_DPT_ApparantEnergy_V64: // LINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt64(64)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcLINT(value), nil
	case datapointType == KnxDatapointType_DPT_ReactiveEnergy_V64: // LINT

		// Simple Field (value)
		value, _valueErr := io.ReadInt64(64)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcLINT(value), nil
	case datapointType == KnxDatapointType_DPT_Channel_Activation_24: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (activationStateOfChannel1)
		activationStateOfChannel1, _activationStateOfChannel1Err := io.ReadBit()
		if _activationStateOfChannel1Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel1' field " + _activationStateOfChannel1Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel1)

		// Simple Field (activationStateOfChannel2)
		activationStateOfChannel2, _activationStateOfChannel2Err := io.ReadBit()
		if _activationStateOfChannel2Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel2' field " + _activationStateOfChannel2Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel2)

		// Simple Field (activationStateOfChannel3)
		activationStateOfChannel3, _activationStateOfChannel3Err := io.ReadBit()
		if _activationStateOfChannel3Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel3' field " + _activationStateOfChannel3Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel3)

		// Simple Field (activationStateOfChannel4)
		activationStateOfChannel4, _activationStateOfChannel4Err := io.ReadBit()
		if _activationStateOfChannel4Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel4' field " + _activationStateOfChannel4Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel4)

		// Simple Field (activationStateOfChannel5)
		activationStateOfChannel5, _activationStateOfChannel5Err := io.ReadBit()
		if _activationStateOfChannel5Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel5' field " + _activationStateOfChannel5Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel5)

		// Simple Field (activationStateOfChannel6)
		activationStateOfChannel6, _activationStateOfChannel6Err := io.ReadBit()
		if _activationStateOfChannel6Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel6' field " + _activationStateOfChannel6Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel6)

		// Simple Field (activationStateOfChannel7)
		activationStateOfChannel7, _activationStateOfChannel7Err := io.ReadBit()
		if _activationStateOfChannel7Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel7' field " + _activationStateOfChannel7Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel7)

		// Simple Field (activationStateOfChannel8)
		activationStateOfChannel8, _activationStateOfChannel8Err := io.ReadBit()
		if _activationStateOfChannel8Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel8' field " + _activationStateOfChannel8Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel8)

		// Simple Field (activationStateOfChannel9)
		activationStateOfChannel9, _activationStateOfChannel9Err := io.ReadBit()
		if _activationStateOfChannel9Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel9' field " + _activationStateOfChannel9Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel9)

		// Simple Field (activationStateOfChannel10)
		activationStateOfChannel10, _activationStateOfChannel10Err := io.ReadBit()
		if _activationStateOfChannel10Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel10' field " + _activationStateOfChannel10Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel10)

		// Simple Field (activationStateOfChannel11)
		activationStateOfChannel11, _activationStateOfChannel11Err := io.ReadBit()
		if _activationStateOfChannel11Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel11' field " + _activationStateOfChannel11Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel11)

		// Simple Field (activationStateOfChannel12)
		activationStateOfChannel12, _activationStateOfChannel12Err := io.ReadBit()
		if _activationStateOfChannel12Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel12' field " + _activationStateOfChannel12Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel12)

		// Simple Field (activationStateOfChannel13)
		activationStateOfChannel13, _activationStateOfChannel13Err := io.ReadBit()
		if _activationStateOfChannel13Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel13' field " + _activationStateOfChannel13Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel13)

		// Simple Field (activationStateOfChannel14)
		activationStateOfChannel14, _activationStateOfChannel14Err := io.ReadBit()
		if _activationStateOfChannel14Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel14' field " + _activationStateOfChannel14Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel14)

		// Simple Field (activationStateOfChannel15)
		activationStateOfChannel15, _activationStateOfChannel15Err := io.ReadBit()
		if _activationStateOfChannel15Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel15' field " + _activationStateOfChannel15Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel15)

		// Simple Field (activationStateOfChannel16)
		activationStateOfChannel16, _activationStateOfChannel16Err := io.ReadBit()
		if _activationStateOfChannel16Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel16' field " + _activationStateOfChannel16Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel16)

		// Simple Field (activationStateOfChannel17)
		activationStateOfChannel17, _activationStateOfChannel17Err := io.ReadBit()
		if _activationStateOfChannel17Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel17' field " + _activationStateOfChannel17Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel17)

		// Simple Field (activationStateOfChannel18)
		activationStateOfChannel18, _activationStateOfChannel18Err := io.ReadBit()
		if _activationStateOfChannel18Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel18' field " + _activationStateOfChannel18Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel18)

		// Simple Field (activationStateOfChannel19)
		activationStateOfChannel19, _activationStateOfChannel19Err := io.ReadBit()
		if _activationStateOfChannel19Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel19' field " + _activationStateOfChannel19Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel19)

		// Simple Field (activationStateOfChannel20)
		activationStateOfChannel20, _activationStateOfChannel20Err := io.ReadBit()
		if _activationStateOfChannel20Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel20' field " + _activationStateOfChannel20Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel20)

		// Simple Field (activationStateOfChannel21)
		activationStateOfChannel21, _activationStateOfChannel21Err := io.ReadBit()
		if _activationStateOfChannel21Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel21' field " + _activationStateOfChannel21Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel21)

		// Simple Field (activationStateOfChannel22)
		activationStateOfChannel22, _activationStateOfChannel22Err := io.ReadBit()
		if _activationStateOfChannel22Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel22' field " + _activationStateOfChannel22Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel22)

		// Simple Field (activationStateOfChannel23)
		activationStateOfChannel23, _activationStateOfChannel23Err := io.ReadBit()
		if _activationStateOfChannel23Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel23' field " + _activationStateOfChannel23Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel23)

		// Simple Field (activationStateOfChannel24)
		activationStateOfChannel24, _activationStateOfChannel24Err := io.ReadBit()
		if _activationStateOfChannel24Err != nil {
			return nil, errors.New("Error parsing 'activationStateOfChannel24' field " + _activationStateOfChannel24Err.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(activationStateOfChannel24)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_HVACModeNext: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (delayTimeMin)
		delayTimeMin, _delayTimeMinErr := io.ReadUint16(16)
		if _delayTimeMinErr != nil {
			return nil, errors.New("Error parsing 'delayTimeMin' field " + _delayTimeMinErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(delayTimeMin)

		// Simple Field (hvacMode)
		hvacMode, _hvacModeErr := io.ReadUint8(8)
		if _hvacModeErr != nil {
			return nil, errors.New("Error parsing 'hvacMode' field " + _hvacModeErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(hvacMode)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_DHWModeNext: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (delayTimeMin)
		delayTimeMin, _delayTimeMinErr := io.ReadUint16(16)
		if _delayTimeMinErr != nil {
			return nil, errors.New("Error parsing 'delayTimeMin' field " + _delayTimeMinErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(delayTimeMin)

		// Simple Field (dhwMode)
		dhwMode, _dhwModeErr := io.ReadUint8(8)
		if _dhwModeErr != nil {
			return nil, errors.New("Error parsing 'dhwMode' field " + _dhwModeErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(dhwMode)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_OccModeNext: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (delayTimeMin)
		delayTimeMin, _delayTimeMinErr := io.ReadUint16(16)
		if _delayTimeMinErr != nil {
			return nil, errors.New("Error parsing 'delayTimeMin' field " + _delayTimeMinErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(delayTimeMin)

		// Simple Field (occupancyMode)
		occupancyMode, _occupancyModeErr := io.ReadUint8(8)
		if _occupancyModeErr != nil {
			return nil, errors.New("Error parsing 'occupancyMode' field " + _occupancyModeErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(occupancyMode)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_BuildingModeNext: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (delayTimeMin)
		delayTimeMin, _delayTimeMinErr := io.ReadUint16(16)
		if _delayTimeMinErr != nil {
			return nil, errors.New("Error parsing 'delayTimeMin' field " + _delayTimeMinErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(delayTimeMin)

		// Simple Field (buildingMode)
		buildingMode, _buildingModeErr := io.ReadUint8(8)
		if _buildingModeErr != nil {
			return nil, errors.New("Error parsing 'buildingMode' field " + _buildingModeErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(buildingMode)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Version: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (magicNumber)
		magicNumber, _magicNumberErr := io.ReadUint8(5)
		if _magicNumberErr != nil {
			return nil, errors.New("Error parsing 'magicNumber' field " + _magicNumberErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(magicNumber)

		// Simple Field (versionNumber)
		versionNumber, _versionNumberErr := io.ReadUint8(5)
		if _versionNumberErr != nil {
			return nil, errors.New("Error parsing 'versionNumber' field " + _versionNumberErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(versionNumber)

		// Simple Field (revisionNumber)
		revisionNumber, _revisionNumberErr := io.ReadUint8(6)
		if _revisionNumberErr != nil {
			return nil, errors.New("Error parsing 'revisionNumber' field " + _revisionNumberErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(revisionNumber)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_AlarmInfo: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (logNumber)
		logNumber, _logNumberErr := io.ReadUint8(8)
		if _logNumberErr != nil {
			return nil, errors.New("Error parsing 'logNumber' field " + _logNumberErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(logNumber)

		// Simple Field (alarmPriority)
		alarmPriority, _alarmPriorityErr := io.ReadUint8(8)
		if _alarmPriorityErr != nil {
			return nil, errors.New("Error parsing 'alarmPriority' field " + _alarmPriorityErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(alarmPriority)

		// Simple Field (applicationArea)
		applicationArea, _applicationAreaErr := io.ReadUint8(8)
		if _applicationAreaErr != nil {
			return nil, errors.New("Error parsing 'applicationArea' field " + _applicationAreaErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(applicationArea)

		// Simple Field (errorClass)
		errorClass, _errorClassErr := io.ReadUint8(8)
		if _errorClassErr != nil {
			return nil, errors.New("Error parsing 'errorClass' field " + _errorClassErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(errorClass)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (errorcodeSup)
		errorcodeSup, _errorcodeSupErr := io.ReadBit()
		if _errorcodeSupErr != nil {
			return nil, errors.New("Error parsing 'errorcodeSup' field " + _errorcodeSupErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(errorcodeSup)

		// Simple Field (alarmtextSup)
		alarmtextSup, _alarmtextSupErr := io.ReadBit()
		if _alarmtextSupErr != nil {
			return nil, errors.New("Error parsing 'alarmtextSup' field " + _alarmtextSupErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(alarmtextSup)

		// Simple Field (timestampSup)
		timestampSup, _timestampSupErr := io.ReadBit()
		if _timestampSupErr != nil {
			return nil, errors.New("Error parsing 'timestampSup' field " + _timestampSupErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(timestampSup)

		// Simple Field (ackSup)
		ackSup, _ackSupErr := io.ReadBit()
		if _ackSupErr != nil {
			return nil, errors.New("Error parsing 'ackSup' field " + _ackSupErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(ackSup)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(5); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (locked)
		locked, _lockedErr := io.ReadBit()
		if _lockedErr != nil {
			return nil, errors.New("Error parsing 'locked' field " + _lockedErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(locked)

		// Simple Field (alarmunack)
		alarmunack, _alarmunackErr := io.ReadBit()
		if _alarmunackErr != nil {
			return nil, errors.New("Error parsing 'alarmunack' field " + _alarmunackErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(alarmunack)

		// Simple Field (inalarm)
		inalarm, _inalarmErr := io.ReadBit()
		if _inalarmErr != nil {
			return nil, errors.New("Error parsing 'inalarm' field " + _inalarmErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(inalarm)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_TempRoomSetpSetF16_3: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (tempsetpcomf)
		tempsetpcomf, _tempsetpcomfErr := io.ReadFloat32(true, 4, 11)
		if _tempsetpcomfErr != nil {
			return nil, errors.New("Error parsing 'tempsetpcomf' field " + _tempsetpcomfErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(tempsetpcomf)

		// Simple Field (tempsetpstdby)
		tempsetpstdby, _tempsetpstdbyErr := io.ReadFloat32(true, 4, 11)
		if _tempsetpstdbyErr != nil {
			return nil, errors.New("Error parsing 'tempsetpstdby' field " + _tempsetpstdbyErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(tempsetpstdby)

		// Simple Field (tempsetpeco)
		tempsetpeco, _tempsetpecoErr := io.ReadFloat32(true, 4, 11)
		if _tempsetpecoErr != nil {
			return nil, errors.New("Error parsing 'tempsetpeco' field " + _tempsetpecoErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(tempsetpeco)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_TempRoomSetpSetShiftF16_3: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (tempsetpshiftcomf)
		tempsetpshiftcomf, _tempsetpshiftcomfErr := io.ReadFloat32(true, 4, 11)
		if _tempsetpshiftcomfErr != nil {
			return nil, errors.New("Error parsing 'tempsetpshiftcomf' field " + _tempsetpshiftcomfErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(tempsetpshiftcomf)

		// Simple Field (tempsetpshiftstdby)
		tempsetpshiftstdby, _tempsetpshiftstdbyErr := io.ReadFloat32(true, 4, 11)
		if _tempsetpshiftstdbyErr != nil {
			return nil, errors.New("Error parsing 'tempsetpshiftstdby' field " + _tempsetpshiftstdbyErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(tempsetpshiftstdby)

		// Simple Field (tempsetpshifteco)
		tempsetpshifteco, _tempsetpshiftecoErr := io.ReadFloat32(true, 4, 11)
		if _tempsetpshiftecoErr != nil {
			return nil, errors.New("Error parsing 'tempsetpshifteco' field " + _tempsetpshiftecoErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(tempsetpshifteco)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Scaling_Speed: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (timePeriod)
		timePeriod, _timePeriodErr := io.ReadUint16(16)
		if _timePeriodErr != nil {
			return nil, errors.New("Error parsing 'timePeriod' field " + _timePeriodErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(timePeriod)

		// Simple Field (percent)
		percent, _percentErr := io.ReadUint8(8)
		if _percentErr != nil {
			return nil, errors.New("Error parsing 'percent' field " + _percentErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(percent)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Scaling_Step_Time: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (timePeriod)
		timePeriod, _timePeriodErr := io.ReadUint16(16)
		if _timePeriodErr != nil {
			return nil, errors.New("Error parsing 'timePeriod' field " + _timePeriodErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(timePeriod)

		// Simple Field (percent)
		percent, _percentErr := io.ReadUint8(8)
		if _percentErr != nil {
			return nil, errors.New("Error parsing 'percent' field " + _percentErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(percent)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_MeteringValue: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (countval)
		countval, _countvalErr := io.ReadInt32(32)
		if _countvalErr != nil {
			return nil, errors.New("Error parsing 'countval' field " + _countvalErr.Error())
		}
		_map["Struct"] = values.NewPlcDINT(countval)

		// Simple Field (valinffield)
		valinffield, _valinffieldErr := io.ReadUint8(8)
		if _valinffieldErr != nil {
			return nil, errors.New("Error parsing 'valinffield' field " + _valinffieldErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(valinffield)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(3); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (alarmunack)
		alarmunack, _alarmunackErr := io.ReadBit()
		if _alarmunackErr != nil {
			return nil, errors.New("Error parsing 'alarmunack' field " + _alarmunackErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(alarmunack)

		// Simple Field (inalarm)
		inalarm, _inalarmErr := io.ReadBit()
		if _inalarmErr != nil {
			return nil, errors.New("Error parsing 'inalarm' field " + _inalarmErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(inalarm)

		// Simple Field (overridden)
		overridden, _overriddenErr := io.ReadBit()
		if _overriddenErr != nil {
			return nil, errors.New("Error parsing 'overridden' field " + _overriddenErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(overridden)

		// Simple Field (fault)
		fault, _faultErr := io.ReadBit()
		if _faultErr != nil {
			return nil, errors.New("Error parsing 'fault' field " + _faultErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(fault)

		// Simple Field (outofservice)
		outofservice, _outofserviceErr := io.ReadBit()
		if _outofserviceErr != nil {
			return nil, errors.New("Error parsing 'outofservice' field " + _outofserviceErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(outofservice)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_MBus_Address: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (manufactid)
		manufactid, _manufactidErr := io.ReadUint16(16)
		if _manufactidErr != nil {
			return nil, errors.New("Error parsing 'manufactid' field " + _manufactidErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(manufactid)

		// Simple Field (identnumber)
		identnumber, _identnumberErr := io.ReadUint32(32)
		if _identnumberErr != nil {
			return nil, errors.New("Error parsing 'identnumber' field " + _identnumberErr.Error())
		}
		_map["Struct"] = values.NewPlcUDINT(identnumber)

		// Simple Field (version)
		version, _versionErr := io.ReadUint8(8)
		if _versionErr != nil {
			return nil, errors.New("Error parsing 'version' field " + _versionErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(version)

		// Simple Field (medium)
		medium, _mediumErr := io.ReadUint8(8)
		if _mediumErr != nil {
			return nil, errors.New("Error parsing 'medium' field " + _mediumErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(medium)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Colour_RGB: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (r)
		r, _rErr := io.ReadUint8(8)
		if _rErr != nil {
			return nil, errors.New("Error parsing 'r' field " + _rErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(r)

		// Simple Field (g)
		g, _gErr := io.ReadUint8(8)
		if _gErr != nil {
			return nil, errors.New("Error parsing 'g' field " + _gErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(g)

		// Simple Field (b)
		b, _bErr := io.ReadUint8(8)
		if _bErr != nil {
			return nil, errors.New("Error parsing 'b' field " + _bErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(b)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_LanguageCodeAlpha2_ASCII: // STRING

		// Simple Field (value)
		value, _valueErr := io.ReadString(16)
		if _valueErr != nil {
			return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
		}
		return values.NewPlcSTRING(value), nil
	case datapointType == KnxDatapointType_DPT_Tariff_ActiveEnergy: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (activeelectricalenergy)
		activeelectricalenergy, _activeelectricalenergyErr := io.ReadInt32(32)
		if _activeelectricalenergyErr != nil {
			return nil, errors.New("Error parsing 'activeelectricalenergy' field " + _activeelectricalenergyErr.Error())
		}
		_map["Struct"] = values.NewPlcDINT(activeelectricalenergy)

		// Simple Field (tariff)
		tariff, _tariffErr := io.ReadUint8(8)
		if _tariffErr != nil {
			return nil, errors.New("Error parsing 'tariff' field " + _tariffErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(tariff)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (electricalengergyvalidity)
		electricalengergyvalidity, _electricalengergyvalidityErr := io.ReadBit()
		if _electricalengergyvalidityErr != nil {
			return nil, errors.New("Error parsing 'electricalengergyvalidity' field " + _electricalengergyvalidityErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(electricalengergyvalidity)

		// Simple Field (tariffvalidity)
		tariffvalidity, _tariffvalidityErr := io.ReadBit()
		if _tariffvalidityErr != nil {
			return nil, errors.New("Error parsing 'tariffvalidity' field " + _tariffvalidityErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(tariffvalidity)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Prioritised_Mode_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (deactivationOfPriority)
		deactivationOfPriority, _deactivationOfPriorityErr := io.ReadBit()
		if _deactivationOfPriorityErr != nil {
			return nil, errors.New("Error parsing 'deactivationOfPriority' field " + _deactivationOfPriorityErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(deactivationOfPriority)

		// Simple Field (priorityLevel)
		priorityLevel, _priorityLevelErr := io.ReadUint8(3)
		if _priorityLevelErr != nil {
			return nil, errors.New("Error parsing 'priorityLevel' field " + _priorityLevelErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(priorityLevel)

		// Simple Field (modeLevel)
		modeLevel, _modeLevelErr := io.ReadUint8(4)
		if _modeLevelErr != nil {
			return nil, errors.New("Error parsing 'modeLevel' field " + _modeLevelErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(modeLevel)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_DALI_Control_Gear_Diagnostic: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(5); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (convertorError)
		convertorError, _convertorErrorErr := io.ReadBit()
		if _convertorErrorErr != nil {
			return nil, errors.New("Error parsing 'convertorError' field " + _convertorErrorErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(convertorError)

		// Simple Field (ballastFailure)
		ballastFailure, _ballastFailureErr := io.ReadBit()
		if _ballastFailureErr != nil {
			return nil, errors.New("Error parsing 'ballastFailure' field " + _ballastFailureErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(ballastFailure)

		// Simple Field (lampFailure)
		lampFailure, _lampFailureErr := io.ReadBit()
		if _lampFailureErr != nil {
			return nil, errors.New("Error parsing 'lampFailure' field " + _lampFailureErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(lampFailure)

		// Simple Field (readOrResponse)
		readOrResponse, _readOrResponseErr := io.ReadBit()
		if _readOrResponseErr != nil {
			return nil, errors.New("Error parsing 'readOrResponse' field " + _readOrResponseErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(readOrResponse)

		// Simple Field (addressIndicator)
		addressIndicator, _addressIndicatorErr := io.ReadBit()
		if _addressIndicatorErr != nil {
			return nil, errors.New("Error parsing 'addressIndicator' field " + _addressIndicatorErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(addressIndicator)

		// Simple Field (daliDeviceAddressOrDaliGroupAddress)
		daliDeviceAddressOrDaliGroupAddress, _daliDeviceAddressOrDaliGroupAddressErr := io.ReadUint8(6)
		if _daliDeviceAddressOrDaliGroupAddressErr != nil {
			return nil, errors.New("Error parsing 'daliDeviceAddressOrDaliGroupAddress' field " + _daliDeviceAddressOrDaliGroupAddressErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(daliDeviceAddressOrDaliGroupAddress)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_DALI_Diagnostics: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (ballastFailure)
		ballastFailure, _ballastFailureErr := io.ReadBit()
		if _ballastFailureErr != nil {
			return nil, errors.New("Error parsing 'ballastFailure' field " + _ballastFailureErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(ballastFailure)

		// Simple Field (lampFailure)
		lampFailure, _lampFailureErr := io.ReadBit()
		if _lampFailureErr != nil {
			return nil, errors.New("Error parsing 'lampFailure' field " + _lampFailureErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(lampFailure)

		// Simple Field (deviceAddress)
		deviceAddress, _deviceAddressErr := io.ReadUint8(6)
		if _deviceAddressErr != nil {
			return nil, errors.New("Error parsing 'deviceAddress' field " + _deviceAddressErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(deviceAddress)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_CombinedPosition: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (heightPosition)
		heightPosition, _heightPositionErr := io.ReadUint8(8)
		if _heightPositionErr != nil {
			return nil, errors.New("Error parsing 'heightPosition' field " + _heightPositionErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(heightPosition)

		// Simple Field (slatsPosition)
		slatsPosition, _slatsPositionErr := io.ReadUint8(8)
		if _slatsPositionErr != nil {
			return nil, errors.New("Error parsing 'slatsPosition' field " + _slatsPositionErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(slatsPosition)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (validityHeightPosition)
		validityHeightPosition, _validityHeightPositionErr := io.ReadBit()
		if _validityHeightPositionErr != nil {
			return nil, errors.New("Error parsing 'validityHeightPosition' field " + _validityHeightPositionErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(validityHeightPosition)

		// Simple Field (validitySlatsPosition)
		validitySlatsPosition, _validitySlatsPositionErr := io.ReadBit()
		if _validitySlatsPositionErr != nil {
			return nil, errors.New("Error parsing 'validitySlatsPosition' field " + _validitySlatsPositionErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(validitySlatsPosition)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_StatusSAB: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (heightPosition)
		heightPosition, _heightPositionErr := io.ReadUint8(8)
		if _heightPositionErr != nil {
			return nil, errors.New("Error parsing 'heightPosition' field " + _heightPositionErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(heightPosition)

		// Simple Field (slatsPosition)
		slatsPosition, _slatsPositionErr := io.ReadUint8(8)
		if _slatsPositionErr != nil {
			return nil, errors.New("Error parsing 'slatsPosition' field " + _slatsPositionErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(slatsPosition)

		// Simple Field (upperEndPosReached)
		upperEndPosReached, _upperEndPosReachedErr := io.ReadBit()
		if _upperEndPosReachedErr != nil {
			return nil, errors.New("Error parsing 'upperEndPosReached' field " + _upperEndPosReachedErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(upperEndPosReached)

		// Simple Field (lowerEndPosReached)
		lowerEndPosReached, _lowerEndPosReachedErr := io.ReadBit()
		if _lowerEndPosReachedErr != nil {
			return nil, errors.New("Error parsing 'lowerEndPosReached' field " + _lowerEndPosReachedErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(lowerEndPosReached)

		// Simple Field (lowerPredefPosReachedTypHeight100PercentSlatsAngle100Percent)
		lowerPredefPosReachedTypHeight100PercentSlatsAngle100Percent, _lowerPredefPosReachedTypHeight100PercentSlatsAngle100PercentErr := io.ReadBit()
		if _lowerPredefPosReachedTypHeight100PercentSlatsAngle100PercentErr != nil {
			return nil, errors.New("Error parsing 'lowerPredefPosReachedTypHeight100PercentSlatsAngle100Percent' field " + _lowerPredefPosReachedTypHeight100PercentSlatsAngle100PercentErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(lowerPredefPosReachedTypHeight100PercentSlatsAngle100Percent)

		// Simple Field (targetPosDrive)
		targetPosDrive, _targetPosDriveErr := io.ReadBit()
		if _targetPosDriveErr != nil {
			return nil, errors.New("Error parsing 'targetPosDrive' field " + _targetPosDriveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(targetPosDrive)

		// Simple Field (restrictionOfTargetHeightPosPosCanNotBeReached)
		restrictionOfTargetHeightPosPosCanNotBeReached, _restrictionOfTargetHeightPosPosCanNotBeReachedErr := io.ReadBit()
		if _restrictionOfTargetHeightPosPosCanNotBeReachedErr != nil {
			return nil, errors.New("Error parsing 'restrictionOfTargetHeightPosPosCanNotBeReached' field " + _restrictionOfTargetHeightPosPosCanNotBeReachedErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(restrictionOfTargetHeightPosPosCanNotBeReached)

		// Simple Field (restrictionOfSlatsHeightPosPosCanNotBeReached)
		restrictionOfSlatsHeightPosPosCanNotBeReached, _restrictionOfSlatsHeightPosPosCanNotBeReachedErr := io.ReadBit()
		if _restrictionOfSlatsHeightPosPosCanNotBeReachedErr != nil {
			return nil, errors.New("Error parsing 'restrictionOfSlatsHeightPosPosCanNotBeReached' field " + _restrictionOfSlatsHeightPosPosCanNotBeReachedErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(restrictionOfSlatsHeightPosPosCanNotBeReached)

		// Simple Field (atLeastOneOfTheInputsWindRainFrostAlarmIsInAlarm)
		atLeastOneOfTheInputsWindRainFrostAlarmIsInAlarm, _atLeastOneOfTheInputsWindRainFrostAlarmIsInAlarmErr := io.ReadBit()
		if _atLeastOneOfTheInputsWindRainFrostAlarmIsInAlarmErr != nil {
			return nil, errors.New("Error parsing 'atLeastOneOfTheInputsWindRainFrostAlarmIsInAlarm' field " + _atLeastOneOfTheInputsWindRainFrostAlarmIsInAlarmErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(atLeastOneOfTheInputsWindRainFrostAlarmIsInAlarm)

		// Simple Field (upDownPositionIsForcedByMoveupdownforcedInput)
		upDownPositionIsForcedByMoveupdownforcedInput, _upDownPositionIsForcedByMoveupdownforcedInputErr := io.ReadBit()
		if _upDownPositionIsForcedByMoveupdownforcedInputErr != nil {
			return nil, errors.New("Error parsing 'upDownPositionIsForcedByMoveupdownforcedInput' field " + _upDownPositionIsForcedByMoveupdownforcedInputErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(upDownPositionIsForcedByMoveupdownforcedInput)

		// Simple Field (movementIsLockedEGByDevicelockedInput)
		movementIsLockedEGByDevicelockedInput, _movementIsLockedEGByDevicelockedInputErr := io.ReadBit()
		if _movementIsLockedEGByDevicelockedInputErr != nil {
			return nil, errors.New("Error parsing 'movementIsLockedEGByDevicelockedInput' field " + _movementIsLockedEGByDevicelockedInputErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(movementIsLockedEGByDevicelockedInput)

		// Simple Field (actuatorSetvalueIsLocallyOverriddenEGViaALocalUserInterface)
		actuatorSetvalueIsLocallyOverriddenEGViaALocalUserInterface, _actuatorSetvalueIsLocallyOverriddenEGViaALocalUserInterfaceErr := io.ReadBit()
		if _actuatorSetvalueIsLocallyOverriddenEGViaALocalUserInterfaceErr != nil {
			return nil, errors.New("Error parsing 'actuatorSetvalueIsLocallyOverriddenEGViaALocalUserInterface' field " + _actuatorSetvalueIsLocallyOverriddenEGViaALocalUserInterfaceErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(actuatorSetvalueIsLocallyOverriddenEGViaALocalUserInterface)

		// Simple Field (generalFailureOfTheActuatorOrTheDrive)
		generalFailureOfTheActuatorOrTheDrive, _generalFailureOfTheActuatorOrTheDriveErr := io.ReadBit()
		if _generalFailureOfTheActuatorOrTheDriveErr != nil {
			return nil, errors.New("Error parsing 'generalFailureOfTheActuatorOrTheDrive' field " + _generalFailureOfTheActuatorOrTheDriveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(generalFailureOfTheActuatorOrTheDrive)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(3); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (validityHeightPos)
		validityHeightPos, _validityHeightPosErr := io.ReadBit()
		if _validityHeightPosErr != nil {
			return nil, errors.New("Error parsing 'validityHeightPos' field " + _validityHeightPosErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(validityHeightPos)

		// Simple Field (validitySlatsPos)
		validitySlatsPos, _validitySlatsPosErr := io.ReadBit()
		if _validitySlatsPosErr != nil {
			return nil, errors.New("Error parsing 'validitySlatsPos' field " + _validitySlatsPosErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(validitySlatsPos)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Colour_xyY: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (xAxis)
		xAxis, _xAxisErr := io.ReadUint16(16)
		if _xAxisErr != nil {
			return nil, errors.New("Error parsing 'xAxis' field " + _xAxisErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(xAxis)

		// Simple Field (yAxis)
		yAxis, _yAxisErr := io.ReadUint16(16)
		if _yAxisErr != nil {
			return nil, errors.New("Error parsing 'yAxis' field " + _yAxisErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(yAxis)

		// Simple Field (brightness)
		brightness, _brightnessErr := io.ReadUint8(8)
		if _brightnessErr != nil {
			return nil, errors.New("Error parsing 'brightness' field " + _brightnessErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(brightness)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (validityXy)
		validityXy, _validityXyErr := io.ReadBit()
		if _validityXyErr != nil {
			return nil, errors.New("Error parsing 'validityXy' field " + _validityXyErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(validityXy)

		// Simple Field (validityBrightness)
		validityBrightness, _validityBrightnessErr := io.ReadBit()
		if _validityBrightnessErr != nil {
			return nil, errors.New("Error parsing 'validityBrightness' field " + _validityBrightnessErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(validityBrightness)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Converter_Status: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (converterModeAccordingToTheDaliConverterStateMachine)
		converterModeAccordingToTheDaliConverterStateMachine, _converterModeAccordingToTheDaliConverterStateMachineErr := io.ReadUint8(4)
		if _converterModeAccordingToTheDaliConverterStateMachineErr != nil {
			return nil, errors.New("Error parsing 'converterModeAccordingToTheDaliConverterStateMachine' field " + _converterModeAccordingToTheDaliConverterStateMachineErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(converterModeAccordingToTheDaliConverterStateMachine)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(2); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (hardwiredSwitchIsActive)
		hardwiredSwitchIsActive, _hardwiredSwitchIsActiveErr := io.ReadBit()
		if _hardwiredSwitchIsActiveErr != nil {
			return nil, errors.New("Error parsing 'hardwiredSwitchIsActive' field " + _hardwiredSwitchIsActiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(hardwiredSwitchIsActive)

		// Simple Field (hardwiredInhibitIsActive)
		hardwiredInhibitIsActive, _hardwiredInhibitIsActiveErr := io.ReadBit()
		if _hardwiredInhibitIsActiveErr != nil {
			return nil, errors.New("Error parsing 'hardwiredInhibitIsActive' field " + _hardwiredInhibitIsActiveErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(hardwiredInhibitIsActive)

		// Simple Field (functionTestPending)
		functionTestPending, _functionTestPendingErr := io.ReadUint8(2)
		if _functionTestPendingErr != nil {
			return nil, errors.New("Error parsing 'functionTestPending' field " + _functionTestPendingErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(functionTestPending)

		// Simple Field (durationTestPending)
		durationTestPending, _durationTestPendingErr := io.ReadUint8(2)
		if _durationTestPendingErr != nil {
			return nil, errors.New("Error parsing 'durationTestPending' field " + _durationTestPendingErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(durationTestPending)

		// Simple Field (partialDurationTestPending)
		partialDurationTestPending, _partialDurationTestPendingErr := io.ReadUint8(2)
		if _partialDurationTestPendingErr != nil {
			return nil, errors.New("Error parsing 'partialDurationTestPending' field " + _partialDurationTestPendingErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(partialDurationTestPending)

		// Simple Field (converterFailure)
		converterFailure, _converterFailureErr := io.ReadUint8(2)
		if _converterFailureErr != nil {
			return nil, errors.New("Error parsing 'converterFailure' field " + _converterFailureErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(converterFailure)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Converter_Test_Result: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (ltrf)
		ltrf, _ltrfErr := io.ReadUint8(4)
		if _ltrfErr != nil {
			return nil, errors.New("Error parsing 'ltrf' field " + _ltrfErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(ltrf)

		// Simple Field (ltrd)
		ltrd, _ltrdErr := io.ReadUint8(4)
		if _ltrdErr != nil {
			return nil, errors.New("Error parsing 'ltrd' field " + _ltrdErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(ltrd)

		// Simple Field (ltrp)
		ltrp, _ltrpErr := io.ReadUint8(4)
		if _ltrpErr != nil {
			return nil, errors.New("Error parsing 'ltrp' field " + _ltrpErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(ltrp)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (sf)
		sf, _sfErr := io.ReadUint8(2)
		if _sfErr != nil {
			return nil, errors.New("Error parsing 'sf' field " + _sfErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(sf)

		// Simple Field (sd)
		sd, _sdErr := io.ReadUint8(2)
		if _sdErr != nil {
			return nil, errors.New("Error parsing 'sd' field " + _sdErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(sd)

		// Simple Field (sp)
		sp, _spErr := io.ReadUint8(2)
		if _spErr != nil {
			return nil, errors.New("Error parsing 'sp' field " + _spErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(sp)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(2); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (ldtr)
		ldtr, _ldtrErr := io.ReadUint16(16)
		if _ldtrErr != nil {
			return nil, errors.New("Error parsing 'ldtr' field " + _ldtrErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(ldtr)

		// Simple Field (lpdtr)
		lpdtr, _lpdtrErr := io.ReadUint8(8)
		if _lpdtrErr != nil {
			return nil, errors.New("Error parsing 'lpdtr' field " + _lpdtrErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(lpdtr)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Battery_Info: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(5); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (batteryFailure)
		batteryFailure, _batteryFailureErr := io.ReadBit()
		if _batteryFailureErr != nil {
			return nil, errors.New("Error parsing 'batteryFailure' field " + _batteryFailureErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(batteryFailure)

		// Simple Field (batteryDurationFailure)
		batteryDurationFailure, _batteryDurationFailureErr := io.ReadBit()
		if _batteryDurationFailureErr != nil {
			return nil, errors.New("Error parsing 'batteryDurationFailure' field " + _batteryDurationFailureErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(batteryDurationFailure)

		// Simple Field (batteryFullyCharged)
		batteryFullyCharged, _batteryFullyChargedErr := io.ReadBit()
		if _batteryFullyChargedErr != nil {
			return nil, errors.New("Error parsing 'batteryFullyCharged' field " + _batteryFullyChargedErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(batteryFullyCharged)

		// Simple Field (batteryChargeLevel)
		batteryChargeLevel, _batteryChargeLevelErr := io.ReadUint8(8)
		if _batteryChargeLevelErr != nil {
			return nil, errors.New("Error parsing 'batteryChargeLevel' field " + _batteryChargeLevelErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(batteryChargeLevel)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Brightness_Colour_Temperature_Transition: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (ms)
		ms, _msErr := io.ReadUint16(16)
		if _msErr != nil {
			return nil, errors.New("Error parsing 'ms' field " + _msErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(ms)

		// Simple Field (temperatureK)
		temperatureK, _temperatureKErr := io.ReadUint16(16)
		if _temperatureKErr != nil {
			return nil, errors.New("Error parsing 'temperatureK' field " + _temperatureKErr.Error())
		}
		_map["Struct"] = values.NewPlcUINT(temperatureK)

		// Simple Field (percent)
		percent, _percentErr := io.ReadUint8(8)
		if _percentErr != nil {
			return nil, errors.New("Error parsing 'percent' field " + _percentErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(percent)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(5); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (validityOfTheTimePeriod)
		validityOfTheTimePeriod, _validityOfTheTimePeriodErr := io.ReadBit()
		if _validityOfTheTimePeriodErr != nil {
			return nil, errors.New("Error parsing 'validityOfTheTimePeriod' field " + _validityOfTheTimePeriodErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(validityOfTheTimePeriod)

		// Simple Field (validityOfTheAbsoluteColourTemperature)
		validityOfTheAbsoluteColourTemperature, _validityOfTheAbsoluteColourTemperatureErr := io.ReadBit()
		if _validityOfTheAbsoluteColourTemperatureErr != nil {
			return nil, errors.New("Error parsing 'validityOfTheAbsoluteColourTemperature' field " + _validityOfTheAbsoluteColourTemperatureErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(validityOfTheAbsoluteColourTemperature)

		// Simple Field (validityOfTheAbsoluteBrightness)
		validityOfTheAbsoluteBrightness, _validityOfTheAbsoluteBrightnessErr := io.ReadBit()
		if _validityOfTheAbsoluteBrightnessErr != nil {
			return nil, errors.New("Error parsing 'validityOfTheAbsoluteBrightness' field " + _validityOfTheAbsoluteBrightnessErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(validityOfTheAbsoluteBrightness)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Brightness_Colour_Temperature_Control: // Struct
		_map := map[string]api.PlcValue{}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (cct)
		cct, _cctErr := io.ReadBit()
		if _cctErr != nil {
			return nil, errors.New("Error parsing 'cct' field " + _cctErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cct)

		// Simple Field (stepCodeColourTemperature)
		stepCodeColourTemperature, _stepCodeColourTemperatureErr := io.ReadUint8(3)
		if _stepCodeColourTemperatureErr != nil {
			return nil, errors.New("Error parsing 'stepCodeColourTemperature' field " + _stepCodeColourTemperatureErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepCodeColourTemperature)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (cb)
		cb, _cbErr := io.ReadBit()
		if _cbErr != nil {
			return nil, errors.New("Error parsing 'cb' field " + _cbErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cb)

		// Simple Field (stepCodeBrightness)
		stepCodeBrightness, _stepCodeBrightnessErr := io.ReadUint8(3)
		if _stepCodeBrightnessErr != nil {
			return nil, errors.New("Error parsing 'stepCodeBrightness' field " + _stepCodeBrightnessErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepCodeBrightness)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(6); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (cctAndStepCodeColourValidity)
		cctAndStepCodeColourValidity, _cctAndStepCodeColourValidityErr := io.ReadBit()
		if _cctAndStepCodeColourValidityErr != nil {
			return nil, errors.New("Error parsing 'cctAndStepCodeColourValidity' field " + _cctAndStepCodeColourValidityErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cctAndStepCodeColourValidity)

		// Simple Field (cbAndStepCodeBrightnessValidity)
		cbAndStepCodeBrightnessValidity, _cbAndStepCodeBrightnessValidityErr := io.ReadBit()
		if _cbAndStepCodeBrightnessValidityErr != nil {
			return nil, errors.New("Error parsing 'cbAndStepCodeBrightnessValidity' field " + _cbAndStepCodeBrightnessValidityErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cbAndStepCodeBrightnessValidity)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Colour_RGBW: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (colourLevelRed)
		colourLevelRed, _colourLevelRedErr := io.ReadUint8(8)
		if _colourLevelRedErr != nil {
			return nil, errors.New("Error parsing 'colourLevelRed' field " + _colourLevelRedErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(colourLevelRed)

		// Simple Field (colourLevelGreen)
		colourLevelGreen, _colourLevelGreenErr := io.ReadUint8(8)
		if _colourLevelGreenErr != nil {
			return nil, errors.New("Error parsing 'colourLevelGreen' field " + _colourLevelGreenErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(colourLevelGreen)

		// Simple Field (colourLevelBlue)
		colourLevelBlue, _colourLevelBlueErr := io.ReadUint8(8)
		if _colourLevelBlueErr != nil {
			return nil, errors.New("Error parsing 'colourLevelBlue' field " + _colourLevelBlueErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(colourLevelBlue)

		// Simple Field (colourLevelWhite)
		colourLevelWhite, _colourLevelWhiteErr := io.ReadUint8(8)
		if _colourLevelWhiteErr != nil {
			return nil, errors.New("Error parsing 'colourLevelWhite' field " + _colourLevelWhiteErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(colourLevelWhite)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(8); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (mr)
		mr, _mrErr := io.ReadBit()
		if _mrErr != nil {
			return nil, errors.New("Error parsing 'mr' field " + _mrErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(mr)

		// Simple Field (mg)
		mg, _mgErr := io.ReadBit()
		if _mgErr != nil {
			return nil, errors.New("Error parsing 'mg' field " + _mgErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(mg)

		// Simple Field (mb)
		mb, _mbErr := io.ReadBit()
		if _mbErr != nil {
			return nil, errors.New("Error parsing 'mb' field " + _mbErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(mb)

		// Simple Field (mw)
		mw, _mwErr := io.ReadBit()
		if _mwErr != nil {
			return nil, errors.New("Error parsing 'mw' field " + _mwErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(mw)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Relative_Control_RGBW: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (maskcw)
		maskcw, _maskcwErr := io.ReadBit()
		if _maskcwErr != nil {
			return nil, errors.New("Error parsing 'maskcw' field " + _maskcwErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskcw)

		// Simple Field (maskcb)
		maskcb, _maskcbErr := io.ReadBit()
		if _maskcbErr != nil {
			return nil, errors.New("Error parsing 'maskcb' field " + _maskcbErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskcb)

		// Simple Field (maskcg)
		maskcg, _maskcgErr := io.ReadBit()
		if _maskcgErr != nil {
			return nil, errors.New("Error parsing 'maskcg' field " + _maskcgErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskcg)

		// Simple Field (maskcr)
		maskcr, _maskcrErr := io.ReadBit()
		if _maskcrErr != nil {
			return nil, errors.New("Error parsing 'maskcr' field " + _maskcrErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(maskcr)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (cw)
		cw, _cwErr := io.ReadBit()
		if _cwErr != nil {
			return nil, errors.New("Error parsing 'cw' field " + _cwErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cw)

		// Simple Field (stepCodeColourWhite)
		stepCodeColourWhite, _stepCodeColourWhiteErr := io.ReadUint8(3)
		if _stepCodeColourWhiteErr != nil {
			return nil, errors.New("Error parsing 'stepCodeColourWhite' field " + _stepCodeColourWhiteErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepCodeColourWhite)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (cb)
		cb, _cbErr := io.ReadBit()
		if _cbErr != nil {
			return nil, errors.New("Error parsing 'cb' field " + _cbErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cb)

		// Simple Field (stepCodeColourBlue)
		stepCodeColourBlue, _stepCodeColourBlueErr := io.ReadUint8(3)
		if _stepCodeColourBlueErr != nil {
			return nil, errors.New("Error parsing 'stepCodeColourBlue' field " + _stepCodeColourBlueErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepCodeColourBlue)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (cg)
		cg, _cgErr := io.ReadBit()
		if _cgErr != nil {
			return nil, errors.New("Error parsing 'cg' field " + _cgErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cg)

		// Simple Field (stepCodeColourGreen)
		stepCodeColourGreen, _stepCodeColourGreenErr := io.ReadUint8(3)
		if _stepCodeColourGreenErr != nil {
			return nil, errors.New("Error parsing 'stepCodeColourGreen' field " + _stepCodeColourGreenErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepCodeColourGreen)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (cr)
		cr, _crErr := io.ReadBit()
		if _crErr != nil {
			return nil, errors.New("Error parsing 'cr' field " + _crErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cr)

		// Simple Field (stepCodeColourRed)
		stepCodeColourRed, _stepCodeColourRedErr := io.ReadUint8(3)
		if _stepCodeColourRedErr != nil {
			return nil, errors.New("Error parsing 'stepCodeColourRed' field " + _stepCodeColourRedErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepCodeColourRed)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_Relative_Control_RGB: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (cb)
		cb, _cbErr := io.ReadBit()
		if _cbErr != nil {
			return nil, errors.New("Error parsing 'cb' field " + _cbErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cb)

		// Simple Field (stepCodeColourBlue)
		stepCodeColourBlue, _stepCodeColourBlueErr := io.ReadUint8(3)
		if _stepCodeColourBlueErr != nil {
			return nil, errors.New("Error parsing 'stepCodeColourBlue' field " + _stepCodeColourBlueErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepCodeColourBlue)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (cg)
		cg, _cgErr := io.ReadBit()
		if _cgErr != nil {
			return nil, errors.New("Error parsing 'cg' field " + _cgErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cg)

		// Simple Field (stepCodeColourGreen)
		stepCodeColourGreen, _stepCodeColourGreenErr := io.ReadUint8(3)
		if _stepCodeColourGreenErr != nil {
			return nil, errors.New("Error parsing 'stepCodeColourGreen' field " + _stepCodeColourGreenErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepCodeColourGreen)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}

		// Simple Field (cr)
		cr, _crErr := io.ReadBit()
		if _crErr != nil {
			return nil, errors.New("Error parsing 'cr' field " + _crErr.Error())
		}
		_map["Struct"] = values.NewPlcBOOL(cr)

		// Simple Field (stepCodeColourRed)
		stepCodeColourRed, _stepCodeColourRedErr := io.ReadUint8(3)
		if _stepCodeColourRedErr != nil {
			return nil, errors.New("Error parsing 'stepCodeColourRed' field " + _stepCodeColourRedErr.Error())
		}
		_map["Struct"] = values.NewPlcUSINT(stepCodeColourRed)

		// Reserved Field (Just skip the bytes)
		if _, _err := io.ReadUint8(4); _err != nil {
			return nil, errors.New("Error parsing reserved field " + _err.Error())
		}
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_GeographicalLocation: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (longitude)
		longitude, _longitudeErr := io.ReadFloat32(true, 8, 23)
		if _longitudeErr != nil {
			return nil, errors.New("Error parsing 'longitude' field " + _longitudeErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(longitude)

		// Simple Field (latitude)
		latitude, _latitudeErr := io.ReadFloat32(true, 8, 23)
		if _latitudeErr != nil {
			return nil, errors.New("Error parsing 'latitude' field " + _latitudeErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(latitude)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_TempRoomSetpSetF16_4: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (roomTemperatureSetpointComfort)
		roomTemperatureSetpointComfort, _roomTemperatureSetpointComfortErr := io.ReadFloat32(true, 4, 11)
		if _roomTemperatureSetpointComfortErr != nil {
			return nil, errors.New("Error parsing 'roomTemperatureSetpointComfort' field " + _roomTemperatureSetpointComfortErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(roomTemperatureSetpointComfort)

		// Simple Field (roomTemperatureSetpointStandby)
		roomTemperatureSetpointStandby, _roomTemperatureSetpointStandbyErr := io.ReadFloat32(true, 4, 11)
		if _roomTemperatureSetpointStandbyErr != nil {
			return nil, errors.New("Error parsing 'roomTemperatureSetpointStandby' field " + _roomTemperatureSetpointStandbyErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(roomTemperatureSetpointStandby)

		// Simple Field (roomTemperatureSetpointEconomy)
		roomTemperatureSetpointEconomy, _roomTemperatureSetpointEconomyErr := io.ReadFloat32(true, 4, 11)
		if _roomTemperatureSetpointEconomyErr != nil {
			return nil, errors.New("Error parsing 'roomTemperatureSetpointEconomy' field " + _roomTemperatureSetpointEconomyErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(roomTemperatureSetpointEconomy)

		// Simple Field (roomTemperatureSetpointBuildingProtection)
		roomTemperatureSetpointBuildingProtection, _roomTemperatureSetpointBuildingProtectionErr := io.ReadFloat32(true, 4, 11)
		if _roomTemperatureSetpointBuildingProtectionErr != nil {
			return nil, errors.New("Error parsing 'roomTemperatureSetpointBuildingProtection' field " + _roomTemperatureSetpointBuildingProtectionErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(roomTemperatureSetpointBuildingProtection)
		return values.NewPlcStruct(_map), nil
	case datapointType == KnxDatapointType_DPT_TempRoomSetpSetShiftF16_4: // Struct
		_map := map[string]api.PlcValue{}

		// Simple Field (roomTemperatureSetpointShiftComfort)
		roomTemperatureSetpointShiftComfort, _roomTemperatureSetpointShiftComfortErr := io.ReadFloat32(true, 4, 11)
		if _roomTemperatureSetpointShiftComfortErr != nil {
			return nil, errors.New("Error parsing 'roomTemperatureSetpointShiftComfort' field " + _roomTemperatureSetpointShiftComfortErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(roomTemperatureSetpointShiftComfort)

		// Simple Field (roomTemperatureSetpointShiftStandby)
		roomTemperatureSetpointShiftStandby, _roomTemperatureSetpointShiftStandbyErr := io.ReadFloat32(true, 4, 11)
		if _roomTemperatureSetpointShiftStandbyErr != nil {
			return nil, errors.New("Error parsing 'roomTemperatureSetpointShiftStandby' field " + _roomTemperatureSetpointShiftStandbyErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(roomTemperatureSetpointShiftStandby)

		// Simple Field (roomTemperatureSetpointShiftEconomy)
		roomTemperatureSetpointShiftEconomy, _roomTemperatureSetpointShiftEconomyErr := io.ReadFloat32(true, 4, 11)
		if _roomTemperatureSetpointShiftEconomyErr != nil {
			return nil, errors.New("Error parsing 'roomTemperatureSetpointShiftEconomy' field " + _roomTemperatureSetpointShiftEconomyErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(roomTemperatureSetpointShiftEconomy)

		// Simple Field (roomTemperatureSetpointShiftBuildingProtection)
		roomTemperatureSetpointShiftBuildingProtection, _roomTemperatureSetpointShiftBuildingProtectionErr := io.ReadFloat32(true, 4, 11)
		if _roomTemperatureSetpointShiftBuildingProtectionErr != nil {
			return nil, errors.New("Error parsing 'roomTemperatureSetpointShiftBuildingProtection' field " + _roomTemperatureSetpointShiftBuildingProtectionErr.Error())
		}
		_map["Struct"] = values.NewPlcREAL(roomTemperatureSetpointShiftBuildingProtection)
		return values.NewPlcStruct(_map), nil
	}
	return nil, errors.New("unsupported type")
}

func KnxDatapointSerialize(io *utils.WriteBuffer, value api.PlcValue, datapointType KnxDatapointType) error {
	switch {
	case datapointType == KnxDatapointType_BOOL: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_BYTE: // BYTE

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_WORD: // WORD

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DWORD: // DWORD

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_LWORD: // LWORD

		// Simple Field (value)
		if _err := io.WriteUint64(64, value.GetUint64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_USINT: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_SINT: // SINT

		// Simple Field (value)
		if _err := io.WriteInt8(8, value.GetInt8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_UINT: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_INT: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_UDINT: // UDINT

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DINT: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_ULINT: // ULINT

		// Simple Field (value)
		if _err := io.WriteUint64(64, value.GetUint64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_LINT: // LINT

		// Simple Field (value)
		if _err := io.WriteInt64(64, value.GetInt64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_REAL: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_LREAL: // LREAL

		// Simple Field (value)
		if _err := io.WriteFloat64(64, value.GetFloat64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_CHAR: // CHAR

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_WCHAR: // WCHAR

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_TIME: // TIME

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_LTIME: // LTIME

		// Simple Field (value)
		if _err := io.WriteUint64(64, value.GetUint64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DATE: // DATE

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_TIME_OF_DAY: // TIME_OF_DAY

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_TOD: // TIME_OF_DAY

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DATE_AND_TIME: // DATE_AND_TIME

		// Simple Field (year)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'year' field " + _err.Error())
		}

		// Simple Field (month)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'month' field " + _err.Error())
		}

		// Simple Field (day)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'day' field " + _err.Error())
		}

		// Simple Field (dayOfWeek)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'dayOfWeek' field " + _err.Error())
		}

		// Simple Field (hour)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'hour' field " + _err.Error())
		}

		// Simple Field (minutes)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'minutes' field " + _err.Error())
		}

		// Simple Field (seconds)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'seconds' field " + _err.Error())
		}

		// Simple Field (nanos)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'nanos' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DT: // DATE_AND_TIME

		// Simple Field (year)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'year' field " + _err.Error())
		}

		// Simple Field (month)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'month' field " + _err.Error())
		}

		// Simple Field (day)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'day' field " + _err.Error())
		}

		// Simple Field (dayOfWeek)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'dayOfWeek' field " + _err.Error())
		}

		// Simple Field (hour)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'hour' field " + _err.Error())
		}

		// Simple Field (minutes)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'minutes' field " + _err.Error())
		}

		// Simple Field (seconds)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'seconds' field " + _err.Error())
		}

		// Simple Field (nanos)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'nanos' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Switch: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Bool: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Enable: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Ramp: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Alarm: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_BinaryValue: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Step: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_UpDown: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_OpenClose: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Start: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_State: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Invert: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DimSendStyle: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_InputSource: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Reset: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Ack: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Trigger: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Occupancy: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Window_Door: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LogicalFunction: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Scene_AB: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ShutterBlinds_Mode: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DayNight: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Heat_Cool: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Switch_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (on)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'on' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Bool_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (valueTrue)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'valueTrue' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Enable_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (enable)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'enable' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Ramp_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (ramp)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'ramp' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Alarm_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (alarm)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'alarm' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_BinaryValue_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (high)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'high' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Step_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (increase)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'increase' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Direction1_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (down)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'down' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Direction2_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (close)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'close' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Start_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (start)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'start' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_State_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (active)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'active' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Invert_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (control)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'control' field " + _err.Error())
		}

		// Simple Field (inverted)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'inverted' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Control_Dimming: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (increase)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'increase' field " + _err.Error())
		}

		// Simple Field (stepcode)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepcode' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Control_Blinds: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (down)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'down' field " + _err.Error())
		}

		// Simple Field (stepcode)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepcode' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Char_ASCII: // STRING

		// Simple Field (value)
		if _err := io.WriteString(8, "ASCII", value.GetString()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Char_8859_1: // STRING

		// Simple Field (value)
		if _err := io.WriteString(8, "ISO-8859-1", value.GetString()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Scaling: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Angle: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Percent_U8: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DecimalFactor: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Tariff: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_1_Ucount: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_FanStage: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Percent_V8: // SINT

		// Simple Field (value)
		if _err := io.WriteInt8(8, value.GetInt8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_1_Count: // SINT

		// Simple Field (value)
		if _err := io.WriteInt8(8, value.GetInt8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Status_Mode3: // Struct

		// Simple Field (statusA)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusA' field " + _err.Error())
		}

		// Simple Field (statusB)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusB' field " + _err.Error())
		}

		// Simple Field (statusC)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusC' field " + _err.Error())
		}

		// Simple Field (statusD)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusD' field " + _err.Error())
		}

		// Simple Field (statusE)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusE' field " + _err.Error())
		}

		// Simple Field (mode)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'mode' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_2_Ucount: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TimePeriodMsec: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TimePeriod10Msec: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TimePeriod100Msec: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TimePeriodSec: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TimePeriodMin: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TimePeriodHrs: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_PropDataType: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Length_mm: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_UElCurrentmA: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Brightness: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Absolute_Colour_Temperature: // UINT

		// Simple Field (value)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_2_Count: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DeltaTimeMsec: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DeltaTime10Msec: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DeltaTime100Msec: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DeltaTimeSec: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DeltaTimeMin: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DeltaTimeHrs: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Percent_V16: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Rotation_Angle: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Length_m: // INT

		// Simple Field (value)
		if _err := io.WriteInt16(16, value.GetInt16()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Temp: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Tempd: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Tempa: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Lux: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Wsp: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Pres: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Humidity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_AirQuality: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_AirFlow: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Time1: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Time2: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Volt: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Curr: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_PowerDensity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_KelvinPerPercent: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Power: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Volume_Flow: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Rain_Amount: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Temp_F: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Wsp_kmh: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Absolute_Humidity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Concentration_ygm3: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TimeOfDay: // Struct

		// Simple Field (day)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'day' field " + _err.Error())
		}

		// Simple Field (hour)
		if _err := io.WriteUint8(5, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'hour' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(2, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (minutes)
		if _err := io.WriteUint8(6, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'minutes' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(2, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (seconds)
		if _err := io.WriteUint8(6, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'seconds' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Date: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(3, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (dayOfMonth)
		if _err := io.WriteUint8(5, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'dayOfMonth' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (month)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'month' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(1, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (year)
		if _err := io.WriteUint8(7, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'year' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_4_Ucount: // UDINT

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LongTimePeriod_Sec: // UDINT

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LongTimePeriod_Min: // UDINT

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LongTimePeriod_Hrs: // UDINT

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_VolumeLiquid_Litre: // UDINT

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Volume_m_3: // UDINT

		// Simple Field (value)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_4_Count: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_FlowRate_m3h: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ActiveEnergy: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ApparantEnergy: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ReactiveEnergy: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ActiveEnergy_kWh: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ApparantEnergy_kVAh: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ReactiveEnergy_kVARh: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ActiveEnergy_MWh: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LongDeltaTimeSec: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DeltaVolumeLiquid_Litre: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DeltaVolume_m_3: // DINT

		// Simple Field (value)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Acceleration: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Acceleration_Angular: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Activation_Energy: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Activity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Mol: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Amplitude: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_AngleRad: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_AngleDeg: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Angular_Momentum: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Angular_Velocity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Area: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Capacitance: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Charge_DensitySurface: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Charge_DensityVolume: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Compressibility: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Conductance: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electrical_Conductivity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Density: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_Charge: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_Current: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_CurrentDensity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_DipoleMoment: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_Displacement: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_FieldStrength: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_Flux: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_FluxDensity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_Polarization: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_Potential: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electric_PotentialDifference: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_ElectromagneticMoment: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Electromotive_Force: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Energy: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Force: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Frequency: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Angular_Frequency: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Heat_Capacity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Heat_FlowRate: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Heat_Quantity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Impedance: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Length: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Light_Quantity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Luminance: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Luminous_Flux: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Luminous_Intensity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Magnetic_FieldStrength: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Magnetic_Flux: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Magnetic_FluxDensity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Magnetic_Moment: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Magnetic_Polarization: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Magnetization: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_MagnetomotiveForce: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Mass: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_MassFlux: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Momentum: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Phase_AngleRad: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Phase_AngleDeg: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Power: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Power_Factor: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Pressure: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Reactance: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Resistance: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Resistivity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_SelfInductance: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_SolidAngle: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Sound_Intensity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Speed: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Stress: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Surface_Tension: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Common_Temperature: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Absolute_Temperature: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_TemperatureDifference: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Thermal_Capacity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Thermal_Conductivity: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_ThermoelectricPower: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Time: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Torque: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Volume: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Volume_Flux: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Weight: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Value_Work: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Volume_Flux_Meter: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Volume_Flux_ls: // REAL

		// Simple Field (value)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Access_Data: // Struct

		// Simple Field (hurz)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'hurz' field " + _err.Error())
		}

		// Simple Field (value1)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value1' field " + _err.Error())
		}

		// Simple Field (value2)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value2' field " + _err.Error())
		}

		// Simple Field (value3)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value3' field " + _err.Error())
		}

		// Simple Field (value4)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value4' field " + _err.Error())
		}

		// Simple Field (value5)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value5' field " + _err.Error())
		}

		// Simple Field (detectionError)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'detectionError' field " + _err.Error())
		}

		// Simple Field (permission)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'permission' field " + _err.Error())
		}

		// Simple Field (readDirection)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'readDirection' field " + _err.Error())
		}

		// Simple Field (encryptionOfAccessInformation)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'encryptionOfAccessInformation' field " + _err.Error())
		}

		// Simple Field (indexOfAccessIdentificationCode)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'indexOfAccessIdentificationCode' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_String_ASCII: // STRING

		// Simple Field (value)
		if _err := io.WriteString(112, "ASCII", value.GetString()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_String_8859_1: // STRING

		// Simple Field (value)
		if _err := io.WriteString(112, "ISO-8859-1", value.GetString()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_SceneNumber: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(6, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_SceneControl: // Struct

		// Simple Field (learnTheSceneCorrespondingToTheFieldSceneNumber)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'learnTheSceneCorrespondingToTheFieldSceneNumber' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(1, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (sceneNumber)
		if _err := io.WriteUint8(6, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'sceneNumber' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DateTime: // Struct

		// Simple Field (year)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'year' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (month)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'month' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(3, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (dayofmonth)
		if _err := io.WriteUint8(5, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'dayofmonth' field " + _err.Error())
		}

		// Simple Field (dayofweek)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'dayofweek' field " + _err.Error())
		}

		// Simple Field (hourofday)
		if _err := io.WriteUint8(5, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'hourofday' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(2, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (minutes)
		if _err := io.WriteUint8(6, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'minutes' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(2, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (seconds)
		if _err := io.WriteUint8(6, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'seconds' field " + _err.Error())
		}

		// Simple Field (fault)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'fault' field " + _err.Error())
		}

		// Simple Field (workingDay)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'workingDay' field " + _err.Error())
		}

		// Simple Field (noWd)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'noWd' field " + _err.Error())
		}

		// Simple Field (noYear)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'noYear' field " + _err.Error())
		}

		// Simple Field (noDate)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'noDate' field " + _err.Error())
		}

		// Simple Field (noDayOfWeek)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'noDayOfWeek' field " + _err.Error())
		}

		// Simple Field (noTime)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'noTime' field " + _err.Error())
		}

		// Simple Field (standardSummerTime)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'standardSummerTime' field " + _err.Error())
		}

		// Simple Field (qualityOfClock)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'qualityOfClock' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_SCLOMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_BuildingMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_OccMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Priority: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LightApplicationMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ApplicationArea: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_AlarmClassType: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_PSUMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ErrorClass_System: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ErrorClass_HVAC: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Time_Delay: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Beaufort_Wind_Force_Scale: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_SensorSelect: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ActuatorConnectType: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Cloud_Cover: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_PowerReturnMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_FuelType: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_BurnerType: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_HVACMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DHWMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LoadPriority: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_HVACContrMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_HVACEmergMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ChangeoverMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ValveMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DamperMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_HeaterMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_FanMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_MasterSlaveMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_StatusRoomSetp: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Metering_DeviceType: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_HumDehumMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_EnableHCStage: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ADAType: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_BackupMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_StartSynchronization: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Behaviour_Lock_Unlock: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Behaviour_Bus_Power_Up_Down: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DALI_Fade_Time: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_BlinkingMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LightControlMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_SwitchPBModel: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_PBAction: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DimmPBModel: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_SwitchOnMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LoadTypeSet: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LoadTypeDetected: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Converter_Test_Control: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_SABExcept_Behaviour: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_SABBehaviour_Lock_Unlock: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_SSSBMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_BlindsControlMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_CommMode: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_AddInfoTypes: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_RF_ModeSelect: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_RF_FilterSelect: // USINT

		// Simple Field (value)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_StatusGen: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(3, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (alarmStatusOfCorrespondingDatapointIsNotAcknowledged)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'alarmStatusOfCorrespondingDatapointIsNotAcknowledged' field " + _err.Error())
		}

		// Simple Field (correspondingDatapointIsInAlarm)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'correspondingDatapointIsInAlarm' field " + _err.Error())
		}

		// Simple Field (correspondingDatapointMainValueIsOverridden)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'correspondingDatapointMainValueIsOverridden' field " + _err.Error())
		}

		// Simple Field (correspondingDatapointMainValueIsCorruptedDueToFailure)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'correspondingDatapointMainValueIsCorruptedDueToFailure' field " + _err.Error())
		}

		// Simple Field (correspondingDatapointValueIsOutOfService)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'correspondingDatapointValueIsOutOfService' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Device_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(5, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (verifyModeIsOn)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'verifyModeIsOn' field " + _err.Error())
		}

		// Simple Field (aDatagramWithTheOwnIndividualAddressAsSourceAddressHasBeenReceived)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'aDatagramWithTheOwnIndividualAddressAsSourceAddressHasBeenReceived' field " + _err.Error())
		}

		// Simple Field (theUserApplicationIsStopped)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'theUserApplicationIsStopped' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ForceSign: // Struct

		// Simple Field (roomhmax)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'roomhmax' field " + _err.Error())
		}

		// Simple Field (roomhconf)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'roomhconf' field " + _err.Error())
		}

		// Simple Field (dhwlegio)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'dhwlegio' field " + _err.Error())
		}

		// Simple Field (dhwnorm)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'dhwnorm' field " + _err.Error())
		}

		// Simple Field (overrun)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'overrun' field " + _err.Error())
		}

		// Simple Field (oversupply)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'oversupply' field " + _err.Error())
		}

		// Simple Field (protection)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'protection' field " + _err.Error())
		}

		// Simple Field (forcerequest)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'forcerequest' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ForceSignCool: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_StatusRHC: // Struct

		// Simple Field (summermode)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'summermode' field " + _err.Error())
		}

		// Simple Field (statusstopoptim)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusstopoptim' field " + _err.Error())
		}

		// Simple Field (statusstartoptim)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusstartoptim' field " + _err.Error())
		}

		// Simple Field (statusmorningboost)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusmorningboost' field " + _err.Error())
		}

		// Simple Field (tempreturnlimit)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'tempreturnlimit' field " + _err.Error())
		}

		// Simple Field (tempflowlimit)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'tempflowlimit' field " + _err.Error())
		}

		// Simple Field (satuseco)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'satuseco' field " + _err.Error())
		}

		// Simple Field (fault)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'fault' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_StatusSDHWC: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(5, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (solarloadsufficient)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'solarloadsufficient' field " + _err.Error())
		}

		// Simple Field (sdhwloadactive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'sdhwloadactive' field " + _err.Error())
		}

		// Simple Field (fault)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'fault' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_FuelTypeSet: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(5, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (solidstate)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'solidstate' field " + _err.Error())
		}

		// Simple Field (gas)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'gas' field " + _err.Error())
		}

		// Simple Field (oil)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'oil' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_StatusRCC: // BOOL

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_StatusAHU: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (cool)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cool' field " + _err.Error())
		}

		// Simple Field (heat)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'heat' field " + _err.Error())
		}

		// Simple Field (fanactive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'fanactive' field " + _err.Error())
		}

		// Simple Field (fault)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'fault' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_CombinedStatus_RTSM: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(3, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (statusOfHvacModeUser)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusOfHvacModeUser' field " + _err.Error())
		}

		// Simple Field (statusOfComfortProlongationUser)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusOfComfortProlongationUser' field " + _err.Error())
		}

		// Simple Field (effectiveValueOfTheComfortPushButton)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'effectiveValueOfTheComfortPushButton' field " + _err.Error())
		}

		// Simple Field (effectiveValueOfThePresenceStatus)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'effectiveValueOfThePresenceStatus' field " + _err.Error())
		}

		// Simple Field (effectiveValueOfTheWindowStatus)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'effectiveValueOfTheWindowStatus' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LightActuatorErrorInfo: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(1, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (overheat)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'overheat' field " + _err.Error())
		}

		// Simple Field (lampfailure)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'lampfailure' field " + _err.Error())
		}

		// Simple Field (defectiveload)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'defectiveload' field " + _err.Error())
		}

		// Simple Field (underload)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'underload' field " + _err.Error())
		}

		// Simple Field (overcurrent)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'overcurrent' field " + _err.Error())
		}

		// Simple Field (undervoltage)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'undervoltage' field " + _err.Error())
		}

		// Simple Field (loaddetectionerror)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'loaddetectionerror' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_RF_ModeInfo: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(5, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (bibatSlave)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'bibatSlave' field " + _err.Error())
		}

		// Simple Field (bibatMaster)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'bibatMaster' field " + _err.Error())
		}

		// Simple Field (asynchronous)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'asynchronous' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_RF_FilterInfo: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(5, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (doa)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'doa' field " + _err.Error())
		}

		// Simple Field (knxSn)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'knxSn' field " + _err.Error())
		}

		// Simple Field (doaAndKnxSn)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'doaAndKnxSn' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Channel_Activation_8: // Struct

		// Simple Field (activationStateOfChannel1)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel1' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel2)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel2' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel3)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel3' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel4)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel4' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel5)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel5' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel6)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel6' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel7)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel7' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel8)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel8' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_StatusDHWC: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(8, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (tempoptimshiftactive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'tempoptimshiftactive' field " + _err.Error())
		}

		// Simple Field (solarenergysupport)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'solarenergysupport' field " + _err.Error())
		}

		// Simple Field (solarenergyonly)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'solarenergyonly' field " + _err.Error())
		}

		// Simple Field (otherenergysourceactive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'otherenergysourceactive' field " + _err.Error())
		}

		// Simple Field (dhwpushactive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'dhwpushactive' field " + _err.Error())
		}

		// Simple Field (legioprotactive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'legioprotactive' field " + _err.Error())
		}

		// Simple Field (dhwloadactive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'dhwloadactive' field " + _err.Error())
		}

		// Simple Field (fault)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'fault' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_StatusRHCC: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(1, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (overheatalarm)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'overheatalarm' field " + _err.Error())
		}

		// Simple Field (frostalarm)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'frostalarm' field " + _err.Error())
		}

		// Simple Field (dewpointstatus)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'dewpointstatus' field " + _err.Error())
		}

		// Simple Field (coolingdisabled)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'coolingdisabled' field " + _err.Error())
		}

		// Simple Field (statusprecool)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusprecool' field " + _err.Error())
		}

		// Simple Field (statusecoc)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusecoc' field " + _err.Error())
		}

		// Simple Field (heatcoolmode)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'heatcoolmode' field " + _err.Error())
		}

		// Simple Field (heatingdiabled)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'heatingdiabled' field " + _err.Error())
		}

		// Simple Field (statusstopoptim)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusstopoptim' field " + _err.Error())
		}

		// Simple Field (statusstartoptim)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusstartoptim' field " + _err.Error())
		}

		// Simple Field (statusmorningboosth)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusmorningboosth' field " + _err.Error())
		}

		// Simple Field (tempflowreturnlimit)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'tempflowreturnlimit' field " + _err.Error())
		}

		// Simple Field (tempflowlimit)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'tempflowlimit' field " + _err.Error())
		}

		// Simple Field (statusecoh)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'statusecoh' field " + _err.Error())
		}

		// Simple Field (fault)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'fault' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_CombinedStatus_HVA: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (calibrationMode)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'calibrationMode' field " + _err.Error())
		}

		// Simple Field (lockedPosition)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'lockedPosition' field " + _err.Error())
		}

		// Simple Field (forcedPosition)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'forcedPosition' field " + _err.Error())
		}

		// Simple Field (manuaOperationOverridden)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'manuaOperationOverridden' field " + _err.Error())
		}

		// Simple Field (serviceMode)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'serviceMode' field " + _err.Error())
		}

		// Simple Field (valveKick)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'valveKick' field " + _err.Error())
		}

		// Simple Field (overload)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'overload' field " + _err.Error())
		}

		// Simple Field (shortCircuit)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'shortCircuit' field " + _err.Error())
		}

		// Simple Field (currentValvePosition)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'currentValvePosition' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_CombinedStatus_RTC: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(7, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (coolingModeEnabled)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'coolingModeEnabled' field " + _err.Error())
		}

		// Simple Field (heatingModeEnabled)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'heatingModeEnabled' field " + _err.Error())
		}

		// Simple Field (additionalHeatingCoolingStage2Stage)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'additionalHeatingCoolingStage2Stage' field " + _err.Error())
		}

		// Simple Field (controllerInactive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'controllerInactive' field " + _err.Error())
		}

		// Simple Field (overheatAlarm)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'overheatAlarm' field " + _err.Error())
		}

		// Simple Field (frostAlarm)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'frostAlarm' field " + _err.Error())
		}

		// Simple Field (dewPointStatus)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'dewPointStatus' field " + _err.Error())
		}

		// Simple Field (activeMode)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activeMode' field " + _err.Error())
		}

		// Simple Field (generalFailureInformation)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'generalFailureInformation' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Media: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint16(10, uint16(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (knxIp)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'knxIp' field " + _err.Error())
		}

		// Simple Field (rf)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'rf' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(1, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (pl110)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'pl110' field " + _err.Error())
		}

		// Simple Field (tp1)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'tp1' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(1, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Channel_Activation_16: // Struct

		// Simple Field (activationStateOfChannel1)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel1' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel2)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel2' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel3)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel3' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel4)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel4' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel5)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel5' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel6)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel6' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel7)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel7' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel8)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel8' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel9)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel9' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel10)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel10' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel11)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel11' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel12)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel12' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel13)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel13' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel14)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel14' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel15)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel15' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel16)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel16' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_OnOffAction: // USINT

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Alarm_Reaction: // USINT

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_UpDown_Action: // USINT

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_HVAC_PB_Action: // USINT

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (value)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DoubleNibble: // Struct

		// Simple Field (busy)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'busy' field " + _err.Error())
		}

		// Simple Field (nak)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'nak' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_SceneInfo: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(1, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (sceneIsInactive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'sceneIsInactive' field " + _err.Error())
		}

		// Simple Field (scenenumber)
		if _err := io.WriteUint8(6, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'scenenumber' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_CombinedInfoOnOff: // Struct

		// Simple Field (maskBitInfoOnOffOutput16)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput16' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput15)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput15' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput14)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput14' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput13)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput13' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput12)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput12' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput11)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput11' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput10)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput10' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput9)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput9' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput8)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput8' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput7)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput7' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput6)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput6' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput5)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput5' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput4)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput4' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput3)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput3' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput2)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput2' field " + _err.Error())
		}

		// Simple Field (maskBitInfoOnOffOutput1)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskBitInfoOnOffOutput1' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput16)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput16' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput15)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput15' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput14)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput14' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput13)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput13' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput12)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput12' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput11)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput11' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput10)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput10' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput9)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput9' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput8)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput8' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput7)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput7' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput6)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput6' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput5)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput5' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput4)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput4' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput3)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput3' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput2)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput2' field " + _err.Error())
		}

		// Simple Field (infoOnOffOutput1)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'infoOnOffOutput1' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ActiveEnergy_V64: // LINT

		// Simple Field (value)
		if _err := io.WriteInt64(64, value.GetInt64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ApparantEnergy_V64: // LINT

		// Simple Field (value)
		if _err := io.WriteInt64(64, value.GetInt64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_ReactiveEnergy_V64: // LINT

		// Simple Field (value)
		if _err := io.WriteInt64(64, value.GetInt64()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Channel_Activation_24: // Struct

		// Simple Field (activationStateOfChannel1)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel1' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel2)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel2' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel3)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel3' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel4)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel4' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel5)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel5' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel6)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel6' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel7)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel7' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel8)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel8' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel9)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel9' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel10)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel10' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel11)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel11' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel12)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel12' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel13)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel13' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel14)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel14' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel15)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel15' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel16)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel16' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel17)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel17' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel18)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel18' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel19)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel19' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel20)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel20' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel21)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel21' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel22)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel22' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel23)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel23' field " + _err.Error())
		}

		// Simple Field (activationStateOfChannel24)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'activationStateOfChannel24' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_HVACModeNext: // Struct

		// Simple Field (delayTimeMin)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'delayTimeMin' field " + _err.Error())
		}

		// Simple Field (hvacMode)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'hvacMode' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DHWModeNext: // Struct

		// Simple Field (delayTimeMin)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'delayTimeMin' field " + _err.Error())
		}

		// Simple Field (dhwMode)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'dhwMode' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_OccModeNext: // Struct

		// Simple Field (delayTimeMin)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'delayTimeMin' field " + _err.Error())
		}

		// Simple Field (occupancyMode)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'occupancyMode' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_BuildingModeNext: // Struct

		// Simple Field (delayTimeMin)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'delayTimeMin' field " + _err.Error())
		}

		// Simple Field (buildingMode)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'buildingMode' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Version: // Struct

		// Simple Field (magicNumber)
		if _err := io.WriteUint8(5, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'magicNumber' field " + _err.Error())
		}

		// Simple Field (versionNumber)
		if _err := io.WriteUint8(5, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'versionNumber' field " + _err.Error())
		}

		// Simple Field (revisionNumber)
		if _err := io.WriteUint8(6, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'revisionNumber' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_AlarmInfo: // Struct

		// Simple Field (logNumber)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'logNumber' field " + _err.Error())
		}

		// Simple Field (alarmPriority)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'alarmPriority' field " + _err.Error())
		}

		// Simple Field (applicationArea)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'applicationArea' field " + _err.Error())
		}

		// Simple Field (errorClass)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'errorClass' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (errorcodeSup)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'errorcodeSup' field " + _err.Error())
		}

		// Simple Field (alarmtextSup)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'alarmtextSup' field " + _err.Error())
		}

		// Simple Field (timestampSup)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'timestampSup' field " + _err.Error())
		}

		// Simple Field (ackSup)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'ackSup' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(5, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (locked)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'locked' field " + _err.Error())
		}

		// Simple Field (alarmunack)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'alarmunack' field " + _err.Error())
		}

		// Simple Field (inalarm)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'inalarm' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TempRoomSetpSetF16_3: // Struct

		// Simple Field (tempsetpcomf)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'tempsetpcomf' field " + _err.Error())
		}

		// Simple Field (tempsetpstdby)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'tempsetpstdby' field " + _err.Error())
		}

		// Simple Field (tempsetpeco)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'tempsetpeco' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TempRoomSetpSetShiftF16_3: // Struct

		// Simple Field (tempsetpshiftcomf)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'tempsetpshiftcomf' field " + _err.Error())
		}

		// Simple Field (tempsetpshiftstdby)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'tempsetpshiftstdby' field " + _err.Error())
		}

		// Simple Field (tempsetpshifteco)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'tempsetpshifteco' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Scaling_Speed: // Struct

		// Simple Field (timePeriod)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'timePeriod' field " + _err.Error())
		}

		// Simple Field (percent)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'percent' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Scaling_Step_Time: // Struct

		// Simple Field (timePeriod)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'timePeriod' field " + _err.Error())
		}

		// Simple Field (percent)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'percent' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_MeteringValue: // Struct

		// Simple Field (countval)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'countval' field " + _err.Error())
		}

		// Simple Field (valinffield)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'valinffield' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(3, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (alarmunack)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'alarmunack' field " + _err.Error())
		}

		// Simple Field (inalarm)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'inalarm' field " + _err.Error())
		}

		// Simple Field (overridden)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'overridden' field " + _err.Error())
		}

		// Simple Field (fault)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'fault' field " + _err.Error())
		}

		// Simple Field (outofservice)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'outofservice' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_MBus_Address: // Struct

		// Simple Field (manufactid)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'manufactid' field " + _err.Error())
		}

		// Simple Field (identnumber)
		if _err := io.WriteUint32(32, value.GetUint32()); _err != nil {
			return errors.New("Error serializing 'identnumber' field " + _err.Error())
		}

		// Simple Field (version)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'version' field " + _err.Error())
		}

		// Simple Field (medium)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'medium' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Colour_RGB: // Struct

		// Simple Field (r)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'r' field " + _err.Error())
		}

		// Simple Field (g)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'g' field " + _err.Error())
		}

		// Simple Field (b)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'b' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_LanguageCodeAlpha2_ASCII: // STRING

		// Simple Field (value)
		if _err := io.WriteString(16, "ASCII", value.GetString()); _err != nil {
			return errors.New("Error serializing 'value' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Tariff_ActiveEnergy: // Struct

		// Simple Field (activeelectricalenergy)
		if _err := io.WriteInt32(32, value.GetInt32()); _err != nil {
			return errors.New("Error serializing 'activeelectricalenergy' field " + _err.Error())
		}

		// Simple Field (tariff)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'tariff' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (electricalengergyvalidity)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'electricalengergyvalidity' field " + _err.Error())
		}

		// Simple Field (tariffvalidity)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'tariffvalidity' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Prioritised_Mode_Control: // Struct

		// Simple Field (deactivationOfPriority)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'deactivationOfPriority' field " + _err.Error())
		}

		// Simple Field (priorityLevel)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'priorityLevel' field " + _err.Error())
		}

		// Simple Field (modeLevel)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'modeLevel' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DALI_Control_Gear_Diagnostic: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(5, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (convertorError)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'convertorError' field " + _err.Error())
		}

		// Simple Field (ballastFailure)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'ballastFailure' field " + _err.Error())
		}

		// Simple Field (lampFailure)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'lampFailure' field " + _err.Error())
		}

		// Simple Field (readOrResponse)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'readOrResponse' field " + _err.Error())
		}

		// Simple Field (addressIndicator)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'addressIndicator' field " + _err.Error())
		}

		// Simple Field (daliDeviceAddressOrDaliGroupAddress)
		if _err := io.WriteUint8(6, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'daliDeviceAddressOrDaliGroupAddress' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_DALI_Diagnostics: // Struct

		// Simple Field (ballastFailure)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'ballastFailure' field " + _err.Error())
		}

		// Simple Field (lampFailure)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'lampFailure' field " + _err.Error())
		}

		// Simple Field (deviceAddress)
		if _err := io.WriteUint8(6, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'deviceAddress' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_CombinedPosition: // Struct

		// Simple Field (heightPosition)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'heightPosition' field " + _err.Error())
		}

		// Simple Field (slatsPosition)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'slatsPosition' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (validityHeightPosition)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'validityHeightPosition' field " + _err.Error())
		}

		// Simple Field (validitySlatsPosition)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'validitySlatsPosition' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_StatusSAB: // Struct

		// Simple Field (heightPosition)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'heightPosition' field " + _err.Error())
		}

		// Simple Field (slatsPosition)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'slatsPosition' field " + _err.Error())
		}

		// Simple Field (upperEndPosReached)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'upperEndPosReached' field " + _err.Error())
		}

		// Simple Field (lowerEndPosReached)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'lowerEndPosReached' field " + _err.Error())
		}

		// Simple Field (lowerPredefPosReachedTypHeight100PercentSlatsAngle100Percent)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'lowerPredefPosReachedTypHeight100PercentSlatsAngle100Percent' field " + _err.Error())
		}

		// Simple Field (targetPosDrive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'targetPosDrive' field " + _err.Error())
		}

		// Simple Field (restrictionOfTargetHeightPosPosCanNotBeReached)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'restrictionOfTargetHeightPosPosCanNotBeReached' field " + _err.Error())
		}

		// Simple Field (restrictionOfSlatsHeightPosPosCanNotBeReached)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'restrictionOfSlatsHeightPosPosCanNotBeReached' field " + _err.Error())
		}

		// Simple Field (atLeastOneOfTheInputsWindRainFrostAlarmIsInAlarm)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'atLeastOneOfTheInputsWindRainFrostAlarmIsInAlarm' field " + _err.Error())
		}

		// Simple Field (upDownPositionIsForcedByMoveupdownforcedInput)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'upDownPositionIsForcedByMoveupdownforcedInput' field " + _err.Error())
		}

		// Simple Field (movementIsLockedEGByDevicelockedInput)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'movementIsLockedEGByDevicelockedInput' field " + _err.Error())
		}

		// Simple Field (actuatorSetvalueIsLocallyOverriddenEGViaALocalUserInterface)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'actuatorSetvalueIsLocallyOverriddenEGViaALocalUserInterface' field " + _err.Error())
		}

		// Simple Field (generalFailureOfTheActuatorOrTheDrive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'generalFailureOfTheActuatorOrTheDrive' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(3, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (validityHeightPos)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'validityHeightPos' field " + _err.Error())
		}

		// Simple Field (validitySlatsPos)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'validitySlatsPos' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Colour_xyY: // Struct

		// Simple Field (xAxis)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'xAxis' field " + _err.Error())
		}

		// Simple Field (yAxis)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'yAxis' field " + _err.Error())
		}

		// Simple Field (brightness)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'brightness' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (validityXy)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'validityXy' field " + _err.Error())
		}

		// Simple Field (validityBrightness)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'validityBrightness' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Converter_Status: // Struct

		// Simple Field (converterModeAccordingToTheDaliConverterStateMachine)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'converterModeAccordingToTheDaliConverterStateMachine' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(2, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (hardwiredSwitchIsActive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'hardwiredSwitchIsActive' field " + _err.Error())
		}

		// Simple Field (hardwiredInhibitIsActive)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'hardwiredInhibitIsActive' field " + _err.Error())
		}

		// Simple Field (functionTestPending)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'functionTestPending' field " + _err.Error())
		}

		// Simple Field (durationTestPending)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'durationTestPending' field " + _err.Error())
		}

		// Simple Field (partialDurationTestPending)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'partialDurationTestPending' field " + _err.Error())
		}

		// Simple Field (converterFailure)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'converterFailure' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Converter_Test_Result: // Struct

		// Simple Field (ltrf)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'ltrf' field " + _err.Error())
		}

		// Simple Field (ltrd)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'ltrd' field " + _err.Error())
		}

		// Simple Field (ltrp)
		if _err := io.WriteUint8(4, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'ltrp' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (sf)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'sf' field " + _err.Error())
		}

		// Simple Field (sd)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'sd' field " + _err.Error())
		}

		// Simple Field (sp)
		if _err := io.WriteUint8(2, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'sp' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(2, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (ldtr)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'ldtr' field " + _err.Error())
		}

		// Simple Field (lpdtr)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'lpdtr' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Battery_Info: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(5, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (batteryFailure)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'batteryFailure' field " + _err.Error())
		}

		// Simple Field (batteryDurationFailure)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'batteryDurationFailure' field " + _err.Error())
		}

		// Simple Field (batteryFullyCharged)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'batteryFullyCharged' field " + _err.Error())
		}

		// Simple Field (batteryChargeLevel)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'batteryChargeLevel' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Brightness_Colour_Temperature_Transition: // Struct

		// Simple Field (ms)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'ms' field " + _err.Error())
		}

		// Simple Field (temperatureK)
		if _err := io.WriteUint16(16, value.GetUint16()); _err != nil {
			return errors.New("Error serializing 'temperatureK' field " + _err.Error())
		}

		// Simple Field (percent)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'percent' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(5, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (validityOfTheTimePeriod)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'validityOfTheTimePeriod' field " + _err.Error())
		}

		// Simple Field (validityOfTheAbsoluteColourTemperature)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'validityOfTheAbsoluteColourTemperature' field " + _err.Error())
		}

		// Simple Field (validityOfTheAbsoluteBrightness)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'validityOfTheAbsoluteBrightness' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Brightness_Colour_Temperature_Control: // Struct

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (cct)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cct' field " + _err.Error())
		}

		// Simple Field (stepCodeColourTemperature)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepCodeColourTemperature' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (cb)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cb' field " + _err.Error())
		}

		// Simple Field (stepCodeBrightness)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepCodeBrightness' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(6, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (cctAndStepCodeColourValidity)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cctAndStepCodeColourValidity' field " + _err.Error())
		}

		// Simple Field (cbAndStepCodeBrightnessValidity)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cbAndStepCodeBrightnessValidity' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Colour_RGBW: // Struct

		// Simple Field (colourLevelRed)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'colourLevelRed' field " + _err.Error())
		}

		// Simple Field (colourLevelGreen)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'colourLevelGreen' field " + _err.Error())
		}

		// Simple Field (colourLevelBlue)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'colourLevelBlue' field " + _err.Error())
		}

		// Simple Field (colourLevelWhite)
		if _err := io.WriteUint8(8, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'colourLevelWhite' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(8, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (mr)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'mr' field " + _err.Error())
		}

		// Simple Field (mg)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'mg' field " + _err.Error())
		}

		// Simple Field (mb)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'mb' field " + _err.Error())
		}

		// Simple Field (mw)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'mw' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Relative_Control_RGBW: // Struct

		// Simple Field (maskcw)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskcw' field " + _err.Error())
		}

		// Simple Field (maskcb)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskcb' field " + _err.Error())
		}

		// Simple Field (maskcg)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskcg' field " + _err.Error())
		}

		// Simple Field (maskcr)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'maskcr' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (cw)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cw' field " + _err.Error())
		}

		// Simple Field (stepCodeColourWhite)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepCodeColourWhite' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (cb)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cb' field " + _err.Error())
		}

		// Simple Field (stepCodeColourBlue)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepCodeColourBlue' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (cg)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cg' field " + _err.Error())
		}

		// Simple Field (stepCodeColourGreen)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepCodeColourGreen' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (cr)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cr' field " + _err.Error())
		}

		// Simple Field (stepCodeColourRed)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepCodeColourRed' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_Relative_Control_RGB: // Struct

		// Simple Field (cb)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cb' field " + _err.Error())
		}

		// Simple Field (stepCodeColourBlue)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepCodeColourBlue' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (cg)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cg' field " + _err.Error())
		}

		// Simple Field (stepCodeColourGreen)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepCodeColourGreen' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}

		// Simple Field (cr)
		if _err := io.WriteBit(value.GetBool()); _err != nil {
			return errors.New("Error serializing 'cr' field " + _err.Error())
		}

		// Simple Field (stepCodeColourRed)
		if _err := io.WriteUint8(3, value.GetUint8()); _err != nil {
			return errors.New("Error serializing 'stepCodeColourRed' field " + _err.Error())
		}

		// Reserved Field (Just skip the bytes)
		if _err := io.WriteUint8(4, uint8(0x00)); _err != nil {
			return errors.New("Error serializing reserved field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_GeographicalLocation: // Struct

		// Simple Field (longitude)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'longitude' field " + _err.Error())
		}

		// Simple Field (latitude)
		if _err := io.WriteFloat32(32, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'latitude' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TempRoomSetpSetF16_4: // Struct

		// Simple Field (roomTemperatureSetpointComfort)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'roomTemperatureSetpointComfort' field " + _err.Error())
		}

		// Simple Field (roomTemperatureSetpointStandby)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'roomTemperatureSetpointStandby' field " + _err.Error())
		}

		// Simple Field (roomTemperatureSetpointEconomy)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'roomTemperatureSetpointEconomy' field " + _err.Error())
		}

		// Simple Field (roomTemperatureSetpointBuildingProtection)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'roomTemperatureSetpointBuildingProtection' field " + _err.Error())
		}
	case datapointType == KnxDatapointType_DPT_TempRoomSetpSetShiftF16_4: // Struct

		// Simple Field (roomTemperatureSetpointShiftComfort)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'roomTemperatureSetpointShiftComfort' field " + _err.Error())
		}

		// Simple Field (roomTemperatureSetpointShiftStandby)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'roomTemperatureSetpointShiftStandby' field " + _err.Error())
		}

		// Simple Field (roomTemperatureSetpointShiftEconomy)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'roomTemperatureSetpointShiftEconomy' field " + _err.Error())
		}

		// Simple Field (roomTemperatureSetpointShiftBuildingProtection)
		if _err := io.WriteFloat32(16, value.GetFloat32()); _err != nil {
			return errors.New("Error serializing 'roomTemperatureSetpointShiftBuildingProtection' field " + _err.Error())
		}
	default:

		return errors.New("unsupported type")
	}
	return nil
}
