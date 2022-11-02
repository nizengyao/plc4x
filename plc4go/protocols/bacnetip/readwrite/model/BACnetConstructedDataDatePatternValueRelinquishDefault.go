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

// BACnetConstructedDataDatePatternValueRelinquishDefault is the corresponding interface of BACnetConstructedDataDatePatternValueRelinquishDefault
type BACnetConstructedDataDatePatternValueRelinquishDefault interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetApplicationTagDate
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDate
}

// BACnetConstructedDataDatePatternValueRelinquishDefaultExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDatePatternValueRelinquishDefault.
// This is useful for switch cases.
type BACnetConstructedDataDatePatternValueRelinquishDefaultExactly interface {
	BACnetConstructedDataDatePatternValueRelinquishDefault
	isBACnetConstructedDataDatePatternValueRelinquishDefault() bool
}

// _BACnetConstructedDataDatePatternValueRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataDatePatternValueRelinquishDefault struct {
	*_BACnetConstructedData
	RelinquishDefault BACnetApplicationTagDate
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATEPATTERN_VALUE
}

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) GetRelinquishDefault() BACnetApplicationTagDate {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) GetActualValue() BACnetApplicationTagDate {
	return CastBACnetApplicationTagDate(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDatePatternValueRelinquishDefault factory function for _BACnetConstructedDataDatePatternValueRelinquishDefault
func NewBACnetConstructedDataDatePatternValueRelinquishDefault(relinquishDefault BACnetApplicationTagDate, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDatePatternValueRelinquishDefault {
	_result := &_BACnetConstructedDataDatePatternValueRelinquishDefault{
		RelinquishDefault:      relinquishDefault,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDatePatternValueRelinquishDefault(structType interface{}) BACnetConstructedDataDatePatternValueRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataDatePatternValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDatePatternValueRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataDatePatternValueRelinquishDefault"
}

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDatePatternValueRelinquishDefaultParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDatePatternValueRelinquishDefault, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDatePatternValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDatePatternValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (relinquishDefault)
	if pullErr := readBuffer.PullContext("relinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for relinquishDefault")
	}
	_relinquishDefault, _relinquishDefaultErr := BACnetApplicationTagParse(readBuffer)
	if _relinquishDefaultErr != nil {
		return nil, errors.Wrap(_relinquishDefaultErr, "Error parsing 'relinquishDefault' field of BACnetConstructedDataDatePatternValueRelinquishDefault")
	}
	relinquishDefault := _relinquishDefault.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("relinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for relinquishDefault")
	}

	// Virtual field
	_actualValue := relinquishDefault
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDatePatternValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDatePatternValueRelinquishDefault")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDatePatternValueRelinquishDefault{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		RelinquishDefault: relinquishDefault,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDatePatternValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDatePatternValueRelinquishDefault")
		}

		// Simple Field (relinquishDefault)
		if pushErr := writeBuffer.PushContext("relinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for relinquishDefault")
		}
		_relinquishDefaultErr := writeBuffer.WriteSerializable(m.GetRelinquishDefault())
		if popErr := writeBuffer.PopContext("relinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for relinquishDefault")
		}
		if _relinquishDefaultErr != nil {
			return errors.Wrap(_relinquishDefaultErr, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDatePatternValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDatePatternValueRelinquishDefault")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) isBACnetConstructedDataDatePatternValueRelinquishDefault() bool {
	return true
}

func (m *_BACnetConstructedDataDatePatternValueRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
