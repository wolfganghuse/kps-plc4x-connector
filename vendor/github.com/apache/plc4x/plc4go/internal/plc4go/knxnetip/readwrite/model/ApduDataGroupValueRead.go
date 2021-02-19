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
	log "github.com/sirupsen/logrus"
	"io"
)

// The data-structure of this message
type ApduDataGroupValueRead struct {
	Parent *ApduData
	IApduDataGroupValueRead
}

// The corresponding interface
type IApduDataGroupValueRead interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataGroupValueRead) ApciType() uint8 {
	return 0x0
}

func (m *ApduDataGroupValueRead) InitializeParent(parent *ApduData) {
}

func NewApduDataGroupValueRead() *ApduData {
	child := &ApduDataGroupValueRead{
		Parent: NewApduData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataGroupValueRead(structType interface{}) *ApduDataGroupValueRead {
	castFunc := func(typ interface{}) *ApduDataGroupValueRead {
		if casted, ok := typ.(ApduDataGroupValueRead); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataGroupValueRead); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataGroupValueRead(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataGroupValueRead(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataGroupValueRead) GetTypeName() string {
	return "ApduDataGroupValueRead"
}

func (m *ApduDataGroupValueRead) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 6

	return lengthInBits
}

func (m *ApduDataGroupValueRead) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataGroupValueReadParse(io *utils.ReadBuffer) (*ApduData, error) {

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8(6)
		if _err != nil {
			return nil, errors.New("Error parsing 'reserved' field " + _err.Error())
		}
		if reserved != uint8(0x00) {
			log.WithFields(log.Fields{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Info("Got unexpected response.")
		}
	}

	// Create a partially initialized instance
	_child := &ApduDataGroupValueRead{
		Parent: &ApduData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataGroupValueRead) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Reserved Field (reserved)
		{
			_err := io.WriteUint8(6, uint8(0x00))
			if _err != nil {
				return errors.New("Error serializing 'reserved' field " + _err.Error())
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataGroupValueRead) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
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

func (m *ApduDataGroupValueRead) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}
