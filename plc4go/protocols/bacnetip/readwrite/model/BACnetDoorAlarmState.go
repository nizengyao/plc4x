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

// BACnetDoorAlarmState is an enum
type BACnetDoorAlarmState uint8

type IBACnetDoorAlarmState interface {
	utils.Serializable
}

const (
	BACnetDoorAlarmState_NORMAL                   BACnetDoorAlarmState = 0
	BACnetDoorAlarmState_ALARM                    BACnetDoorAlarmState = 1
	BACnetDoorAlarmState_DOOR_OPEN_TOO_LONG       BACnetDoorAlarmState = 2
	BACnetDoorAlarmState_FORCED_OPEN              BACnetDoorAlarmState = 3
	BACnetDoorAlarmState_TAMPER                   BACnetDoorAlarmState = 4
	BACnetDoorAlarmState_DOOR_FAULT               BACnetDoorAlarmState = 5
	BACnetDoorAlarmState_LOCK_DOWN                BACnetDoorAlarmState = 6
	BACnetDoorAlarmState_FREE_ACCESS              BACnetDoorAlarmState = 7
	BACnetDoorAlarmState_EGRESS_OPEN              BACnetDoorAlarmState = 8
	BACnetDoorAlarmState_VENDOR_PROPRIETARY_VALUE BACnetDoorAlarmState = 0xFF
)

var BACnetDoorAlarmStateValues []BACnetDoorAlarmState

func init() {
	_ = errors.New
	BACnetDoorAlarmStateValues = []BACnetDoorAlarmState{
		BACnetDoorAlarmState_NORMAL,
		BACnetDoorAlarmState_ALARM,
		BACnetDoorAlarmState_DOOR_OPEN_TOO_LONG,
		BACnetDoorAlarmState_FORCED_OPEN,
		BACnetDoorAlarmState_TAMPER,
		BACnetDoorAlarmState_DOOR_FAULT,
		BACnetDoorAlarmState_LOCK_DOWN,
		BACnetDoorAlarmState_FREE_ACCESS,
		BACnetDoorAlarmState_EGRESS_OPEN,
		BACnetDoorAlarmState_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetDoorAlarmStateByValue(value uint8) (enum BACnetDoorAlarmState, ok bool) {
	switch value {
	case 0:
		return BACnetDoorAlarmState_NORMAL, true
	case 0xFF:
		return BACnetDoorAlarmState_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetDoorAlarmState_ALARM, true
	case 2:
		return BACnetDoorAlarmState_DOOR_OPEN_TOO_LONG, true
	case 3:
		return BACnetDoorAlarmState_FORCED_OPEN, true
	case 4:
		return BACnetDoorAlarmState_TAMPER, true
	case 5:
		return BACnetDoorAlarmState_DOOR_FAULT, true
	case 6:
		return BACnetDoorAlarmState_LOCK_DOWN, true
	case 7:
		return BACnetDoorAlarmState_FREE_ACCESS, true
	case 8:
		return BACnetDoorAlarmState_EGRESS_OPEN, true
	}
	return 0, false
}

func BACnetDoorAlarmStateByName(value string) (enum BACnetDoorAlarmState, ok bool) {
	switch value {
	case "NORMAL":
		return BACnetDoorAlarmState_NORMAL, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetDoorAlarmState_VENDOR_PROPRIETARY_VALUE, true
	case "ALARM":
		return BACnetDoorAlarmState_ALARM, true
	case "DOOR_OPEN_TOO_LONG":
		return BACnetDoorAlarmState_DOOR_OPEN_TOO_LONG, true
	case "FORCED_OPEN":
		return BACnetDoorAlarmState_FORCED_OPEN, true
	case "TAMPER":
		return BACnetDoorAlarmState_TAMPER, true
	case "DOOR_FAULT":
		return BACnetDoorAlarmState_DOOR_FAULT, true
	case "LOCK_DOWN":
		return BACnetDoorAlarmState_LOCK_DOWN, true
	case "FREE_ACCESS":
		return BACnetDoorAlarmState_FREE_ACCESS, true
	case "EGRESS_OPEN":
		return BACnetDoorAlarmState_EGRESS_OPEN, true
	}
	return 0, false
}

func BACnetDoorAlarmStateKnows(value uint8) bool {
	for _, typeValue := range BACnetDoorAlarmStateValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetDoorAlarmState(structType interface{}) BACnetDoorAlarmState {
	castFunc := func(typ interface{}) BACnetDoorAlarmState {
		if sBACnetDoorAlarmState, ok := typ.(BACnetDoorAlarmState); ok {
			return sBACnetDoorAlarmState
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetDoorAlarmState) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetDoorAlarmState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDoorAlarmStateParse(readBuffer utils.ReadBuffer) (BACnetDoorAlarmState, error) {
	val, err := readBuffer.ReadUint8("BACnetDoorAlarmState", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetDoorAlarmState")
	}
	if enum, ok := BACnetDoorAlarmStateByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetDoorAlarmState(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetDoorAlarmState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetDoorAlarmState) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetDoorAlarmState", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetDoorAlarmState) PLC4XEnumName() string {
	switch e {
	case BACnetDoorAlarmState_NORMAL:
		return "NORMAL"
	case BACnetDoorAlarmState_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetDoorAlarmState_ALARM:
		return "ALARM"
	case BACnetDoorAlarmState_DOOR_OPEN_TOO_LONG:
		return "DOOR_OPEN_TOO_LONG"
	case BACnetDoorAlarmState_FORCED_OPEN:
		return "FORCED_OPEN"
	case BACnetDoorAlarmState_TAMPER:
		return "TAMPER"
	case BACnetDoorAlarmState_DOOR_FAULT:
		return "DOOR_FAULT"
	case BACnetDoorAlarmState_LOCK_DOWN:
		return "LOCK_DOWN"
	case BACnetDoorAlarmState_FREE_ACCESS:
		return "FREE_ACCESS"
	case BACnetDoorAlarmState_EGRESS_OPEN:
		return "EGRESS_OPEN"
	}
	return ""
}

func (e BACnetDoorAlarmState) String() string {
	return e.PLC4XEnumName()
}
