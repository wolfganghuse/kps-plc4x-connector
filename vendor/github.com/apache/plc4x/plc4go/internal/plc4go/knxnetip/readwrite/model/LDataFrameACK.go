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
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
)

// The data-structure of this message
type LDataFrameACK struct {
    Parent *LDataFrame
    ILDataFrameACK
}

// The corresponding interface
type ILDataFrameACK interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LDataFrameACK) NotAckFrame() bool {
    return false
}

func (m *LDataFrameACK) Polling() bool {
    return false
}


func (m *LDataFrameACK) InitializeParent(parent *LDataFrame, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) {
    m.Parent.FrameType = frameType
    m.Parent.NotRepeated = notRepeated
    m.Parent.Priority = priority
    m.Parent.AcknowledgeRequested = acknowledgeRequested
    m.Parent.ErrorFlag = errorFlag
}

func NewLDataFrameACK(frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *LDataFrame {
    child := &LDataFrameACK{
        Parent: NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastLDataFrameACK(structType interface{}) *LDataFrameACK {
    castFunc := func(typ interface{}) *LDataFrameACK {
        if casted, ok := typ.(LDataFrameACK); ok {
            return &casted
        }
        if casted, ok := typ.(*LDataFrameACK); ok {
            return casted
        }
        if casted, ok := typ.(LDataFrame); ok {
            return CastLDataFrameACK(casted.Child)
        }
        if casted, ok := typ.(*LDataFrame); ok {
            return CastLDataFrameACK(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *LDataFrameACK) GetTypeName() string {
    return "LDataFrameACK"
}

func (m *LDataFrameACK) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *LDataFrameACK) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func LDataFrameACKParse(io *utils.ReadBuffer) (*LDataFrame, error) {

    // Create a partially initialized instance
    _child := &LDataFrameACK{
        Parent: &LDataFrame{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *LDataFrameACK) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *LDataFrameACK) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *LDataFrameACK) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

