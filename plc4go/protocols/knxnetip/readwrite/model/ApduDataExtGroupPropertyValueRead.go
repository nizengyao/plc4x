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

// ApduDataExtGroupPropertyValueRead is the corresponding interface of ApduDataExtGroupPropertyValueRead
type ApduDataExtGroupPropertyValueRead interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtGroupPropertyValueReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtGroupPropertyValueRead.
// This is useful for switch cases.
type ApduDataExtGroupPropertyValueReadExactly interface {
	ApduDataExtGroupPropertyValueRead
	isApduDataExtGroupPropertyValueRead() bool
}

// _ApduDataExtGroupPropertyValueRead is the data-structure of this message
type _ApduDataExtGroupPropertyValueRead struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtGroupPropertyValueRead) GetExtApciType() uint8 {
	return 0x28
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtGroupPropertyValueRead) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtGroupPropertyValueRead) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtGroupPropertyValueRead factory function for _ApduDataExtGroupPropertyValueRead
func NewApduDataExtGroupPropertyValueRead(length uint8) *_ApduDataExtGroupPropertyValueRead {
	_result := &_ApduDataExtGroupPropertyValueRead{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtGroupPropertyValueRead(structType interface{}) ApduDataExtGroupPropertyValueRead {
	if casted, ok := structType.(ApduDataExtGroupPropertyValueRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtGroupPropertyValueRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtGroupPropertyValueRead) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueRead"
}

func (m *_ApduDataExtGroupPropertyValueRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtGroupPropertyValueRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtGroupPropertyValueRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtGroupPropertyValueReadParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtGroupPropertyValueRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtGroupPropertyValueRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtGroupPropertyValueRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtGroupPropertyValueRead{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtGroupPropertyValueRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtGroupPropertyValueRead) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtGroupPropertyValueRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtGroupPropertyValueRead")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtGroupPropertyValueRead) isApduDataExtGroupPropertyValueRead() bool {
	return true
}

func (m *_ApduDataExtGroupPropertyValueRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
