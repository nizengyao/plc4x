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

// BACnetConstructedDataVendorName is the data-structure of this message
type BACnetConstructedDataVendorName struct {
	*BACnetConstructedData
	VendorName *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataVendorName is the corresponding interface of BACnetConstructedDataVendorName
type IBACnetConstructedDataVendorName interface {
	IBACnetConstructedData
	// GetVendorName returns VendorName (property field)
	GetVendorName() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataVendorName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataVendorName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VENDOR_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataVendorName) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataVendorName) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataVendorName) GetVendorName() *BACnetApplicationTagCharacterString {
	return m.VendorName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVendorName factory function for BACnetConstructedDataVendorName
func NewBACnetConstructedDataVendorName(vendorName *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataVendorName {
	_result := &BACnetConstructedDataVendorName{
		VendorName:            vendorName,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataVendorName(structType interface{}) *BACnetConstructedDataVendorName {
	if casted, ok := structType.(BACnetConstructedDataVendorName); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVendorName); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataVendorName(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataVendorName(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataVendorName) GetTypeName() string {
	return "BACnetConstructedDataVendorName"
}

func (m *BACnetConstructedDataVendorName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataVendorName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (vendorName)
	lengthInBits += m.VendorName.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataVendorName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataVendorNameParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataVendorName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVendorName"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vendorName)
	if pullErr := readBuffer.PullContext("vendorName"); pullErr != nil {
		return nil, pullErr
	}
	_vendorName, _vendorNameErr := BACnetApplicationTagParse(readBuffer)
	if _vendorNameErr != nil {
		return nil, errors.Wrap(_vendorNameErr, "Error parsing 'vendorName' field")
	}
	vendorName := CastBACnetApplicationTagCharacterString(_vendorName)
	if closeErr := readBuffer.CloseContext("vendorName"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVendorName"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataVendorName{
		VendorName:            CastBACnetApplicationTagCharacterString(vendorName),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataVendorName) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVendorName"); pushErr != nil {
			return pushErr
		}

		// Simple Field (vendorName)
		if pushErr := writeBuffer.PushContext("vendorName"); pushErr != nil {
			return pushErr
		}
		_vendorNameErr := m.VendorName.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("vendorName"); popErr != nil {
			return popErr
		}
		if _vendorNameErr != nil {
			return errors.Wrap(_vendorNameErr, "Error serializing 'vendorName' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVendorName"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataVendorName) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
