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
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BVLCForwardedNPDU struct {
	Ip     []uint8
	Port   uint16
	Npdu   *NPDU
	Parent *BVLC
}

// The corresponding interface
type IBVLCForwardedNPDU interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCForwardedNPDU) BvlcFunction() uint8 {
	return 0x04
}

func (m *BVLCForwardedNPDU) InitializeParent(parent *BVLC) {
}

func NewBVLCForwardedNPDU(ip []uint8, port uint16, npdu *NPDU) *BVLC {
	child := &BVLCForwardedNPDU{
		Ip:     ip,
		Port:   port,
		Npdu:   npdu,
		Parent: NewBVLC(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBVLCForwardedNPDU(structType interface{}) *BVLCForwardedNPDU {
	castFunc := func(typ interface{}) *BVLCForwardedNPDU {
		if casted, ok := typ.(BVLCForwardedNPDU); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCForwardedNPDU); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCForwardedNPDU(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCForwardedNPDU(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCForwardedNPDU) GetTypeName() string {
	return "BVLCForwardedNPDU"
}

func (m *BVLCForwardedNPDU) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCForwardedNPDU) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Simple field (npdu)
	lengthInBits += m.Npdu.LengthInBits()

	return lengthInBits
}

func (m *BVLCForwardedNPDU) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCForwardedNPDUParse(io *utils.ReadBuffer, bvlcLength uint16) (*BVLC, error) {

	// Array field (ip)
	// Count array
	ip := make([]uint8, uint16(4))
	for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
		_item, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'ip' field")
		}
		ip[curItem] = _item
	}

	// Simple Field (port)
	port, _portErr := io.ReadUint16(16)
	if _portErr != nil {
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field")
	}

	// Simple Field (npdu)
	npdu, _npduErr := NPDUParse(io, uint16(bvlcLength)-uint16(uint16(10)))
	if _npduErr != nil {
		return nil, errors.Wrap(_npduErr, "Error parsing 'npdu' field")
	}

	// Create a partially initialized instance
	_child := &BVLCForwardedNPDU{
		Ip:     ip,
		Port:   port,
		Npdu:   npdu,
		Parent: &BVLC{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BVLCForwardedNPDU) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Array Field (ip)
		if m.Ip != nil {
			for _, _element := range m.Ip {
				_elementErr := io.WriteUint8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'ip' field")
				}
			}
		}

		// Simple Field (port)
		port := uint16(m.Port)
		_portErr := io.WriteUint16(16, (port))
		if _portErr != nil {
			return errors.Wrap(_portErr, "Error serializing 'port' field")
		}

		// Simple Field (npdu)
		_npduErr := m.Npdu.Serialize(io)
		if _npduErr != nil {
			return errors.Wrap(_npduErr, "Error serializing 'npdu' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BVLCForwardedNPDU) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "ip":
				var data []uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Ip = data
			case "port":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Port = data
			case "npdu":
				var data NPDU
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Npdu = &data
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

func (m *BVLCForwardedNPDU) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Ip, xml.StartElement{Name: xml.Name{Local: "ip"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Port, xml.StartElement{Name: xml.Name{Local: "port"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Npdu, xml.StartElement{Name: xml.Name{Local: "npdu"}}); err != nil {
		return err
	}
	return nil
}

func (m BVLCForwardedNPDU) String() string {
	return string(m.Box("", 120))
}

func (m BVLCForwardedNPDU) Box(name string, width int) utils.AsciiBox {
	boxName := "BVLCForwardedNPDU"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Array Field (ip)
		if m.Ip != nil {
			// Simple array base type uint8 will be hex dumped
			boxes = append(boxes, utils.BoxedDumpAnything("Ip", m.Ip))
			// Simple array base type uint8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Ip {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Ip", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Port", m.Port, -1))
		// Complex field (case complex)
		boxes = append(boxes, m.Npdu.Box("npdu", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
