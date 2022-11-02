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

// BACnetSecurityPolicy is an enum
type BACnetSecurityPolicy uint8

type IBACnetSecurityPolicy interface {
	utils.Serializable
}

const (
	BACnetSecurityPolicy_PLAIN_NON_TRUSTED BACnetSecurityPolicy = 0
	BACnetSecurityPolicy_PLAIN_TRUSTED     BACnetSecurityPolicy = 1
	BACnetSecurityPolicy_SIGNED_TRUSTED    BACnetSecurityPolicy = 2
	BACnetSecurityPolicy_ENCRYPTED_TRUSTED BACnetSecurityPolicy = 3
)

var BACnetSecurityPolicyValues []BACnetSecurityPolicy

func init() {
	_ = errors.New
	BACnetSecurityPolicyValues = []BACnetSecurityPolicy{
		BACnetSecurityPolicy_PLAIN_NON_TRUSTED,
		BACnetSecurityPolicy_PLAIN_TRUSTED,
		BACnetSecurityPolicy_SIGNED_TRUSTED,
		BACnetSecurityPolicy_ENCRYPTED_TRUSTED,
	}
}

func BACnetSecurityPolicyByValue(value uint8) (enum BACnetSecurityPolicy, ok bool) {
	switch value {
	case 0:
		return BACnetSecurityPolicy_PLAIN_NON_TRUSTED, true
	case 1:
		return BACnetSecurityPolicy_PLAIN_TRUSTED, true
	case 2:
		return BACnetSecurityPolicy_SIGNED_TRUSTED, true
	case 3:
		return BACnetSecurityPolicy_ENCRYPTED_TRUSTED, true
	}
	return 0, false
}

func BACnetSecurityPolicyByName(value string) (enum BACnetSecurityPolicy, ok bool) {
	switch value {
	case "PLAIN_NON_TRUSTED":
		return BACnetSecurityPolicy_PLAIN_NON_TRUSTED, true
	case "PLAIN_TRUSTED":
		return BACnetSecurityPolicy_PLAIN_TRUSTED, true
	case "SIGNED_TRUSTED":
		return BACnetSecurityPolicy_SIGNED_TRUSTED, true
	case "ENCRYPTED_TRUSTED":
		return BACnetSecurityPolicy_ENCRYPTED_TRUSTED, true
	}
	return 0, false
}

func BACnetSecurityPolicyKnows(value uint8) bool {
	for _, typeValue := range BACnetSecurityPolicyValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetSecurityPolicy(structType interface{}) BACnetSecurityPolicy {
	castFunc := func(typ interface{}) BACnetSecurityPolicy {
		if sBACnetSecurityPolicy, ok := typ.(BACnetSecurityPolicy); ok {
			return sBACnetSecurityPolicy
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetSecurityPolicy) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetSecurityPolicy) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetSecurityPolicyParse(readBuffer utils.ReadBuffer) (BACnetSecurityPolicy, error) {
	val, err := readBuffer.ReadUint8("BACnetSecurityPolicy", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetSecurityPolicy")
	}
	if enum, ok := BACnetSecurityPolicyByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetSecurityPolicy(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetSecurityPolicy) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetSecurityPolicy) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetSecurityPolicy", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetSecurityPolicy) PLC4XEnumName() string {
	switch e {
	case BACnetSecurityPolicy_PLAIN_NON_TRUSTED:
		return "PLAIN_NON_TRUSTED"
	case BACnetSecurityPolicy_PLAIN_TRUSTED:
		return "PLAIN_TRUSTED"
	case BACnetSecurityPolicy_SIGNED_TRUSTED:
		return "SIGNED_TRUSTED"
	case BACnetSecurityPolicy_ENCRYPTED_TRUSTED:
		return "ENCRYPTED_TRUSTED"
	}
	return ""
}

func (e BACnetSecurityPolicy) String() string {
	return e.PLC4XEnumName()
}
