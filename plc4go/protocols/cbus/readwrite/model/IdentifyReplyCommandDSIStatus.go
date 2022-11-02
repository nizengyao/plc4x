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

// IdentifyReplyCommandDSIStatus is the corresponding interface of IdentifyReplyCommandDSIStatus
type IdentifyReplyCommandDSIStatus interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetChannelStatus1 returns ChannelStatus1 (property field)
	GetChannelStatus1() ChannelStatus
	// GetChannelStatus2 returns ChannelStatus2 (property field)
	GetChannelStatus2() ChannelStatus
	// GetChannelStatus3 returns ChannelStatus3 (property field)
	GetChannelStatus3() ChannelStatus
	// GetChannelStatus4 returns ChannelStatus4 (property field)
	GetChannelStatus4() ChannelStatus
	// GetChannelStatus5 returns ChannelStatus5 (property field)
	GetChannelStatus5() ChannelStatus
	// GetChannelStatus6 returns ChannelStatus6 (property field)
	GetChannelStatus6() ChannelStatus
	// GetChannelStatus7 returns ChannelStatus7 (property field)
	GetChannelStatus7() ChannelStatus
	// GetChannelStatus8 returns ChannelStatus8 (property field)
	GetChannelStatus8() ChannelStatus
	// GetUnitStatus returns UnitStatus (property field)
	GetUnitStatus() UnitStatus
	// GetDimmingUCRevisionNumber returns DimmingUCRevisionNumber (property field)
	GetDimmingUCRevisionNumber() byte
}

// IdentifyReplyCommandDSIStatusExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandDSIStatus.
// This is useful for switch cases.
type IdentifyReplyCommandDSIStatusExactly interface {
	IdentifyReplyCommandDSIStatus
	isIdentifyReplyCommandDSIStatus() bool
}

// _IdentifyReplyCommandDSIStatus is the data-structure of this message
type _IdentifyReplyCommandDSIStatus struct {
	*_IdentifyReplyCommand
	ChannelStatus1          ChannelStatus
	ChannelStatus2          ChannelStatus
	ChannelStatus3          ChannelStatus
	ChannelStatus4          ChannelStatus
	ChannelStatus5          ChannelStatus
	ChannelStatus6          ChannelStatus
	ChannelStatus7          ChannelStatus
	ChannelStatus8          ChannelStatus
	UnitStatus              UnitStatus
	DimmingUCRevisionNumber byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandDSIStatus) GetAttribute() Attribute {
	return Attribute_DSIStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandDSIStatus) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandDSIStatus) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus1() ChannelStatus {
	return m.ChannelStatus1
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus2() ChannelStatus {
	return m.ChannelStatus2
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus3() ChannelStatus {
	return m.ChannelStatus3
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus4() ChannelStatus {
	return m.ChannelStatus4
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus5() ChannelStatus {
	return m.ChannelStatus5
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus6() ChannelStatus {
	return m.ChannelStatus6
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus7() ChannelStatus {
	return m.ChannelStatus7
}

func (m *_IdentifyReplyCommandDSIStatus) GetChannelStatus8() ChannelStatus {
	return m.ChannelStatus8
}

func (m *_IdentifyReplyCommandDSIStatus) GetUnitStatus() UnitStatus {
	return m.UnitStatus
}

func (m *_IdentifyReplyCommandDSIStatus) GetDimmingUCRevisionNumber() byte {
	return m.DimmingUCRevisionNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandDSIStatus factory function for _IdentifyReplyCommandDSIStatus
func NewIdentifyReplyCommandDSIStatus(channelStatus1 ChannelStatus, channelStatus2 ChannelStatus, channelStatus3 ChannelStatus, channelStatus4 ChannelStatus, channelStatus5 ChannelStatus, channelStatus6 ChannelStatus, channelStatus7 ChannelStatus, channelStatus8 ChannelStatus, unitStatus UnitStatus, dimmingUCRevisionNumber byte, numBytes uint8) *_IdentifyReplyCommandDSIStatus {
	_result := &_IdentifyReplyCommandDSIStatus{
		ChannelStatus1:          channelStatus1,
		ChannelStatus2:          channelStatus2,
		ChannelStatus3:          channelStatus3,
		ChannelStatus4:          channelStatus4,
		ChannelStatus5:          channelStatus5,
		ChannelStatus6:          channelStatus6,
		ChannelStatus7:          channelStatus7,
		ChannelStatus8:          channelStatus8,
		UnitStatus:              unitStatus,
		DimmingUCRevisionNumber: dimmingUCRevisionNumber,
		_IdentifyReplyCommand:   NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandDSIStatus(structType interface{}) IdentifyReplyCommandDSIStatus {
	if casted, ok := structType.(IdentifyReplyCommandDSIStatus); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandDSIStatus); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandDSIStatus) GetTypeName() string {
	return "IdentifyReplyCommandDSIStatus"
}

func (m *_IdentifyReplyCommandDSIStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandDSIStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (channelStatus1)
	lengthInBits += 8

	// Simple field (channelStatus2)
	lengthInBits += 8

	// Simple field (channelStatus3)
	lengthInBits += 8

	// Simple field (channelStatus4)
	lengthInBits += 8

	// Simple field (channelStatus5)
	lengthInBits += 8

	// Simple field (channelStatus6)
	lengthInBits += 8

	// Simple field (channelStatus7)
	lengthInBits += 8

	// Simple field (channelStatus8)
	lengthInBits += 8

	// Simple field (unitStatus)
	lengthInBits += 8

	// Simple field (dimmingUCRevisionNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_IdentifyReplyCommandDSIStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandDSIStatusParse(readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandDSIStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandDSIStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandDSIStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (channelStatus1)
	if pullErr := readBuffer.PullContext("channelStatus1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelStatus1")
	}
	_channelStatus1, _channelStatus1Err := ChannelStatusParse(readBuffer)
	if _channelStatus1Err != nil {
		return nil, errors.Wrap(_channelStatus1Err, "Error parsing 'channelStatus1' field of IdentifyReplyCommandDSIStatus")
	}
	channelStatus1 := _channelStatus1
	if closeErr := readBuffer.CloseContext("channelStatus1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelStatus1")
	}

	// Simple Field (channelStatus2)
	if pullErr := readBuffer.PullContext("channelStatus2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelStatus2")
	}
	_channelStatus2, _channelStatus2Err := ChannelStatusParse(readBuffer)
	if _channelStatus2Err != nil {
		return nil, errors.Wrap(_channelStatus2Err, "Error parsing 'channelStatus2' field of IdentifyReplyCommandDSIStatus")
	}
	channelStatus2 := _channelStatus2
	if closeErr := readBuffer.CloseContext("channelStatus2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelStatus2")
	}

	// Simple Field (channelStatus3)
	if pullErr := readBuffer.PullContext("channelStatus3"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelStatus3")
	}
	_channelStatus3, _channelStatus3Err := ChannelStatusParse(readBuffer)
	if _channelStatus3Err != nil {
		return nil, errors.Wrap(_channelStatus3Err, "Error parsing 'channelStatus3' field of IdentifyReplyCommandDSIStatus")
	}
	channelStatus3 := _channelStatus3
	if closeErr := readBuffer.CloseContext("channelStatus3"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelStatus3")
	}

	// Simple Field (channelStatus4)
	if pullErr := readBuffer.PullContext("channelStatus4"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelStatus4")
	}
	_channelStatus4, _channelStatus4Err := ChannelStatusParse(readBuffer)
	if _channelStatus4Err != nil {
		return nil, errors.Wrap(_channelStatus4Err, "Error parsing 'channelStatus4' field of IdentifyReplyCommandDSIStatus")
	}
	channelStatus4 := _channelStatus4
	if closeErr := readBuffer.CloseContext("channelStatus4"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelStatus4")
	}

	// Simple Field (channelStatus5)
	if pullErr := readBuffer.PullContext("channelStatus5"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelStatus5")
	}
	_channelStatus5, _channelStatus5Err := ChannelStatusParse(readBuffer)
	if _channelStatus5Err != nil {
		return nil, errors.Wrap(_channelStatus5Err, "Error parsing 'channelStatus5' field of IdentifyReplyCommandDSIStatus")
	}
	channelStatus5 := _channelStatus5
	if closeErr := readBuffer.CloseContext("channelStatus5"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelStatus5")
	}

	// Simple Field (channelStatus6)
	if pullErr := readBuffer.PullContext("channelStatus6"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelStatus6")
	}
	_channelStatus6, _channelStatus6Err := ChannelStatusParse(readBuffer)
	if _channelStatus6Err != nil {
		return nil, errors.Wrap(_channelStatus6Err, "Error parsing 'channelStatus6' field of IdentifyReplyCommandDSIStatus")
	}
	channelStatus6 := _channelStatus6
	if closeErr := readBuffer.CloseContext("channelStatus6"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelStatus6")
	}

	// Simple Field (channelStatus7)
	if pullErr := readBuffer.PullContext("channelStatus7"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelStatus7")
	}
	_channelStatus7, _channelStatus7Err := ChannelStatusParse(readBuffer)
	if _channelStatus7Err != nil {
		return nil, errors.Wrap(_channelStatus7Err, "Error parsing 'channelStatus7' field of IdentifyReplyCommandDSIStatus")
	}
	channelStatus7 := _channelStatus7
	if closeErr := readBuffer.CloseContext("channelStatus7"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelStatus7")
	}

	// Simple Field (channelStatus8)
	if pullErr := readBuffer.PullContext("channelStatus8"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelStatus8")
	}
	_channelStatus8, _channelStatus8Err := ChannelStatusParse(readBuffer)
	if _channelStatus8Err != nil {
		return nil, errors.Wrap(_channelStatus8Err, "Error parsing 'channelStatus8' field of IdentifyReplyCommandDSIStatus")
	}
	channelStatus8 := _channelStatus8
	if closeErr := readBuffer.CloseContext("channelStatus8"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelStatus8")
	}

	// Simple Field (unitStatus)
	if pullErr := readBuffer.PullContext("unitStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unitStatus")
	}
	_unitStatus, _unitStatusErr := UnitStatusParse(readBuffer)
	if _unitStatusErr != nil {
		return nil, errors.Wrap(_unitStatusErr, "Error parsing 'unitStatus' field of IdentifyReplyCommandDSIStatus")
	}
	unitStatus := _unitStatus
	if closeErr := readBuffer.CloseContext("unitStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unitStatus")
	}

	// Simple Field (dimmingUCRevisionNumber)
	_dimmingUCRevisionNumber, _dimmingUCRevisionNumberErr := readBuffer.ReadByte("dimmingUCRevisionNumber")
	if _dimmingUCRevisionNumberErr != nil {
		return nil, errors.Wrap(_dimmingUCRevisionNumberErr, "Error parsing 'dimmingUCRevisionNumber' field of IdentifyReplyCommandDSIStatus")
	}
	dimmingUCRevisionNumber := _dimmingUCRevisionNumber

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandDSIStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandDSIStatus")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandDSIStatus{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		ChannelStatus1:          channelStatus1,
		ChannelStatus2:          channelStatus2,
		ChannelStatus3:          channelStatus3,
		ChannelStatus4:          channelStatus4,
		ChannelStatus5:          channelStatus5,
		ChannelStatus6:          channelStatus6,
		ChannelStatus7:          channelStatus7,
		ChannelStatus8:          channelStatus8,
		UnitStatus:              unitStatus,
		DimmingUCRevisionNumber: dimmingUCRevisionNumber,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandDSIStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandDSIStatus) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandDSIStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandDSIStatus")
		}

		// Simple Field (channelStatus1)
		if pushErr := writeBuffer.PushContext("channelStatus1"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for channelStatus1")
		}
		_channelStatus1Err := writeBuffer.WriteSerializable(m.GetChannelStatus1())
		if popErr := writeBuffer.PopContext("channelStatus1"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for channelStatus1")
		}
		if _channelStatus1Err != nil {
			return errors.Wrap(_channelStatus1Err, "Error serializing 'channelStatus1' field")
		}

		// Simple Field (channelStatus2)
		if pushErr := writeBuffer.PushContext("channelStatus2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for channelStatus2")
		}
		_channelStatus2Err := writeBuffer.WriteSerializable(m.GetChannelStatus2())
		if popErr := writeBuffer.PopContext("channelStatus2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for channelStatus2")
		}
		if _channelStatus2Err != nil {
			return errors.Wrap(_channelStatus2Err, "Error serializing 'channelStatus2' field")
		}

		// Simple Field (channelStatus3)
		if pushErr := writeBuffer.PushContext("channelStatus3"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for channelStatus3")
		}
		_channelStatus3Err := writeBuffer.WriteSerializable(m.GetChannelStatus3())
		if popErr := writeBuffer.PopContext("channelStatus3"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for channelStatus3")
		}
		if _channelStatus3Err != nil {
			return errors.Wrap(_channelStatus3Err, "Error serializing 'channelStatus3' field")
		}

		// Simple Field (channelStatus4)
		if pushErr := writeBuffer.PushContext("channelStatus4"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for channelStatus4")
		}
		_channelStatus4Err := writeBuffer.WriteSerializable(m.GetChannelStatus4())
		if popErr := writeBuffer.PopContext("channelStatus4"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for channelStatus4")
		}
		if _channelStatus4Err != nil {
			return errors.Wrap(_channelStatus4Err, "Error serializing 'channelStatus4' field")
		}

		// Simple Field (channelStatus5)
		if pushErr := writeBuffer.PushContext("channelStatus5"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for channelStatus5")
		}
		_channelStatus5Err := writeBuffer.WriteSerializable(m.GetChannelStatus5())
		if popErr := writeBuffer.PopContext("channelStatus5"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for channelStatus5")
		}
		if _channelStatus5Err != nil {
			return errors.Wrap(_channelStatus5Err, "Error serializing 'channelStatus5' field")
		}

		// Simple Field (channelStatus6)
		if pushErr := writeBuffer.PushContext("channelStatus6"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for channelStatus6")
		}
		_channelStatus6Err := writeBuffer.WriteSerializable(m.GetChannelStatus6())
		if popErr := writeBuffer.PopContext("channelStatus6"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for channelStatus6")
		}
		if _channelStatus6Err != nil {
			return errors.Wrap(_channelStatus6Err, "Error serializing 'channelStatus6' field")
		}

		// Simple Field (channelStatus7)
		if pushErr := writeBuffer.PushContext("channelStatus7"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for channelStatus7")
		}
		_channelStatus7Err := writeBuffer.WriteSerializable(m.GetChannelStatus7())
		if popErr := writeBuffer.PopContext("channelStatus7"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for channelStatus7")
		}
		if _channelStatus7Err != nil {
			return errors.Wrap(_channelStatus7Err, "Error serializing 'channelStatus7' field")
		}

		// Simple Field (channelStatus8)
		if pushErr := writeBuffer.PushContext("channelStatus8"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for channelStatus8")
		}
		_channelStatus8Err := writeBuffer.WriteSerializable(m.GetChannelStatus8())
		if popErr := writeBuffer.PopContext("channelStatus8"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for channelStatus8")
		}
		if _channelStatus8Err != nil {
			return errors.Wrap(_channelStatus8Err, "Error serializing 'channelStatus8' field")
		}

		// Simple Field (unitStatus)
		if pushErr := writeBuffer.PushContext("unitStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unitStatus")
		}
		_unitStatusErr := writeBuffer.WriteSerializable(m.GetUnitStatus())
		if popErr := writeBuffer.PopContext("unitStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unitStatus")
		}
		if _unitStatusErr != nil {
			return errors.Wrap(_unitStatusErr, "Error serializing 'unitStatus' field")
		}

		// Simple Field (dimmingUCRevisionNumber)
		dimmingUCRevisionNumber := byte(m.GetDimmingUCRevisionNumber())
		_dimmingUCRevisionNumberErr := writeBuffer.WriteByte("dimmingUCRevisionNumber", (dimmingUCRevisionNumber))
		if _dimmingUCRevisionNumberErr != nil {
			return errors.Wrap(_dimmingUCRevisionNumberErr, "Error serializing 'dimmingUCRevisionNumber' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandDSIStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandDSIStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandDSIStatus) isIdentifyReplyCommandDSIStatus() bool {
	return true
}

func (m *_IdentifyReplyCommandDSIStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
