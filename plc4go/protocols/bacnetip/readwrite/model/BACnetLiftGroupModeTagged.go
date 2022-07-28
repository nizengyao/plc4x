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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLiftGroupModeTagged is the corresponding interface of BACnetLiftGroupModeTagged
type BACnetLiftGroupModeTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLiftGroupMode
}

// BACnetLiftGroupModeTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetLiftGroupModeTagged.
// This is useful for switch cases.
type BACnetLiftGroupModeTaggedExactly interface {
	BACnetLiftGroupModeTagged
	isBACnetLiftGroupModeTagged() bool
}

// _BACnetLiftGroupModeTagged is the data-structure of this message
type _BACnetLiftGroupModeTagged struct {
	Header BACnetTagHeader
	Value  BACnetLiftGroupMode

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLiftGroupModeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLiftGroupModeTagged) GetValue() BACnetLiftGroupMode {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLiftGroupModeTagged factory function for _BACnetLiftGroupModeTagged
func NewBACnetLiftGroupModeTagged(header BACnetTagHeader, value BACnetLiftGroupMode, tagNumber uint8, tagClass TagClass) *_BACnetLiftGroupModeTagged {
	return &_BACnetLiftGroupModeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetLiftGroupModeTagged(structType interface{}) BACnetLiftGroupModeTagged {
	if casted, ok := structType.(BACnetLiftGroupModeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLiftGroupModeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLiftGroupModeTagged) GetTypeName() string {
	return "BACnetLiftGroupModeTagged"
}

func (m *_BACnetLiftGroupModeTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLiftGroupModeTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetLiftGroupModeTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLiftGroupModeTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLiftGroupModeTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLiftGroupModeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLiftGroupModeTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetLiftGroupModeTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetLiftGroupMode_UNKNOWN)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetLiftGroupModeTagged")
	}
	var value BACnetLiftGroupMode
	if _value != nil {
		value = _value.(BACnetLiftGroupMode)
	}

	if closeErr := readBuffer.CloseContext("BACnetLiftGroupModeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftGroupModeTagged")
	}

	// Create the instance
	return NewBACnetLiftGroupModeTagged(header, value, tagNumber, tagClass), nil
}

func (m *_BACnetLiftGroupModeTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLiftGroupModeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftGroupModeTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLiftGroupModeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLiftGroupModeTagged")
	}
	return nil
}

func (m *_BACnetLiftGroupModeTagged) isBACnetLiftGroupModeTagged() bool {
	return true
}

func (m *_BACnetLiftGroupModeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
