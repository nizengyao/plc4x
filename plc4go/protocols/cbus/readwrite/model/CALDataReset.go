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

// CALDataReset is the corresponding interface of CALDataReset
type CALDataReset interface {
	utils.LengthAware
	utils.Serializable
	CALData
}

// CALDataResetExactly can be used when we want exactly this type and not a type which fulfills CALDataReset.
// This is useful for switch cases.
type CALDataResetExactly interface {
	CALDataReset
	isCALDataReset() bool
}

// _CALDataReset is the data-structure of this message
type _CALDataReset struct {
	*_CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataReset) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer, additionalData CALData) {
	m.CommandTypeContainer = commandTypeContainer
	m.AdditionalData = additionalData
}

func (m *_CALDataReset) GetParent() CALData {
	return m._CALData
}

// NewCALDataReset factory function for _CALDataReset
func NewCALDataReset(commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataReset {
	_result := &_CALDataReset{
		_CALData: NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataReset(structType interface{}) CALDataReset {
	if casted, ok := structType.(CALDataReset); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataReset); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataReset) GetTypeName() string {
	return "CALDataReset"
}

func (m *_CALDataReset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataReset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_CALDataReset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataResetParse(readBuffer utils.ReadBuffer, requestContext RequestContext) (CALDataReset, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CALDataReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataReset")
	}

	// Create a partially initialized instance
	_child := &_CALDataReset{
		_CALData: &_CALData{
			RequestContext: requestContext,
		},
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataReset) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataReset")
		}

		if popErr := writeBuffer.PopContext("CALDataReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataReset")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataReset) isCALDataReset() bool {
	return true
}

func (m *_CALDataReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
