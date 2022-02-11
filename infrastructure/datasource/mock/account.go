// Code generated by MockGen. DO NOT EDIT.
// Source: ./infrastructure/datasource/account.go

// Package mock_datasource is a generated GoMock package.
package mock_datasource

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	inframodel "github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

// MockAccountDatasource is a mock of AccountDatasource interface.
type MockAccountDatasource struct {
	ctrl     *gomock.Controller
	recorder *MockAccountDatasourceMockRecorder
}

// MockAccountDatasourceMockRecorder is the mock recorder for MockAccountDatasource.
type MockAccountDatasourceMockRecorder struct {
	mock *MockAccountDatasource
}

// NewMockAccountDatasource creates a new mock instance.
func NewMockAccountDatasource(ctrl *gomock.Controller) *MockAccountDatasource {
	mock := &MockAccountDatasource{ctrl: ctrl}
	mock.recorder = &MockAccountDatasourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountDatasource) EXPECT() *MockAccountDatasourceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountDatasource) Create(ctx context.Context, accountModel *inframodel.AccountModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, accountModel)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAccountDatasourceMockRecorder) Create(ctx, accountModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountDatasource)(nil).Create), ctx, accountModel)
}

// GetById mocks base method.
func (m *MockAccountDatasource) GetById(ctx context.Context, id string) (*inframodel.AccountModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*inframodel.AccountModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockAccountDatasourceMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockAccountDatasource)(nil).GetById), ctx, id)
}
