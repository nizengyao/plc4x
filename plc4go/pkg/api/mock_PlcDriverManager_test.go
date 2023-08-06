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

package plc4go

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcDriverManager is an autogenerated mock type for the PlcDriverManager type
type MockPlcDriverManager struct {
	mock.Mock
}

type MockPlcDriverManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcDriverManager) EXPECT() *MockPlcDriverManager_Expecter {
	return &MockPlcDriverManager_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockPlcDriverManager) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcDriverManager_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockPlcDriverManager_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockPlcDriverManager_Expecter) Close() *MockPlcDriverManager_Close_Call {
	return &MockPlcDriverManager_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockPlcDriverManager_Close_Call) Run(run func()) *MockPlcDriverManager_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcDriverManager_Close_Call) Return(_a0 error) *MockPlcDriverManager_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriverManager_Close_Call) RunAndReturn(run func() error) *MockPlcDriverManager_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Discover provides a mock function with given fields: callback, discoveryOptions
func (_m *MockPlcDriverManager) Discover(callback func(model.PlcDiscoveryItem), discoveryOptions ...WithDiscoveryOption) error {
	_va := make([]interface{}, len(discoveryOptions))
	for _i := range discoveryOptions {
		_va[_i] = discoveryOptions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, callback)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(model.PlcDiscoveryItem), ...WithDiscoveryOption) error); ok {
		r0 = rf(callback, discoveryOptions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcDriverManager_Discover_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Discover'
type MockPlcDriverManager_Discover_Call struct {
	*mock.Call
}

// Discover is a helper method to define mock.On call
//   - callback func(model.PlcDiscoveryItem)
//   - discoveryOptions ...WithDiscoveryOption
func (_e *MockPlcDriverManager_Expecter) Discover(callback interface{}, discoveryOptions ...interface{}) *MockPlcDriverManager_Discover_Call {
	return &MockPlcDriverManager_Discover_Call{Call: _e.mock.On("Discover",
		append([]interface{}{callback}, discoveryOptions...)...)}
}

func (_c *MockPlcDriverManager_Discover_Call) Run(run func(callback func(model.PlcDiscoveryItem), discoveryOptions ...WithDiscoveryOption)) *MockPlcDriverManager_Discover_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]WithDiscoveryOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(WithDiscoveryOption)
			}
		}
		run(args[0].(func(model.PlcDiscoveryItem)), variadicArgs...)
	})
	return _c
}

func (_c *MockPlcDriverManager_Discover_Call) Return(_a0 error) *MockPlcDriverManager_Discover_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriverManager_Discover_Call) RunAndReturn(run func(func(model.PlcDiscoveryItem), ...WithDiscoveryOption) error) *MockPlcDriverManager_Discover_Call {
	_c.Call.Return(run)
	return _c
}

// DiscoverWithContext provides a mock function with given fields: ctx, callback, discoveryOptions
func (_m *MockPlcDriverManager) DiscoverWithContext(ctx context.Context, callback func(model.PlcDiscoveryItem), discoveryOptions ...WithDiscoveryOption) error {
	_va := make([]interface{}, len(discoveryOptions))
	for _i := range discoveryOptions {
		_va[_i] = discoveryOptions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, callback)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(model.PlcDiscoveryItem), ...WithDiscoveryOption) error); ok {
		r0 = rf(ctx, callback, discoveryOptions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcDriverManager_DiscoverWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DiscoverWithContext'
type MockPlcDriverManager_DiscoverWithContext_Call struct {
	*mock.Call
}

// DiscoverWithContext is a helper method to define mock.On call
//   - ctx context.Context
//   - callback func(model.PlcDiscoveryItem)
//   - discoveryOptions ...WithDiscoveryOption
func (_e *MockPlcDriverManager_Expecter) DiscoverWithContext(ctx interface{}, callback interface{}, discoveryOptions ...interface{}) *MockPlcDriverManager_DiscoverWithContext_Call {
	return &MockPlcDriverManager_DiscoverWithContext_Call{Call: _e.mock.On("DiscoverWithContext",
		append([]interface{}{ctx, callback}, discoveryOptions...)...)}
}

func (_c *MockPlcDriverManager_DiscoverWithContext_Call) Run(run func(ctx context.Context, callback func(model.PlcDiscoveryItem), discoveryOptions ...WithDiscoveryOption)) *MockPlcDriverManager_DiscoverWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]WithDiscoveryOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(WithDiscoveryOption)
			}
		}
		run(args[0].(context.Context), args[1].(func(model.PlcDiscoveryItem)), variadicArgs...)
	})
	return _c
}

func (_c *MockPlcDriverManager_DiscoverWithContext_Call) Return(_a0 error) *MockPlcDriverManager_DiscoverWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriverManager_DiscoverWithContext_Call) RunAndReturn(run func(context.Context, func(model.PlcDiscoveryItem), ...WithDiscoveryOption) error) *MockPlcDriverManager_DiscoverWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnection provides a mock function with given fields: connectionString
func (_m *MockPlcDriverManager) GetConnection(connectionString string) <-chan PlcConnectionConnectResult {
	ret := _m.Called(connectionString)

	var r0 <-chan PlcConnectionConnectResult
	if rf, ok := ret.Get(0).(func(string) <-chan PlcConnectionConnectResult); ok {
		r0 = rf(connectionString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcConnectionConnectResult)
		}
	}

	return r0
}

// MockPlcDriverManager_GetConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnection'
type MockPlcDriverManager_GetConnection_Call struct {
	*mock.Call
}

// GetConnection is a helper method to define mock.On call
//   - connectionString string
func (_e *MockPlcDriverManager_Expecter) GetConnection(connectionString interface{}) *MockPlcDriverManager_GetConnection_Call {
	return &MockPlcDriverManager_GetConnection_Call{Call: _e.mock.On("GetConnection", connectionString)}
}

func (_c *MockPlcDriverManager_GetConnection_Call) Run(run func(connectionString string)) *MockPlcDriverManager_GetConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcDriverManager_GetConnection_Call) Return(_a0 <-chan PlcConnectionConnectResult) *MockPlcDriverManager_GetConnection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriverManager_GetConnection_Call) RunAndReturn(run func(string) <-chan PlcConnectionConnectResult) *MockPlcDriverManager_GetConnection_Call {
	_c.Call.Return(run)
	return _c
}

// GetDriver provides a mock function with given fields: driverName
func (_m *MockPlcDriverManager) GetDriver(driverName string) (PlcDriver, error) {
	ret := _m.Called(driverName)

	var r0 PlcDriver
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (PlcDriver, error)); ok {
		return rf(driverName)
	}
	if rf, ok := ret.Get(0).(func(string) PlcDriver); ok {
		r0 = rf(driverName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcDriver)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(driverName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlcDriverManager_GetDriver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDriver'
type MockPlcDriverManager_GetDriver_Call struct {
	*mock.Call
}

// GetDriver is a helper method to define mock.On call
//   - driverName string
func (_e *MockPlcDriverManager_Expecter) GetDriver(driverName interface{}) *MockPlcDriverManager_GetDriver_Call {
	return &MockPlcDriverManager_GetDriver_Call{Call: _e.mock.On("GetDriver", driverName)}
}

func (_c *MockPlcDriverManager_GetDriver_Call) Run(run func(driverName string)) *MockPlcDriverManager_GetDriver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcDriverManager_GetDriver_Call) Return(_a0 PlcDriver, _a1 error) *MockPlcDriverManager_GetDriver_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlcDriverManager_GetDriver_Call) RunAndReturn(run func(string) (PlcDriver, error)) *MockPlcDriverManager_GetDriver_Call {
	_c.Call.Return(run)
	return _c
}

// ListDriverNames provides a mock function with given fields:
func (_m *MockPlcDriverManager) ListDriverNames() []string {
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

// MockPlcDriverManager_ListDriverNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDriverNames'
type MockPlcDriverManager_ListDriverNames_Call struct {
	*mock.Call
}

// ListDriverNames is a helper method to define mock.On call
func (_e *MockPlcDriverManager_Expecter) ListDriverNames() *MockPlcDriverManager_ListDriverNames_Call {
	return &MockPlcDriverManager_ListDriverNames_Call{Call: _e.mock.On("ListDriverNames")}
}

func (_c *MockPlcDriverManager_ListDriverNames_Call) Run(run func()) *MockPlcDriverManager_ListDriverNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcDriverManager_ListDriverNames_Call) Return(_a0 []string) *MockPlcDriverManager_ListDriverNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriverManager_ListDriverNames_Call) RunAndReturn(run func() []string) *MockPlcDriverManager_ListDriverNames_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterDriver provides a mock function with given fields: driver
func (_m *MockPlcDriverManager) RegisterDriver(driver PlcDriver) {
	_m.Called(driver)
}

// MockPlcDriverManager_RegisterDriver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterDriver'
type MockPlcDriverManager_RegisterDriver_Call struct {
	*mock.Call
}

// RegisterDriver is a helper method to define mock.On call
//   - driver PlcDriver
func (_e *MockPlcDriverManager_Expecter) RegisterDriver(driver interface{}) *MockPlcDriverManager_RegisterDriver_Call {
	return &MockPlcDriverManager_RegisterDriver_Call{Call: _e.mock.On("RegisterDriver", driver)}
}

func (_c *MockPlcDriverManager_RegisterDriver_Call) Run(run func(driver PlcDriver)) *MockPlcDriverManager_RegisterDriver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(PlcDriver))
	})
	return _c
}

func (_c *MockPlcDriverManager_RegisterDriver_Call) Return() *MockPlcDriverManager_RegisterDriver_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPlcDriverManager_RegisterDriver_Call) RunAndReturn(run func(PlcDriver)) *MockPlcDriverManager_RegisterDriver_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcDriverManager creates a new instance of MockPlcDriverManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcDriverManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcDriverManager {
	mock := &MockPlcDriverManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
