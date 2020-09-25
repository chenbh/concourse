// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"context"
	"github.com/concourse/concourse/atc/types"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/atc/exec/build"
)

type FakeTaskConfigSource struct {
	FetchConfigStub        func(context.Context, lager.Logger, *build.Repository) (types.TaskConfig, error)
	fetchConfigMutex       sync.RWMutex
	fetchConfigArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *build.Repository
	}
	fetchConfigReturns struct {
		result1 types.TaskConfig
		result2 error
	}
	fetchConfigReturnsOnCall map[int]struct {
		result1 types.TaskConfig
		result2 error
	}
	WarningsStub        func() []string
	warningsMutex       sync.RWMutex
	warningsArgsForCall []struct {
	}
	warningsReturns struct {
		result1 []string
	}
	warningsReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskConfigSource) FetchConfig(arg1 context.Context, arg2 lager.Logger, arg3 *build.Repository) (types.TaskConfig, error) {
	fake.fetchConfigMutex.Lock()
	ret, specificReturn := fake.fetchConfigReturnsOnCall[len(fake.fetchConfigArgsForCall)]
	fake.fetchConfigArgsForCall = append(fake.fetchConfigArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *build.Repository
	}{arg1, arg2, arg3})
	fake.recordInvocation("FetchConfig", []interface{}{arg1, arg2, arg3})
	fake.fetchConfigMutex.Unlock()
	if fake.FetchConfigStub != nil {
		return fake.FetchConfigStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fetchConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskConfigSource) FetchConfigCallCount() int {
	fake.fetchConfigMutex.RLock()
	defer fake.fetchConfigMutex.RUnlock()
	return len(fake.fetchConfigArgsForCall)
}

func (fake *FakeTaskConfigSource) FetchConfigCalls(stub func(context.Context, lager.Logger, *build.Repository) (types.TaskConfig, error)) {
	fake.fetchConfigMutex.Lock()
	defer fake.fetchConfigMutex.Unlock()
	fake.FetchConfigStub = stub
}

func (fake *FakeTaskConfigSource) FetchConfigArgsForCall(i int) (context.Context, lager.Logger, *build.Repository) {
	fake.fetchConfigMutex.RLock()
	defer fake.fetchConfigMutex.RUnlock()
	argsForCall := fake.fetchConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTaskConfigSource) FetchConfigReturns(result1 types.TaskConfig, result2 error) {
	fake.fetchConfigMutex.Lock()
	defer fake.fetchConfigMutex.Unlock()
	fake.FetchConfigStub = nil
	fake.fetchConfigReturns = struct {
		result1 types.TaskConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskConfigSource) FetchConfigReturnsOnCall(i int, result1 types.TaskConfig, result2 error) {
	fake.fetchConfigMutex.Lock()
	defer fake.fetchConfigMutex.Unlock()
	fake.FetchConfigStub = nil
	if fake.fetchConfigReturnsOnCall == nil {
		fake.fetchConfigReturnsOnCall = make(map[int]struct {
			result1 types.TaskConfig
			result2 error
		})
	}
	fake.fetchConfigReturnsOnCall[i] = struct {
		result1 types.TaskConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskConfigSource) Warnings() []string {
	fake.warningsMutex.Lock()
	ret, specificReturn := fake.warningsReturnsOnCall[len(fake.warningsArgsForCall)]
	fake.warningsArgsForCall = append(fake.warningsArgsForCall, struct {
	}{})
	fake.recordInvocation("Warnings", []interface{}{})
	fake.warningsMutex.Unlock()
	if fake.WarningsStub != nil {
		return fake.WarningsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.warningsReturns
	return fakeReturns.result1
}

func (fake *FakeTaskConfigSource) WarningsCallCount() int {
	fake.warningsMutex.RLock()
	defer fake.warningsMutex.RUnlock()
	return len(fake.warningsArgsForCall)
}

func (fake *FakeTaskConfigSource) WarningsCalls(stub func() []string) {
	fake.warningsMutex.Lock()
	defer fake.warningsMutex.Unlock()
	fake.WarningsStub = stub
}

func (fake *FakeTaskConfigSource) WarningsReturns(result1 []string) {
	fake.warningsMutex.Lock()
	defer fake.warningsMutex.Unlock()
	fake.WarningsStub = nil
	fake.warningsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeTaskConfigSource) WarningsReturnsOnCall(i int, result1 []string) {
	fake.warningsMutex.Lock()
	defer fake.warningsMutex.Unlock()
	fake.WarningsStub = nil
	if fake.warningsReturnsOnCall == nil {
		fake.warningsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.warningsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeTaskConfigSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchConfigMutex.RLock()
	defer fake.fetchConfigMutex.RUnlock()
	fake.warningsMutex.RLock()
	defer fake.warningsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskConfigSource) recordInvocation(key string, args []interface{}) {
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

var _ exec.TaskConfigSource = new(FakeTaskConfigSource)
