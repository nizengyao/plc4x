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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const DF1SymbolMessageFrame_MESSAGEEND uint8 = 0x10
const DF1SymbolMessageFrame_ENDTRANSACTION uint8 = 0x03

// DF1SymbolMessageFrame is the corresponding interface of DF1SymbolMessageFrame
type DF1SymbolMessageFrame interface {
	utils.LengthAware
	utils.Serializable
	DF1Symbol
	// GetDestinationAddress returns DestinationAddress (property field)
	GetDestinationAddress() uint8
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() uint8
	// GetCommand returns Command (property field)
	GetCommand() DF1Command
}

// DF1SymbolMessageFrameExactly can be used when we want exactly this type and not a type which fulfills DF1SymbolMessageFrame.
// This is useful for switch cases.
type DF1SymbolMessageFrameExactly interface {
	DF1SymbolMessageFrame
	isDF1SymbolMessageFrame() bool
}

// _DF1SymbolMessageFrame is the data-structure of this message
type _DF1SymbolMessageFrame struct {
	*_DF1Symbol
	DestinationAddress uint8
	SourceAddress      uint8
	Command            DF1Command
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1SymbolMessageFrame) GetSymbolType() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1SymbolMessageFrame) InitializeParent(parent DF1Symbol) {}

func (m *_DF1SymbolMessageFrame) GetParent() DF1Symbol {
	return m._DF1Symbol
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1SymbolMessageFrame) GetDestinationAddress() uint8 {
	return m.DestinationAddress
}

func (m *_DF1SymbolMessageFrame) GetSourceAddress() uint8 {
	return m.SourceAddress
}

func (m *_DF1SymbolMessageFrame) GetCommand() DF1Command {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_DF1SymbolMessageFrame) GetMessageEnd() uint8 {
	return DF1SymbolMessageFrame_MESSAGEEND
}

func (m *_DF1SymbolMessageFrame) GetEndTransaction() uint8 {
	return DF1SymbolMessageFrame_ENDTRANSACTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1SymbolMessageFrame factory function for _DF1SymbolMessageFrame
func NewDF1SymbolMessageFrame(destinationAddress uint8, sourceAddress uint8, command DF1Command) *_DF1SymbolMessageFrame {
	_result := &_DF1SymbolMessageFrame{
		DestinationAddress: destinationAddress,
		SourceAddress:      sourceAddress,
		Command:            command,
		_DF1Symbol:         NewDF1Symbol(),
	}
	_result._DF1Symbol._DF1SymbolChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1SymbolMessageFrame(structType interface{}) DF1SymbolMessageFrame {
	if casted, ok := structType.(DF1SymbolMessageFrame); ok {
		return casted
	}
	if casted, ok := structType.(*DF1SymbolMessageFrame); ok {
		return *casted
	}
	return nil
}

func (m *_DF1SymbolMessageFrame) GetTypeName() string {
	return "DF1SymbolMessageFrame"
}

func (m *_DF1SymbolMessageFrame) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_DF1SymbolMessageFrame) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationAddress)
	lengthInBits += 8

	// Simple field (sourceAddress)
	lengthInBits += 8

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits()

	// Const Field (messageEnd)
	lengthInBits += 8

	// Const Field (endTransaction)
	lengthInBits += 8

	// Checksum Field (checksum)
	lengthInBits += 16

	return lengthInBits
}

func (m *_DF1SymbolMessageFrame) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1SymbolMessageFrameParse(readBuffer utils.ReadBuffer) (DF1SymbolMessageFrame, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1SymbolMessageFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1SymbolMessageFrame")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationAddress)
	_destinationAddress, _destinationAddressErr := readBuffer.ReadUint8("destinationAddress", 8)
	if _destinationAddressErr != nil {
		return nil, errors.Wrap(_destinationAddressErr, "Error parsing 'destinationAddress' field of DF1SymbolMessageFrame")
	}
	destinationAddress := _destinationAddress

	// Simple Field (sourceAddress)
	_sourceAddress, _sourceAddressErr := readBuffer.ReadUint8("sourceAddress", 8)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field of DF1SymbolMessageFrame")
	}
	sourceAddress := _sourceAddress

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for command")
	}
	_command, _commandErr := DF1CommandParse(readBuffer)
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field of DF1SymbolMessageFrame")
	}
	command := _command.(DF1Command)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for command")
	}

	// Const Field (messageEnd)
	messageEnd, _messageEndErr := readBuffer.ReadUint8("messageEnd", 8)
	if _messageEndErr != nil {
		return nil, errors.Wrap(_messageEndErr, "Error parsing 'messageEnd' field of DF1SymbolMessageFrame")
	}
	if messageEnd != DF1SymbolMessageFrame_MESSAGEEND {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", DF1SymbolMessageFrame_MESSAGEEND) + " but got " + fmt.Sprintf("%d", messageEnd))
	}

	// Const Field (endTransaction)
	endTransaction, _endTransactionErr := readBuffer.ReadUint8("endTransaction", 8)
	if _endTransactionErr != nil {
		return nil, errors.Wrap(_endTransactionErr, "Error parsing 'endTransaction' field of DF1SymbolMessageFrame")
	}
	if endTransaction != DF1SymbolMessageFrame_ENDTRANSACTION {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", DF1SymbolMessageFrame_ENDTRANSACTION) + " but got " + fmt.Sprintf("%d", endTransaction))
	}

	// Checksum Field (checksum)
	{
		checksumRef, _checksumRefErr := readBuffer.ReadUint16("checksum", 16)
		if _checksumRefErr != nil {
			return nil, errors.Wrap(_checksumRefErr, "Error parsing 'checksum' field of DF1SymbolMessageFrame")
		}
		checksum, _checksumErr := CrcCheck(destinationAddress, sourceAddress, command)
		if _checksumErr != nil {
			return nil, errors.Wrap(_checksumErr, "Checksum verification failed")
		}
		if checksum != checksumRef {
			return nil, errors.Errorf("Checksum verification failed. Expected %x but got %x", checksumRef, checksum)
		}
	}

	if closeErr := readBuffer.CloseContext("DF1SymbolMessageFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1SymbolMessageFrame")
	}

	// Create a partially initialized instance
	_child := &_DF1SymbolMessageFrame{
		_DF1Symbol:         &_DF1Symbol{},
		DestinationAddress: destinationAddress,
		SourceAddress:      sourceAddress,
		Command:            command,
	}
	_child._DF1Symbol._DF1SymbolChildRequirements = _child
	return _child, nil
}

func (m *_DF1SymbolMessageFrame) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1SymbolMessageFrame) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1SymbolMessageFrame"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1SymbolMessageFrame")
		}

		// Simple Field (destinationAddress)
		destinationAddress := uint8(m.GetDestinationAddress())
		_destinationAddressErr := writeBuffer.WriteUint8("destinationAddress", 8, (destinationAddress))
		if _destinationAddressErr != nil {
			return errors.Wrap(_destinationAddressErr, "Error serializing 'destinationAddress' field")
		}

		// Simple Field (sourceAddress)
		sourceAddress := uint8(m.GetSourceAddress())
		_sourceAddressErr := writeBuffer.WriteUint8("sourceAddress", 8, (sourceAddress))
		if _sourceAddressErr != nil {
			return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
		}

		// Simple Field (command)
		if pushErr := writeBuffer.PushContext("command"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for command")
		}
		_commandErr := writeBuffer.WriteSerializable(m.GetCommand())
		if popErr := writeBuffer.PopContext("command"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for command")
		}
		if _commandErr != nil {
			return errors.Wrap(_commandErr, "Error serializing 'command' field")
		}

		// Const Field (messageEnd)
		_messageEndErr := writeBuffer.WriteUint8("messageEnd", 8, 0x10)
		if _messageEndErr != nil {
			return errors.Wrap(_messageEndErr, "Error serializing 'messageEnd' field")
		}

		// Const Field (endTransaction)
		_endTransactionErr := writeBuffer.WriteUint8("endTransaction", 8, 0x03)
		if _endTransactionErr != nil {
			return errors.Wrap(_endTransactionErr, "Error serializing 'endTransaction' field")
		}

		// Checksum Field (checksum) (Calculated)
		{
			_checksum, _checksumErr := CrcCheck(m.GetDestinationAddress(), m.GetSourceAddress(), m.GetCommand())
			if _checksumErr != nil {
				return errors.Wrap(_checksumErr, "Checksum calculation failed")
			}
			_checksumWriteErr := writeBuffer.WriteUint16("checksum", 16, (_checksum))
			if _checksumWriteErr != nil {
				return errors.Wrap(_checksumWriteErr, "Error serializing 'checksum' field")
			}
		}

		if popErr := writeBuffer.PopContext("DF1SymbolMessageFrame"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1SymbolMessageFrame")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_DF1SymbolMessageFrame) isDF1SymbolMessageFrame() bool {
	return true
}

func (m *_DF1SymbolMessageFrame) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
