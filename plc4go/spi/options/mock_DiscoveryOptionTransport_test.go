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

package options

import mock "github.com/stretchr/testify/mock"

// MockDiscoveryOptionTransport is an autogenerated mock type for the DiscoveryOptionTransport type
type MockDiscoveryOptionTransport struct {
	mock.Mock
}

type MockDiscoveryOptionTransport_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDiscoveryOptionTransport) EXPECT() *MockDiscoveryOptionTransport_Expecter {
	return &MockDiscoveryOptionTransport_Expecter{mock: &_m.Mock}
}

// GetTransportName provides a mock function with given fields:
func (_m *MockDiscoveryOptionTransport) GetTransportName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockDiscoveryOptionTransport_GetTransportName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransportName'
type MockDiscoveryOptionTransport_GetTransportName_Call struct {
	*mock.Call
}

// GetTransportName is a helper method to define mock.On call
func (_e *MockDiscoveryOptionTransport_Expecter) GetTransportName() *MockDiscoveryOptionTransport_GetTransportName_Call {
	return &MockDiscoveryOptionTransport_GetTransportName_Call{Call: _e.mock.On("GetTransportName")}
}

func (_c *MockDiscoveryOptionTransport_GetTransportName_Call) Run(run func()) *MockDiscoveryOptionTransport_GetTransportName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDiscoveryOptionTransport_GetTransportName_Call) Return(_a0 string) *MockDiscoveryOptionTransport_GetTransportName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDiscoveryOptionTransport_GetTransportName_Call) RunAndReturn(run func() string) *MockDiscoveryOptionTransport_GetTransportName_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDiscoveryOptionTransport creates a new instance of MockDiscoveryOptionTransport. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDiscoveryOptionTransport(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDiscoveryOptionTransport {
	mock := &MockDiscoveryOptionTransport{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
