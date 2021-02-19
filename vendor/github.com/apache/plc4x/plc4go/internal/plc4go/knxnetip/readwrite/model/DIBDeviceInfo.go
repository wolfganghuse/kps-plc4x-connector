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
    "encoding/base64"
    "encoding/xml"
    "errors"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
)

// The data-structure of this message
type DIBDeviceInfo struct {
    DescriptionType uint8
    KnxMedium KnxMedium
    DeviceStatus *DeviceStatus
    KnxAddress *KnxAddress
    ProjectInstallationIdentifier *ProjectInstallationIdentifier
    KnxNetIpDeviceSerialNumber []int8
    KnxNetIpDeviceMulticastAddress *IPAddress
    KnxNetIpDeviceMacAddress *MACAddress
    DeviceFriendlyName []int8
    IDIBDeviceInfo
}

// The corresponding interface
type IDIBDeviceInfo interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

func NewDIBDeviceInfo(descriptionType uint8, knxMedium KnxMedium, deviceStatus *DeviceStatus, knxAddress *KnxAddress, projectInstallationIdentifier *ProjectInstallationIdentifier, knxNetIpDeviceSerialNumber []int8, knxNetIpDeviceMulticastAddress *IPAddress, knxNetIpDeviceMacAddress *MACAddress, deviceFriendlyName []int8) *DIBDeviceInfo {
    return &DIBDeviceInfo{DescriptionType: descriptionType, KnxMedium: knxMedium, DeviceStatus: deviceStatus, KnxAddress: knxAddress, ProjectInstallationIdentifier: projectInstallationIdentifier, KnxNetIpDeviceSerialNumber: knxNetIpDeviceSerialNumber, KnxNetIpDeviceMulticastAddress: knxNetIpDeviceMulticastAddress, KnxNetIpDeviceMacAddress: knxNetIpDeviceMacAddress, DeviceFriendlyName: deviceFriendlyName}
}

func CastDIBDeviceInfo(structType interface{}) *DIBDeviceInfo {
    castFunc := func(typ interface{}) *DIBDeviceInfo {
        if casted, ok := typ.(DIBDeviceInfo); ok {
            return &casted
        }
        if casted, ok := typ.(*DIBDeviceInfo); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *DIBDeviceInfo) GetTypeName() string {
    return "DIBDeviceInfo"
}

func (m *DIBDeviceInfo) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Implicit Field (structureLength)
    lengthInBits += 8

    // Simple field (descriptionType)
    lengthInBits += 8

    // Simple field (knxMedium)
    lengthInBits += 8

    // Simple field (deviceStatus)
    lengthInBits += m.DeviceStatus.LengthInBits()

    // Simple field (knxAddress)
    lengthInBits += m.KnxAddress.LengthInBits()

    // Simple field (projectInstallationIdentifier)
    lengthInBits += m.ProjectInstallationIdentifier.LengthInBits()

    // Array field
    if len(m.KnxNetIpDeviceSerialNumber) > 0 {
        lengthInBits += 8 * uint16(len(m.KnxNetIpDeviceSerialNumber))
    }

    // Simple field (knxNetIpDeviceMulticastAddress)
    lengthInBits += m.KnxNetIpDeviceMulticastAddress.LengthInBits()

    // Simple field (knxNetIpDeviceMacAddress)
    lengthInBits += m.KnxNetIpDeviceMacAddress.LengthInBits()

    // Array field
    if len(m.DeviceFriendlyName) > 0 {
        lengthInBits += 8 * uint16(len(m.DeviceFriendlyName))
    }

    return lengthInBits
}

func (m *DIBDeviceInfo) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func DIBDeviceInfoParse(io *utils.ReadBuffer) (*DIBDeviceInfo, error) {

    // Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    _, _structureLengthErr := io.ReadUint8(8)
    if _structureLengthErr != nil {
        return nil, errors.New("Error parsing 'structureLength' field " + _structureLengthErr.Error())
    }

    // Simple Field (descriptionType)
    descriptionType, _descriptionTypeErr := io.ReadUint8(8)
    if _descriptionTypeErr != nil {
        return nil, errors.New("Error parsing 'descriptionType' field " + _descriptionTypeErr.Error())
    }

    // Simple Field (knxMedium)
    knxMedium, _knxMediumErr := KnxMediumParse(io)
    if _knxMediumErr != nil {
        return nil, errors.New("Error parsing 'knxMedium' field " + _knxMediumErr.Error())
    }

    // Simple Field (deviceStatus)
    deviceStatus, _deviceStatusErr := DeviceStatusParse(io)
    if _deviceStatusErr != nil {
        return nil, errors.New("Error parsing 'deviceStatus' field " + _deviceStatusErr.Error())
    }

    // Simple Field (knxAddress)
    knxAddress, _knxAddressErr := KnxAddressParse(io)
    if _knxAddressErr != nil {
        return nil, errors.New("Error parsing 'knxAddress' field " + _knxAddressErr.Error())
    }

    // Simple Field (projectInstallationIdentifier)
    projectInstallationIdentifier, _projectInstallationIdentifierErr := ProjectInstallationIdentifierParse(io)
    if _projectInstallationIdentifierErr != nil {
        return nil, errors.New("Error parsing 'projectInstallationIdentifier' field " + _projectInstallationIdentifierErr.Error())
    }

    // Array field (knxNetIpDeviceSerialNumber)
    // Count array
    knxNetIpDeviceSerialNumber := make([]int8, uint16(6))
    for curItem := uint16(0); curItem < uint16(uint16(6)); curItem++ {
        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'knxNetIpDeviceSerialNumber' field " + _err.Error())
        }
        knxNetIpDeviceSerialNumber[curItem] = _item
    }

    // Simple Field (knxNetIpDeviceMulticastAddress)
    knxNetIpDeviceMulticastAddress, _knxNetIpDeviceMulticastAddressErr := IPAddressParse(io)
    if _knxNetIpDeviceMulticastAddressErr != nil {
        return nil, errors.New("Error parsing 'knxNetIpDeviceMulticastAddress' field " + _knxNetIpDeviceMulticastAddressErr.Error())
    }

    // Simple Field (knxNetIpDeviceMacAddress)
    knxNetIpDeviceMacAddress, _knxNetIpDeviceMacAddressErr := MACAddressParse(io)
    if _knxNetIpDeviceMacAddressErr != nil {
        return nil, errors.New("Error parsing 'knxNetIpDeviceMacAddress' field " + _knxNetIpDeviceMacAddressErr.Error())
    }

    // Array field (deviceFriendlyName)
    // Count array
    deviceFriendlyName := make([]int8, uint16(30))
    for curItem := uint16(0); curItem < uint16(uint16(30)); curItem++ {
        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'deviceFriendlyName' field " + _err.Error())
        }
        deviceFriendlyName[curItem] = _item
    }

    // Create the instance
    return NewDIBDeviceInfo(descriptionType, knxMedium, deviceStatus, knxAddress, projectInstallationIdentifier, knxNetIpDeviceSerialNumber, knxNetIpDeviceMulticastAddress, knxNetIpDeviceMacAddress, deviceFriendlyName), nil
}

func (m *DIBDeviceInfo) Serialize(io utils.WriteBuffer) error {

    // Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    structureLength := uint8(uint8(m.LengthInBytes()))
    _structureLengthErr := io.WriteUint8(8, (structureLength))
    if _structureLengthErr != nil {
        return errors.New("Error serializing 'structureLength' field " + _structureLengthErr.Error())
    }

    // Simple Field (descriptionType)
    descriptionType := uint8(m.DescriptionType)
    _descriptionTypeErr := io.WriteUint8(8, (descriptionType))
    if _descriptionTypeErr != nil {
        return errors.New("Error serializing 'descriptionType' field " + _descriptionTypeErr.Error())
    }

    // Simple Field (knxMedium)
    _knxMediumErr := m.KnxMedium.Serialize(io)
    if _knxMediumErr != nil {
        return errors.New("Error serializing 'knxMedium' field " + _knxMediumErr.Error())
    }

    // Simple Field (deviceStatus)
    _deviceStatusErr := m.DeviceStatus.Serialize(io)
    if _deviceStatusErr != nil {
        return errors.New("Error serializing 'deviceStatus' field " + _deviceStatusErr.Error())
    }

    // Simple Field (knxAddress)
    _knxAddressErr := m.KnxAddress.Serialize(io)
    if _knxAddressErr != nil {
        return errors.New("Error serializing 'knxAddress' field " + _knxAddressErr.Error())
    }

    // Simple Field (projectInstallationIdentifier)
    _projectInstallationIdentifierErr := m.ProjectInstallationIdentifier.Serialize(io)
    if _projectInstallationIdentifierErr != nil {
        return errors.New("Error serializing 'projectInstallationIdentifier' field " + _projectInstallationIdentifierErr.Error())
    }

    // Array Field (knxNetIpDeviceSerialNumber)
    if m.KnxNetIpDeviceSerialNumber != nil {
        for _, _element := range m.KnxNetIpDeviceSerialNumber {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'knxNetIpDeviceSerialNumber' field " + _elementErr.Error())
            }
        }
    }

    // Simple Field (knxNetIpDeviceMulticastAddress)
    _knxNetIpDeviceMulticastAddressErr := m.KnxNetIpDeviceMulticastAddress.Serialize(io)
    if _knxNetIpDeviceMulticastAddressErr != nil {
        return errors.New("Error serializing 'knxNetIpDeviceMulticastAddress' field " + _knxNetIpDeviceMulticastAddressErr.Error())
    }

    // Simple Field (knxNetIpDeviceMacAddress)
    _knxNetIpDeviceMacAddressErr := m.KnxNetIpDeviceMacAddress.Serialize(io)
    if _knxNetIpDeviceMacAddressErr != nil {
        return errors.New("Error serializing 'knxNetIpDeviceMacAddress' field " + _knxNetIpDeviceMacAddressErr.Error())
    }

    // Array Field (deviceFriendlyName)
    if m.DeviceFriendlyName != nil {
        for _, _element := range m.DeviceFriendlyName {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'deviceFriendlyName' field " + _elementErr.Error())
            }
        }
    }

    return nil
}

func (m *DIBDeviceInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "descriptionType":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.DescriptionType = data
            case "knxMedium":
                var data KnxMedium
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.KnxMedium = data
            case "deviceStatus":
                var data *DeviceStatus
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.DeviceStatus = data
            case "knxAddress":
                var data *KnxAddress
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.KnxAddress = data
            case "projectInstallationIdentifier":
                var data *ProjectInstallationIdentifier
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.ProjectInstallationIdentifier = data
            case "knxNetIpDeviceSerialNumber":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.KnxNetIpDeviceSerialNumber = utils.ByteArrayToInt8Array(_decoded[0:_len])
            case "knxNetIpDeviceMulticastAddress":
                var data *IPAddress
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.KnxNetIpDeviceMulticastAddress = data
            case "knxNetIpDeviceMacAddress":
                var data *MACAddress
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.KnxNetIpDeviceMacAddress = data
            case "deviceFriendlyName":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.DeviceFriendlyName = utils.ByteArrayToInt8Array(_decoded[0:_len])
            }
        }
    }
}

func (m *DIBDeviceInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := "org.apache.plc4x.java.knxnetip.readwrite.DIBDeviceInfo"
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.DescriptionType, xml.StartElement{Name: xml.Name{Local: "descriptionType"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.KnxMedium, xml.StartElement{Name: xml.Name{Local: "knxMedium"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.DeviceStatus, xml.StartElement{Name: xml.Name{Local: "deviceStatus"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.KnxAddress, xml.StartElement{Name: xml.Name{Local: "knxAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ProjectInstallationIdentifier, xml.StartElement{Name: xml.Name{Local: "projectInstallationIdentifier"}}); err != nil {
        return err
    }
    _encodedKnxNetIpDeviceSerialNumber := make([]byte, base64.StdEncoding.EncodedLen(len(m.KnxNetIpDeviceSerialNumber)))
    base64.StdEncoding.Encode(_encodedKnxNetIpDeviceSerialNumber, utils.Int8ArrayToByteArray(m.KnxNetIpDeviceSerialNumber))
    if err := e.EncodeElement(_encodedKnxNetIpDeviceSerialNumber, xml.StartElement{Name: xml.Name{Local: "knxNetIpDeviceSerialNumber"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.KnxNetIpDeviceMulticastAddress, xml.StartElement{Name: xml.Name{Local: "knxNetIpDeviceMulticastAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.KnxNetIpDeviceMacAddress, xml.StartElement{Name: xml.Name{Local: "knxNetIpDeviceMacAddress"}}); err != nil {
        return err
    }
    _encodedDeviceFriendlyName := make([]byte, base64.StdEncoding.EncodedLen(len(m.DeviceFriendlyName)))
    base64.StdEncoding.Encode(_encodedDeviceFriendlyName, utils.Int8ArrayToByteArray(m.DeviceFriendlyName))
    if err := e.EncodeElement(_encodedDeviceFriendlyName, xml.StartElement{Name: xml.Name{Local: "deviceFriendlyName"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

