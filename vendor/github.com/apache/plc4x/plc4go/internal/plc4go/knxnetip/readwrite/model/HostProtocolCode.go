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
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type HostProtocolCode uint8

type IHostProtocolCode interface {
    Serialize(io utils.WriteBuffer) error
}

const(
    HostProtocolCode_IPV4_UDP HostProtocolCode = 0x01
    HostProtocolCode_IPV4_TCP HostProtocolCode = 0x02
)

func HostProtocolCodeByValue(value uint8) HostProtocolCode {
    switch value {
        case 0x01:
            return HostProtocolCode_IPV4_UDP
        case 0x02:
            return HostProtocolCode_IPV4_TCP
    }
    return 0
}

func HostProtocolCodeByName(value string) HostProtocolCode {
    switch value {
    case "IPV4_UDP":
        return HostProtocolCode_IPV4_UDP
    case "IPV4_TCP":
        return HostProtocolCode_IPV4_TCP
    }
    return 0
}

func CastHostProtocolCode(structType interface{}) HostProtocolCode {
    castFunc := func(typ interface{}) HostProtocolCode {
        if sHostProtocolCode, ok := typ.(HostProtocolCode); ok {
            return sHostProtocolCode
        }
        return 0
    }
    return castFunc(structType)
}

func (m HostProtocolCode) LengthInBits() uint16 {
    return 8
}

func (m HostProtocolCode) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func HostProtocolCodeParse(io *utils.ReadBuffer) (HostProtocolCode, error) {
    val, err := io.ReadUint8(8)
    if err != nil {
        return 0, nil
    }
    return HostProtocolCodeByValue(val), nil
}

func (e HostProtocolCode) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint8(8, uint8(e))
    return err
}

func (e HostProtocolCode) String() string {
    switch e {
    case HostProtocolCode_IPV4_UDP:
        return "IPV4_UDP"
    case HostProtocolCode_IPV4_TCP:
        return "IPV4_TCP"
    }
    return ""
}
