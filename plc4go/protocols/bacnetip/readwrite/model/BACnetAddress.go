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

// BACnetAddress is the corresponding interface of BACnetAddress
type BACnetAddress interface {
	utils.LengthAware
	utils.Serializable
	// GetNetworkNumber returns NetworkNumber (property field)
	GetNetworkNumber() BACnetApplicationTagUnsignedInteger
	// GetMacAddress returns MacAddress (property field)
	GetMacAddress() BACnetApplicationTagOctetString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetIsLocalNetwork returns IsLocalNetwork (virtual field)
	GetIsLocalNetwork() bool
	// GetIsBroadcast returns IsBroadcast (virtual field)
	GetIsBroadcast() bool
}

// BACnetAddressExactly can be used when we want exactly this type and not a type which fulfills BACnetAddress.
// This is useful for switch cases.
type BACnetAddressExactly interface {
	BACnetAddress
	isBACnetAddress() bool
}

// _BACnetAddress is the data-structure of this message
type _BACnetAddress struct {
	NetworkNumber BACnetApplicationTagUnsignedInteger
	MacAddress    BACnetApplicationTagOctetString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAddress) GetNetworkNumber() BACnetApplicationTagUnsignedInteger {
	return m.NetworkNumber
}

func (m *_BACnetAddress) GetMacAddress() BACnetApplicationTagOctetString {
	return m.MacAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetAddress) GetZero() uint64 {
	return uint64(uint64(0))
}

func (m *_BACnetAddress) GetIsLocalNetwork() bool {
	return bool(bool((m.GetNetworkNumber().GetActualValue()) == (m.GetZero())))
}

func (m *_BACnetAddress) GetIsBroadcast() bool {
	return bool(bool((m.GetMacAddress().GetActualLength()) == (0)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAddress factory function for _BACnetAddress
func NewBACnetAddress(networkNumber BACnetApplicationTagUnsignedInteger, macAddress BACnetApplicationTagOctetString) *_BACnetAddress {
	return &_BACnetAddress{NetworkNumber: networkNumber, MacAddress: macAddress}
}

// Deprecated: use the interface for direct cast
func CastBACnetAddress(structType interface{}) BACnetAddress {
	if casted, ok := structType.(BACnetAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAddress) GetTypeName() string {
	return "BACnetAddress"
}

func (m *_BACnetAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (networkNumber)
	lengthInBits += m.NetworkNumber.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (macAddress)
	lengthInBits += m.MacAddress.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAddressParse(readBuffer utils.ReadBuffer) (BACnetAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkNumber)
	if pullErr := readBuffer.PullContext("networkNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkNumber")
	}
	_networkNumber, _networkNumberErr := BACnetApplicationTagParse(readBuffer)
	if _networkNumberErr != nil {
		return nil, errors.Wrap(_networkNumberErr, "Error parsing 'networkNumber' field of BACnetAddress")
	}
	networkNumber := _networkNumber.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("networkNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkNumber")
	}

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Virtual field
	_isLocalNetwork := bool((networkNumber.GetActualValue()) == (zero))
	isLocalNetwork := bool(_isLocalNetwork)
	_ = isLocalNetwork

	// Simple Field (macAddress)
	if pullErr := readBuffer.PullContext("macAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for macAddress")
	}
	_macAddress, _macAddressErr := BACnetApplicationTagParse(readBuffer)
	if _macAddressErr != nil {
		return nil, errors.Wrap(_macAddressErr, "Error parsing 'macAddress' field of BACnetAddress")
	}
	macAddress := _macAddress.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("macAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for macAddress")
	}

	// Virtual field
	_isBroadcast := bool((macAddress.GetActualLength()) == (0))
	isBroadcast := bool(_isBroadcast)
	_ = isBroadcast

	if closeErr := readBuffer.CloseContext("BACnetAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAddress")
	}

	// Create the instance
	return &_BACnetAddress{
		NetworkNumber: networkNumber,
		MacAddress:    macAddress,
	}, nil
}

func (m *_BACnetAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAddress) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAddress")
	}

	// Simple Field (networkNumber)
	if pushErr := writeBuffer.PushContext("networkNumber"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for networkNumber")
	}
	_networkNumberErr := writeBuffer.WriteSerializable(m.GetNetworkNumber())
	if popErr := writeBuffer.PopContext("networkNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for networkNumber")
	}
	if _networkNumberErr != nil {
		return errors.Wrap(_networkNumberErr, "Error serializing 'networkNumber' field")
	}
	// Virtual field
	if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
		return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
	}
	// Virtual field
	if _isLocalNetworkErr := writeBuffer.WriteVirtual("isLocalNetwork", m.GetIsLocalNetwork()); _isLocalNetworkErr != nil {
		return errors.Wrap(_isLocalNetworkErr, "Error serializing 'isLocalNetwork' field")
	}

	// Simple Field (macAddress)
	if pushErr := writeBuffer.PushContext("macAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for macAddress")
	}
	_macAddressErr := writeBuffer.WriteSerializable(m.GetMacAddress())
	if popErr := writeBuffer.PopContext("macAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for macAddress")
	}
	if _macAddressErr != nil {
		return errors.Wrap(_macAddressErr, "Error serializing 'macAddress' field")
	}
	// Virtual field
	if _isBroadcastErr := writeBuffer.WriteVirtual("isBroadcast", m.GetIsBroadcast()); _isBroadcastErr != nil {
		return errors.Wrap(_isBroadcastErr, "Error serializing 'isBroadcast' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAddress")
	}
	return nil
}

func (m *_BACnetAddress) isBACnetAddress() bool {
	return true
}

func (m *_BACnetAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
