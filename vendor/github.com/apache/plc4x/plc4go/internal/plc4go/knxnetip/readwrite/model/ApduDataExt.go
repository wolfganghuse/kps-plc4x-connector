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
type ApduDataExt struct {
    Child IApduDataExtChild
    IApduDataExt
    IApduDataExtParent
}

// The corresponding interface
type IApduDataExt interface {
    ExtApciType() uint8
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type IApduDataExtParent interface {
    SerializeParent(io utils.WriteBuffer, child IApduDataExt, serializeChildFunction func() error) error
    GetTypeName() string
}

type IApduDataExtChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *ApduDataExt)
    GetTypeName() string
    IApduDataExt
}

func NewApduDataExt() *ApduDataExt {
    return &ApduDataExt{}
}

func CastApduDataExt(structType interface{}) *ApduDataExt {
    castFunc := func(typ interface{}) *ApduDataExt {
        if casted, ok := typ.(ApduDataExt); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataExt); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataExt) GetTypeName() string {
    return "ApduDataExt"
}

func (m *ApduDataExt) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Discriminator Field (extApciType)
    lengthInBits += 6

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *ApduDataExt) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataExtParse(io *utils.ReadBuffer, length uint8) (*ApduDataExt, error) {

    // Discriminator Field (extApciType) (Used as input to a switch field)
    extApciType, _extApciTypeErr := io.ReadUint8(6)
    if _extApciTypeErr != nil {
        return nil, errors.New("Error parsing 'extApciType' field " + _extApciTypeErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *ApduDataExt
    var typeSwitchError error
    switch {
    case extApciType == 0x00:
        _parent, typeSwitchError = ApduDataExtOpenRoutingTableRequestParse(io)
    case extApciType == 0x01:
        _parent, typeSwitchError = ApduDataExtReadRoutingTableRequestParse(io)
    case extApciType == 0x02:
        _parent, typeSwitchError = ApduDataExtReadRoutingTableResponseParse(io)
    case extApciType == 0x03:
        _parent, typeSwitchError = ApduDataExtWriteRoutingTableRequestParse(io)
    case extApciType == 0x08:
        _parent, typeSwitchError = ApduDataExtReadRouterMemoryRequestParse(io)
    case extApciType == 0x09:
        _parent, typeSwitchError = ApduDataExtReadRouterMemoryResponseParse(io)
    case extApciType == 0x0A:
        _parent, typeSwitchError = ApduDataExtWriteRouterMemoryRequestParse(io)
    case extApciType == 0x0D:
        _parent, typeSwitchError = ApduDataExtReadRouterStatusRequestParse(io)
    case extApciType == 0x0E:
        _parent, typeSwitchError = ApduDataExtReadRouterStatusResponseParse(io)
    case extApciType == 0x0F:
        _parent, typeSwitchError = ApduDataExtWriteRouterStatusRequestParse(io)
    case extApciType == 0x10:
        _parent, typeSwitchError = ApduDataExtMemoryBitWriteParse(io)
    case extApciType == 0x11:
        _parent, typeSwitchError = ApduDataExtAuthorizeRequestParse(io)
    case extApciType == 0x12:
        _parent, typeSwitchError = ApduDataExtAuthorizeResponseParse(io)
    case extApciType == 0x13:
        _parent, typeSwitchError = ApduDataExtKeyWriteParse(io)
    case extApciType == 0x14:
        _parent, typeSwitchError = ApduDataExtKeyResponseParse(io)
    case extApciType == 0x15:
        _parent, typeSwitchError = ApduDataExtPropertyValueReadParse(io)
    case extApciType == 0x16:
        _parent, typeSwitchError = ApduDataExtPropertyValueResponseParse(io, length)
    case extApciType == 0x17:
        _parent, typeSwitchError = ApduDataExtPropertyValueWriteParse(io)
    case extApciType == 0x18:
        _parent, typeSwitchError = ApduDataExtPropertyDescriptionReadParse(io)
    case extApciType == 0x19:
        _parent, typeSwitchError = ApduDataExtPropertyDescriptionResponseParse(io)
    case extApciType == 0x1A:
        _parent, typeSwitchError = ApduDataExtNetworkParameterReadParse(io)
    case extApciType == 0x1B:
        _parent, typeSwitchError = ApduDataExtNetworkParameterResponseParse(io)
    case extApciType == 0x1C:
        _parent, typeSwitchError = ApduDataExtIndividualAddressSerialNumberReadParse(io)
    case extApciType == 0x1D:
        _parent, typeSwitchError = ApduDataExtIndividualAddressSerialNumberResponseParse(io)
    case extApciType == 0x1E:
        _parent, typeSwitchError = ApduDataExtIndividualAddressSerialNumberWriteParse(io)
    case extApciType == 0x20:
        _parent, typeSwitchError = ApduDataExtDomainAddressWriteParse(io)
    case extApciType == 0x21:
        _parent, typeSwitchError = ApduDataExtDomainAddressReadParse(io)
    case extApciType == 0x22:
        _parent, typeSwitchError = ApduDataExtDomainAddressResponseParse(io)
    case extApciType == 0x23:
        _parent, typeSwitchError = ApduDataExtDomainAddressSelectiveReadParse(io)
    case extApciType == 0x24:
        _parent, typeSwitchError = ApduDataExtNetworkParameterWriteParse(io)
    case extApciType == 0x25:
        _parent, typeSwitchError = ApduDataExtLinkReadParse(io)
    case extApciType == 0x26:
        _parent, typeSwitchError = ApduDataExtLinkResponseParse(io)
    case extApciType == 0x27:
        _parent, typeSwitchError = ApduDataExtLinkWriteParse(io)
    case extApciType == 0x28:
        _parent, typeSwitchError = ApduDataExtGroupPropertyValueReadParse(io)
    case extApciType == 0x29:
        _parent, typeSwitchError = ApduDataExtGroupPropertyValueResponseParse(io)
    case extApciType == 0x2A:
        _parent, typeSwitchError = ApduDataExtGroupPropertyValueWriteParse(io)
    case extApciType == 0x2B:
        _parent, typeSwitchError = ApduDataExtGroupPropertyValueInfoReportParse(io)
    case extApciType == 0x2C:
        _parent, typeSwitchError = ApduDataExtDomainAddressSerialNumberReadParse(io)
    case extApciType == 0x2D:
        _parent, typeSwitchError = ApduDataExtDomainAddressSerialNumberResponseParse(io)
    case extApciType == 0x2E:
        _parent, typeSwitchError = ApduDataExtDomainAddressSerialNumberWriteParse(io)
    case extApciType == 0x30:
        _parent, typeSwitchError = ApduDataExtFileStreamInfoReportParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent)
    return _parent, nil
}

func (m *ApduDataExt) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *ApduDataExt) SerializeParent(io utils.WriteBuffer, child IApduDataExt, serializeChildFunction func() error) error {

    // Discriminator Field (extApciType) (Used as input to a switch field)
    extApciType := uint8(child.ExtApciType())
    _extApciTypeErr := io.WriteUint8(6, (extApciType))
    if _extApciTypeErr != nil {
        return errors.New("Error serializing 'extApciType' field " + _extApciTypeErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *ApduDataExt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtOpenRoutingTableRequest":
                        var dt *ApduDataExtOpenRoutingTableRequest
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtOpenRoutingTableRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtReadRoutingTableRequest":
                        var dt *ApduDataExtReadRoutingTableRequest
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtReadRoutingTableRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtReadRoutingTableResponse":
                        var dt *ApduDataExtReadRoutingTableResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtReadRoutingTableResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtWriteRoutingTableRequest":
                        var dt *ApduDataExtWriteRoutingTableRequest
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtWriteRoutingTableRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtReadRouterMemoryRequest":
                        var dt *ApduDataExtReadRouterMemoryRequest
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtReadRouterMemoryRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtReadRouterMemoryResponse":
                        var dt *ApduDataExtReadRouterMemoryResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtReadRouterMemoryResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtWriteRouterMemoryRequest":
                        var dt *ApduDataExtWriteRouterMemoryRequest
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtWriteRouterMemoryRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtReadRouterStatusRequest":
                        var dt *ApduDataExtReadRouterStatusRequest
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtReadRouterStatusRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtReadRouterStatusResponse":
                        var dt *ApduDataExtReadRouterStatusResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtReadRouterStatusResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtWriteRouterStatusRequest":
                        var dt *ApduDataExtWriteRouterStatusRequest
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtWriteRouterStatusRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtMemoryBitWrite":
                        var dt *ApduDataExtMemoryBitWrite
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtMemoryBitWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtAuthorizeRequest":
                        var dt *ApduDataExtAuthorizeRequest
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtAuthorizeRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtAuthorizeResponse":
                        var dt *ApduDataExtAuthorizeResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtAuthorizeResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtKeyWrite":
                        var dt *ApduDataExtKeyWrite
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtKeyWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtKeyResponse":
                        var dt *ApduDataExtKeyResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtKeyResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtPropertyValueRead":
                        var dt *ApduDataExtPropertyValueRead
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtPropertyValueRead)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtPropertyValueResponse":
                        var dt *ApduDataExtPropertyValueResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtPropertyValueResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtPropertyValueWrite":
                        var dt *ApduDataExtPropertyValueWrite
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtPropertyValueWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtPropertyDescriptionRead":
                        var dt *ApduDataExtPropertyDescriptionRead
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtPropertyDescriptionRead)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtPropertyDescriptionResponse":
                        var dt *ApduDataExtPropertyDescriptionResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtPropertyDescriptionResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtNetworkParameterRead":
                        var dt *ApduDataExtNetworkParameterRead
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtNetworkParameterRead)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtNetworkParameterResponse":
                        var dt *ApduDataExtNetworkParameterResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtNetworkParameterResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtIndividualAddressSerialNumberRead":
                        var dt *ApduDataExtIndividualAddressSerialNumberRead
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtIndividualAddressSerialNumberRead)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtIndividualAddressSerialNumberResponse":
                        var dt *ApduDataExtIndividualAddressSerialNumberResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtIndividualAddressSerialNumberResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtIndividualAddressSerialNumberWrite":
                        var dt *ApduDataExtIndividualAddressSerialNumberWrite
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtIndividualAddressSerialNumberWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtDomainAddressWrite":
                        var dt *ApduDataExtDomainAddressWrite
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtDomainAddressWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtDomainAddressRead":
                        var dt *ApduDataExtDomainAddressRead
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtDomainAddressRead)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtDomainAddressResponse":
                        var dt *ApduDataExtDomainAddressResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtDomainAddressResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtDomainAddressSelectiveRead":
                        var dt *ApduDataExtDomainAddressSelectiveRead
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtDomainAddressSelectiveRead)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtNetworkParameterWrite":
                        var dt *ApduDataExtNetworkParameterWrite
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtNetworkParameterWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtLinkRead":
                        var dt *ApduDataExtLinkRead
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtLinkRead)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtLinkResponse":
                        var dt *ApduDataExtLinkResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtLinkResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtLinkWrite":
                        var dt *ApduDataExtLinkWrite
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtLinkWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtGroupPropertyValueRead":
                        var dt *ApduDataExtGroupPropertyValueRead
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtGroupPropertyValueRead)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtGroupPropertyValueResponse":
                        var dt *ApduDataExtGroupPropertyValueResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtGroupPropertyValueResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtGroupPropertyValueWrite":
                        var dt *ApduDataExtGroupPropertyValueWrite
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtGroupPropertyValueWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtGroupPropertyValueInfoReport":
                        var dt *ApduDataExtGroupPropertyValueInfoReport
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtGroupPropertyValueInfoReport)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtDomainAddressSerialNumberRead":
                        var dt *ApduDataExtDomainAddressSerialNumberRead
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtDomainAddressSerialNumberRead)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtDomainAddressSerialNumberResponse":
                        var dt *ApduDataExtDomainAddressSerialNumberResponse
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtDomainAddressSerialNumberResponse)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtDomainAddressSerialNumberWrite":
                        var dt *ApduDataExtDomainAddressSerialNumberWrite
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtDomainAddressSerialNumberWrite)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.knxnetip.readwrite.ApduDataExtFileStreamInfoReport":
                        var dt *ApduDataExtFileStreamInfoReport
                        if m.Child != nil {
                            dt = m.Child.(*ApduDataExtFileStreamInfoReport)
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

func (m *ApduDataExt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := reflect.TypeOf(m.Child).String()
    className = "org.apache.plc4x.java.knxnetip.readwrite." + className[strings.LastIndex(className, ".") + 1:]
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

