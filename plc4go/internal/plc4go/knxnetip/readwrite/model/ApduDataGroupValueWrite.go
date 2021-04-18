//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduDataGroupValueWrite struct {
	DataFirstByte int8
	Data          []int8
	Parent        *ApduData
}

// The corresponding interface
type IApduDataGroupValueWrite interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataGroupValueWrite) ApciType() uint8 {
	return 0x2
}

func (m *ApduDataGroupValueWrite) InitializeParent(parent *ApduData) {
}

func NewApduDataGroupValueWrite(dataFirstByte int8, data []int8) *ApduData {
	child := &ApduDataGroupValueWrite{
		DataFirstByte: dataFirstByte,
		Data:          data,
		Parent:        NewApduData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataGroupValueWrite(structType interface{}) *ApduDataGroupValueWrite {
	castFunc := func(typ interface{}) *ApduDataGroupValueWrite {
		if casted, ok := typ.(ApduDataGroupValueWrite); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataGroupValueWrite); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataGroupValueWrite(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataGroupValueWrite(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataGroupValueWrite) GetTypeName() string {
	return "ApduDataGroupValueWrite"
}

func (m *ApduDataGroupValueWrite) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataGroupValueWrite) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (dataFirstByte)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataGroupValueWrite) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataGroupValueWriteParse(io *utils.ReadBuffer, dataLength uint8) (*ApduData, error) {

	// Simple Field (dataFirstByte)
	dataFirstByte, _dataFirstByteErr := io.ReadInt8(6)
	if _dataFirstByteErr != nil {
		return nil, errors.Wrap(_dataFirstByteErr, "Error parsing 'dataFirstByte' field")
	}

	// Array field (data)
	// Count array
	data := make([]int8, utils.InlineIf(bool(bool((dataLength) < (1))), func() uint16 { return uint16(uint16(0)) }, func() uint16 { return uint16(uint16(dataLength) - uint16(uint16(1))) }))
	for curItem := uint16(0); curItem < uint16(utils.InlineIf(bool(bool((dataLength) < (1))), func() uint16 { return uint16(uint16(0)) }, func() uint16 { return uint16(uint16(dataLength) - uint16(uint16(1))) })); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &ApduDataGroupValueWrite{
		DataFirstByte: dataFirstByte,
		Data:          data,
		Parent:        &ApduData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataGroupValueWrite) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (dataFirstByte)
		dataFirstByte := int8(m.DataFirstByte)
		_dataFirstByteErr := io.WriteInt8(6, (dataFirstByte))
		if _dataFirstByteErr != nil {
			return errors.Wrap(_dataFirstByteErr, "Error serializing 'dataFirstByte' field")
		}

		// Array Field (data)
		if m.Data != nil {
			for _, _element := range m.Data {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataGroupValueWrite) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "dataFirstByte":
				var data int8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DataFirstByte = data
			case "data":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *ApduDataGroupValueWrite) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.DataFirstByte, xml.StartElement{Name: xml.Name{Local: "dataFirstByte"}}); err != nil {
		return err
	}
	_encodedData := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Data))
	_encodedData = strings.ToUpper(_encodedData)
	if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m ApduDataGroupValueWrite) String() string {
	return string(m.Box("", 120))
}

func (m ApduDataGroupValueWrite) Box(name string, width int) utils.AsciiBox {
	boxName := "ApduDataGroupValueWrite"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// int8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("DataFirstByte", m.DataFirstByte, -1))
		// Array Field (data)
		if m.Data != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Data {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Data", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
