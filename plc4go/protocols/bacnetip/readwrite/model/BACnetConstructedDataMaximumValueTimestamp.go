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

// BACnetConstructedDataMaximumValueTimestamp is the corresponding interface of BACnetConstructedDataMaximumValueTimestamp
type BACnetConstructedDataMaximumValueTimestamp interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaximumValueTimestamp returns MaximumValueTimestamp (property field)
	GetMaximumValueTimestamp() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataMaximumValueTimestampExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMaximumValueTimestamp.
// This is useful for switch cases.
type BACnetConstructedDataMaximumValueTimestampExactly interface {
	BACnetConstructedDataMaximumValueTimestamp
	isBACnetConstructedDataMaximumValueTimestamp() bool
}

// _BACnetConstructedDataMaximumValueTimestamp is the data-structure of this message
type _BACnetConstructedDataMaximumValueTimestamp struct {
	*_BACnetConstructedData
	MaximumValueTimestamp BACnetDateTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAXIMUM_VALUE_TIMESTAMP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaximumValueTimestamp) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetMaximumValueTimestamp() BACnetDateTime {
	return m.MaximumValueTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetActualValue() BACnetDateTime {
	return CastBACnetDateTime(m.GetMaximumValueTimestamp())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaximumValueTimestamp factory function for _BACnetConstructedDataMaximumValueTimestamp
func NewBACnetConstructedDataMaximumValueTimestamp(maximumValueTimestamp BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaximumValueTimestamp {
	_result := &_BACnetConstructedDataMaximumValueTimestamp{
		MaximumValueTimestamp:  maximumValueTimestamp,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaximumValueTimestamp(structType interface{}) BACnetConstructedDataMaximumValueTimestamp {
	if casted, ok := structType.(BACnetConstructedDataMaximumValueTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaximumValueTimestamp); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetTypeName() string {
	return "BACnetConstructedDataMaximumValueTimestamp"
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maximumValueTimestamp)
	lengthInBits += m.MaximumValueTimestamp.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaximumValueTimestampParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaximumValueTimestamp, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaximumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaximumValueTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maximumValueTimestamp)
	if pullErr := readBuffer.PullContext("maximumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maximumValueTimestamp")
	}
	_maximumValueTimestamp, _maximumValueTimestampErr := BACnetDateTimeParse(readBuffer)
	if _maximumValueTimestampErr != nil {
		return nil, errors.Wrap(_maximumValueTimestampErr, "Error parsing 'maximumValueTimestamp' field of BACnetConstructedDataMaximumValueTimestamp")
	}
	maximumValueTimestamp := _maximumValueTimestamp.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("maximumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maximumValueTimestamp")
	}

	// Virtual field
	_actualValue := maximumValueTimestamp
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaximumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaximumValueTimestamp")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMaximumValueTimestamp{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaximumValueTimestamp: maximumValueTimestamp,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaximumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaximumValueTimestamp")
		}

		// Simple Field (maximumValueTimestamp)
		if pushErr := writeBuffer.PushContext("maximumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maximumValueTimestamp")
		}
		_maximumValueTimestampErr := writeBuffer.WriteSerializable(m.GetMaximumValueTimestamp())
		if popErr := writeBuffer.PopContext("maximumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maximumValueTimestamp")
		}
		if _maximumValueTimestampErr != nil {
			return errors.Wrap(_maximumValueTimestampErr, "Error serializing 'maximumValueTimestamp' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaximumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaximumValueTimestamp")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) isBACnetConstructedDataMaximumValueTimestamp() bool {
	return true
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
