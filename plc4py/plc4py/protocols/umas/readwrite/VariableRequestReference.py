#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

from dataclasses import dataclass

from plc4py.api.exceptions.exceptions import PlcRuntimeException
from plc4py.api.exceptions.exceptions import SerializationException
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.spi.generation.ReadBuffer import ReadBuffer
from plc4py.spi.generation.WriteBuffer import WriteBuffer
import math


@dataclass
class VariableRequestReference:
    is_array: int
    data_size_index: int
    block: int
    base_offset: int
    offset: int
    array_length: int
    UNKNOWN1: int = 0x01

    def serialize(self, write_buffer: WriteBuffer):
        write_buffer.push_context("VariableRequestReference")

        # Simple Field (isArray)
        write_buffer.write_unsigned_byte(
            self.is_array, bit_length=4, logical_name="isArray"
        )

        # Simple Field (dataSizeIndex)
        write_buffer.write_unsigned_byte(
            self.data_size_index, bit_length=4, logical_name="dataSizeIndex"
        )

        # Simple Field (block)
        write_buffer.write_unsigned_short(
            self.block, bit_length=16, logical_name="block"
        )

        # Const Field (unknown1)
        write_buffer.write_unsigned_byte(self.UNKNOWN1, logical_name="unknown1")

        # Simple Field (baseOffset)
        write_buffer.write_unsigned_short(
            self.base_offset, bit_length=16, logical_name="baseOffset"
        )

        # Simple Field (offset)
        write_buffer.write_unsigned_byte(
            self.offset, bit_length=8, logical_name="offset"
        )

        # Optional Field (arrayLength) (Can be skipped, if the value is null)
        write_buffer.write_unsigned_short(self.array_length, logical_name="arrayLength")

        write_buffer.pop_context("VariableRequestReference")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.length_in_bits() / 8.0)))

    def length_in_bits(self) -> int:
        length_in_bits: int = 0
        _value: VariableRequestReference = self

        # Simple field (isArray)
        length_in_bits += 4

        # Simple field (dataSizeIndex)
        length_in_bits += 4

        # Simple field (block)
        length_in_bits += 16

        # Const Field (unknown1)
        length_in_bits += 8

        # Simple field (baseOffset)
        length_in_bits += 16

        # Simple field (offset)
        length_in_bits += 8

        # Optional Field (arrayLength)
        length_in_bits += 16

        return length_in_bits

    @staticmethod
    def static_parse(read_buffer: ReadBuffer, **kwargs):
        return VariableRequestReference.static_parse_context(read_buffer)

    @staticmethod
    def static_parse_context(read_buffer: ReadBuffer):
        read_buffer.push_context("VariableRequestReference")

        is_array: int = read_buffer.read_unsigned_byte(
            logical_name="isArray", bit_length=4
        )

        data_size_index: int = read_buffer.read_unsigned_byte(
            logical_name="dataSizeIndex", bit_length=4
        )

        block: int = read_buffer.read_unsigned_short(
            logical_name="block", bit_length=16
        )

        UNKNOWN1: int = read_buffer.read_unsigned_byte(logical_name="unknown1")

        base_offset: int = read_buffer.read_unsigned_short(
            logical_name="baseOffset", bit_length=16
        )

        offset: int = read_buffer.read_unsigned_byte(
            logical_name="offset", bit_length=8
        )

        array_length: int = None
        if self.is_array:
            array_length = read_buffer.read_unsigned_short(logical_name="arrayLength")

        read_buffer.pop_context("VariableRequestReference")
        # Create the instance
        _variable_request_reference: VariableRequestReference = (
            VariableRequestReference(
                is_array, data_size_index, block, base_offset, offset, array_length
            )
        )
        return _variable_request_reference

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, VariableRequestReference):
            return False

        that: VariableRequestReference = VariableRequestReference(o)
        return (
            (self.is_array == that.is_array)
            and (self.data_size_index == that.data_size_index)
            and (self.block == that.block)
            and (self.base_offset == that.base_offset)
            and (self.offset == that.offset)
            and (self.array_length == that.array_length)
            and True
        )

    def hash_code(self) -> int:
        return hash(self)

    def __str__(self) -> str:
        pass
        # write_buffer_box_based: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        # try:
        #    write_buffer_box_based.writeSerializable(self)
        # except SerializationException as e:
        #    raise PlcRuntimeException(e)

        # return "\n" + str(write_buffer_box_based.get_box()) + "\n"
