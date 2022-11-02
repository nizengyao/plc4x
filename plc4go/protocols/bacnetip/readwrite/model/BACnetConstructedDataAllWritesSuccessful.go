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

// BACnetConstructedDataAllWritesSuccessful is the corresponding interface of BACnetConstructedDataAllWritesSuccessful
type BACnetConstructedDataAllWritesSuccessful interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAllWritesSuccessful returns AllWritesSuccessful (property field)
	GetAllWritesSuccessful() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataAllWritesSuccessfulExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAllWritesSuccessful.
// This is useful for switch cases.
type BACnetConstructedDataAllWritesSuccessfulExactly interface {
	BACnetConstructedDataAllWritesSuccessful
	isBACnetConstructedDataAllWritesSuccessful() bool
}

// _BACnetConstructedDataAllWritesSuccessful is the data-structure of this message
type _BACnetConstructedDataAllWritesSuccessful struct {
	*_BACnetConstructedData
	AllWritesSuccessful BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAllWritesSuccessful) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAllWritesSuccessful) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL_WRITES_SUCCESSFUL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAllWritesSuccessful) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAllWritesSuccessful) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAllWritesSuccessful) GetAllWritesSuccessful() BACnetApplicationTagBoolean {
	return m.AllWritesSuccessful
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAllWritesSuccessful) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetAllWritesSuccessful())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAllWritesSuccessful factory function for _BACnetConstructedDataAllWritesSuccessful
func NewBACnetConstructedDataAllWritesSuccessful(allWritesSuccessful BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAllWritesSuccessful {
	_result := &_BACnetConstructedDataAllWritesSuccessful{
		AllWritesSuccessful:    allWritesSuccessful,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAllWritesSuccessful(structType interface{}) BACnetConstructedDataAllWritesSuccessful {
	if casted, ok := structType.(BACnetConstructedDataAllWritesSuccessful); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAllWritesSuccessful); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAllWritesSuccessful) GetTypeName() string {
	return "BACnetConstructedDataAllWritesSuccessful"
}

func (m *_BACnetConstructedDataAllWritesSuccessful) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAllWritesSuccessful) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (allWritesSuccessful)
	lengthInBits += m.AllWritesSuccessful.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAllWritesSuccessful) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAllWritesSuccessfulParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAllWritesSuccessful, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAllWritesSuccessful"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAllWritesSuccessful")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (allWritesSuccessful)
	if pullErr := readBuffer.PullContext("allWritesSuccessful"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for allWritesSuccessful")
	}
	_allWritesSuccessful, _allWritesSuccessfulErr := BACnetApplicationTagParse(readBuffer)
	if _allWritesSuccessfulErr != nil {
		return nil, errors.Wrap(_allWritesSuccessfulErr, "Error parsing 'allWritesSuccessful' field of BACnetConstructedDataAllWritesSuccessful")
	}
	allWritesSuccessful := _allWritesSuccessful.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("allWritesSuccessful"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for allWritesSuccessful")
	}

	// Virtual field
	_actualValue := allWritesSuccessful
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAllWritesSuccessful"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAllWritesSuccessful")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAllWritesSuccessful{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AllWritesSuccessful: allWritesSuccessful,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAllWritesSuccessful) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAllWritesSuccessful) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAllWritesSuccessful"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAllWritesSuccessful")
		}

		// Simple Field (allWritesSuccessful)
		if pushErr := writeBuffer.PushContext("allWritesSuccessful"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for allWritesSuccessful")
		}
		_allWritesSuccessfulErr := writeBuffer.WriteSerializable(m.GetAllWritesSuccessful())
		if popErr := writeBuffer.PopContext("allWritesSuccessful"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for allWritesSuccessful")
		}
		if _allWritesSuccessfulErr != nil {
			return errors.Wrap(_allWritesSuccessfulErr, "Error serializing 'allWritesSuccessful' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAllWritesSuccessful"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAllWritesSuccessful")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAllWritesSuccessful) isBACnetConstructedDataAllWritesSuccessful() bool {
	return true
}

func (m *_BACnetConstructedDataAllWritesSuccessful) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
