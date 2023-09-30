// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/campaign/service_interface.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/campaign/service_interface.go -destination=internal/domain/campaign/mock/service_mock.go
//
// Package mock_campaign is a generated GoMock package.
package mock_campaign

import (
	contract "email-dispatch-gateway/internal/contract"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockServiceInterface is a mock of ServiceInterface interface.
type MockServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInterfaceMockRecorder
}

// MockServiceInterfaceMockRecorder is the mock recorder for MockServiceInterface.
type MockServiceInterfaceMockRecorder struct {
	mock *MockServiceInterface
}

// NewMockServiceInterface creates a new mock instance.
func NewMockServiceInterface(ctrl *gomock.Controller) *MockServiceInterface {
	mock := &MockServiceInterface{ctrl: ctrl}
	mock.recorder = &MockServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInterface) EXPECT() *MockServiceInterfaceMockRecorder {
	return m.recorder
}

// CancelByID mocks base method.
func (m *MockServiceInterface) CancelByID(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelByID indicates an expected call of CancelByID.
func (mr *MockServiceInterfaceMockRecorder) CancelByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelByID", reflect.TypeOf((*MockServiceInterface)(nil).CancelByID), id)
}

// Create mocks base method.
func (m *MockServiceInterface) Create(dto contract.NewCampaignDTO) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", dto)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceInterfaceMockRecorder) Create(dto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServiceInterface)(nil).Create), dto)
}

// GetByID mocks base method.
func (m *MockServiceInterface) GetByID(id string) (contract.CampaignResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(contract.CampaignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockServiceInterfaceMockRecorder) GetByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockServiceInterface)(nil).GetByID), id)
}
