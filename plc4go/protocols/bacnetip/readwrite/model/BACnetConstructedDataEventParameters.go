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

// BACnetConstructedDataEventParameters is the corresponding interface of BACnetConstructedDataEventParameters
type BACnetConstructedDataEventParameters interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEventParameter returns EventParameter (property field)
	GetEventParameter() BACnetEventParameter
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventParameter
}

// BACnetConstructedDataEventParametersExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventParameters.
// This is useful for switch cases.
type BACnetConstructedDataEventParametersExactly interface {
	BACnetConstructedDataEventParameters
	isBACnetConstructedDataEventParameters() bool
}

// _BACnetConstructedDataEventParameters is the data-structure of this message
type _BACnetConstructedDataEventParameters struct {
	*_BACnetConstructedData
	EventParameter BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventParameters) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventParameters) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_PARAMETERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventParameters) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEventParameters) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventParameters) GetEventParameter() BACnetEventParameter {
	return m.EventParameter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventParameters) GetActualValue() BACnetEventParameter {
	return CastBACnetEventParameter(m.GetEventParameter())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventParameters factory function for _BACnetConstructedDataEventParameters
func NewBACnetConstructedDataEventParameters(eventParameter BACnetEventParameter, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventParameters {
	_result := &_BACnetConstructedDataEventParameters{
		EventParameter:         eventParameter,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventParameters(structType interface{}) BACnetConstructedDataEventParameters {
	if casted, ok := structType.(BACnetConstructedDataEventParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventParameters) GetTypeName() string {
	return "BACnetConstructedDataEventParameters"
}

func (m *_BACnetConstructedDataEventParameters) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEventParameters) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventParameter)
	lengthInBits += m.EventParameter.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventParameters) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventParametersParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventParameters, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventParameter)
	if pullErr := readBuffer.PullContext("eventParameter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventParameter")
	}
	_eventParameter, _eventParameterErr := BACnetEventParameterParse(readBuffer)
	if _eventParameterErr != nil {
		return nil, errors.Wrap(_eventParameterErr, "Error parsing 'eventParameter' field of BACnetConstructedDataEventParameters")
	}
	eventParameter := _eventParameter.(BACnetEventParameter)
	if closeErr := readBuffer.CloseContext("eventParameter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventParameter")
	}

	// Virtual field
	_actualValue := eventParameter
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventParameters")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEventParameters{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		EventParameter: eventParameter,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEventParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventParameters) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventParameters")
		}

		// Simple Field (eventParameter)
		if pushErr := writeBuffer.PushContext("eventParameter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventParameter")
		}
		_eventParameterErr := writeBuffer.WriteSerializable(m.GetEventParameter())
		if popErr := writeBuffer.PopContext("eventParameter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventParameter")
		}
		if _eventParameterErr != nil {
			return errors.Wrap(_eventParameterErr, "Error serializing 'eventParameter' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventParameters")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventParameters) isBACnetConstructedDataEventParameters() bool {
	return true
}

func (m *_BACnetConstructedDataEventParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
