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
type LDataFrame struct {
    FrameType bool
    NotRepeated bool
    Priority CEMIPriority
    AcknowledgeRequested bool
    ErrorFlag bool
    Child ILDataFrameChild
    ILDataFrame
    ILDataFrameParent
}

// The corresponding interface
type ILDataFrame interface {
    NotAckFrame() bool
    Polling() bool
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type ILDataFrameParent interface {
    SerializeParent(io utils.WriteBuffer, child ILDataFrame, serializeChildFunction func() error) error
    GetTypeName() string
}

type ILDataFrameChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *LDataFrame, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool)
    GetTypeName() string
    ILDataFrame
}

func NewLDataFrame(frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *LDataFrame {
    return &LDataFrame{FrameType: frameType, NotRepeated: notRepeated, Priority: priority, AcknowledgeRequested: acknowledgeRequested, ErrorFlag: errorFlag}
}

func CastLDataFrame(structType interface{}) *LDataFrame {
    castFunc := func(typ interface{}) *LDataFrame {
        if casted, ok := typ.(LDataFrame); ok {
            return &casted
        }
        if casted, ok := typ.(*LDataFrame); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *LDataFrame) GetTypeName() string {
    return "LDataFrame"
}

func (m *LDataFrame) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (frameType)
    lengthInBits += 1

    // Discriminator Field (polling)
    lengthInBits += 1

    // Simple field (notRepeated)
    lengthInBits += 1

    // Discriminator Field (notAckFrame)
    lengthInBits += 1

    // Enum Field (priority)
    lengthInBits += 2

    // Simple field (acknowledgeRequested)
    lengthInBits += 1

    // Simple field (errorFlag)
    lengthInBits += 1

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *LDataFrame) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func LDataFrameParse(io *utils.ReadBuffer) (*LDataFrame, error) {

    // Simple Field (frameType)
    frameType, _frameTypeErr := io.ReadBit()
    if _frameTypeErr != nil {
        return nil, errors.New("Error parsing 'frameType' field " + _frameTypeErr.Error())
    }

    // Discriminator Field (polling) (Used as input to a switch field)
    polling, _pollingErr := io.ReadBit()
    if _pollingErr != nil {
        return nil, errors.New("Error parsing 'polling' field " + _pollingErr.Error())
    }

    // Simple Field (notRepeated)
    notRepeated, _notRepeatedErr := io.ReadBit()
    if _notRepeatedErr != nil {
        return nil, errors.New("Error parsing 'notRepeated' field " + _notRepeatedErr.Error())
    }

    // Discriminator Field (notAckFrame) (Used as input to a switch field)
    notAckFrame, _notAckFrameErr := io.ReadBit()
    if _notAckFrameErr != nil {
        return nil, errors.New("Error parsing 'notAckFrame' field " + _notAckFrameErr.Error())
    }

    // Enum field (priority)
    priority, _priorityErr := CEMIPriorityParse(io)
    if _priorityErr != nil {
        return nil, errors.New("Error parsing 'priority' field " + _priorityErr.Error())
    }

    // Simple Field (acknowledgeRequested)
    acknowledgeRequested, _acknowledgeRequestedErr := io.ReadBit()
    if _acknowledgeRequestedErr != nil {
        return nil, errors.New("Error parsing 'acknowledgeRequested' field " + _acknowledgeRequestedErr.Error())
    }

    // Simple Field (errorFlag)
    errorFlag, _errorFlagErr := io.ReadBit()
    if _errorFlagErr != nil {
        return nil, errors.New("Error parsing 'errorFlag' field " + _errorFlagErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *LDataFrame
    var typeSwitchError error
    switch {
    case notAckFrame == true && polling == false:
        _parent, typeSwitchError = LDataExtendedParse(io)
    case notAckFrame == true && polling == true:
        _parent, typeSwitchError = LPollDataParse(io)
    case notAckFrame == false:
        _parent, typeSwitchError = LDataFrameACKParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent, frameType, notRepeated, priority, acknowledgeRequested, errorFlag)
    return _parent, nil
}

func (m *LDataFrame) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *LDataFrame) SerializeParent(io utils.WriteBuffer, child ILDataFrame, serializeChildFunction func() error) error {

    // Simple Field (frameType)
    frameType := bool(m.FrameType)
    _frameTypeErr := io.WriteBit((frameType))
    if _frameTypeErr != nil {
        return errors.New("Error serializing 'frameType' field " + _frameTypeErr.Error())
    }

    // Discriminator Field (polling) (Used as input to a switch field)
    polling := bool(child.Polling())
    _pollingErr := io.WriteBit((polling))
    if _pollingErr != nil {
        return errors.New("Error serializing 'polling' field " + _pollingErr.Error())
    }

    // Simple Field (notRepeated)
    notRepeated := bool(m.NotRepeated)
    _notRepeatedErr := io.WriteBit((notRepeated))
    if _notRepeatedErr != nil {
        return errors.New("Error serializing 'notRepeated' field " + _notRepeatedErr.Error())
    }

    // Discriminator Field (notAckFrame) (Used as input to a switch field)
    notAckFrame := bool(child.NotAckFrame())
    _notAckFrameErr := io.WriteBit((notAckFrame))
    if _notAckFrameErr != nil {
        return errors.New("Error serializing 'notAckFrame' field " + _notAckFrameErr.Error())
    }

    // Enum field (priority)
    priority := CastCEMIPriority(m.Priority)
    _priorityErr := priority.Serialize(io)
    if _priorityErr != nil {
        return errors.New("Error serializing 'priority' field " + _priorityErr.Error())
    }

    // Simple Field (acknowledgeRequested)
    acknowledgeRequested := bool(m.AcknowledgeRequested)
    _acknowledgeRequestedErr := io.WriteBit((acknowledgeRequested))
    if _acknowledgeRequestedErr != nil {
        return errors.New("Error serializing 'acknowledgeRequested' field " + _acknowledgeRequestedErr.Error())
    }

    // Simple Field (errorFlag)
    errorFlag := bool(m.ErrorFlag)
    _errorFlagErr := io.WriteBit((errorFlag))
    if _errorFlagErr != nil {
        return errors.New("Error serializing 'errorFlag' field " + _errorFlagErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *LDataFrame) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "frameType":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.FrameType = data
            case "notRepeated":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.NotRepeated = data
            case "priority":
                var data CEMIPriority
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Priority = data
            case "acknowledgeRequested":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.AcknowledgeRequested = data
            case "errorFlag":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ErrorFlag = data
            default:
                switch start.Attr[0].Value {
                    case "org.apache.plc4x.java.knxnetip.readwrite.LDataExtended":
                        var dt *LDataExtended
                        if m.Child != nil {
                            dt = m.Child.(*LDataExtended)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.LPollData":
                        var dt *LPollData
                        if m.Child != nil {
                            dt = m.Child.(*LPollData)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.LDataFrameACK":
                        var dt *LDataFrameACK
                        if m.Child != nil {
                            dt = m.Child.(*LDataFrameACK)
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

func (m *LDataFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := reflect.TypeOf(m.Child).String()
    className = "org.apache.plc4x.java.knxnetip.readwrite." + className[strings.LastIndex(className, ".") + 1:]
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.FrameType, xml.StartElement{Name: xml.Name{Local: "frameType"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.NotRepeated, xml.StartElement{Name: xml.Name{Local: "notRepeated"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Priority, xml.StartElement{Name: xml.Name{Local: "priority"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.AcknowledgeRequested, xml.StartElement{Name: xml.Name{Local: "acknowledgeRequested"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ErrorFlag, xml.StartElement{Name: xml.Name{Local: "errorFlag"}}); err != nil {
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

