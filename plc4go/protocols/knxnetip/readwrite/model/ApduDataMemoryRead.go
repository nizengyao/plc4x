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

// ApduDataMemoryRead is the corresponding interface of ApduDataMemoryRead
type ApduDataMemoryRead interface {
	utils.LengthAware
	utils.Serializable
	ApduData
	// GetNumBytes returns NumBytes (property field)
	GetNumBytes() uint8
	// GetAddress returns Address (property field)
	GetAddress() uint16
}

// ApduDataMemoryReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataMemoryRead.
// This is useful for switch cases.
type ApduDataMemoryReadExactly interface {
	ApduDataMemoryRead
	isApduDataMemoryRead() bool
}

// _ApduDataMemoryRead is the data-structure of this message
type _ApduDataMemoryRead struct {
	*_ApduData
	NumBytes uint8
	Address  uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataMemoryRead) GetApciType() uint8 {
	return 0x8
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataMemoryRead) InitializeParent(parent ApduData) {}

func (m *_ApduDataMemoryRead) GetParent() ApduData {
	return m._ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataMemoryRead) GetNumBytes() uint8 {
	return m.NumBytes
}

func (m *_ApduDataMemoryRead) GetAddress() uint16 {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataMemoryRead factory function for _ApduDataMemoryRead
func NewApduDataMemoryRead(numBytes uint8, address uint16, dataLength uint8) *_ApduDataMemoryRead {
	_result := &_ApduDataMemoryRead{
		NumBytes:  numBytes,
		Address:   address,
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataMemoryRead(structType interface{}) ApduDataMemoryRead {
	if casted, ok := structType.(ApduDataMemoryRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataMemoryRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataMemoryRead) GetTypeName() string {
	return "ApduDataMemoryRead"
}

func (m *_ApduDataMemoryRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataMemoryRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numBytes)
	lengthInBits += 6

	// Simple field (address)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ApduDataMemoryRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataMemoryReadParse(readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataMemoryRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataMemoryRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataMemoryRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numBytes)
	_numBytes, _numBytesErr := readBuffer.ReadUint8("numBytes", 6)
	if _numBytesErr != nil {
		return nil, errors.Wrap(_numBytesErr, "Error parsing 'numBytes' field of ApduDataMemoryRead")
	}
	numBytes := _numBytes

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of ApduDataMemoryRead")
	}
	address := _address

	if closeErr := readBuffer.CloseContext("ApduDataMemoryRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataMemoryRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataMemoryRead{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
		NumBytes: numBytes,
		Address:  address,
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataMemoryRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataMemoryRead) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataMemoryRead")
		}

		// Simple Field (numBytes)
		numBytes := uint8(m.GetNumBytes())
		_numBytesErr := writeBuffer.WriteUint8("numBytes", 6, (numBytes))
		if _numBytesErr != nil {
			return errors.Wrap(_numBytesErr, "Error serializing 'numBytes' field")
		}

		// Simple Field (address)
		address := uint16(m.GetAddress())
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataMemoryRead")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataMemoryRead) isApduDataMemoryRead() bool {
	return true
}

func (m *_ApduDataMemoryRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
