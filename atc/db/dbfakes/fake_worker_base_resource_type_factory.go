// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/chenbh/concourse/v6/atc/db"
)

type FakeWorkerBaseResourceTypeFactory struct {
	FindStub        func(string, db.Worker) (*db.UsedWorkerBaseResourceType, bool, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		arg1 string
		arg2 db.Worker
	}
	findReturns struct {
		result1 *db.UsedWorkerBaseResourceType
		result2 bool
		result3 error
	}
	findReturnsOnCall map[int]struct {
		result1 *db.UsedWorkerBaseResourceType
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkerBaseResourceTypeFactory) Find(arg1 string, arg2 db.Worker) (*db.UsedWorkerBaseResourceType, bool, error) {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		arg1 string
		arg2 db.Worker
	}{arg1, arg2})
	fake.recordInvocation("Find", []interface{}{arg1, arg2})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.findReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeWorkerBaseResourceTypeFactory) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeWorkerBaseResourceTypeFactory) FindCalls(stub func(string, db.Worker) (*db.UsedWorkerBaseResourceType, bool, error)) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = stub
}

func (fake *FakeWorkerBaseResourceTypeFactory) FindArgsForCall(i int) (string, db.Worker) {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	argsForCall := fake.findArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeWorkerBaseResourceTypeFactory) FindReturns(result1 *db.UsedWorkerBaseResourceType, result2 bool, result3 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 *db.UsedWorkerBaseResourceType
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerBaseResourceTypeFactory) FindReturnsOnCall(i int, result1 *db.UsedWorkerBaseResourceType, result2 bool, result3 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 *db.UsedWorkerBaseResourceType
			result2 bool
			result3 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 *db.UsedWorkerBaseResourceType
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerBaseResourceTypeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorkerBaseResourceTypeFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.WorkerBaseResourceTypeFactory = new(FakeWorkerBaseResourceTypeFactory)
