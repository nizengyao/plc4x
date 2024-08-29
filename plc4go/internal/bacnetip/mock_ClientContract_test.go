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

// Code generated by mockery v2.42.2. DO NOT EDIT.

package bacnetip

import mock "github.com/stretchr/testify/mock"

// MockClientContract is an autogenerated mock type for the ClientContract type
type MockClientContract struct {
	mock.Mock
}

type MockClientContract_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientContract) EXPECT() *MockClientContract_Expecter {
	return &MockClientContract_Expecter{mock: &_m.Mock}
}

// Confirmation provides a mock function with given fields: args, kwargs
func (_m *MockClientContract) Confirmation(args Args, kwargs KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Confirmation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClientContract_Confirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Confirmation'
type MockClientContract_Confirmation_Call struct {
	*mock.Call
}

// Confirmation is a helper method to define mock.On call
//   - args Args
//   - kwargs KWArgs
func (_e *MockClientContract_Expecter) Confirmation(args interface{}, kwargs interface{}) *MockClientContract_Confirmation_Call {
	return &MockClientContract_Confirmation_Call{Call: _e.mock.On("Confirmation", args, kwargs)}
}

func (_c *MockClientContract_Confirmation_Call) Run(run func(args Args, kwargs KWArgs)) *MockClientContract_Confirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *MockClientContract_Confirmation_Call) Return(_a0 error) *MockClientContract_Confirmation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClientContract_Confirmation_Call) RunAndReturn(run func(Args, KWArgs) error) *MockClientContract_Confirmation_Call {
	_c.Call.Return(run)
	return _c
}

// Request provides a mock function with given fields: args, kwargs
func (_m *MockClientContract) Request(args Args, kwargs KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Request")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClientContract_Request_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Request'
type MockClientContract_Request_Call struct {
	*mock.Call
}

// Request is a helper method to define mock.On call
//   - args Args
//   - kwargs KWArgs
func (_e *MockClientContract_Expecter) Request(args interface{}, kwargs interface{}) *MockClientContract_Request_Call {
	return &MockClientContract_Request_Call{Call: _e.mock.On("Request", args, kwargs)}
}

func (_c *MockClientContract_Request_Call) Run(run func(args Args, kwargs KWArgs)) *MockClientContract_Request_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *MockClientContract_Request_Call) Return(_a0 error) *MockClientContract_Request_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClientContract_Request_Call) RunAndReturn(run func(Args, KWArgs) error) *MockClientContract_Request_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockClientContract) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockClientContract_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockClientContract_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockClientContract_Expecter) String() *MockClientContract_String_Call {
	return &MockClientContract_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockClientContract_String_Call) Run(run func()) *MockClientContract_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClientContract_String_Call) Return(_a0 string) *MockClientContract_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClientContract_String_Call) RunAndReturn(run func() string) *MockClientContract_String_Call {
	_c.Call.Return(run)
	return _c
}

// _setClientPeer provides a mock function with given fields: server
func (_m *MockClientContract) _setClientPeer(server Server) {
	_m.Called(server)
}

// MockClientContract__setClientPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method '_setClientPeer'
type MockClientContract__setClientPeer_Call struct {
	*mock.Call
}

// _setClientPeer is a helper method to define mock.On call
//   - server Server
func (_e *MockClientContract_Expecter) _setClientPeer(server interface{}) *MockClientContract__setClientPeer_Call {
	return &MockClientContract__setClientPeer_Call{Call: _e.mock.On("_setClientPeer", server)}
}

func (_c *MockClientContract__setClientPeer_Call) Run(run func(server Server)) *MockClientContract__setClientPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Server))
	})
	return _c
}

func (_c *MockClientContract__setClientPeer_Call) Return() *MockClientContract__setClientPeer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockClientContract__setClientPeer_Call) RunAndReturn(run func(Server)) *MockClientContract__setClientPeer_Call {
	_c.Call.Return(run)
	return _c
}

// getClientId provides a mock function with given fields:
func (_m *MockClientContract) getClientId() *int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getClientId")
	}

	var r0 *int
	if rf, ok := ret.Get(0).(func() *int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	return r0
}

// MockClientContract_getClientId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getClientId'
type MockClientContract_getClientId_Call struct {
	*mock.Call
}

// getClientId is a helper method to define mock.On call
func (_e *MockClientContract_Expecter) getClientId() *MockClientContract_getClientId_Call {
	return &MockClientContract_getClientId_Call{Call: _e.mock.On("getClientId")}
}

func (_c *MockClientContract_getClientId_Call) Run(run func()) *MockClientContract_getClientId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClientContract_getClientId_Call) Return(_a0 *int) *MockClientContract_getClientId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClientContract_getClientId_Call) RunAndReturn(run func() *int) *MockClientContract_getClientId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClientContract creates a new instance of MockClientContract. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientContract(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientContract {
	mock := &MockClientContract{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}