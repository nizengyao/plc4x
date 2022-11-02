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

// BACnetLogDataLogDataEntryFailure is the corresponding interface of BACnetLogDataLogDataEntryFailure
type BACnetLogDataLogDataEntryFailure interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogDataLogDataEntry
	// GetFailure returns Failure (property field)
	GetFailure() ErrorEnclosed
}

// BACnetLogDataLogDataEntryFailureExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataEntryFailure.
// This is useful for switch cases.
type BACnetLogDataLogDataEntryFailureExactly interface {
	BACnetLogDataLogDataEntryFailure
	isBACnetLogDataLogDataEntryFailure() bool
}

// _BACnetLogDataLogDataEntryFailure is the data-structure of this message
type _BACnetLogDataLogDataEntryFailure struct {
	*_BACnetLogDataLogDataEntry
	Failure ErrorEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataEntryFailure) InitializeParent(parent BACnetLogDataLogDataEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetLogDataLogDataEntryFailure) GetParent() BACnetLogDataLogDataEntry {
	return m._BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryFailure) GetFailure() ErrorEnclosed {
	return m.Failure
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryFailure factory function for _BACnetLogDataLogDataEntryFailure
func NewBACnetLogDataLogDataEntryFailure(failure ErrorEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetLogDataLogDataEntryFailure {
	_result := &_BACnetLogDataLogDataEntryFailure{
		Failure:                    failure,
		_BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryFailure(structType interface{}) BACnetLogDataLogDataEntryFailure {
	if casted, ok := structType.(BACnetLogDataLogDataEntryFailure); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryFailure); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryFailure) GetTypeName() string {
	return "BACnetLogDataLogDataEntryFailure"
}

func (m *_BACnetLogDataLogDataEntryFailure) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogDataLogDataEntryFailure) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (failure)
	lengthInBits += m.Failure.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryFailure) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogDataLogDataEntryFailureParse(readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntryFailure, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (failure)
	if pullErr := readBuffer.PullContext("failure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for failure")
	}
	_failure, _failureErr := ErrorEnclosedParse(readBuffer, uint8(uint8(7)))
	if _failureErr != nil {
		return nil, errors.Wrap(_failureErr, "Error parsing 'failure' field of BACnetLogDataLogDataEntryFailure")
	}
	failure := _failure.(ErrorEnclosed)
	if closeErr := readBuffer.CloseContext("failure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for failure")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryFailure")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogDataEntryFailure{
		_BACnetLogDataLogDataEntry: &_BACnetLogDataLogDataEntry{},
		Failure:                    failure,
	}
	_child._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogDataEntryFailure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryFailure) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryFailure")
		}

		// Simple Field (failure)
		if pushErr := writeBuffer.PushContext("failure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for failure")
		}
		_failureErr := writeBuffer.WriteSerializable(m.GetFailure())
		if popErr := writeBuffer.PopContext("failure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for failure")
		}
		if _failureErr != nil {
			return errors.Wrap(_failureErr, "Error serializing 'failure' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryFailure")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryFailure) isBACnetLogDataLogDataEntryFailure() bool {
	return true
}

func (m *_BACnetLogDataLogDataEntryFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
