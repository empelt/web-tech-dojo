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

// MockChatService is a mock of ChatService interface.
type MockChatService struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceMockRecorder
}

// MockChatServiceMockRecorder is the mock recorder for MockChatService.
type MockChatServiceMockRecorder struct {
	mock *MockChatService
}

// NewMockChatService creates a new mock instance.
func NewMockChatService(ctrl *gomock.Controller) *MockChatService {
	mock := &MockChatService{ctrl: ctrl}
	mock.recorder = &MockChatServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatService) EXPECT() *MockChatServiceMockRecorder {
	return m.recorder
}

// PostChatMessage mocks base method.
func (m *MockChatService) PostChatMessage(ctx context.Context, message string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostChatMessage", ctx, message)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostChatMessage indicates an expected call of PostChatMessage.
func (mr *MockChatServiceMockRecorder) PostChatMessage(ctx, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostChatMessage", reflect.TypeOf((*MockChatService)(nil).PostChatMessage), ctx, message)
}
