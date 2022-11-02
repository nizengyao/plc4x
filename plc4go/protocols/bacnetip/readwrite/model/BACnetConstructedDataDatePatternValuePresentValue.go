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

// BACnetConstructedDataDatePatternValuePresentValue is the corresponding interface of BACnetConstructedDataDatePatternValuePresentValue
type BACnetConstructedDataDatePatternValuePresentValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagDate
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDate
}

// BACnetConstructedDataDatePatternValuePresentValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDatePatternValuePresentValue.
// This is useful for switch cases.
type BACnetConstructedDataDatePatternValuePresentValueExactly interface {
	BACnetConstructedDataDatePatternValuePresentValue
	isBACnetConstructedDataDatePatternValuePresentValue() bool
}

// _BACnetConstructedDataDatePatternValuePresentValue is the data-structure of this message
type _BACnetConstructedDataDatePatternValuePresentValue struct {
	*_BACnetConstructedData
	PresentValue BACnetApplicationTagDate
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATEPATTERN_VALUE
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDatePatternValuePresentValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetPresentValue() BACnetApplicationTagDate {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetActualValue() BACnetApplicationTagDate {
	return CastBACnetApplicationTagDate(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDatePatternValuePresentValue factory function for _BACnetConstructedDataDatePatternValuePresentValue
func NewBACnetConstructedDataDatePatternValuePresentValue(presentValue BACnetApplicationTagDate, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDatePatternValuePresentValue {
	_result := &_BACnetConstructedDataDatePatternValuePresentValue{
		PresentValue:           presentValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDatePatternValuePresentValue(structType interface{}) BACnetConstructedDataDatePatternValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataDatePatternValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDatePatternValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataDatePatternValuePresentValue"
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDatePatternValuePresentValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDatePatternValuePresentValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDatePatternValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDatePatternValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (presentValue)
	if pullErr := readBuffer.PullContext("presentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for presentValue")
	}
	_presentValue, _presentValueErr := BACnetApplicationTagParse(readBuffer)
	if _presentValueErr != nil {
		return nil, errors.Wrap(_presentValueErr, "Error parsing 'presentValue' field of BACnetConstructedDataDatePatternValuePresentValue")
	}
	presentValue := _presentValue.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("presentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for presentValue")
	}

	// Virtual field
	_actualValue := presentValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDatePatternValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDatePatternValuePresentValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDatePatternValuePresentValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PresentValue: presentValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDatePatternValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDatePatternValuePresentValue")
		}

		// Simple Field (presentValue)
		if pushErr := writeBuffer.PushContext("presentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for presentValue")
		}
		_presentValueErr := writeBuffer.WriteSerializable(m.GetPresentValue())
		if popErr := writeBuffer.PopContext("presentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for presentValue")
		}
		if _presentValueErr != nil {
			return errors.Wrap(_presentValueErr, "Error serializing 'presentValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDatePatternValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDatePatternValuePresentValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) isBACnetConstructedDataDatePatternValuePresentValue() bool {
	return true
}

func (m *_BACnetConstructedDataDatePatternValuePresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
