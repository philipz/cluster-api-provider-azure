/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-03-01/compute/computeapi (interfaces: DisksClientAPI)

// Package mock_disks is a generated GoMock package.
package mock_disks

import (
	context "context"
	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-03-01/compute"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDisksClientAPI is a mock of DisksClientAPI interface
type MockDisksClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDisksClientAPIMockRecorder
}

// MockDisksClientAPIMockRecorder is the mock recorder for MockDisksClientAPI
type MockDisksClientAPIMockRecorder struct {
	mock *MockDisksClientAPI
}

// NewMockDisksClientAPI creates a new mock instance
func NewMockDisksClientAPI(ctrl *gomock.Controller) *MockDisksClientAPI {
	mock := &MockDisksClientAPI{ctrl: ctrl}
	mock.recorder = &MockDisksClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDisksClientAPI) EXPECT() *MockDisksClientAPIMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockDisksClientAPI) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 compute.Disk) (compute.DisksCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.DisksCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockDisksClientAPIMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockDisksClientAPI)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockDisksClientAPI) Delete(arg0 context.Context, arg1, arg2 string) (compute.DisksDeleteFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.DisksDeleteFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockDisksClientAPIMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDisksClientAPI)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockDisksClientAPI) Get(arg0 context.Context, arg1, arg2 string) (compute.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDisksClientAPIMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDisksClientAPI)(nil).Get), arg0, arg1, arg2)
}

// GrantAccess mocks base method
func (m *MockDisksClientAPI) GrantAccess(arg0 context.Context, arg1, arg2 string, arg3 compute.GrantAccessData) (compute.DisksGrantAccessFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantAccess", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.DisksGrantAccessFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GrantAccess indicates an expected call of GrantAccess
func (mr *MockDisksClientAPIMockRecorder) GrantAccess(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantAccess", reflect.TypeOf((*MockDisksClientAPI)(nil).GrantAccess), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockDisksClientAPI) List(arg0 context.Context) (compute.DiskListPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(compute.DiskListPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockDisksClientAPIMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDisksClientAPI)(nil).List), arg0)
}

// ListByResourceGroup mocks base method
func (m *MockDisksClientAPI) ListByResourceGroup(arg0 context.Context, arg1 string) (compute.DiskListPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByResourceGroup", arg0, arg1)
	ret0, _ := ret[0].(compute.DiskListPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByResourceGroup indicates an expected call of ListByResourceGroup
func (mr *MockDisksClientAPIMockRecorder) ListByResourceGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByResourceGroup", reflect.TypeOf((*MockDisksClientAPI)(nil).ListByResourceGroup), arg0, arg1)
}

// RevokeAccess mocks base method
func (m *MockDisksClientAPI) RevokeAccess(arg0 context.Context, arg1, arg2 string) (compute.DisksRevokeAccessFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.DisksRevokeAccessFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAccess indicates an expected call of RevokeAccess
func (mr *MockDisksClientAPIMockRecorder) RevokeAccess(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAccess", reflect.TypeOf((*MockDisksClientAPI)(nil).RevokeAccess), arg0, arg1, arg2)
}

// Update mocks base method
func (m *MockDisksClientAPI) Update(arg0 context.Context, arg1, arg2 string, arg3 compute.DiskUpdate) (compute.DisksUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.DisksUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockDisksClientAPIMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDisksClientAPI)(nil).Update), arg0, arg1, arg2, arg3)
}
