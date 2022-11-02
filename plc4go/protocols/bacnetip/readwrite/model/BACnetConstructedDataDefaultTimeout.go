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

// BACnetConstructedDataDefaultTimeout is the corresponding interface of BACnetConstructedDataDefaultTimeout
type BACnetConstructedDataDefaultTimeout interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDefaultTimeout returns DefaultTimeout (property field)
	GetDefaultTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataDefaultTimeoutExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDefaultTimeout.
// This is useful for switch cases.
type BACnetConstructedDataDefaultTimeoutExactly interface {
	BACnetConstructedDataDefaultTimeout
	isBACnetConstructedDataDefaultTimeout() bool
}

// _BACnetConstructedDataDefaultTimeout is the data-structure of this message
type _BACnetConstructedDataDefaultTimeout struct {
	*_BACnetConstructedData
	DefaultTimeout BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDefaultTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDefaultTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEFAULT_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDefaultTimeout) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDefaultTimeout) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultTimeout) GetDefaultTimeout() BACnetApplicationTagUnsignedInteger {
	return m.DefaultTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetDefaultTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDefaultTimeout factory function for _BACnetConstructedDataDefaultTimeout
func NewBACnetConstructedDataDefaultTimeout(defaultTimeout BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDefaultTimeout {
	_result := &_BACnetConstructedDataDefaultTimeout{
		DefaultTimeout:         defaultTimeout,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDefaultTimeout(structType interface{}) BACnetConstructedDataDefaultTimeout {
	if casted, ok := structType.(BACnetConstructedDataDefaultTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDefaultTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDefaultTimeout) GetTypeName() string {
	return "BACnetConstructedDataDefaultTimeout"
}

func (m *_BACnetConstructedDataDefaultTimeout) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDefaultTimeout) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (defaultTimeout)
	lengthInBits += m.DefaultTimeout.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDefaultTimeout) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDefaultTimeoutParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDefaultTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDefaultTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDefaultTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (defaultTimeout)
	if pullErr := readBuffer.PullContext("defaultTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for defaultTimeout")
	}
	_defaultTimeout, _defaultTimeoutErr := BACnetApplicationTagParse(readBuffer)
	if _defaultTimeoutErr != nil {
		return nil, errors.Wrap(_defaultTimeoutErr, "Error parsing 'defaultTimeout' field of BACnetConstructedDataDefaultTimeout")
	}
	defaultTimeout := _defaultTimeout.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("defaultTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for defaultTimeout")
	}

	// Virtual field
	_actualValue := defaultTimeout
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDefaultTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDefaultTimeout")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDefaultTimeout{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DefaultTimeout: defaultTimeout,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDefaultTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDefaultTimeout) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDefaultTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDefaultTimeout")
		}

		// Simple Field (defaultTimeout)
		if pushErr := writeBuffer.PushContext("defaultTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for defaultTimeout")
		}
		_defaultTimeoutErr := writeBuffer.WriteSerializable(m.GetDefaultTimeout())
		if popErr := writeBuffer.PopContext("defaultTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for defaultTimeout")
		}
		if _defaultTimeoutErr != nil {
			return errors.Wrap(_defaultTimeoutErr, "Error serializing 'defaultTimeout' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDefaultTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDefaultTimeout")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDefaultTimeout) isBACnetConstructedDataDefaultTimeout() bool {
	return true
}

func (m *_BACnetConstructedDataDefaultTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
