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

// BACnetConstructedDataVendorIdentifier is the data-structure of this message
type BACnetConstructedDataVendorIdentifier struct {
	*BACnetConstructedData
	VendorIdentifier *BACnetVendorIdTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataVendorIdentifier is the corresponding interface of BACnetConstructedDataVendorIdentifier
type IBACnetConstructedDataVendorIdentifier interface {
	IBACnetConstructedData
	// GetVendorIdentifier returns VendorIdentifier (property field)
	GetVendorIdentifier() *BACnetVendorIdTagged
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

func (m *BACnetConstructedDataVendorIdentifier) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataVendorIdentifier) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VENDOR_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataVendorIdentifier) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataVendorIdentifier) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataVendorIdentifier) GetVendorIdentifier() *BACnetVendorIdTagged {
	return m.VendorIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVendorIdentifier factory function for BACnetConstructedDataVendorIdentifier
func NewBACnetConstructedDataVendorIdentifier(vendorIdentifier *BACnetVendorIdTagged, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataVendorIdentifier {
	_result := &BACnetConstructedDataVendorIdentifier{
		VendorIdentifier:      vendorIdentifier,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataVendorIdentifier(structType interface{}) *BACnetConstructedDataVendorIdentifier {
	if casted, ok := structType.(BACnetConstructedDataVendorIdentifier); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVendorIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataVendorIdentifier(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataVendorIdentifier(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataVendorIdentifier) GetTypeName() string {
	return "BACnetConstructedDataVendorIdentifier"
}

func (m *BACnetConstructedDataVendorIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataVendorIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (vendorIdentifier)
	lengthInBits += m.VendorIdentifier.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataVendorIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataVendorIdentifierParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataVendorIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVendorIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vendorIdentifier)
	if pullErr := readBuffer.PullContext("vendorIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_vendorIdentifier, _vendorIdentifierErr := BACnetVendorIdTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _vendorIdentifierErr != nil {
		return nil, errors.Wrap(_vendorIdentifierErr, "Error parsing 'vendorIdentifier' field")
	}
	vendorIdentifier := CastBACnetVendorIdTagged(_vendorIdentifier)
	if closeErr := readBuffer.CloseContext("vendorIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVendorIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataVendorIdentifier{
		VendorIdentifier:      CastBACnetVendorIdTagged(vendorIdentifier),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataVendorIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVendorIdentifier"); pushErr != nil {
			return pushErr
		}

		// Simple Field (vendorIdentifier)
		if pushErr := writeBuffer.PushContext("vendorIdentifier"); pushErr != nil {
			return pushErr
		}
		_vendorIdentifierErr := m.VendorIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("vendorIdentifier"); popErr != nil {
			return popErr
		}
		if _vendorIdentifierErr != nil {
			return errors.Wrap(_vendorIdentifierErr, "Error serializing 'vendorIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVendorIdentifier"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataVendorIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
