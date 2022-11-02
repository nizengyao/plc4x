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

// BACnetResultFlagsTagged is the corresponding interface of BACnetResultFlagsTagged
type BACnetResultFlagsTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetFirstItem returns FirstItem (virtual field)
	GetFirstItem() bool
	// GetLastItem returns LastItem (virtual field)
	GetLastItem() bool
	// GetMoreItems returns MoreItems (virtual field)
	GetMoreItems() bool
}

// BACnetResultFlagsTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetResultFlagsTagged.
// This is useful for switch cases.
type BACnetResultFlagsTaggedExactly interface {
	BACnetResultFlagsTagged
	isBACnetResultFlagsTagged() bool
}

// _BACnetResultFlagsTagged is the data-structure of this message
type _BACnetResultFlagsTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetResultFlagsTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetResultFlagsTagged) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetResultFlagsTagged) GetFirstItem() bool {
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (0))), func() interface{} { return bool(m.GetPayload().GetData()[0]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetResultFlagsTagged) GetLastItem() bool {
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (1))), func() interface{} { return bool(m.GetPayload().GetData()[1]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetResultFlagsTagged) GetMoreItems() bool {
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (2))), func() interface{} { return bool(m.GetPayload().GetData()[2]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetResultFlagsTagged factory function for _BACnetResultFlagsTagged
func NewBACnetResultFlagsTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetResultFlagsTagged {
	return &_BACnetResultFlagsTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetResultFlagsTagged(structType interface{}) BACnetResultFlagsTagged {
	if casted, ok := structType.(BACnetResultFlagsTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetResultFlagsTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetResultFlagsTagged) GetTypeName() string {
	return "BACnetResultFlagsTagged"
}

func (m *_BACnetResultFlagsTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetResultFlagsTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetResultFlagsTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetResultFlagsTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetResultFlagsTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetResultFlagsTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetResultFlagsTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetResultFlagsTagged")
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

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadBitStringParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetResultFlagsTagged")
	}
	payload := _payload.(BACnetTagPayloadBitString)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_firstItem := utils.InlineIf((bool((len(payload.GetData())) > (0))), func() interface{} { return bool(payload.GetData()[0]) }, func() interface{} { return bool(bool(false)) }).(bool)
	firstItem := bool(_firstItem)
	_ = firstItem

	// Virtual field
	_lastItem := utils.InlineIf((bool((len(payload.GetData())) > (1))), func() interface{} { return bool(payload.GetData()[1]) }, func() interface{} { return bool(bool(false)) }).(bool)
	lastItem := bool(_lastItem)
	_ = lastItem

	// Virtual field
	_moreItems := utils.InlineIf((bool((len(payload.GetData())) > (2))), func() interface{} { return bool(payload.GetData()[2]) }, func() interface{} { return bool(bool(false)) }).(bool)
	moreItems := bool(_moreItems)
	_ = moreItems

	if closeErr := readBuffer.CloseContext("BACnetResultFlagsTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetResultFlagsTagged")
	}

	// Create the instance
	return &_BACnetResultFlagsTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Payload:   payload,
	}, nil
}

func (m *_BACnetResultFlagsTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetResultFlagsTagged) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetResultFlagsTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetResultFlagsTagged")
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

	// Simple Field (payload)
	if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for payload")
	}
	_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
	if popErr := writeBuffer.PopContext("payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for payload")
	}
	if _payloadErr != nil {
		return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
	}
	// Virtual field
	if _firstItemErr := writeBuffer.WriteVirtual("firstItem", m.GetFirstItem()); _firstItemErr != nil {
		return errors.Wrap(_firstItemErr, "Error serializing 'firstItem' field")
	}
	// Virtual field
	if _lastItemErr := writeBuffer.WriteVirtual("lastItem", m.GetLastItem()); _lastItemErr != nil {
		return errors.Wrap(_lastItemErr, "Error serializing 'lastItem' field")
	}
	// Virtual field
	if _moreItemsErr := writeBuffer.WriteVirtual("moreItems", m.GetMoreItems()); _moreItemsErr != nil {
		return errors.Wrap(_moreItemsErr, "Error serializing 'moreItems' field")
	}

	if popErr := writeBuffer.PopContext("BACnetResultFlagsTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetResultFlagsTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetResultFlagsTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetResultFlagsTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetResultFlagsTagged) isBACnetResultFlagsTagged() bool {
	return true
}

func (m *_BACnetResultFlagsTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
