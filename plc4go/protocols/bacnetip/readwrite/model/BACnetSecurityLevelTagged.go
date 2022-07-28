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

// BACnetSecurityLevelTagged is the corresponding interface of BACnetSecurityLevelTagged
type BACnetSecurityLevelTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetSecurityLevel
}

// BACnetSecurityLevelTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetSecurityLevelTagged.
// This is useful for switch cases.
type BACnetSecurityLevelTaggedExactly interface {
	BACnetSecurityLevelTagged
	isBACnetSecurityLevelTagged() bool
}

// _BACnetSecurityLevelTagged is the data-structure of this message
type _BACnetSecurityLevelTagged struct {
	Header BACnetTagHeader
	Value  BACnetSecurityLevel

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSecurityLevelTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetSecurityLevelTagged) GetValue() BACnetSecurityLevel {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSecurityLevelTagged factory function for _BACnetSecurityLevelTagged
func NewBACnetSecurityLevelTagged(header BACnetTagHeader, value BACnetSecurityLevel, tagNumber uint8, tagClass TagClass) *_BACnetSecurityLevelTagged {
	return &_BACnetSecurityLevelTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetSecurityLevelTagged(structType interface{}) BACnetSecurityLevelTagged {
	if casted, ok := structType.(BACnetSecurityLevelTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSecurityLevelTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSecurityLevelTagged) GetTypeName() string {
	return "BACnetSecurityLevelTagged"
}

func (m *_BACnetSecurityLevelTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetSecurityLevelTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetSecurityLevelTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetSecurityLevelTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetSecurityLevelTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSecurityLevelTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSecurityLevelTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetSecurityLevelTagged")
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
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetSecurityLevel_INCAPABLE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetSecurityLevelTagged")
	}
	var value BACnetSecurityLevel
	if _value != nil {
		value = _value.(BACnetSecurityLevel)
	}

	if closeErr := readBuffer.CloseContext("BACnetSecurityLevelTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSecurityLevelTagged")
	}

	// Create the instance
	return NewBACnetSecurityLevelTagged(header, value, tagNumber, tagClass), nil
}

func (m *_BACnetSecurityLevelTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetSecurityLevelTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSecurityLevelTagged")
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

	if popErr := writeBuffer.PopContext("BACnetSecurityLevelTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSecurityLevelTagged")
	}
	return nil
}

func (m *_BACnetSecurityLevelTagged) isBACnetSecurityLevelTagged() bool {
	return true
}

func (m *_BACnetSecurityLevelTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
