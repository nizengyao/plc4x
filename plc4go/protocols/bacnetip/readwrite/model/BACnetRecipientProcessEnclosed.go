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

// BACnetRecipientProcessEnclosed is the corresponding interface of BACnetRecipientProcessEnclosed
type BACnetRecipientProcessEnclosed interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetRecipientProcess returns RecipientProcess (property field)
	GetRecipientProcess() BACnetRecipientProcess
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetRecipientProcessEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetRecipientProcessEnclosed.
// This is useful for switch cases.
type BACnetRecipientProcessEnclosedExactly interface {
	BACnetRecipientProcessEnclosed
	isBACnetRecipientProcessEnclosed() bool
}

// _BACnetRecipientProcessEnclosed is the data-structure of this message
type _BACnetRecipientProcessEnclosed struct {
	OpeningTag       BACnetOpeningTag
	RecipientProcess BACnetRecipientProcess
	ClosingTag       BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientProcessEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetRecipientProcessEnclosed) GetRecipientProcess() BACnetRecipientProcess {
	return m.RecipientProcess
}

func (m *_BACnetRecipientProcessEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRecipientProcessEnclosed factory function for _BACnetRecipientProcessEnclosed
func NewBACnetRecipientProcessEnclosed(openingTag BACnetOpeningTag, recipientProcess BACnetRecipientProcess, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetRecipientProcessEnclosed {
	return &_BACnetRecipientProcessEnclosed{OpeningTag: openingTag, RecipientProcess: recipientProcess, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetRecipientProcessEnclosed(structType interface{}) BACnetRecipientProcessEnclosed {
	if casted, ok := structType.(BACnetRecipientProcessEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientProcessEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientProcessEnclosed) GetTypeName() string {
	return "BACnetRecipientProcessEnclosed"
}

func (m *_BACnetRecipientProcessEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetRecipientProcessEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (recipientProcess)
	lengthInBits += m.RecipientProcess.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetRecipientProcessEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRecipientProcessEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetRecipientProcessEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientProcessEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientProcessEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetRecipientProcessEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (recipientProcess)
	if pullErr := readBuffer.PullContext("recipientProcess"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for recipientProcess")
	}
	_recipientProcess, _recipientProcessErr := BACnetRecipientProcessParse(readBuffer)
	if _recipientProcessErr != nil {
		return nil, errors.Wrap(_recipientProcessErr, "Error parsing 'recipientProcess' field of BACnetRecipientProcessEnclosed")
	}
	recipientProcess := _recipientProcess.(BACnetRecipientProcess)
	if closeErr := readBuffer.CloseContext("recipientProcess"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for recipientProcess")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetRecipientProcessEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipientProcessEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientProcessEnclosed")
	}

	// Create the instance
	return &_BACnetRecipientProcessEnclosed{
		TagNumber:        tagNumber,
		OpeningTag:       openingTag,
		RecipientProcess: recipientProcess,
		ClosingTag:       closingTag,
	}, nil
}

func (m *_BACnetRecipientProcessEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientProcessEnclosed) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetRecipientProcessEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRecipientProcessEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (recipientProcess)
	if pushErr := writeBuffer.PushContext("recipientProcess"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for recipientProcess")
	}
	_recipientProcessErr := writeBuffer.WriteSerializable(m.GetRecipientProcess())
	if popErr := writeBuffer.PopContext("recipientProcess"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for recipientProcess")
	}
	if _recipientProcessErr != nil {
		return errors.Wrap(_recipientProcessErr, "Error serializing 'recipientProcess' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRecipientProcessEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRecipientProcessEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetRecipientProcessEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetRecipientProcessEnclosed) isBACnetRecipientProcessEnclosed() bool {
	return true
}

func (m *_BACnetRecipientProcessEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
