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

// ApduControlDisconnect is the corresponding interface of ApduControlDisconnect
type ApduControlDisconnect interface {
	utils.LengthAware
	utils.Serializable
	ApduControl
}

// ApduControlDisconnectExactly can be used when we want exactly this type and not a type which fulfills ApduControlDisconnect.
// This is useful for switch cases.
type ApduControlDisconnectExactly interface {
	ApduControlDisconnect
	isApduControlDisconnect() bool
}

// _ApduControlDisconnect is the data-structure of this message
type _ApduControlDisconnect struct {
	*_ApduControl
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduControlDisconnect) GetControlType() uint8 {
	return 0x1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduControlDisconnect) InitializeParent(parent ApduControl) {}

func (m *_ApduControlDisconnect) GetParent() ApduControl {
	return m._ApduControl
}

// NewApduControlDisconnect factory function for _ApduControlDisconnect
func NewApduControlDisconnect() *_ApduControlDisconnect {
	_result := &_ApduControlDisconnect{
		_ApduControl: NewApduControl(),
	}
	_result._ApduControl._ApduControlChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduControlDisconnect(structType interface{}) ApduControlDisconnect {
	if casted, ok := structType.(ApduControlDisconnect); ok {
		return casted
	}
	if casted, ok := structType.(*ApduControlDisconnect); ok {
		return *casted
	}
	return nil
}

func (m *_ApduControlDisconnect) GetTypeName() string {
	return "ApduControlDisconnect"
}

func (m *_ApduControlDisconnect) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduControlDisconnect) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduControlDisconnect) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduControlDisconnectParse(readBuffer utils.ReadBuffer) (ApduControlDisconnect, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduControlDisconnect"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControlDisconnect")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduControlDisconnect"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControlDisconnect")
	}

	// Create a partially initialized instance
	_child := &_ApduControlDisconnect{
		_ApduControl: &_ApduControl{},
	}
	_child._ApduControl._ApduControlChildRequirements = _child
	return _child, nil
}

func (m *_ApduControlDisconnect) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduControlDisconnect) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlDisconnect"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduControlDisconnect")
		}

		if popErr := writeBuffer.PopContext("ApduControlDisconnect"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduControlDisconnect")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduControlDisconnect) isApduControlDisconnect() bool {
	return true
}

func (m *_ApduControlDisconnect) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
