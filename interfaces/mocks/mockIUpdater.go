// Code generated by MockGen. DO NOT EDIT.
// Source: IUpdater.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/gregod-com/grgd/interfaces"
	reflect "reflect"
)

// MockIUpdater is a mock of IUpdater interface
type MockIUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockIUpdaterMockRecorder
}

// MockIUpdaterMockRecorder is the mock recorder for MockIUpdater
type MockIUpdaterMockRecorder struct {
	mock *MockIUpdater
}

// NewMockIUpdater creates a new mock instance
func NewMockIUpdater(ctrl *gomock.Controller) *MockIUpdater {
	mock := &MockIUpdater{ctrl: ctrl}
	mock.recorder = &MockIUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUpdater) EXPECT() *MockIUpdaterMockRecorder {
	return m.recorder
}

// CheckUpdate mocks base method
func (m *MockIUpdater) CheckUpdate(core interfaces.ICore) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUpdate", core)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUpdate indicates an expected call of CheckUpdate
func (mr *MockIUpdaterMockRecorder) CheckUpdate(core interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUpdate", reflect.TypeOf((*MockIUpdater)(nil).CheckUpdate), core)
}
