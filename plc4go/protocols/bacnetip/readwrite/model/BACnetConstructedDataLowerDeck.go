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

// BACnetConstructedDataLowerDeck is the corresponding interface of BACnetConstructedDataLowerDeck
type BACnetConstructedDataLowerDeck interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLowerDeck returns LowerDeck (property field)
	GetLowerDeck() BACnetApplicationTagObjectIdentifier
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagObjectIdentifier
}

// BACnetConstructedDataLowerDeckExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLowerDeck.
// This is useful for switch cases.
type BACnetConstructedDataLowerDeckExactly interface {
	BACnetConstructedDataLowerDeck
	isBACnetConstructedDataLowerDeck() bool
}

// _BACnetConstructedDataLowerDeck is the data-structure of this message
type _BACnetConstructedDataLowerDeck struct {
	*_BACnetConstructedData
	LowerDeck BACnetApplicationTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLowerDeck) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLowerDeck) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOWER_DECK
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLowerDeck) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLowerDeck) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLowerDeck) GetLowerDeck() BACnetApplicationTagObjectIdentifier {
	return m.LowerDeck
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLowerDeck) GetActualValue() BACnetApplicationTagObjectIdentifier {
	return CastBACnetApplicationTagObjectIdentifier(m.GetLowerDeck())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLowerDeck factory function for _BACnetConstructedDataLowerDeck
func NewBACnetConstructedDataLowerDeck(lowerDeck BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLowerDeck {
	_result := &_BACnetConstructedDataLowerDeck{
		LowerDeck:              lowerDeck,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLowerDeck(structType interface{}) BACnetConstructedDataLowerDeck {
	if casted, ok := structType.(BACnetConstructedDataLowerDeck); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLowerDeck); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLowerDeck) GetTypeName() string {
	return "BACnetConstructedDataLowerDeck"
}

func (m *_BACnetConstructedDataLowerDeck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLowerDeck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lowerDeck)
	lengthInBits += m.LowerDeck.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLowerDeck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLowerDeckParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLowerDeck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLowerDeck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLowerDeck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lowerDeck)
	if pullErr := readBuffer.PullContext("lowerDeck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lowerDeck")
	}
	_lowerDeck, _lowerDeckErr := BACnetApplicationTagParse(readBuffer)
	if _lowerDeckErr != nil {
		return nil, errors.Wrap(_lowerDeckErr, "Error parsing 'lowerDeck' field of BACnetConstructedDataLowerDeck")
	}
	lowerDeck := _lowerDeck.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("lowerDeck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lowerDeck")
	}

	// Virtual field
	_actualValue := lowerDeck
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLowerDeck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLowerDeck")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLowerDeck{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LowerDeck: lowerDeck,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLowerDeck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLowerDeck) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLowerDeck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLowerDeck")
		}

		// Simple Field (lowerDeck)
		if pushErr := writeBuffer.PushContext("lowerDeck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lowerDeck")
		}
		_lowerDeckErr := writeBuffer.WriteSerializable(m.GetLowerDeck())
		if popErr := writeBuffer.PopContext("lowerDeck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lowerDeck")
		}
		if _lowerDeckErr != nil {
			return errors.Wrap(_lowerDeckErr, "Error serializing 'lowerDeck' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLowerDeck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLowerDeck")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLowerDeck) isBACnetConstructedDataLowerDeck() bool {
	return true
}

func (m *_BACnetConstructedDataLowerDeck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
