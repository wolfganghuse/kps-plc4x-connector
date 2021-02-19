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

type KnxLayer uint8

type IKnxLayer interface {
    Serialize(io utils.WriteBuffer) error
}

const(
    KnxLayer_TUNNEL_LINK_LAYER KnxLayer = 0x02
    KnxLayer_TUNNEL_RAW KnxLayer = 0x04
    KnxLayer_TUNNEL_BUSMONITOR KnxLayer = 0x80
)

func KnxLayerByValue(value uint8) KnxLayer {
    switch value {
        case 0x02:
            return KnxLayer_TUNNEL_LINK_LAYER
        case 0x04:
            return KnxLayer_TUNNEL_RAW
        case 0x80:
            return KnxLayer_TUNNEL_BUSMONITOR
    }
    return 0
}

func KnxLayerByName(value string) KnxLayer {
    switch value {
    case "TUNNEL_LINK_LAYER":
        return KnxLayer_TUNNEL_LINK_LAYER
    case "TUNNEL_RAW":
        return KnxLayer_TUNNEL_RAW
    case "TUNNEL_BUSMONITOR":
        return KnxLayer_TUNNEL_BUSMONITOR
    }
    return 0
}

func CastKnxLayer(structType interface{}) KnxLayer {
    castFunc := func(typ interface{}) KnxLayer {
        if sKnxLayer, ok := typ.(KnxLayer); ok {
            return sKnxLayer
        }
        return 0
    }
    return castFunc(structType)
}

func (m KnxLayer) LengthInBits() uint16 {
    return 8
}

func (m KnxLayer) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func KnxLayerParse(io *utils.ReadBuffer) (KnxLayer, error) {
    val, err := io.ReadUint8(8)
    if err != nil {
        return 0, nil
    }
    return KnxLayerByValue(val), nil
}

func (e KnxLayer) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint8(8, uint8(e))
    return err
}

func (e KnxLayer) String() string {
    switch e {
    case KnxLayer_TUNNEL_LINK_LAYER:
        return "TUNNEL_LINK_LAYER"
    case KnxLayer_TUNNEL_RAW:
        return "TUNNEL_RAW"
    case KnxLayer_TUNNEL_BUSMONITOR:
        return "TUNNEL_BUSMONITOR"
    }
    return ""
}
