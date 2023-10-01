// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go
//
// Generated by this command:
//
//	mockgen -source=usecase.go -destination=mocks/mock.go
//
// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	dto "persserv/internal/dto"
	myerror "persserv/internal/error-my"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPerson is a mock of Person interface.
type MockPerson struct {
	ctrl     *gomock.Controller
	recorder *MockPersonMockRecorder
}

// MockPersonMockRecorder is the mock recorder for MockPerson.
type MockPersonMockRecorder struct {
	mock *MockPerson
}

// NewMockPerson creates a new mock instance.
func NewMockPerson(ctrl *gomock.Controller) *MockPerson {
	mock := &MockPerson{ctrl: ctrl}
	mock.recorder = &MockPersonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerson) EXPECT() *MockPersonMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPerson) Create(arg0 dto.PersonCreate) (int, *myerror.ErrorFull) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*myerror.ErrorFull)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPersonMockRecorder) Create(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPerson)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockPerson) Delete(personId int) *myerror.ErrorFull {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", personId)
	ret0, _ := ret[0].(*myerror.ErrorFull)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPersonMockRecorder) Delete(personId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPerson)(nil).Delete), personId)
}

// GetAll mocks base method.
func (m *MockPerson) GetAll() ([]dto.Person, *myerror.ErrorFull) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]dto.Person)
	ret1, _ := ret[1].(*myerror.ErrorFull)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPersonMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPerson)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockPerson) GetById(personId int) (dto.Person, *myerror.ErrorFull) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", personId)
	ret0, _ := ret[0].(dto.Person)
	ret1, _ := ret[1].(*myerror.ErrorFull)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockPersonMockRecorder) GetById(personId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPerson)(nil).GetById), personId)
}

// Update mocks base method.
func (m *MockPerson) Update(personId int, person dto.PersonUpdate) (dto.Person, *myerror.ErrorFull) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", personId, person)
	ret0, _ := ret[0].(dto.Person)
	ret1, _ := ret[1].(*myerror.ErrorFull)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPersonMockRecorder) Update(personId, person any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPerson)(nil).Update), personId, person)
}