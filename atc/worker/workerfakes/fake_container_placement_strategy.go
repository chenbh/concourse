// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/chenbh/concourse/v6/atc/worker"
)

type FakeContainerPlacementStrategy struct {
	ChooseStub        func(lager.Logger, []worker.Worker, worker.ContainerSpec) (worker.Worker, error)
	chooseMutex       sync.RWMutex
	chooseArgsForCall []struct {
		arg1 lager.Logger
		arg2 []worker.Worker
		arg3 worker.ContainerSpec
	}
	chooseReturns struct {
		result1 worker.Worker
		result2 error
	}
	chooseReturnsOnCall map[int]struct {
		result1 worker.Worker
		result2 error
	}
	ModifiesActiveTasksStub        func() bool
	modifiesActiveTasksMutex       sync.RWMutex
	modifiesActiveTasksArgsForCall []struct {
	}
	modifiesActiveTasksReturns struct {
		result1 bool
	}
	modifiesActiveTasksReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerPlacementStrategy) Choose(arg1 lager.Logger, arg2 []worker.Worker, arg3 worker.ContainerSpec) (worker.Worker, error) {
	var arg2Copy []worker.Worker
	if arg2 != nil {
		arg2Copy = make([]worker.Worker, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.chooseMutex.Lock()
	ret, specificReturn := fake.chooseReturnsOnCall[len(fake.chooseArgsForCall)]
	fake.chooseArgsForCall = append(fake.chooseArgsForCall, struct {
		arg1 lager.Logger
		arg2 []worker.Worker
		arg3 worker.ContainerSpec
	}{arg1, arg2Copy, arg3})
	fake.recordInvocation("Choose", []interface{}{arg1, arg2Copy, arg3})
	fake.chooseMutex.Unlock()
	if fake.ChooseStub != nil {
		return fake.ChooseStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.chooseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeContainerPlacementStrategy) ChooseCallCount() int {
	fake.chooseMutex.RLock()
	defer fake.chooseMutex.RUnlock()
	return len(fake.chooseArgsForCall)
}

func (fake *FakeContainerPlacementStrategy) ChooseCalls(stub func(lager.Logger, []worker.Worker, worker.ContainerSpec) (worker.Worker, error)) {
	fake.chooseMutex.Lock()
	defer fake.chooseMutex.Unlock()
	fake.ChooseStub = stub
}

func (fake *FakeContainerPlacementStrategy) ChooseArgsForCall(i int) (lager.Logger, []worker.Worker, worker.ContainerSpec) {
	fake.chooseMutex.RLock()
	defer fake.chooseMutex.RUnlock()
	argsForCall := fake.chooseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeContainerPlacementStrategy) ChooseReturns(result1 worker.Worker, result2 error) {
	fake.chooseMutex.Lock()
	defer fake.chooseMutex.Unlock()
	fake.ChooseStub = nil
	fake.chooseReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerPlacementStrategy) ChooseReturnsOnCall(i int, result1 worker.Worker, result2 error) {
	fake.chooseMutex.Lock()
	defer fake.chooseMutex.Unlock()
	fake.ChooseStub = nil
	if fake.chooseReturnsOnCall == nil {
		fake.chooseReturnsOnCall = make(map[int]struct {
			result1 worker.Worker
			result2 error
		})
	}
	fake.chooseReturnsOnCall[i] = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerPlacementStrategy) ModifiesActiveTasks() bool {
	fake.modifiesActiveTasksMutex.Lock()
	ret, specificReturn := fake.modifiesActiveTasksReturnsOnCall[len(fake.modifiesActiveTasksArgsForCall)]
	fake.modifiesActiveTasksArgsForCall = append(fake.modifiesActiveTasksArgsForCall, struct {
	}{})
	fake.recordInvocation("ModifiesActiveTasks", []interface{}{})
	fake.modifiesActiveTasksMutex.Unlock()
	if fake.ModifiesActiveTasksStub != nil {
		return fake.ModifiesActiveTasksStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.modifiesActiveTasksReturns
	return fakeReturns.result1
}

func (fake *FakeContainerPlacementStrategy) ModifiesActiveTasksCallCount() int {
	fake.modifiesActiveTasksMutex.RLock()
	defer fake.modifiesActiveTasksMutex.RUnlock()
	return len(fake.modifiesActiveTasksArgsForCall)
}

func (fake *FakeContainerPlacementStrategy) ModifiesActiveTasksCalls(stub func() bool) {
	fake.modifiesActiveTasksMutex.Lock()
	defer fake.modifiesActiveTasksMutex.Unlock()
	fake.ModifiesActiveTasksStub = stub
}

func (fake *FakeContainerPlacementStrategy) ModifiesActiveTasksReturns(result1 bool) {
	fake.modifiesActiveTasksMutex.Lock()
	defer fake.modifiesActiveTasksMutex.Unlock()
	fake.ModifiesActiveTasksStub = nil
	fake.modifiesActiveTasksReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeContainerPlacementStrategy) ModifiesActiveTasksReturnsOnCall(i int, result1 bool) {
	fake.modifiesActiveTasksMutex.Lock()
	defer fake.modifiesActiveTasksMutex.Unlock()
	fake.ModifiesActiveTasksStub = nil
	if fake.modifiesActiveTasksReturnsOnCall == nil {
		fake.modifiesActiveTasksReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.modifiesActiveTasksReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeContainerPlacementStrategy) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chooseMutex.RLock()
	defer fake.chooseMutex.RUnlock()
	fake.modifiesActiveTasksMutex.RLock()
	defer fake.modifiesActiveTasksMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerPlacementStrategy) recordInvocation(key string, args []interface{}) {
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

var _ worker.ContainerPlacementStrategy = new(FakeContainerPlacementStrategy)
