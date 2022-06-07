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

// BACnetConstructedDataProtocolObjectTypesSupported is the data-structure of this message
type BACnetConstructedDataProtocolObjectTypesSupported struct {
	*BACnetConstructedData
	ProtocolObjectTypesSupported *BACnetObjectTypesSupportedTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataProtocolObjectTypesSupported is the corresponding interface of BACnetConstructedDataProtocolObjectTypesSupported
type IBACnetConstructedDataProtocolObjectTypesSupported interface {
	IBACnetConstructedData
	// GetProtocolObjectTypesSupported returns ProtocolObjectTypesSupported (property field)
	GetProtocolObjectTypesSupported() *BACnetObjectTypesSupportedTagged
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

func (m *BACnetConstructedDataProtocolObjectTypesSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataProtocolObjectTypesSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROTOCOL_OBJECT_TYPES_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataProtocolObjectTypesSupported) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataProtocolObjectTypesSupported) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataProtocolObjectTypesSupported) GetProtocolObjectTypesSupported() *BACnetObjectTypesSupportedTagged {
	return m.ProtocolObjectTypesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProtocolObjectTypesSupported factory function for BACnetConstructedDataProtocolObjectTypesSupported
func NewBACnetConstructedDataProtocolObjectTypesSupported(protocolObjectTypesSupported *BACnetObjectTypesSupportedTagged, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataProtocolObjectTypesSupported {
	_result := &BACnetConstructedDataProtocolObjectTypesSupported{
		ProtocolObjectTypesSupported: protocolObjectTypesSupported,
		BACnetConstructedData:        NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataProtocolObjectTypesSupported(structType interface{}) *BACnetConstructedDataProtocolObjectTypesSupported {
	if casted, ok := structType.(BACnetConstructedDataProtocolObjectTypesSupported); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProtocolObjectTypesSupported); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataProtocolObjectTypesSupported(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataProtocolObjectTypesSupported(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataProtocolObjectTypesSupported) GetTypeName() string {
	return "BACnetConstructedDataProtocolObjectTypesSupported"
}

func (m *BACnetConstructedDataProtocolObjectTypesSupported) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataProtocolObjectTypesSupported) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (protocolObjectTypesSupported)
	lengthInBits += m.ProtocolObjectTypesSupported.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataProtocolObjectTypesSupported) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProtocolObjectTypesSupportedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataProtocolObjectTypesSupported, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProtocolObjectTypesSupported"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (protocolObjectTypesSupported)
	if pullErr := readBuffer.PullContext("protocolObjectTypesSupported"); pullErr != nil {
		return nil, pullErr
	}
	_protocolObjectTypesSupported, _protocolObjectTypesSupportedErr := BACnetObjectTypesSupportedTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _protocolObjectTypesSupportedErr != nil {
		return nil, errors.Wrap(_protocolObjectTypesSupportedErr, "Error parsing 'protocolObjectTypesSupported' field")
	}
	protocolObjectTypesSupported := CastBACnetObjectTypesSupportedTagged(_protocolObjectTypesSupported)
	if closeErr := readBuffer.CloseContext("protocolObjectTypesSupported"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProtocolObjectTypesSupported"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataProtocolObjectTypesSupported{
		ProtocolObjectTypesSupported: CastBACnetObjectTypesSupportedTagged(protocolObjectTypesSupported),
		BACnetConstructedData:        &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataProtocolObjectTypesSupported) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProtocolObjectTypesSupported"); pushErr != nil {
			return pushErr
		}

		// Simple Field (protocolObjectTypesSupported)
		if pushErr := writeBuffer.PushContext("protocolObjectTypesSupported"); pushErr != nil {
			return pushErr
		}
		_protocolObjectTypesSupportedErr := m.ProtocolObjectTypesSupported.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("protocolObjectTypesSupported"); popErr != nil {
			return popErr
		}
		if _protocolObjectTypesSupportedErr != nil {
			return errors.Wrap(_protocolObjectTypesSupportedErr, "Error serializing 'protocolObjectTypesSupported' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProtocolObjectTypesSupported"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataProtocolObjectTypesSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
