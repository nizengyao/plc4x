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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCReadBroadcastDistributionTableAck is the data-structure of this message
type BVLCReadBroadcastDistributionTableAck struct {
	*BVLC
	BdtEntries []byte

	// Arguments.
	BvlcPayloadLength uint16
}

// IBVLCReadBroadcastDistributionTableAck is the corresponding interface of BVLCReadBroadcastDistributionTableAck
type IBVLCReadBroadcastDistributionTableAck interface {
	IBVLC
	// GetBdtEntries returns BdtEntries (property field)
	GetBdtEntries() []byte
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

func (m *BVLCReadBroadcastDistributionTableAck) GetBvlcFunction() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BVLCReadBroadcastDistributionTableAck) InitializeParent(parent *BVLC) {}

func (m *BVLCReadBroadcastDistributionTableAck) GetParent() *BVLC {
	return m.BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BVLCReadBroadcastDistributionTableAck) GetBdtEntries() []byte {
	return m.BdtEntries
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCReadBroadcastDistributionTableAck factory function for BVLCReadBroadcastDistributionTableAck
func NewBVLCReadBroadcastDistributionTableAck(bdtEntries []byte, bvlcPayloadLength uint16) *BVLCReadBroadcastDistributionTableAck {
	_result := &BVLCReadBroadcastDistributionTableAck{
		BdtEntries: bdtEntries,
		BVLC:       NewBVLC(),
	}
	_result.Child = _result
	return _result
}

func CastBVLCReadBroadcastDistributionTableAck(structType interface{}) *BVLCReadBroadcastDistributionTableAck {
	if casted, ok := structType.(BVLCReadBroadcastDistributionTableAck); ok {
		return &casted
	}
	if casted, ok := structType.(*BVLCReadBroadcastDistributionTableAck); ok {
		return casted
	}
	if casted, ok := structType.(BVLC); ok {
		return CastBVLCReadBroadcastDistributionTableAck(casted.Child)
	}
	if casted, ok := structType.(*BVLC); ok {
		return CastBVLCReadBroadcastDistributionTableAck(casted.Child)
	}
	return nil
}

func (m *BVLCReadBroadcastDistributionTableAck) GetTypeName() string {
	return "BVLCReadBroadcastDistributionTableAck"
}

func (m *BVLCReadBroadcastDistributionTableAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BVLCReadBroadcastDistributionTableAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.BdtEntries) > 0 {
		lengthInBits += 8 * uint16(len(m.BdtEntries))
	}

	return lengthInBits
}

func (m *BVLCReadBroadcastDistributionTableAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCReadBroadcastDistributionTableAckParse(readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (*BVLCReadBroadcastDistributionTableAck, error) {
	if pullErr := readBuffer.PullContext("BVLCReadBroadcastDistributionTableAck"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos
	// Byte Array field (bdtEntries)
	numberOfBytesbdtEntries := int(bvlcPayloadLength)
	bdtEntries, _readArrayErr := readBuffer.ReadByteArray("bdtEntries", numberOfBytesbdtEntries)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'bdtEntries' field")
	}

	if closeErr := readBuffer.CloseContext("BVLCReadBroadcastDistributionTableAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCReadBroadcastDistributionTableAck{
		BdtEntries: bdtEntries,
		BVLC:       &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child, nil
}

func (m *BVLCReadBroadcastDistributionTableAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadBroadcastDistributionTableAck"); pushErr != nil {
			return pushErr
		}

		// Array Field (bdtEntries)
		if m.BdtEntries != nil {
			// Byte Array field (bdtEntries)
			_writeArrayErr := writeBuffer.WriteByteArray("bdtEntries", m.BdtEntries)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'bdtEntries' field")
			}
		}

		if popErr := writeBuffer.PopContext("BVLCReadBroadcastDistributionTableAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCReadBroadcastDistributionTableAck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
