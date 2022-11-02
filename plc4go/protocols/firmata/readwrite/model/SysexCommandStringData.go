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

// SysexCommandStringData is the corresponding interface of SysexCommandStringData
type SysexCommandStringData interface {
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// SysexCommandStringDataExactly can be used when we want exactly this type and not a type which fulfills SysexCommandStringData.
// This is useful for switch cases.
type SysexCommandStringDataExactly interface {
	SysexCommandStringData
	isSysexCommandStringData() bool
}

// _SysexCommandStringData is the data-structure of this message
type _SysexCommandStringData struct {
	*_SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandStringData) GetCommandType() uint8 {
	return 0x71
}

func (m *_SysexCommandStringData) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandStringData) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandStringData) GetParent() SysexCommand {
	return m._SysexCommand
}

// NewSysexCommandStringData factory function for _SysexCommandStringData
func NewSysexCommandStringData() *_SysexCommandStringData {
	_result := &_SysexCommandStringData{
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandStringData(structType interface{}) SysexCommandStringData {
	if casted, ok := structType.(SysexCommandStringData); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandStringData); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandStringData) GetTypeName() string {
	return "SysexCommandStringData"
}

func (m *_SysexCommandStringData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SysexCommandStringData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SysexCommandStringData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandStringDataParse(readBuffer utils.ReadBuffer, response bool) (SysexCommandStringData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandStringData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandStringData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandStringData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandStringData")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandStringData{
		_SysexCommand: &_SysexCommand{},
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandStringData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandStringData) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandStringData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandStringData")
		}

		if popErr := writeBuffer.PopContext("SysexCommandStringData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandStringData")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SysexCommandStringData) isSysexCommandStringData() bool {
	return true
}

func (m *_SysexCommandStringData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
