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

// SecurityDataRequestZoneName is the corresponding interface of SecurityDataRequestZoneName
type SecurityDataRequestZoneName interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
}

// SecurityDataRequestZoneNameExactly can be used when we want exactly this type and not a type which fulfills SecurityDataRequestZoneName.
// This is useful for switch cases.
type SecurityDataRequestZoneNameExactly interface {
	SecurityDataRequestZoneName
	isSecurityDataRequestZoneName() bool
}

// _SecurityDataRequestZoneName is the data-structure of this message
type _SecurityDataRequestZoneName struct {
	*_SecurityData
	ZoneNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataRequestZoneName) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataRequestZoneName) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataRequestZoneName) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataRequestZoneName factory function for _SecurityDataRequestZoneName
func NewSecurityDataRequestZoneName(zoneNumber uint8, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataRequestZoneName {
	_result := &_SecurityDataRequestZoneName{
		ZoneNumber:    zoneNumber,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataRequestZoneName(structType interface{}) SecurityDataRequestZoneName {
	if casted, ok := structType.(SecurityDataRequestZoneName); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataRequestZoneName); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataRequestZoneName) GetTypeName() string {
	return "SecurityDataRequestZoneName"
}

func (m *_SecurityDataRequestZoneName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataRequestZoneName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityDataRequestZoneName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataRequestZoneNameParse(readBuffer utils.ReadBuffer) (SecurityDataRequestZoneName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataRequestZoneName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataRequestZoneName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneNumber)
	_zoneNumber, _zoneNumberErr := readBuffer.ReadUint8("zoneNumber", 8)
	if _zoneNumberErr != nil {
		return nil, errors.Wrap(_zoneNumberErr, "Error parsing 'zoneNumber' field of SecurityDataRequestZoneName")
	}
	zoneNumber := _zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataRequestZoneName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataRequestZoneName")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataRequestZoneName{
		_SecurityData: &_SecurityData{},
		ZoneNumber:    zoneNumber,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataRequestZoneName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataRequestZoneName) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataRequestZoneName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataRequestZoneName")
		}

		// Simple Field (zoneNumber)
		zoneNumber := uint8(m.GetZoneNumber())
		_zoneNumberErr := writeBuffer.WriteUint8("zoneNumber", 8, (zoneNumber))
		if _zoneNumberErr != nil {
			return errors.Wrap(_zoneNumberErr, "Error serializing 'zoneNumber' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataRequestZoneName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataRequestZoneName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataRequestZoneName) isSecurityDataRequestZoneName() bool {
	return true
}

func (m *_SecurityDataRequestZoneName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
