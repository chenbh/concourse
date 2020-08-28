// Code generated by counterfeiter. DO NOT EDIT.
package syslogfakes

import (
	"context"
	"sync"

	"github.com/chenbh/concourse/atc/syslog"
)

type FakeDrainer struct {
	RunStub        func(context.Context) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 context.Context
	}
	runReturns struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDrainer) Run(arg1 context.Context) error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Run", []interface{}{arg1})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.runReturns
	return fakeReturns.result1
}

func (fake *FakeDrainer) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeDrainer) RunCalls(stub func(context.Context) error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeDrainer) RunArgsForCall(i int) context.Context {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDrainer) RunReturns(result1 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDrainer) RunReturnsOnCall(i int, result1 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDrainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDrainer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ syslog.Drainer = new(FakeDrainer)
