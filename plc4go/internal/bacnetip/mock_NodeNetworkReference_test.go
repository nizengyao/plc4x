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

// MockNodeNetworkReference is an autogenerated mock type for the NodeNetworkReference type
type MockNodeNetworkReference struct {
	mock.Mock
}

type MockNodeNetworkReference_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNodeNetworkReference) EXPECT() *MockNodeNetworkReference_Expecter {
	return &MockNodeNetworkReference_Expecter{mock: &_m.Mock}
}

// AddNode provides a mock function with given fields: node
func (_m *MockNodeNetworkReference) AddNode(node NetworkNode) {
	_m.Called(node)
}

// MockNodeNetworkReference_AddNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNode'
type MockNodeNetworkReference_AddNode_Call struct {
	*mock.Call
}

// AddNode is a helper method to define mock.On call
//   - node NetworkNode
func (_e *MockNodeNetworkReference_Expecter) AddNode(node interface{}) *MockNodeNetworkReference_AddNode_Call {
	return &MockNodeNetworkReference_AddNode_Call{Call: _e.mock.On("AddNode", node)}
}

func (_c *MockNodeNetworkReference_AddNode_Call) Run(run func(node NetworkNode)) *MockNodeNetworkReference_AddNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(NetworkNode))
	})
	return _c
}

func (_c *MockNodeNetworkReference_AddNode_Call) Return() *MockNodeNetworkReference_AddNode_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNodeNetworkReference_AddNode_Call) RunAndReturn(run func(NetworkNode)) *MockNodeNetworkReference_AddNode_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessPDU provides a mock function with given fields: pdu
func (_m *MockNodeNetworkReference) ProcessPDU(pdu PDU) error {
	ret := _m.Called(pdu)

	if len(ret) == 0 {
		panic("no return value specified for ProcessPDU")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNodeNetworkReference_ProcessPDU_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessPDU'
type MockNodeNetworkReference_ProcessPDU_Call struct {
	*mock.Call
}

// ProcessPDU is a helper method to define mock.On call
//   - pdu PDU
func (_e *MockNodeNetworkReference_Expecter) ProcessPDU(pdu interface{}) *MockNodeNetworkReference_ProcessPDU_Call {
	return &MockNodeNetworkReference_ProcessPDU_Call{Call: _e.mock.On("ProcessPDU", pdu)}
}

func (_c *MockNodeNetworkReference_ProcessPDU_Call) Run(run func(pdu PDU)) *MockNodeNetworkReference_ProcessPDU_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(PDU))
	})
	return _c
}

func (_c *MockNodeNetworkReference_ProcessPDU_Call) Return(_a0 error) *MockNodeNetworkReference_ProcessPDU_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNodeNetworkReference_ProcessPDU_Call) RunAndReturn(run func(PDU) error) *MockNodeNetworkReference_ProcessPDU_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNodeNetworkReference creates a new instance of MockNodeNetworkReference. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNodeNetworkReference(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNodeNetworkReference {
	mock := &MockNodeNetworkReference{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
