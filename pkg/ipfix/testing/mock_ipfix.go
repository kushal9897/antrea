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
// Source: antrea.io/antrea/pkg/ipfix (interfaces: IPFIXExportingProcess,IPFIXRegistry,IPFIXCollectingProcess,IPFIXAggregationProcess)
//
// Generated by this command:
//
//	mockgen -copyright_file hack/boilerplate/license_header.raw.txt -destination pkg/ipfix/testing/mock_ipfix.go -package testing antrea.io/antrea/pkg/ipfix IPFIXExportingProcess,IPFIXRegistry,IPFIXCollectingProcess,IPFIXAggregationProcess
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"
	time "time"

	entities "github.com/vmware/go-ipfix/pkg/entities"
	intermediate "github.com/vmware/go-ipfix/pkg/intermediate"
	gomock "go.uber.org/mock/gomock"
)

// MockIPFIXExportingProcess is a mock of IPFIXExportingProcess interface.
type MockIPFIXExportingProcess struct {
	ctrl     *gomock.Controller
	recorder *MockIPFIXExportingProcessMockRecorder
	isgomock struct{}
}

// MockIPFIXExportingProcessMockRecorder is the mock recorder for MockIPFIXExportingProcess.
type MockIPFIXExportingProcessMockRecorder struct {
	mock *MockIPFIXExportingProcess
}

// NewMockIPFIXExportingProcess creates a new mock instance.
func NewMockIPFIXExportingProcess(ctrl *gomock.Controller) *MockIPFIXExportingProcess {
	mock := &MockIPFIXExportingProcess{ctrl: ctrl}
	mock.recorder = &MockIPFIXExportingProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPFIXExportingProcess) EXPECT() *MockIPFIXExportingProcessMockRecorder {
	return m.recorder
}

// CloseConnToCollector mocks base method.
func (m *MockIPFIXExportingProcess) CloseConnToCollector() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseConnToCollector")
}

// CloseConnToCollector indicates an expected call of CloseConnToCollector.
func (mr *MockIPFIXExportingProcessMockRecorder) CloseConnToCollector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseConnToCollector", reflect.TypeOf((*MockIPFIXExportingProcess)(nil).CloseConnToCollector))
}

// NewTemplateID mocks base method.
func (m *MockIPFIXExportingProcess) NewTemplateID() uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTemplateID")
	ret0, _ := ret[0].(uint16)
	return ret0
}

// NewTemplateID indicates an expected call of NewTemplateID.
func (mr *MockIPFIXExportingProcessMockRecorder) NewTemplateID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTemplateID", reflect.TypeOf((*MockIPFIXExportingProcess)(nil).NewTemplateID))
}

// SendSet mocks base method.
func (m *MockIPFIXExportingProcess) SendSet(set entities.Set) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSet", set)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendSet indicates an expected call of SendSet.
func (mr *MockIPFIXExportingProcessMockRecorder) SendSet(set any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSet", reflect.TypeOf((*MockIPFIXExportingProcess)(nil).SendSet), set)
}

// MockIPFIXRegistry is a mock of IPFIXRegistry interface.
type MockIPFIXRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockIPFIXRegistryMockRecorder
	isgomock struct{}
}

// MockIPFIXRegistryMockRecorder is the mock recorder for MockIPFIXRegistry.
type MockIPFIXRegistryMockRecorder struct {
	mock *MockIPFIXRegistry
}

// NewMockIPFIXRegistry creates a new mock instance.
func NewMockIPFIXRegistry(ctrl *gomock.Controller) *MockIPFIXRegistry {
	mock := &MockIPFIXRegistry{ctrl: ctrl}
	mock.recorder = &MockIPFIXRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPFIXRegistry) EXPECT() *MockIPFIXRegistryMockRecorder {
	return m.recorder
}

// GetInfoElement mocks base method.
func (m *MockIPFIXRegistry) GetInfoElement(name string, enterpriseID uint32) (*entities.InfoElement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfoElement", name, enterpriseID)
	ret0, _ := ret[0].(*entities.InfoElement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfoElement indicates an expected call of GetInfoElement.
func (mr *MockIPFIXRegistryMockRecorder) GetInfoElement(name, enterpriseID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfoElement", reflect.TypeOf((*MockIPFIXRegistry)(nil).GetInfoElement), name, enterpriseID)
}

// LoadRegistry mocks base method.
func (m *MockIPFIXRegistry) LoadRegistry() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LoadRegistry")
}

// LoadRegistry indicates an expected call of LoadRegistry.
func (mr *MockIPFIXRegistryMockRecorder) LoadRegistry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadRegistry", reflect.TypeOf((*MockIPFIXRegistry)(nil).LoadRegistry))
}

// MockIPFIXCollectingProcess is a mock of IPFIXCollectingProcess interface.
type MockIPFIXCollectingProcess struct {
	ctrl     *gomock.Controller
	recorder *MockIPFIXCollectingProcessMockRecorder
	isgomock struct{}
}

// MockIPFIXCollectingProcessMockRecorder is the mock recorder for MockIPFIXCollectingProcess.
type MockIPFIXCollectingProcessMockRecorder struct {
	mock *MockIPFIXCollectingProcess
}

// NewMockIPFIXCollectingProcess creates a new mock instance.
func NewMockIPFIXCollectingProcess(ctrl *gomock.Controller) *MockIPFIXCollectingProcess {
	mock := &MockIPFIXCollectingProcess{ctrl: ctrl}
	mock.recorder = &MockIPFIXCollectingProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPFIXCollectingProcess) EXPECT() *MockIPFIXCollectingProcessMockRecorder {
	return m.recorder
}

// GetMsgChan mocks base method.
func (m *MockIPFIXCollectingProcess) GetMsgChan() <-chan *entities.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMsgChan")
	ret0, _ := ret[0].(<-chan *entities.Message)
	return ret0
}

// GetMsgChan indicates an expected call of GetMsgChan.
func (mr *MockIPFIXCollectingProcessMockRecorder) GetMsgChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMsgChan", reflect.TypeOf((*MockIPFIXCollectingProcess)(nil).GetMsgChan))
}

// GetNumConnToCollector mocks base method.
func (m *MockIPFIXCollectingProcess) GetNumConnToCollector() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumConnToCollector")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetNumConnToCollector indicates an expected call of GetNumConnToCollector.
func (mr *MockIPFIXCollectingProcessMockRecorder) GetNumConnToCollector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumConnToCollector", reflect.TypeOf((*MockIPFIXCollectingProcess)(nil).GetNumConnToCollector))
}

// GetNumRecordsReceived mocks base method.
func (m *MockIPFIXCollectingProcess) GetNumRecordsReceived() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumRecordsReceived")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetNumRecordsReceived indicates an expected call of GetNumRecordsReceived.
func (mr *MockIPFIXCollectingProcessMockRecorder) GetNumRecordsReceived() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumRecordsReceived", reflect.TypeOf((*MockIPFIXCollectingProcess)(nil).GetNumRecordsReceived))
}

// Start mocks base method.
func (m *MockIPFIXCollectingProcess) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockIPFIXCollectingProcessMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockIPFIXCollectingProcess)(nil).Start))
}

// Stop mocks base method.
func (m *MockIPFIXCollectingProcess) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockIPFIXCollectingProcessMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockIPFIXCollectingProcess)(nil).Stop))
}

// MockIPFIXAggregationProcess is a mock of IPFIXAggregationProcess interface.
type MockIPFIXAggregationProcess struct {
	ctrl     *gomock.Controller
	recorder *MockIPFIXAggregationProcessMockRecorder
	isgomock struct{}
}

// MockIPFIXAggregationProcessMockRecorder is the mock recorder for MockIPFIXAggregationProcess.
type MockIPFIXAggregationProcessMockRecorder struct {
	mock *MockIPFIXAggregationProcess
}

// NewMockIPFIXAggregationProcess creates a new mock instance.
func NewMockIPFIXAggregationProcess(ctrl *gomock.Controller) *MockIPFIXAggregationProcess {
	mock := &MockIPFIXAggregationProcess{ctrl: ctrl}
	mock.recorder = &MockIPFIXAggregationProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPFIXAggregationProcess) EXPECT() *MockIPFIXAggregationProcessMockRecorder {
	return m.recorder
}

// AreCorrelatedFieldsFilled mocks base method.
func (m *MockIPFIXAggregationProcess) AreCorrelatedFieldsFilled(record intermediate.AggregationFlowRecord) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AreCorrelatedFieldsFilled", record)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AreCorrelatedFieldsFilled indicates an expected call of AreCorrelatedFieldsFilled.
func (mr *MockIPFIXAggregationProcessMockRecorder) AreCorrelatedFieldsFilled(record any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AreCorrelatedFieldsFilled", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).AreCorrelatedFieldsFilled), record)
}

// AreExternalFieldsFilled mocks base method.
func (m *MockIPFIXAggregationProcess) AreExternalFieldsFilled(record intermediate.AggregationFlowRecord) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AreExternalFieldsFilled", record)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AreExternalFieldsFilled indicates an expected call of AreExternalFieldsFilled.
func (mr *MockIPFIXAggregationProcessMockRecorder) AreExternalFieldsFilled(record any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AreExternalFieldsFilled", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).AreExternalFieldsFilled), record)
}

// ForAllExpiredFlowRecordsDo mocks base method.
func (m *MockIPFIXAggregationProcess) ForAllExpiredFlowRecordsDo(callback intermediate.FlowKeyRecordMapCallBack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForAllExpiredFlowRecordsDo", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForAllExpiredFlowRecordsDo indicates an expected call of ForAllExpiredFlowRecordsDo.
func (mr *MockIPFIXAggregationProcessMockRecorder) ForAllExpiredFlowRecordsDo(callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForAllExpiredFlowRecordsDo", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).ForAllExpiredFlowRecordsDo), callback)
}

// GetExpiryFromExpirePriorityQueue mocks base method.
func (m *MockIPFIXAggregationProcess) GetExpiryFromExpirePriorityQueue() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpiryFromExpirePriorityQueue")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetExpiryFromExpirePriorityQueue indicates an expected call of GetExpiryFromExpirePriorityQueue.
func (mr *MockIPFIXAggregationProcessMockRecorder) GetExpiryFromExpirePriorityQueue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpiryFromExpirePriorityQueue", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).GetExpiryFromExpirePriorityQueue))
}

// GetNumFlows mocks base method.
func (m *MockIPFIXAggregationProcess) GetNumFlows() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumFlows")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetNumFlows indicates an expected call of GetNumFlows.
func (mr *MockIPFIXAggregationProcessMockRecorder) GetNumFlows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumFlows", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).GetNumFlows))
}

// GetRecords mocks base method.
func (m *MockIPFIXAggregationProcess) GetRecords(flowKey *intermediate.FlowKey) []map[string]any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecords", flowKey)
	ret0, _ := ret[0].([]map[string]any)
	return ret0
}

// GetRecords indicates an expected call of GetRecords.
func (mr *MockIPFIXAggregationProcessMockRecorder) GetRecords(flowKey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecords", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).GetRecords), flowKey)
}

// IsAggregatedRecordIPv4 mocks base method.
func (m *MockIPFIXAggregationProcess) IsAggregatedRecordIPv4(record intermediate.AggregationFlowRecord) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAggregatedRecordIPv4", record)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAggregatedRecordIPv4 indicates an expected call of IsAggregatedRecordIPv4.
func (mr *MockIPFIXAggregationProcessMockRecorder) IsAggregatedRecordIPv4(record any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAggregatedRecordIPv4", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).IsAggregatedRecordIPv4), record)
}

// ResetStatAndThroughputElementsInRecord mocks base method.
func (m *MockIPFIXAggregationProcess) ResetStatAndThroughputElementsInRecord(record entities.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetStatAndThroughputElementsInRecord", record)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetStatAndThroughputElementsInRecord indicates an expected call of ResetStatAndThroughputElementsInRecord.
func (mr *MockIPFIXAggregationProcessMockRecorder) ResetStatAndThroughputElementsInRecord(record any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetStatAndThroughputElementsInRecord", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).ResetStatAndThroughputElementsInRecord), record)
}

// SetCorrelatedFieldsFilled mocks base method.
func (m *MockIPFIXAggregationProcess) SetCorrelatedFieldsFilled(record *intermediate.AggregationFlowRecord, isFilled bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCorrelatedFieldsFilled", record, isFilled)
}

// SetCorrelatedFieldsFilled indicates an expected call of SetCorrelatedFieldsFilled.
func (mr *MockIPFIXAggregationProcessMockRecorder) SetCorrelatedFieldsFilled(record, isFilled any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCorrelatedFieldsFilled", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).SetCorrelatedFieldsFilled), record, isFilled)
}

// SetExternalFieldsFilled mocks base method.
func (m *MockIPFIXAggregationProcess) SetExternalFieldsFilled(record *intermediate.AggregationFlowRecord, isFilled bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetExternalFieldsFilled", record, isFilled)
}

// SetExternalFieldsFilled indicates an expected call of SetExternalFieldsFilled.
func (mr *MockIPFIXAggregationProcessMockRecorder) SetExternalFieldsFilled(record, isFilled any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExternalFieldsFilled", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).SetExternalFieldsFilled), record, isFilled)
}

// Start mocks base method.
func (m *MockIPFIXAggregationProcess) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockIPFIXAggregationProcessMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).Start))
}

// Stop mocks base method.
func (m *MockIPFIXAggregationProcess) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockIPFIXAggregationProcessMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockIPFIXAggregationProcess)(nil).Stop))
}
