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

// DialOutFailureReason is an enum
type DialOutFailureReason uint8

type IDialOutFailureReason interface {
	utils.Serializable
}

const (
	DialOutFailureReason_NO_DIAL_TONE                           DialOutFailureReason = 0x01
	DialOutFailureReason_NO_ANSWER                              DialOutFailureReason = 0x02
	DialOutFailureReason_NO_VALID_ACKNOWLEDGEMENT_OF_PROMPTS    DialOutFailureReason = 0x03
	DialOutFailureReason_NUMBER_WAS_UNOBTAINABLE_DOES_NOT_EXIST DialOutFailureReason = 0x04
	DialOutFailureReason_NUMBER_WAS_BUSY                        DialOutFailureReason = 0x05
	DialOutFailureReason_INTERNAL_FAILURE                       DialOutFailureReason = 0x06
)

var DialOutFailureReasonValues []DialOutFailureReason

func init() {
	_ = errors.New
	DialOutFailureReasonValues = []DialOutFailureReason{
		DialOutFailureReason_NO_DIAL_TONE,
		DialOutFailureReason_NO_ANSWER,
		DialOutFailureReason_NO_VALID_ACKNOWLEDGEMENT_OF_PROMPTS,
		DialOutFailureReason_NUMBER_WAS_UNOBTAINABLE_DOES_NOT_EXIST,
		DialOutFailureReason_NUMBER_WAS_BUSY,
		DialOutFailureReason_INTERNAL_FAILURE,
	}
}

func DialOutFailureReasonByValue(value uint8) (enum DialOutFailureReason, ok bool) {
	switch value {
	case 0x01:
		return DialOutFailureReason_NO_DIAL_TONE, true
	case 0x02:
		return DialOutFailureReason_NO_ANSWER, true
	case 0x03:
		return DialOutFailureReason_NO_VALID_ACKNOWLEDGEMENT_OF_PROMPTS, true
	case 0x04:
		return DialOutFailureReason_NUMBER_WAS_UNOBTAINABLE_DOES_NOT_EXIST, true
	case 0x05:
		return DialOutFailureReason_NUMBER_WAS_BUSY, true
	case 0x06:
		return DialOutFailureReason_INTERNAL_FAILURE, true
	}
	return 0, false
}

func DialOutFailureReasonByName(value string) (enum DialOutFailureReason, ok bool) {
	switch value {
	case "NO_DIAL_TONE":
		return DialOutFailureReason_NO_DIAL_TONE, true
	case "NO_ANSWER":
		return DialOutFailureReason_NO_ANSWER, true
	case "NO_VALID_ACKNOWLEDGEMENT_OF_PROMPTS":
		return DialOutFailureReason_NO_VALID_ACKNOWLEDGEMENT_OF_PROMPTS, true
	case "NUMBER_WAS_UNOBTAINABLE_DOES_NOT_EXIST":
		return DialOutFailureReason_NUMBER_WAS_UNOBTAINABLE_DOES_NOT_EXIST, true
	case "NUMBER_WAS_BUSY":
		return DialOutFailureReason_NUMBER_WAS_BUSY, true
	case "INTERNAL_FAILURE":
		return DialOutFailureReason_INTERNAL_FAILURE, true
	}
	return 0, false
}

func DialOutFailureReasonKnows(value uint8) bool {
	for _, typeValue := range DialOutFailureReasonValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDialOutFailureReason(structType interface{}) DialOutFailureReason {
	castFunc := func(typ interface{}) DialOutFailureReason {
		if sDialOutFailureReason, ok := typ.(DialOutFailureReason); ok {
			return sDialOutFailureReason
		}
		return 0
	}
	return castFunc(structType)
}

func (m DialOutFailureReason) GetLengthInBits() uint16 {
	return 8
}

func (m DialOutFailureReason) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DialOutFailureReasonParse(readBuffer utils.ReadBuffer) (DialOutFailureReason, error) {
	val, err := readBuffer.ReadUint8("DialOutFailureReason", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DialOutFailureReason")
	}
	if enum, ok := DialOutFailureReasonByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return DialOutFailureReason(val), nil
	} else {
		return enum, nil
	}
}

func (e DialOutFailureReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DialOutFailureReason) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("DialOutFailureReason", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DialOutFailureReason) PLC4XEnumName() string {
	switch e {
	case DialOutFailureReason_NO_DIAL_TONE:
		return "NO_DIAL_TONE"
	case DialOutFailureReason_NO_ANSWER:
		return "NO_ANSWER"
	case DialOutFailureReason_NO_VALID_ACKNOWLEDGEMENT_OF_PROMPTS:
		return "NO_VALID_ACKNOWLEDGEMENT_OF_PROMPTS"
	case DialOutFailureReason_NUMBER_WAS_UNOBTAINABLE_DOES_NOT_EXIST:
		return "NUMBER_WAS_UNOBTAINABLE_DOES_NOT_EXIST"
	case DialOutFailureReason_NUMBER_WAS_BUSY:
		return "NUMBER_WAS_BUSY"
	case DialOutFailureReason_INTERNAL_FAILURE:
		return "INTERNAL_FAILURE"
	}
	return ""
}

func (e DialOutFailureReason) String() string {
	return e.PLC4XEnumName()
}
