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

// BACnetRestartReasonTagged is the corresponding interface of BACnetRestartReasonTagged
type BACnetRestartReasonTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetRestartReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetRestartReasonTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetRestartReasonTagged.
// This is useful for switch cases.
type BACnetRestartReasonTaggedExactly interface {
	BACnetRestartReasonTagged
	isBACnetRestartReasonTagged() bool
}

// _BACnetRestartReasonTagged is the data-structure of this message
type _BACnetRestartReasonTagged struct {
	Header           BACnetTagHeader
	Value            BACnetRestartReason
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRestartReasonTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetRestartReasonTagged) GetValue() BACnetRestartReason {
	return m.Value
}

func (m *_BACnetRestartReasonTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetRestartReasonTagged) GetIsProprietary() bool {
	return bool(bool((m.GetValue()) == (BACnetRestartReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRestartReasonTagged factory function for _BACnetRestartReasonTagged
func NewBACnetRestartReasonTagged(header BACnetTagHeader, value BACnetRestartReason, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetRestartReasonTagged {
	return &_BACnetRestartReasonTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetRestartReasonTagged(structType interface{}) BACnetRestartReasonTagged {
	if casted, ok := structType.(BACnetRestartReasonTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRestartReasonTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRestartReasonTagged) GetTypeName() string {
	return "BACnetRestartReasonTagged"
}

func (m *_BACnetRestartReasonTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetRestartReasonTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} { return int32(int32(0)) }, func() interface{} { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() interface{} { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetRestartReasonTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRestartReasonTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetRestartReasonTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRestartReasonTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRestartReasonTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetRestartReasonTagged")
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
	_value, _valueErr := ReadEnumGeneric(readBuffer, header.GetActualLength(), BACnetRestartReason_VENDOR_PROPRIETARY_VALUE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetRestartReasonTagged")
	}
	var value BACnetRestartReason
	if _value != nil {
		value = _value.(BACnetRestartReason)
	}

	// Virtual field
	_isProprietary := bool((value) == (BACnetRestartReason_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Manual Field (proprietaryValue)
	_proprietaryValue, _proprietaryValueErr := ReadProprietaryEnumGeneric(readBuffer, header.GetActualLength(), isProprietary)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field of BACnetRestartReasonTagged")
	}
	var proprietaryValue uint32
	if _proprietaryValue != nil {
		proprietaryValue = _proprietaryValue.(uint32)
	}

	if closeErr := readBuffer.CloseContext("BACnetRestartReasonTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRestartReasonTagged")
	}

	// Create the instance
	return NewBACnetRestartReasonTagged(header, value, proprietaryValue, tagNumber, tagClass), nil
}

func (m *_BACnetRestartReasonTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetRestartReasonTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRestartReasonTagged")
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
	// Virtual field
	if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryEnumGeneric(writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRestartReasonTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRestartReasonTagged")
	}
	return nil
}

func (m *_BACnetRestartReasonTagged) isBACnetRestartReasonTagged() bool {
	return true
}

func (m *_BACnetRestartReasonTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
