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

// MockIOControllerRequirements is an autogenerated mock type for the IOControllerRequirements type
type MockIOControllerRequirements struct {
	mock.Mock
}

type MockIOControllerRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIOControllerRequirements) EXPECT() *MockIOControllerRequirements_Expecter {
	return &MockIOControllerRequirements_Expecter{mock: &_m.Mock}
}

// Abort provides a mock function with given fields: err
func (_m *MockIOControllerRequirements) Abort(err error) error {
	ret := _m.Called(err)

	if len(ret) == 0 {
		panic("no return value specified for Abort")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOControllerRequirements_Abort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Abort'
type MockIOControllerRequirements_Abort_Call struct {
	*mock.Call
}

// Abort is a helper method to define mock.On call
//   - err error
func (_e *MockIOControllerRequirements_Expecter) Abort(err interface{}) *MockIOControllerRequirements_Abort_Call {
	return &MockIOControllerRequirements_Abort_Call{Call: _e.mock.On("Abort", err)}
}

func (_c *MockIOControllerRequirements_Abort_Call) Run(run func(err error)) *MockIOControllerRequirements_Abort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *MockIOControllerRequirements_Abort_Call) Return(_a0 error) *MockIOControllerRequirements_Abort_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOControllerRequirements_Abort_Call) RunAndReturn(run func(error) error) *MockIOControllerRequirements_Abort_Call {
	_c.Call.Return(run)
	return _c
}

// AbortIO provides a mock function with given fields: iocb, err
func (_m *MockIOControllerRequirements) AbortIO(iocb _IOCB, err error) error {
	ret := _m.Called(iocb, err)

	if len(ret) == 0 {
		panic("no return value specified for AbortIO")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(_IOCB, error) error); ok {
		r0 = rf(iocb, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOControllerRequirements_AbortIO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AbortIO'
type MockIOControllerRequirements_AbortIO_Call struct {
	*mock.Call
}

// AbortIO is a helper method to define mock.On call
//   - iocb _IOCB
//   - err error
func (_e *MockIOControllerRequirements_Expecter) AbortIO(iocb interface{}, err interface{}) *MockIOControllerRequirements_AbortIO_Call {
	return &MockIOControllerRequirements_AbortIO_Call{Call: _e.mock.On("AbortIO", iocb, err)}
}

func (_c *MockIOControllerRequirements_AbortIO_Call) Run(run func(iocb _IOCB, err error)) *MockIOControllerRequirements_AbortIO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_IOCB), args[1].(error))
	})
	return _c
}

func (_c *MockIOControllerRequirements_AbortIO_Call) Return(_a0 error) *MockIOControllerRequirements_AbortIO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOControllerRequirements_AbortIO_Call) RunAndReturn(run func(_IOCB, error) error) *MockIOControllerRequirements_AbortIO_Call {
	_c.Call.Return(run)
	return _c
}

// CompleteIO provides a mock function with given fields: iocb, pdu
func (_m *MockIOControllerRequirements) CompleteIO(iocb _IOCB, pdu PDU) error {
	ret := _m.Called(iocb, pdu)

	if len(ret) == 0 {
		panic("no return value specified for CompleteIO")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(_IOCB, PDU) error); ok {
		r0 = rf(iocb, pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOControllerRequirements_CompleteIO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompleteIO'
type MockIOControllerRequirements_CompleteIO_Call struct {
	*mock.Call
}

// CompleteIO is a helper method to define mock.On call
//   - iocb _IOCB
//   - pdu PDU
func (_e *MockIOControllerRequirements_Expecter) CompleteIO(iocb interface{}, pdu interface{}) *MockIOControllerRequirements_CompleteIO_Call {
	return &MockIOControllerRequirements_CompleteIO_Call{Call: _e.mock.On("CompleteIO", iocb, pdu)}
}

func (_c *MockIOControllerRequirements_CompleteIO_Call) Run(run func(iocb _IOCB, pdu PDU)) *MockIOControllerRequirements_CompleteIO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_IOCB), args[1].(PDU))
	})
	return _c
}

func (_c *MockIOControllerRequirements_CompleteIO_Call) Return(_a0 error) *MockIOControllerRequirements_CompleteIO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOControllerRequirements_CompleteIO_Call) RunAndReturn(run func(_IOCB, PDU) error) *MockIOControllerRequirements_CompleteIO_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessIO provides a mock function with given fields: iocb
func (_m *MockIOControllerRequirements) ProcessIO(iocb _IOCB) error {
	ret := _m.Called(iocb)

	if len(ret) == 0 {
		panic("no return value specified for ProcessIO")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(_IOCB) error); ok {
		r0 = rf(iocb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOControllerRequirements_ProcessIO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessIO'
type MockIOControllerRequirements_ProcessIO_Call struct {
	*mock.Call
}

// ProcessIO is a helper method to define mock.On call
//   - iocb _IOCB
func (_e *MockIOControllerRequirements_Expecter) ProcessIO(iocb interface{}) *MockIOControllerRequirements_ProcessIO_Call {
	return &MockIOControllerRequirements_ProcessIO_Call{Call: _e.mock.On("ProcessIO", iocb)}
}

func (_c *MockIOControllerRequirements_ProcessIO_Call) Run(run func(iocb _IOCB)) *MockIOControllerRequirements_ProcessIO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_IOCB))
	})
	return _c
}

func (_c *MockIOControllerRequirements_ProcessIO_Call) Return(_a0 error) *MockIOControllerRequirements_ProcessIO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOControllerRequirements_ProcessIO_Call) RunAndReturn(run func(_IOCB) error) *MockIOControllerRequirements_ProcessIO_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIOControllerRequirements creates a new instance of MockIOControllerRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIOControllerRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIOControllerRequirements {
	mock := &MockIOControllerRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
