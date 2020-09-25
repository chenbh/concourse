// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"github.com/concourse/concourse/atc/types"
	"sync"

	"github.com/concourse/concourse/atc/resource"
)

type FakeResourceFactory struct {
	NewResourceStub        func(types.Source, types.Params, types.Version) resource.Resource
	newResourceMutex       sync.RWMutex
	newResourceArgsForCall []struct {
		arg1 types.Source
		arg2 types.Params
		arg3 types.Version
	}
	newResourceReturns struct {
		result1 resource.Resource
	}
	newResourceReturnsOnCall map[int]struct {
		result1 resource.Resource
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceFactory) NewResource(arg1 types.Source, arg2 types.Params, arg3 types.Version) resource.Resource {
	fake.newResourceMutex.Lock()
	ret, specificReturn := fake.newResourceReturnsOnCall[len(fake.newResourceArgsForCall)]
	fake.newResourceArgsForCall = append(fake.newResourceArgsForCall, struct {
		arg1 types.Source
		arg2 types.Params
		arg3 types.Version
	}{arg1, arg2, arg3})
	fake.recordInvocation("NewResource", []interface{}{arg1, arg2, arg3})
	fake.newResourceMutex.Unlock()
	if fake.NewResourceStub != nil {
		return fake.NewResourceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newResourceReturns
	return fakeReturns.result1
}

func (fake *FakeResourceFactory) NewResourceCallCount() int {
	fake.newResourceMutex.RLock()
	defer fake.newResourceMutex.RUnlock()
	return len(fake.newResourceArgsForCall)
}

func (fake *FakeResourceFactory) NewResourceCalls(stub func(types.Source, types.Params, types.Version) resource.Resource) {
	fake.newResourceMutex.Lock()
	defer fake.newResourceMutex.Unlock()
	fake.NewResourceStub = stub
}

func (fake *FakeResourceFactory) NewResourceArgsForCall(i int) (types.Source, types.Params, types.Version) {
	fake.newResourceMutex.RLock()
	defer fake.newResourceMutex.RUnlock()
	argsForCall := fake.newResourceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeResourceFactory) NewResourceReturns(result1 resource.Resource) {
	fake.newResourceMutex.Lock()
	defer fake.newResourceMutex.Unlock()
	fake.NewResourceStub = nil
	fake.newResourceReturns = struct {
		result1 resource.Resource
	}{result1}
}

func (fake *FakeResourceFactory) NewResourceReturnsOnCall(i int, result1 resource.Resource) {
	fake.newResourceMutex.Lock()
	defer fake.newResourceMutex.Unlock()
	fake.NewResourceStub = nil
	if fake.newResourceReturnsOnCall == nil {
		fake.newResourceReturnsOnCall = make(map[int]struct {
			result1 resource.Resource
		})
	}
	fake.newResourceReturnsOnCall[i] = struct {
		result1 resource.Resource
	}{result1}
}

func (fake *FakeResourceFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newResourceMutex.RLock()
	defer fake.newResourceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceFactory) recordInvocation(key string, args []interface{}) {
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

var _ resource.ResourceFactory = new(FakeResourceFactory)
