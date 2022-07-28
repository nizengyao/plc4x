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

// BACnetAbortReasonTagged is the corresponding interface of BACnetAbortReasonTagged
type BACnetAbortReasonTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetValue returns Value (property field)
	GetValue() BACnetAbortReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetAbortReasonTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetAbortReasonTagged.
// This is useful for switch cases.
type BACnetAbortReasonTaggedExactly interface {
	BACnetAbortReasonTagged
	isBACnetAbortReasonTagged() bool
}

// _BACnetAbortReasonTagged is the data-structure of this message
type _BACnetAbortReasonTagged struct {
	Value            BACnetAbortReason
	ProprietaryValue uint32

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAbortReasonTagged) GetValue() BACnetAbortReason {
	return m.Value
}

func (m *_BACnetAbortReasonTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetAbortReasonTagged) GetIsProprietary() bool {
	return bool(bool((m.GetValue()) == (BACnetAbortReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAbortReasonTagged factory function for _BACnetAbortReasonTagged
func NewBACnetAbortReasonTagged(value BACnetAbortReason, proprietaryValue uint32, actualLength uint32) *_BACnetAbortReasonTagged {
	return &_BACnetAbortReasonTagged{Value: value, ProprietaryValue: proprietaryValue, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetAbortReasonTagged(structType interface{}) BACnetAbortReasonTagged {
	if casted, ok := structType.(BACnetAbortReasonTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAbortReasonTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAbortReasonTagged) GetTypeName() string {
	return "BACnetAbortReasonTagged"
}

func (m *_BACnetAbortReasonTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetAbortReasonTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} { return int32(int32(0)) }, func() interface{} { return int32((int32(m.ActualLength) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} { return int32((int32(m.ActualLength) * int32(int32(8)))) }, func() interface{} { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetAbortReasonTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAbortReasonTaggedParse(readBuffer utils.ReadBuffer, actualLength uint32) (BACnetAbortReasonTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAbortReasonTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAbortReasonTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Manual Field (value)
	_value, _valueErr := ReadEnumGeneric(readBuffer, actualLength, BACnetAbortReason_VENDOR_PROPRIETARY_VALUE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetAbortReasonTagged")
	}
	var value BACnetAbortReason
	if _value != nil {
		value = _value.(BACnetAbortReason)
	}

	// Virtual field
	_isProprietary := bool((value) == (BACnetAbortReason_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Manual Field (proprietaryValue)
	_proprietaryValue, _proprietaryValueErr := ReadProprietaryEnumGeneric(readBuffer, actualLength, isProprietary)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field of BACnetAbortReasonTagged")
	}
	var proprietaryValue uint32
	if _proprietaryValue != nil {
		proprietaryValue = _proprietaryValue.(uint32)
	}

	if closeErr := readBuffer.CloseContext("BACnetAbortReasonTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAbortReasonTagged")
	}

	// Create the instance
	return NewBACnetAbortReasonTagged(value, proprietaryValue, actualLength), nil
}

func (m *_BACnetAbortReasonTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAbortReasonTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAbortReasonTagged")
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

	if popErr := writeBuffer.PopContext("BACnetAbortReasonTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAbortReasonTagged")
	}
	return nil
}

func (m *_BACnetAbortReasonTagged) isBACnetAbortReasonTagged() bool {
	return true
}

func (m *_BACnetAbortReasonTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
