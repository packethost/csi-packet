// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cloud_provider/volume.go

// Package mock_cloud_provider is a generated GoMock package.
package test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	packngo "github.com/packethost/packngo"
)

// MockVolumeProvider is a mock of VolumeProvider interface
type MockVolumeProvider struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeProviderMockRecorder
}

// MockVolumeProviderMockRecorder is the mock recorder for MockVolumeProvider
type MockVolumeProviderMockRecorder struct {
	mock *MockVolumeProvider
}

// NewMockVolumeProvider creates a new mock instance
func NewMockVolumeProvider(ctrl *gomock.Controller) *MockVolumeProvider {
	mock := &MockVolumeProvider{ctrl: ctrl}
	mock.recorder = &MockVolumeProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVolumeProvider) EXPECT() *MockVolumeProviderMockRecorder {
	return m.recorder
}

// ListVolumes mocks base method
func (m *MockVolumeProvider) ListVolumes(options *packngo.ListOptions) ([]packngo.Volume, *packngo.Response, error) {
	ret := m.ctrl.Call(m, "ListVolumes")
	ret0, _ := ret[0].([]packngo.Volume)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListVolumes indicates an expected call of ListVolumes
func (mr *MockVolumeProviderMockRecorder) ListVolumes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumes", reflect.TypeOf((*MockVolumeProvider)(nil).ListVolumes))
}

// Get mocks base method
func (m *MockVolumeProvider) Get(volumeID string) (*packngo.Volume, *packngo.Response, error) {
	ret := m.ctrl.Call(m, "Get", volumeID)
	ret0, _ := ret[0].(*packngo.Volume)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get
func (mr *MockVolumeProviderMockRecorder) Get(volumeID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVolumeProvider)(nil).Get), volumeID)
}

// Delete mocks base method
func (m *MockVolumeProvider) Delete(volumeID string) (*packngo.Response, error) {
	ret := m.ctrl.Call(m, "Delete", volumeID)
	ret0, _ := ret[0].(*packngo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockVolumeProviderMockRecorder) Delete(volumeID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVolumeProvider)(nil).Delete), volumeID)
}

// Create mocks base method
func (m *MockVolumeProvider) Create(arg0 *packngo.VolumeCreateRequest) (*packngo.Volume, *packngo.Response, error) {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*packngo.Volume)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create
func (mr *MockVolumeProviderMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVolumeProvider)(nil).Create), arg0)
}

// Attach mocks base method
func (m *MockVolumeProvider) Attach(volumeID, deviceID string) (*packngo.VolumeAttachment, *packngo.Response, error) {
	ret := m.ctrl.Call(m, "Attach", volumeID, deviceID)
	ret0, _ := ret[0].(*packngo.VolumeAttachment)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Attach indicates an expected call of Attach
func (mr *MockVolumeProviderMockRecorder) Attach(volumeID, deviceID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attach", reflect.TypeOf((*MockVolumeProvider)(nil).Attach), volumeID, deviceID)
}

// Detach mocks base method
func (m *MockVolumeProvider) Detach(attachmentID string) (*packngo.Response, error) {
	ret := m.ctrl.Call(m, "Detach", attachmentID)
	ret0, _ := ret[0].(*packngo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Detach indicates an expected call of Detach
func (mr *MockVolumeProviderMockRecorder) Detach(attachmentID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detach", reflect.TypeOf((*MockVolumeProvider)(nil).Detach), attachmentID)
}

// GetNodes mocks base method
func (m *MockVolumeProvider) GetNodes() ([]packngo.Device, *packngo.Response, error) {
	ret := m.ctrl.Call(m, "GetNodes")
	ret0, _ := ret[0].([]packngo.Device)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNodes indicates an expected call of GetNodes
func (mr *MockVolumeProviderMockRecorder) GetNodes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodes", reflect.TypeOf((*MockVolumeProvider)(nil).GetNodes))
}

// MockNodeVolumeManager is a mock of NodeVolumeManager interface
type MockNodeVolumeManager struct {
	ctrl     *gomock.Controller
	recorder *MockNodeVolumeManagerMockRecorder
}

// MockNodeVolumeManagerMockRecorder is the mock recorder for MockNodeVolumeManager
type MockNodeVolumeManagerMockRecorder struct {
	mock *MockNodeVolumeManager
}

// NewMockNodeVolumeManager creates a new mock instance
func NewMockNodeVolumeManager(ctrl *gomock.Controller) *MockNodeVolumeManager {
	mock := &MockNodeVolumeManager{ctrl: ctrl}
	mock.recorder = &MockNodeVolumeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeVolumeManager) EXPECT() *MockNodeVolumeManagerMockRecorder {
	return m.recorder
}
