// Code generated by MockGen. DO NOT EDIT.
// Source: limitbreak.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/iamnande/cardmod/internal/models"
)

// MockLimitBreakDAO is a mock of LimitBreakDAO interface.
type MockLimitBreakDAO struct {
	ctrl     *gomock.Controller
	recorder *MockLimitBreakDAOMockRecorder
}

// MockLimitBreakDAOMockRecorder is the mock recorder for MockLimitBreakDAO.
type MockLimitBreakDAOMockRecorder struct {
	mock *MockLimitBreakDAO
}

// NewMockLimitBreakDAO creates a new mock instance.
func NewMockLimitBreakDAO(ctrl *gomock.Controller) *MockLimitBreakDAO {
	mock := &MockLimitBreakDAO{ctrl: ctrl}
	mock.recorder = &MockLimitBreakDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLimitBreakDAO) EXPECT() *MockLimitBreakDAOMockRecorder {
	return m.recorder
}

// GetLimitBreak mocks base method.
func (m *MockLimitBreakDAO) GetLimitBreak(name string) (models.LimitBreak, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLimitBreak", name)
	ret0, _ := ret[0].(models.LimitBreak)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLimitBreak indicates an expected call of GetLimitBreak.
func (mr *MockLimitBreakDAOMockRecorder) GetLimitBreak(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLimitBreak", reflect.TypeOf((*MockLimitBreakDAO)(nil).GetLimitBreak), name)
}

// ListLimitBreaks mocks base method.
func (m *MockLimitBreakDAO) ListLimitBreaks() []models.LimitBreak {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLimitBreaks")
	ret0, _ := ret[0].([]models.LimitBreak)
	return ret0
}

// ListLimitBreaks indicates an expected call of ListLimitBreaks.
func (mr *MockLimitBreakDAOMockRecorder) ListLimitBreaks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLimitBreaks", reflect.TypeOf((*MockLimitBreakDAO)(nil).ListLimitBreaks))
}
