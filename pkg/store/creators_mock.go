// Code generated by MockGen. DO NOT EDIT.
// Source: creators.go

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/nooble/task/audio-short-api/pkg/api/model"
)

// MockCreatorsStore is a mock of CreatorsStore interface.
type MockCreatorsStore struct {
	ctrl     *gomock.Controller
	recorder *MockCreatorsStoreMockRecorder
}

// MockCreatorsStoreMockRecorder is the mock recorder for MockCreatorsStore.
type MockCreatorsStoreMockRecorder struct {
	mock *MockCreatorsStore
}

// NewMockCreatorsStore creates a new mock instance.
func NewMockCreatorsStore(ctrl *gomock.Controller) *MockCreatorsStore {
	mock := &MockCreatorsStore{ctrl: ctrl}
	mock.recorder = &MockCreatorsStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreatorsStore) EXPECT() *MockCreatorsStoreMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockCreatorsStore) GetAll(ctx context.Context, page, limit uint16) ([]*model.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, page, limit)
	ret0, _ := ret[0].([]*model.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCreatorsStoreMockRecorder) GetAll(ctx, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCreatorsStore)(nil).GetAll), ctx, page, limit)
}
