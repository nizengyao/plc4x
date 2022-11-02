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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// CALReplyLong is the corresponding interface of CALReplyLong
type CALReplyLong interface {
	utils.LengthAware
	utils.Serializable
	CALReply
	// GetTerminatingByte returns TerminatingByte (property field)
	GetTerminatingByte() uint32
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetSerialInterfaceAddress returns SerialInterfaceAddress (property field)
	GetSerialInterfaceAddress() SerialInterfaceAddress
	// GetReservedByte returns ReservedByte (property field)
	GetReservedByte() *byte
	// GetReplyNetwork returns ReplyNetwork (property field)
	GetReplyNetwork() ReplyNetwork
	// GetIsUnitAddress returns IsUnitAddress (virtual field)
	GetIsUnitAddress() bool
}

// CALReplyLongExactly can be used when we want exactly this type and not a type which fulfills CALReplyLong.
// This is useful for switch cases.
type CALReplyLongExactly interface {
	CALReplyLong
	isCALReplyLong() bool
}

// _CALReplyLong is the data-structure of this message
type _CALReplyLong struct {
	*_CALReply
	TerminatingByte        uint32
	UnitAddress            UnitAddress
	BridgeAddress          BridgeAddress
	SerialInterfaceAddress SerialInterfaceAddress
	ReservedByte           *byte
	ReplyNetwork           ReplyNetwork
	// Reserved Fields
	reservedField0 *byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALReplyLong) InitializeParent(parent CALReply, calType byte, calData CALData) {
	m.CalType = calType
	m.CalData = calData
}

func (m *_CALReplyLong) GetParent() CALReply {
	return m._CALReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALReplyLong) GetTerminatingByte() uint32 {
	return m.TerminatingByte
}

func (m *_CALReplyLong) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

func (m *_CALReplyLong) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_CALReplyLong) GetSerialInterfaceAddress() SerialInterfaceAddress {
	return m.SerialInterfaceAddress
}

func (m *_CALReplyLong) GetReservedByte() *byte {
	return m.ReservedByte
}

func (m *_CALReplyLong) GetReplyNetwork() ReplyNetwork {
	return m.ReplyNetwork
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_CALReplyLong) GetIsUnitAddress() bool {
	unitAddress := m.UnitAddress
	_ = unitAddress
	bridgeAddress := m.BridgeAddress
	_ = bridgeAddress
	reservedByte := m.ReservedByte
	_ = reservedByte
	replyNetwork := m.ReplyNetwork
	_ = replyNetwork
	return bool(bool((m.GetTerminatingByte() & 0xff) == (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALReplyLong factory function for _CALReplyLong
func NewCALReplyLong(terminatingByte uint32, unitAddress UnitAddress, bridgeAddress BridgeAddress, serialInterfaceAddress SerialInterfaceAddress, reservedByte *byte, replyNetwork ReplyNetwork, calType byte, calData CALData, cBusOptions CBusOptions, requestContext RequestContext) *_CALReplyLong {
	_result := &_CALReplyLong{
		TerminatingByte:        terminatingByte,
		UnitAddress:            unitAddress,
		BridgeAddress:          bridgeAddress,
		SerialInterfaceAddress: serialInterfaceAddress,
		ReservedByte:           reservedByte,
		ReplyNetwork:           replyNetwork,
		_CALReply:              NewCALReply(calType, calData, cBusOptions, requestContext),
	}
	_result._CALReply._CALReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALReplyLong(structType interface{}) CALReplyLong {
	if casted, ok := structType.(CALReplyLong); ok {
		return casted
	}
	if casted, ok := structType.(*CALReplyLong); ok {
		return *casted
	}
	return nil
}

func (m *_CALReplyLong) GetTypeName() string {
	return "CALReplyLong"
}

func (m *_CALReplyLong) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALReplyLong) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Optional Field (unitAddress)
	if m.UnitAddress != nil {
		lengthInBits += m.UnitAddress.GetLengthInBits()
	}

	// Optional Field (bridgeAddress)
	if m.BridgeAddress != nil {
		lengthInBits += m.BridgeAddress.GetLengthInBits()
	}

	// Simple field (serialInterfaceAddress)
	lengthInBits += m.SerialInterfaceAddress.GetLengthInBits()

	// Optional Field (reservedByte)
	if m.ReservedByte != nil {
		lengthInBits += 8
	}

	// Optional Field (replyNetwork)
	if m.ReplyNetwork != nil {
		lengthInBits += m.ReplyNetwork.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_CALReplyLong) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALReplyLongParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (CALReplyLong, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALReplyLong"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALReplyLong")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *byte
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CALReplyLong")
		}
		if reserved != byte(0x86) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": byte(0x86),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Peek Field (terminatingByte)
	currentPos = positionAware.GetPos()
	terminatingByte, _err := readBuffer.ReadUint32("terminatingByte", 24)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'terminatingByte' field of CALReplyLong")
	}

	readBuffer.Reset(currentPos)

	// Virtual field
	_isUnitAddress := bool((terminatingByte & 0xff) == (0x00))
	isUnitAddress := bool(_isUnitAddress)
	_ = isUnitAddress

	// Optional Field (unitAddress) (Can be skipped, if a given expression evaluates to false)
	var unitAddress UnitAddress = nil
	if isUnitAddress {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("unitAddress"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for unitAddress")
		}
		_val, _err := UnitAddressParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'unitAddress' field of CALReplyLong")
		default:
			unitAddress = _val.(UnitAddress)
			if closeErr := readBuffer.CloseContext("unitAddress"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for unitAddress")
			}
		}
	}

	// Optional Field (bridgeAddress) (Can be skipped, if a given expression evaluates to false)
	var bridgeAddress BridgeAddress = nil
	if !(isUnitAddress) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("bridgeAddress"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for bridgeAddress")
		}
		_val, _err := BridgeAddressParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'bridgeAddress' field of CALReplyLong")
		default:
			bridgeAddress = _val.(BridgeAddress)
			if closeErr := readBuffer.CloseContext("bridgeAddress"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for bridgeAddress")
			}
		}
	}

	// Simple Field (serialInterfaceAddress)
	if pullErr := readBuffer.PullContext("serialInterfaceAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serialInterfaceAddress")
	}
	_serialInterfaceAddress, _serialInterfaceAddressErr := SerialInterfaceAddressParse(readBuffer)
	if _serialInterfaceAddressErr != nil {
		return nil, errors.Wrap(_serialInterfaceAddressErr, "Error parsing 'serialInterfaceAddress' field of CALReplyLong")
	}
	serialInterfaceAddress := _serialInterfaceAddress.(SerialInterfaceAddress)
	if closeErr := readBuffer.CloseContext("serialInterfaceAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serialInterfaceAddress")
	}

	// Optional Field (reservedByte) (Can be skipped, if a given expression evaluates to false)
	var reservedByte *byte = nil
	if isUnitAddress {
		_val, _err := readBuffer.ReadByte("reservedByte")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reservedByte' field of CALReplyLong")
		}
		reservedByte = &_val
	}

	// Validation
	if !(bool(bool(isUnitAddress) && bool(bool((*reservedByte) == (0x00)))) || bool(!(isUnitAddress))) {
		return nil, errors.WithStack(utils.ParseValidationError{"wrong reservedByte"})
	}

	// Optional Field (replyNetwork) (Can be skipped, if a given expression evaluates to false)
	var replyNetwork ReplyNetwork = nil
	if !(isUnitAddress) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("replyNetwork"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for replyNetwork")
		}
		_val, _err := ReplyNetworkParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'replyNetwork' field of CALReplyLong")
		default:
			replyNetwork = _val.(ReplyNetwork)
			if closeErr := readBuffer.CloseContext("replyNetwork"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for replyNetwork")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("CALReplyLong"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALReplyLong")
	}

	// Create a partially initialized instance
	_child := &_CALReplyLong{
		_CALReply: &_CALReply{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
		TerminatingByte:        terminatingByte,
		UnitAddress:            unitAddress,
		BridgeAddress:          bridgeAddress,
		SerialInterfaceAddress: serialInterfaceAddress,
		ReservedByte:           reservedByte,
		ReplyNetwork:           replyNetwork,
		reservedField0:         reservedField0,
	}
	_child._CALReply._CALReplyChildRequirements = _child
	return _child, nil
}

func (m *_CALReplyLong) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALReplyLong) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALReplyLong"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALReplyLong")
		}

		// Reserved Field (reserved)
		{
			var reserved byte = byte(0x86)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": byte(0x86),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteByte("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}
		// Virtual field
		if _isUnitAddressErr := writeBuffer.WriteVirtual("isUnitAddress", m.GetIsUnitAddress()); _isUnitAddressErr != nil {
			return errors.Wrap(_isUnitAddressErr, "Error serializing 'isUnitAddress' field")
		}

		// Optional Field (unitAddress) (Can be skipped, if the value is null)
		var unitAddress UnitAddress = nil
		if m.GetUnitAddress() != nil {
			if pushErr := writeBuffer.PushContext("unitAddress"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for unitAddress")
			}
			unitAddress = m.GetUnitAddress()
			_unitAddressErr := writeBuffer.WriteSerializable(unitAddress)
			if popErr := writeBuffer.PopContext("unitAddress"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for unitAddress")
			}
			if _unitAddressErr != nil {
				return errors.Wrap(_unitAddressErr, "Error serializing 'unitAddress' field")
			}
		}

		// Optional Field (bridgeAddress) (Can be skipped, if the value is null)
		var bridgeAddress BridgeAddress = nil
		if m.GetBridgeAddress() != nil {
			if pushErr := writeBuffer.PushContext("bridgeAddress"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for bridgeAddress")
			}
			bridgeAddress = m.GetBridgeAddress()
			_bridgeAddressErr := writeBuffer.WriteSerializable(bridgeAddress)
			if popErr := writeBuffer.PopContext("bridgeAddress"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for bridgeAddress")
			}
			if _bridgeAddressErr != nil {
				return errors.Wrap(_bridgeAddressErr, "Error serializing 'bridgeAddress' field")
			}
		}

		// Simple Field (serialInterfaceAddress)
		if pushErr := writeBuffer.PushContext("serialInterfaceAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serialInterfaceAddress")
		}
		_serialInterfaceAddressErr := writeBuffer.WriteSerializable(m.GetSerialInterfaceAddress())
		if popErr := writeBuffer.PopContext("serialInterfaceAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serialInterfaceAddress")
		}
		if _serialInterfaceAddressErr != nil {
			return errors.Wrap(_serialInterfaceAddressErr, "Error serializing 'serialInterfaceAddress' field")
		}

		// Optional Field (reservedByte) (Can be skipped, if the value is null)
		var reservedByte *byte = nil
		if m.GetReservedByte() != nil {
			reservedByte = m.GetReservedByte()
			_reservedByteErr := writeBuffer.WriteByte("reservedByte", *(reservedByte))
			if _reservedByteErr != nil {
				return errors.Wrap(_reservedByteErr, "Error serializing 'reservedByte' field")
			}
		}

		// Optional Field (replyNetwork) (Can be skipped, if the value is null)
		var replyNetwork ReplyNetwork = nil
		if m.GetReplyNetwork() != nil {
			if pushErr := writeBuffer.PushContext("replyNetwork"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for replyNetwork")
			}
			replyNetwork = m.GetReplyNetwork()
			_replyNetworkErr := writeBuffer.WriteSerializable(replyNetwork)
			if popErr := writeBuffer.PopContext("replyNetwork"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for replyNetwork")
			}
			if _replyNetworkErr != nil {
				return errors.Wrap(_replyNetworkErr, "Error serializing 'replyNetwork' field")
			}
		}

		if popErr := writeBuffer.PopContext("CALReplyLong"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALReplyLong")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALReplyLong) isCALReplyLong() bool {
	return true
}

func (m *_CALReplyLong) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
