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

// BACnetConstructedDataMinimumOnTime is the corresponding interface of BACnetConstructedDataMinimumOnTime
type BACnetConstructedDataMinimumOnTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMinimumOnTime returns MinimumOnTime (property field)
	GetMinimumOnTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataMinimumOnTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMinimumOnTime.
// This is useful for switch cases.
type BACnetConstructedDataMinimumOnTimeExactly interface {
	BACnetConstructedDataMinimumOnTime
	isBACnetConstructedDataMinimumOnTime() bool
}

// _BACnetConstructedDataMinimumOnTime is the data-structure of this message
type _BACnetConstructedDataMinimumOnTime struct {
	*_BACnetConstructedData
	MinimumOnTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMinimumOnTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMinimumOnTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_ON_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMinimumOnTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMinimumOnTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumOnTime) GetMinimumOnTime() BACnetApplicationTagUnsignedInteger {
	return m.MinimumOnTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumOnTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetMinimumOnTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinimumOnTime factory function for _BACnetConstructedDataMinimumOnTime
func NewBACnetConstructedDataMinimumOnTime(minimumOnTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMinimumOnTime {
	_result := &_BACnetConstructedDataMinimumOnTime{
		MinimumOnTime:          minimumOnTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMinimumOnTime(structType interface{}) BACnetConstructedDataMinimumOnTime {
	if casted, ok := structType.(BACnetConstructedDataMinimumOnTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumOnTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMinimumOnTime) GetTypeName() string {
	return "BACnetConstructedDataMinimumOnTime"
}

func (m *_BACnetConstructedDataMinimumOnTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMinimumOnTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (minimumOnTime)
	lengthInBits += m.MinimumOnTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMinimumOnTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMinimumOnTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMinimumOnTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumOnTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumOnTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minimumOnTime)
	if pullErr := readBuffer.PullContext("minimumOnTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minimumOnTime")
	}
	_minimumOnTime, _minimumOnTimeErr := BACnetApplicationTagParse(readBuffer)
	if _minimumOnTimeErr != nil {
		return nil, errors.Wrap(_minimumOnTimeErr, "Error parsing 'minimumOnTime' field of BACnetConstructedDataMinimumOnTime")
	}
	minimumOnTime := _minimumOnTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("minimumOnTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minimumOnTime")
	}

	// Virtual field
	_actualValue := minimumOnTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumOnTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumOnTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMinimumOnTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MinimumOnTime: minimumOnTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMinimumOnTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMinimumOnTime) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumOnTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumOnTime")
		}

		// Simple Field (minimumOnTime)
		if pushErr := writeBuffer.PushContext("minimumOnTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minimumOnTime")
		}
		_minimumOnTimeErr := writeBuffer.WriteSerializable(m.GetMinimumOnTime())
		if popErr := writeBuffer.PopContext("minimumOnTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minimumOnTime")
		}
		if _minimumOnTimeErr != nil {
			return errors.Wrap(_minimumOnTimeErr, "Error serializing 'minimumOnTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumOnTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumOnTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMinimumOnTime) isBACnetConstructedDataMinimumOnTime() bool {
	return true
}

func (m *_BACnetConstructedDataMinimumOnTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
