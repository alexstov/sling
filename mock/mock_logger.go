// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/4alexey/sling/slog (interfaces: Logger)

// Package mock is a generated GoMock package.
package mock

import (
	conf "github.com/4alexey/sling/conf"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockLogger) Configure(arg0 *conf.Log) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Configure indicates an expected call of Configure
func (mr *MockLoggerMockRecorder) Configure(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockLogger)(nil).Configure), arg0)
}

// GetLogger mocks base method
func (m *MockLogger) GetLogger() *logrus.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*logrus.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger
func (mr *MockLoggerMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockLogger)(nil).GetLogger))
}

// Out mocks base method
func (m *MockLogger) Out(arg0 logrus.Level, arg1 logrus.Fields, arg2 ...interface{}) error {
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
func (mr *MockLoggerMockRecorder) Out(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Out", reflect.TypeOf((*MockLogger)(nil).Out), varargs...)
}

// SetLevel mocks base method
func (m *MockLogger) SetLevel(arg0 logrus.Level) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLevel", arg0)
}

// SetLevel indicates an expected call of SetLevel
func (mr *MockLoggerMockRecorder) SetLevel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLevel", reflect.TypeOf((*MockLogger)(nil).SetLevel), arg0)
}
