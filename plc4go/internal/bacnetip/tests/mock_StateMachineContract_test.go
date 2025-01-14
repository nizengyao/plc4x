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

// MockStateMachineContract is an autogenerated mock type for the StateMachineContract type
type MockStateMachineContract struct {
	mock.Mock
}

type MockStateMachineContract_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStateMachineContract) EXPECT() *MockStateMachineContract_Expecter {
	return &MockStateMachineContract_Expecter{mock: &_m.Mock}
}

// AfterReceive provides a mock function with given fields: pdu
func (_m *MockStateMachineContract) AfterReceive(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockStateMachineContract_AfterReceive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AfterReceive'
type MockStateMachineContract_AfterReceive_Call struct {
	*mock.Call
}

// AfterReceive is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockStateMachineContract_Expecter) AfterReceive(pdu interface{}) *MockStateMachineContract_AfterReceive_Call {
	return &MockStateMachineContract_AfterReceive_Call{Call: _e.mock.On("AfterReceive", pdu)}
}

func (_c *MockStateMachineContract_AfterReceive_Call) Run(run func(pdu bacnetip.PDU)) *MockStateMachineContract_AfterReceive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockStateMachineContract_AfterReceive_Call) Return() *MockStateMachineContract_AfterReceive_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineContract_AfterReceive_Call) RunAndReturn(run func(bacnetip.PDU)) *MockStateMachineContract_AfterReceive_Call {
	_c.Call.Return(run)
	return _c
}

// AfterSend provides a mock function with given fields: pdu
func (_m *MockStateMachineContract) AfterSend(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockStateMachineContract_AfterSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AfterSend'
type MockStateMachineContract_AfterSend_Call struct {
	*mock.Call
}

// AfterSend is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockStateMachineContract_Expecter) AfterSend(pdu interface{}) *MockStateMachineContract_AfterSend_Call {
	return &MockStateMachineContract_AfterSend_Call{Call: _e.mock.On("AfterSend", pdu)}
}

func (_c *MockStateMachineContract_AfterSend_Call) Run(run func(pdu bacnetip.PDU)) *MockStateMachineContract_AfterSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockStateMachineContract_AfterSend_Call) Return() *MockStateMachineContract_AfterSend_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineContract_AfterSend_Call) RunAndReturn(run func(bacnetip.PDU)) *MockStateMachineContract_AfterSend_Call {
	_c.Call.Return(run)
	return _c
}

// BeforeReceive provides a mock function with given fields: pdu
func (_m *MockStateMachineContract) BeforeReceive(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockStateMachineContract_BeforeReceive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeforeReceive'
type MockStateMachineContract_BeforeReceive_Call struct {
	*mock.Call
}

// BeforeReceive is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockStateMachineContract_Expecter) BeforeReceive(pdu interface{}) *MockStateMachineContract_BeforeReceive_Call {
	return &MockStateMachineContract_BeforeReceive_Call{Call: _e.mock.On("BeforeReceive", pdu)}
}

func (_c *MockStateMachineContract_BeforeReceive_Call) Run(run func(pdu bacnetip.PDU)) *MockStateMachineContract_BeforeReceive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockStateMachineContract_BeforeReceive_Call) Return() *MockStateMachineContract_BeforeReceive_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineContract_BeforeReceive_Call) RunAndReturn(run func(bacnetip.PDU)) *MockStateMachineContract_BeforeReceive_Call {
	_c.Call.Return(run)
	return _c
}

// BeforeSend provides a mock function with given fields: pdu
func (_m *MockStateMachineContract) BeforeSend(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockStateMachineContract_BeforeSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeforeSend'
type MockStateMachineContract_BeforeSend_Call struct {
	*mock.Call
}

// BeforeSend is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockStateMachineContract_Expecter) BeforeSend(pdu interface{}) *MockStateMachineContract_BeforeSend_Call {
	return &MockStateMachineContract_BeforeSend_Call{Call: _e.mock.On("BeforeSend", pdu)}
}

func (_c *MockStateMachineContract_BeforeSend_Call) Run(run func(pdu bacnetip.PDU)) *MockStateMachineContract_BeforeSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockStateMachineContract_BeforeSend_Call) Return() *MockStateMachineContract_BeforeSend_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineContract_BeforeSend_Call) RunAndReturn(run func(bacnetip.PDU)) *MockStateMachineContract_BeforeSend_Call {
	_c.Call.Return(run)
	return _c
}

// EventSet provides a mock function with given fields: id
func (_m *MockStateMachineContract) EventSet(id string) {
	_m.Called(id)
}

// MockStateMachineContract_EventSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EventSet'
type MockStateMachineContract_EventSet_Call struct {
	*mock.Call
}

// EventSet is a helper method to define mock.On call
//   - id string
func (_e *MockStateMachineContract_Expecter) EventSet(id interface{}) *MockStateMachineContract_EventSet_Call {
	return &MockStateMachineContract_EventSet_Call{Call: _e.mock.On("EventSet", id)}
}

func (_c *MockStateMachineContract_EventSet_Call) Run(run func(id string)) *MockStateMachineContract_EventSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStateMachineContract_EventSet_Call) Return() *MockStateMachineContract_EventSet_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineContract_EventSet_Call) RunAndReturn(run func(string)) *MockStateMachineContract_EventSet_Call {
	_c.Call.Return(run)
	return _c
}

// GetCurrentState provides a mock function with given fields:
func (_m *MockStateMachineContract) GetCurrentState() State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentState")
	}

	var r0 State
	if rf, ok := ret.Get(0).(func() State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(State)
		}
	}

	return r0
}

// MockStateMachineContract_GetCurrentState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrentState'
type MockStateMachineContract_GetCurrentState_Call struct {
	*mock.Call
}

// GetCurrentState is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) GetCurrentState() *MockStateMachineContract_GetCurrentState_Call {
	return &MockStateMachineContract_GetCurrentState_Call{Call: _e.mock.On("GetCurrentState")}
}

func (_c *MockStateMachineContract_GetCurrentState_Call) Run(run func()) *MockStateMachineContract_GetCurrentState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_GetCurrentState_Call) Return(_a0 State) *MockStateMachineContract_GetCurrentState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_GetCurrentState_Call) RunAndReturn(run func() State) *MockStateMachineContract_GetCurrentState_Call {
	_c.Call.Return(run)
	return _c
}

// GetStartState provides a mock function with given fields:
func (_m *MockStateMachineContract) GetStartState() State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStartState")
	}

	var r0 State
	if rf, ok := ret.Get(0).(func() State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(State)
		}
	}

	return r0
}

// MockStateMachineContract_GetStartState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStartState'
type MockStateMachineContract_GetStartState_Call struct {
	*mock.Call
}

// GetStartState is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) GetStartState() *MockStateMachineContract_GetStartState_Call {
	return &MockStateMachineContract_GetStartState_Call{Call: _e.mock.On("GetStartState")}
}

func (_c *MockStateMachineContract_GetStartState_Call) Run(run func()) *MockStateMachineContract_GetStartState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_GetStartState_Call) Return(_a0 State) *MockStateMachineContract_GetStartState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_GetStartState_Call) RunAndReturn(run func() State) *MockStateMachineContract_GetStartState_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransactionLog provides a mock function with given fields:
func (_m *MockStateMachineContract) GetTransactionLog() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionLog")
	}

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

// MockStateMachineContract_GetTransactionLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransactionLog'
type MockStateMachineContract_GetTransactionLog_Call struct {
	*mock.Call
}

// GetTransactionLog is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) GetTransactionLog() *MockStateMachineContract_GetTransactionLog_Call {
	return &MockStateMachineContract_GetTransactionLog_Call{Call: _e.mock.On("GetTransactionLog")}
}

func (_c *MockStateMachineContract_GetTransactionLog_Call) Run(run func()) *MockStateMachineContract_GetTransactionLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_GetTransactionLog_Call) Return(_a0 []string) *MockStateMachineContract_GetTransactionLog_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_GetTransactionLog_Call) RunAndReturn(run func() []string) *MockStateMachineContract_GetTransactionLog_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnexpectedReceiveState provides a mock function with given fields:
func (_m *MockStateMachineContract) GetUnexpectedReceiveState() State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetUnexpectedReceiveState")
	}

	var r0 State
	if rf, ok := ret.Get(0).(func() State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(State)
		}
	}

	return r0
}

// MockStateMachineContract_GetUnexpectedReceiveState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnexpectedReceiveState'
type MockStateMachineContract_GetUnexpectedReceiveState_Call struct {
	*mock.Call
}

// GetUnexpectedReceiveState is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) GetUnexpectedReceiveState() *MockStateMachineContract_GetUnexpectedReceiveState_Call {
	return &MockStateMachineContract_GetUnexpectedReceiveState_Call{Call: _e.mock.On("GetUnexpectedReceiveState")}
}

func (_c *MockStateMachineContract_GetUnexpectedReceiveState_Call) Run(run func()) *MockStateMachineContract_GetUnexpectedReceiveState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_GetUnexpectedReceiveState_Call) Return(_a0 State) *MockStateMachineContract_GetUnexpectedReceiveState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_GetUnexpectedReceiveState_Call) RunAndReturn(run func() State) *MockStateMachineContract_GetUnexpectedReceiveState_Call {
	_c.Call.Return(run)
	return _c
}

// IsFailState provides a mock function with given fields:
func (_m *MockStateMachineContract) IsFailState() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsFailState")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockStateMachineContract_IsFailState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsFailState'
type MockStateMachineContract_IsFailState_Call struct {
	*mock.Call
}

// IsFailState is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) IsFailState() *MockStateMachineContract_IsFailState_Call {
	return &MockStateMachineContract_IsFailState_Call{Call: _e.mock.On("IsFailState")}
}

func (_c *MockStateMachineContract_IsFailState_Call) Run(run func()) *MockStateMachineContract_IsFailState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_IsFailState_Call) Return(_a0 bool) *MockStateMachineContract_IsFailState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_IsFailState_Call) RunAndReturn(run func() bool) *MockStateMachineContract_IsFailState_Call {
	_c.Call.Return(run)
	return _c
}

// IsRunning provides a mock function with given fields:
func (_m *MockStateMachineContract) IsRunning() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsRunning")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockStateMachineContract_IsRunning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRunning'
type MockStateMachineContract_IsRunning_Call struct {
	*mock.Call
}

// IsRunning is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) IsRunning() *MockStateMachineContract_IsRunning_Call {
	return &MockStateMachineContract_IsRunning_Call{Call: _e.mock.On("IsRunning")}
}

func (_c *MockStateMachineContract_IsRunning_Call) Run(run func()) *MockStateMachineContract_IsRunning_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_IsRunning_Call) Return(_a0 bool) *MockStateMachineContract_IsRunning_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_IsRunning_Call) RunAndReturn(run func() bool) *MockStateMachineContract_IsRunning_Call {
	_c.Call.Return(run)
	return _c
}

// IsSuccessState provides a mock function with given fields:
func (_m *MockStateMachineContract) IsSuccessState() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSuccessState")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockStateMachineContract_IsSuccessState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSuccessState'
type MockStateMachineContract_IsSuccessState_Call struct {
	*mock.Call
}

// IsSuccessState is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) IsSuccessState() *MockStateMachineContract_IsSuccessState_Call {
	return &MockStateMachineContract_IsSuccessState_Call{Call: _e.mock.On("IsSuccessState")}
}

func (_c *MockStateMachineContract_IsSuccessState_Call) Run(run func()) *MockStateMachineContract_IsSuccessState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_IsSuccessState_Call) Return(_a0 bool) *MockStateMachineContract_IsSuccessState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_IsSuccessState_Call) RunAndReturn(run func() bool) *MockStateMachineContract_IsSuccessState_Call {
	_c.Call.Return(run)
	return _c
}

// NewState provides a mock function with given fields: _a0
func (_m *MockStateMachineContract) NewState(_a0 string) State {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for NewState")
	}

	var r0 State
	if rf, ok := ret.Get(0).(func(string) State); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(State)
		}
	}

	return r0
}

// MockStateMachineContract_NewState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewState'
type MockStateMachineContract_NewState_Call struct {
	*mock.Call
}

// NewState is a helper method to define mock.On call
//   - _a0 string
func (_e *MockStateMachineContract_Expecter) NewState(_a0 interface{}) *MockStateMachineContract_NewState_Call {
	return &MockStateMachineContract_NewState_Call{Call: _e.mock.On("NewState", _a0)}
}

func (_c *MockStateMachineContract_NewState_Call) Run(run func(_a0 string)) *MockStateMachineContract_NewState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStateMachineContract_NewState_Call) Return(_a0 State) *MockStateMachineContract_NewState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_NewState_Call) RunAndReturn(run func(string) State) *MockStateMachineContract_NewState_Call {
	_c.Call.Return(run)
	return _c
}

// Receive provides a mock function with given fields: args, kwargs
func (_m *MockStateMachineContract) Receive(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Receive")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bacnetip.Args, bacnetip.KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStateMachineContract_Receive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Receive'
type MockStateMachineContract_Receive_Call struct {
	*mock.Call
}

// Receive is a helper method to define mock.On call
//   - args bacnetip.Args
//   - kwargs bacnetip.KWArgs
func (_e *MockStateMachineContract_Expecter) Receive(args interface{}, kwargs interface{}) *MockStateMachineContract_Receive_Call {
	return &MockStateMachineContract_Receive_Call{Call: _e.mock.On("Receive", args, kwargs)}
}

func (_c *MockStateMachineContract_Receive_Call) Run(run func(args bacnetip.Args, kwargs bacnetip.KWArgs)) *MockStateMachineContract_Receive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.Args), args[1].(bacnetip.KWArgs))
	})
	return _c
}

func (_c *MockStateMachineContract_Receive_Call) Return(_a0 error) *MockStateMachineContract_Receive_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_Receive_Call) RunAndReturn(run func(bacnetip.Args, bacnetip.KWArgs) error) *MockStateMachineContract_Receive_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields:
func (_m *MockStateMachineContract) Reset() {
	_m.Called()
}

// MockStateMachineContract_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type MockStateMachineContract_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) Reset() *MockStateMachineContract_Reset_Call {
	return &MockStateMachineContract_Reset_Call{Call: _e.mock.On("Reset")}
}

func (_c *MockStateMachineContract_Reset_Call) Run(run func()) *MockStateMachineContract_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_Reset_Call) Return() *MockStateMachineContract_Reset_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineContract_Reset_Call) RunAndReturn(run func()) *MockStateMachineContract_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields:
func (_m *MockStateMachineContract) Run() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStateMachineContract_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockStateMachineContract_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) Run() *MockStateMachineContract_Run_Call {
	return &MockStateMachineContract_Run_Call{Call: _e.mock.On("Run")}
}

func (_c *MockStateMachineContract_Run_Call) Run(run func()) *MockStateMachineContract_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_Run_Call) Return(_a0 error) *MockStateMachineContract_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_Run_Call) RunAndReturn(run func() error) *MockStateMachineContract_Run_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockStateMachineContract) String() string {
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

// MockStateMachineContract_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockStateMachineContract_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) String() *MockStateMachineContract_String_Call {
	return &MockStateMachineContract_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockStateMachineContract_String_Call) Run(run func()) *MockStateMachineContract_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_String_Call) Return(_a0 string) *MockStateMachineContract_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_String_Call) RunAndReturn(run func() string) *MockStateMachineContract_String_Call {
	_c.Call.Return(run)
	return _c
}

// UnexpectedReceive provides a mock function with given fields: pdu
func (_m *MockStateMachineContract) UnexpectedReceive(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockStateMachineContract_UnexpectedReceive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnexpectedReceive'
type MockStateMachineContract_UnexpectedReceive_Call struct {
	*mock.Call
}

// UnexpectedReceive is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockStateMachineContract_Expecter) UnexpectedReceive(pdu interface{}) *MockStateMachineContract_UnexpectedReceive_Call {
	return &MockStateMachineContract_UnexpectedReceive_Call{Call: _e.mock.On("UnexpectedReceive", pdu)}
}

func (_c *MockStateMachineContract_UnexpectedReceive_Call) Run(run func(pdu bacnetip.PDU)) *MockStateMachineContract_UnexpectedReceive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockStateMachineContract_UnexpectedReceive_Call) Return() *MockStateMachineContract_UnexpectedReceive_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineContract_UnexpectedReceive_Call) RunAndReturn(run func(bacnetip.PDU)) *MockStateMachineContract_UnexpectedReceive_Call {
	_c.Call.Return(run)
	return _c
}

// getCurrentState provides a mock function with given fields:
func (_m *MockStateMachineContract) getCurrentState() State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getCurrentState")
	}

	var r0 State
	if rf, ok := ret.Get(0).(func() State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(State)
		}
	}

	return r0
}

// MockStateMachineContract_getCurrentState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getCurrentState'
type MockStateMachineContract_getCurrentState_Call struct {
	*mock.Call
}

// getCurrentState is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) getCurrentState() *MockStateMachineContract_getCurrentState_Call {
	return &MockStateMachineContract_getCurrentState_Call{Call: _e.mock.On("getCurrentState")}
}

func (_c *MockStateMachineContract_getCurrentState_Call) Run(run func()) *MockStateMachineContract_getCurrentState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_getCurrentState_Call) Return(_a0 State) *MockStateMachineContract_getCurrentState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_getCurrentState_Call) RunAndReturn(run func() State) *MockStateMachineContract_getCurrentState_Call {
	_c.Call.Return(run)
	return _c
}

// getMachineGroup provides a mock function with given fields:
func (_m *MockStateMachineContract) getMachineGroup() *StateMachineGroup {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getMachineGroup")
	}

	var r0 *StateMachineGroup
	if rf, ok := ret.Get(0).(func() *StateMachineGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StateMachineGroup)
		}
	}

	return r0
}

// MockStateMachineContract_getMachineGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getMachineGroup'
type MockStateMachineContract_getMachineGroup_Call struct {
	*mock.Call
}

// getMachineGroup is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) getMachineGroup() *MockStateMachineContract_getMachineGroup_Call {
	return &MockStateMachineContract_getMachineGroup_Call{Call: _e.mock.On("getMachineGroup")}
}

func (_c *MockStateMachineContract_getMachineGroup_Call) Run(run func()) *MockStateMachineContract_getMachineGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_getMachineGroup_Call) Return(_a0 *StateMachineGroup) *MockStateMachineContract_getMachineGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_getMachineGroup_Call) RunAndReturn(run func() *StateMachineGroup) *MockStateMachineContract_getMachineGroup_Call {
	_c.Call.Return(run)
	return _c
}

// getStateTimeoutTask provides a mock function with given fields:
func (_m *MockStateMachineContract) getStateTimeoutTask() *TimeoutTask {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getStateTimeoutTask")
	}

	var r0 *TimeoutTask
	if rf, ok := ret.Get(0).(func() *TimeoutTask); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TimeoutTask)
		}
	}

	return r0
}

// MockStateMachineContract_getStateTimeoutTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getStateTimeoutTask'
type MockStateMachineContract_getStateTimeoutTask_Call struct {
	*mock.Call
}

// getStateTimeoutTask is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) getStateTimeoutTask() *MockStateMachineContract_getStateTimeoutTask_Call {
	return &MockStateMachineContract_getStateTimeoutTask_Call{Call: _e.mock.On("getStateTimeoutTask")}
}

func (_c *MockStateMachineContract_getStateTimeoutTask_Call) Run(run func()) *MockStateMachineContract_getStateTimeoutTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_getStateTimeoutTask_Call) Return(_a0 *TimeoutTask) *MockStateMachineContract_getStateTimeoutTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_getStateTimeoutTask_Call) RunAndReturn(run func() *TimeoutTask) *MockStateMachineContract_getStateTimeoutTask_Call {
	_c.Call.Return(run)
	return _c
}

// getStates provides a mock function with given fields:
func (_m *MockStateMachineContract) getStates() []State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getStates")
	}

	var r0 []State
	if rf, ok := ret.Get(0).(func() []State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]State)
		}
	}

	return r0
}

// MockStateMachineContract_getStates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getStates'
type MockStateMachineContract_getStates_Call struct {
	*mock.Call
}

// getStates is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) getStates() *MockStateMachineContract_getStates_Call {
	return &MockStateMachineContract_getStates_Call{Call: _e.mock.On("getStates")}
}

func (_c *MockStateMachineContract_getStates_Call) Run(run func()) *MockStateMachineContract_getStates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_getStates_Call) Return(_a0 []State) *MockStateMachineContract_getStates_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineContract_getStates_Call) RunAndReturn(run func() []State) *MockStateMachineContract_getStates_Call {
	_c.Call.Return(run)
	return _c
}

// halt provides a mock function with given fields:
func (_m *MockStateMachineContract) halt() {
	_m.Called()
}

// MockStateMachineContract_halt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'halt'
type MockStateMachineContract_halt_Call struct {
	*mock.Call
}

// halt is a helper method to define mock.On call
func (_e *MockStateMachineContract_Expecter) halt() *MockStateMachineContract_halt_Call {
	return &MockStateMachineContract_halt_Call{Call: _e.mock.On("halt")}
}

func (_c *MockStateMachineContract_halt_Call) Run(run func()) *MockStateMachineContract_halt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineContract_halt_Call) Return() *MockStateMachineContract_halt_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineContract_halt_Call) RunAndReturn(run func()) *MockStateMachineContract_halt_Call {
	_c.Call.Return(run)
	return _c
}

// setMachineGroup provides a mock function with given fields: machineGroup
func (_m *MockStateMachineContract) setMachineGroup(machineGroup *StateMachineGroup) {
	_m.Called(machineGroup)
}

// MockStateMachineContract_setMachineGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setMachineGroup'
type MockStateMachineContract_setMachineGroup_Call struct {
	*mock.Call
}

// setMachineGroup is a helper method to define mock.On call
//   - machineGroup *StateMachineGroup
func (_e *MockStateMachineContract_Expecter) setMachineGroup(machineGroup interface{}) *MockStateMachineContract_setMachineGroup_Call {
	return &MockStateMachineContract_setMachineGroup_Call{Call: _e.mock.On("setMachineGroup", machineGroup)}
}

func (_c *MockStateMachineContract_setMachineGroup_Call) Run(run func(machineGroup *StateMachineGroup)) *MockStateMachineContract_setMachineGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*StateMachineGroup))
	})
	return _c
}

func (_c *MockStateMachineContract_setMachineGroup_Call) Return() *MockStateMachineContract_setMachineGroup_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineContract_setMachineGroup_Call) RunAndReturn(run func(*StateMachineGroup)) *MockStateMachineContract_setMachineGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStateMachineContract creates a new instance of MockStateMachineContract. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStateMachineContract(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStateMachineContract {
	mock := &MockStateMachineContract{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
