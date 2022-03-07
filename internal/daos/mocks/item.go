// Code generated by MockGen. DO NOT EDIT.
// Source: item.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/iamnande/cardmod/internal/models"
)

// MockItemDAO is a mock of ItemDAO interface.
type MockItemDAO struct {
	ctrl     *gomock.Controller
	recorder *MockItemDAOMockRecorder
}

// MockItemDAOMockRecorder is the mock recorder for MockItemDAO.
type MockItemDAOMockRecorder struct {
	mock *MockItemDAO
}

// NewMockItemDAO creates a new mock instance.
func NewMockItemDAO(ctrl *gomock.Controller) *MockItemDAO {
	mock := &MockItemDAO{ctrl: ctrl}
	mock.recorder = &MockItemDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemDAO) EXPECT() *MockItemDAOMockRecorder {
	return m.recorder
}

// GetItem mocks base method.
func (m *MockItemDAO) GetItem(name string) (models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", name)
	ret0, _ := ret[0].(models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockItemDAOMockRecorder) GetItem(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockItemDAO)(nil).GetItem), name)
}

// ListItems mocks base method.
func (m *MockItemDAO) ListItems() []models.Item {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListItems")
	ret0, _ := ret[0].([]models.Item)
	return ret0
}

// ListItems indicates an expected call of ListItems.
func (mr *MockItemDAOMockRecorder) ListItems() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItems", reflect.TypeOf((*MockItemDAO)(nil).ListItems))
}
