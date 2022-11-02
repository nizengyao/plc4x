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

// BACnetEventParameterChangeOfStateListOfValues is the corresponding interface of BACnetEventParameterChangeOfStateListOfValues
type BACnetEventParameterChangeOfStateListOfValues interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() []BACnetPropertyStates
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfStateListOfValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfStateListOfValues.
// This is useful for switch cases.
type BACnetEventParameterChangeOfStateListOfValuesExactly interface {
	BACnetEventParameterChangeOfStateListOfValues
	isBACnetEventParameterChangeOfStateListOfValues() bool
}

// _BACnetEventParameterChangeOfStateListOfValues is the data-structure of this message
type _BACnetEventParameterChangeOfStateListOfValues struct {
	OpeningTag   BACnetOpeningTag
	ListOfValues []BACnetPropertyStates
	ClosingTag   BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetListOfValues() []BACnetPropertyStates {
	return m.ListOfValues
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfStateListOfValues factory function for _BACnetEventParameterChangeOfStateListOfValues
func NewBACnetEventParameterChangeOfStateListOfValues(openingTag BACnetOpeningTag, listOfValues []BACnetPropertyStates, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfStateListOfValues {
	return &_BACnetEventParameterChangeOfStateListOfValues{OpeningTag: openingTag, ListOfValues: listOfValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfStateListOfValues(structType interface{}) BACnetEventParameterChangeOfStateListOfValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfStateListOfValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfStateListOfValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfStateListOfValues"
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfValues) > 0 {
		for _, element := range m.ListOfValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterChangeOfStateListOfValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfStateListOfValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfStateListOfValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfStateListOfValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterChangeOfStateListOfValues")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfValues)
	if pullErr := readBuffer.PullContext("listOfValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfValues")
	}
	// Terminated array
	var listOfValues []BACnetPropertyStates
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetPropertyStatesParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfValues' field of BACnetEventParameterChangeOfStateListOfValues")
			}
			listOfValues = append(listOfValues, _item.(BACnetPropertyStates))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterChangeOfStateListOfValues")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfStateListOfValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfStateListOfValues")
	}

	// Create the instance
	return &_BACnetEventParameterChangeOfStateListOfValues{
		TagNumber:    tagNumber,
		OpeningTag:   openingTag,
		ListOfValues: listOfValues,
		ClosingTag:   closingTag,
	}, nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfStateListOfValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfStateListOfValues")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfValues)
	if pushErr := writeBuffer.PushContext("listOfValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfValues")
	}
	for _, _element := range m.GetListOfValues() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfValues")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfStateListOfValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfStateListOfValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfStateListOfValues) isBACnetEventParameterChangeOfStateListOfValues() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
