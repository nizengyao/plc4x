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

// BACnetConstructedDataPassbackTimeout is the corresponding interface of BACnetConstructedDataPassbackTimeout
type BACnetConstructedDataPassbackTimeout interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPassbackTimeout returns PassbackTimeout (property field)
	GetPassbackTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataPassbackTimeoutExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPassbackTimeout.
// This is useful for switch cases.
type BACnetConstructedDataPassbackTimeoutExactly interface {
	BACnetConstructedDataPassbackTimeout
	isBACnetConstructedDataPassbackTimeout() bool
}

// _BACnetConstructedDataPassbackTimeout is the data-structure of this message
type _BACnetConstructedDataPassbackTimeout struct {
	*_BACnetConstructedData
	PassbackTimeout BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPassbackTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPassbackTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PASSBACK_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPassbackTimeout) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPassbackTimeout) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPassbackTimeout) GetPassbackTimeout() BACnetApplicationTagUnsignedInteger {
	return m.PassbackTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPassbackTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetPassbackTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPassbackTimeout factory function for _BACnetConstructedDataPassbackTimeout
func NewBACnetConstructedDataPassbackTimeout(passbackTimeout BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPassbackTimeout {
	_result := &_BACnetConstructedDataPassbackTimeout{
		PassbackTimeout:        passbackTimeout,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPassbackTimeout(structType interface{}) BACnetConstructedDataPassbackTimeout {
	if casted, ok := structType.(BACnetConstructedDataPassbackTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPassbackTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPassbackTimeout) GetTypeName() string {
	return "BACnetConstructedDataPassbackTimeout"
}

func (m *_BACnetConstructedDataPassbackTimeout) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataPassbackTimeout) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (passbackTimeout)
	lengthInBits += m.PassbackTimeout.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPassbackTimeout) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPassbackTimeoutParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPassbackTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPassbackTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPassbackTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (passbackTimeout)
	if pullErr := readBuffer.PullContext("passbackTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for passbackTimeout")
	}
	_passbackTimeout, _passbackTimeoutErr := BACnetApplicationTagParse(readBuffer)
	if _passbackTimeoutErr != nil {
		return nil, errors.Wrap(_passbackTimeoutErr, "Error parsing 'passbackTimeout' field of BACnetConstructedDataPassbackTimeout")
	}
	passbackTimeout := _passbackTimeout.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("passbackTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for passbackTimeout")
	}

	// Virtual field
	_actualValue := passbackTimeout
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPassbackTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPassbackTimeout")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPassbackTimeout{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PassbackTimeout: passbackTimeout,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPassbackTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPassbackTimeout) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPassbackTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPassbackTimeout")
		}

		// Simple Field (passbackTimeout)
		if pushErr := writeBuffer.PushContext("passbackTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for passbackTimeout")
		}
		_passbackTimeoutErr := writeBuffer.WriteSerializable(m.GetPassbackTimeout())
		if popErr := writeBuffer.PopContext("passbackTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for passbackTimeout")
		}
		if _passbackTimeoutErr != nil {
			return errors.Wrap(_passbackTimeoutErr, "Error serializing 'passbackTimeout' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPassbackTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPassbackTimeout")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPassbackTimeout) isBACnetConstructedDataPassbackTimeout() bool {
	return true
}

func (m *_BACnetConstructedDataPassbackTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
