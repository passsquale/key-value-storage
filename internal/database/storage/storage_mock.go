// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package storage is a generated GoMock package.
package storage

import (
	context "context"
	reflect "reflect"
	wal "github.com/passsquale/key-value-storage/internal/database/storage/wal"
	tools "github.com/passsquale/key-value-storage/internal/tools"

	gomock "github.com/golang/mock/gomock"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockEngine) Del(arg0 context.Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Del", arg0, arg1)
}

// Del indicates an expected call of Del.
func (mr *MockEngineMockRecorder) Del(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockEngine)(nil).Del), arg0, arg1)
}

// Get mocks base method.
func (m *MockEngine) Get(arg0 context.Context, arg1 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockEngineMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEngine)(nil).Get), arg0, arg1)
}

// Set mocks base method.
func (m *MockEngine) Set(arg0 context.Context, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1, arg2)
}

// Set indicates an expected call of Set.
func (mr *MockEngineMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockEngine)(nil).Set), arg0, arg1, arg2)
}

// MockWAL is a mock of WAL interface.
type MockWAL struct {
	ctrl     *gomock.Controller
	recorder *MockWALMockRecorder
}

// MockWALMockRecorder is the mock recorder for MockWAL.
type MockWALMockRecorder struct {
	mock *MockWAL
}

// NewMockWAL creates a new mock instance.
func NewMockWAL(ctrl *gomock.Controller) *MockWAL {
	mock := &MockWAL{ctrl: ctrl}
	mock.recorder = &MockWALMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWAL) EXPECT() *MockWALMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockWAL) Del(arg0 context.Context, arg1 string) tools.FutureError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", arg0, arg1)
	ret0, _ := ret[0].(tools.FutureError)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockWALMockRecorder) Del(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockWAL)(nil).Del), arg0, arg1)
}

// Recover mocks base method.
func (m *MockWAL) Recover() ([]wal.LogData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recover")
	ret0, _ := ret[0].([]wal.LogData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recover indicates an expected call of Recover.
func (mr *MockWALMockRecorder) Recover() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recover", reflect.TypeOf((*MockWAL)(nil).Recover))
}

// Set mocks base method.
func (m *MockWAL) Set(arg0 context.Context, arg1, arg2 string) tools.FutureError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(tools.FutureError)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockWALMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockWAL)(nil).Set), arg0, arg1, arg2)
}

// Shutdown mocks base method.
func (m *MockWAL) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockWALMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockWAL)(nil).Shutdown))
}

// Start mocks base method.
func (m *MockWAL) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockWALMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockWAL)(nil).Start))
}
