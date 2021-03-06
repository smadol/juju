// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/charmhub (interfaces: InfoCommandAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	charmhub "github.com/juju/juju/api/charmhub"
	reflect "reflect"
)

// MockInfoCommandAPI is a mock of InfoCommandAPI interface
type MockInfoCommandAPI struct {
	ctrl     *gomock.Controller
	recorder *MockInfoCommandAPIMockRecorder
}

// MockInfoCommandAPIMockRecorder is the mock recorder for MockInfoCommandAPI
type MockInfoCommandAPIMockRecorder struct {
	mock *MockInfoCommandAPI
}

// NewMockInfoCommandAPI creates a new mock instance
func NewMockInfoCommandAPI(ctrl *gomock.Controller) *MockInfoCommandAPI {
	mock := &MockInfoCommandAPI{ctrl: ctrl}
	mock.recorder = &MockInfoCommandAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfoCommandAPI) EXPECT() *MockInfoCommandAPIMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockInfoCommandAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockInfoCommandAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockInfoCommandAPI)(nil).Close))
}

// Info mocks base method
func (m *MockInfoCommandAPI) Info(arg0 string) (charmhub.InfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0)
	ret0, _ := ret[0].(charmhub.InfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockInfoCommandAPIMockRecorder) Info(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockInfoCommandAPI)(nil).Info), arg0)
}
