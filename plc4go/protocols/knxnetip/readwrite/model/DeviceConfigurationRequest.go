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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// DeviceConfigurationRequest is the corresponding interface of DeviceConfigurationRequest
type DeviceConfigurationRequest interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetDeviceConfigurationRequestDataBlock returns DeviceConfigurationRequestDataBlock (property field)
	GetDeviceConfigurationRequestDataBlock() DeviceConfigurationRequestDataBlock
	// GetCemi returns Cemi (property field)
	GetCemi() CEMI
}

// DeviceConfigurationRequestExactly can be used when we want exactly this type and not a type which fulfills DeviceConfigurationRequest.
// This is useful for switch cases.
type DeviceConfigurationRequestExactly interface {
	DeviceConfigurationRequest
	isDeviceConfigurationRequest() bool
}

// _DeviceConfigurationRequest is the data-structure of this message
type _DeviceConfigurationRequest struct {
	*_KnxNetIpMessage
	DeviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock
	Cemi                                CEMI

	// Arguments.
	TotalLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeviceConfigurationRequest) GetMsgType() uint16 {
	return 0x0310
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeviceConfigurationRequest) InitializeParent(parent KnxNetIpMessage) {}

func (m *_DeviceConfigurationRequest) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationRequest) GetDeviceConfigurationRequestDataBlock() DeviceConfigurationRequestDataBlock {
	return m.DeviceConfigurationRequestDataBlock
}

func (m *_DeviceConfigurationRequest) GetCemi() CEMI {
	return m.Cemi
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeviceConfigurationRequest factory function for _DeviceConfigurationRequest
func NewDeviceConfigurationRequest(deviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock, cemi CEMI, totalLength uint16) *_DeviceConfigurationRequest {
	_result := &_DeviceConfigurationRequest{
		DeviceConfigurationRequestDataBlock: deviceConfigurationRequestDataBlock,
		Cemi:                                cemi,
		_KnxNetIpMessage:                    NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationRequest(structType interface{}) DeviceConfigurationRequest {
	if casted, ok := structType.(DeviceConfigurationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationRequest) GetTypeName() string {
	return "DeviceConfigurationRequest"
}

func (m *_DeviceConfigurationRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_DeviceConfigurationRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (deviceConfigurationRequestDataBlock)
	lengthInBits += m.DeviceConfigurationRequestDataBlock.GetLengthInBits()

	// Simple field (cemi)
	lengthInBits += m.Cemi.GetLengthInBits()

	return lengthInBits
}

func (m *_DeviceConfigurationRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DeviceConfigurationRequestParse(readBuffer utils.ReadBuffer, totalLength uint16) (DeviceConfigurationRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceConfigurationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deviceConfigurationRequestDataBlock)
	if pullErr := readBuffer.PullContext("deviceConfigurationRequestDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deviceConfigurationRequestDataBlock")
	}
	_deviceConfigurationRequestDataBlock, _deviceConfigurationRequestDataBlockErr := DeviceConfigurationRequestDataBlockParse(readBuffer)
	if _deviceConfigurationRequestDataBlockErr != nil {
		return nil, errors.Wrap(_deviceConfigurationRequestDataBlockErr, "Error parsing 'deviceConfigurationRequestDataBlock' field of DeviceConfigurationRequest")
	}
	deviceConfigurationRequestDataBlock := _deviceConfigurationRequestDataBlock.(DeviceConfigurationRequestDataBlock)
	if closeErr := readBuffer.CloseContext("deviceConfigurationRequestDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deviceConfigurationRequestDataBlock")
	}

	// Simple Field (cemi)
	if pullErr := readBuffer.PullContext("cemi"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for cemi")
	}
	_cemi, _cemiErr := CEMIParse(readBuffer, uint16(uint16(totalLength)-uint16((uint16(uint16(6))+uint16(deviceConfigurationRequestDataBlock.GetLengthInBytes())))))
	if _cemiErr != nil {
		return nil, errors.Wrap(_cemiErr, "Error parsing 'cemi' field of DeviceConfigurationRequest")
	}
	cemi := _cemi.(CEMI)
	if closeErr := readBuffer.CloseContext("cemi"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for cemi")
	}

	if closeErr := readBuffer.CloseContext("DeviceConfigurationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationRequest")
	}

	// Create a partially initialized instance
	_child := &_DeviceConfigurationRequest{
		DeviceConfigurationRequestDataBlock: deviceConfigurationRequestDataBlock,
		Cemi:                                cemi,
		_KnxNetIpMessage:                    &_KnxNetIpMessage{},
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_DeviceConfigurationRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeviceConfigurationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationRequest")
		}

		// Simple Field (deviceConfigurationRequestDataBlock)
		if pushErr := writeBuffer.PushContext("deviceConfigurationRequestDataBlock"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceConfigurationRequestDataBlock")
		}
		_deviceConfigurationRequestDataBlockErr := writeBuffer.WriteSerializable(m.GetDeviceConfigurationRequestDataBlock())
		if popErr := writeBuffer.PopContext("deviceConfigurationRequestDataBlock"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceConfigurationRequestDataBlock")
		}
		if _deviceConfigurationRequestDataBlockErr != nil {
			return errors.Wrap(_deviceConfigurationRequestDataBlockErr, "Error serializing 'deviceConfigurationRequestDataBlock' field")
		}

		// Simple Field (cemi)
		if pushErr := writeBuffer.PushContext("cemi"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for cemi")
		}
		_cemiErr := writeBuffer.WriteSerializable(m.GetCemi())
		if popErr := writeBuffer.PopContext("cemi"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for cemi")
		}
		if _cemiErr != nil {
			return errors.Wrap(_cemiErr, "Error serializing 'cemi' field")
		}

		if popErr := writeBuffer.PopContext("DeviceConfigurationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeviceConfigurationRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_DeviceConfigurationRequest) isDeviceConfigurationRequest() bool {
	return true
}

func (m *_DeviceConfigurationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
