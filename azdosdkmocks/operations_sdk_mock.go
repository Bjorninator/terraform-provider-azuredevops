// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/azure-devops-go-api/azuredevops/operations (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	operations "github.com/microsoft/azure-devops-go-api/azuredevops/operations"
	reflect "reflect"
)

// MockOperationsClient is a mock of Client interface
type MockOperationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockOperationsClientMockRecorder
}

// MockOperationsClientMockRecorder is the mock recorder for MockOperationsClient
type MockOperationsClientMockRecorder struct {
	mock *MockOperationsClient
}

// NewMockOperationsClient creates a new mock instance
func NewMockOperationsClient(ctrl *gomock.Controller) *MockOperationsClient {
	mock := &MockOperationsClient{ctrl: ctrl}
	mock.recorder = &MockOperationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOperationsClient) EXPECT() *MockOperationsClientMockRecorder {
	return m.recorder
}

// GetOperation mocks base method
func (m *MockOperationsClient) GetOperation(arg0 context.Context, arg1 operations.GetOperationArgs) (*operations.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation", arg0, arg1)
	ret0, _ := ret[0].(*operations.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperation indicates an expected call of GetOperation
func (mr *MockOperationsClientMockRecorder) GetOperation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockOperationsClient)(nil).GetOperation), arg0, arg1)
}