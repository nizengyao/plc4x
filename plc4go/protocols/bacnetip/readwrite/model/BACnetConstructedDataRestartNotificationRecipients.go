/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataRestartNotificationRecipients is the data-structure of this message
type BACnetConstructedDataRestartNotificationRecipients struct {
	*BACnetConstructedData
	RestartNotificationRecipients []*BACnetRecipient

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataRestartNotificationRecipients is the corresponding interface of BACnetConstructedDataRestartNotificationRecipients
type IBACnetConstructedDataRestartNotificationRecipients interface {
	IBACnetConstructedData
	// GetRestartNotificationRecipients returns RestartNotificationRecipients (property field)
	GetRestartNotificationRecipients() []*BACnetRecipient
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataRestartNotificationRecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataRestartNotificationRecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESTART_NOTIFICATION_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataRestartNotificationRecipients) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataRestartNotificationRecipients) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataRestartNotificationRecipients) GetRestartNotificationRecipients() []*BACnetRecipient {
	return m.RestartNotificationRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRestartNotificationRecipients factory function for BACnetConstructedDataRestartNotificationRecipients
func NewBACnetConstructedDataRestartNotificationRecipients(restartNotificationRecipients []*BACnetRecipient, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataRestartNotificationRecipients {
	_result := &BACnetConstructedDataRestartNotificationRecipients{
		RestartNotificationRecipients: restartNotificationRecipients,
		BACnetConstructedData:         NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataRestartNotificationRecipients(structType interface{}) *BACnetConstructedDataRestartNotificationRecipients {
	if casted, ok := structType.(BACnetConstructedDataRestartNotificationRecipients); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRestartNotificationRecipients); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataRestartNotificationRecipients(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataRestartNotificationRecipients(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataRestartNotificationRecipients) GetTypeName() string {
	return "BACnetConstructedDataRestartNotificationRecipients"
}

func (m *BACnetConstructedDataRestartNotificationRecipients) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataRestartNotificationRecipients) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.RestartNotificationRecipients) > 0 {
		for _, element := range m.RestartNotificationRecipients {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataRestartNotificationRecipients) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRestartNotificationRecipientsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataRestartNotificationRecipients, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRestartNotificationRecipients"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (restartNotificationRecipients)
	if pullErr := readBuffer.PullContext("restartNotificationRecipients", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	restartNotificationRecipients := make([]*BACnetRecipient, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetRecipientParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'restartNotificationRecipients' field")
			}
			restartNotificationRecipients = append(restartNotificationRecipients, CastBACnetRecipient(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("restartNotificationRecipients", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRestartNotificationRecipients"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataRestartNotificationRecipients{
		RestartNotificationRecipients: restartNotificationRecipients,
		BACnetConstructedData:         &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataRestartNotificationRecipients) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRestartNotificationRecipients"); pushErr != nil {
			return pushErr
		}

		// Array Field (restartNotificationRecipients)
		if m.RestartNotificationRecipients != nil {
			if pushErr := writeBuffer.PushContext("restartNotificationRecipients", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.RestartNotificationRecipients {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'restartNotificationRecipients' field")
				}
			}
			if popErr := writeBuffer.PopContext("restartNotificationRecipients", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRestartNotificationRecipients"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataRestartNotificationRecipients) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
