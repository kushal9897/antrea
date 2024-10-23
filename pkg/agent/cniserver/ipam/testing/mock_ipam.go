// Copyright 2024 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/pkg/agent/cniserver/ipam (interfaces: IPAMDriver)
//
// Generated by this command:
//
//	mockgen -copyright_file hack/boilerplate/license_header.raw.txt -destination pkg/agent/cniserver/ipam/testing/mock_ipam.go -package testing antrea.io/antrea/pkg/agent/cniserver/ipam IPAMDriver
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	ipam "antrea.io/antrea/pkg/agent/cniserver/ipam"
	types "antrea.io/antrea/pkg/agent/cniserver/types"
	invoke "github.com/containernetworking/cni/pkg/invoke"
	gomock "go.uber.org/mock/gomock"
)

// MockIPAMDriver is a mock of IPAMDriver interface.
type MockIPAMDriver struct {
	ctrl     *gomock.Controller
	recorder *MockIPAMDriverMockRecorder
	isgomock struct{}
}

// MockIPAMDriverMockRecorder is the mock recorder for MockIPAMDriver.
type MockIPAMDriverMockRecorder struct {
	mock *MockIPAMDriver
}

// NewMockIPAMDriver creates a new mock instance.
func NewMockIPAMDriver(ctrl *gomock.Controller) *MockIPAMDriver {
	mock := &MockIPAMDriver{ctrl: ctrl}
	mock.recorder = &MockIPAMDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPAMDriver) EXPECT() *MockIPAMDriverMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIPAMDriver) Add(args *invoke.Args, k8sArgs *types.K8sArgs, networkConfig []byte) (bool, *ipam.IPAMResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", args, k8sArgs, networkConfig)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*ipam.IPAMResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Add indicates an expected call of Add.
func (mr *MockIPAMDriverMockRecorder) Add(args, k8sArgs, networkConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIPAMDriver)(nil).Add), args, k8sArgs, networkConfig)
}

// Check mocks base method.
func (m *MockIPAMDriver) Check(args *invoke.Args, k8sArgs *types.K8sArgs, networkConfig []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", args, k8sArgs, networkConfig)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockIPAMDriverMockRecorder) Check(args, k8sArgs, networkConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockIPAMDriver)(nil).Check), args, k8sArgs, networkConfig)
}

// Del mocks base method.
func (m *MockIPAMDriver) Del(args *invoke.Args, k8sArgs *types.K8sArgs, networkConfig []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", args, k8sArgs, networkConfig)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Del indicates an expected call of Del.
func (mr *MockIPAMDriverMockRecorder) Del(args, k8sArgs, networkConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockIPAMDriver)(nil).Del), args, k8sArgs, networkConfig)
}
