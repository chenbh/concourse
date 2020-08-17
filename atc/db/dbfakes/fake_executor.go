// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"database/sql"
	"sync"

	"github.com/chenbh/concourse/v6/atc/db"
)

type FakeExecutor struct {
	ExecStub        func(string, ...interface{}) (sql.Result, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	execReturns struct {
		result1 sql.Result
		result2 error
	}
	execReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExecutor) Exec(arg1 string, arg2 ...interface{}) (sql.Result, error) {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Exec", []interface{}{arg1, arg2})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.execReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeExecutor) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeExecutor) ExecCalls(stub func(string, ...interface{}) (sql.Result, error)) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = stub
}

func (fake *FakeExecutor) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	argsForCall := fake.execArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExecutor) ExecReturns(result1 sql.Result, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeExecutor) ExecReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeExecutor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExecutor) recordInvocation(key string, args []interface{}) {
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

var _ db.Executor = new(FakeExecutor)
