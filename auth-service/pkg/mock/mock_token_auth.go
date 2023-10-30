// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/token/token.go

// Package mock is a generated GoMock package.
package mock

import (
	token "auth-service/pkg/service/token"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTokenAuth is a mock of TokenAuth interface.
type MockTokenAuth struct {
	ctrl     *gomock.Controller
	recorder *MockTokenAuthMockRecorder
}

// MockTokenAuthMockRecorder is the mock recorder for MockTokenAuth.
type MockTokenAuthMockRecorder struct {
	mock *MockTokenAuth
}

// NewMockTokenAuth creates a new mock instance.
func NewMockTokenAuth(ctrl *gomock.Controller) *MockTokenAuth {
	mock := &MockTokenAuth{ctrl: ctrl}
	mock.recorder = &MockTokenAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenAuth) EXPECT() *MockTokenAuthMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockTokenAuth) GenerateToken(req token.Payload) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockTokenAuthMockRecorder) GenerateToken(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockTokenAuth)(nil).GenerateToken), req)
}

// VerifyToken mocks base method.
func (m *MockTokenAuth) VerifyToken(tokenString string) (token.Payload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", tokenString)
	ret0, _ := ret[0].(token.Payload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockTokenAuthMockRecorder) VerifyToken(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockTokenAuth)(nil).VerifyToken), tokenString)
}
