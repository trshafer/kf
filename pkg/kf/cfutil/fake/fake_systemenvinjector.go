// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/kf/pkg/kf/cfutil/fake (interfaces: SystemEnvInjector)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/google/kf/pkg/apis/kf/v1alpha1"
	cfutil "github.com/google/kf/pkg/kf/cfutil"
	v1beta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
)

// FakeSystemEnvInjector is a mock of SystemEnvInjector interface
type FakeSystemEnvInjector struct {
	ctrl     *gomock.Controller
	recorder *FakeSystemEnvInjectorMockRecorder
}

// FakeSystemEnvInjectorMockRecorder is the mock recorder for FakeSystemEnvInjector
type FakeSystemEnvInjectorMockRecorder struct {
	mock *FakeSystemEnvInjector
}

// NewFakeSystemEnvInjector creates a new mock instance
func NewFakeSystemEnvInjector(ctrl *gomock.Controller) *FakeSystemEnvInjector {
	mock := &FakeSystemEnvInjector{ctrl: ctrl}
	mock.recorder = &FakeSystemEnvInjectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *FakeSystemEnvInjector) EXPECT() *FakeSystemEnvInjectorMockRecorder {
	return m.recorder
}

// ComputeSystemEnv mocks base method
func (m *FakeSystemEnvInjector) ComputeSystemEnv(arg0 *v1alpha1.App, arg1 []v1beta1.ServiceBinding) ([]v1.EnvVar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeSystemEnv", arg0, arg1)
	ret0, _ := ret[0].([]v1.EnvVar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeSystemEnv indicates an expected call of ComputeSystemEnv
func (mr *FakeSystemEnvInjectorMockRecorder) ComputeSystemEnv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeSystemEnv", reflect.TypeOf((*FakeSystemEnvInjector)(nil).ComputeSystemEnv), arg0, arg1)
}

// GetVcapService mocks base method
func (m *FakeSystemEnvInjector) GetVcapService(arg0 string, arg1 *v1beta1.ServiceBinding) (cfutil.VcapService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVcapService", arg0, arg1)
	ret0, _ := ret[0].(cfutil.VcapService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVcapService indicates an expected call of GetVcapService
func (mr *FakeSystemEnvInjectorMockRecorder) GetVcapService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVcapService", reflect.TypeOf((*FakeSystemEnvInjector)(nil).GetVcapService), arg0, arg1)
}

// GetVcapServices mocks base method
func (m *FakeSystemEnvInjector) GetVcapServices(arg0 string, arg1 []v1beta1.ServiceBinding) ([]cfutil.VcapService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVcapServices", arg0, arg1)
	ret0, _ := ret[0].([]cfutil.VcapService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVcapServices indicates an expected call of GetVcapServices
func (mr *FakeSystemEnvInjectorMockRecorder) GetVcapServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVcapServices", reflect.TypeOf((*FakeSystemEnvInjector)(nil).GetVcapServices), arg0, arg1)
}
