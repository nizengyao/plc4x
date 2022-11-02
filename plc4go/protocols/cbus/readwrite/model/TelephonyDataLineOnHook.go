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

// TelephonyDataLineOnHook is the corresponding interface of TelephonyDataLineOnHook
type TelephonyDataLineOnHook interface {
	utils.LengthAware
	utils.Serializable
	TelephonyData
}

// TelephonyDataLineOnHookExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataLineOnHook.
// This is useful for switch cases.
type TelephonyDataLineOnHookExactly interface {
	TelephonyDataLineOnHook
	isTelephonyDataLineOnHook() bool
}

// _TelephonyDataLineOnHook is the data-structure of this message
type _TelephonyDataLineOnHook struct {
	*_TelephonyData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataLineOnHook) InitializeParent(parent TelephonyData, commandTypeContainer TelephonyCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_TelephonyDataLineOnHook) GetParent() TelephonyData {
	return m._TelephonyData
}

// NewTelephonyDataLineOnHook factory function for _TelephonyDataLineOnHook
func NewTelephonyDataLineOnHook(commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataLineOnHook {
	_result := &_TelephonyDataLineOnHook{
		_TelephonyData: NewTelephonyData(commandTypeContainer, argument),
	}
	_result._TelephonyData._TelephonyDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataLineOnHook(structType interface{}) TelephonyDataLineOnHook {
	if casted, ok := structType.(TelephonyDataLineOnHook); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataLineOnHook); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataLineOnHook) GetTypeName() string {
	return "TelephonyDataLineOnHook"
}

func (m *_TelephonyDataLineOnHook) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TelephonyDataLineOnHook) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_TelephonyDataLineOnHook) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TelephonyDataLineOnHookParse(readBuffer utils.ReadBuffer) (TelephonyDataLineOnHook, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataLineOnHook"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataLineOnHook")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TelephonyDataLineOnHook"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataLineOnHook")
	}

	// Create a partially initialized instance
	_child := &_TelephonyDataLineOnHook{
		_TelephonyData: &_TelephonyData{},
	}
	_child._TelephonyData._TelephonyDataChildRequirements = _child
	return _child, nil
}

func (m *_TelephonyDataLineOnHook) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataLineOnHook) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataLineOnHook"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataLineOnHook")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataLineOnHook"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataLineOnHook")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_TelephonyDataLineOnHook) isTelephonyDataLineOnHook() bool {
	return true
}

func (m *_TelephonyDataLineOnHook) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
