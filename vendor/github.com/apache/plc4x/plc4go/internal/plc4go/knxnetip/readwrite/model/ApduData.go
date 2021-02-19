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
type ApduData struct {
	Child IApduDataChild
	IApduData
	IApduDataParent
}

// The corresponding interface
type IApduData interface {
	ApciType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

type IApduDataParent interface {
	SerializeParent(io utils.WriteBuffer, child IApduData, serializeChildFunction func() error) error
	GetTypeName() string
}

type IApduDataChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *ApduData)
	GetTypeName() string
	IApduData
}

func NewApduData() *ApduData {
	return &ApduData{}
}

func CastApduData(structType interface{}) *ApduData {
	castFunc := func(typ interface{}) *ApduData {
		if casted, ok := typ.(ApduData); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduData); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduData) GetTypeName() string {
	return "ApduData"
}

func (m *ApduData) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Discriminator Field (apciType)
	lengthInBits += 4

	// Length of sub-type elements will be added by sub-type...
	lengthInBits += m.Child.LengthInBits()

	return lengthInBits
}

func (m *ApduData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataParse(io *utils.ReadBuffer, dataLength uint8) (*ApduData, error) {

	// Discriminator Field (apciType) (Used as input to a switch field)
	apciType, _apciTypeErr := io.ReadUint8(4)
	if _apciTypeErr != nil {
		return nil, errors.New("Error parsing 'apciType' field " + _apciTypeErr.Error())
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ApduData
	var typeSwitchError error
	switch {
	case apciType == 0x0:
		_parent, typeSwitchError = ApduDataGroupValueReadParse(io)
	case apciType == 0x1:
		_parent, typeSwitchError = ApduDataGroupValueResponseParse(io, dataLength)
	case apciType == 0x2:
		_parent, typeSwitchError = ApduDataGroupValueWriteParse(io, dataLength)
	case apciType == 0x3:
		_parent, typeSwitchError = ApduDataIndividualAddressWriteParse(io)
	case apciType == 0x4:
		_parent, typeSwitchError = ApduDataIndividualAddressReadParse(io)
	case apciType == 0x5:
		_parent, typeSwitchError = ApduDataIndividualAddressResponseParse(io)
	case apciType == 0x6:
		_parent, typeSwitchError = ApduDataAdcReadParse(io)
	case apciType == 0x7:
		_parent, typeSwitchError = ApduDataAdcResponseParse(io)
	case apciType == 0x8:
		_parent, typeSwitchError = ApduDataMemoryReadParse(io)
	case apciType == 0x9:
		_parent, typeSwitchError = ApduDataMemoryResponseParse(io)
	case apciType == 0xA:
		_parent, typeSwitchError = ApduDataMemoryWriteParse(io)
	case apciType == 0xB:
		_parent, typeSwitchError = ApduDataUserMessageParse(io)
	case apciType == 0xC:
		_parent, typeSwitchError = ApduDataDeviceDescriptorReadParse(io)
	case apciType == 0xD:
		_parent, typeSwitchError = ApduDataDeviceDescriptorResponseParse(io, dataLength)
	case apciType == 0xE:
		_parent, typeSwitchError = ApduDataRestartParse(io)
	case apciType == 0xF:
		_parent, typeSwitchError = ApduDataOtherParse(io, dataLength)
	}
	if typeSwitchError != nil {
		return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ApduData) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *ApduData) SerializeParent(io utils.WriteBuffer, child IApduData, serializeChildFunction func() error) error {

	// Discriminator Field (apciType) (Used as input to a switch field)
	apciType := uint8(child.ApciType())
	_apciTypeErr := io.WriteUint8(4, (apciType))
	if _apciTypeErr != nil {
		return errors.New("Error serializing 'apciType' field " + _apciTypeErr.Error())
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
	}

	return nil
}

func (m *ApduData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataGroupValueRead":
					var dt *ApduDataGroupValueRead
					if m.Child != nil {
						dt = m.Child.(*ApduDataGroupValueRead)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataGroupValueResponse":
					var dt *ApduDataGroupValueResponse
					if m.Child != nil {
						dt = m.Child.(*ApduDataGroupValueResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataGroupValueWrite":
					var dt *ApduDataGroupValueWrite
					if m.Child != nil {
						dt = m.Child.(*ApduDataGroupValueWrite)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataIndividualAddressWrite":
					var dt *ApduDataIndividualAddressWrite
					if m.Child != nil {
						dt = m.Child.(*ApduDataIndividualAddressWrite)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataIndividualAddressRead":
					var dt *ApduDataIndividualAddressRead
					if m.Child != nil {
						dt = m.Child.(*ApduDataIndividualAddressRead)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataIndividualAddressResponse":
					var dt *ApduDataIndividualAddressResponse
					if m.Child != nil {
						dt = m.Child.(*ApduDataIndividualAddressResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataAdcRead":
					var dt *ApduDataAdcRead
					if m.Child != nil {
						dt = m.Child.(*ApduDataAdcRead)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataAdcResponse":
					var dt *ApduDataAdcResponse
					if m.Child != nil {
						dt = m.Child.(*ApduDataAdcResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataMemoryRead":
					var dt *ApduDataMemoryRead
					if m.Child != nil {
						dt = m.Child.(*ApduDataMemoryRead)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataMemoryResponse":
					var dt *ApduDataMemoryResponse
					if m.Child != nil {
						dt = m.Child.(*ApduDataMemoryResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataMemoryWrite":
					var dt *ApduDataMemoryWrite
					if m.Child != nil {
						dt = m.Child.(*ApduDataMemoryWrite)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataUserMessage":
					var dt *ApduDataUserMessage
					if m.Child != nil {
						dt = m.Child.(*ApduDataUserMessage)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataDeviceDescriptorRead":
					var dt *ApduDataDeviceDescriptorRead
					if m.Child != nil {
						dt = m.Child.(*ApduDataDeviceDescriptorRead)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataDeviceDescriptorResponse":
					var dt *ApduDataDeviceDescriptorResponse
					if m.Child != nil {
						dt = m.Child.(*ApduDataDeviceDescriptorResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataRestart":
					var dt *ApduDataRestart
					if m.Child != nil {
						dt = m.Child.(*ApduDataRestart)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataOther":
					var dt *ApduDataOther
					if m.Child != nil {
						dt = m.Child.(*ApduDataOther)
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

func (m *ApduData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.knxnetip.readwrite." + className[strings.LastIndex(className, ".")+1:]
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
