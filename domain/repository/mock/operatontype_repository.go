// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/repository/operationtype_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	entity "github.com/nelsonlpco/transactions/domain/entity"
)

// MockOperationTypeRepository is a mock of OperationTypeRepository interface.
type MockOperationTypeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOperationTypeRepositoryMockRecorder
}

// MockOperationTypeRepositoryMockRecorder is the mock recorder for MockOperationTypeRepository.
type MockOperationTypeRepositoryMockRecorder struct {
	mock *MockOperationTypeRepository
}

// NewMockOperationTypeRepository creates a new mock instance.
func NewMockOperationTypeRepository(ctrl *gomock.Controller) *MockOperationTypeRepository {
	mock := &MockOperationTypeRepository{ctrl: ctrl}
	mock.recorder = &MockOperationTypeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationTypeRepository) EXPECT() *MockOperationTypeRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOperationTypeRepository) Create(ctx context.Context, operationType *entity.OperationType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, operationType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOperationTypeRepositoryMockRecorder) Create(ctx, operationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOperationTypeRepository)(nil).Create), ctx, operationType)
}

// GetById mocks base method.
func (m *MockOperationTypeRepository) GetById(ctx context.Context, id uuid.UUID) (*entity.OperationType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*entity.OperationType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockOperationTypeRepositoryMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockOperationTypeRepository)(nil).GetById), ctx, id)
}
