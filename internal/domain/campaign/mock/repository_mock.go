// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/campaign/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/campaign/repository.go -destination=internal/domain/campaign/mock/repository_mock.go
//
// Package mock_campaign is a generated GoMock package.
package mock_campaign

import (
	campaign "email-dispatch-gateway/internal/domain/campaign"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockRepository) GetByID(id string) (*campaign.Campaign, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*campaign.Campaign)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRepositoryMockRecorder) GetByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), id)
}

// Save mocks base method.
func (m *MockRepository) Save(campaign *campaign.Campaign) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", campaign)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRepositoryMockRecorder) Save(campaign any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), campaign)
}
