// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go
//
// Generated by this command:
//
//	mockgen -source=auth.go -destination=auth_mock.go -package=auth
//
// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	jwt "github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthentication is a mock of Authentication interface.
type MockAuthentication struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationMockRecorder
}

// MockAuthenticationMockRecorder is the mock recorder for MockAuthentication.
type MockAuthenticationMockRecorder struct {
	mock *MockAuthentication
}

// NewMockAuthentication creates a new mock instance.
func NewMockAuthentication(ctrl *gomock.Controller) *MockAuthentication {
	mock := &MockAuthentication{ctrl: ctrl}
	mock.recorder = &MockAuthenticationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthentication) EXPECT() *MockAuthenticationMockRecorder {
	return m.recorder
}

// GenerateAuthToken mocks base method.
func (m *MockAuthentication) GenerateAuthToken(claims jwt.RegisteredClaims) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAuthToken", claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAuthToken indicates an expected call of GenerateAuthToken.
func (mr *MockAuthenticationMockRecorder) GenerateAuthToken(claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAuthToken", reflect.TypeOf((*MockAuthentication)(nil).GenerateAuthToken), claims)
}

// ValidateToken mocks base method.
func (m *MockAuthentication) ValidateToken(token string) (jwt.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", token)
	ret0, _ := ret[0].(jwt.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockAuthenticationMockRecorder) ValidateToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockAuthentication)(nil).ValidateToken), token)
}
