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

// BACnetFaultParameterFaultExtendedParametersEntry is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntry
type BACnetFaultParameterFaultExtendedParametersEntry interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
}

// BACnetFaultParameterFaultExtendedParametersEntryExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParametersEntry.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersEntryExactly interface {
	BACnetFaultParameterFaultExtendedParametersEntry
	isBACnetFaultParameterFaultExtendedParametersEntry() bool
}

// _BACnetFaultParameterFaultExtendedParametersEntry is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntry struct {
	_BACnetFaultParameterFaultExtendedParametersEntryChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetFaultParameterFaultExtendedParametersEntryChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type BACnetFaultParameterFaultExtendedParametersEntryParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetFaultParameterFaultExtendedParametersEntry, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetFaultParameterFaultExtendedParametersEntryChild interface {
	utils.Serializable
	InitializeParent(parent BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetFaultParameterFaultExtendedParametersEntry

	GetTypeName() string
	BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetPeekedIsContextTag() bool {
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntry factory function for _BACnetFaultParameterFaultExtendedParametersEntry
func NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtendedParametersEntry {
	return &_BACnetFaultParameterFaultExtendedParametersEntry{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntry(structType interface{}) BACnetFaultParameterFaultExtendedParametersEntry {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntry"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryParse(readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Virtual field
	_peekedIsContextTag := bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))
	peekedIsContextTag := bool(_peekedIsContextTag)
	_ = peekedIsContextTag

	// Validation
	if !(bool((!(peekedIsContextTag))) || bool((bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"unexpected opening or closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetFaultParameterFaultExtendedParametersEntryChildSerializeRequirement interface {
		BACnetFaultParameterFaultExtendedParametersEntry
		InitializeParent(BACnetFaultParameterFaultExtendedParametersEntry, BACnetTagHeader)
		GetParent() BACnetFaultParameterFaultExtendedParametersEntry
	}
	var _childTemp interface{}
	var _child BACnetFaultParameterFaultExtendedParametersEntryChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x0 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryNull
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryNullParse(readBuffer)
	case peekedTagNumber == 0x4 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryReal
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryRealParse(readBuffer)
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryUnsigned
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryUnsignedParse(readBuffer)
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryBoolean
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryBooleanParse(readBuffer)
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryInteger
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryIntegerParse(readBuffer)
	case peekedTagNumber == 0x5 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryDouble
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryDoubleParse(readBuffer)
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryOctetString
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryOctetStringParse(readBuffer)
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryCharacterString
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryCharacterStringParse(readBuffer)
	case peekedTagNumber == 0x8 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryBitString
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryBitStringParse(readBuffer)
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryEnumerated
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryEnumeratedParse(readBuffer)
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryDate
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryDateParse(readBuffer)
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryTime
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryTimeParse(readBuffer)
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryObjectidentifier
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryObjectidentifierParse(readBuffer)
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetFaultParameterFaultExtendedParametersEntryReference
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryReferenceParse(readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v, peekedIsContextTag=%v]", peekedTagNumber, peekedIsContextTag)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
	}
	_child = _childTemp.(BACnetFaultParameterFaultExtendedParametersEntryChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntry")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetFaultParameterFaultExtendedParametersEntry) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetFaultParameterFaultExtendedParametersEntry, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntry")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual("peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntry")
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) isBACnetFaultParameterFaultExtendedParametersEntry() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
