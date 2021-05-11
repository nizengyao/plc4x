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
type S7Payload struct {
	Child IS7PayloadChild
}

// The corresponding interface
type IS7Payload interface {
	MessageType() uint8
	ParameterParameterType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7PayloadParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7Payload, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7PayloadChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7Payload)
	GetTypeName() string
	IS7Payload
}

func NewS7Payload() *S7Payload {
	return &S7Payload{}
}

func CastS7Payload(structType interface{}) *S7Payload {
	castFunc := func(typ interface{}) *S7Payload {
		if casted, ok := typ.(S7Payload); ok {
			return &casted
		}
		if casted, ok := typ.(*S7Payload); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7Payload) GetTypeName() string {
	return "S7Payload"
}

func (m *S7Payload) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7Payload) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *S7Payload) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *S7Payload) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadParse(readBuffer utils.ReadBuffer, messageType uint8, parameter *S7Parameter) (*S7Payload, error) {
	if pullErr := readBuffer.PullContext("S7Payload"); pullErr != nil {
		return nil, pullErr
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *S7Payload
	var typeSwitchError error
	switch {
	case CastS7Parameter(parameter).Child.ParameterType() == 0x04 && messageType == 0x03: // S7PayloadReadVarResponse
		_parent, typeSwitchError = S7PayloadReadVarResponseParse(readBuffer, parameter)
	case CastS7Parameter(parameter).Child.ParameterType() == 0x05 && messageType == 0x01: // S7PayloadWriteVarRequest
		_parent, typeSwitchError = S7PayloadWriteVarRequestParse(readBuffer, parameter)
	case CastS7Parameter(parameter).Child.ParameterType() == 0x05 && messageType == 0x03: // S7PayloadWriteVarResponse
		_parent, typeSwitchError = S7PayloadWriteVarResponseParse(readBuffer, parameter)
	case CastS7Parameter(parameter).Child.ParameterType() == 0x00 && messageType == 0x07: // S7PayloadUserData
		_parent, typeSwitchError = S7PayloadUserDataParse(readBuffer, parameter)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7Payload"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *S7Payload) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7Payload) SerializeParent(writeBuffer utils.WriteBuffer, child IS7Payload, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("S7Payload"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Payload"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7Payload) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
