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

// Code generated by mockery v2.23.4. DO NOT EDIT.

package bacnetip

import (
	model "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	spi "github.com/apache/plc4x/plc4go/spi"
	mock "github.com/stretchr/testify/mock"
)

// mock_PDU is an autogenerated mock type for the _PDU type
type mock_PDU struct {
	mock.Mock
}

type mock_PDU_Expecter struct {
	mock *mock.Mock
}

func (_m *mock_PDU) EXPECT() *mock_PDU_Expecter {
	return &mock_PDU_Expecter{mock: &_m.Mock}
}

// GetExpectingReply provides a mock function with given fields:
func (_m *mock_PDU) GetExpectingReply() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// mock_PDU_GetExpectingReply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExpectingReply'
type mock_PDU_GetExpectingReply_Call struct {
	*mock.Call
}

// GetExpectingReply is a helper method to define mock.On call
func (_e *mock_PDU_Expecter) GetExpectingReply() *mock_PDU_GetExpectingReply_Call {
	return &mock_PDU_GetExpectingReply_Call{Call: _e.mock.On("GetExpectingReply")}
}

func (_c *mock_PDU_GetExpectingReply_Call) Run(run func()) *mock_PDU_GetExpectingReply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_PDU_GetExpectingReply_Call) Return(_a0 bool) *mock_PDU_GetExpectingReply_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_PDU_GetExpectingReply_Call) RunAndReturn(run func() bool) *mock_PDU_GetExpectingReply_Call {
	_c.Call.Return(run)
	return _c
}

// GetMessage provides a mock function with given fields:
func (_m *mock_PDU) GetMessage() spi.Message {
	ret := _m.Called()

	var r0 spi.Message
	if rf, ok := ret.Get(0).(func() spi.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.Message)
		}
	}

	return r0
}

// mock_PDU_GetMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessage'
type mock_PDU_GetMessage_Call struct {
	*mock.Call
}

// GetMessage is a helper method to define mock.On call
func (_e *mock_PDU_Expecter) GetMessage() *mock_PDU_GetMessage_Call {
	return &mock_PDU_GetMessage_Call{Call: _e.mock.On("GetMessage")}
}

func (_c *mock_PDU_GetMessage_Call) Run(run func()) *mock_PDU_GetMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_PDU_GetMessage_Call) Return(_a0 spi.Message) *mock_PDU_GetMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_PDU_GetMessage_Call) RunAndReturn(run func() spi.Message) *mock_PDU_GetMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetNetworkPriority provides a mock function with given fields:
func (_m *mock_PDU) GetNetworkPriority() model.NPDUNetworkPriority {
	ret := _m.Called()

	var r0 model.NPDUNetworkPriority
	if rf, ok := ret.Get(0).(func() model.NPDUNetworkPriority); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.NPDUNetworkPriority)
	}

	return r0
}

// mock_PDU_GetNetworkPriority_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNetworkPriority'
type mock_PDU_GetNetworkPriority_Call struct {
	*mock.Call
}

// GetNetworkPriority is a helper method to define mock.On call
func (_e *mock_PDU_Expecter) GetNetworkPriority() *mock_PDU_GetNetworkPriority_Call {
	return &mock_PDU_GetNetworkPriority_Call{Call: _e.mock.On("GetNetworkPriority")}
}

func (_c *mock_PDU_GetNetworkPriority_Call) Run(run func()) *mock_PDU_GetNetworkPriority_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_PDU_GetNetworkPriority_Call) Return(_a0 model.NPDUNetworkPriority) *mock_PDU_GetNetworkPriority_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_PDU_GetNetworkPriority_Call) RunAndReturn(run func() model.NPDUNetworkPriority) *mock_PDU_GetNetworkPriority_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUDestination provides a mock function with given fields:
func (_m *mock_PDU) GetPDUDestination() *Address {
	ret := _m.Called()

	var r0 *Address
	if rf, ok := ret.Get(0).(func() *Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Address)
		}
	}

	return r0
}

// mock_PDU_GetPDUDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUDestination'
type mock_PDU_GetPDUDestination_Call struct {
	*mock.Call
}

// GetPDUDestination is a helper method to define mock.On call
func (_e *mock_PDU_Expecter) GetPDUDestination() *mock_PDU_GetPDUDestination_Call {
	return &mock_PDU_GetPDUDestination_Call{Call: _e.mock.On("GetPDUDestination")}
}

func (_c *mock_PDU_GetPDUDestination_Call) Run(run func()) *mock_PDU_GetPDUDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_PDU_GetPDUDestination_Call) Return(_a0 *Address) *mock_PDU_GetPDUDestination_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_PDU_GetPDUDestination_Call) RunAndReturn(run func() *Address) *mock_PDU_GetPDUDestination_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUSource provides a mock function with given fields:
func (_m *mock_PDU) GetPDUSource() *Address {
	ret := _m.Called()

	var r0 *Address
	if rf, ok := ret.Get(0).(func() *Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Address)
		}
	}

	return r0
}

// mock_PDU_GetPDUSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUSource'
type mock_PDU_GetPDUSource_Call struct {
	*mock.Call
}

// GetPDUSource is a helper method to define mock.On call
func (_e *mock_PDU_Expecter) GetPDUSource() *mock_PDU_GetPDUSource_Call {
	return &mock_PDU_GetPDUSource_Call{Call: _e.mock.On("GetPDUSource")}
}

func (_c *mock_PDU_GetPDUSource_Call) Run(run func()) *mock_PDU_GetPDUSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_PDU_GetPDUSource_Call) Return(_a0 *Address) *mock_PDU_GetPDUSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_PDU_GetPDUSource_Call) RunAndReturn(run func() *Address) *mock_PDU_GetPDUSource_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUDestination provides a mock function with given fields: _a0
func (_m *mock_PDU) SetPDUDestination(_a0 *Address) {
	_m.Called(_a0)
}

// mock_PDU_SetPDUDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUDestination'
type mock_PDU_SetPDUDestination_Call struct {
	*mock.Call
}

// SetPDUDestination is a helper method to define mock.On call
//   - _a0 *Address
func (_e *mock_PDU_Expecter) SetPDUDestination(_a0 interface{}) *mock_PDU_SetPDUDestination_Call {
	return &mock_PDU_SetPDUDestination_Call{Call: _e.mock.On("SetPDUDestination", _a0)}
}

func (_c *mock_PDU_SetPDUDestination_Call) Run(run func(_a0 *Address)) *mock_PDU_SetPDUDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Address))
	})
	return _c
}

func (_c *mock_PDU_SetPDUDestination_Call) Return() *mock_PDU_SetPDUDestination_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_PDU_SetPDUDestination_Call) RunAndReturn(run func(*Address)) *mock_PDU_SetPDUDestination_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *mock_PDU) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// mock_PDU_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type mock_PDU_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *mock_PDU_Expecter) String() *mock_PDU_String_Call {
	return &mock_PDU_String_Call{Call: _e.mock.On("String")}
}

func (_c *mock_PDU_String_Call) Run(run func()) *mock_PDU_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_PDU_String_Call) Return(_a0 string) *mock_PDU_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_PDU_String_Call) RunAndReturn(run func() string) *mock_PDU_String_Call {
	_c.Call.Return(run)
	return _c
}

// newMock_PDU creates a new instance of mock_PDU. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMock_PDU(t interface {
	mock.TestingT
	Cleanup(func())
}) *mock_PDU {
	mock := &mock_PDU{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
