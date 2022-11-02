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

// AirConditioningDataSetHumidityUpperGuardLimit is the corresponding interface of AirConditioningDataSetHumidityUpperGuardLimit
type AirConditioningDataSetHumidityUpperGuardLimit interface {
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetLimit returns Limit (property field)
	GetLimit() HVACHumidity
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACHumidityModeAndFlags
}

// AirConditioningDataSetHumidityUpperGuardLimitExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataSetHumidityUpperGuardLimit.
// This is useful for switch cases.
type AirConditioningDataSetHumidityUpperGuardLimitExactly interface {
	AirConditioningDataSetHumidityUpperGuardLimit
	isAirConditioningDataSetHumidityUpperGuardLimit() bool
}

// _AirConditioningDataSetHumidityUpperGuardLimit is the data-structure of this message
type _AirConditioningDataSetHumidityUpperGuardLimit struct {
	*_AirConditioningData
	ZoneGroup        byte
	ZoneList         HVACZoneList
	Limit            HVACHumidity
	HvacModeAndFlags HVACHumidityModeAndFlags
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetLimit() HVACHumidity {
	return m.Limit
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetHvacModeAndFlags() HVACHumidityModeAndFlags {
	return m.HvacModeAndFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataSetHumidityUpperGuardLimit factory function for _AirConditioningDataSetHumidityUpperGuardLimit
func NewAirConditioningDataSetHumidityUpperGuardLimit(zoneGroup byte, zoneList HVACZoneList, limit HVACHumidity, hvacModeAndFlags HVACHumidityModeAndFlags, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataSetHumidityUpperGuardLimit {
	_result := &_AirConditioningDataSetHumidityUpperGuardLimit{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Limit:                limit,
		HvacModeAndFlags:     hvacModeAndFlags,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetHumidityUpperGuardLimit(structType interface{}) AirConditioningDataSetHumidityUpperGuardLimit {
	if casted, ok := structType.(AirConditioningDataSetHumidityUpperGuardLimit); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetHumidityUpperGuardLimit); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetTypeName() string {
	return "AirConditioningDataSetHumidityUpperGuardLimit"
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits()

	// Simple field (limit)
	lengthInBits += m.Limit.GetLengthInBits()

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits()

	return lengthInBits
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AirConditioningDataSetHumidityUpperGuardLimitParse(readBuffer utils.ReadBuffer) (AirConditioningDataSetHumidityUpperGuardLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetHumidityUpperGuardLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetHumidityUpperGuardLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
	_zoneGroup, _zoneGroupErr := readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataSetHumidityUpperGuardLimit")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
	_zoneList, _zoneListErr := HVACZoneListParse(readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataSetHumidityUpperGuardLimit")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (limit)
	if pullErr := readBuffer.PullContext("limit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for limit")
	}
	_limit, _limitErr := HVACHumidityParse(readBuffer)
	if _limitErr != nil {
		return nil, errors.Wrap(_limitErr, "Error parsing 'limit' field of AirConditioningDataSetHumidityUpperGuardLimit")
	}
	limit := _limit.(HVACHumidity)
	if closeErr := readBuffer.CloseContext("limit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for limit")
	}

	// Simple Field (hvacModeAndFlags)
	if pullErr := readBuffer.PullContext("hvacModeAndFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hvacModeAndFlags")
	}
	_hvacModeAndFlags, _hvacModeAndFlagsErr := HVACHumidityModeAndFlagsParse(readBuffer)
	if _hvacModeAndFlagsErr != nil {
		return nil, errors.Wrap(_hvacModeAndFlagsErr, "Error parsing 'hvacModeAndFlags' field of AirConditioningDataSetHumidityUpperGuardLimit")
	}
	hvacModeAndFlags := _hvacModeAndFlags.(HVACHumidityModeAndFlags)
	if closeErr := readBuffer.CloseContext("hvacModeAndFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hvacModeAndFlags")
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetHumidityUpperGuardLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetHumidityUpperGuardLimit")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataSetHumidityUpperGuardLimit{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Limit:                limit,
		HvacModeAndFlags:     hvacModeAndFlags,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetHumidityUpperGuardLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetHumidityUpperGuardLimit")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		// Simple Field (zoneList)
		if pushErr := writeBuffer.PushContext("zoneList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneList")
		}
		_zoneListErr := writeBuffer.WriteSerializable(m.GetZoneList())
		if popErr := writeBuffer.PopContext("zoneList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneList")
		}
		if _zoneListErr != nil {
			return errors.Wrap(_zoneListErr, "Error serializing 'zoneList' field")
		}

		// Simple Field (limit)
		if pushErr := writeBuffer.PushContext("limit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for limit")
		}
		_limitErr := writeBuffer.WriteSerializable(m.GetLimit())
		if popErr := writeBuffer.PopContext("limit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for limit")
		}
		if _limitErr != nil {
			return errors.Wrap(_limitErr, "Error serializing 'limit' field")
		}

		// Simple Field (hvacModeAndFlags)
		if pushErr := writeBuffer.PushContext("hvacModeAndFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hvacModeAndFlags")
		}
		_hvacModeAndFlagsErr := writeBuffer.WriteSerializable(m.GetHvacModeAndFlags())
		if popErr := writeBuffer.PopContext("hvacModeAndFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hvacModeAndFlags")
		}
		if _hvacModeAndFlagsErr != nil {
			return errors.Wrap(_hvacModeAndFlagsErr, "Error serializing 'hvacModeAndFlags' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetHumidityUpperGuardLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetHumidityUpperGuardLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) isAirConditioningDataSetHumidityUpperGuardLimit() bool {
	return true
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
