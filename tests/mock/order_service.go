// Code generated by MockGen. DO NOT EDIT.
// Source: api/service/order_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	request "github.com/mariojuzar/kitara-store/api/entity/rest-web/request"
	response "github.com/mariojuzar/kitara-store/api/entity/rest-web/response"
	reflect "reflect"
)

// MockOrderService is a mock of OrderService interface
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// LockOrder mocks base method
func (m *MockOrderService) LockOrder(request request.LockOrderRequest) (response.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockOrder", request)
	ret0, _ := ret[0].(response.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LockOrder indicates an expected call of LockOrder
func (mr *MockOrderServiceMockRecorder) LockOrder(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockOrder", reflect.TypeOf((*MockOrderService)(nil).LockOrder), request)
}