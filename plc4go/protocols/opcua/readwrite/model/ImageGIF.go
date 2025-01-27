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

// ImageGIF is the corresponding interface of ImageGIF
type ImageGIF interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// ImageGIFExactly can be used when we want exactly this type and not a type which fulfills ImageGIF.
// This is useful for switch cases.
type ImageGIFExactly interface {
	ImageGIF
	isImageGIF() bool
}

// _ImageGIF is the data-structure of this message
type _ImageGIF struct {
}

// NewImageGIF factory function for _ImageGIF
func NewImageGIF() *_ImageGIF {
	return &_ImageGIF{}
}

// Deprecated: use the interface for direct cast
func CastImageGIF(structType any) ImageGIF {
	if casted, ok := structType.(ImageGIF); ok {
		return casted
	}
	if casted, ok := structType.(*ImageGIF); ok {
		return *casted
	}
	return nil
}

func (m *_ImageGIF) GetTypeName() string {
	return "ImageGIF"
}

func (m *_ImageGIF) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ImageGIF) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ImageGIFParse(ctx context.Context, theBytes []byte) (ImageGIF, error) {
	return ImageGIFParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ImageGIFParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ImageGIF, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ImageGIF"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ImageGIF")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ImageGIF"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ImageGIF")
	}

	// Create the instance
	return &_ImageGIF{}, nil
}

func (m *_ImageGIF) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ImageGIF) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ImageGIF"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ImageGIF")
	}

	if popErr := writeBuffer.PopContext("ImageGIF"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ImageGIF")
	}
	return nil
}

func (m *_ImageGIF) isImageGIF() bool {
	return true
}

func (m *_ImageGIF) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
