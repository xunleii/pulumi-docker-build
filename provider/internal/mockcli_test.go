// Code generated by MockGen. DO NOT EDIT.
// Source: cli.go
//
// Generated by this command:
//
//	mockgen -typed -package internal -source cli.go -destination mockcli_test.go --self_package github.com/pulumi/pulumi-docker-build/provider/internal
//
// Package internal is a generated GoMock package.
package internal

import (
	io "io"
	reflect "reflect"

	command "github.com/docker/cli/cli/command"
	configfile "github.com/docker/cli/cli/config/configfile"
	docker "github.com/docker/cli/cli/context/docker"
	store "github.com/docker/cli/cli/context/store"
	store0 "github.com/docker/cli/cli/manifest/store"
	client "github.com/docker/cli/cli/registry/client"
	streams "github.com/docker/cli/cli/streams"
	trust "github.com/docker/cli/cli/trust"
	client0 "github.com/docker/docker/client"
	client1 "github.com/theupdateframework/notary/client"
	gomock "go.uber.org/mock/gomock"
)

// MockCli is a mock of Cli interface.
type MockCli struct {
	ctrl     *gomock.Controller
	recorder *MockCliMockRecorder
}

// MockCliMockRecorder is the mock recorder for MockCli.
type MockCliMockRecorder struct {
	mock *MockCli
}

// NewMockCli creates a new mock instance.
func NewMockCli(ctrl *gomock.Controller) *MockCli {
	mock := &MockCli{ctrl: ctrl}
	mock.recorder = &MockCliMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCli) EXPECT() *MockCliMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockCli) Apply(ops ...command.CLIOption) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range ops {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Apply", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockCliMockRecorder) Apply(ops ...any) *CliApplyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockCli)(nil).Apply), ops...)
	return &CliApplyCall{Call: call}
}

// CliApplyCall wrap *gomock.Call
type CliApplyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliApplyCall) Return(arg0 error) *CliApplyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliApplyCall) Do(f func(...command.CLIOption) error) *CliApplyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliApplyCall) DoAndReturn(f func(...command.CLIOption) error) *CliApplyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// BuildKitEnabled mocks base method.
func (m *MockCli) BuildKitEnabled() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildKitEnabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildKitEnabled indicates an expected call of BuildKitEnabled.
func (mr *MockCliMockRecorder) BuildKitEnabled() *CliBuildKitEnabledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildKitEnabled", reflect.TypeOf((*MockCli)(nil).BuildKitEnabled))
	return &CliBuildKitEnabledCall{Call: call}
}

// CliBuildKitEnabledCall wrap *gomock.Call
type CliBuildKitEnabledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliBuildKitEnabledCall) Return(arg0 bool, arg1 error) *CliBuildKitEnabledCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliBuildKitEnabledCall) Do(f func() (bool, error)) *CliBuildKitEnabledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliBuildKitEnabledCall) DoAndReturn(f func() (bool, error)) *CliBuildKitEnabledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Client mocks base method.
func (m *MockCli) Client() client0.APIClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(client0.APIClient)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockCliMockRecorder) Client() *CliClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockCli)(nil).Client))
	return &CliClientCall{Call: call}
}

// CliClientCall wrap *gomock.Call
type CliClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliClientCall) Return(arg0 client0.APIClient) *CliClientCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliClientCall) Do(f func() client0.APIClient) *CliClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliClientCall) DoAndReturn(f func() client0.APIClient) *CliClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConfigFile mocks base method.
func (m *MockCli) ConfigFile() *configfile.ConfigFile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigFile")
	ret0, _ := ret[0].(*configfile.ConfigFile)
	return ret0
}

// ConfigFile indicates an expected call of ConfigFile.
func (mr *MockCliMockRecorder) ConfigFile() *CliConfigFileCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigFile", reflect.TypeOf((*MockCli)(nil).ConfigFile))
	return &CliConfigFileCall{Call: call}
}

// CliConfigFileCall wrap *gomock.Call
type CliConfigFileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliConfigFileCall) Return(arg0 *configfile.ConfigFile) *CliConfigFileCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliConfigFileCall) Do(f func() *configfile.ConfigFile) *CliConfigFileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliConfigFileCall) DoAndReturn(f func() *configfile.ConfigFile) *CliConfigFileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ContentTrustEnabled mocks base method.
func (m *MockCli) ContentTrustEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContentTrustEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ContentTrustEnabled indicates an expected call of ContentTrustEnabled.
func (mr *MockCliMockRecorder) ContentTrustEnabled() *CliContentTrustEnabledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentTrustEnabled", reflect.TypeOf((*MockCli)(nil).ContentTrustEnabled))
	return &CliContentTrustEnabledCall{Call: call}
}

// CliContentTrustEnabledCall wrap *gomock.Call
type CliContentTrustEnabledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliContentTrustEnabledCall) Return(arg0 bool) *CliContentTrustEnabledCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliContentTrustEnabledCall) Do(f func() bool) *CliContentTrustEnabledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliContentTrustEnabledCall) DoAndReturn(f func() bool) *CliContentTrustEnabledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ContextStore mocks base method.
func (m *MockCli) ContextStore() store.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContextStore")
	ret0, _ := ret[0].(store.Store)
	return ret0
}

// ContextStore indicates an expected call of ContextStore.
func (mr *MockCliMockRecorder) ContextStore() *CliContextStoreCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContextStore", reflect.TypeOf((*MockCli)(nil).ContextStore))
	return &CliContextStoreCall{Call: call}
}

// CliContextStoreCall wrap *gomock.Call
type CliContextStoreCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliContextStoreCall) Return(arg0 store.Store) *CliContextStoreCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliContextStoreCall) Do(f func() store.Store) *CliContextStoreCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliContextStoreCall) DoAndReturn(f func() store.Store) *CliContextStoreCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CurrentContext mocks base method.
func (m *MockCli) CurrentContext() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentContext")
	ret0, _ := ret[0].(string)
	return ret0
}

// CurrentContext indicates an expected call of CurrentContext.
func (mr *MockCliMockRecorder) CurrentContext() *CliCurrentContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentContext", reflect.TypeOf((*MockCli)(nil).CurrentContext))
	return &CliCurrentContextCall{Call: call}
}

// CliCurrentContextCall wrap *gomock.Call
type CliCurrentContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliCurrentContextCall) Return(arg0 string) *CliCurrentContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliCurrentContextCall) Do(f func() string) *CliCurrentContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliCurrentContextCall) DoAndReturn(f func() string) *CliCurrentContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CurrentVersion mocks base method.
func (m *MockCli) CurrentVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// CurrentVersion indicates an expected call of CurrentVersion.
func (mr *MockCliMockRecorder) CurrentVersion() *CliCurrentVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentVersion", reflect.TypeOf((*MockCli)(nil).CurrentVersion))
	return &CliCurrentVersionCall{Call: call}
}

// CliCurrentVersionCall wrap *gomock.Call
type CliCurrentVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliCurrentVersionCall) Return(arg0 string) *CliCurrentVersionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliCurrentVersionCall) Do(f func() string) *CliCurrentVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliCurrentVersionCall) DoAndReturn(f func() string) *CliCurrentVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DefaultVersion mocks base method.
func (m *MockCli) DefaultVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// DefaultVersion indicates an expected call of DefaultVersion.
func (mr *MockCliMockRecorder) DefaultVersion() *CliDefaultVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultVersion", reflect.TypeOf((*MockCli)(nil).DefaultVersion))
	return &CliDefaultVersionCall{Call: call}
}

// CliDefaultVersionCall wrap *gomock.Call
type CliDefaultVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliDefaultVersionCall) Return(arg0 string) *CliDefaultVersionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliDefaultVersionCall) Do(f func() string) *CliDefaultVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliDefaultVersionCall) DoAndReturn(f func() string) *CliDefaultVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DockerEndpoint mocks base method.
func (m *MockCli) DockerEndpoint() docker.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DockerEndpoint")
	ret0, _ := ret[0].(docker.Endpoint)
	return ret0
}

// DockerEndpoint indicates an expected call of DockerEndpoint.
func (mr *MockCliMockRecorder) DockerEndpoint() *CliDockerEndpointCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DockerEndpoint", reflect.TypeOf((*MockCli)(nil).DockerEndpoint))
	return &CliDockerEndpointCall{Call: call}
}

// CliDockerEndpointCall wrap *gomock.Call
type CliDockerEndpointCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliDockerEndpointCall) Return(arg0 docker.Endpoint) *CliDockerEndpointCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliDockerEndpointCall) Do(f func() docker.Endpoint) *CliDockerEndpointCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliDockerEndpointCall) DoAndReturn(f func() docker.Endpoint) *CliDockerEndpointCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Err mocks base method.
func (m *MockCli) Err() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockCliMockRecorder) Err() *CliErrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockCli)(nil).Err))
	return &CliErrCall{Call: call}
}

// CliErrCall wrap *gomock.Call
type CliErrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliErrCall) Return(arg0 io.Writer) *CliErrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliErrCall) Do(f func() io.Writer) *CliErrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliErrCall) DoAndReturn(f func() io.Writer) *CliErrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// In mocks base method.
func (m *MockCli) In() *streams.In {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "In")
	ret0, _ := ret[0].(*streams.In)
	return ret0
}

// In indicates an expected call of In.
func (mr *MockCliMockRecorder) In() *CliInCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "In", reflect.TypeOf((*MockCli)(nil).In))
	return &CliInCall{Call: call}
}

// CliInCall wrap *gomock.Call
type CliInCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliInCall) Return(arg0 *streams.In) *CliInCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliInCall) Do(f func() *streams.In) *CliInCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliInCall) DoAndReturn(f func() *streams.In) *CliInCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ManifestStore mocks base method.
func (m *MockCli) ManifestStore() store0.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManifestStore")
	ret0, _ := ret[0].(store0.Store)
	return ret0
}

// ManifestStore indicates an expected call of ManifestStore.
func (mr *MockCliMockRecorder) ManifestStore() *CliManifestStoreCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManifestStore", reflect.TypeOf((*MockCli)(nil).ManifestStore))
	return &CliManifestStoreCall{Call: call}
}

// CliManifestStoreCall wrap *gomock.Call
type CliManifestStoreCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliManifestStoreCall) Return(arg0 store0.Store) *CliManifestStoreCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliManifestStoreCall) Do(f func() store0.Store) *CliManifestStoreCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliManifestStoreCall) DoAndReturn(f func() store0.Store) *CliManifestStoreCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NotaryClient mocks base method.
func (m *MockCli) NotaryClient(imgRefAndAuth trust.ImageRefAndAuth, actions []string) (client1.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotaryClient", imgRefAndAuth, actions)
	ret0, _ := ret[0].(client1.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotaryClient indicates an expected call of NotaryClient.
func (mr *MockCliMockRecorder) NotaryClient(imgRefAndAuth, actions any) *CliNotaryClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotaryClient", reflect.TypeOf((*MockCli)(nil).NotaryClient), imgRefAndAuth, actions)
	return &CliNotaryClientCall{Call: call}
}

// CliNotaryClientCall wrap *gomock.Call
type CliNotaryClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliNotaryClientCall) Return(arg0 client1.Repository, arg1 error) *CliNotaryClientCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliNotaryClientCall) Do(f func(trust.ImageRefAndAuth, []string) (client1.Repository, error)) *CliNotaryClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliNotaryClientCall) DoAndReturn(f func(trust.ImageRefAndAuth, []string) (client1.Repository, error)) *CliNotaryClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Out mocks base method.
func (m *MockCli) Out() *streams.Out {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Out")
	ret0, _ := ret[0].(*streams.Out)
	return ret0
}

// Out indicates an expected call of Out.
func (mr *MockCliMockRecorder) Out() *CliOutCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Out", reflect.TypeOf((*MockCli)(nil).Out))
	return &CliOutCall{Call: call}
}

// CliOutCall wrap *gomock.Call
type CliOutCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliOutCall) Return(arg0 *streams.Out) *CliOutCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliOutCall) Do(f func() *streams.Out) *CliOutCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliOutCall) DoAndReturn(f func() *streams.Out) *CliOutCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RegistryClient mocks base method.
func (m *MockCli) RegistryClient(arg0 bool) client.RegistryClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryClient", arg0)
	ret0, _ := ret[0].(client.RegistryClient)
	return ret0
}

// RegistryClient indicates an expected call of RegistryClient.
func (mr *MockCliMockRecorder) RegistryClient(arg0 any) *CliRegistryClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryClient", reflect.TypeOf((*MockCli)(nil).RegistryClient), arg0)
	return &CliRegistryClientCall{Call: call}
}

// CliRegistryClientCall wrap *gomock.Call
type CliRegistryClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliRegistryClientCall) Return(arg0 client.RegistryClient) *CliRegistryClientCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliRegistryClientCall) Do(f func(bool) client.RegistryClient) *CliRegistryClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliRegistryClientCall) DoAndReturn(f func(bool) client.RegistryClient) *CliRegistryClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ServerInfo mocks base method.
func (m *MockCli) ServerInfo() command.ServerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerInfo")
	ret0, _ := ret[0].(command.ServerInfo)
	return ret0
}

// ServerInfo indicates an expected call of ServerInfo.
func (mr *MockCliMockRecorder) ServerInfo() *CliServerInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerInfo", reflect.TypeOf((*MockCli)(nil).ServerInfo))
	return &CliServerInfoCall{Call: call}
}

// CliServerInfoCall wrap *gomock.Call
type CliServerInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliServerInfoCall) Return(arg0 command.ServerInfo) *CliServerInfoCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliServerInfoCall) Do(f func() command.ServerInfo) *CliServerInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliServerInfoCall) DoAndReturn(f func() command.ServerInfo) *CliServerInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetIn mocks base method.
func (m *MockCli) SetIn(in *streams.In) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIn", in)
}

// SetIn indicates an expected call of SetIn.
func (mr *MockCliMockRecorder) SetIn(in any) *CliSetInCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIn", reflect.TypeOf((*MockCli)(nil).SetIn), in)
	return &CliSetInCall{Call: call}
}

// CliSetInCall wrap *gomock.Call
type CliSetInCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CliSetInCall) Return() *CliSetInCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CliSetInCall) Do(f func(*streams.In)) *CliSetInCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CliSetInCall) DoAndReturn(f func(*streams.In)) *CliSetInCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
