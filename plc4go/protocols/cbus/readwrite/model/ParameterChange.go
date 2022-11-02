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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const ParameterChange_SPECIALCHAR1 byte = 0x3D
const ParameterChange_SPECIALCHAR2 byte = 0x3D

// ParameterChange is the corresponding interface of ParameterChange
type ParameterChange interface {
	utils.LengthAware
	utils.Serializable
}

// ParameterChangeExactly can be used when we want exactly this type and not a type which fulfills ParameterChange.
// This is useful for switch cases.
type ParameterChangeExactly interface {
	ParameterChange
	isParameterChange() bool
}

// _ParameterChange is the data-structure of this message
type _ParameterChange struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ParameterChange) GetSpecialChar1() byte {
	return ParameterChange_SPECIALCHAR1
}

func (m *_ParameterChange) GetSpecialChar2() byte {
	return ParameterChange_SPECIALCHAR2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterChange factory function for _ParameterChange
func NewParameterChange() *_ParameterChange {
	return &_ParameterChange{}
}

// Deprecated: use the interface for direct cast
func CastParameterChange(structType interface{}) ParameterChange {
	if casted, ok := structType.(ParameterChange); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterChange); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterChange) GetTypeName() string {
	return "ParameterChange"
}

func (m *_ParameterChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ParameterChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (specialChar1)
	lengthInBits += 8

	// Const Field (specialChar2)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ParameterChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ParameterChangeParse(readBuffer utils.ReadBuffer) (ParameterChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (specialChar1)
	specialChar1, _specialChar1Err := readBuffer.ReadByte("specialChar1")
	if _specialChar1Err != nil {
		return nil, errors.Wrap(_specialChar1Err, "Error parsing 'specialChar1' field of ParameterChange")
	}
	if specialChar1 != ParameterChange_SPECIALCHAR1 {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ParameterChange_SPECIALCHAR1) + " but got " + fmt.Sprintf("%d", specialChar1))
	}

	// Const Field (specialChar2)
	specialChar2, _specialChar2Err := readBuffer.ReadByte("specialChar2")
	if _specialChar2Err != nil {
		return nil, errors.Wrap(_specialChar2Err, "Error parsing 'specialChar2' field of ParameterChange")
	}
	if specialChar2 != ParameterChange_SPECIALCHAR2 {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ParameterChange_SPECIALCHAR2) + " but got " + fmt.Sprintf("%d", specialChar2))
	}

	if closeErr := readBuffer.CloseContext("ParameterChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterChange")
	}

	// Create the instance
	return &_ParameterChange{}, nil
}

func (m *_ParameterChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterChange) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ParameterChange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ParameterChange")
	}

	// Const Field (specialChar1)
	_specialChar1Err := writeBuffer.WriteByte("specialChar1", 0x3D)
	if _specialChar1Err != nil {
		return errors.Wrap(_specialChar1Err, "Error serializing 'specialChar1' field")
	}

	// Const Field (specialChar2)
	_specialChar2Err := writeBuffer.WriteByte("specialChar2", 0x3D)
	if _specialChar2Err != nil {
		return errors.Wrap(_specialChar2Err, "Error serializing 'specialChar2' field")
	}

	if popErr := writeBuffer.PopContext("ParameterChange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ParameterChange")
	}
	return nil
}

func (m *_ParameterChange) isParameterChange() bool {
	return true
}

func (m *_ParameterChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
