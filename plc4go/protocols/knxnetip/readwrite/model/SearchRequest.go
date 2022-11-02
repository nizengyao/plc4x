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

// SearchRequest is the corresponding interface of SearchRequest
type SearchRequest interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetHpaiIDiscoveryEndpoint returns HpaiIDiscoveryEndpoint (property field)
	GetHpaiIDiscoveryEndpoint() HPAIDiscoveryEndpoint
}

// SearchRequestExactly can be used when we want exactly this type and not a type which fulfills SearchRequest.
// This is useful for switch cases.
type SearchRequestExactly interface {
	SearchRequest
	isSearchRequest() bool
}

// _SearchRequest is the data-structure of this message
type _SearchRequest struct {
	*_KnxNetIpMessage
	HpaiIDiscoveryEndpoint HPAIDiscoveryEndpoint
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SearchRequest) GetMsgType() uint16 {
	return 0x0201
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SearchRequest) InitializeParent(parent KnxNetIpMessage) {}

func (m *_SearchRequest) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SearchRequest) GetHpaiIDiscoveryEndpoint() HPAIDiscoveryEndpoint {
	return m.HpaiIDiscoveryEndpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSearchRequest factory function for _SearchRequest
func NewSearchRequest(hpaiIDiscoveryEndpoint HPAIDiscoveryEndpoint) *_SearchRequest {
	_result := &_SearchRequest{
		HpaiIDiscoveryEndpoint: hpaiIDiscoveryEndpoint,
		_KnxNetIpMessage:       NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSearchRequest(structType interface{}) SearchRequest {
	if casted, ok := structType.(SearchRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SearchRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SearchRequest) GetTypeName() string {
	return "SearchRequest"
}

func (m *_SearchRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SearchRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (hpaiIDiscoveryEndpoint)
	lengthInBits += m.HpaiIDiscoveryEndpoint.GetLengthInBits()

	return lengthInBits
}

func (m *_SearchRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SearchRequestParse(readBuffer utils.ReadBuffer) (SearchRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SearchRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SearchRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (hpaiIDiscoveryEndpoint)
	if pullErr := readBuffer.PullContext("hpaiIDiscoveryEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hpaiIDiscoveryEndpoint")
	}
	_hpaiIDiscoveryEndpoint, _hpaiIDiscoveryEndpointErr := HPAIDiscoveryEndpointParse(readBuffer)
	if _hpaiIDiscoveryEndpointErr != nil {
		return nil, errors.Wrap(_hpaiIDiscoveryEndpointErr, "Error parsing 'hpaiIDiscoveryEndpoint' field of SearchRequest")
	}
	hpaiIDiscoveryEndpoint := _hpaiIDiscoveryEndpoint.(HPAIDiscoveryEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiIDiscoveryEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hpaiIDiscoveryEndpoint")
	}

	if closeErr := readBuffer.CloseContext("SearchRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SearchRequest")
	}

	// Create a partially initialized instance
	_child := &_SearchRequest{
		_KnxNetIpMessage:       &_KnxNetIpMessage{},
		HpaiIDiscoveryEndpoint: hpaiIDiscoveryEndpoint,
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_SearchRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SearchRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SearchRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SearchRequest")
		}

		// Simple Field (hpaiIDiscoveryEndpoint)
		if pushErr := writeBuffer.PushContext("hpaiIDiscoveryEndpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hpaiIDiscoveryEndpoint")
		}
		_hpaiIDiscoveryEndpointErr := writeBuffer.WriteSerializable(m.GetHpaiIDiscoveryEndpoint())
		if popErr := writeBuffer.PopContext("hpaiIDiscoveryEndpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hpaiIDiscoveryEndpoint")
		}
		if _hpaiIDiscoveryEndpointErr != nil {
			return errors.Wrap(_hpaiIDiscoveryEndpointErr, "Error serializing 'hpaiIDiscoveryEndpoint' field")
		}

		if popErr := writeBuffer.PopContext("SearchRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SearchRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SearchRequest) isSearchRequest() bool {
	return true
}

func (m *_SearchRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
