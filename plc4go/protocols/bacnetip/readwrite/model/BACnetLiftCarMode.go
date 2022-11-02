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

// BACnetLiftCarMode is an enum
type BACnetLiftCarMode uint16

type IBACnetLiftCarMode interface {
	utils.Serializable
}

const (
	BACnetLiftCarMode_UNKNOWN                  BACnetLiftCarMode = 0
	BACnetLiftCarMode_NORMAL                   BACnetLiftCarMode = 1
	BACnetLiftCarMode_VIP                      BACnetLiftCarMode = 2
	BACnetLiftCarMode_HOMING                   BACnetLiftCarMode = 3
	BACnetLiftCarMode_PARKING                  BACnetLiftCarMode = 4
	BACnetLiftCarMode_ATTENDANT_CONTROL        BACnetLiftCarMode = 5
	BACnetLiftCarMode_FIREFIGHTER_CONTROL      BACnetLiftCarMode = 6
	BACnetLiftCarMode_EMERGENCY_POWER          BACnetLiftCarMode = 7
	BACnetLiftCarMode_INSPECTION               BACnetLiftCarMode = 8
	BACnetLiftCarMode_CABINET_RECALL           BACnetLiftCarMode = 9
	BACnetLiftCarMode_EARTHQUAKE_OPERATION     BACnetLiftCarMode = 10
	BACnetLiftCarMode_FIRE_OPERATION           BACnetLiftCarMode = 11
	BACnetLiftCarMode_OUT_OF_SERVICE           BACnetLiftCarMode = 12
	BACnetLiftCarMode_OCCUPANT_EVACUATION      BACnetLiftCarMode = 13
	BACnetLiftCarMode_VENDOR_PROPRIETARY_VALUE BACnetLiftCarMode = 0xFFFF
)

var BACnetLiftCarModeValues []BACnetLiftCarMode

func init() {
	_ = errors.New
	BACnetLiftCarModeValues = []BACnetLiftCarMode{
		BACnetLiftCarMode_UNKNOWN,
		BACnetLiftCarMode_NORMAL,
		BACnetLiftCarMode_VIP,
		BACnetLiftCarMode_HOMING,
		BACnetLiftCarMode_PARKING,
		BACnetLiftCarMode_ATTENDANT_CONTROL,
		BACnetLiftCarMode_FIREFIGHTER_CONTROL,
		BACnetLiftCarMode_EMERGENCY_POWER,
		BACnetLiftCarMode_INSPECTION,
		BACnetLiftCarMode_CABINET_RECALL,
		BACnetLiftCarMode_EARTHQUAKE_OPERATION,
		BACnetLiftCarMode_FIRE_OPERATION,
		BACnetLiftCarMode_OUT_OF_SERVICE,
		BACnetLiftCarMode_OCCUPANT_EVACUATION,
		BACnetLiftCarMode_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLiftCarModeByValue(value uint16) (enum BACnetLiftCarMode, ok bool) {
	switch value {
	case 0:
		return BACnetLiftCarMode_UNKNOWN, true
	case 0xFFFF:
		return BACnetLiftCarMode_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetLiftCarMode_NORMAL, true
	case 10:
		return BACnetLiftCarMode_EARTHQUAKE_OPERATION, true
	case 11:
		return BACnetLiftCarMode_FIRE_OPERATION, true
	case 12:
		return BACnetLiftCarMode_OUT_OF_SERVICE, true
	case 13:
		return BACnetLiftCarMode_OCCUPANT_EVACUATION, true
	case 2:
		return BACnetLiftCarMode_VIP, true
	case 3:
		return BACnetLiftCarMode_HOMING, true
	case 4:
		return BACnetLiftCarMode_PARKING, true
	case 5:
		return BACnetLiftCarMode_ATTENDANT_CONTROL, true
	case 6:
		return BACnetLiftCarMode_FIREFIGHTER_CONTROL, true
	case 7:
		return BACnetLiftCarMode_EMERGENCY_POWER, true
	case 8:
		return BACnetLiftCarMode_INSPECTION, true
	case 9:
		return BACnetLiftCarMode_CABINET_RECALL, true
	}
	return 0, false
}

func BACnetLiftCarModeByName(value string) (enum BACnetLiftCarMode, ok bool) {
	switch value {
	case "UNKNOWN":
		return BACnetLiftCarMode_UNKNOWN, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetLiftCarMode_VENDOR_PROPRIETARY_VALUE, true
	case "NORMAL":
		return BACnetLiftCarMode_NORMAL, true
	case "EARTHQUAKE_OPERATION":
		return BACnetLiftCarMode_EARTHQUAKE_OPERATION, true
	case "FIRE_OPERATION":
		return BACnetLiftCarMode_FIRE_OPERATION, true
	case "OUT_OF_SERVICE":
		return BACnetLiftCarMode_OUT_OF_SERVICE, true
	case "OCCUPANT_EVACUATION":
		return BACnetLiftCarMode_OCCUPANT_EVACUATION, true
	case "VIP":
		return BACnetLiftCarMode_VIP, true
	case "HOMING":
		return BACnetLiftCarMode_HOMING, true
	case "PARKING":
		return BACnetLiftCarMode_PARKING, true
	case "ATTENDANT_CONTROL":
		return BACnetLiftCarMode_ATTENDANT_CONTROL, true
	case "FIREFIGHTER_CONTROL":
		return BACnetLiftCarMode_FIREFIGHTER_CONTROL, true
	case "EMERGENCY_POWER":
		return BACnetLiftCarMode_EMERGENCY_POWER, true
	case "INSPECTION":
		return BACnetLiftCarMode_INSPECTION, true
	case "CABINET_RECALL":
		return BACnetLiftCarMode_CABINET_RECALL, true
	}
	return 0, false
}

func BACnetLiftCarModeKnows(value uint16) bool {
	for _, typeValue := range BACnetLiftCarModeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLiftCarMode(structType interface{}) BACnetLiftCarMode {
	castFunc := func(typ interface{}) BACnetLiftCarMode {
		if sBACnetLiftCarMode, ok := typ.(BACnetLiftCarMode); ok {
			return sBACnetLiftCarMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLiftCarMode) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetLiftCarMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLiftCarModeParse(readBuffer utils.ReadBuffer) (BACnetLiftCarMode, error) {
	val, err := readBuffer.ReadUint16("BACnetLiftCarMode", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLiftCarMode")
	}
	if enum, ok := BACnetLiftCarModeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetLiftCarMode(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLiftCarMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLiftCarMode) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetLiftCarMode", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLiftCarMode) PLC4XEnumName() string {
	switch e {
	case BACnetLiftCarMode_UNKNOWN:
		return "UNKNOWN"
	case BACnetLiftCarMode_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLiftCarMode_NORMAL:
		return "NORMAL"
	case BACnetLiftCarMode_EARTHQUAKE_OPERATION:
		return "EARTHQUAKE_OPERATION"
	case BACnetLiftCarMode_FIRE_OPERATION:
		return "FIRE_OPERATION"
	case BACnetLiftCarMode_OUT_OF_SERVICE:
		return "OUT_OF_SERVICE"
	case BACnetLiftCarMode_OCCUPANT_EVACUATION:
		return "OCCUPANT_EVACUATION"
	case BACnetLiftCarMode_VIP:
		return "VIP"
	case BACnetLiftCarMode_HOMING:
		return "HOMING"
	case BACnetLiftCarMode_PARKING:
		return "PARKING"
	case BACnetLiftCarMode_ATTENDANT_CONTROL:
		return "ATTENDANT_CONTROL"
	case BACnetLiftCarMode_FIREFIGHTER_CONTROL:
		return "FIREFIGHTER_CONTROL"
	case BACnetLiftCarMode_EMERGENCY_POWER:
		return "EMERGENCY_POWER"
	case BACnetLiftCarMode_INSPECTION:
		return "INSPECTION"
	case BACnetLiftCarMode_CABINET_RECALL:
		return "CABINET_RECALL"
	}
	return ""
}

func (e BACnetLiftCarMode) String() string {
	return e.PLC4XEnumName()
}
