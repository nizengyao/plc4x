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

// MockObjectTypeContract is an autogenerated mock type for the ObjectTypeContract type
type MockObjectTypeContract struct {
	mock.Mock
}

type MockObjectTypeContract_Expecter struct {
	mock *mock.Mock
}

func (_m *MockObjectTypeContract) EXPECT() *MockObjectTypeContract_Expecter {
	return &MockObjectTypeContract_Expecter{mock: &_m.Mock}
}

// GetEnumerations provides a mock function with given fields:
func (_m *MockObjectTypeContract) GetEnumerations() map[string]uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEnumerations")
	}

	var r0 map[string]uint64
	if rf, ok := ret.Get(0).(func() map[string]uint64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]uint64)
		}
	}

	return r0
}

// MockObjectTypeContract_GetEnumerations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEnumerations'
type MockObjectTypeContract_GetEnumerations_Call struct {
	*mock.Call
}

// GetEnumerations is a helper method to define mock.On call
func (_e *MockObjectTypeContract_Expecter) GetEnumerations() *MockObjectTypeContract_GetEnumerations_Call {
	return &MockObjectTypeContract_GetEnumerations_Call{Call: _e.mock.On("GetEnumerations")}
}

func (_c *MockObjectTypeContract_GetEnumerations_Call) Run(run func()) *MockObjectTypeContract_GetEnumerations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockObjectTypeContract_GetEnumerations_Call) Return(_a0 map[string]uint64) *MockObjectTypeContract_GetEnumerations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockObjectTypeContract_GetEnumerations_Call) RunAndReturn(run func() map[string]uint64) *MockObjectTypeContract_GetEnumerations_Call {
	_c.Call.Return(run)
	return _c
}

// GetXlateTable provides a mock function with given fields:
func (_m *MockObjectTypeContract) GetXlateTable() map[interface{}]interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetXlateTable")
	}

	var r0 map[interface{}]interface{}
	if rf, ok := ret.Get(0).(func() map[interface{}]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[interface{}]interface{})
		}
	}

	return r0
}

// MockObjectTypeContract_GetXlateTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetXlateTable'
type MockObjectTypeContract_GetXlateTable_Call struct {
	*mock.Call
}

// GetXlateTable is a helper method to define mock.On call
func (_e *MockObjectTypeContract_Expecter) GetXlateTable() *MockObjectTypeContract_GetXlateTable_Call {
	return &MockObjectTypeContract_GetXlateTable_Call{Call: _e.mock.On("GetXlateTable")}
}

func (_c *MockObjectTypeContract_GetXlateTable_Call) Run(run func()) *MockObjectTypeContract_GetXlateTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockObjectTypeContract_GetXlateTable_Call) Return(_a0 map[interface{}]interface{}) *MockObjectTypeContract_GetXlateTable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockObjectTypeContract_GetXlateTable_Call) RunAndReturn(run func() map[interface{}]interface{}) *MockObjectTypeContract_GetXlateTable_Call {
	_c.Call.Return(run)
	return _c
}

// SetEnumerated provides a mock function with given fields: enumerated
func (_m *MockObjectTypeContract) SetEnumerated(enumerated *Enumerated) {
	_m.Called(enumerated)
}

// MockObjectTypeContract_SetEnumerated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetEnumerated'
type MockObjectTypeContract_SetEnumerated_Call struct {
	*mock.Call
}

// SetEnumerated is a helper method to define mock.On call
//   - enumerated *Enumerated
func (_e *MockObjectTypeContract_Expecter) SetEnumerated(enumerated interface{}) *MockObjectTypeContract_SetEnumerated_Call {
	return &MockObjectTypeContract_SetEnumerated_Call{Call: _e.mock.On("SetEnumerated", enumerated)}
}

func (_c *MockObjectTypeContract_SetEnumerated_Call) Run(run func(enumerated *Enumerated)) *MockObjectTypeContract_SetEnumerated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Enumerated))
	})
	return _c
}

func (_c *MockObjectTypeContract_SetEnumerated_Call) Return() *MockObjectTypeContract_SetEnumerated_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockObjectTypeContract_SetEnumerated_Call) RunAndReturn(run func(*Enumerated)) *MockObjectTypeContract_SetEnumerated_Call {
	_c.Call.Return(run)
	return _c
}

// SetObjectType provides a mock function with given fields: objectType
func (_m *MockObjectTypeContract) SetObjectType(objectType *ObjectType) {
	_m.Called(objectType)
}

// MockObjectTypeContract_SetObjectType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetObjectType'
type MockObjectTypeContract_SetObjectType_Call struct {
	*mock.Call
}

// SetObjectType is a helper method to define mock.On call
//   - objectType *ObjectType
func (_e *MockObjectTypeContract_Expecter) SetObjectType(objectType interface{}) *MockObjectTypeContract_SetObjectType_Call {
	return &MockObjectTypeContract_SetObjectType_Call{Call: _e.mock.On("SetObjectType", objectType)}
}

func (_c *MockObjectTypeContract_SetObjectType_Call) Run(run func(objectType *ObjectType)) *MockObjectTypeContract_SetObjectType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*ObjectType))
	})
	return _c
}

func (_c *MockObjectTypeContract_SetObjectType_Call) Return() *MockObjectTypeContract_SetObjectType_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockObjectTypeContract_SetObjectType_Call) RunAndReturn(run func(*ObjectType)) *MockObjectTypeContract_SetObjectType_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockObjectTypeContract creates a new instance of MockObjectTypeContract. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockObjectTypeContract(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockObjectTypeContract {
	mock := &MockObjectTypeContract{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
