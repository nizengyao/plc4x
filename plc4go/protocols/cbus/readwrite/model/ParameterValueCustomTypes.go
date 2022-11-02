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

// ParameterValueCustomTypes is the corresponding interface of ParameterValueCustomTypes
type ParameterValueCustomTypes interface {
	utils.LengthAware
	utils.Serializable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() CustomTypes
}

// ParameterValueCustomTypesExactly can be used when we want exactly this type and not a type which fulfills ParameterValueCustomTypes.
// This is useful for switch cases.
type ParameterValueCustomTypesExactly interface {
	ParameterValueCustomTypes
	isParameterValueCustomTypes() bool
}

// _ParameterValueCustomTypes is the data-structure of this message
type _ParameterValueCustomTypes struct {
	*_ParameterValue
	Value CustomTypes
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueCustomTypes) GetParameterType() ParameterType {
	return ParameterType_CUSTOM_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueCustomTypes) InitializeParent(parent ParameterValue) {}

func (m *_ParameterValueCustomTypes) GetParent() ParameterValue {
	return m._ParameterValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueCustomTypes) GetValue() CustomTypes {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterValueCustomTypes factory function for _ParameterValueCustomTypes
func NewParameterValueCustomTypes(value CustomTypes, numBytes uint8) *_ParameterValueCustomTypes {
	_result := &_ParameterValueCustomTypes{
		Value:           value,
		_ParameterValue: NewParameterValue(numBytes),
	}
	_result._ParameterValue._ParameterValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterValueCustomTypes(structType interface{}) ParameterValueCustomTypes {
	if casted, ok := structType.(ParameterValueCustomTypes); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueCustomTypes); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueCustomTypes) GetTypeName() string {
	return "ParameterValueCustomTypes"
}

func (m *_ParameterValueCustomTypes) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ParameterValueCustomTypes) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits()

	return lengthInBits
}

func (m *_ParameterValueCustomTypes) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ParameterValueCustomTypesParse(readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (ParameterValueCustomTypes, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueCustomTypes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueCustomTypes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := CustomTypesParse(readBuffer, uint8(numBytes))
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ParameterValueCustomTypes")
	}
	value := _value.(CustomTypes)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("ParameterValueCustomTypes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueCustomTypes")
	}

	// Create a partially initialized instance
	_child := &_ParameterValueCustomTypes{
		_ParameterValue: &_ParameterValue{
			NumBytes: numBytes,
		},
		Value: value,
	}
	_child._ParameterValue._ParameterValueChildRequirements = _child
	return _child, nil
}

func (m *_ParameterValueCustomTypes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueCustomTypes) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueCustomTypes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueCustomTypes")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueCustomTypes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueCustomTypes")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ParameterValueCustomTypes) isParameterValueCustomTypes() bool {
	return true
}

func (m *_ParameterValueCustomTypes) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
