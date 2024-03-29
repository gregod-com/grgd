// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/IUIPlugin.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIUIPlugin is a mock of IUIPlugin interface
type MockIUIPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockIUIPluginMockRecorder
}

// MockIUIPluginMockRecorder is the mock recorder for MockIUIPlugin
type MockIUIPluginMockRecorder struct {
	mock *MockIUIPlugin
}

// NewMockIUIPlugin creates a new mock instance
func NewMockIUIPlugin(ctrl *gomock.Controller) *MockIUIPlugin {
	mock := &MockIUIPlugin{ctrl: ctrl}
	mock.recorder = &MockIUIPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUIPlugin) EXPECT() *MockIUIPluginMockRecorder {
	return m.recorder
}

// ClearScreen mocks base method
func (m *MockIUIPlugin) ClearScreen(i ...interface{}) interface{} {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClearScreen", varargs...)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// ClearScreen indicates an expected call of ClearScreen
func (mr *MockIUIPluginMockRecorder) ClearScreen(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearScreen", reflect.TypeOf((*MockIUIPlugin)(nil).ClearScreen), i...)
}

// PrintPercentOfScreen mocks base method
func (m *MockIUIPlugin) PrintPercentOfScreen(startPercent, endPercent int, i ...interface{}) interface{} {
	m.ctrl.T.Helper()
	varargs := []interface{}{startPercent, endPercent}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrintPercentOfScreen", varargs...)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// PrintPercentOfScreen indicates an expected call of PrintPercentOfScreen
func (mr *MockIUIPluginMockRecorder) PrintPercentOfScreen(startPercent, endPercent interface{}, i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{startPercent, endPercent}, i...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintPercentOfScreen", reflect.TypeOf((*MockIUIPlugin)(nil).PrintPercentOfScreen), varargs...)
}

// PrintBanner mocks base method
func (m *MockIUIPlugin) PrintBanner(i ...interface{}) interface{} {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrintBanner", varargs...)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// PrintBanner indicates an expected call of PrintBanner
func (mr *MockIUIPluginMockRecorder) PrintBanner(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintBanner", reflect.TypeOf((*MockIUIPlugin)(nil).PrintBanner), i...)
}

// PrintTable mocks base method
func (m *MockIUIPlugin) PrintTable(heads []string, rows [][]string, i ...interface{}) interface{} {
	m.ctrl.T.Helper()
	varargs := []interface{}{heads, rows}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrintTable", varargs...)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// PrintTable indicates an expected call of PrintTable
func (mr *MockIUIPluginMockRecorder) PrintTable(heads, rows interface{}, i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{heads, rows}, i...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintTable", reflect.TypeOf((*MockIUIPlugin)(nil).PrintTable), varargs...)
}

// Println mocks base method
func (m *MockIUIPlugin) Println(i ...interface{}) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Println", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Println indicates an expected call of Println
func (mr *MockIUIPluginMockRecorder) Println(i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Println", reflect.TypeOf((*MockIUIPlugin)(nil).Println), i...)
}

// Printf mocks base method
func (m *MockIUIPlugin) Printf(format string, i ...interface{}) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Printf", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Printf indicates an expected call of Printf
func (mr *MockIUIPluginMockRecorder) Printf(format interface{}, i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, i...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockIUIPlugin)(nil).Printf), varargs...)
}

// YesNoQuestion mocks base method
func (m *MockIUIPlugin) YesNoQuestion(question string, i ...interface{}) bool {
	m.ctrl.T.Helper()
	varargs := []interface{}{question}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "YesNoQuestion", varargs...)
	ret0, _ := ret[0].(bool)
	return ret0
}

// YesNoQuestion indicates an expected call of YesNoQuestion
func (mr *MockIUIPluginMockRecorder) YesNoQuestion(question interface{}, i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{question}, i...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "YesNoQuestion", reflect.TypeOf((*MockIUIPlugin)(nil).YesNoQuestion), varargs...)
}

// YesNoQuestionf mocks base method
func (m *MockIUIPlugin) YesNoQuestionf(questionf string, i ...interface{}) bool {
	m.ctrl.T.Helper()
	varargs := []interface{}{questionf}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "YesNoQuestionf", varargs...)
	ret0, _ := ret[0].(bool)
	return ret0
}

// YesNoQuestionf indicates an expected call of YesNoQuestionf
func (mr *MockIUIPluginMockRecorder) YesNoQuestionf(questionf interface{}, i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{questionf}, i...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "YesNoQuestionf", reflect.TypeOf((*MockIUIPlugin)(nil).YesNoQuestionf), varargs...)
}

// Question mocks base method
func (m *MockIUIPlugin) Question(question string, i ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{question}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Question", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Question indicates an expected call of Question
func (mr *MockIUIPluginMockRecorder) Question(question interface{}, i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{question}, i...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Question", reflect.TypeOf((*MockIUIPlugin)(nil).Question), varargs...)
}

// Questionf mocks base method
func (m *MockIUIPlugin) Questionf(questionf string, answer *string, i ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{questionf, answer}
	for _, a := range i {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Questionf", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Questionf indicates an expected call of Questionf
func (mr *MockIUIPluginMockRecorder) Questionf(questionf, answer interface{}, i ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{questionf, answer}, i...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Questionf", reflect.TypeOf((*MockIUIPlugin)(nil).Questionf), varargs...)
}
