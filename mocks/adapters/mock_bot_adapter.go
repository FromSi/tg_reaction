// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/tg_reaction/internal/adapters (interfaces: BotAdapter)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/adapters/mock_bot_adapter.go -package=adapters_mocks github.com/fromsi/tg_reaction/internal/adapters BotAdapter
//

// Package adapters_mocks is a generated GoMock package.
package adapters_mocks

import (
	reflect "reflect"

	json "github.com/fromsi/tg_reaction/pkg/json"
	gomock "go.uber.org/mock/gomock"
)

// MockBotAdapter is a mock of BotAdapter interface.
type MockBotAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockBotAdapterMockRecorder
	isgomock struct{}
}

// MockBotAdapterMockRecorder is the mock recorder for MockBotAdapter.
type MockBotAdapterMockRecorder struct {
	mock *MockBotAdapter
}

// NewMockBotAdapter creates a new mock instance.
func NewMockBotAdapter(ctrl *gomock.Controller) *MockBotAdapter {
	mock := &MockBotAdapter{ctrl: ctrl}
	mock.recorder = &MockBotAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBotAdapter) EXPECT() *MockBotAdapterMockRecorder {
	return m.recorder
}

// SetMessageReaction mocks base method.
func (m *MockBotAdapter) SetMessageReaction(chatId int64, messageId int, reaction json.Reaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMessageReaction", chatId, messageId, reaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMessageReaction indicates an expected call of SetMessageReaction.
func (mr *MockBotAdapterMockRecorder) SetMessageReaction(chatId, messageId, reaction any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMessageReaction", reflect.TypeOf((*MockBotAdapter)(nil).SetMessageReaction), chatId, messageId, reaction)
}
