// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/handlers (interfaces: ProjectCommandOutputHandler)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockProjectCommandOutputHandler struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectCommandOutputHandler(options ...pegomock.Option) *MockProjectCommandOutputHandler {
	mock := &MockProjectCommandOutputHandler{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockProjectCommandOutputHandler) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockProjectCommandOutputHandler) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockProjectCommandOutputHandler) Clear(ctx models.ProjectCommandContext) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{ctx}
	pegomock.GetGenericMockFrom(mock).Invoke("Clear", params, []reflect.Type{})
}

func (mock *MockProjectCommandOutputHandler) Send(ctx models.ProjectCommandContext, msg string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{ctx, msg}
	pegomock.GetGenericMockFrom(mock).Invoke("Send", params, []reflect.Type{})
}

func (mock *MockProjectCommandOutputHandler) Receive(projectInfo string, receiver chan string, callback func(string) error) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{projectInfo, receiver, callback}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Receive", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockProjectCommandOutputHandler) Handle() {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{}
	pegomock.GetGenericMockFrom(mock).Invoke("Handle", params, []reflect.Type{})
}

func (mock *MockProjectCommandOutputHandler) SetJobURLWithStatus(ctx models.ProjectCommandContext, cmdName models.CommandName, status models.CommitStatus) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{ctx, cmdName, status}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SetJobURLWithStatus", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockProjectCommandOutputHandler) VerifyWasCalledOnce() *VerifierMockProjectCommandOutputHandler {
	return &VerifierMockProjectCommandOutputHandler{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockProjectCommandOutputHandler) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockProjectCommandOutputHandler {
	return &VerifierMockProjectCommandOutputHandler{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockProjectCommandOutputHandler) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockProjectCommandOutputHandler {
	return &VerifierMockProjectCommandOutputHandler{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockProjectCommandOutputHandler) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockProjectCommandOutputHandler {
	return &VerifierMockProjectCommandOutputHandler{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockProjectCommandOutputHandler struct {
	mock                   *MockProjectCommandOutputHandler
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockProjectCommandOutputHandler) Clear(ctx models.ProjectCommandContext) *MockProjectCommandOutputHandler_Clear_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Clear", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_Clear_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_Clear_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_Clear_OngoingVerification) GetCapturedArguments() models.ProjectCommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandOutputHandler_Clear_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandOutputHandler) Send(ctx models.ProjectCommandContext, msg string) *MockProjectCommandOutputHandler_Send_OngoingVerification {
	params := []pegomock.Param{ctx, msg}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Send", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_Send_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_Send_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_Send_OngoingVerification) GetCapturedArguments() (models.ProjectCommandContext, string) {
	ctx, msg := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], msg[len(msg)-1]
}

func (c *MockProjectCommandOutputHandler_Send_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandOutputHandler) Receive(projectInfo string, receiver chan string, callback func(string) error) *MockProjectCommandOutputHandler_Receive_OngoingVerification {
	params := []pegomock.Param{projectInfo, receiver, callback}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Receive", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_Receive_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_Receive_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_Receive_OngoingVerification) GetCapturedArguments() (string, chan string, func(string) error) {
	projectInfo, receiver, callback := c.GetAllCapturedArguments()
	return projectInfo[len(projectInfo)-1], receiver[len(receiver)-1], callback[len(callback)-1]
}

func (c *MockProjectCommandOutputHandler_Receive_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []chan string, _param2 []func(string) error) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]chan string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(chan string)
		}
		_param2 = make([]func(string) error, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(func(string) error)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandOutputHandler) Handle() *MockProjectCommandOutputHandler_Handle_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Handle", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_Handle_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_Handle_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_Handle_OngoingVerification) GetCapturedArguments() {
}

func (c *MockProjectCommandOutputHandler_Handle_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockProjectCommandOutputHandler) SetJobURLWithStatus(ctx models.ProjectCommandContext, cmdName models.CommandName, status models.CommitStatus) *MockProjectCommandOutputHandler_SetJobURLWithStatus_OngoingVerification {
	params := []pegomock.Param{ctx, cmdName, status}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetJobURLWithStatus", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_SetJobURLWithStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_SetJobURLWithStatus_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_SetJobURLWithStatus_OngoingVerification) GetCapturedArguments() (models.ProjectCommandContext, models.CommandName, models.CommitStatus) {
	ctx, cmdName, status := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], cmdName[len(cmdName)-1], status[len(status)-1]
}

func (c *MockProjectCommandOutputHandler_SetJobURLWithStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext, _param1 []models.CommandName, _param2 []models.CommitStatus) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
		_param1 = make([]models.CommandName, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.CommandName)
		}
		_param2 = make([]models.CommitStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
	}
	return
}
