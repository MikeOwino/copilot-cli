// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/cli/deploy/worker.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	deploy "github.com/aws/copilot-cli/internal/pkg/deploy"
	gomock "github.com/golang/mock/gomock"
)

// MocksnsTopicsLister is a mock of snsTopicsLister interface.
type MocksnsTopicsLister struct {
	ctrl     *gomock.Controller
	recorder *MocksnsTopicsListerMockRecorder
}

// MocksnsTopicsListerMockRecorder is the mock recorder for MocksnsTopicsLister.
type MocksnsTopicsListerMockRecorder struct {
	mock *MocksnsTopicsLister
}

// NewMocksnsTopicsLister creates a new mock instance.
func NewMocksnsTopicsLister(ctrl *gomock.Controller) *MocksnsTopicsLister {
	mock := &MocksnsTopicsLister{ctrl: ctrl}
	mock.recorder = &MocksnsTopicsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksnsTopicsLister) EXPECT() *MocksnsTopicsListerMockRecorder {
	return m.recorder
}

// ListSNSTopics mocks base method.
func (m *MocksnsTopicsLister) ListSNSTopics(appName, envName string) ([]deploy.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSNSTopics", appName, envName)
	ret0, _ := ret[0].([]deploy.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSNSTopics indicates an expected call of ListSNSTopics.
func (mr *MocksnsTopicsListerMockRecorder) ListSNSTopics(appName, envName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSNSTopics", reflect.TypeOf((*MocksnsTopicsLister)(nil).ListSNSTopics), appName, envName)
}