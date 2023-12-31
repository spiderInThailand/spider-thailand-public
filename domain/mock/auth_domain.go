// Code generated by MockGen. DO NOT EDIT.
// Source: auth_domain.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"
	model "spider-go/model"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorities is a mock of Authorities interface.
type MockAuthorities struct {
	ctrl     *gomock.Controller
	recorder *MockAuthoritiesMockRecorder
}

// MockAuthoritiesMockRecorder is the mock recorder for MockAuthorities.
type MockAuthoritiesMockRecorder struct {
	mock *MockAuthorities
}

// NewMockAuthorities creates a new mock instance.
func NewMockAuthorities(ctrl *gomock.Controller) *MockAuthorities {
	mock := &MockAuthorities{ctrl: ctrl}
	mock.recorder = &MockAuthoritiesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorities) EXPECT() *MockAuthoritiesMockRecorder {
	return m.recorder
}

// CreateAccout mocks base method.
func (m *MockAuthorities) CreateAccout(ctx context.Context, data model.Account, password, confirmPassowrd string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccout", ctx, data, password, confirmPassowrd)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccout indicates an expected call of CreateAccout.
func (mr *MockAuthoritiesMockRecorder) CreateAccout(ctx, data, password, confirmPassowrd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccout", reflect.TypeOf((*MockAuthorities)(nil).CreateAccout), ctx, data, password, confirmPassowrd)
}

// Login mocks base method.
func (m *MockAuthorities) Login(ctx context.Context, username, password string) (*model.Account, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, username, password)
	ret0, _ := ret[0].(*model.Account)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockAuthoritiesMockRecorder) Login(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthorities)(nil).Login), ctx, username, password)
}
