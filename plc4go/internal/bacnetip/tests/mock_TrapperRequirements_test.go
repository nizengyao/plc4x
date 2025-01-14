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

// MockTrapperRequirements is an autogenerated mock type for the TrapperRequirements type
type MockTrapperRequirements struct {
	mock.Mock
}

type MockTrapperRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTrapperRequirements) EXPECT() *MockTrapperRequirements_Expecter {
	return &MockTrapperRequirements_Expecter{mock: &_m.Mock}
}

// AfterReceive provides a mock function with given fields: pdu
func (_m *MockTrapperRequirements) AfterReceive(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockTrapperRequirements_AfterReceive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AfterReceive'
type MockTrapperRequirements_AfterReceive_Call struct {
	*mock.Call
}

// AfterReceive is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockTrapperRequirements_Expecter) AfterReceive(pdu interface{}) *MockTrapperRequirements_AfterReceive_Call {
	return &MockTrapperRequirements_AfterReceive_Call{Call: _e.mock.On("AfterReceive", pdu)}
}

func (_c *MockTrapperRequirements_AfterReceive_Call) Run(run func(pdu bacnetip.PDU)) *MockTrapperRequirements_AfterReceive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockTrapperRequirements_AfterReceive_Call) Return() *MockTrapperRequirements_AfterReceive_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTrapperRequirements_AfterReceive_Call) RunAndReturn(run func(bacnetip.PDU)) *MockTrapperRequirements_AfterReceive_Call {
	_c.Call.Return(run)
	return _c
}

// AfterSend provides a mock function with given fields: pdu
func (_m *MockTrapperRequirements) AfterSend(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockTrapperRequirements_AfterSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AfterSend'
type MockTrapperRequirements_AfterSend_Call struct {
	*mock.Call
}

// AfterSend is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockTrapperRequirements_Expecter) AfterSend(pdu interface{}) *MockTrapperRequirements_AfterSend_Call {
	return &MockTrapperRequirements_AfterSend_Call{Call: _e.mock.On("AfterSend", pdu)}
}

func (_c *MockTrapperRequirements_AfterSend_Call) Run(run func(pdu bacnetip.PDU)) *MockTrapperRequirements_AfterSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockTrapperRequirements_AfterSend_Call) Return() *MockTrapperRequirements_AfterSend_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTrapperRequirements_AfterSend_Call) RunAndReturn(run func(bacnetip.PDU)) *MockTrapperRequirements_AfterSend_Call {
	_c.Call.Return(run)
	return _c
}

// BeforeReceive provides a mock function with given fields: pdu
func (_m *MockTrapperRequirements) BeforeReceive(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockTrapperRequirements_BeforeReceive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeforeReceive'
type MockTrapperRequirements_BeforeReceive_Call struct {
	*mock.Call
}

// BeforeReceive is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockTrapperRequirements_Expecter) BeforeReceive(pdu interface{}) *MockTrapperRequirements_BeforeReceive_Call {
	return &MockTrapperRequirements_BeforeReceive_Call{Call: _e.mock.On("BeforeReceive", pdu)}
}

func (_c *MockTrapperRequirements_BeforeReceive_Call) Run(run func(pdu bacnetip.PDU)) *MockTrapperRequirements_BeforeReceive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockTrapperRequirements_BeforeReceive_Call) Return() *MockTrapperRequirements_BeforeReceive_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTrapperRequirements_BeforeReceive_Call) RunAndReturn(run func(bacnetip.PDU)) *MockTrapperRequirements_BeforeReceive_Call {
	_c.Call.Return(run)
	return _c
}

// BeforeSend provides a mock function with given fields: pdu
func (_m *MockTrapperRequirements) BeforeSend(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockTrapperRequirements_BeforeSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeforeSend'
type MockTrapperRequirements_BeforeSend_Call struct {
	*mock.Call
}

// BeforeSend is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockTrapperRequirements_Expecter) BeforeSend(pdu interface{}) *MockTrapperRequirements_BeforeSend_Call {
	return &MockTrapperRequirements_BeforeSend_Call{Call: _e.mock.On("BeforeSend", pdu)}
}

func (_c *MockTrapperRequirements_BeforeSend_Call) Run(run func(pdu bacnetip.PDU)) *MockTrapperRequirements_BeforeSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockTrapperRequirements_BeforeSend_Call) Return() *MockTrapperRequirements_BeforeSend_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTrapperRequirements_BeforeSend_Call) RunAndReturn(run func(bacnetip.PDU)) *MockTrapperRequirements_BeforeSend_Call {
	_c.Call.Return(run)
	return _c
}

// UnexpectedReceive provides a mock function with given fields: pdu
func (_m *MockTrapperRequirements) UnexpectedReceive(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockTrapperRequirements_UnexpectedReceive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnexpectedReceive'
type MockTrapperRequirements_UnexpectedReceive_Call struct {
	*mock.Call
}

// UnexpectedReceive is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockTrapperRequirements_Expecter) UnexpectedReceive(pdu interface{}) *MockTrapperRequirements_UnexpectedReceive_Call {
	return &MockTrapperRequirements_UnexpectedReceive_Call{Call: _e.mock.On("UnexpectedReceive", pdu)}
}

func (_c *MockTrapperRequirements_UnexpectedReceive_Call) Run(run func(pdu bacnetip.PDU)) *MockTrapperRequirements_UnexpectedReceive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockTrapperRequirements_UnexpectedReceive_Call) Return() *MockTrapperRequirements_UnexpectedReceive_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTrapperRequirements_UnexpectedReceive_Call) RunAndReturn(run func(bacnetip.PDU)) *MockTrapperRequirements_UnexpectedReceive_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTrapperRequirements creates a new instance of MockTrapperRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTrapperRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTrapperRequirements {
	mock := &MockTrapperRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
