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

// BACnetTimerStateChangeValueCharacterString is the corresponding interface of BACnetTimerStateChangeValueCharacterString
type BACnetTimerStateChangeValueCharacterString interface {
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetCharacterStringValue returns CharacterStringValue (property field)
	GetCharacterStringValue() BACnetApplicationTagCharacterString
}

// BACnetTimerStateChangeValueCharacterStringExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueCharacterString.
// This is useful for switch cases.
type BACnetTimerStateChangeValueCharacterStringExactly interface {
	BACnetTimerStateChangeValueCharacterString
	isBACnetTimerStateChangeValueCharacterString() bool
}

// _BACnetTimerStateChangeValueCharacterString is the data-structure of this message
type _BACnetTimerStateChangeValueCharacterString struct {
	*_BACnetTimerStateChangeValue
	CharacterStringValue BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueCharacterString) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueCharacterString) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueCharacterString) GetCharacterStringValue() BACnetApplicationTagCharacterString {
	return m.CharacterStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueCharacterString factory function for _BACnetTimerStateChangeValueCharacterString
func NewBACnetTimerStateChangeValueCharacterString(characterStringValue BACnetApplicationTagCharacterString, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueCharacterString {
	_result := &_BACnetTimerStateChangeValueCharacterString{
		CharacterStringValue:         characterStringValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueCharacterString(structType interface{}) BACnetTimerStateChangeValueCharacterString {
	if casted, ok := structType.(BACnetTimerStateChangeValueCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueCharacterString) GetTypeName() string {
	return "BACnetTimerStateChangeValueCharacterString"
}

func (m *_BACnetTimerStateChangeValueCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTimerStateChangeValueCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (characterStringValue)
	lengthInBits += m.CharacterStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateChangeValueCharacterStringParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (characterStringValue)
	if pullErr := readBuffer.PullContext("characterStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for characterStringValue")
	}
	_characterStringValue, _characterStringValueErr := BACnetApplicationTagParse(readBuffer)
	if _characterStringValueErr != nil {
		return nil, errors.Wrap(_characterStringValueErr, "Error parsing 'characterStringValue' field of BACnetTimerStateChangeValueCharacterString")
	}
	characterStringValue := _characterStringValue.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("characterStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for characterStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueCharacterString")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueCharacterString{
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		CharacterStringValue: characterStringValue,
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueCharacterString) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueCharacterString")
		}

		// Simple Field (characterStringValue)
		if pushErr := writeBuffer.PushContext("characterStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for characterStringValue")
		}
		_characterStringValueErr := writeBuffer.WriteSerializable(m.GetCharacterStringValue())
		if popErr := writeBuffer.PopContext("characterStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for characterStringValue")
		}
		if _characterStringValueErr != nil {
			return errors.Wrap(_characterStringValueErr, "Error serializing 'characterStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueCharacterString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueCharacterString) isBACnetTimerStateChangeValueCharacterString() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
