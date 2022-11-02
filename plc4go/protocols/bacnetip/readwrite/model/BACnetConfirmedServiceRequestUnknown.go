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

// BACnetConfirmedServiceRequestUnknown is the corresponding interface of BACnetConfirmedServiceRequestUnknown
type BACnetConfirmedServiceRequestUnknown interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetUnknownBytes returns UnknownBytes (property field)
	GetUnknownBytes() []byte
}

// BACnetConfirmedServiceRequestUnknownExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestUnknown.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestUnknownExactly interface {
	BACnetConfirmedServiceRequestUnknown
	isBACnetConfirmedServiceRequestUnknown() bool
}

// _BACnetConfirmedServiceRequestUnknown is the data-structure of this message
type _BACnetConfirmedServiceRequestUnknown struct {
	*_BACnetConfirmedServiceRequest
	UnknownBytes []byte

	// Arguments.
	ServiceRequestPayloadLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestUnknown) GetServiceChoice() BACnetConfirmedServiceChoice {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestUnknown) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestUnknown) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestUnknown) GetUnknownBytes() []byte {
	return m.UnknownBytes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestUnknown factory function for _BACnetConfirmedServiceRequestUnknown
func NewBACnetConfirmedServiceRequestUnknown(unknownBytes []byte, serviceRequestLength uint32, serviceRequestPayloadLength uint32) *_BACnetConfirmedServiceRequestUnknown {
	_result := &_BACnetConfirmedServiceRequestUnknown{
		UnknownBytes:                   unknownBytes,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestUnknown(structType interface{}) BACnetConfirmedServiceRequestUnknown {
	if casted, ok := structType.(BACnetConfirmedServiceRequestUnknown); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestUnknown); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestUnknown) GetTypeName() string {
	return "BACnetConfirmedServiceRequestUnknown"
}

func (m *_BACnetConfirmedServiceRequestUnknown) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestUnknown) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.UnknownBytes) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownBytes))
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestUnknown) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestUnknownParse(readBuffer utils.ReadBuffer, serviceRequestLength uint32, serviceRequestPayloadLength uint32) (BACnetConfirmedServiceRequestUnknown, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestUnknown"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestUnknown")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (unknownBytes)
	numberOfBytesunknownBytes := int(serviceRequestPayloadLength)
	unknownBytes, _readArrayErr := readBuffer.ReadByteArray("unknownBytes", numberOfBytesunknownBytes)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownBytes' field of BACnetConfirmedServiceRequestUnknown")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestUnknown"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestUnknown")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestUnknown{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		UnknownBytes: unknownBytes,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestUnknown) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestUnknown) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestUnknown"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestUnknown")
		}

		// Array Field (unknownBytes)
		// Byte Array field (unknownBytes)
		if err := writeBuffer.WriteByteArray("unknownBytes", m.GetUnknownBytes()); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownBytes' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestUnknown"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestUnknown")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestUnknown) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestUnknown) isBACnetConfirmedServiceRequestUnknown() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
