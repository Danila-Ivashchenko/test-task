// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/ports/storage/user.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	model "go-kafka/internal/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserStorage is a mock of UserStorage interface.
type MockUserStorage struct {
	ctrl     *gomock.Controller
	recorder *MockUserStorageMockRecorder
}

// MockUserStorageMockRecorder is the mock recorder for MockUserStorage.
type MockUserStorageMockRecorder struct {
	mock *MockUserStorage
}

// NewMockUserStorage creates a new mock instance.
func NewMockUserStorage(ctrl *gomock.Controller) *MockUserStorage {
	mock := &MockUserStorage{ctrl: ctrl}
	mock.recorder = &MockUserStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStorage) EXPECT() *MockUserStorageMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserStorage) AddUser(arg0 context.Context, arg1 *model.FullUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserStorageMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserStorage)(nil).AddUser), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockUserStorage) GetUsers(arg0 context.Context, arg1 *model.GetUsersDTO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUsers", arg0, arg1)
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserStorageMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserStorage)(nil).GetUsers), arg0, arg1)
}
