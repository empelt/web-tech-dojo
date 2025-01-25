// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=mock/interface.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGenaiClient is a mock of GenaiClient interface.
type MockGenaiClient struct {
	ctrl     *gomock.Controller
	recorder *MockGenaiClientMockRecorder
}

// MockGenaiClientMockRecorder is the mock recorder for MockGenaiClient.
type MockGenaiClientMockRecorder struct {
	mock *MockGenaiClient
}

// NewMockGenaiClient creates a new mock instance.
func NewMockGenaiClient(ctrl *gomock.Controller) *MockGenaiClient {
	mock := &MockGenaiClient{ctrl: ctrl}
	mock.recorder = &MockGenaiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenaiClient) EXPECT() *MockGenaiClientMockRecorder {
	return m.recorder
}

// GenerateContentFromText mocks base method.
func (m *MockGenaiClient) GenerateContentFromText(ctx context.Context, message string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateContentFromText", ctx, message)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateContentFromText indicates an expected call of GenerateContentFromText.
func (mr *MockGenaiClientMockRecorder) GenerateContentFromText(ctx, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateContentFromText", reflect.TypeOf((*MockGenaiClient)(nil).GenerateContentFromText), ctx, message)
}
