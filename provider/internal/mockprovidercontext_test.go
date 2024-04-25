// Code generated by MockGen. DO NOT EDIT.
// Source: providercontext.go
//
// Generated by this command:
//
//	mockgen -typed -package internal -source providercontext.go -destination mockprovidercontext_test.go --self_package github.com/pulumi/pulumi-docker/provider/v4/internal
//
// Package internal is a generated GoMock package.
package internal

import (
	reflect "reflect"
	time "time"

	provider "github.com/pulumi/pulumi-go-provider"
	diag "github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	gomock "go.uber.org/mock/gomock"
)

// MockProviderContext is a mock of ProviderContext interface.
type MockProviderContext struct {
	ctrl     *gomock.Controller
	recorder *MockProviderContextMockRecorder
}

// MockProviderContextMockRecorder is the mock recorder for MockProviderContext.
type MockProviderContextMockRecorder struct {
	mock *MockProviderContext
}

// NewMockProviderContext creates a new mock instance.
func NewMockProviderContext(ctrl *gomock.Controller) *MockProviderContext {
	mock := &MockProviderContext{ctrl: ctrl}
	mock.recorder = &MockProviderContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderContext) EXPECT() *MockProviderContextMockRecorder {
	return m.recorder
}

// Deadline mocks base method.
func (m *MockProviderContext) Deadline() (time.Time, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deadline")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Deadline indicates an expected call of Deadline.
func (mr *MockProviderContextMockRecorder) Deadline() *ProviderContextDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deadline", reflect.TypeOf((*MockProviderContext)(nil).Deadline))
	return &ProviderContextDeadlineCall{Call: call}
}

// ProviderContextDeadlineCall wrap *gomock.Call
type ProviderContextDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderContextDeadlineCall) Return(deadline time.Time, ok bool) *ProviderContextDeadlineCall {
	c.Call = c.Call.Return(deadline, ok)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderContextDeadlineCall) Do(f func() (time.Time, bool)) *ProviderContextDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderContextDeadlineCall) DoAndReturn(f func() (time.Time, bool)) *ProviderContextDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Done mocks base method.
func (m *MockProviderContext) Done() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockProviderContextMockRecorder) Done() *ProviderContextDoneCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockProviderContext)(nil).Done))
	return &ProviderContextDoneCall{Call: call}
}

// ProviderContextDoneCall wrap *gomock.Call
type ProviderContextDoneCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderContextDoneCall) Return(arg0 <-chan struct{}) *ProviderContextDoneCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderContextDoneCall) Do(f func() <-chan struct{}) *ProviderContextDoneCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderContextDoneCall) DoAndReturn(f func() <-chan struct{}) *ProviderContextDoneCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Err mocks base method.
func (m *MockProviderContext) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockProviderContextMockRecorder) Err() *ProviderContextErrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockProviderContext)(nil).Err))
	return &ProviderContextErrCall{Call: call}
}

// ProviderContextErrCall wrap *gomock.Call
type ProviderContextErrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderContextErrCall) Return(arg0 error) *ProviderContextErrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderContextErrCall) Do(f func() error) *ProviderContextErrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderContextErrCall) DoAndReturn(f func() error) *ProviderContextErrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Log mocks base method.
func (m *MockProviderContext) Log(severity diag.Severity, msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Log", severity, msg)
}

// Log indicates an expected call of Log.
func (mr *MockProviderContextMockRecorder) Log(severity, msg any) *ProviderContextLogCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockProviderContext)(nil).Log), severity, msg)
	return &ProviderContextLogCall{Call: call}
}

// ProviderContextLogCall wrap *gomock.Call
type ProviderContextLogCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderContextLogCall) Return() *ProviderContextLogCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderContextLogCall) Do(f func(diag.Severity, string)) *ProviderContextLogCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderContextLogCall) DoAndReturn(f func(diag.Severity, string)) *ProviderContextLogCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LogStatus mocks base method.
func (m *MockProviderContext) LogStatus(severity diag.Severity, msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogStatus", severity, msg)
}

// LogStatus indicates an expected call of LogStatus.
func (mr *MockProviderContextMockRecorder) LogStatus(severity, msg any) *ProviderContextLogStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogStatus", reflect.TypeOf((*MockProviderContext)(nil).LogStatus), severity, msg)
	return &ProviderContextLogStatusCall{Call: call}
}

// ProviderContextLogStatusCall wrap *gomock.Call
type ProviderContextLogStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderContextLogStatusCall) Return() *ProviderContextLogStatusCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderContextLogStatusCall) Do(f func(diag.Severity, string)) *ProviderContextLogStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderContextLogStatusCall) DoAndReturn(f func(diag.Severity, string)) *ProviderContextLogStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LogStatusf mocks base method.
func (m *MockProviderContext) LogStatusf(severity diag.Severity, msg string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{severity, msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogStatusf", varargs...)
}

// LogStatusf indicates an expected call of LogStatusf.
func (mr *MockProviderContextMockRecorder) LogStatusf(severity, msg any, args ...any) *ProviderContextLogStatusfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{severity, msg}, args...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogStatusf", reflect.TypeOf((*MockProviderContext)(nil).LogStatusf), varargs...)
	return &ProviderContextLogStatusfCall{Call: call}
}

// ProviderContextLogStatusfCall wrap *gomock.Call
type ProviderContextLogStatusfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderContextLogStatusfCall) Return() *ProviderContextLogStatusfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderContextLogStatusfCall) Do(f func(diag.Severity, string, ...any)) *ProviderContextLogStatusfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderContextLogStatusfCall) DoAndReturn(f func(diag.Severity, string, ...any)) *ProviderContextLogStatusfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Logf mocks base method.
func (m *MockProviderContext) Logf(severity diag.Severity, msg string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{severity, msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Logf", varargs...)
}

// Logf indicates an expected call of Logf.
func (mr *MockProviderContextMockRecorder) Logf(severity, msg any, args ...any) *ProviderContextLogfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{severity, msg}, args...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logf", reflect.TypeOf((*MockProviderContext)(nil).Logf), varargs...)
	return &ProviderContextLogfCall{Call: call}
}

// ProviderContextLogfCall wrap *gomock.Call
type ProviderContextLogfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderContextLogfCall) Return() *ProviderContextLogfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderContextLogfCall) Do(f func(diag.Severity, string, ...any)) *ProviderContextLogfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderContextLogfCall) DoAndReturn(f func(diag.Severity, string, ...any)) *ProviderContextLogfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RuntimeInformation mocks base method.
func (m *MockProviderContext) RuntimeInformation() provider.RunInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuntimeInformation")
	ret0, _ := ret[0].(provider.RunInfo)
	return ret0
}

// RuntimeInformation indicates an expected call of RuntimeInformation.
func (mr *MockProviderContextMockRecorder) RuntimeInformation() *ProviderContextRuntimeInformationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuntimeInformation", reflect.TypeOf((*MockProviderContext)(nil).RuntimeInformation))
	return &ProviderContextRuntimeInformationCall{Call: call}
}

// ProviderContextRuntimeInformationCall wrap *gomock.Call
type ProviderContextRuntimeInformationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderContextRuntimeInformationCall) Return(arg0 provider.RunInfo) *ProviderContextRuntimeInformationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderContextRuntimeInformationCall) Do(f func() provider.RunInfo) *ProviderContextRuntimeInformationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderContextRuntimeInformationCall) DoAndReturn(f func() provider.RunInfo) *ProviderContextRuntimeInformationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Value mocks base method.
func (m *MockProviderContext) Value(key any) any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value", key)
	ret0, _ := ret[0].(any)
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockProviderContextMockRecorder) Value(key any) *ProviderContextValueCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockProviderContext)(nil).Value), key)
	return &ProviderContextValueCall{Call: call}
}

// ProviderContextValueCall wrap *gomock.Call
type ProviderContextValueCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderContextValueCall) Return(arg0 any) *ProviderContextValueCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderContextValueCall) Do(f func(any) any) *ProviderContextValueCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderContextValueCall) DoAndReturn(f func(any) any) *ProviderContextValueCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
