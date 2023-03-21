// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"

	entity "github.com/demig00d/order-service/internal/entity"
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
func (m *MockOrder) GetById(arg0 context.Context, arg1 string) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockOrderMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockOrder)(nil).GetById), arg0, arg1)
}

// Save mocks base method.
func (m *MockOrder) Save(arg0 context.Context, arg1 entity.Order) error {
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

// MockOrderRepo is a mock of OrderRepo interface.
type MockOrderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepoMockRecorder
}

// MockOrderRepoMockRecorder is the mock recorder for MockOrderRepo.
type MockOrderRepoMockRecorder struct {
	mock *MockOrderRepo
}

// NewMockOrderRepo creates a new mock instance.
func NewMockOrderRepo(ctrl *gomock.Controller) *MockOrderRepo {
	mock := &MockOrderRepo{ctrl: ctrl}
	mock.recorder = &MockOrderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepo) EXPECT() *MockOrderRepoMockRecorder {
	return m.recorder
}

// SelectById mocks base method.
func (m *MockOrderRepo) SelectById(arg0 context.Context, arg1 string) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", arg0, arg1)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectById indicates an expected call of SelectById.
func (mr *MockOrderRepoMockRecorder) SelectById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockOrderRepo)(nil).SelectById), arg0, arg1)
}

// Store mocks base method.
func (m *MockOrderRepo) Store(arg0 context.Context, arg1 entity.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockOrderRepoMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockOrderRepo)(nil).Store), arg0, arg1)
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
func (m *MockOrderWebAPI) Translate(arg0 entity.Order) (entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", arg0)
	ret0, _ := ret[0].(entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Translate indicates an expected call of Translate.
func (mr *MockOrderWebAPIMockRecorder) Translate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockOrderWebAPI)(nil).Translate), arg0)
}
