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

// VersionTime is the corresponding interface of VersionTime
type VersionTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// VersionTimeExactly can be used when we want exactly this type and not a type which fulfills VersionTime.
// This is useful for switch cases.
type VersionTimeExactly interface {
	VersionTime
	isVersionTime() bool
}

// _VersionTime is the data-structure of this message
type _VersionTime struct {
}

// NewVersionTime factory function for _VersionTime
func NewVersionTime() *_VersionTime {
	return &_VersionTime{}
}

// Deprecated: use the interface for direct cast
func CastVersionTime(structType any) VersionTime {
	if casted, ok := structType.(VersionTime); ok {
		return casted
	}
	if casted, ok := structType.(*VersionTime); ok {
		return *casted
	}
	return nil
}

func (m *_VersionTime) GetTypeName() string {
	return "VersionTime"
}

func (m *_VersionTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_VersionTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func VersionTimeParse(ctx context.Context, theBytes []byte) (VersionTime, error) {
	return VersionTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func VersionTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (VersionTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("VersionTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VersionTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("VersionTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VersionTime")
	}

	// Create the instance
	return &_VersionTime{}, nil
}

func (m *_VersionTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VersionTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("VersionTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for VersionTime")
	}

	if popErr := writeBuffer.PopContext("VersionTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for VersionTime")
	}
	return nil
}

func (m *_VersionTime) isVersionTime() bool {
	return true
}

func (m *_VersionTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
