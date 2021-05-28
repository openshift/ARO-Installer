// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/refreshable (interfaces: Authorizer)

// Package mock_refreshable is a generated GoMock package.
package mock_refreshable

import (
	context "context"
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
)

// MockAuthorizer is a mock of Authorizer interface.
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer.
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance.
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// OAuthToken mocks base method.
func (m *MockAuthorizer) OAuthToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OAuthToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// OAuthToken indicates an expected call of OAuthToken.
func (mr *MockAuthorizerMockRecorder) OAuthToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuthToken", reflect.TypeOf((*MockAuthorizer)(nil).OAuthToken))
}

// RefreshWithContext mocks base method.
func (m *MockAuthorizer) RefreshWithContext(arg0 context.Context, arg1 *logrus.Entry) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshWithContext", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshWithContext indicates an expected call of RefreshWithContext.
func (mr *MockAuthorizerMockRecorder) RefreshWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshWithContext", reflect.TypeOf((*MockAuthorizer)(nil).RefreshWithContext), arg0, arg1)
}

// WithAuthorization mocks base method.
func (m *MockAuthorizer) WithAuthorization() autorest.PrepareDecorator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithAuthorization")
	ret0, _ := ret[0].(autorest.PrepareDecorator)
	return ret0
}

// WithAuthorization indicates an expected call of WithAuthorization.
func (mr *MockAuthorizerMockRecorder) WithAuthorization() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithAuthorization", reflect.TypeOf((*MockAuthorizer)(nil).WithAuthorization))
}
