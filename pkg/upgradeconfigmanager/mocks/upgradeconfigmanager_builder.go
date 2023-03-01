// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/upgradeconfigmanager (interfaces: UpgradeConfigManagerBuilder)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	upgradeconfigmanager "github.com/openshift/managed-upgrade-operator/pkg/upgradeconfigmanager"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockUpgradeConfigManagerBuilder is a mock of UpgradeConfigManagerBuilder interface.
type MockUpgradeConfigManagerBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeConfigManagerBuilderMockRecorder
}

// MockUpgradeConfigManagerBuilderMockRecorder is the mock recorder for MockUpgradeConfigManagerBuilder.
type MockUpgradeConfigManagerBuilderMockRecorder struct {
	mock *MockUpgradeConfigManagerBuilder
}

// NewMockUpgradeConfigManagerBuilder creates a new mock instance.
func NewMockUpgradeConfigManagerBuilder(ctrl *gomock.Controller) *MockUpgradeConfigManagerBuilder {
	mock := &MockUpgradeConfigManagerBuilder{ctrl: ctrl}
	mock.recorder = &MockUpgradeConfigManagerBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeConfigManagerBuilder) EXPECT() *MockUpgradeConfigManagerBuilderMockRecorder {
	return m.recorder
}

// NewManager mocks base method.
func (m *MockUpgradeConfigManagerBuilder) NewManager(arg0 client.Client) (upgradeconfigmanager.UpgradeConfigManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewManager", arg0)
	ret0, _ := ret[0].(upgradeconfigmanager.UpgradeConfigManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewManager indicates an expected call of NewManager.
func (mr *MockUpgradeConfigManagerBuilderMockRecorder) NewManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewManager", reflect.TypeOf((*MockUpgradeConfigManagerBuilder)(nil).NewManager), arg0)
}
