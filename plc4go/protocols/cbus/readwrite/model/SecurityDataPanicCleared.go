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

// SecurityDataPanicCleared is the corresponding interface of SecurityDataPanicCleared
type SecurityDataPanicCleared interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataPanicClearedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataPanicCleared.
// This is useful for switch cases.
type SecurityDataPanicClearedExactly interface {
	SecurityDataPanicCleared
	isSecurityDataPanicCleared() bool
}

// _SecurityDataPanicCleared is the data-structure of this message
type _SecurityDataPanicCleared struct {
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

func (m *_SecurityDataPanicCleared) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataPanicCleared) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataPanicCleared factory function for _SecurityDataPanicCleared
func NewSecurityDataPanicCleared(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataPanicCleared {
	_result := &_SecurityDataPanicCleared{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataPanicCleared(structType interface{}) SecurityDataPanicCleared {
	if casted, ok := structType.(SecurityDataPanicCleared); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataPanicCleared); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataPanicCleared) GetTypeName() string {
	return "SecurityDataPanicCleared"
}

func (m *_SecurityDataPanicCleared) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataPanicCleared) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SecurityDataPanicCleared) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataPanicClearedParse(readBuffer utils.ReadBuffer) (SecurityDataPanicCleared, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataPanicCleared"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataPanicCleared")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataPanicCleared"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataPanicCleared")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataPanicCleared{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataPanicCleared) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataPanicCleared) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataPanicCleared"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataPanicCleared")
		}

		if popErr := writeBuffer.PopContext("SecurityDataPanicCleared"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataPanicCleared")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataPanicCleared) isSecurityDataPanicCleared() bool {
	return true
}

func (m *_SecurityDataPanicCleared) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
