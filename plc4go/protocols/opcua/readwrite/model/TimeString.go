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

// TimeString is the corresponding interface of TimeString
type TimeString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// TimeStringExactly can be used when we want exactly this type and not a type which fulfills TimeString.
// This is useful for switch cases.
type TimeStringExactly interface {
	TimeString
	isTimeString() bool
}

// _TimeString is the data-structure of this message
type _TimeString struct {
}

// NewTimeString factory function for _TimeString
func NewTimeString() *_TimeString {
	return &_TimeString{}
}

// Deprecated: use the interface for direct cast
func CastTimeString(structType any) TimeString {
	if casted, ok := structType.(TimeString); ok {
		return casted
	}
	if casted, ok := structType.(*TimeString); ok {
		return *casted
	}
	return nil
}

func (m *_TimeString) GetTypeName() string {
	return "TimeString"
}

func (m *_TimeString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_TimeString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TimeStringParse(ctx context.Context, theBytes []byte) (TimeString, error) {
	return TimeStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TimeStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TimeString, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TimeString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TimeString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TimeString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TimeString")
	}

	// Create the instance
	return &_TimeString{}, nil
}

func (m *_TimeString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TimeString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TimeString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TimeString")
	}

	if popErr := writeBuffer.PopContext("TimeString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TimeString")
	}
	return nil
}

func (m *_TimeString) isTimeString() bool {
	return true
}

func (m *_TimeString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
