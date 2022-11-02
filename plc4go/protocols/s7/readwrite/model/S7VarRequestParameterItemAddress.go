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

// S7VarRequestParameterItemAddress is the corresponding interface of S7VarRequestParameterItemAddress
type S7VarRequestParameterItemAddress interface {
	utils.LengthAware
	utils.Serializable
	S7VarRequestParameterItem
	// GetAddress returns Address (property field)
	GetAddress() S7Address
}

// S7VarRequestParameterItemAddressExactly can be used when we want exactly this type and not a type which fulfills S7VarRequestParameterItemAddress.
// This is useful for switch cases.
type S7VarRequestParameterItemAddressExactly interface {
	S7VarRequestParameterItemAddress
	isS7VarRequestParameterItemAddress() bool
}

// _S7VarRequestParameterItemAddress is the data-structure of this message
type _S7VarRequestParameterItemAddress struct {
	*_S7VarRequestParameterItem
	Address S7Address
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7VarRequestParameterItemAddress) GetItemType() uint8 {
	return 0x12
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7VarRequestParameterItemAddress) InitializeParent(parent S7VarRequestParameterItem) {}

func (m *_S7VarRequestParameterItemAddress) GetParent() S7VarRequestParameterItem {
	return m._S7VarRequestParameterItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7VarRequestParameterItemAddress) GetAddress() S7Address {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7VarRequestParameterItemAddress factory function for _S7VarRequestParameterItemAddress
func NewS7VarRequestParameterItemAddress(address S7Address) *_S7VarRequestParameterItemAddress {
	_result := &_S7VarRequestParameterItemAddress{
		Address:                    address,
		_S7VarRequestParameterItem: NewS7VarRequestParameterItem(),
	}
	_result._S7VarRequestParameterItem._S7VarRequestParameterItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7VarRequestParameterItemAddress(structType interface{}) S7VarRequestParameterItemAddress {
	if casted, ok := structType.(S7VarRequestParameterItemAddress); ok {
		return casted
	}
	if casted, ok := structType.(*S7VarRequestParameterItemAddress); ok {
		return *casted
	}
	return nil
}

func (m *_S7VarRequestParameterItemAddress) GetTypeName() string {
	return "S7VarRequestParameterItemAddress"
}

func (m *_S7VarRequestParameterItemAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7VarRequestParameterItemAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (address)
	lengthInBits += m.Address.GetLengthInBits()

	return lengthInBits
}

func (m *_S7VarRequestParameterItemAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7VarRequestParameterItemAddressParse(readBuffer utils.ReadBuffer) (S7VarRequestParameterItemAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7VarRequestParameterItemAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7VarRequestParameterItemAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (itemLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	itemLength, _itemLengthErr := readBuffer.ReadUint8("itemLength", 8)
	_ = itemLength
	if _itemLengthErr != nil {
		return nil, errors.Wrap(_itemLengthErr, "Error parsing 'itemLength' field of S7VarRequestParameterItemAddress")
	}

	// Simple Field (address)
	if pullErr := readBuffer.PullContext("address"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for address")
	}
	_address, _addressErr := S7AddressParse(readBuffer)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of S7VarRequestParameterItemAddress")
	}
	address := _address.(S7Address)
	if closeErr := readBuffer.CloseContext("address"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for address")
	}

	if closeErr := readBuffer.CloseContext("S7VarRequestParameterItemAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7VarRequestParameterItemAddress")
	}

	// Create a partially initialized instance
	_child := &_S7VarRequestParameterItemAddress{
		_S7VarRequestParameterItem: &_S7VarRequestParameterItem{},
		Address:                    address,
	}
	_child._S7VarRequestParameterItem._S7VarRequestParameterItemChildRequirements = _child
	return _child, nil
}

func (m *_S7VarRequestParameterItemAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7VarRequestParameterItemAddress) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7VarRequestParameterItemAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7VarRequestParameterItemAddress")
		}

		// Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		itemLength := uint8(m.GetAddress().GetLengthInBytes())
		_itemLengthErr := writeBuffer.WriteUint8("itemLength", 8, (itemLength))
		if _itemLengthErr != nil {
			return errors.Wrap(_itemLengthErr, "Error serializing 'itemLength' field")
		}

		// Simple Field (address)
		if pushErr := writeBuffer.PushContext("address"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for address")
		}
		_addressErr := writeBuffer.WriteSerializable(m.GetAddress())
		if popErr := writeBuffer.PopContext("address"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for address")
		}
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("S7VarRequestParameterItemAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7VarRequestParameterItemAddress")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7VarRequestParameterItemAddress) isS7VarRequestParameterItemAddress() bool {
	return true
}

func (m *_S7VarRequestParameterItemAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
