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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple is the corresponding interface of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetLifetime returns Lifetime (property field)
	GetLifetime() BACnetContextTagUnsignedInteger
	// GetMaxNotificationDelay returns MaxNotificationDelay (property field)
	GetMaxNotificationDelay() BACnetContextTagUnsignedInteger
	// GetListOfCovSubscriptionSpecifications returns ListOfCovSubscriptionSpecifications (property field)
	GetListOfCovSubscriptionSpecifications() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList
}

// BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleExactly interface {
	BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
	isBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple() bool
}

// _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple struct {
	*_BACnetConfirmedServiceRequest
	SubscriberProcessIdentifier         BACnetContextTagUnsignedInteger
	IssueConfirmedNotifications         BACnetContextTagBoolean
	Lifetime                            BACnetContextTagUnsignedInteger
	MaxNotificationDelay                BACnetContextTagUnsignedInteger
	ListOfCovSubscriptionSpecifications BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetLifetime() BACnetContextTagUnsignedInteger {
	return m.Lifetime
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetMaxNotificationDelay() BACnetContextTagUnsignedInteger {
	return m.MaxNotificationDelay
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetListOfCovSubscriptionSpecifications() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList {
	return m.ListOfCovSubscriptionSpecifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple factory function for _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
func NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, issueConfirmedNotifications BACnetContextTagBoolean, lifetime BACnetContextTagUnsignedInteger, maxNotificationDelay BACnetContextTagUnsignedInteger, listOfCovSubscriptionSpecifications BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
	_result := &_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple{
		SubscriberProcessIdentifier:         subscriberProcessIdentifier,
		IssueConfirmedNotifications:         issueConfirmedNotifications,
		Lifetime:                            lifetime,
		MaxNotificationDelay:                maxNotificationDelay,
		ListOfCovSubscriptionSpecifications: listOfCovSubscriptionSpecifications,
		_BACnetConfirmedServiceRequest:      NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple(structType interface{}) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits()

	// Optional Field (issueConfirmedNotifications)
	if m.IssueConfirmedNotifications != nil {
		lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits()
	}

	// Optional Field (lifetime)
	if m.Lifetime != nil {
		lengthInBits += m.Lifetime.GetLengthInBits()
	}

	// Optional Field (maxNotificationDelay)
	if m.MaxNotificationDelay != nil {
		lengthInBits += m.MaxNotificationDelay.GetLengthInBits()
	}

	// Simple field (listOfCovSubscriptionSpecifications)
	lengthInBits += m.ListOfCovSubscriptionSpecifications.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleParse(readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (subscriberProcessIdentifier)
	if pullErr := readBuffer.PullContext("subscriberProcessIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subscriberProcessIdentifier")
	}
	_subscriberProcessIdentifier, _subscriberProcessIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
	}
	subscriberProcessIdentifier := _subscriberProcessIdentifier.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("subscriberProcessIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subscriberProcessIdentifier")
	}

	// Optional Field (issueConfirmedNotifications) (Can be skipped, if a given expression evaluates to false)
	var issueConfirmedNotifications BACnetContextTagBoolean = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("issueConfirmedNotifications"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for issueConfirmedNotifications")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_BOOLEAN)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'issueConfirmedNotifications' field of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
		default:
			issueConfirmedNotifications = _val.(BACnetContextTagBoolean)
			if closeErr := readBuffer.CloseContext("issueConfirmedNotifications"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for issueConfirmedNotifications")
			}
		}
	}

	// Optional Field (lifetime) (Can be skipped, if a given expression evaluates to false)
	var lifetime BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("lifetime"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for lifetime")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'lifetime' field of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
		default:
			lifetime = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("lifetime"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for lifetime")
			}
		}
	}

	// Optional Field (maxNotificationDelay) (Can be skipped, if a given expression evaluates to false)
	var maxNotificationDelay BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("maxNotificationDelay"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for maxNotificationDelay")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'maxNotificationDelay' field of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
		default:
			maxNotificationDelay = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("maxNotificationDelay"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for maxNotificationDelay")
			}
		}
	}

	// Simple Field (listOfCovSubscriptionSpecifications)
	if pullErr := readBuffer.PullContext("listOfCovSubscriptionSpecifications"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfCovSubscriptionSpecifications")
	}
	_listOfCovSubscriptionSpecifications, _listOfCovSubscriptionSpecificationsErr := BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsListParse(readBuffer, uint8(uint8(4)))
	if _listOfCovSubscriptionSpecificationsErr != nil {
		return nil, errors.Wrap(_listOfCovSubscriptionSpecificationsErr, "Error parsing 'listOfCovSubscriptionSpecifications' field of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
	}
	listOfCovSubscriptionSpecifications := _listOfCovSubscriptionSpecifications.(BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList)
	if closeErr := readBuffer.CloseContext("listOfCovSubscriptionSpecifications"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfCovSubscriptionSpecifications")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		SubscriberProcessIdentifier:         subscriberProcessIdentifier,
		IssueConfirmedNotifications:         issueConfirmedNotifications,
		Lifetime:                            lifetime,
		MaxNotificationDelay:                maxNotificationDelay,
		ListOfCovSubscriptionSpecifications: listOfCovSubscriptionSpecifications,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
		}

		// Simple Field (subscriberProcessIdentifier)
		if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subscriberProcessIdentifier")
		}
		_subscriberProcessIdentifierErr := writeBuffer.WriteSerializable(m.GetSubscriberProcessIdentifier())
		if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subscriberProcessIdentifier")
		}
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Optional Field (issueConfirmedNotifications) (Can be skipped, if the value is null)
		var issueConfirmedNotifications BACnetContextTagBoolean = nil
		if m.GetIssueConfirmedNotifications() != nil {
			if pushErr := writeBuffer.PushContext("issueConfirmedNotifications"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for issueConfirmedNotifications")
			}
			issueConfirmedNotifications = m.GetIssueConfirmedNotifications()
			_issueConfirmedNotificationsErr := writeBuffer.WriteSerializable(issueConfirmedNotifications)
			if popErr := writeBuffer.PopContext("issueConfirmedNotifications"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for issueConfirmedNotifications")
			}
			if _issueConfirmedNotificationsErr != nil {
				return errors.Wrap(_issueConfirmedNotificationsErr, "Error serializing 'issueConfirmedNotifications' field")
			}
		}

		// Optional Field (lifetime) (Can be skipped, if the value is null)
		var lifetime BACnetContextTagUnsignedInteger = nil
		if m.GetLifetime() != nil {
			if pushErr := writeBuffer.PushContext("lifetime"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for lifetime")
			}
			lifetime = m.GetLifetime()
			_lifetimeErr := writeBuffer.WriteSerializable(lifetime)
			if popErr := writeBuffer.PopContext("lifetime"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for lifetime")
			}
			if _lifetimeErr != nil {
				return errors.Wrap(_lifetimeErr, "Error serializing 'lifetime' field")
			}
		}

		// Optional Field (maxNotificationDelay) (Can be skipped, if the value is null)
		var maxNotificationDelay BACnetContextTagUnsignedInteger = nil
		if m.GetMaxNotificationDelay() != nil {
			if pushErr := writeBuffer.PushContext("maxNotificationDelay"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for maxNotificationDelay")
			}
			maxNotificationDelay = m.GetMaxNotificationDelay()
			_maxNotificationDelayErr := writeBuffer.WriteSerializable(maxNotificationDelay)
			if popErr := writeBuffer.PopContext("maxNotificationDelay"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for maxNotificationDelay")
			}
			if _maxNotificationDelayErr != nil {
				return errors.Wrap(_maxNotificationDelayErr, "Error serializing 'maxNotificationDelay' field")
			}
		}

		// Simple Field (listOfCovSubscriptionSpecifications)
		if pushErr := writeBuffer.PushContext("listOfCovSubscriptionSpecifications"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfCovSubscriptionSpecifications")
		}
		_listOfCovSubscriptionSpecificationsErr := writeBuffer.WriteSerializable(m.GetListOfCovSubscriptionSpecifications())
		if popErr := writeBuffer.PopContext("listOfCovSubscriptionSpecifications"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfCovSubscriptionSpecifications")
		}
		if _listOfCovSubscriptionSpecificationsErr != nil {
			return errors.Wrap(_listOfCovSubscriptionSpecificationsErr, "Error serializing 'listOfCovSubscriptionSpecifications' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) isBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
