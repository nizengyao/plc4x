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

// MeteringCommandTypeContainer is an enum
type MeteringCommandTypeContainer uint8

type IMeteringCommandTypeContainer interface {
	utils.Serializable
	NumBytes() uint8
	CommandType() MeteringCommandType
}

const (
	MeteringCommandTypeContainer_MeteringCommandEvent_0Bytes MeteringCommandTypeContainer = 0x08
	MeteringCommandTypeContainer_MeteringCommandEvent_1Bytes MeteringCommandTypeContainer = 0x09
	MeteringCommandTypeContainer_MeteringCommandEvent_2Bytes MeteringCommandTypeContainer = 0x0A
	MeteringCommandTypeContainer_MeteringCommandEvent_3Bytes MeteringCommandTypeContainer = 0x0B
	MeteringCommandTypeContainer_MeteringCommandEvent_4Bytes MeteringCommandTypeContainer = 0x0C
	MeteringCommandTypeContainer_MeteringCommandEvent_5Bytes MeteringCommandTypeContainer = 0x0D
	MeteringCommandTypeContainer_MeteringCommandEvent_6Bytes MeteringCommandTypeContainer = 0x0E
	MeteringCommandTypeContainer_MeteringCommandEvent_7Bytes MeteringCommandTypeContainer = 0x0F
)

var MeteringCommandTypeContainerValues []MeteringCommandTypeContainer

func init() {
	_ = errors.New
	MeteringCommandTypeContainerValues = []MeteringCommandTypeContainer{
		MeteringCommandTypeContainer_MeteringCommandEvent_0Bytes,
		MeteringCommandTypeContainer_MeteringCommandEvent_1Bytes,
		MeteringCommandTypeContainer_MeteringCommandEvent_2Bytes,
		MeteringCommandTypeContainer_MeteringCommandEvent_3Bytes,
		MeteringCommandTypeContainer_MeteringCommandEvent_4Bytes,
		MeteringCommandTypeContainer_MeteringCommandEvent_5Bytes,
		MeteringCommandTypeContainer_MeteringCommandEvent_6Bytes,
		MeteringCommandTypeContainer_MeteringCommandEvent_7Bytes,
	}
}

func (e MeteringCommandTypeContainer) NumBytes() uint8 {
	switch e {
	case 0x08:
		{ /* '0x08' */
			return 0
		}
	case 0x09:
		{ /* '0x09' */
			return 1
		}
	case 0x0A:
		{ /* '0x0A' */
			return 2
		}
	case 0x0B:
		{ /* '0x0B' */
			return 3
		}
	case 0x0C:
		{ /* '0x0C' */
			return 4
		}
	case 0x0D:
		{ /* '0x0D' */
			return 5
		}
	case 0x0E:
		{ /* '0x0E' */
			return 6
		}
	case 0x0F:
		{ /* '0x0F' */
			return 7
		}
	default:
		{
			return 0
		}
	}
}

func MeteringCommandTypeContainerFirstEnumForFieldNumBytes(value uint8) (MeteringCommandTypeContainer, error) {
	for _, sizeValue := range MeteringCommandTypeContainerValues {
		if sizeValue.NumBytes() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumBytes not found", value)
}

func (e MeteringCommandTypeContainer) CommandType() MeteringCommandType {
	switch e {
	case 0x08:
		{ /* '0x08' */
			return MeteringCommandType_EVENT
		}
	case 0x09:
		{ /* '0x09' */
			return MeteringCommandType_EVENT
		}
	case 0x0A:
		{ /* '0x0A' */
			return MeteringCommandType_EVENT
		}
	case 0x0B:
		{ /* '0x0B' */
			return MeteringCommandType_EVENT
		}
	case 0x0C:
		{ /* '0x0C' */
			return MeteringCommandType_EVENT
		}
	case 0x0D:
		{ /* '0x0D' */
			return MeteringCommandType_EVENT
		}
	case 0x0E:
		{ /* '0x0E' */
			return MeteringCommandType_EVENT
		}
	case 0x0F:
		{ /* '0x0F' */
			return MeteringCommandType_EVENT
		}
	default:
		{
			return 0
		}
	}
}

func MeteringCommandTypeContainerFirstEnumForFieldCommandType(value MeteringCommandType) (MeteringCommandTypeContainer, error) {
	for _, sizeValue := range MeteringCommandTypeContainerValues {
		if sizeValue.CommandType() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing CommandType not found", value)
}
func MeteringCommandTypeContainerByValue(value uint8) (enum MeteringCommandTypeContainer, ok bool) {
	switch value {
	case 0x08:
		return MeteringCommandTypeContainer_MeteringCommandEvent_0Bytes, true
	case 0x09:
		return MeteringCommandTypeContainer_MeteringCommandEvent_1Bytes, true
	case 0x0A:
		return MeteringCommandTypeContainer_MeteringCommandEvent_2Bytes, true
	case 0x0B:
		return MeteringCommandTypeContainer_MeteringCommandEvent_3Bytes, true
	case 0x0C:
		return MeteringCommandTypeContainer_MeteringCommandEvent_4Bytes, true
	case 0x0D:
		return MeteringCommandTypeContainer_MeteringCommandEvent_5Bytes, true
	case 0x0E:
		return MeteringCommandTypeContainer_MeteringCommandEvent_6Bytes, true
	case 0x0F:
		return MeteringCommandTypeContainer_MeteringCommandEvent_7Bytes, true
	}
	return 0, false
}

func MeteringCommandTypeContainerByName(value string) (enum MeteringCommandTypeContainer, ok bool) {
	switch value {
	case "MeteringCommandEvent_0Bytes":
		return MeteringCommandTypeContainer_MeteringCommandEvent_0Bytes, true
	case "MeteringCommandEvent_1Bytes":
		return MeteringCommandTypeContainer_MeteringCommandEvent_1Bytes, true
	case "MeteringCommandEvent_2Bytes":
		return MeteringCommandTypeContainer_MeteringCommandEvent_2Bytes, true
	case "MeteringCommandEvent_3Bytes":
		return MeteringCommandTypeContainer_MeteringCommandEvent_3Bytes, true
	case "MeteringCommandEvent_4Bytes":
		return MeteringCommandTypeContainer_MeteringCommandEvent_4Bytes, true
	case "MeteringCommandEvent_5Bytes":
		return MeteringCommandTypeContainer_MeteringCommandEvent_5Bytes, true
	case "MeteringCommandEvent_6Bytes":
		return MeteringCommandTypeContainer_MeteringCommandEvent_6Bytes, true
	case "MeteringCommandEvent_7Bytes":
		return MeteringCommandTypeContainer_MeteringCommandEvent_7Bytes, true
	}
	return 0, false
}

func MeteringCommandTypeContainerKnows(value uint8) bool {
	for _, typeValue := range MeteringCommandTypeContainerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMeteringCommandTypeContainer(structType interface{}) MeteringCommandTypeContainer {
	castFunc := func(typ interface{}) MeteringCommandTypeContainer {
		if sMeteringCommandTypeContainer, ok := typ.(MeteringCommandTypeContainer); ok {
			return sMeteringCommandTypeContainer
		}
		return 0
	}
	return castFunc(structType)
}

func (m MeteringCommandTypeContainer) GetLengthInBits() uint16 {
	return 8
}

func (m MeteringCommandTypeContainer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MeteringCommandTypeContainerParse(readBuffer utils.ReadBuffer) (MeteringCommandTypeContainer, error) {
	val, err := readBuffer.ReadUint8("MeteringCommandTypeContainer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MeteringCommandTypeContainer")
	}
	if enum, ok := MeteringCommandTypeContainerByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return MeteringCommandTypeContainer(val), nil
	} else {
		return enum, nil
	}
}

func (e MeteringCommandTypeContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e MeteringCommandTypeContainer) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("MeteringCommandTypeContainer", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MeteringCommandTypeContainer) PLC4XEnumName() string {
	switch e {
	case MeteringCommandTypeContainer_MeteringCommandEvent_0Bytes:
		return "MeteringCommandEvent_0Bytes"
	case MeteringCommandTypeContainer_MeteringCommandEvent_1Bytes:
		return "MeteringCommandEvent_1Bytes"
	case MeteringCommandTypeContainer_MeteringCommandEvent_2Bytes:
		return "MeteringCommandEvent_2Bytes"
	case MeteringCommandTypeContainer_MeteringCommandEvent_3Bytes:
		return "MeteringCommandEvent_3Bytes"
	case MeteringCommandTypeContainer_MeteringCommandEvent_4Bytes:
		return "MeteringCommandEvent_4Bytes"
	case MeteringCommandTypeContainer_MeteringCommandEvent_5Bytes:
		return "MeteringCommandEvent_5Bytes"
	case MeteringCommandTypeContainer_MeteringCommandEvent_6Bytes:
		return "MeteringCommandEvent_6Bytes"
	case MeteringCommandTypeContainer_MeteringCommandEvent_7Bytes:
		return "MeteringCommandEvent_7Bytes"
	}
	return ""
}

func (e MeteringCommandTypeContainer) String() string {
	return e.PLC4XEnumName()
}
