// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/kt/cluster/types.go

// Package cluster is a generated GoMock package.
package cluster

import (
	context "context"
	reflect "reflect"

	options "github.com/alibaba/kt-connect/pkg/kt/options"
	util "github.com/alibaba/kt-connect/pkg/kt/util"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
)

// MockKubernetesInterface is a mock of KubernetesInterface interface.
type MockKubernetesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesInterfaceMockRecorder
}

// MockKubernetesInterfaceMockRecorder is the mock recorder for MockKubernetesInterface.
type MockKubernetesInterfaceMockRecorder struct {
	mock *MockKubernetesInterface
}

// NewMockKubernetesInterface creates a new mock instance.
func NewMockKubernetesInterface(ctrl *gomock.Controller) *MockKubernetesInterface {
	mock := &MockKubernetesInterface{ctrl: ctrl}
	mock.recorder = &MockKubernetesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesInterface) EXPECT() *MockKubernetesInterfaceMockRecorder {
	return m.recorder
}

// AddEphemeralContainer mocks base method.
func (m *MockKubernetesInterface) AddEphemeralContainer(ctx context.Context, containerName, podName string, options *options.DaemonOptions, envs map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEphemeralContainer", ctx, containerName, podName, options, envs)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddEphemeralContainer indicates an expected call of AddEphemeralContainer.
func (mr *MockKubernetesInterfaceMockRecorder) AddEphemeralContainer(ctx, containerName, podName, options, envs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEphemeralContainer", reflect.TypeOf((*MockKubernetesInterface)(nil).AddEphemeralContainer), ctx, containerName, podName, options, envs)
}

// ClusterCidrs mocks base method.
func (m *MockKubernetesInterface) ClusterCidrs(ctx context.Context, namespace string, connectOptions *options.ConnectOptions) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterCidrs", ctx, namespace, connectOptions)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterCidrs indicates an expected call of ClusterCidrs.
func (mr *MockKubernetesInterfaceMockRecorder) ClusterCidrs(ctx, namespace, connectOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterCidrs", reflect.TypeOf((*MockKubernetesInterface)(nil).ClusterCidrs), ctx, namespace, connectOptions)
}

// CreateConfigMapWithSshKey mocks base method.
func (m *MockKubernetesInterface) CreateConfigMapWithSshKey(ctx context.Context, labels map[string]string, sshcm, namespace string, generator *util.SSHGenerator) (*v10.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConfigMapWithSshKey", ctx, labels, sshcm, namespace, generator)
	ret0, _ := ret[0].(*v10.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConfigMapWithSshKey indicates an expected call of CreateConfigMapWithSshKey.
func (mr *MockKubernetesInterfaceMockRecorder) CreateConfigMapWithSshKey(ctx, labels, sshcm, namespace, generator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConfigMapWithSshKey", reflect.TypeOf((*MockKubernetesInterface)(nil).CreateConfigMapWithSshKey), ctx, labels, sshcm, namespace, generator)
}

// CreatePod mocks base method.
func (m *MockKubernetesInterface) CreatePod(ctx context.Context, metaAndSpec *PodMetaAndSpec, options *options.DaemonOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePod", ctx, metaAndSpec, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePod indicates an expected call of CreatePod.
func (mr *MockKubernetesInterfaceMockRecorder) CreatePod(ctx, metaAndSpec, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePod", reflect.TypeOf((*MockKubernetesInterface)(nil).CreatePod), ctx, metaAndSpec, options)
}

// CreateService mocks base method.
func (m *MockKubernetesInterface) CreateService(ctx context.Context, metaAndSpec *SvcMetaAndSpec) (*v10.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", ctx, metaAndSpec)
	ret0, _ := ret[0].(*v10.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService.
func (mr *MockKubernetesInterfaceMockRecorder) CreateService(ctx, metaAndSpec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockKubernetesInterface)(nil).CreateService), ctx, metaAndSpec)
}

// CreateShadowPod mocks base method.
func (m *MockKubernetesInterface) CreateShadowPod(ctx context.Context, metaAndSpec *PodMetaAndSpec, sshcm string, options *options.DaemonOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShadowPod", ctx, metaAndSpec, sshcm, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateShadowPod indicates an expected call of CreateShadowPod.
func (mr *MockKubernetesInterfaceMockRecorder) CreateShadowPod(ctx, metaAndSpec, sshcm, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShadowPod", reflect.TypeOf((*MockKubernetesInterface)(nil).CreateShadowPod), ctx, metaAndSpec, sshcm, options)
}

// DecreaseRef mocks base method.
func (m *MockKubernetesInterface) DecreaseRef(ctx context.Context, name, namespace string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseRef", ctx, name, namespace)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecreaseRef indicates an expected call of DecreaseRef.
func (mr *MockKubernetesInterfaceMockRecorder) DecreaseRef(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseRef", reflect.TypeOf((*MockKubernetesInterface)(nil).DecreaseRef), ctx, name, namespace)
}

// ExecInPod mocks base method.
func (m *MockKubernetesInterface) ExecInPod(containerName, podName, namespace string, opts options.RuntimeOptions, cmd ...string) (string, string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{containerName, podName, namespace, opts}
	for _, a := range cmd {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecInPod", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExecInPod indicates an expected call of ExecInPod.
func (mr *MockKubernetesInterfaceMockRecorder) ExecInPod(containerName, podName, namespace, opts interface{}, cmd ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{containerName, podName, namespace, opts}, cmd...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecInPod", reflect.TypeOf((*MockKubernetesInterface)(nil).ExecInPod), varargs...)
}

// GetConfigMap mocks base method.
func (m *MockKubernetesInterface) GetConfigMap(ctx context.Context, name, namespace string) (*v10.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigMap", ctx, name, namespace)
	ret0, _ := ret[0].(*v10.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigMap indicates an expected call of GetConfigMap.
func (mr *MockKubernetesInterfaceMockRecorder) GetConfigMap(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigMap", reflect.TypeOf((*MockKubernetesInterface)(nil).GetConfigMap), ctx, name, namespace)
}

// GetDeployment mocks base method.
func (m *MockKubernetesInterface) GetDeployment(ctx context.Context, name, namespace string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", ctx, name, namespace)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockKubernetesInterfaceMockRecorder) GetDeployment(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockKubernetesInterface)(nil).GetDeployment), ctx, name, namespace)
}

// GetPod mocks base method.
func (m *MockKubernetesInterface) GetPod(ctx context.Context, name, namespace string) (*v10.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPod", ctx, name, namespace)
	ret0, _ := ret[0].(*v10.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPod indicates an expected call of GetPod.
func (mr *MockKubernetesInterfaceMockRecorder) GetPod(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPod", reflect.TypeOf((*MockKubernetesInterface)(nil).GetPod), ctx, name, namespace)
}

// GetPods mocks base method.
func (m *MockKubernetesInterface) GetPods(ctx context.Context, labels map[string]string, namespace string) (*v10.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPods", ctx, labels, namespace)
	ret0, _ := ret[0].(*v10.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPods indicates an expected call of GetPods.
func (mr *MockKubernetesInterfaceMockRecorder) GetPods(ctx, labels, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPods", reflect.TypeOf((*MockKubernetesInterface)(nil).GetPods), ctx, labels, namespace)
}

// GetService mocks base method.
func (m *MockKubernetesInterface) GetService(ctx context.Context, name, namespace string) (*v10.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", ctx, name, namespace)
	ret0, _ := ret[0].(*v10.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockKubernetesInterfaceMockRecorder) GetService(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockKubernetesInterface)(nil).GetService), ctx, name, namespace)
}

// GetServiceHosts mocks base method.
func (m *MockKubernetesInterface) GetServiceHosts(ctx context.Context, namespace string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceHosts", ctx, namespace)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetServiceHosts indicates an expected call of GetServiceHosts.
func (mr *MockKubernetesInterfaceMockRecorder) GetServiceHosts(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceHosts", reflect.TypeOf((*MockKubernetesInterface)(nil).GetServiceHosts), ctx, namespace)
}

// GetServices mocks base method.
func (m *MockKubernetesInterface) GetServices(ctx context.Context, matchLabels map[string]string, namespace string) ([]v10.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServices", ctx, matchLabels, namespace)
	ret0, _ := ret[0].([]v10.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServices indicates an expected call of GetServices.
func (mr *MockKubernetesInterfaceMockRecorder) GetServices(ctx, matchLabels, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServices", reflect.TypeOf((*MockKubernetesInterface)(nil).GetServices), ctx, matchLabels, namespace)
}

// IncreaseRef mocks base method.
func (m *MockKubernetesInterface) IncreaseRef(ctx context.Context, name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseRef", ctx, name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseRef indicates an expected call of IncreaseRef.
func (mr *MockKubernetesInterfaceMockRecorder) IncreaseRef(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseRef", reflect.TypeOf((*MockKubernetesInterface)(nil).IncreaseRef), ctx, name, namespace)
}

// RemoveConfigMap mocks base method.
func (m *MockKubernetesInterface) RemoveConfigMap(ctx context.Context, name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveConfigMap", ctx, name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveConfigMap indicates an expected call of RemoveConfigMap.
func (mr *MockKubernetesInterfaceMockRecorder) RemoveConfigMap(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveConfigMap", reflect.TypeOf((*MockKubernetesInterface)(nil).RemoveConfigMap), ctx, name, namespace)
}

// RemoveEphemeralContainer mocks base method.
func (m *MockKubernetesInterface) RemoveEphemeralContainer(ctx context.Context, containerName, podName, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEphemeralContainer", ctx, containerName, podName, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEphemeralContainer indicates an expected call of RemoveEphemeralContainer.
func (mr *MockKubernetesInterfaceMockRecorder) RemoveEphemeralContainer(ctx, containerName, podName, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEphemeralContainer", reflect.TypeOf((*MockKubernetesInterface)(nil).RemoveEphemeralContainer), ctx, containerName, podName, namespace)
}

// RemovePod mocks base method.
func (m *MockKubernetesInterface) RemovePod(ctx context.Context, name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePod", ctx, name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePod indicates an expected call of RemovePod.
func (mr *MockKubernetesInterfaceMockRecorder) RemovePod(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePod", reflect.TypeOf((*MockKubernetesInterface)(nil).RemovePod), ctx, name, namespace)
}

// RemoveService mocks base method.
func (m *MockKubernetesInterface) RemoveService(ctx context.Context, name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveService", ctx, name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveService indicates an expected call of RemoveService.
func (mr *MockKubernetesInterfaceMockRecorder) RemoveService(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveService", reflect.TypeOf((*MockKubernetesInterface)(nil).RemoveService), ctx, name, namespace)
}

// ScaleTo mocks base method.
func (m *MockKubernetesInterface) ScaleTo(ctx context.Context, deployment, namespace string, replicas *int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleTo", ctx, deployment, namespace, replicas)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScaleTo indicates an expected call of ScaleTo.
func (mr *MockKubernetesInterfaceMockRecorder) ScaleTo(ctx, deployment, namespace, replicas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleTo", reflect.TypeOf((*MockKubernetesInterface)(nil).ScaleTo), ctx, deployment, namespace, replicas)
}

// UpdateDeployment mocks base method.
func (m *MockKubernetesInterface) UpdateDeployment(ctx context.Context, deployment *v1.Deployment) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeployment", ctx, deployment)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDeployment indicates an expected call of UpdateDeployment.
func (mr *MockKubernetesInterfaceMockRecorder) UpdateDeployment(ctx, deployment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeployment", reflect.TypeOf((*MockKubernetesInterface)(nil).UpdateDeployment), ctx, deployment)
}

// UpdatePod mocks base method.
func (m *MockKubernetesInterface) UpdatePod(ctx context.Context, pod *v10.Pod) (*v10.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePod", ctx, pod)
	ret0, _ := ret[0].(*v10.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePod indicates an expected call of UpdatePod.
func (mr *MockKubernetesInterfaceMockRecorder) UpdatePod(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePod", reflect.TypeOf((*MockKubernetesInterface)(nil).UpdatePod), ctx, pod)
}

// UpdateService mocks base method.
func (m *MockKubernetesInterface) UpdateService(ctx context.Context, svc *v10.Service) (*v10.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", ctx, svc)
	ret0, _ := ret[0].(*v10.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockKubernetesInterfaceMockRecorder) UpdateService(ctx, svc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockKubernetesInterface)(nil).UpdateService), ctx, svc)
}

// WaitPodReady mocks base method.
func (m *MockKubernetesInterface) WaitPodReady(name, namespace string) (*v10.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitPodReady", name, namespace)
	ret0, _ := ret[0].(*v10.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitPodReady indicates an expected call of WaitPodReady.
func (mr *MockKubernetesInterfaceMockRecorder) WaitPodReady(name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitPodReady", reflect.TypeOf((*MockKubernetesInterface)(nil).WaitPodReady), name, namespace)
}
