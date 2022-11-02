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

// MediaTransportControlDataNextPreviousSelection is the corresponding interface of MediaTransportControlDataNextPreviousSelection
type MediaTransportControlDataNextPreviousSelection interface {
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetOperation returns Operation (property field)
	GetOperation() byte
	// GetIsSetThePreviousSelection returns IsSetThePreviousSelection (virtual field)
	GetIsSetThePreviousSelection() bool
	// GetIsSetTheNextSelection returns IsSetTheNextSelection (virtual field)
	GetIsSetTheNextSelection() bool
}

// MediaTransportControlDataNextPreviousSelectionExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataNextPreviousSelection.
// This is useful for switch cases.
type MediaTransportControlDataNextPreviousSelectionExactly interface {
	MediaTransportControlDataNextPreviousSelection
	isMediaTransportControlDataNextPreviousSelection() bool
}

// _MediaTransportControlDataNextPreviousSelection is the data-structure of this message
type _MediaTransportControlDataNextPreviousSelection struct {
	*_MediaTransportControlData
	Operation byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataNextPreviousSelection) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataNextPreviousSelection) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataNextPreviousSelection) GetOperation() byte {
	return m.Operation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataNextPreviousSelection) GetIsSetThePreviousSelection() bool {
	return bool(bool((m.GetOperation()) == (0x00)))
}

func (m *_MediaTransportControlDataNextPreviousSelection) GetIsSetTheNextSelection() bool {
	return bool(bool((m.GetOperation()) != (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataNextPreviousSelection factory function for _MediaTransportControlDataNextPreviousSelection
func NewMediaTransportControlDataNextPreviousSelection(operation byte, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataNextPreviousSelection {
	_result := &_MediaTransportControlDataNextPreviousSelection{
		Operation:                  operation,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataNextPreviousSelection(structType interface{}) MediaTransportControlDataNextPreviousSelection {
	if casted, ok := structType.(MediaTransportControlDataNextPreviousSelection); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataNextPreviousSelection); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataNextPreviousSelection) GetTypeName() string {
	return "MediaTransportControlDataNextPreviousSelection"
}

func (m *_MediaTransportControlDataNextPreviousSelection) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MediaTransportControlDataNextPreviousSelection) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (operation)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataNextPreviousSelection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MediaTransportControlDataNextPreviousSelectionParse(readBuffer utils.ReadBuffer) (MediaTransportControlDataNextPreviousSelection, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataNextPreviousSelection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataNextPreviousSelection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (operation)
	_operation, _operationErr := readBuffer.ReadByte("operation")
	if _operationErr != nil {
		return nil, errors.Wrap(_operationErr, "Error parsing 'operation' field of MediaTransportControlDataNextPreviousSelection")
	}
	operation := _operation

	// Virtual field
	_isSetThePreviousSelection := bool((operation) == (0x00))
	isSetThePreviousSelection := bool(_isSetThePreviousSelection)
	_ = isSetThePreviousSelection

	// Virtual field
	_isSetTheNextSelection := bool((operation) != (0x00))
	isSetTheNextSelection := bool(_isSetTheNextSelection)
	_ = isSetTheNextSelection

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataNextPreviousSelection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataNextPreviousSelection")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataNextPreviousSelection{
		_MediaTransportControlData: &_MediaTransportControlData{},
		Operation:                  operation,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataNextPreviousSelection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataNextPreviousSelection) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataNextPreviousSelection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataNextPreviousSelection")
		}

		// Simple Field (operation)
		operation := byte(m.GetOperation())
		_operationErr := writeBuffer.WriteByte("operation", (operation))
		if _operationErr != nil {
			return errors.Wrap(_operationErr, "Error serializing 'operation' field")
		}
		// Virtual field
		if _isSetThePreviousSelectionErr := writeBuffer.WriteVirtual("isSetThePreviousSelection", m.GetIsSetThePreviousSelection()); _isSetThePreviousSelectionErr != nil {
			return errors.Wrap(_isSetThePreviousSelectionErr, "Error serializing 'isSetThePreviousSelection' field")
		}
		// Virtual field
		if _isSetTheNextSelectionErr := writeBuffer.WriteVirtual("isSetTheNextSelection", m.GetIsSetTheNextSelection()); _isSetTheNextSelectionErr != nil {
			return errors.Wrap(_isSetTheNextSelectionErr, "Error serializing 'isSetTheNextSelection' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataNextPreviousSelection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataNextPreviousSelection")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataNextPreviousSelection) isMediaTransportControlDataNextPreviousSelection() bool {
	return true
}

func (m *_MediaTransportControlDataNextPreviousSelection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
