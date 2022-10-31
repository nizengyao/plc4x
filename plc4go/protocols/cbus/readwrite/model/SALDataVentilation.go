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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SALDataVentilation is the corresponding interface of SALDataVentilation
type SALDataVentilation interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetVentilationData returns VentilationData (property field)
	GetVentilationData() LightingData
}

// SALDataVentilationExactly can be used when we want exactly this type and not a type which fulfills SALDataVentilation.
// This is useful for switch cases.
type SALDataVentilationExactly interface {
	SALDataVentilation
	isSALDataVentilation() bool
}

// _SALDataVentilation is the data-structure of this message
type _SALDataVentilation struct {
	*_SALData
	VentilationData LightingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataVentilation) GetApplicationId() ApplicationId {
	return ApplicationId_VENTILATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataVentilation) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataVentilation) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataVentilation) GetVentilationData() LightingData {
	return m.VentilationData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataVentilation factory function for _SALDataVentilation
func NewSALDataVentilation(ventilationData LightingData, salData SALData) *_SALDataVentilation {
	_result := &_SALDataVentilation{
		VentilationData: ventilationData,
		_SALData:        NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataVentilation(structType interface{}) SALDataVentilation {
	if casted, ok := structType.(SALDataVentilation); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataVentilation); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataVentilation) GetTypeName() string {
	return "SALDataVentilation"
}

func (m *_SALDataVentilation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataVentilation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ventilationData)
	lengthInBits += m.VentilationData.GetLengthInBits()

	return lengthInBits
}

func (m *_SALDataVentilation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataVentilationParse(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataVentilation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataVentilation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataVentilation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ventilationData)
	if pullErr := readBuffer.PullContext("ventilationData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ventilationData")
	}
	_ventilationData, _ventilationDataErr := LightingDataParse(readBuffer)
	if _ventilationDataErr != nil {
		return nil, errors.Wrap(_ventilationDataErr, "Error parsing 'ventilationData' field of SALDataVentilation")
	}
	ventilationData := _ventilationData.(LightingData)
	if closeErr := readBuffer.CloseContext("ventilationData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ventilationData")
	}

	if closeErr := readBuffer.CloseContext("SALDataVentilation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataVentilation")
	}

	// Create a partially initialized instance
	_child := &_SALDataVentilation{
		_SALData:        &_SALData{},
		VentilationData: ventilationData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataVentilation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataVentilation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataVentilation")
		}

		// Simple Field (ventilationData)
		if pushErr := writeBuffer.PushContext("ventilationData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ventilationData")
		}
		_ventilationDataErr := writeBuffer.WriteSerializable(m.GetVentilationData())
		if popErr := writeBuffer.PopContext("ventilationData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ventilationData")
		}
		if _ventilationDataErr != nil {
			return errors.Wrap(_ventilationDataErr, "Error serializing 'ventilationData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataVentilation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataVentilation")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SALDataVentilation) isSALDataVentilation() bool {
	return true
}

func (m *_SALDataVentilation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}