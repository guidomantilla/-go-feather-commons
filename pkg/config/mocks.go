// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sethvargo/go-envconfig (interfaces: Lookuper)

// Package config is a generated GoMock package.
package config

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLookuper is a mock of Lookuper interface.
type MockLookuper struct {
	ctrl     *gomock.Controller
	recorder *MockLookuperMockRecorder
}

// MockLookuperMockRecorder is the mock recorder for MockLookuper.
type MockLookuperMockRecorder struct {
	mock *MockLookuper
}

// NewMockLookuper creates a new mock instance.
func NewMockLookuper(ctrl *gomock.Controller) *MockLookuper {
	mock := &MockLookuper{ctrl: ctrl}
	mock.recorder = &MockLookuperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLookuper) EXPECT() *MockLookuperMockRecorder {
	return m.recorder
}

// Lookup mocks base method.
func (m *MockLookuper) Lookup(arg0 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup.
func (mr *MockLookuperMockRecorder) Lookup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockLookuper)(nil).Lookup), arg0)
}
