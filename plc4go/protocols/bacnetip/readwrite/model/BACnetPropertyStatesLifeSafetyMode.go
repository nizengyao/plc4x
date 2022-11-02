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

// BACnetPropertyStatesLifeSafetyMode is the corresponding interface of BACnetPropertyStatesLifeSafetyMode
type BACnetPropertyStatesLifeSafetyMode interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLifeSafetyMode returns LifeSafetyMode (property field)
	GetLifeSafetyMode() BACnetLifeSafetyModeTagged
}

// BACnetPropertyStatesLifeSafetyModeExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLifeSafetyMode.
// This is useful for switch cases.
type BACnetPropertyStatesLifeSafetyModeExactly interface {
	BACnetPropertyStatesLifeSafetyMode
	isBACnetPropertyStatesLifeSafetyMode() bool
}

// _BACnetPropertyStatesLifeSafetyMode is the data-structure of this message
type _BACnetPropertyStatesLifeSafetyMode struct {
	*_BACnetPropertyStates
	LifeSafetyMode BACnetLifeSafetyModeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLifeSafetyMode) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLifeSafetyMode) GetLifeSafetyMode() BACnetLifeSafetyModeTagged {
	return m.LifeSafetyMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLifeSafetyMode factory function for _BACnetPropertyStatesLifeSafetyMode
func NewBACnetPropertyStatesLifeSafetyMode(lifeSafetyMode BACnetLifeSafetyModeTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLifeSafetyMode {
	_result := &_BACnetPropertyStatesLifeSafetyMode{
		LifeSafetyMode:        lifeSafetyMode,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLifeSafetyMode(structType interface{}) BACnetPropertyStatesLifeSafetyMode {
	if casted, ok := structType.(BACnetPropertyStatesLifeSafetyMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLifeSafetyMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetTypeName() string {
	return "BACnetPropertyStatesLifeSafetyMode"
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lifeSafetyMode)
	lengthInBits += m.LifeSafetyMode.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLifeSafetyModeParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLifeSafetyMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLifeSafetyMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLifeSafetyMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lifeSafetyMode)
	if pullErr := readBuffer.PullContext("lifeSafetyMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lifeSafetyMode")
	}
	_lifeSafetyMode, _lifeSafetyModeErr := BACnetLifeSafetyModeTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _lifeSafetyModeErr != nil {
		return nil, errors.Wrap(_lifeSafetyModeErr, "Error parsing 'lifeSafetyMode' field of BACnetPropertyStatesLifeSafetyMode")
	}
	lifeSafetyMode := _lifeSafetyMode.(BACnetLifeSafetyModeTagged)
	if closeErr := readBuffer.CloseContext("lifeSafetyMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lifeSafetyMode")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLifeSafetyMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLifeSafetyMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLifeSafetyMode{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		LifeSafetyMode:        lifeSafetyMode,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLifeSafetyMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLifeSafetyMode) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLifeSafetyMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLifeSafetyMode")
		}

		// Simple Field (lifeSafetyMode)
		if pushErr := writeBuffer.PushContext("lifeSafetyMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lifeSafetyMode")
		}
		_lifeSafetyModeErr := writeBuffer.WriteSerializable(m.GetLifeSafetyMode())
		if popErr := writeBuffer.PopContext("lifeSafetyMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lifeSafetyMode")
		}
		if _lifeSafetyModeErr != nil {
			return errors.Wrap(_lifeSafetyModeErr, "Error serializing 'lifeSafetyMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLifeSafetyMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLifeSafetyMode")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLifeSafetyMode) isBACnetPropertyStatesLifeSafetyMode() bool {
	return true
}

func (m *_BACnetPropertyStatesLifeSafetyMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
