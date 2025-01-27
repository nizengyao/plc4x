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

// ChunkType is an enum
type ChunkType string

type IChunkType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ChunkType_CONTINUE ChunkType = "C"
	ChunkType_FINAL    ChunkType = "F"
	ChunkType_ABORT    ChunkType = "A"
)

var ChunkTypeValues []ChunkType

func init() {
	_ = errors.New
	ChunkTypeValues = []ChunkType{
		ChunkType_CONTINUE,
		ChunkType_FINAL,
		ChunkType_ABORT,
	}
}

func ChunkTypeByValue(value string) (enum ChunkType, ok bool) {
	switch value {
	case "A":
		return ChunkType_ABORT, true
	case "C":
		return ChunkType_CONTINUE, true
	case "F":
		return ChunkType_FINAL, true
	}
	return "", false
}

func ChunkTypeByName(value string) (enum ChunkType, ok bool) {
	switch value {
	case "ABORT":
		return ChunkType_ABORT, true
	case "CONTINUE":
		return ChunkType_CONTINUE, true
	case "FINAL":
		return ChunkType_FINAL, true
	}
	return "", false
}

func ChunkTypeKnows(value string) bool {
	for _, typeValue := range ChunkTypeValues {
		if string(typeValue) == value {
			return true
		}
	}
	return false
}

func CastChunkType(structType any) ChunkType {
	castFunc := func(typ any) ChunkType {
		if sChunkType, ok := typ.(ChunkType); ok {
			return sChunkType
		}
		return ""
	}
	return castFunc(structType)
}

func (m ChunkType) GetLengthInBits(ctx context.Context) uint16 {
	return 0
}

func (m ChunkType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ChunkTypeParse(ctx context.Context, theBytes []byte) (ChunkType, error) {
	return ChunkTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ChunkTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ChunkType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadString("ChunkType", uint32(8), "UTF-8")
	if err != nil {
		return "", errors.Wrap(err, "error reading ChunkType")
	}
	if enum, ok := ChunkTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ChunkType")
		return ChunkType(val), nil
	} else {
		return enum, nil
	}
}

func (e ChunkType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ChunkType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteString("ChunkType", uint32(8), "UTF-8", string(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ChunkType) PLC4XEnumName() string {
	switch e {
	case ChunkType_ABORT:
		return "ABORT"
	case ChunkType_CONTINUE:
		return "CONTINUE"
	case ChunkType_FINAL:
		return "FINAL"
	}
	return fmt.Sprintf("Unknown(%v)", string(e))
}

func (e ChunkType) String() string {
	return e.PLC4XEnumName()
}
