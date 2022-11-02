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

// BACnetConstructedDataDefaultRampRate is the corresponding interface of BACnetConstructedDataDefaultRampRate
type BACnetConstructedDataDefaultRampRate interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDefaultRampRate returns DefaultRampRate (property field)
	GetDefaultRampRate() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataDefaultRampRateExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDefaultRampRate.
// This is useful for switch cases.
type BACnetConstructedDataDefaultRampRateExactly interface {
	BACnetConstructedDataDefaultRampRate
	isBACnetConstructedDataDefaultRampRate() bool
}

// _BACnetConstructedDataDefaultRampRate is the data-structure of this message
type _BACnetConstructedDataDefaultRampRate struct {
	*_BACnetConstructedData
	DefaultRampRate BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDefaultRampRate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDefaultRampRate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEFAULT_RAMP_RATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDefaultRampRate) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDefaultRampRate) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultRampRate) GetDefaultRampRate() BACnetApplicationTagReal {
	return m.DefaultRampRate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultRampRate) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetDefaultRampRate())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDefaultRampRate factory function for _BACnetConstructedDataDefaultRampRate
func NewBACnetConstructedDataDefaultRampRate(defaultRampRate BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDefaultRampRate {
	_result := &_BACnetConstructedDataDefaultRampRate{
		DefaultRampRate:        defaultRampRate,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDefaultRampRate(structType interface{}) BACnetConstructedDataDefaultRampRate {
	if casted, ok := structType.(BACnetConstructedDataDefaultRampRate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDefaultRampRate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDefaultRampRate) GetTypeName() string {
	return "BACnetConstructedDataDefaultRampRate"
}

func (m *_BACnetConstructedDataDefaultRampRate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDefaultRampRate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (defaultRampRate)
	lengthInBits += m.DefaultRampRate.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDefaultRampRate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDefaultRampRateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDefaultRampRate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDefaultRampRate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDefaultRampRate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (defaultRampRate)
	if pullErr := readBuffer.PullContext("defaultRampRate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for defaultRampRate")
	}
	_defaultRampRate, _defaultRampRateErr := BACnetApplicationTagParse(readBuffer)
	if _defaultRampRateErr != nil {
		return nil, errors.Wrap(_defaultRampRateErr, "Error parsing 'defaultRampRate' field of BACnetConstructedDataDefaultRampRate")
	}
	defaultRampRate := _defaultRampRate.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("defaultRampRate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for defaultRampRate")
	}

	// Virtual field
	_actualValue := defaultRampRate
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDefaultRampRate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDefaultRampRate")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDefaultRampRate{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DefaultRampRate: defaultRampRate,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDefaultRampRate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDefaultRampRate) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDefaultRampRate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDefaultRampRate")
		}

		// Simple Field (defaultRampRate)
		if pushErr := writeBuffer.PushContext("defaultRampRate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for defaultRampRate")
		}
		_defaultRampRateErr := writeBuffer.WriteSerializable(m.GetDefaultRampRate())
		if popErr := writeBuffer.PopContext("defaultRampRate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for defaultRampRate")
		}
		if _defaultRampRateErr != nil {
			return errors.Wrap(_defaultRampRateErr, "Error serializing 'defaultRampRate' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDefaultRampRate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDefaultRampRate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDefaultRampRate) isBACnetConstructedDataDefaultRampRate() bool {
	return true
}

func (m *_BACnetConstructedDataDefaultRampRate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
