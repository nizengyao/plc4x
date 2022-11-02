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

// SALDataTriggerControl is the corresponding interface of SALDataTriggerControl
type SALDataTriggerControl interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetTriggerControlData returns TriggerControlData (property field)
	GetTriggerControlData() TriggerControlData
}

// SALDataTriggerControlExactly can be used when we want exactly this type and not a type which fulfills SALDataTriggerControl.
// This is useful for switch cases.
type SALDataTriggerControlExactly interface {
	SALDataTriggerControl
	isSALDataTriggerControl() bool
}

// _SALDataTriggerControl is the data-structure of this message
type _SALDataTriggerControl struct {
	*_SALData
	TriggerControlData TriggerControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataTriggerControl) GetApplicationId() ApplicationId {
	return ApplicationId_TRIGGER_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTriggerControl) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataTriggerControl) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataTriggerControl) GetTriggerControlData() TriggerControlData {
	return m.TriggerControlData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataTriggerControl factory function for _SALDataTriggerControl
func NewSALDataTriggerControl(triggerControlData TriggerControlData, salData SALData) *_SALDataTriggerControl {
	_result := &_SALDataTriggerControl{
		TriggerControlData: triggerControlData,
		_SALData:           NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataTriggerControl(structType interface{}) SALDataTriggerControl {
	if casted, ok := structType.(SALDataTriggerControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTriggerControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTriggerControl) GetTypeName() string {
	return "SALDataTriggerControl"
}

func (m *_SALDataTriggerControl) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataTriggerControl) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (triggerControlData)
	lengthInBits += m.TriggerControlData.GetLengthInBits()

	return lengthInBits
}

func (m *_SALDataTriggerControl) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataTriggerControlParse(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataTriggerControl, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataTriggerControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTriggerControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (triggerControlData)
	if pullErr := readBuffer.PullContext("triggerControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for triggerControlData")
	}
	_triggerControlData, _triggerControlDataErr := TriggerControlDataParse(readBuffer)
	if _triggerControlDataErr != nil {
		return nil, errors.Wrap(_triggerControlDataErr, "Error parsing 'triggerControlData' field of SALDataTriggerControl")
	}
	triggerControlData := _triggerControlData.(TriggerControlData)
	if closeErr := readBuffer.CloseContext("triggerControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for triggerControlData")
	}

	if closeErr := readBuffer.CloseContext("SALDataTriggerControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTriggerControl")
	}

	// Create a partially initialized instance
	_child := &_SALDataTriggerControl{
		_SALData:           &_SALData{},
		TriggerControlData: triggerControlData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataTriggerControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataTriggerControl) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTriggerControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTriggerControl")
		}

		// Simple Field (triggerControlData)
		if pushErr := writeBuffer.PushContext("triggerControlData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for triggerControlData")
		}
		_triggerControlDataErr := writeBuffer.WriteSerializable(m.GetTriggerControlData())
		if popErr := writeBuffer.PopContext("triggerControlData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for triggerControlData")
		}
		if _triggerControlDataErr != nil {
			return errors.Wrap(_triggerControlDataErr, "Error serializing 'triggerControlData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataTriggerControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTriggerControl")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SALDataTriggerControl) isSALDataTriggerControl() bool {
	return true
}

func (m *_SALDataTriggerControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
