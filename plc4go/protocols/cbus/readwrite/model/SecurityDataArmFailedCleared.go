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

// SecurityDataArmFailedCleared is the corresponding interface of SecurityDataArmFailedCleared
type SecurityDataArmFailedCleared interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataArmFailedClearedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataArmFailedCleared.
// This is useful for switch cases.
type SecurityDataArmFailedClearedExactly interface {
	SecurityDataArmFailedCleared
	isSecurityDataArmFailedCleared() bool
}

// _SecurityDataArmFailedCleared is the data-structure of this message
type _SecurityDataArmFailedCleared struct {
	*_SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataArmFailedCleared) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataArmFailedCleared) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataArmFailedCleared factory function for _SecurityDataArmFailedCleared
func NewSecurityDataArmFailedCleared(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataArmFailedCleared {
	_result := &_SecurityDataArmFailedCleared{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataArmFailedCleared(structType interface{}) SecurityDataArmFailedCleared {
	if casted, ok := structType.(SecurityDataArmFailedCleared); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataArmFailedCleared); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataArmFailedCleared) GetTypeName() string {
	return "SecurityDataArmFailedCleared"
}

func (m *_SecurityDataArmFailedCleared) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataArmFailedCleared) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SecurityDataArmFailedCleared) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataArmFailedClearedParse(readBuffer utils.ReadBuffer) (SecurityDataArmFailedCleared, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataArmFailedCleared"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataArmFailedCleared")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataArmFailedCleared"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataArmFailedCleared")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataArmFailedCleared{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataArmFailedCleared) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataArmFailedCleared) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataArmFailedCleared"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataArmFailedCleared")
		}

		if popErr := writeBuffer.PopContext("SecurityDataArmFailedCleared"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataArmFailedCleared")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataArmFailedCleared) isSecurityDataArmFailedCleared() bool {
	return true
}

func (m *_SecurityDataArmFailedCleared) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
