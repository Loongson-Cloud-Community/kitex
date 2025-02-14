/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
*/


// Code generated by MockGen. DO NOT EDIT.
// Source: ../../../netpoll/connection.go

// Package netpoll is a generated GoMock package.
package netpoll

import (
	net "net"
	reflect "reflect"
	time "time"

	netpoll "github.com/cloudwego/netpoll"
	gomock "github.com/golang/mock/gomock"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// AddCloseCallback mocks base method.
func (m *MockConnection) AddCloseCallback(callback netpoll.CloseCallback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCloseCallback", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCloseCallback indicates an expected call of AddCloseCallback.
func (mr *MockConnectionMockRecorder) AddCloseCallback(callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCloseCallback", reflect.TypeOf((*MockConnection)(nil).AddCloseCallback), callback)
}

// Close mocks base method.
func (m *MockConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// IsActive mocks base method.
func (m *MockConnection) IsActive() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsActive indicates an expected call of IsActive.
func (mr *MockConnectionMockRecorder) IsActive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockConnection)(nil).IsActive))
}

// LocalAddr mocks base method.
func (m *MockConnection) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockConnectionMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockConnection)(nil).LocalAddr))
}

// Read mocks base method.
func (m *MockConnection) Read(b []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockConnectionMockRecorder) Read(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockConnection)(nil).Read), b)
}

// Reader mocks base method.
func (m *MockConnection) Reader() netpoll.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reader")
	ret0, _ := ret[0].(netpoll.Reader)
	return ret0
}

// Reader indicates an expected call of Reader.
func (mr *MockConnectionMockRecorder) Reader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockConnection)(nil).Reader))
}

// RemoteAddr mocks base method.
func (m *MockConnection) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockConnectionMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockConnection)(nil).RemoteAddr))
}

// SetDeadline mocks base method.
func (m *MockConnection) SetDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockConnectionMockRecorder) SetDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockConnection)(nil).SetDeadline), t)
}

// SetIdleTimeout mocks base method.
func (m *MockConnection) SetIdleTimeout(timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIdleTimeout", timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetIdleTimeout indicates an expected call of SetIdleTimeout.
func (mr *MockConnectionMockRecorder) SetIdleTimeout(timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIdleTimeout", reflect.TypeOf((*MockConnection)(nil).SetIdleTimeout), timeout)
}

// SetOnRequest mocks base method.
func (m *MockConnection) SetOnRequest(on netpoll.OnRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOnRequest", on)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOnRequest indicates an expected call of SetOnRequest.
func (mr *MockConnectionMockRecorder) SetOnRequest(on interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOnRequest", reflect.TypeOf((*MockConnection)(nil).SetOnRequest), on)
}

// SetReadDeadline mocks base method.
func (m *MockConnection) SetReadDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockConnectionMockRecorder) SetReadDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockConnection)(nil).SetReadDeadline), t)
}

// SetReadTimeout mocks base method.
func (m *MockConnection) SetReadTimeout(timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadTimeout", timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadTimeout indicates an expected call of SetReadTimeout.
func (mr *MockConnectionMockRecorder) SetReadTimeout(timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadTimeout", reflect.TypeOf((*MockConnection)(nil).SetReadTimeout), timeout)
}

// SetWriteDeadline mocks base method.
func (m *MockConnection) SetWriteDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockConnectionMockRecorder) SetWriteDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockConnection)(nil).SetWriteDeadline), t)
}

// Write mocks base method.
func (m *MockConnection) Write(b []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockConnectionMockRecorder) Write(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockConnection)(nil).Write), b)
}

// Writer mocks base method.
func (m *MockConnection) Writer() netpoll.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Writer")
	ret0, _ := ret[0].(netpoll.Writer)
	return ret0
}

// Writer indicates an expected call of Writer.
func (mr *MockConnectionMockRecorder) Writer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Writer", reflect.TypeOf((*MockConnection)(nil).Writer))
}
