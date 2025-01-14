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

package bacnetip

// NPDUTypes is a dictionary of message type values and structs
var NPDUTypes map[uint8]func() interface{ Decode(Arg) error }

func init() {
	NPDUTypes = map[uint8]func() interface{ Decode(Arg) error }{
		0x00: func() interface{ Decode(Arg) error } {
			v, _ := NewWhoIsRouterToNetwork()
			return v
		},
		0x01: func() interface{ Decode(Arg) error } {
			v, _ := NewIAmRouterToNetwork()
			return v
		},
		0x02: func() interface{ Decode(Arg) error } {
			v, _ := NewICouldBeRouterToNetwork()
			return v
		},
		0x03: func() interface{ Decode(Arg) error } {
			v, _ := NewRejectMessageToNetwork()
			return v
		},
		0x04: func() interface{ Decode(Arg) error } {
			v, _ := NewRouterBusyToNetwork()
			return v
		},
		0x05: func() interface{ Decode(Arg) error } {
			v, _ := NewRouterAvailableToNetwork()
			return v
		},
		0x06: func() interface{ Decode(Arg) error } {
			v, _ := NewInitializeRoutingTable()
			return v
		},
		0x07: func() interface{ Decode(Arg) error } {
			v, _ := NewInitializeRoutingTableAck()
			return v
		},
		0x08: func() interface{ Decode(Arg) error } {
			v, _ := NewEstablishConnectionToNetwork()
			return v
		},
		0x09: func() interface{ Decode(Arg) error } {
			v, _ := NewDisconnectConnectionToNetwork()
			return v
		},
		// 0x0A: NewChallengeRequest, // TODO: not present upstream
		// 0x0B: NewSecurityPayload, // TODO: not present upstream
		// 0x0C: NewSecurityResponse, // TODO: not present upstream
		// 0x0D: NewRequestKeyUpdate, // TODO: not present upstream
		// 0x0E: NewUpdateKeyUpdate, // TODO: not present upstream
		// 0x0F: NewUpdateKeyDistributionKey, // TODO: not present upstream
		// 0x10: NewRequestMasterKey, // TODO: not present upstream
		// 0x11: NewSetMasterKey, // TODO: not present upstream
		0x12: func() interface{ Decode(Arg) error } {
			v, _ := NewWhatIsNetworkNumber()
			return v
		},
		0x13: func() interface{ Decode(Arg) error } {
			v, _ := NewNetworkNumberIs()
			return v
		},
	}
}
