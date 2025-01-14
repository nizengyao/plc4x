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

// NLMRejectMessageToNetworkRejectReason is an enum
type NLMRejectMessageToNetworkRejectReason uint8

type INLMRejectMessageToNetworkRejectReason interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	NLMRejectMessageToNetworkRejectReason_OTHER                  NLMRejectMessageToNetworkRejectReason = 0
	NLMRejectMessageToNetworkRejectReason_NOT_DIRECTLY_CONNECTED NLMRejectMessageToNetworkRejectReason = 1
	NLMRejectMessageToNetworkRejectReason_BUSY                   NLMRejectMessageToNetworkRejectReason = 2
	NLMRejectMessageToNetworkRejectReason_UNKNOWN_NLMT           NLMRejectMessageToNetworkRejectReason = 3
	NLMRejectMessageToNetworkRejectReason_TOO_LONG               NLMRejectMessageToNetworkRejectReason = 4
	NLMRejectMessageToNetworkRejectReason_SECURITY_ERROR         NLMRejectMessageToNetworkRejectReason = 5
	NLMRejectMessageToNetworkRejectReason_ADDRESSING_ERROR       NLMRejectMessageToNetworkRejectReason = 6
)

var NLMRejectMessageToNetworkRejectReasonValues []NLMRejectMessageToNetworkRejectReason

func init() {
	_ = errors.New
	NLMRejectMessageToNetworkRejectReasonValues = []NLMRejectMessageToNetworkRejectReason{
		NLMRejectMessageToNetworkRejectReason_OTHER,
		NLMRejectMessageToNetworkRejectReason_NOT_DIRECTLY_CONNECTED,
		NLMRejectMessageToNetworkRejectReason_BUSY,
		NLMRejectMessageToNetworkRejectReason_UNKNOWN_NLMT,
		NLMRejectMessageToNetworkRejectReason_TOO_LONG,
		NLMRejectMessageToNetworkRejectReason_SECURITY_ERROR,
		NLMRejectMessageToNetworkRejectReason_ADDRESSING_ERROR,
	}
}

func NLMRejectMessageToNetworkRejectReasonByValue(value uint8) (enum NLMRejectMessageToNetworkRejectReason, ok bool) {
	switch value {
	case 0:
		return NLMRejectMessageToNetworkRejectReason_OTHER, true
	case 1:
		return NLMRejectMessageToNetworkRejectReason_NOT_DIRECTLY_CONNECTED, true
	case 2:
		return NLMRejectMessageToNetworkRejectReason_BUSY, true
	case 3:
		return NLMRejectMessageToNetworkRejectReason_UNKNOWN_NLMT, true
	case 4:
		return NLMRejectMessageToNetworkRejectReason_TOO_LONG, true
	case 5:
		return NLMRejectMessageToNetworkRejectReason_SECURITY_ERROR, true
	case 6:
		return NLMRejectMessageToNetworkRejectReason_ADDRESSING_ERROR, true
	}
	return 0, false
}

func NLMRejectMessageToNetworkRejectReasonByName(value string) (enum NLMRejectMessageToNetworkRejectReason, ok bool) {
	switch value {
	case "OTHER":
		return NLMRejectMessageToNetworkRejectReason_OTHER, true
	case "NOT_DIRECTLY_CONNECTED":
		return NLMRejectMessageToNetworkRejectReason_NOT_DIRECTLY_CONNECTED, true
	case "BUSY":
		return NLMRejectMessageToNetworkRejectReason_BUSY, true
	case "UNKNOWN_NLMT":
		return NLMRejectMessageToNetworkRejectReason_UNKNOWN_NLMT, true
	case "TOO_LONG":
		return NLMRejectMessageToNetworkRejectReason_TOO_LONG, true
	case "SECURITY_ERROR":
		return NLMRejectMessageToNetworkRejectReason_SECURITY_ERROR, true
	case "ADDRESSING_ERROR":
		return NLMRejectMessageToNetworkRejectReason_ADDRESSING_ERROR, true
	}
	return 0, false
}

func NLMRejectMessageToNetworkRejectReasonKnows(value uint8) bool {
	for _, typeValue := range NLMRejectMessageToNetworkRejectReasonValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastNLMRejectMessageToNetworkRejectReason(structType any) NLMRejectMessageToNetworkRejectReason {
	castFunc := func(typ any) NLMRejectMessageToNetworkRejectReason {
		if sNLMRejectMessageToNetworkRejectReason, ok := typ.(NLMRejectMessageToNetworkRejectReason); ok {
			return sNLMRejectMessageToNetworkRejectReason
		}
		return 0
	}
	return castFunc(structType)
}

func (m NLMRejectMessageToNetworkRejectReason) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m NLMRejectMessageToNetworkRejectReason) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMRejectMessageToNetworkRejectReasonParse(ctx context.Context, theBytes []byte) (NLMRejectMessageToNetworkRejectReason, error) {
	return NLMRejectMessageToNetworkRejectReasonParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NLMRejectMessageToNetworkRejectReasonParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NLMRejectMessageToNetworkRejectReason, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("NLMRejectMessageToNetworkRejectReason", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading NLMRejectMessageToNetworkRejectReason")
	}
	if enum, ok := NLMRejectMessageToNetworkRejectReasonByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for NLMRejectMessageToNetworkRejectReason")
		return NLMRejectMessageToNetworkRejectReason(val), nil
	} else {
		return enum, nil
	}
}

func (e NLMRejectMessageToNetworkRejectReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e NLMRejectMessageToNetworkRejectReason) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("NLMRejectMessageToNetworkRejectReason", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e NLMRejectMessageToNetworkRejectReason) PLC4XEnumName() string {
	switch e {
	case NLMRejectMessageToNetworkRejectReason_OTHER:
		return "OTHER"
	case NLMRejectMessageToNetworkRejectReason_NOT_DIRECTLY_CONNECTED:
		return "NOT_DIRECTLY_CONNECTED"
	case NLMRejectMessageToNetworkRejectReason_BUSY:
		return "BUSY"
	case NLMRejectMessageToNetworkRejectReason_UNKNOWN_NLMT:
		return "UNKNOWN_NLMT"
	case NLMRejectMessageToNetworkRejectReason_TOO_LONG:
		return "TOO_LONG"
	case NLMRejectMessageToNetworkRejectReason_SECURITY_ERROR:
		return "SECURITY_ERROR"
	case NLMRejectMessageToNetworkRejectReason_ADDRESSING_ERROR:
		return "ADDRESSING_ERROR"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e NLMRejectMessageToNetworkRejectReason) String() string {
	return e.PLC4XEnumName()
}
