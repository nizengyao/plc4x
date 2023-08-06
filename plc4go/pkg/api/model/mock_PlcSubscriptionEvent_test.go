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

import (
	values "github.com/apache/plc4x/plc4go/pkg/api/values"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcSubscriptionEvent is an autogenerated mock type for the PlcSubscriptionEvent type
type MockPlcSubscriptionEvent struct {
	mock.Mock
}

type MockPlcSubscriptionEvent_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcSubscriptionEvent) EXPECT() *MockPlcSubscriptionEvent_Expecter {
	return &MockPlcSubscriptionEvent_Expecter{mock: &_m.Mock}
}

// GetAddress provides a mock function with given fields: tagName
func (_m *MockPlcSubscriptionEvent) GetAddress(tagName string) string {
	ret := _m.Called(tagName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(tagName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcSubscriptionEvent_GetAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddress'
type MockPlcSubscriptionEvent_GetAddress_Call struct {
	*mock.Call
}

// GetAddress is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcSubscriptionEvent_Expecter) GetAddress(tagName interface{}) *MockPlcSubscriptionEvent_GetAddress_Call {
	return &MockPlcSubscriptionEvent_GetAddress_Call{Call: _e.mock.On("GetAddress", tagName)}
}

func (_c *MockPlcSubscriptionEvent_GetAddress_Call) Run(run func(tagName string)) *MockPlcSubscriptionEvent_GetAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcSubscriptionEvent_GetAddress_Call) Return(_a0 string) *MockPlcSubscriptionEvent_GetAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionEvent_GetAddress_Call) RunAndReturn(run func(string) string) *MockPlcSubscriptionEvent_GetAddress_Call {
	_c.Call.Return(run)
	return _c
}

// GetResponseCode provides a mock function with given fields: tagName
func (_m *MockPlcSubscriptionEvent) GetResponseCode(tagName string) PlcResponseCode {
	ret := _m.Called(tagName)

	var r0 PlcResponseCode
	if rf, ok := ret.Get(0).(func(string) PlcResponseCode); ok {
		r0 = rf(tagName)
	} else {
		r0 = ret.Get(0).(PlcResponseCode)
	}

	return r0
}

// MockPlcSubscriptionEvent_GetResponseCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResponseCode'
type MockPlcSubscriptionEvent_GetResponseCode_Call struct {
	*mock.Call
}

// GetResponseCode is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcSubscriptionEvent_Expecter) GetResponseCode(tagName interface{}) *MockPlcSubscriptionEvent_GetResponseCode_Call {
	return &MockPlcSubscriptionEvent_GetResponseCode_Call{Call: _e.mock.On("GetResponseCode", tagName)}
}

func (_c *MockPlcSubscriptionEvent_GetResponseCode_Call) Run(run func(tagName string)) *MockPlcSubscriptionEvent_GetResponseCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcSubscriptionEvent_GetResponseCode_Call) Return(_a0 PlcResponseCode) *MockPlcSubscriptionEvent_GetResponseCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionEvent_GetResponseCode_Call) RunAndReturn(run func(string) PlcResponseCode) *MockPlcSubscriptionEvent_GetResponseCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetSource provides a mock function with given fields: tagName
func (_m *MockPlcSubscriptionEvent) GetSource(tagName string) string {
	ret := _m.Called(tagName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(tagName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcSubscriptionEvent_GetSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSource'
type MockPlcSubscriptionEvent_GetSource_Call struct {
	*mock.Call
}

// GetSource is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcSubscriptionEvent_Expecter) GetSource(tagName interface{}) *MockPlcSubscriptionEvent_GetSource_Call {
	return &MockPlcSubscriptionEvent_GetSource_Call{Call: _e.mock.On("GetSource", tagName)}
}

func (_c *MockPlcSubscriptionEvent_GetSource_Call) Run(run func(tagName string)) *MockPlcSubscriptionEvent_GetSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcSubscriptionEvent_GetSource_Call) Return(_a0 string) *MockPlcSubscriptionEvent_GetSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionEvent_GetSource_Call) RunAndReturn(run func(string) string) *MockPlcSubscriptionEvent_GetSource_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagNames provides a mock function with given fields:
func (_m *MockPlcSubscriptionEvent) GetTagNames() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockPlcSubscriptionEvent_GetTagNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagNames'
type MockPlcSubscriptionEvent_GetTagNames_Call struct {
	*mock.Call
}

// GetTagNames is a helper method to define mock.On call
func (_e *MockPlcSubscriptionEvent_Expecter) GetTagNames() *MockPlcSubscriptionEvent_GetTagNames_Call {
	return &MockPlcSubscriptionEvent_GetTagNames_Call{Call: _e.mock.On("GetTagNames")}
}

func (_c *MockPlcSubscriptionEvent_GetTagNames_Call) Run(run func()) *MockPlcSubscriptionEvent_GetTagNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionEvent_GetTagNames_Call) Return(_a0 []string) *MockPlcSubscriptionEvent_GetTagNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionEvent_GetTagNames_Call) RunAndReturn(run func() []string) *MockPlcSubscriptionEvent_GetTagNames_Call {
	_c.Call.Return(run)
	return _c
}

// GetValue provides a mock function with given fields: tagName
func (_m *MockPlcSubscriptionEvent) GetValue(tagName string) values.PlcValue {
	ret := _m.Called(tagName)

	var r0 values.PlcValue
	if rf, ok := ret.Get(0).(func(string) values.PlcValue); ok {
		r0 = rf(tagName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(values.PlcValue)
		}
	}

	return r0
}

// MockPlcSubscriptionEvent_GetValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValue'
type MockPlcSubscriptionEvent_GetValue_Call struct {
	*mock.Call
}

// GetValue is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcSubscriptionEvent_Expecter) GetValue(tagName interface{}) *MockPlcSubscriptionEvent_GetValue_Call {
	return &MockPlcSubscriptionEvent_GetValue_Call{Call: _e.mock.On("GetValue", tagName)}
}

func (_c *MockPlcSubscriptionEvent_GetValue_Call) Run(run func(tagName string)) *MockPlcSubscriptionEvent_GetValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcSubscriptionEvent_GetValue_Call) Return(_a0 values.PlcValue) *MockPlcSubscriptionEvent_GetValue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionEvent_GetValue_Call) RunAndReturn(run func(string) values.PlcValue) *MockPlcSubscriptionEvent_GetValue_Call {
	_c.Call.Return(run)
	return _c
}

// IsAPlcMessage provides a mock function with given fields:
func (_m *MockPlcSubscriptionEvent) IsAPlcMessage() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcSubscriptionEvent_IsAPlcMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAPlcMessage'
type MockPlcSubscriptionEvent_IsAPlcMessage_Call struct {
	*mock.Call
}

// IsAPlcMessage is a helper method to define mock.On call
func (_e *MockPlcSubscriptionEvent_Expecter) IsAPlcMessage() *MockPlcSubscriptionEvent_IsAPlcMessage_Call {
	return &MockPlcSubscriptionEvent_IsAPlcMessage_Call{Call: _e.mock.On("IsAPlcMessage")}
}

func (_c *MockPlcSubscriptionEvent_IsAPlcMessage_Call) Run(run func()) *MockPlcSubscriptionEvent_IsAPlcMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionEvent_IsAPlcMessage_Call) Return(_a0 bool) *MockPlcSubscriptionEvent_IsAPlcMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionEvent_IsAPlcMessage_Call) RunAndReturn(run func() bool) *MockPlcSubscriptionEvent_IsAPlcMessage_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcSubscriptionEvent) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcSubscriptionEvent_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcSubscriptionEvent_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcSubscriptionEvent_Expecter) String() *MockPlcSubscriptionEvent_String_Call {
	return &MockPlcSubscriptionEvent_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcSubscriptionEvent_String_Call) Run(run func()) *MockPlcSubscriptionEvent_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionEvent_String_Call) Return(_a0 string) *MockPlcSubscriptionEvent_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionEvent_String_Call) RunAndReturn(run func() string) *MockPlcSubscriptionEvent_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcSubscriptionEvent creates a new instance of MockPlcSubscriptionEvent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcSubscriptionEvent(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcSubscriptionEvent {
	mock := &MockPlcSubscriptionEvent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
