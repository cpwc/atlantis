// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/legacy/events/command (interfaces: Runner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	command "github.com/runatlantis/atlantis/server/legacy/events/command"
	"reflect"
	"time"
)

type MockRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockRunner(options ...pegomock.Option) *MockRunner {
	mock := &MockRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockRunner) Run(ctx *command.Context, cmd *command.Comment) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockRunner().")
	}
	params := []pegomock.Param{ctx, cmd}
	pegomock.GetGenericMockFrom(mock).Invoke("Run", params, []reflect.Type{})
}

func (mock *MockRunner) VerifyWasCalledOnce() *VerifierMockRunner {
	return &VerifierMockRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockRunner) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockRunner {
	return &VerifierMockRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockRunner {
	return &VerifierMockRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockRunner {
	return &VerifierMockRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockRunner struct {
	mock                   *MockRunner
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockRunner) Run(ctx *command.Context, cmd *command.Comment) *MockRunner_Run_OngoingVerification {
	params := []pegomock.Param{ctx, cmd}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", params, verifier.timeout)
	return &MockRunner_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockRunner_Run_OngoingVerification struct {
	mock              *MockRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockRunner_Run_OngoingVerification) GetCapturedArguments() (*command.Context, *command.Comment) {
	ctx, cmd := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], cmd[len(cmd)-1]
}

func (c *MockRunner_Run_OngoingVerification) GetAllCapturedArguments() (_param0 []*command.Context, _param1 []*command.Comment) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*command.Context, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*command.Context)
		}
		_param1 = make([]*command.Comment, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(*command.Comment)
		}
	}
	return
}