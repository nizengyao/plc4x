//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AdsData struct {
	Child IAdsDataChild
}

// The corresponding interface
type IAdsData interface {
	CommandId() CommandId
	Response() bool
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IAdsDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IAdsData, serializeChildFunction func() error) error
	GetTypeName() string
}

type IAdsDataChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *AdsData)
	GetTypeName() string
	IAdsData
}

func NewAdsData() *AdsData {
	return &AdsData{}
}

func CastAdsData(structType interface{}) *AdsData {
	castFunc := func(typ interface{}) *AdsData {
		if casted, ok := typ.(AdsData); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsData); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsData) GetTypeName() string {
	return "AdsData"
}

func (m *AdsData) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsData) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *AdsData) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *AdsData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsDataParse(readBuffer utils.ReadBuffer, commandId *CommandId, response bool) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsData"); pullErr != nil {
		return nil, pullErr
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *AdsData
	var typeSwitchError error
	switch {
	case *commandId == CommandId_INVALID && response == false: // AdsInvalidRequest
		_parent, typeSwitchError = AdsInvalidRequestParse(readBuffer)
	case *commandId == CommandId_INVALID && response == true: // AdsInvalidResponse
		_parent, typeSwitchError = AdsInvalidResponseParse(readBuffer)
	case *commandId == CommandId_ADS_READ_DEVICE_INFO && response == false: // AdsReadDeviceInfoRequest
		_parent, typeSwitchError = AdsReadDeviceInfoRequestParse(readBuffer)
	case *commandId == CommandId_ADS_READ_DEVICE_INFO && response == true: // AdsReadDeviceInfoResponse
		_parent, typeSwitchError = AdsReadDeviceInfoResponseParse(readBuffer)
	case *commandId == CommandId_ADS_READ && response == false: // AdsReadRequest
		_parent, typeSwitchError = AdsReadRequestParse(readBuffer)
	case *commandId == CommandId_ADS_READ && response == true: // AdsReadResponse
		_parent, typeSwitchError = AdsReadResponseParse(readBuffer)
	case *commandId == CommandId_ADS_WRITE && response == false: // AdsWriteRequest
		_parent, typeSwitchError = AdsWriteRequestParse(readBuffer)
	case *commandId == CommandId_ADS_WRITE && response == true: // AdsWriteResponse
		_parent, typeSwitchError = AdsWriteResponseParse(readBuffer)
	case *commandId == CommandId_ADS_READ_STATE && response == false: // AdsReadStateRequest
		_parent, typeSwitchError = AdsReadStateRequestParse(readBuffer)
	case *commandId == CommandId_ADS_READ_STATE && response == true: // AdsReadStateResponse
		_parent, typeSwitchError = AdsReadStateResponseParse(readBuffer)
	case *commandId == CommandId_ADS_WRITE_CONTROL && response == false: // AdsWriteControlRequest
		_parent, typeSwitchError = AdsWriteControlRequestParse(readBuffer)
	case *commandId == CommandId_ADS_WRITE_CONTROL && response == true: // AdsWriteControlResponse
		_parent, typeSwitchError = AdsWriteControlResponseParse(readBuffer)
	case *commandId == CommandId_ADS_ADD_DEVICE_NOTIFICATION && response == false: // AdsAddDeviceNotificationRequest
		_parent, typeSwitchError = AdsAddDeviceNotificationRequestParse(readBuffer)
	case *commandId == CommandId_ADS_ADD_DEVICE_NOTIFICATION && response == true: // AdsAddDeviceNotificationResponse
		_parent, typeSwitchError = AdsAddDeviceNotificationResponseParse(readBuffer)
	case *commandId == CommandId_ADS_DELETE_DEVICE_NOTIFICATION && response == false: // AdsDeleteDeviceNotificationRequest
		_parent, typeSwitchError = AdsDeleteDeviceNotificationRequestParse(readBuffer)
	case *commandId == CommandId_ADS_DELETE_DEVICE_NOTIFICATION && response == true: // AdsDeleteDeviceNotificationResponse
		_parent, typeSwitchError = AdsDeleteDeviceNotificationResponseParse(readBuffer)
	case *commandId == CommandId_ADS_DEVICE_NOTIFICATION && response == false: // AdsDeviceNotificationRequest
		_parent, typeSwitchError = AdsDeviceNotificationRequestParse(readBuffer)
	case *commandId == CommandId_ADS_DEVICE_NOTIFICATION && response == true: // AdsDeviceNotificationResponse
		_parent, typeSwitchError = AdsDeviceNotificationResponseParse(readBuffer)
	case *commandId == CommandId_ADS_READ_WRITE && response == false: // AdsReadWriteRequest
		_parent, typeSwitchError = AdsReadWriteRequestParse(readBuffer)
	case *commandId == CommandId_ADS_READ_WRITE && response == true: // AdsReadWriteResponse
		_parent, typeSwitchError = AdsReadWriteResponseParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("AdsData"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *AdsData) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *AdsData) SerializeParent(writeBuffer utils.WriteBuffer, child IAdsData, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("AdsData"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AdsData"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AdsData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
