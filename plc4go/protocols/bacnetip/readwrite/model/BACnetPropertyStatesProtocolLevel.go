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

// BACnetPropertyStatesProtocolLevel is the corresponding interface of BACnetPropertyStatesProtocolLevel
type BACnetPropertyStatesProtocolLevel interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetProtocolLevel returns ProtocolLevel (property field)
	GetProtocolLevel() BACnetProtocolLevelTagged
}

// BACnetPropertyStatesProtocolLevelExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesProtocolLevel.
// This is useful for switch cases.
type BACnetPropertyStatesProtocolLevelExactly interface {
	BACnetPropertyStatesProtocolLevel
	isBACnetPropertyStatesProtocolLevel() bool
}

// _BACnetPropertyStatesProtocolLevel is the data-structure of this message
type _BACnetPropertyStatesProtocolLevel struct {
	*_BACnetPropertyStates
	ProtocolLevel BACnetProtocolLevelTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesProtocolLevel) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesProtocolLevel) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesProtocolLevel) GetProtocolLevel() BACnetProtocolLevelTagged {
	return m.ProtocolLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesProtocolLevel factory function for _BACnetPropertyStatesProtocolLevel
func NewBACnetPropertyStatesProtocolLevel(protocolLevel BACnetProtocolLevelTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesProtocolLevel {
	_result := &_BACnetPropertyStatesProtocolLevel{
		ProtocolLevel:         protocolLevel,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesProtocolLevel(structType interface{}) BACnetPropertyStatesProtocolLevel {
	if casted, ok := structType.(BACnetPropertyStatesProtocolLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesProtocolLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesProtocolLevel) GetTypeName() string {
	return "BACnetPropertyStatesProtocolLevel"
}

func (m *_BACnetPropertyStatesProtocolLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesProtocolLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (protocolLevel)
	lengthInBits += m.ProtocolLevel.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesProtocolLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesProtocolLevelParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesProtocolLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesProtocolLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesProtocolLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (protocolLevel)
	if pullErr := readBuffer.PullContext("protocolLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for protocolLevel")
	}
	_protocolLevel, _protocolLevelErr := BACnetProtocolLevelTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _protocolLevelErr != nil {
		return nil, errors.Wrap(_protocolLevelErr, "Error parsing 'protocolLevel' field of BACnetPropertyStatesProtocolLevel")
	}
	protocolLevel := _protocolLevel.(BACnetProtocolLevelTagged)
	if closeErr := readBuffer.CloseContext("protocolLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for protocolLevel")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesProtocolLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesProtocolLevel")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesProtocolLevel{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		ProtocolLevel:         protocolLevel,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesProtocolLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesProtocolLevel) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesProtocolLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesProtocolLevel")
		}

		// Simple Field (protocolLevel)
		if pushErr := writeBuffer.PushContext("protocolLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for protocolLevel")
		}
		_protocolLevelErr := writeBuffer.WriteSerializable(m.GetProtocolLevel())
		if popErr := writeBuffer.PopContext("protocolLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for protocolLevel")
		}
		if _protocolLevelErr != nil {
			return errors.Wrap(_protocolLevelErr, "Error serializing 'protocolLevel' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesProtocolLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesProtocolLevel")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesProtocolLevel) isBACnetPropertyStatesProtocolLevel() bool {
	return true
}

func (m *_BACnetPropertyStatesProtocolLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
