// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	entity "github.com/demig00d/order-service/internal/repo"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrder is a mock of Order interface.
type MockOrder struct {
	ctrl     *gomock.Controller
	recorder *MockOrderMockRecorder
}

// MockOrderMockRecorder is the mock recorder for MockOrder.
type MockOrderMockRecorder struct {
	mock *MockOrder
}

// NewMockOrder creates a new mock instance.
func NewMockOrder(ctrl *gomock.Controller) *MockOrder {
	mock := &MockOrder{ctrl: ctrl}
	mock.recorder = &MockOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrder) EXPECT() *MockOrderMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockOrder) GetById(arg0 context.Context, arg1 string) (*entity.OrderDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*entity.OrderDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockOrderMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockOrder)(nil).GetById), arg0, arg1)
}

// Save mocks base method.
func (m *MockOrder) Save(arg0 context.Context, arg1 entity.OrderDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockOrderMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOrder)(nil).Save), arg0, arg1)
}

// MockOrderDb is a mock of OrderDb interface.
type MockOrderDb struct {
	ctrl     *gomock.Controller
	recorder *MockOrderDbMockRecorder
}

// MockOrderDbMockRecorder is the mock recorder for MockOrderDb.
type MockOrderDbMockRecorder struct {
	mock *MockOrderDb
}

// NewMockOrderDb creates a new mock instance.
func NewMockOrderDb(ctrl *gomock.Controller) *MockOrderDb {
	mock := &MockOrderDb{ctrl: ctrl}
	mock.recorder = &MockOrderDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderDb) EXPECT() *MockOrderDbMockRecorder {
	return m.recorder
}

// SelectById mocks base method.
func (m *MockOrderDb) SelectById(arg0 context.Context, arg1 string) (*entity.OrderDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", arg0, arg1)
	ret0, _ := ret[0].(*entity.OrderDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectById indicates an expected call of SelectById.
func (mr *MockOrderDbMockRecorder) SelectById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockOrderDb)(nil).SelectById), arg0, arg1)
}

// Store mocks base method.
func (m *MockOrderDb) Store(arg0 context.Context, arg1 entity.OrderDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockOrderDbMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockOrderDb)(nil).Store), arg0, arg1)
}

// MockOrderWebAPI is a mock of OrderWebAPI interface.
type MockOrderWebAPI struct {
	ctrl     *gomock.Controller
	recorder *MockOrderWebAPIMockRecorder
}

// MockOrderWebAPIMockRecorder is the mock recorder for MockOrderWebAPI.
type MockOrderWebAPIMockRecorder struct {
	mock *MockOrderWebAPI
}

// NewMockOrderWebAPI creates a new mock instance.
func NewMockOrderWebAPI(ctrl *gomock.Controller) *MockOrderWebAPI {
	mock := &MockOrderWebAPI{ctrl: ctrl}
	mock.recorder = &MockOrderWebAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderWebAPI) EXPECT() *MockOrderWebAPIMockRecorder {
	return m.recorder
}

// Translate mocks base method.
func (m *MockOrderWebAPI) Translate(arg0 entity.OrderDto) (entity.OrderDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", arg0)
	ret0, _ := ret[0].(entity.OrderDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Translate indicates an expected call of Translate.
func (mr *MockOrderWebAPIMockRecorder) Translate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockOrderWebAPI)(nil).Translate), arg0)
}
