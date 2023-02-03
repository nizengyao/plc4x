
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

// Code generated by code-generation. DO NOT EDIT.

#include <unity.h>

#include "plc4c/spi/read_buffer.h"
#include "plc4c/spi/write_buffer.h"

#include "tpkt_packet.h"


void internal_assert_arrays_equal(uint8_t* expected_array, uint8_t* actual_array, uint8_t num_bytes);

        
void parser_serializer_test_s7_read_write_cotp_connection_request() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x16, 0x11, 0xe0, 0x00, 0x00, 0x00, 0x0f, 0x00, 0xc2, 0x02, 0x01, 0x00, 0xc1, 0x02, 0x03, 0x11, 0xc0, 0x01, 0x0a
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_s7_read_write_cotp_connection_response() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x16, 0x11, 0xd0, 0x00, 0x0f, 0x00, 0x0b, 0x00, 0xc0, 0x01, 0x0a, 0xc1, 0x02, 0x03, 0x11, 0xc2, 0x02, 0x01, 0x00
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_s7_read_write_s7_setup_communication_request() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x19, 0x02, 0xf0, 0x81, 0x32, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x08, 0x00, 0x08, 0x03, 0xf0
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_s7_read_write_s7_setup_communication_response() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x1b, 0x02, 0xf0, 0x80, 0x32, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x03, 0x00, 0x03, 0x00, 0xf0
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_s7_read_write_s7_read_plc_type_request() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x21, 0x02, 0xf0, 0x82, 0x32, 0x07, 0x00, 0x00, 0x00, 0x01, 0x00, 0x08, 0x00, 0x08, 0x00, 0x01, 0x12, 0x04, 0x11, 0x44, 0x01, 0x00, 0xff, 0x09, 0x00, 0x04, 0x00, 0x11, 0x00, 0x00
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_s7_read_write_s7_read_plc_type_response() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x7d, 0x02, 0xf0, 0x80, 0x32, 0x07, 0x00, 0x00, 0x00, 0x01, 0x00, 0x0c, 0x00, 0x60, 0x00, 0x01, 0x12, 0x08, 0x12, 0x84, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0xff, 0x09, 0x00, 0x5c, 0x00, 0x11, 0x00, 0x00, 0x00, 0x1c, 0x00, 0x03, 0x00, 0x01, 0x36, 0x45, 0x53, 0x37, 0x20, 0x32, 0x31, 0x32, 0x2d, 0x31, 0x42, 0x44, 0x33, 0x30, 0x2d, 0x30, 0x58, 0x42, 0x30, 0x20, 0x20, 0x20, 0x00, 0x01, 0x20, 0x20, 0x00, 0x06, 0x36, 0x45, 0x53, 0x37, 0x20, 0x32, 0x31, 0x32, 0x2d, 0x31, 0x42, 0x44, 0x33, 0x30, 0x2d, 0x30, 0x58, 0x42, 0x30, 0x20, 0x20, 0x20, 0x00, 0x01, 0x20, 0x20, 0x00, 0x07, 0x36, 0x45, 0x53, 0x37, 0x20, 0x32, 0x31, 0x32, 0x2d, 0x31, 0x42, 0x44, 0x33, 0x30, 0x2d, 0x30, 0x58, 0x42, 0x30, 0x20, 0x20, 0x20, 0x56, 0x02, 0x00, 0x02
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_s7_read_write_s7_read_request() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x43, 0x02, 0xf0, 0x8b, 0x32, 0x01, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x32, 0x00, 0x00, 0x04, 0x04, 0x12, 0x0a, 0x10, 0x01, 0x00, 0x01, 0x00, 0x00, 0x82, 0x00, 0x00, 0x00, 0x12, 0x0a, 0x10, 0x01, 0x00, 0x01, 0x00, 0x00, 0x82, 0x00, 0x00, 0x00, 0x12, 0x0a, 0x10, 0x01, 0x00, 0x01, 0x00, 0x00, 0x82, 0x00, 0x00, 0x00, 0x12, 0x0a, 0x10, 0x01, 0x00, 0x01, 0x00, 0x00, 0x82, 0x00, 0x00, 0x00
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_s7_read_write_s7_read_response() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x2d, 0x02, 0xf0, 0x80, 0x32, 0x03, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x02, 0x00, 0x18, 0x00, 0x00, 0x04, 0x04, 0xff, 0x03, 0x00, 0x01, 0x01, 0x00, 0xff, 0x03, 0x00, 0x01, 0x01, 0x00, 0xff, 0x03, 0x00, 0x01, 0x01, 0x00, 0xff, 0x03, 0x00, 0x01, 0x01, 0x00
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_s7_read_write_s7_read_error_response() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x13, 0x02, 0xf0, 0x80, 0x32, 0x02, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x85, 0x00
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_s7_read_write_s7_write_request() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x5b, 0x02, 0xf0, 0x8e, 0x32, 0x01, 0x00, 0x00, 0x00, 0x0e, 0x00, 0x32, 0x00, 0x18, 0x05, 0x04, 0x12, 0x0a, 0x10, 0x01, 0x00, 0x01, 0x00, 0x00, 0x82, 0x00, 0x00, 0x00, 0x12, 0x0a, 0x10, 0x01, 0x00, 0x01, 0x00, 0x00, 0x82, 0x00, 0x00, 0x01, 0x12, 0x0a, 0x10, 0x01, 0x00, 0x01, 0x00, 0x00, 0x82, 0x00, 0x00, 0x02, 0x12, 0x0a, 0x10, 0x01, 0x00, 0x01, 0x00, 0x00, 0x82, 0x00, 0x00, 0x03, 0xff, 0x03, 0x00, 0x01, 0x01, 0x00, 0xff, 0x03, 0x00, 0x01, 0x01, 0x00, 0xff, 0x03, 0x00, 0x01, 0x01, 0x00, 0xff, 0x03, 0x00, 0x01, 0x01, 0x00
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_s7_read_write_s7_write_response() {
            
    uint8_t payload[] = {
        0x03, 0x00, 0x00, 0x19, 0x02, 0xf0, 0x80, 0x32, 0x03, 0x00, 0x00, 0x00, 0x0e, 0x00, 0x02, 0x00, 0x04, 0x00, 0x00, 0x05, 0x04, 0xff, 0xff, 0xff, 0xff
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_s7_read_write_tpkt_packet* message = NULL;
    return_code = plc4c_s7_read_write_tpkt_packet_parse(plc4x_spi_context_background(), read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_s7_read_write_tpkt_packet_serialize(plc4x_spi_context_background(), write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        


void parser_serializer_test_s7_read_write(void) {
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_cotp_connection_request);
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_cotp_connection_response);
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_s7_setup_communication_request);
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_s7_setup_communication_response);
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_s7_read_plc_type_request);
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_s7_read_plc_type_response);
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_s7_read_request);
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_s7_read_response);
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_s7_read_error_response);
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_s7_write_request);
        
    RUN_TEST(
        parser_serializer_test_s7_read_write_s7_write_response);
        
}
    