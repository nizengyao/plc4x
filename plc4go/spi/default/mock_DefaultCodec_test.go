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

package _default

import (
	context "context"

	spi "github.com/apache/plc4x/plc4go/spi"
	mock "github.com/stretchr/testify/mock"

	time "time"

	transports "github.com/apache/plc4x/plc4go/spi/transports"

	utils "github.com/apache/plc4x/plc4go/spi/utils"
)

// MockDefaultCodec is an autogenerated mock type for the DefaultCodec type
type MockDefaultCodec struct {
	mock.Mock
}

type MockDefaultCodec_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDefaultCodec) EXPECT() *MockDefaultCodec_Expecter {
	return &MockDefaultCodec_Expecter{mock: &_m.Mock}
}

// Connect provides a mock function with given fields:
func (_m *MockDefaultCodec) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultCodec_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockDefaultCodec_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockDefaultCodec_Expecter) Connect() *MockDefaultCodec_Connect_Call {
	return &MockDefaultCodec_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockDefaultCodec_Connect_Call) Run(run func()) *MockDefaultCodec_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultCodec_Connect_Call) Return(_a0 error) *MockDefaultCodec_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultCodec_Connect_Call) RunAndReturn(run func() error) *MockDefaultCodec_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// ConnectWithContext provides a mock function with given fields: ctx
func (_m *MockDefaultCodec) ConnectWithContext(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultCodec_ConnectWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectWithContext'
type MockDefaultCodec_ConnectWithContext_Call struct {
	*mock.Call
}

// ConnectWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockDefaultCodec_Expecter) ConnectWithContext(ctx interface{}) *MockDefaultCodec_ConnectWithContext_Call {
	return &MockDefaultCodec_ConnectWithContext_Call{Call: _e.mock.On("ConnectWithContext", ctx)}
}

func (_c *MockDefaultCodec_ConnectWithContext_Call) Run(run func(ctx context.Context)) *MockDefaultCodec_ConnectWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockDefaultCodec_ConnectWithContext_Call) Return(_a0 error) *MockDefaultCodec_ConnectWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultCodec_ConnectWithContext_Call) RunAndReturn(run func(context.Context) error) *MockDefaultCodec_ConnectWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// Disconnect provides a mock function with given fields:
func (_m *MockDefaultCodec) Disconnect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultCodec_Disconnect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Disconnect'
type MockDefaultCodec_Disconnect_Call struct {
	*mock.Call
}

// Disconnect is a helper method to define mock.On call
func (_e *MockDefaultCodec_Expecter) Disconnect() *MockDefaultCodec_Disconnect_Call {
	return &MockDefaultCodec_Disconnect_Call{Call: _e.mock.On("Disconnect")}
}

func (_c *MockDefaultCodec_Disconnect_Call) Run(run func()) *MockDefaultCodec_Disconnect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultCodec_Disconnect_Call) Return(_a0 error) *MockDefaultCodec_Disconnect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultCodec_Disconnect_Call) RunAndReturn(run func() error) *MockDefaultCodec_Disconnect_Call {
	_c.Call.Return(run)
	return _c
}

// Expect provides a mock function with given fields: ctx, acceptsMessage, handleMessage, handleError, ttl
func (_m *MockDefaultCodec) Expect(ctx context.Context, acceptsMessage spi.AcceptsMessage, handleMessage spi.HandleMessage, handleError spi.HandleError, ttl time.Duration) error {
	ret := _m.Called(ctx, acceptsMessage, handleMessage, handleError, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, spi.AcceptsMessage, spi.HandleMessage, spi.HandleError, time.Duration) error); ok {
		r0 = rf(ctx, acceptsMessage, handleMessage, handleError, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultCodec_Expect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Expect'
type MockDefaultCodec_Expect_Call struct {
	*mock.Call
}

// Expect is a helper method to define mock.On call
//   - ctx context.Context
//   - acceptsMessage spi.AcceptsMessage
//   - handleMessage spi.HandleMessage
//   - handleError spi.HandleError
//   - ttl time.Duration
func (_e *MockDefaultCodec_Expecter) Expect(ctx interface{}, acceptsMessage interface{}, handleMessage interface{}, handleError interface{}, ttl interface{}) *MockDefaultCodec_Expect_Call {
	return &MockDefaultCodec_Expect_Call{Call: _e.mock.On("Expect", ctx, acceptsMessage, handleMessage, handleError, ttl)}
}

func (_c *MockDefaultCodec_Expect_Call) Run(run func(ctx context.Context, acceptsMessage spi.AcceptsMessage, handleMessage spi.HandleMessage, handleError spi.HandleError, ttl time.Duration)) *MockDefaultCodec_Expect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(spi.AcceptsMessage), args[2].(spi.HandleMessage), args[3].(spi.HandleError), args[4].(time.Duration))
	})
	return _c
}

func (_c *MockDefaultCodec_Expect_Call) Return(_a0 error) *MockDefaultCodec_Expect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultCodec_Expect_Call) RunAndReturn(run func(context.Context, spi.AcceptsMessage, spi.HandleMessage, spi.HandleError, time.Duration) error) *MockDefaultCodec_Expect_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultIncomingMessageChannel provides a mock function with given fields:
func (_m *MockDefaultCodec) GetDefaultIncomingMessageChannel() chan spi.Message {
	ret := _m.Called()

	var r0 chan spi.Message
	if rf, ok := ret.Get(0).(func() chan spi.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan spi.Message)
		}
	}

	return r0
}

// MockDefaultCodec_GetDefaultIncomingMessageChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultIncomingMessageChannel'
type MockDefaultCodec_GetDefaultIncomingMessageChannel_Call struct {
	*mock.Call
}

// GetDefaultIncomingMessageChannel is a helper method to define mock.On call
func (_e *MockDefaultCodec_Expecter) GetDefaultIncomingMessageChannel() *MockDefaultCodec_GetDefaultIncomingMessageChannel_Call {
	return &MockDefaultCodec_GetDefaultIncomingMessageChannel_Call{Call: _e.mock.On("GetDefaultIncomingMessageChannel")}
}

func (_c *MockDefaultCodec_GetDefaultIncomingMessageChannel_Call) Run(run func()) *MockDefaultCodec_GetDefaultIncomingMessageChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultCodec_GetDefaultIncomingMessageChannel_Call) Return(_a0 chan spi.Message) *MockDefaultCodec_GetDefaultIncomingMessageChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultCodec_GetDefaultIncomingMessageChannel_Call) RunAndReturn(run func() chan spi.Message) *MockDefaultCodec_GetDefaultIncomingMessageChannel_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransportInstance provides a mock function with given fields:
func (_m *MockDefaultCodec) GetTransportInstance() transports.TransportInstance {
	ret := _m.Called()

	var r0 transports.TransportInstance
	if rf, ok := ret.Get(0).(func() transports.TransportInstance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transports.TransportInstance)
		}
	}

	return r0
}

// MockDefaultCodec_GetTransportInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransportInstance'
type MockDefaultCodec_GetTransportInstance_Call struct {
	*mock.Call
}

// GetTransportInstance is a helper method to define mock.On call
func (_e *MockDefaultCodec_Expecter) GetTransportInstance() *MockDefaultCodec_GetTransportInstance_Call {
	return &MockDefaultCodec_GetTransportInstance_Call{Call: _e.mock.On("GetTransportInstance")}
}

func (_c *MockDefaultCodec_GetTransportInstance_Call) Run(run func()) *MockDefaultCodec_GetTransportInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultCodec_GetTransportInstance_Call) Return(_a0 transports.TransportInstance) *MockDefaultCodec_GetTransportInstance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultCodec_GetTransportInstance_Call) RunAndReturn(run func() transports.TransportInstance) *MockDefaultCodec_GetTransportInstance_Call {
	_c.Call.Return(run)
	return _c
}

// IsRunning provides a mock function with given fields:
func (_m *MockDefaultCodec) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockDefaultCodec_IsRunning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRunning'
type MockDefaultCodec_IsRunning_Call struct {
	*mock.Call
}

// IsRunning is a helper method to define mock.On call
func (_e *MockDefaultCodec_Expecter) IsRunning() *MockDefaultCodec_IsRunning_Call {
	return &MockDefaultCodec_IsRunning_Call{Call: _e.mock.On("IsRunning")}
}

func (_c *MockDefaultCodec_IsRunning_Call) Run(run func()) *MockDefaultCodec_IsRunning_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultCodec_IsRunning_Call) Return(_a0 bool) *MockDefaultCodec_IsRunning_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultCodec_IsRunning_Call) RunAndReturn(run func() bool) *MockDefaultCodec_IsRunning_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: message
func (_m *MockDefaultCodec) Send(message spi.Message) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(spi.Message) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultCodec_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockDefaultCodec_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - message spi.Message
func (_e *MockDefaultCodec_Expecter) Send(message interface{}) *MockDefaultCodec_Send_Call {
	return &MockDefaultCodec_Send_Call{Call: _e.mock.On("Send", message)}
}

func (_c *MockDefaultCodec_Send_Call) Run(run func(message spi.Message)) *MockDefaultCodec_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spi.Message))
	})
	return _c
}

func (_c *MockDefaultCodec_Send_Call) Return(_a0 error) *MockDefaultCodec_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultCodec_Send_Call) RunAndReturn(run func(spi.Message) error) *MockDefaultCodec_Send_Call {
	_c.Call.Return(run)
	return _c
}

// SendRequest provides a mock function with given fields: ctx, message, acceptsMessage, handleMessage, handleError, ttl
func (_m *MockDefaultCodec) SendRequest(ctx context.Context, message spi.Message, acceptsMessage spi.AcceptsMessage, handleMessage spi.HandleMessage, handleError spi.HandleError, ttl time.Duration) error {
	ret := _m.Called(ctx, message, acceptsMessage, handleMessage, handleError, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, spi.Message, spi.AcceptsMessage, spi.HandleMessage, spi.HandleError, time.Duration) error); ok {
		r0 = rf(ctx, message, acceptsMessage, handleMessage, handleError, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultCodec_SendRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendRequest'
type MockDefaultCodec_SendRequest_Call struct {
	*mock.Call
}

// SendRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - message spi.Message
//   - acceptsMessage spi.AcceptsMessage
//   - handleMessage spi.HandleMessage
//   - handleError spi.HandleError
//   - ttl time.Duration
func (_e *MockDefaultCodec_Expecter) SendRequest(ctx interface{}, message interface{}, acceptsMessage interface{}, handleMessage interface{}, handleError interface{}, ttl interface{}) *MockDefaultCodec_SendRequest_Call {
	return &MockDefaultCodec_SendRequest_Call{Call: _e.mock.On("SendRequest", ctx, message, acceptsMessage, handleMessage, handleError, ttl)}
}

func (_c *MockDefaultCodec_SendRequest_Call) Run(run func(ctx context.Context, message spi.Message, acceptsMessage spi.AcceptsMessage, handleMessage spi.HandleMessage, handleError spi.HandleError, ttl time.Duration)) *MockDefaultCodec_SendRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(spi.Message), args[2].(spi.AcceptsMessage), args[3].(spi.HandleMessage), args[4].(spi.HandleError), args[5].(time.Duration))
	})
	return _c
}

func (_c *MockDefaultCodec_SendRequest_Call) Return(_a0 error) *MockDefaultCodec_SendRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultCodec_SendRequest_Call) RunAndReturn(run func(context.Context, spi.Message, spi.AcceptsMessage, spi.HandleMessage, spi.HandleError, time.Duration) error) *MockDefaultCodec_SendRequest_Call {
	_c.Call.Return(run)
	return _c
}

// Serialize provides a mock function with given fields:
func (_m *MockDefaultCodec) Serialize() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDefaultCodec_Serialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serialize'
type MockDefaultCodec_Serialize_Call struct {
	*mock.Call
}

// Serialize is a helper method to define mock.On call
func (_e *MockDefaultCodec_Expecter) Serialize() *MockDefaultCodec_Serialize_Call {
	return &MockDefaultCodec_Serialize_Call{Call: _e.mock.On("Serialize")}
}

func (_c *MockDefaultCodec_Serialize_Call) Run(run func()) *MockDefaultCodec_Serialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultCodec_Serialize_Call) Return(_a0 []byte, _a1 error) *MockDefaultCodec_Serialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDefaultCodec_Serialize_Call) RunAndReturn(run func() ([]byte, error)) *MockDefaultCodec_Serialize_Call {
	_c.Call.Return(run)
	return _c
}

// SerializeWithWriteBuffer provides a mock function with given fields: ctx, writeBuffer
func (_m *MockDefaultCodec) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	ret := _m.Called(ctx, writeBuffer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, utils.WriteBuffer) error); ok {
		r0 = rf(ctx, writeBuffer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultCodec_SerializeWithWriteBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SerializeWithWriteBuffer'
type MockDefaultCodec_SerializeWithWriteBuffer_Call struct {
	*mock.Call
}

// SerializeWithWriteBuffer is a helper method to define mock.On call
//   - ctx context.Context
//   - writeBuffer utils.WriteBuffer
func (_e *MockDefaultCodec_Expecter) SerializeWithWriteBuffer(ctx interface{}, writeBuffer interface{}) *MockDefaultCodec_SerializeWithWriteBuffer_Call {
	return &MockDefaultCodec_SerializeWithWriteBuffer_Call{Call: _e.mock.On("SerializeWithWriteBuffer", ctx, writeBuffer)}
}

func (_c *MockDefaultCodec_SerializeWithWriteBuffer_Call) Run(run func(ctx context.Context, writeBuffer utils.WriteBuffer)) *MockDefaultCodec_SerializeWithWriteBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(utils.WriteBuffer))
	})
	return _c
}

func (_c *MockDefaultCodec_SerializeWithWriteBuffer_Call) Return(_a0 error) *MockDefaultCodec_SerializeWithWriteBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultCodec_SerializeWithWriteBuffer_Call) RunAndReturn(run func(context.Context, utils.WriteBuffer) error) *MockDefaultCodec_SerializeWithWriteBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDefaultCodec creates a new instance of MockDefaultCodec. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDefaultCodec(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDefaultCodec {
	mock := &MockDefaultCodec{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
