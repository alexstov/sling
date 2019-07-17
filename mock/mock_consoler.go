// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/4alexey/sling/cui (interfaces: Consoler)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
)

// MockConsoler is a mock of Consoler interface
type MockConsoler struct {
	ctrl     *gomock.Controller
	recorder *MockConsolerMockRecorder
}

// MockConsolerMockRecorder is the mock recorder for MockConsoler
type MockConsolerMockRecorder struct {
	mock *MockConsoler
}

// NewMockConsoler creates a new mock instance
func NewMockConsoler(ctrl *gomock.Controller) *MockConsoler {
	mock := &MockConsoler{ctrl: ctrl}
	mock.recorder = &MockConsolerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsoler) EXPECT() *MockConsolerMockRecorder {
	return m.recorder
}

// GetLogger mocks base method
func (m *MockConsoler) GetLogger() *logrus.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*logrus.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger
func (mr *MockConsolerMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockConsoler)(nil).GetLogger))
}

// Out mocks base method
func (m *MockConsoler) Out(arg0 logrus.Level, arg1 logrus.Fields, arg2 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Out", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Out indicates an expected call of Out
func (mr *MockConsolerMockRecorder) Out(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Out", reflect.TypeOf((*MockConsoler)(nil).Out), varargs...)
}

// OutLogAndConsole mocks base method
func (m *MockConsoler) OutLogAndConsole(arg0 logrus.Level, arg1 logrus.Fields, arg2 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OutLogAndConsole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// OutLogAndConsole indicates an expected call of OutLogAndConsole
func (mr *MockConsolerMockRecorder) OutLogAndConsole(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutLogAndConsole", reflect.TypeOf((*MockConsoler)(nil).OutLogAndConsole), varargs...)
}

// SetFlat mocks base method
func (m *MockConsoler) SetFlat(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFlat", arg0)
}

// SetFlat indicates an expected call of SetFlat
func (mr *MockConsolerMockRecorder) SetFlat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFlat", reflect.TypeOf((*MockConsoler)(nil).SetFlat), arg0)
}

// SetLevel mocks base method
func (m *MockConsoler) SetLevel(arg0 logrus.Level) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLevel", arg0)
}

// SetLevel indicates an expected call of SetLevel
func (mr *MockConsolerMockRecorder) SetLevel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLevel", reflect.TypeOf((*MockConsoler)(nil).SetLevel), arg0)
}