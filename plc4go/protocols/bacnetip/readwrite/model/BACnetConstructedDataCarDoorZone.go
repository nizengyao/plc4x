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

// BACnetConstructedDataCarDoorZone is the corresponding interface of BACnetConstructedDataCarDoorZone
type BACnetConstructedDataCarDoorZone interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCarDoorZone returns CarDoorZone (property field)
	GetCarDoorZone() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataCarDoorZoneExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCarDoorZone.
// This is useful for switch cases.
type BACnetConstructedDataCarDoorZoneExactly interface {
	BACnetConstructedDataCarDoorZone
	isBACnetConstructedDataCarDoorZone() bool
}

// _BACnetConstructedDataCarDoorZone is the data-structure of this message
type _BACnetConstructedDataCarDoorZone struct {
	*_BACnetConstructedData
	CarDoorZone BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarDoorZone) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCarDoorZone) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_DOOR_ZONE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarDoorZone) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCarDoorZone) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarDoorZone) GetCarDoorZone() BACnetApplicationTagBoolean {
	return m.CarDoorZone
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarDoorZone) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetCarDoorZone())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCarDoorZone factory function for _BACnetConstructedDataCarDoorZone
func NewBACnetConstructedDataCarDoorZone(carDoorZone BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCarDoorZone {
	_result := &_BACnetConstructedDataCarDoorZone{
		CarDoorZone:            carDoorZone,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarDoorZone(structType interface{}) BACnetConstructedDataCarDoorZone {
	if casted, ok := structType.(BACnetConstructedDataCarDoorZone); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarDoorZone); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarDoorZone) GetTypeName() string {
	return "BACnetConstructedDataCarDoorZone"
}

func (m *_BACnetConstructedDataCarDoorZone) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCarDoorZone) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (carDoorZone)
	lengthInBits += m.CarDoorZone.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCarDoorZone) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCarDoorZoneParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarDoorZone, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarDoorZone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarDoorZone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (carDoorZone)
	if pullErr := readBuffer.PullContext("carDoorZone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for carDoorZone")
	}
	_carDoorZone, _carDoorZoneErr := BACnetApplicationTagParse(readBuffer)
	if _carDoorZoneErr != nil {
		return nil, errors.Wrap(_carDoorZoneErr, "Error parsing 'carDoorZone' field of BACnetConstructedDataCarDoorZone")
	}
	carDoorZone := _carDoorZone.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("carDoorZone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for carDoorZone")
	}

	// Virtual field
	_actualValue := carDoorZone
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarDoorZone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarDoorZone")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCarDoorZone{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CarDoorZone: carDoorZone,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCarDoorZone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCarDoorZone) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarDoorZone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarDoorZone")
		}

		// Simple Field (carDoorZone)
		if pushErr := writeBuffer.PushContext("carDoorZone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for carDoorZone")
		}
		_carDoorZoneErr := writeBuffer.WriteSerializable(m.GetCarDoorZone())
		if popErr := writeBuffer.PopContext("carDoorZone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for carDoorZone")
		}
		if _carDoorZoneErr != nil {
			return errors.Wrap(_carDoorZoneErr, "Error serializing 'carDoorZone' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarDoorZone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarDoorZone")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCarDoorZone) isBACnetConstructedDataCarDoorZone() bool {
	return true
}

func (m *_BACnetConstructedDataCarDoorZone) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
