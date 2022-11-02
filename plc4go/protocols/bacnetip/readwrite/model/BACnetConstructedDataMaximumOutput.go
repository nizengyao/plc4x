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

// BACnetConstructedDataMaximumOutput is the corresponding interface of BACnetConstructedDataMaximumOutput
type BACnetConstructedDataMaximumOutput interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaximumOutput returns MaximumOutput (property field)
	GetMaximumOutput() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataMaximumOutputExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMaximumOutput.
// This is useful for switch cases.
type BACnetConstructedDataMaximumOutputExactly interface {
	BACnetConstructedDataMaximumOutput
	isBACnetConstructedDataMaximumOutput() bool
}

// _BACnetConstructedDataMaximumOutput is the data-structure of this message
type _BACnetConstructedDataMaximumOutput struct {
	*_BACnetConstructedData
	MaximumOutput BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaximumOutput) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaximumOutput) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAXIMUM_OUTPUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaximumOutput) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMaximumOutput) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaximumOutput) GetMaximumOutput() BACnetApplicationTagReal {
	return m.MaximumOutput
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaximumOutput) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetMaximumOutput())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaximumOutput factory function for _BACnetConstructedDataMaximumOutput
func NewBACnetConstructedDataMaximumOutput(maximumOutput BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaximumOutput {
	_result := &_BACnetConstructedDataMaximumOutput{
		MaximumOutput:          maximumOutput,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaximumOutput(structType interface{}) BACnetConstructedDataMaximumOutput {
	if casted, ok := structType.(BACnetConstructedDataMaximumOutput); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaximumOutput); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaximumOutput) GetTypeName() string {
	return "BACnetConstructedDataMaximumOutput"
}

func (m *_BACnetConstructedDataMaximumOutput) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMaximumOutput) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maximumOutput)
	lengthInBits += m.MaximumOutput.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaximumOutput) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaximumOutputParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaximumOutput, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaximumOutput"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaximumOutput")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maximumOutput)
	if pullErr := readBuffer.PullContext("maximumOutput"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maximumOutput")
	}
	_maximumOutput, _maximumOutputErr := BACnetApplicationTagParse(readBuffer)
	if _maximumOutputErr != nil {
		return nil, errors.Wrap(_maximumOutputErr, "Error parsing 'maximumOutput' field of BACnetConstructedDataMaximumOutput")
	}
	maximumOutput := _maximumOutput.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("maximumOutput"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maximumOutput")
	}

	// Virtual field
	_actualValue := maximumOutput
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaximumOutput"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaximumOutput")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMaximumOutput{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaximumOutput: maximumOutput,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMaximumOutput) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaximumOutput) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaximumOutput"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaximumOutput")
		}

		// Simple Field (maximumOutput)
		if pushErr := writeBuffer.PushContext("maximumOutput"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maximumOutput")
		}
		_maximumOutputErr := writeBuffer.WriteSerializable(m.GetMaximumOutput())
		if popErr := writeBuffer.PopContext("maximumOutput"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maximumOutput")
		}
		if _maximumOutputErr != nil {
			return errors.Wrap(_maximumOutputErr, "Error serializing 'maximumOutput' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaximumOutput"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaximumOutput")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaximumOutput) isBACnetConstructedDataMaximumOutput() bool {
	return true
}

func (m *_BACnetConstructedDataMaximumOutput) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
