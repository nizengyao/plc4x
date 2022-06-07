/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataDaylightSavingsStatus is the data-structure of this message
type BACnetConstructedDataDaylightSavingsStatus struct {
	*BACnetConstructedData
	DaylightSavingsStatus *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataDaylightSavingsStatus is the corresponding interface of BACnetConstructedDataDaylightSavingsStatus
type IBACnetConstructedDataDaylightSavingsStatus interface {
	IBACnetConstructedData
	// GetDaylightSavingsStatus returns DaylightSavingsStatus (property field)
	GetDaylightSavingsStatus() *BACnetApplicationTagBoolean
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataDaylightSavingsStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDaylightSavingsStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DAYLIGHT_SAVINGS_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDaylightSavingsStatus) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDaylightSavingsStatus) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDaylightSavingsStatus) GetDaylightSavingsStatus() *BACnetApplicationTagBoolean {
	return m.DaylightSavingsStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDaylightSavingsStatus factory function for BACnetConstructedDataDaylightSavingsStatus
func NewBACnetConstructedDataDaylightSavingsStatus(daylightSavingsStatus *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataDaylightSavingsStatus {
	_result := &BACnetConstructedDataDaylightSavingsStatus{
		DaylightSavingsStatus: daylightSavingsStatus,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDaylightSavingsStatus(structType interface{}) *BACnetConstructedDataDaylightSavingsStatus {
	if casted, ok := structType.(BACnetConstructedDataDaylightSavingsStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDaylightSavingsStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDaylightSavingsStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDaylightSavingsStatus(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDaylightSavingsStatus) GetTypeName() string {
	return "BACnetConstructedDataDaylightSavingsStatus"
}

func (m *BACnetConstructedDataDaylightSavingsStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDaylightSavingsStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (daylightSavingsStatus)
	lengthInBits += m.DaylightSavingsStatus.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDaylightSavingsStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDaylightSavingsStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataDaylightSavingsStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDaylightSavingsStatus"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (daylightSavingsStatus)
	if pullErr := readBuffer.PullContext("daylightSavingsStatus"); pullErr != nil {
		return nil, pullErr
	}
	_daylightSavingsStatus, _daylightSavingsStatusErr := BACnetApplicationTagParse(readBuffer)
	if _daylightSavingsStatusErr != nil {
		return nil, errors.Wrap(_daylightSavingsStatusErr, "Error parsing 'daylightSavingsStatus' field")
	}
	daylightSavingsStatus := CastBACnetApplicationTagBoolean(_daylightSavingsStatus)
	if closeErr := readBuffer.CloseContext("daylightSavingsStatus"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDaylightSavingsStatus"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDaylightSavingsStatus{
		DaylightSavingsStatus: CastBACnetApplicationTagBoolean(daylightSavingsStatus),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDaylightSavingsStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDaylightSavingsStatus"); pushErr != nil {
			return pushErr
		}

		// Simple Field (daylightSavingsStatus)
		if pushErr := writeBuffer.PushContext("daylightSavingsStatus"); pushErr != nil {
			return pushErr
		}
		_daylightSavingsStatusErr := m.DaylightSavingsStatus.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("daylightSavingsStatus"); popErr != nil {
			return popErr
		}
		if _daylightSavingsStatusErr != nil {
			return errors.Wrap(_daylightSavingsStatusErr, "Error serializing 'daylightSavingsStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDaylightSavingsStatus"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDaylightSavingsStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
