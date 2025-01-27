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

package tests

import (
	bacnetip "github.com/apache/plc4x/plc4go/internal/bacnetip"
	mock "github.com/stretchr/testify/mock"
)

// MockTrappedApplicationServiceElementRequirements is an autogenerated mock type for the TrappedApplicationServiceElementRequirements type
type MockTrappedApplicationServiceElementRequirements struct {
	mock.Mock
}

type MockTrappedApplicationServiceElementRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTrappedApplicationServiceElementRequirements) EXPECT() *MockTrappedApplicationServiceElementRequirements_Expecter {
	return &MockTrappedApplicationServiceElementRequirements_Expecter{mock: &_m.Mock}
}

// Confirmation provides a mock function with given fields: args, kwargs
func (_m *MockTrappedApplicationServiceElementRequirements) Confirmation(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Confirmation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bacnetip.Args, bacnetip.KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTrappedApplicationServiceElementRequirements_Confirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Confirmation'
type MockTrappedApplicationServiceElementRequirements_Confirmation_Call struct {
	*mock.Call
}

// Confirmation is a helper method to define mock.On call
//   - args bacnetip.Args
//   - kwargs bacnetip.KWArgs
func (_e *MockTrappedApplicationServiceElementRequirements_Expecter) Confirmation(args interface{}, kwargs interface{}) *MockTrappedApplicationServiceElementRequirements_Confirmation_Call {
	return &MockTrappedApplicationServiceElementRequirements_Confirmation_Call{Call: _e.mock.On("Confirmation", args, kwargs)}
}

func (_c *MockTrappedApplicationServiceElementRequirements_Confirmation_Call) Run(run func(args bacnetip.Args, kwargs bacnetip.KWArgs)) *MockTrappedApplicationServiceElementRequirements_Confirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.Args), args[1].(bacnetip.KWArgs))
	})
	return _c
}

func (_c *MockTrappedApplicationServiceElementRequirements_Confirmation_Call) Return(_a0 error) *MockTrappedApplicationServiceElementRequirements_Confirmation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrappedApplicationServiceElementRequirements_Confirmation_Call) RunAndReturn(run func(bacnetip.Args, bacnetip.KWArgs) error) *MockTrappedApplicationServiceElementRequirements_Confirmation_Call {
	_c.Call.Return(run)
	return _c
}

// Indication provides a mock function with given fields: args, kwargs
func (_m *MockTrappedApplicationServiceElementRequirements) Indication(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Indication")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bacnetip.Args, bacnetip.KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTrappedApplicationServiceElementRequirements_Indication_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Indication'
type MockTrappedApplicationServiceElementRequirements_Indication_Call struct {
	*mock.Call
}

// Indication is a helper method to define mock.On call
//   - args bacnetip.Args
//   - kwargs bacnetip.KWArgs
func (_e *MockTrappedApplicationServiceElementRequirements_Expecter) Indication(args interface{}, kwargs interface{}) *MockTrappedApplicationServiceElementRequirements_Indication_Call {
	return &MockTrappedApplicationServiceElementRequirements_Indication_Call{Call: _e.mock.On("Indication", args, kwargs)}
}

func (_c *MockTrappedApplicationServiceElementRequirements_Indication_Call) Run(run func(args bacnetip.Args, kwargs bacnetip.KWArgs)) *MockTrappedApplicationServiceElementRequirements_Indication_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.Args), args[1].(bacnetip.KWArgs))
	})
	return _c
}

func (_c *MockTrappedApplicationServiceElementRequirements_Indication_Call) Return(_a0 error) *MockTrappedApplicationServiceElementRequirements_Indication_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrappedApplicationServiceElementRequirements_Indication_Call) RunAndReturn(run func(bacnetip.Args, bacnetip.KWArgs) error) *MockTrappedApplicationServiceElementRequirements_Indication_Call {
	_c.Call.Return(run)
	return _c
}

// Request provides a mock function with given fields: args, kwargs
func (_m *MockTrappedApplicationServiceElementRequirements) Request(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Request")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bacnetip.Args, bacnetip.KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTrappedApplicationServiceElementRequirements_Request_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Request'
type MockTrappedApplicationServiceElementRequirements_Request_Call struct {
	*mock.Call
}

// Request is a helper method to define mock.On call
//   - args bacnetip.Args
//   - kwargs bacnetip.KWArgs
func (_e *MockTrappedApplicationServiceElementRequirements_Expecter) Request(args interface{}, kwargs interface{}) *MockTrappedApplicationServiceElementRequirements_Request_Call {
	return &MockTrappedApplicationServiceElementRequirements_Request_Call{Call: _e.mock.On("Request", args, kwargs)}
}

func (_c *MockTrappedApplicationServiceElementRequirements_Request_Call) Run(run func(args bacnetip.Args, kwargs bacnetip.KWArgs)) *MockTrappedApplicationServiceElementRequirements_Request_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.Args), args[1].(bacnetip.KWArgs))
	})
	return _c
}

func (_c *MockTrappedApplicationServiceElementRequirements_Request_Call) Return(_a0 error) *MockTrappedApplicationServiceElementRequirements_Request_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrappedApplicationServiceElementRequirements_Request_Call) RunAndReturn(run func(bacnetip.Args, bacnetip.KWArgs) error) *MockTrappedApplicationServiceElementRequirements_Request_Call {
	_c.Call.Return(run)
	return _c
}

// Response provides a mock function with given fields: args, kwargs
func (_m *MockTrappedApplicationServiceElementRequirements) Response(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Response")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bacnetip.Args, bacnetip.KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTrappedApplicationServiceElementRequirements_Response_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Response'
type MockTrappedApplicationServiceElementRequirements_Response_Call struct {
	*mock.Call
}

// Response is a helper method to define mock.On call
//   - args bacnetip.Args
//   - kwargs bacnetip.KWArgs
func (_e *MockTrappedApplicationServiceElementRequirements_Expecter) Response(args interface{}, kwargs interface{}) *MockTrappedApplicationServiceElementRequirements_Response_Call {
	return &MockTrappedApplicationServiceElementRequirements_Response_Call{Call: _e.mock.On("Response", args, kwargs)}
}

func (_c *MockTrappedApplicationServiceElementRequirements_Response_Call) Run(run func(args bacnetip.Args, kwargs bacnetip.KWArgs)) *MockTrappedApplicationServiceElementRequirements_Response_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.Args), args[1].(bacnetip.KWArgs))
	})
	return _c
}

func (_c *MockTrappedApplicationServiceElementRequirements_Response_Call) Return(_a0 error) *MockTrappedApplicationServiceElementRequirements_Response_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrappedApplicationServiceElementRequirements_Response_Call) RunAndReturn(run func(bacnetip.Args, bacnetip.KWArgs) error) *MockTrappedApplicationServiceElementRequirements_Response_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTrappedApplicationServiceElementRequirements creates a new instance of MockTrappedApplicationServiceElementRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTrappedApplicationServiceElementRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTrappedApplicationServiceElementRequirements {
	mock := &MockTrappedApplicationServiceElementRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
