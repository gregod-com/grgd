// Code generated by MockGen. DO NOT EDIT.
// Source: IProject.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/gregod-com/grgd/interfaces"
	reflect "reflect"
)

// MockIProject is a mock of IProject interface
type MockIProject struct {
	ctrl     *gomock.Controller
	recorder *MockIProjectMockRecorder
}

// MockIProjectMockRecorder is the mock recorder for MockIProject
type MockIProjectMockRecorder struct {
	mock *MockIProject
}

// NewMockIProject creates a new mock instance
func NewMockIProject(ctrl *gomock.Controller) *MockIProject {
	mock := &MockIProject{ctrl: ctrl}
	mock.recorder = &MockIProjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIProject) EXPECT() *MockIProjectMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockIProject) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockIProjectMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockIProject)(nil).GetName))
}

// GetID mocks base method
func (m *MockIProject) GetID(i ...interface{}) uint {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetID", varargs...)
	ret0, _ := ret[0].(uint)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockIProjectMockRecorder) GetID(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockIProject)(nil).GetID), i...)
}

// GetPath mocks base method
func (m *MockIProject) GetPath(i ...interface{}) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPath", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath
func (mr *MockIProjectMockRecorder) GetPath(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockIProject)(nil).GetPath), i...)
}

// SetPath mocks base method
func (m *MockIProject) SetPath(path string, i ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetPath", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPath indicates an expected call of SetPath
func (mr *MockIProjectMockRecorder) SetPath(path interface{}, i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, i...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPath", reflect.TypeOf((*MockIProject)(nil).SetPath), varargs...)
}

// GetServices mocks base method
func (m *MockIProject) GetServices(i ...interface{}) map[string]interfaces.IService {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetServices", varargs...)
	ret0, _ := ret[0].(map[string]interfaces.IService)
	return ret0
}

// GetServices indicates an expected call of GetServices
func (mr *MockIProjectMockRecorder) GetServices(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServices", reflect.TypeOf((*MockIProject)(nil).GetServices), i...)
}

// GetServiceByName mocks base method
func (m *MockIProject) GetServiceByName(serviceName string, i ...interface{}) interfaces.IService {
	m.ctrl.T.Helper()
	varargs := []interface{}{serviceName}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetServiceByName", varargs...)
	ret0, _ := ret[0].(interfaces.IService)
	return ret0
}

// GetServiceByName indicates an expected call of GetServiceByName
func (mr *MockIProjectMockRecorder) GetServiceByName(serviceName interface{}, i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{serviceName}, i...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceByName", reflect.TypeOf((*MockIProject)(nil).GetServiceByName), varargs...)
}

// GetValues mocks base method
func (m *MockIProject) GetValues(i ...interface{}) []string {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetValues", varargs...)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetValues indicates an expected call of GetValues
func (mr *MockIProjectMockRecorder) GetValues(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValues", reflect.TypeOf((*MockIProject)(nil).GetValues), i...)
}
