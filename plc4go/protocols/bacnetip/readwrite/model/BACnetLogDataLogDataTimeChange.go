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

// BACnetLogDataLogDataTimeChange is the corresponding interface of BACnetLogDataLogDataTimeChange
type BACnetLogDataLogDataTimeChange interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogData
	// GetTimeChange returns TimeChange (property field)
	GetTimeChange() BACnetContextTagReal
}

// BACnetLogDataLogDataTimeChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataTimeChange.
// This is useful for switch cases.
type BACnetLogDataLogDataTimeChangeExactly interface {
	BACnetLogDataLogDataTimeChange
	isBACnetLogDataLogDataTimeChange() bool
}

// _BACnetLogDataLogDataTimeChange is the data-structure of this message
type _BACnetLogDataLogDataTimeChange struct {
	*_BACnetLogData
	TimeChange BACnetContextTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataTimeChange) InitializeParent(parent BACnetLogData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogDataLogDataTimeChange) GetParent() BACnetLogData {
	return m._BACnetLogData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataTimeChange) GetTimeChange() BACnetContextTagReal {
	return m.TimeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataTimeChange factory function for _BACnetLogDataLogDataTimeChange
func NewBACnetLogDataLogDataTimeChange(timeChange BACnetContextTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogDataLogDataTimeChange {
	_result := &_BACnetLogDataLogDataTimeChange{
		TimeChange:     timeChange,
		_BACnetLogData: NewBACnetLogData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogData._BACnetLogDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataTimeChange(structType interface{}) BACnetLogDataLogDataTimeChange {
	if casted, ok := structType.(BACnetLogDataLogDataTimeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataTimeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataTimeChange) GetTypeName() string {
	return "BACnetLogDataLogDataTimeChange"
}

func (m *_BACnetLogDataLogDataTimeChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogDataLogDataTimeChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeChange)
	lengthInBits += m.TimeChange.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogDataLogDataTimeChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogDataLogDataTimeChangeParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogDataLogDataTimeChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataTimeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataTimeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeChange)
	if pullErr := readBuffer.PullContext("timeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeChange")
	}
	_timeChange, _timeChangeErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_REAL))
	if _timeChangeErr != nil {
		return nil, errors.Wrap(_timeChangeErr, "Error parsing 'timeChange' field of BACnetLogDataLogDataTimeChange")
	}
	timeChange := _timeChange.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("timeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeChange")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataTimeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataTimeChange")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogDataTimeChange{
		_BACnetLogData: &_BACnetLogData{
			TagNumber: tagNumber,
		},
		TimeChange: timeChange,
	}
	_child._BACnetLogData._BACnetLogDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogDataTimeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataTimeChange) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataTimeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataTimeChange")
		}

		// Simple Field (timeChange)
		if pushErr := writeBuffer.PushContext("timeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeChange")
		}
		_timeChangeErr := writeBuffer.WriteSerializable(m.GetTimeChange())
		if popErr := writeBuffer.PopContext("timeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeChange")
		}
		if _timeChangeErr != nil {
			return errors.Wrap(_timeChangeErr, "Error serializing 'timeChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataTimeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataTimeChange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataTimeChange) isBACnetLogDataLogDataTimeChange() bool {
	return true
}

func (m *_BACnetLogDataLogDataTimeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
