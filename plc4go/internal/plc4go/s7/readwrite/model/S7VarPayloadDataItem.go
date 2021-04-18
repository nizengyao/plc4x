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
	"math"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type S7VarPayloadDataItem struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	Data          []int8
}

// The corresponding interface
type IS7VarPayloadDataItem interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewS7VarPayloadDataItem(returnCode DataTransportErrorCode, transportSize DataTransportSize, data []int8) *S7VarPayloadDataItem {
	return &S7VarPayloadDataItem{ReturnCode: returnCode, TransportSize: transportSize, Data: data}
}

func CastS7VarPayloadDataItem(structType interface{}) *S7VarPayloadDataItem {
	castFunc := func(typ interface{}) *S7VarPayloadDataItem {
		if casted, ok := typ.(S7VarPayloadDataItem); ok {
			return &casted
		}
		if casted, ok := typ.(*S7VarPayloadDataItem); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7VarPayloadDataItem) GetTypeName() string {
	return "S7VarPayloadDataItem"
}

func (m *S7VarPayloadDataItem) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7VarPayloadDataItem) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Enum Field (returnCode)
	lengthInBits += 8

	// Enum Field (transportSize)
	lengthInBits += 8

	// Implicit Field (dataLength)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	// Padding Field (padding)
	_timesPadding := uint8(utils.InlineIf(lastItem, func() uint16 { return uint16(uint8(0)) }, func() uint16 { return uint16(uint8(uint8(len(m.Data))) % uint8(uint8(2))) }))
	for ; _timesPadding > 0; _timesPadding-- {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *S7VarPayloadDataItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7VarPayloadDataItemParse(io *utils.ReadBuffer, lastItem bool) (*S7VarPayloadDataItem, error) {

	// Enum field (returnCode)
	returnCode, _returnCodeErr := DataTransportErrorCodeParse(io)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field")
	}

	// Enum field (transportSize)
	transportSize, _transportSizeErr := DataTransportSizeParse(io)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := io.ReadUint16(16)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field")
	}

	// Array field (data)
	// Count array
	data := make([]int8, utils.InlineIf(transportSize.SizeInBits(), func() uint16 { return uint16(math.Ceil(float64(dataLength) / float64(float64(8.0)))) }, func() uint16 { return uint16(dataLength) }))
	for curItem := uint16(0); curItem < uint16(utils.InlineIf(transportSize.SizeInBits(), func() uint16 { return uint16(math.Ceil(float64(dataLength) / float64(float64(8.0)))) }, func() uint16 { return uint16(dataLength) })); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}

	// Padding Field (padding)
	{
		_timesPadding := (utils.InlineIf(lastItem, func() uint16 { return uint16(uint8(0)) }, func() uint16 { return uint16(uint8(uint8(len(data))) % uint8(uint8(2))) }))
		for ; (io.HasMore(8)) && (_timesPadding > 0); _timesPadding-- {
			// Just read the padding data and ignore it
			_, _err := io.ReadUint8(8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'padding' field")
			}
		}
	}

	// Create the instance
	return NewS7VarPayloadDataItem(returnCode, transportSize, data), nil
}

func (m *S7VarPayloadDataItem) Serialize(io utils.WriteBuffer, lastItem bool) error {

	// Enum field (returnCode)
	returnCode := CastDataTransportErrorCode(m.ReturnCode)
	_returnCodeErr := returnCode.Serialize(io)
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Enum field (transportSize)
	transportSize := CastDataTransportSize(m.TransportSize)
	_transportSizeErr := transportSize.Serialize(io)
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength := uint16(uint16(uint16(len(m.Data))) * uint16(uint16(utils.InlineIf(bool(bool((m.TransportSize) == (DataTransportSize_BIT))), func() uint16 { return uint16(uint16(1)) }, func() uint16 {
		return uint16(uint16(utils.InlineIf(transportSize.SizeInBits(), func() uint16 { return uint16(uint16(8)) }, func() uint16 { return uint16(uint16(1)) })))
	}))))
	_dataLengthErr := io.WriteUint16(16, (dataLength))
	if _dataLengthErr != nil {
		return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
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

	// Padding Field (padding)
	{
		_timesPadding := uint8(utils.InlineIf(lastItem, func() uint16 { return uint16(uint8(0)) }, func() uint16 { return uint16(uint8(uint8(len(m.Data))) % uint8(uint8(2))) }))
		for ; _timesPadding > 0; _timesPadding-- {
			_paddingValue := uint8(uint8(0))
			_paddingErr := io.WriteUint8(8, (_paddingValue))
			if _paddingErr != nil {
				return errors.Wrap(_paddingErr, "Error serializing 'padding' field")
			}
		}
	}

	return nil
}

func (m *S7VarPayloadDataItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "returnCode":
				var data DataTransportErrorCode
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ReturnCode = data
			case "transportSize":
				var data DataTransportSize
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TransportSize = data
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
	}
}

func (m *S7VarPayloadDataItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.s7.readwrite.S7VarPayloadDataItem"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ReturnCode, xml.StartElement{Name: xml.Name{Local: "returnCode"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TransportSize, xml.StartElement{Name: xml.Name{Local: "transportSize"}}); err != nil {
		return err
	}
	_encodedData := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Data))
	_encodedData = strings.ToUpper(_encodedData)
	if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m S7VarPayloadDataItem) String() string {
	return string(m.Box("", 120))
}

func (m S7VarPayloadDataItem) Box(name string, width int) utils.AsciiBox {
	boxName := "S7VarPayloadDataItem"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Enum field (returnCode)
	returnCode := CastDataTransportErrorCode(m.ReturnCode)
	boxes = append(boxes, returnCode.Box("returnCode", -1))
	// Enum field (transportSize)
	transportSize := CastDataTransportSize(m.TransportSize)
	boxes = append(boxes, transportSize.Box("transportSize", -1))
	// Implicit Field (dataLength)
	dataLength := uint16(uint16(uint16(len(m.Data))) * uint16(uint16(utils.InlineIf(bool(bool((m.TransportSize) == (DataTransportSize_BIT))), func() uint16 { return uint16(uint16(1)) }, func() uint16 {
		return uint16(uint16(utils.InlineIf(transportSize.SizeInBits(), func() uint16 { return uint16(uint16(8)) }, func() uint16 { return uint16(uint16(1)) })))
	}))))
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("DataLength", dataLength, -1))
	// Array Field (data)
	if m.Data != nil {
		// Simple array base type int8 will be rendered one by one
		arrayBoxes := make([]utils.AsciiBox, 0)
		for _, _element := range m.Data {
			arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
		}
		boxes = append(boxes, utils.BoxBox("Data", utils.AlignBoxes(arrayBoxes, width-4), 0))
	}
	// Padding Field (padding)
	lastItem := false // TODO: not sure if its worth to calculate here
	_ = lastItem
	_timesPadding := uint8(utils.InlineIf(lastItem, func() uint16 { return uint16(uint8(0)) }, func() uint16 { return uint16(uint8(uint8(len(m.Data))) % uint8(uint8(2))) }))
	for ; _timesPadding > 0; _timesPadding-- {
		_paddingValue := uint8(uint8(0))
		boxes = append(boxes, utils.BoxAnything("padding", _paddingValue, -1))
	}
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
