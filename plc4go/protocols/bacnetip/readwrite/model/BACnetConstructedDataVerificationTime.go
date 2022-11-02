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

// BACnetConstructedDataVerificationTime is the corresponding interface of BACnetConstructedDataVerificationTime
type BACnetConstructedDataVerificationTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetVerificationTime returns VerificationTime (property field)
	GetVerificationTime() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataVerificationTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataVerificationTime.
// This is useful for switch cases.
type BACnetConstructedDataVerificationTimeExactly interface {
	BACnetConstructedDataVerificationTime
	isBACnetConstructedDataVerificationTime() bool
}

// _BACnetConstructedDataVerificationTime is the data-structure of this message
type _BACnetConstructedDataVerificationTime struct {
	*_BACnetConstructedData
	VerificationTime BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVerificationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVerificationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VERIFICATION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVerificationTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataVerificationTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVerificationTime) GetVerificationTime() BACnetApplicationTagSignedInteger {
	return m.VerificationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataVerificationTime) GetActualValue() BACnetApplicationTagSignedInteger {
	return CastBACnetApplicationTagSignedInteger(m.GetVerificationTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVerificationTime factory function for _BACnetConstructedDataVerificationTime
func NewBACnetConstructedDataVerificationTime(verificationTime BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVerificationTime {
	_result := &_BACnetConstructedDataVerificationTime{
		VerificationTime:       verificationTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVerificationTime(structType interface{}) BACnetConstructedDataVerificationTime {
	if casted, ok := structType.(BACnetConstructedDataVerificationTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVerificationTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVerificationTime) GetTypeName() string {
	return "BACnetConstructedDataVerificationTime"
}

func (m *_BACnetConstructedDataVerificationTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataVerificationTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (verificationTime)
	lengthInBits += m.VerificationTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataVerificationTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataVerificationTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVerificationTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVerificationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVerificationTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (verificationTime)
	if pullErr := readBuffer.PullContext("verificationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for verificationTime")
	}
	_verificationTime, _verificationTimeErr := BACnetApplicationTagParse(readBuffer)
	if _verificationTimeErr != nil {
		return nil, errors.Wrap(_verificationTimeErr, "Error parsing 'verificationTime' field of BACnetConstructedDataVerificationTime")
	}
	verificationTime := _verificationTime.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("verificationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for verificationTime")
	}

	// Virtual field
	_actualValue := verificationTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVerificationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVerificationTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataVerificationTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		VerificationTime: verificationTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataVerificationTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataVerificationTime) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVerificationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVerificationTime")
		}

		// Simple Field (verificationTime)
		if pushErr := writeBuffer.PushContext("verificationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for verificationTime")
		}
		_verificationTimeErr := writeBuffer.WriteSerializable(m.GetVerificationTime())
		if popErr := writeBuffer.PopContext("verificationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for verificationTime")
		}
		if _verificationTimeErr != nil {
			return errors.Wrap(_verificationTimeErr, "Error serializing 'verificationTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVerificationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVerificationTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVerificationTime) isBACnetConstructedDataVerificationTime() bool {
	return true
}

func (m *_BACnetConstructedDataVerificationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
