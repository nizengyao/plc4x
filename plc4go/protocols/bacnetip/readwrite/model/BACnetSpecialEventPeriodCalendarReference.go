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

// BACnetSpecialEventPeriodCalendarReference is the corresponding interface of BACnetSpecialEventPeriodCalendarReference
type BACnetSpecialEventPeriodCalendarReference interface {
	utils.LengthAware
	utils.Serializable
	BACnetSpecialEventPeriod
	// GetCalendarReference returns CalendarReference (property field)
	GetCalendarReference() BACnetContextTagObjectIdentifier
}

// BACnetSpecialEventPeriodCalendarReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetSpecialEventPeriodCalendarReference.
// This is useful for switch cases.
type BACnetSpecialEventPeriodCalendarReferenceExactly interface {
	BACnetSpecialEventPeriodCalendarReference
	isBACnetSpecialEventPeriodCalendarReference() bool
}

// _BACnetSpecialEventPeriodCalendarReference is the data-structure of this message
type _BACnetSpecialEventPeriodCalendarReference struct {
	*_BACnetSpecialEventPeriod
	CalendarReference BACnetContextTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetSpecialEventPeriodCalendarReference) InitializeParent(parent BACnetSpecialEventPeriod, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetSpecialEventPeriodCalendarReference) GetParent() BACnetSpecialEventPeriod {
	return m._BACnetSpecialEventPeriod
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSpecialEventPeriodCalendarReference) GetCalendarReference() BACnetContextTagObjectIdentifier {
	return m.CalendarReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSpecialEventPeriodCalendarReference factory function for _BACnetSpecialEventPeriodCalendarReference
func NewBACnetSpecialEventPeriodCalendarReference(calendarReference BACnetContextTagObjectIdentifier, peekedTagHeader BACnetTagHeader) *_BACnetSpecialEventPeriodCalendarReference {
	_result := &_BACnetSpecialEventPeriodCalendarReference{
		CalendarReference:         calendarReference,
		_BACnetSpecialEventPeriod: NewBACnetSpecialEventPeriod(peekedTagHeader),
	}
	_result._BACnetSpecialEventPeriod._BACnetSpecialEventPeriodChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetSpecialEventPeriodCalendarReference(structType interface{}) BACnetSpecialEventPeriodCalendarReference {
	if casted, ok := structType.(BACnetSpecialEventPeriodCalendarReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSpecialEventPeriodCalendarReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSpecialEventPeriodCalendarReference) GetTypeName() string {
	return "BACnetSpecialEventPeriodCalendarReference"
}

func (m *_BACnetSpecialEventPeriodCalendarReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetSpecialEventPeriodCalendarReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (calendarReference)
	lengthInBits += m.CalendarReference.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetSpecialEventPeriodCalendarReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetSpecialEventPeriodCalendarReferenceParse(readBuffer utils.ReadBuffer) (BACnetSpecialEventPeriodCalendarReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSpecialEventPeriodCalendarReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEventPeriodCalendarReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (calendarReference)
	if pullErr := readBuffer.PullContext("calendarReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for calendarReference")
	}
	_calendarReference, _calendarReferenceErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _calendarReferenceErr != nil {
		return nil, errors.Wrap(_calendarReferenceErr, "Error parsing 'calendarReference' field of BACnetSpecialEventPeriodCalendarReference")
	}
	calendarReference := _calendarReference.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("calendarReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for calendarReference")
	}

	if closeErr := readBuffer.CloseContext("BACnetSpecialEventPeriodCalendarReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEventPeriodCalendarReference")
	}

	// Create a partially initialized instance
	_child := &_BACnetSpecialEventPeriodCalendarReference{
		_BACnetSpecialEventPeriod: &_BACnetSpecialEventPeriod{},
		CalendarReference:         calendarReference,
	}
	_child._BACnetSpecialEventPeriod._BACnetSpecialEventPeriodChildRequirements = _child
	return _child, nil
}

func (m *_BACnetSpecialEventPeriodCalendarReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSpecialEventPeriodCalendarReference) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetSpecialEventPeriodCalendarReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEventPeriodCalendarReference")
		}

		// Simple Field (calendarReference)
		if pushErr := writeBuffer.PushContext("calendarReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for calendarReference")
		}
		_calendarReferenceErr := writeBuffer.WriteSerializable(m.GetCalendarReference())
		if popErr := writeBuffer.PopContext("calendarReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for calendarReference")
		}
		if _calendarReferenceErr != nil {
			return errors.Wrap(_calendarReferenceErr, "Error serializing 'calendarReference' field")
		}

		if popErr := writeBuffer.PopContext("BACnetSpecialEventPeriodCalendarReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetSpecialEventPeriodCalendarReference")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetSpecialEventPeriodCalendarReference) isBACnetSpecialEventPeriodCalendarReference() bool {
	return true
}

func (m *_BACnetSpecialEventPeriodCalendarReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
