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

package model

import mock "github.com/stretchr/testify/mock"

// MockPlcBrowseRequestResult is an autogenerated mock type for the PlcBrowseRequestResult type
type MockPlcBrowseRequestResult struct {
	mock.Mock
}

type MockPlcBrowseRequestResult_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcBrowseRequestResult) EXPECT() *MockPlcBrowseRequestResult_Expecter {
	return &MockPlcBrowseRequestResult_Expecter{mock: &_m.Mock}
}

// GetErr provides a mock function with given fields:
func (_m *MockPlcBrowseRequestResult) GetErr() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcBrowseRequestResult_GetErr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetErr'
type MockPlcBrowseRequestResult_GetErr_Call struct {
	*mock.Call
}

// GetErr is a helper method to define mock.On call
func (_e *MockPlcBrowseRequestResult_Expecter) GetErr() *MockPlcBrowseRequestResult_GetErr_Call {
	return &MockPlcBrowseRequestResult_GetErr_Call{Call: _e.mock.On("GetErr")}
}

func (_c *MockPlcBrowseRequestResult_GetErr_Call) Run(run func()) *MockPlcBrowseRequestResult_GetErr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseRequestResult_GetErr_Call) Return(_a0 error) *MockPlcBrowseRequestResult_GetErr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseRequestResult_GetErr_Call) RunAndReturn(run func() error) *MockPlcBrowseRequestResult_GetErr_Call {
	_c.Call.Return(run)
	return _c
}

// GetRequest provides a mock function with given fields:
func (_m *MockPlcBrowseRequestResult) GetRequest() PlcBrowseRequest {
	ret := _m.Called()

	var r0 PlcBrowseRequest
	if rf, ok := ret.Get(0).(func() PlcBrowseRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcBrowseRequest)
		}
	}

	return r0
}

// MockPlcBrowseRequestResult_GetRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequest'
type MockPlcBrowseRequestResult_GetRequest_Call struct {
	*mock.Call
}

// GetRequest is a helper method to define mock.On call
func (_e *MockPlcBrowseRequestResult_Expecter) GetRequest() *MockPlcBrowseRequestResult_GetRequest_Call {
	return &MockPlcBrowseRequestResult_GetRequest_Call{Call: _e.mock.On("GetRequest")}
}

func (_c *MockPlcBrowseRequestResult_GetRequest_Call) Run(run func()) *MockPlcBrowseRequestResult_GetRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseRequestResult_GetRequest_Call) Return(_a0 PlcBrowseRequest) *MockPlcBrowseRequestResult_GetRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseRequestResult_GetRequest_Call) RunAndReturn(run func() PlcBrowseRequest) *MockPlcBrowseRequestResult_GetRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetResponse provides a mock function with given fields:
func (_m *MockPlcBrowseRequestResult) GetResponse() PlcBrowseResponse {
	ret := _m.Called()

	var r0 PlcBrowseResponse
	if rf, ok := ret.Get(0).(func() PlcBrowseResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcBrowseResponse)
		}
	}

	return r0
}

// MockPlcBrowseRequestResult_GetResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResponse'
type MockPlcBrowseRequestResult_GetResponse_Call struct {
	*mock.Call
}

// GetResponse is a helper method to define mock.On call
func (_e *MockPlcBrowseRequestResult_Expecter) GetResponse() *MockPlcBrowseRequestResult_GetResponse_Call {
	return &MockPlcBrowseRequestResult_GetResponse_Call{Call: _e.mock.On("GetResponse")}
}

func (_c *MockPlcBrowseRequestResult_GetResponse_Call) Run(run func()) *MockPlcBrowseRequestResult_GetResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseRequestResult_GetResponse_Call) Return(_a0 PlcBrowseResponse) *MockPlcBrowseRequestResult_GetResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseRequestResult_GetResponse_Call) RunAndReturn(run func() PlcBrowseResponse) *MockPlcBrowseRequestResult_GetResponse_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcBrowseRequestResult) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcBrowseRequestResult_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcBrowseRequestResult_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcBrowseRequestResult_Expecter) String() *MockPlcBrowseRequestResult_String_Call {
	return &MockPlcBrowseRequestResult_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcBrowseRequestResult_String_Call) Run(run func()) *MockPlcBrowseRequestResult_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseRequestResult_String_Call) Return(_a0 string) *MockPlcBrowseRequestResult_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseRequestResult_String_Call) RunAndReturn(run func() string) *MockPlcBrowseRequestResult_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcBrowseRequestResult creates a new instance of MockPlcBrowseRequestResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcBrowseRequestResult(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcBrowseRequestResult {
	mock := &MockPlcBrowseRequestResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
