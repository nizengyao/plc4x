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

// BACnetConstructedDataScaleFactor is the corresponding interface of BACnetConstructedDataScaleFactor
type BACnetConstructedDataScaleFactor interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetScaleFactor returns ScaleFactor (property field)
	GetScaleFactor() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataScaleFactorExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataScaleFactor.
// This is useful for switch cases.
type BACnetConstructedDataScaleFactorExactly interface {
	BACnetConstructedDataScaleFactor
	isBACnetConstructedDataScaleFactor() bool
}

// _BACnetConstructedDataScaleFactor is the data-structure of this message
type _BACnetConstructedDataScaleFactor struct {
	*_BACnetConstructedData
	ScaleFactor BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataScaleFactor) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataScaleFactor) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SCALE_FACTOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataScaleFactor) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataScaleFactor) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataScaleFactor) GetScaleFactor() BACnetApplicationTagReal {
	return m.ScaleFactor
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataScaleFactor) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetScaleFactor())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataScaleFactor factory function for _BACnetConstructedDataScaleFactor
func NewBACnetConstructedDataScaleFactor(scaleFactor BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataScaleFactor {
	_result := &_BACnetConstructedDataScaleFactor{
		ScaleFactor:            scaleFactor,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataScaleFactor(structType interface{}) BACnetConstructedDataScaleFactor {
	if casted, ok := structType.(BACnetConstructedDataScaleFactor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataScaleFactor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataScaleFactor) GetTypeName() string {
	return "BACnetConstructedDataScaleFactor"
}

func (m *_BACnetConstructedDataScaleFactor) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataScaleFactor) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (scaleFactor)
	lengthInBits += m.ScaleFactor.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataScaleFactor) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataScaleFactorParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataScaleFactor, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataScaleFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataScaleFactor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (scaleFactor)
	if pullErr := readBuffer.PullContext("scaleFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for scaleFactor")
	}
	_scaleFactor, _scaleFactorErr := BACnetApplicationTagParse(readBuffer)
	if _scaleFactorErr != nil {
		return nil, errors.Wrap(_scaleFactorErr, "Error parsing 'scaleFactor' field of BACnetConstructedDataScaleFactor")
	}
	scaleFactor := _scaleFactor.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("scaleFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for scaleFactor")
	}

	// Virtual field
	_actualValue := scaleFactor
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataScaleFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataScaleFactor")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataScaleFactor{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ScaleFactor: scaleFactor,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataScaleFactor) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataScaleFactor) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataScaleFactor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataScaleFactor")
		}

		// Simple Field (scaleFactor)
		if pushErr := writeBuffer.PushContext("scaleFactor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for scaleFactor")
		}
		_scaleFactorErr := writeBuffer.WriteSerializable(m.GetScaleFactor())
		if popErr := writeBuffer.PopContext("scaleFactor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for scaleFactor")
		}
		if _scaleFactorErr != nil {
			return errors.Wrap(_scaleFactorErr, "Error serializing 'scaleFactor' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataScaleFactor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataScaleFactor")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataScaleFactor) isBACnetConstructedDataScaleFactor() bool {
	return true
}

func (m *_BACnetConstructedDataScaleFactor) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
