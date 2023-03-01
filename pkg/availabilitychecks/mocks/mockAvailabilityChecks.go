// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/availabilitychecks (interfaces: AvailabilityChecker)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAvailabilityChecker is a mock of AvailabilityChecker interface.
type MockAvailabilityChecker struct {
	ctrl     *gomock.Controller
	recorder *MockAvailabilityCheckerMockRecorder
}

// MockAvailabilityCheckerMockRecorder is the mock recorder for MockAvailabilityChecker.
type MockAvailabilityCheckerMockRecorder struct {
	mock *MockAvailabilityChecker
}

// NewMockAvailabilityChecker creates a new mock instance.
func NewMockAvailabilityChecker(ctrl *gomock.Controller) *MockAvailabilityChecker {
	mock := &MockAvailabilityChecker{ctrl: ctrl}
	mock.recorder = &MockAvailabilityCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvailabilityChecker) EXPECT() *MockAvailabilityCheckerMockRecorder {
	return m.recorder
}

// AvailabilityCheck mocks base method.
func (m *MockAvailabilityChecker) AvailabilityCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilityCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// AvailabilityCheck indicates an expected call of AvailabilityCheck.
func (mr *MockAvailabilityCheckerMockRecorder) AvailabilityCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilityCheck", reflect.TypeOf((*MockAvailabilityChecker)(nil).AvailabilityCheck))
}
