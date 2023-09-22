// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/ports/enricher/enricher.go

// Package mock_enricher is a generated GoMock package.
package mock_enricher

import (
	context "context"
	model "go-kafka/internal/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEnricher is a mock of Enricher interface.
type MockEnricher struct {
	ctrl     *gomock.Controller
	recorder *MockEnricherMockRecorder
}

// MockEnricherMockRecorder is the mock recorder for MockEnricher.
type MockEnricherMockRecorder struct {
	mock *MockEnricher
}

// NewMockEnricher creates a new mock instance.
func NewMockEnricher(ctrl *gomock.Controller) *MockEnricher {
	mock := &MockEnricher{ctrl: ctrl}
	mock.recorder = &MockEnricherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnricher) EXPECT() *MockEnricherMockRecorder {
	return m.recorder
}

// GetAge mocks base method.
func (m *MockEnricher) GetAge(arg0 context.Context, arg1 string) (*model.AgeDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAge", arg0, arg1)
	ret0, _ := ret[0].(*model.AgeDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAge indicates an expected call of GetAge.
func (mr *MockEnricherMockRecorder) GetAge(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAge", reflect.TypeOf((*MockEnricher)(nil).GetAge), arg0, arg1)
}

// GetGender mocks base method.
func (m *MockEnricher) GetGender(arg0 context.Context, arg1 string) (*model.GenderDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGender", arg0, arg1)
	ret0, _ := ret[0].(*model.GenderDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGender indicates an expected call of GetGender.
func (mr *MockEnricherMockRecorder) GetGender(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGender", reflect.TypeOf((*MockEnricher)(nil).GetGender), arg0, arg1)
}

// GetNationalities mocks base method.
func (m *MockEnricher) GetNationalities(arg0 context.Context, arg1 string) (*model.NationalitiesDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNationalities", arg0, arg1)
	ret0, _ := ret[0].(*model.NationalitiesDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNationalities indicates an expected call of GetNationalities.
func (mr *MockEnricherMockRecorder) GetNationalities(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNationalities", reflect.TypeOf((*MockEnricher)(nil).GetNationalities), arg0, arg1)
}
