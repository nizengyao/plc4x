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

// RelativeTimestamp is the corresponding interface of RelativeTimestamp
type RelativeTimestamp interface {
	utils.LengthAware
	utils.Serializable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() uint16
}

// RelativeTimestampExactly can be used when we want exactly this type and not a type which fulfills RelativeTimestamp.
// This is useful for switch cases.
type RelativeTimestampExactly interface {
	RelativeTimestamp
	isRelativeTimestamp() bool
}

// _RelativeTimestamp is the data-structure of this message
type _RelativeTimestamp struct {
	Timestamp uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RelativeTimestamp) GetTimestamp() uint16 {
	return m.Timestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRelativeTimestamp factory function for _RelativeTimestamp
func NewRelativeTimestamp(timestamp uint16) *_RelativeTimestamp {
	return &_RelativeTimestamp{Timestamp: timestamp}
}

// Deprecated: use the interface for direct cast
func CastRelativeTimestamp(structType interface{}) RelativeTimestamp {
	if casted, ok := structType.(RelativeTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(*RelativeTimestamp); ok {
		return *casted
	}
	return nil
}

func (m *_RelativeTimestamp) GetTypeName() string {
	return "RelativeTimestamp"
}

func (m *_RelativeTimestamp) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RelativeTimestamp) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += 16

	return lengthInBits
}

func (m *_RelativeTimestamp) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RelativeTimestampParse(readBuffer utils.ReadBuffer) (RelativeTimestamp, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RelativeTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RelativeTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timestamp)
	_timestamp, _timestampErr := readBuffer.ReadUint16("timestamp", 16)
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field of RelativeTimestamp")
	}
	timestamp := _timestamp

	if closeErr := readBuffer.CloseContext("RelativeTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RelativeTimestamp")
	}

	// Create the instance
	return &_RelativeTimestamp{
		Timestamp: timestamp,
	}, nil
}

func (m *_RelativeTimestamp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RelativeTimestamp) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("RelativeTimestamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for RelativeTimestamp")
	}

	// Simple Field (timestamp)
	timestamp := uint16(m.GetTimestamp())
	_timestampErr := writeBuffer.WriteUint16("timestamp", 16, (timestamp))
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	if popErr := writeBuffer.PopContext("RelativeTimestamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for RelativeTimestamp")
	}
	return nil
}

func (m *_RelativeTimestamp) isRelativeTimestamp() bool {
	return true
}

func (m *_RelativeTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
