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

// BACnetProcessIdSelectionValue is the corresponding interface of BACnetProcessIdSelectionValue
type BACnetProcessIdSelectionValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetProcessIdSelection
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetApplicationTagUnsignedInteger
}

// BACnetProcessIdSelectionValueExactly can be used when we want exactly this type and not a type which fulfills BACnetProcessIdSelectionValue.
// This is useful for switch cases.
type BACnetProcessIdSelectionValueExactly interface {
	BACnetProcessIdSelectionValue
	isBACnetProcessIdSelectionValue() bool
}

// _BACnetProcessIdSelectionValue is the data-structure of this message
type _BACnetProcessIdSelectionValue struct {
	*_BACnetProcessIdSelection
	ProcessIdentifier BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetProcessIdSelectionValue) InitializeParent(parent BACnetProcessIdSelection, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetProcessIdSelectionValue) GetParent() BACnetProcessIdSelection {
	return m._BACnetProcessIdSelection
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProcessIdSelectionValue) GetProcessIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.ProcessIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetProcessIdSelectionValue factory function for _BACnetProcessIdSelectionValue
func NewBACnetProcessIdSelectionValue(processIdentifier BACnetApplicationTagUnsignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetProcessIdSelectionValue {
	_result := &_BACnetProcessIdSelectionValue{
		ProcessIdentifier:         processIdentifier,
		_BACnetProcessIdSelection: NewBACnetProcessIdSelection(peekedTagHeader),
	}
	_result._BACnetProcessIdSelection._BACnetProcessIdSelectionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetProcessIdSelectionValue(structType interface{}) BACnetProcessIdSelectionValue {
	if casted, ok := structType.(BACnetProcessIdSelectionValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProcessIdSelectionValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProcessIdSelectionValue) GetTypeName() string {
	return "BACnetProcessIdSelectionValue"
}

func (m *_BACnetProcessIdSelectionValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetProcessIdSelectionValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetProcessIdSelectionValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetProcessIdSelectionValueParse(readBuffer utils.ReadBuffer) (BACnetProcessIdSelectionValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProcessIdSelectionValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProcessIdSelectionValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (processIdentifier)
	if pullErr := readBuffer.PullContext("processIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for processIdentifier")
	}
	_processIdentifier, _processIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _processIdentifierErr != nil {
		return nil, errors.Wrap(_processIdentifierErr, "Error parsing 'processIdentifier' field of BACnetProcessIdSelectionValue")
	}
	processIdentifier := _processIdentifier.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("processIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for processIdentifier")
	}

	if closeErr := readBuffer.CloseContext("BACnetProcessIdSelectionValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProcessIdSelectionValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetProcessIdSelectionValue{
		_BACnetProcessIdSelection: &_BACnetProcessIdSelection{},
		ProcessIdentifier:         processIdentifier,
	}
	_child._BACnetProcessIdSelection._BACnetProcessIdSelectionChildRequirements = _child
	return _child, nil
}

func (m *_BACnetProcessIdSelectionValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProcessIdSelectionValue) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetProcessIdSelectionValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetProcessIdSelectionValue")
		}

		// Simple Field (processIdentifier)
		if pushErr := writeBuffer.PushContext("processIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for processIdentifier")
		}
		_processIdentifierErr := writeBuffer.WriteSerializable(m.GetProcessIdentifier())
		if popErr := writeBuffer.PopContext("processIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for processIdentifier")
		}
		if _processIdentifierErr != nil {
			return errors.Wrap(_processIdentifierErr, "Error serializing 'processIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetProcessIdSelectionValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetProcessIdSelectionValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetProcessIdSelectionValue) isBACnetProcessIdSelectionValue() bool {
	return true
}

func (m *_BACnetProcessIdSelectionValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
