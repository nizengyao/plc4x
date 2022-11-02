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

// NLMInitalizeRoutingTable is the corresponding interface of NLMInitalizeRoutingTable
type NLMInitalizeRoutingTable interface {
	utils.LengthAware
	utils.Serializable
	NLM
	// GetNumberOfPorts returns NumberOfPorts (property field)
	GetNumberOfPorts() uint8
	// GetPortMappings returns PortMappings (property field)
	GetPortMappings() []NLMInitalizeRoutingTablePortMapping
}

// NLMInitalizeRoutingTableExactly can be used when we want exactly this type and not a type which fulfills NLMInitalizeRoutingTable.
// This is useful for switch cases.
type NLMInitalizeRoutingTableExactly interface {
	NLMInitalizeRoutingTable
	isNLMInitalizeRoutingTable() bool
}

// _NLMInitalizeRoutingTable is the data-structure of this message
type _NLMInitalizeRoutingTable struct {
	*_NLM
	NumberOfPorts uint8
	PortMappings  []NLMInitalizeRoutingTablePortMapping
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMInitalizeRoutingTable) GetMessageType() uint8 {
	return 0x06
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMInitalizeRoutingTable) InitializeParent(parent NLM, vendorId *BACnetVendorId) {
	m.VendorId = vendorId
}

func (m *_NLMInitalizeRoutingTable) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMInitalizeRoutingTable) GetNumberOfPorts() uint8 {
	return m.NumberOfPorts
}

func (m *_NLMInitalizeRoutingTable) GetPortMappings() []NLMInitalizeRoutingTablePortMapping {
	return m.PortMappings
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMInitalizeRoutingTable factory function for _NLMInitalizeRoutingTable
func NewNLMInitalizeRoutingTable(numberOfPorts uint8, portMappings []NLMInitalizeRoutingTablePortMapping, vendorId *BACnetVendorId, apduLength uint16) *_NLMInitalizeRoutingTable {
	_result := &_NLMInitalizeRoutingTable{
		NumberOfPorts: numberOfPorts,
		PortMappings:  portMappings,
		_NLM:          NewNLM(vendorId, apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMInitalizeRoutingTable(structType interface{}) NLMInitalizeRoutingTable {
	if casted, ok := structType.(NLMInitalizeRoutingTable); ok {
		return casted
	}
	if casted, ok := structType.(*NLMInitalizeRoutingTable); ok {
		return *casted
	}
	return nil
}

func (m *_NLMInitalizeRoutingTable) GetTypeName() string {
	return "NLMInitalizeRoutingTable"
}

func (m *_NLMInitalizeRoutingTable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NLMInitalizeRoutingTable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numberOfPorts)
	lengthInBits += 8

	// Array field
	if len(m.PortMappings) > 0 {
		for i, element := range m.PortMappings {
			last := i == len(m.PortMappings)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_NLMInitalizeRoutingTable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMInitalizeRoutingTableParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (NLMInitalizeRoutingTable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMInitalizeRoutingTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMInitalizeRoutingTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numberOfPorts)
	_numberOfPorts, _numberOfPortsErr := readBuffer.ReadUint8("numberOfPorts", 8)
	if _numberOfPortsErr != nil {
		return nil, errors.Wrap(_numberOfPortsErr, "Error parsing 'numberOfPorts' field of NLMInitalizeRoutingTable")
	}
	numberOfPorts := _numberOfPorts

	// Array field (portMappings)
	if pullErr := readBuffer.PullContext("portMappings", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for portMappings")
	}
	// Count array
	portMappings := make([]NLMInitalizeRoutingTablePortMapping, numberOfPorts)
	// This happens when the size is set conditional to 0
	if len(portMappings) == 0 {
		portMappings = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(numberOfPorts); curItem++ {
			_item, _err := NLMInitalizeRoutingTablePortMappingParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'portMappings' field of NLMInitalizeRoutingTable")
			}
			portMappings[curItem] = _item.(NLMInitalizeRoutingTablePortMapping)
		}
	}
	if closeErr := readBuffer.CloseContext("portMappings", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for portMappings")
	}

	if closeErr := readBuffer.CloseContext("NLMInitalizeRoutingTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMInitalizeRoutingTable")
	}

	// Create a partially initialized instance
	_child := &_NLMInitalizeRoutingTable{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		NumberOfPorts: numberOfPorts,
		PortMappings:  portMappings,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMInitalizeRoutingTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMInitalizeRoutingTable) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMInitalizeRoutingTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMInitalizeRoutingTable")
		}

		// Simple Field (numberOfPorts)
		numberOfPorts := uint8(m.GetNumberOfPorts())
		_numberOfPortsErr := writeBuffer.WriteUint8("numberOfPorts", 8, (numberOfPorts))
		if _numberOfPortsErr != nil {
			return errors.Wrap(_numberOfPortsErr, "Error serializing 'numberOfPorts' field")
		}

		// Array Field (portMappings)
		if pushErr := writeBuffer.PushContext("portMappings", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for portMappings")
		}
		for _, _element := range m.GetPortMappings() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'portMappings' field")
			}
		}
		if popErr := writeBuffer.PopContext("portMappings", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for portMappings")
		}

		if popErr := writeBuffer.PopContext("NLMInitalizeRoutingTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMInitalizeRoutingTable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_NLMInitalizeRoutingTable) isNLMInitalizeRoutingTable() bool {
	return true
}

func (m *_NLMInitalizeRoutingTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
