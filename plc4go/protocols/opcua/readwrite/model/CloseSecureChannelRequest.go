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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// CloseSecureChannelRequest is the corresponding interface of CloseSecureChannelRequest
type CloseSecureChannelRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
}

// CloseSecureChannelRequestExactly can be used when we want exactly this type and not a type which fulfills CloseSecureChannelRequest.
// This is useful for switch cases.
type CloseSecureChannelRequestExactly interface {
	CloseSecureChannelRequest
	isCloseSecureChannelRequest() bool
}

// _CloseSecureChannelRequest is the data-structure of this message
type _CloseSecureChannelRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CloseSecureChannelRequest) GetIdentifier() string {
	return "452"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CloseSecureChannelRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CloseSecureChannelRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CloseSecureChannelRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCloseSecureChannelRequest factory function for _CloseSecureChannelRequest
func NewCloseSecureChannelRequest(requestHeader ExtensionObjectDefinition) *_CloseSecureChannelRequest {
	_result := &_CloseSecureChannelRequest{
		RequestHeader:              requestHeader,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCloseSecureChannelRequest(structType any) CloseSecureChannelRequest {
	if casted, ok := structType.(CloseSecureChannelRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CloseSecureChannelRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CloseSecureChannelRequest) GetTypeName() string {
	return "CloseSecureChannelRequest"
}

func (m *_CloseSecureChannelRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CloseSecureChannelRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CloseSecureChannelRequestParse(ctx context.Context, theBytes []byte, identifier string) (CloseSecureChannelRequest, error) {
	return CloseSecureChannelRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CloseSecureChannelRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CloseSecureChannelRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CloseSecureChannelRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CloseSecureChannelRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of CloseSecureChannelRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	if closeErr := readBuffer.CloseContext("CloseSecureChannelRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CloseSecureChannelRequest")
	}

	// Create a partially initialized instance
	_child := &_CloseSecureChannelRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CloseSecureChannelRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CloseSecureChannelRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CloseSecureChannelRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CloseSecureChannelRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		if popErr := writeBuffer.PopContext("CloseSecureChannelRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CloseSecureChannelRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CloseSecureChannelRequest) isCloseSecureChannelRequest() bool {
	return true
}

func (m *_CloseSecureChannelRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
