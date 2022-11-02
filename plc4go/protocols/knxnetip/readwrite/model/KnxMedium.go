/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"encoding/binary"

	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// KnxMedium is an enum
type KnxMedium uint8

type IKnxMedium interface {
	utils.Serializable
}

const (
	KnxMedium_MEDIUM_RESERVED_1 KnxMedium = 0x01
	KnxMedium_MEDIUM_TP1        KnxMedium = 0x02
	KnxMedium_MEDIUM_PL110      KnxMedium = 0x04
	KnxMedium_MEDIUM_RESERVED_2 KnxMedium = 0x08
	KnxMedium_MEDIUM_RF         KnxMedium = 0x10
	KnxMedium_MEDIUM_KNX_IP     KnxMedium = 0x20
)

var KnxMediumValues []KnxMedium

func init() {
	_ = errors.New
	KnxMediumValues = []KnxMedium{
		KnxMedium_MEDIUM_RESERVED_1,
		KnxMedium_MEDIUM_TP1,
		KnxMedium_MEDIUM_PL110,
		KnxMedium_MEDIUM_RESERVED_2,
		KnxMedium_MEDIUM_RF,
		KnxMedium_MEDIUM_KNX_IP,
	}
}

func KnxMediumByValue(value uint8) (enum KnxMedium, ok bool) {
	switch value {
	case 0x01:
		return KnxMedium_MEDIUM_RESERVED_1, true
	case 0x02:
		return KnxMedium_MEDIUM_TP1, true
	case 0x04:
		return KnxMedium_MEDIUM_PL110, true
	case 0x08:
		return KnxMedium_MEDIUM_RESERVED_2, true
	case 0x10:
		return KnxMedium_MEDIUM_RF, true
	case 0x20:
		return KnxMedium_MEDIUM_KNX_IP, true
	}
	return 0, false
}

func KnxMediumByName(value string) (enum KnxMedium, ok bool) {
	switch value {
	case "MEDIUM_RESERVED_1":
		return KnxMedium_MEDIUM_RESERVED_1, true
	case "MEDIUM_TP1":
		return KnxMedium_MEDIUM_TP1, true
	case "MEDIUM_PL110":
		return KnxMedium_MEDIUM_PL110, true
	case "MEDIUM_RESERVED_2":
		return KnxMedium_MEDIUM_RESERVED_2, true
	case "MEDIUM_RF":
		return KnxMedium_MEDIUM_RF, true
	case "MEDIUM_KNX_IP":
		return KnxMedium_MEDIUM_KNX_IP, true
	}
	return 0, false
}

func KnxMediumKnows(value uint8) bool {
	for _, typeValue := range KnxMediumValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastKnxMedium(structType interface{}) KnxMedium {
	castFunc := func(typ interface{}) KnxMedium {
		if sKnxMedium, ok := typ.(KnxMedium); ok {
			return sKnxMedium
		}
		return 0
	}
	return castFunc(structType)
}

func (m KnxMedium) GetLengthInBits() uint16 {
	return 8
}

func (m KnxMedium) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxMediumParse(readBuffer utils.ReadBuffer) (KnxMedium, error) {
	val, err := readBuffer.ReadUint8("KnxMedium", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading KnxMedium")
	}
	if enum, ok := KnxMediumByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return KnxMedium(val), nil
	} else {
		return enum, nil
	}
}

func (e KnxMedium) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e KnxMedium) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("KnxMedium", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e KnxMedium) PLC4XEnumName() string {
	switch e {
	case KnxMedium_MEDIUM_RESERVED_1:
		return "MEDIUM_RESERVED_1"
	case KnxMedium_MEDIUM_TP1:
		return "MEDIUM_TP1"
	case KnxMedium_MEDIUM_PL110:
		return "MEDIUM_PL110"
	case KnxMedium_MEDIUM_RESERVED_2:
		return "MEDIUM_RESERVED_2"
	case KnxMedium_MEDIUM_RF:
		return "MEDIUM_RF"
	case KnxMedium_MEDIUM_KNX_IP:
		return "MEDIUM_KNX_IP"
	}
	return ""
}

func (e KnxMedium) String() string {
	return e.PLC4XEnumName()
}
