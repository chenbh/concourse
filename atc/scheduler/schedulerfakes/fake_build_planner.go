// Code generated by counterfeiter. DO NOT EDIT.
package schedulerfakes

import (
	"sync"

	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/atc/db"
	"github.com/chenbh/concourse/v6/atc/scheduler"
)

type FakeBuildPlanner struct {
	CreateStub        func(atc.StepConfig, db.SchedulerResources, atc.VersionedResourceTypes, []db.BuildInput) (atc.Plan, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 atc.StepConfig
		arg2 db.SchedulerResources
		arg3 atc.VersionedResourceTypes
		arg4 []db.BuildInput
	}
	createReturns struct {
		result1 atc.Plan
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 atc.Plan
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildPlanner) Create(arg1 atc.StepConfig, arg2 db.SchedulerResources, arg3 atc.VersionedResourceTypes, arg4 []db.BuildInput) (atc.Plan, error) {
	var arg4Copy []db.BuildInput
	if arg4 != nil {
		arg4Copy = make([]db.BuildInput, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 atc.StepConfig
		arg2 db.SchedulerResources
		arg3 atc.VersionedResourceTypes
		arg4 []db.BuildInput
	}{arg1, arg2, arg3, arg4Copy})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4Copy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBuildPlanner) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeBuildPlanner) CreateCalls(stub func(atc.StepConfig, db.SchedulerResources, atc.VersionedResourceTypes, []db.BuildInput) (atc.Plan, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeBuildPlanner) CreateArgsForCall(i int) (atc.StepConfig, db.SchedulerResources, atc.VersionedResourceTypes, []db.BuildInput) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeBuildPlanner) CreateReturns(result1 atc.Plan, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 atc.Plan
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildPlanner) CreateReturnsOnCall(i int, result1 atc.Plan, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 atc.Plan
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 atc.Plan
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildPlanner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildPlanner) recordInvocation(key string, args []interface{}) {
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

var _ scheduler.BuildPlanner = new(FakeBuildPlanner)
