// Code generated by MockGen. DO NOT EDIT.
// Source: IPinger.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIPinger is a mock of IPinger interface
type MockIPinger struct {
	ctrl     *gomock.Controller
	recorder *MockIPingerMockRecorder
}

// MockIPingerMockRecorder is the mock recorder for MockIPinger
type MockIPingerMockRecorder struct {
	mock *MockIPinger
}

// NewMockIPinger creates a new mock instance
func NewMockIPinger(ctrl *gomock.Controller) *MockIPinger {
	mock := &MockIPinger{ctrl: ctrl}
	mock.recorder = &MockIPingerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPinger) EXPECT() *MockIPingerMockRecorder {
	return m.recorder
}

// CheckConnections mocks base method
func (m *MockIPinger) CheckConnections(conns map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CheckConnections", conns)
}

// CheckConnections indicates an expected call of CheckConnections
func (mr *MockIPingerMockRecorder) CheckConnections(conns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckConnections", reflect.TypeOf((*MockIPinger)(nil).CheckConnections), conns)
}