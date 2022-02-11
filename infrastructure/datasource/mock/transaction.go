// Code generated by MockGen. DO NOT EDIT.
// Source: ./infrastructure/datasource/transaction.go

// Package mock_datasource is a generated GoMock package.
package mock_datasource

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	inframodel "github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

// MockTransactionDatasource is a mock of TransactionDatasource interface.
type MockTransactionDatasource struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionDatasourceMockRecorder
}

// MockTransactionDatasourceMockRecorder is the mock recorder for MockTransactionDatasource.
type MockTransactionDatasourceMockRecorder struct {
	mock *MockTransactionDatasource
}

// NewMockTransactionDatasource creates a new mock instance.
func NewMockTransactionDatasource(ctrl *gomock.Controller) *MockTransactionDatasource {
	mock := &MockTransactionDatasource{ctrl: ctrl}
	mock.recorder = &MockTransactionDatasourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionDatasource) EXPECT() *MockTransactionDatasourceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTransactionDatasource) Create(ctx context.Context, transaction *inframodel.TransactionModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTransactionDatasourceMockRecorder) Create(ctx, transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransactionDatasource)(nil).Create), ctx, transaction)
}

// GetTransactionsByAccountId mocks base method.
func (m *MockTransactionDatasource) GetTransactionsByAccountId(ctx context.Context, accountId string) ([]*inframodel.TransactionModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsByAccountId", ctx, accountId)
	ret0, _ := ret[0].([]*inframodel.TransactionModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsByAccountId indicates an expected call of GetTransactionsByAccountId.
func (mr *MockTransactionDatasourceMockRecorder) GetTransactionsByAccountId(ctx, accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsByAccountId", reflect.TypeOf((*MockTransactionDatasource)(nil).GetTransactionsByAccountId), ctx, accountId)
}
