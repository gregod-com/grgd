// Code generated by MockGen. DO NOT EDIT.
// Source: IHelper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	interfaces "grgd/interfaces"
	reflect "reflect"
)

// MockIHelper is a mock of IHelper interface
type MockIHelper struct {
	ctrl     *gomock.Controller
	recorder *MockIHelperMockRecorder
}

// MockIHelperMockRecorder is the mock recorder for MockIHelper
type MockIHelperMockRecorder struct {
	mock *MockIHelper
}

// NewMockIHelper creates a new mock instance
func NewMockIHelper(ctrl *gomock.Controller) *MockIHelper {
	mock := &MockIHelper{ctrl: ctrl}
	mock.recorder = &MockIHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIHelper) EXPECT() *MockIHelperMockRecorder {
	return m.recorder
}

// CheckUserProfile mocks base method
func (m *MockIHelper) CheckUserProfile(logger interfaces.ILogger) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserProfile", logger)
	ret0, _ := ret[0].(string)
	return ret0
}

// CheckUserProfile indicates an expected call of CheckUserProfile
func (mr *MockIHelperMockRecorder) CheckUserProfile(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserProfile", reflect.TypeOf((*MockIHelper)(nil).CheckUserProfile), logger)
}

// CheckFlag mocks base method
func (m *MockIHelper) CheckFlag(flag string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckFlag", flag)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckFlag indicates an expected call of CheckFlag
func (mr *MockIHelperMockRecorder) CheckFlag(flag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckFlag", reflect.TypeOf((*MockIHelper)(nil).CheckFlag), flag)
}

// CheckFlagArg mocks base method
func (m *MockIHelper) CheckFlagArg(flag string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckFlagArg", flag)
	ret0, _ := ret[0].(string)
	return ret0
}

// CheckFlagArg indicates an expected call of CheckFlagArg
func (mr *MockIHelperMockRecorder) CheckFlagArg(flag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckFlagArg", reflect.TypeOf((*MockIHelper)(nil).CheckFlagArg), flag)
}