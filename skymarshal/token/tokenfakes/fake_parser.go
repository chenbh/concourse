// Code generated by counterfeiter. DO NOT EDIT.
package tokenfakes

import (
	"sync"
	"time"

	"github.com/chenbh/concourse/v6/skymarshal/token"
)

type FakeParser struct {
	ParseExpiryStub        func(string) (time.Time, error)
	parseExpiryMutex       sync.RWMutex
	parseExpiryArgsForCall []struct {
		arg1 string
	}
	parseExpiryReturns struct {
		result1 time.Time
		result2 error
	}
	parseExpiryReturnsOnCall map[int]struct {
		result1 time.Time
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeParser) ParseExpiry(arg1 string) (time.Time, error) {
	fake.parseExpiryMutex.Lock()
	ret, specificReturn := fake.parseExpiryReturnsOnCall[len(fake.parseExpiryArgsForCall)]
	fake.parseExpiryArgsForCall = append(fake.parseExpiryArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ParseExpiry", []interface{}{arg1})
	fake.parseExpiryMutex.Unlock()
	if fake.ParseExpiryStub != nil {
		return fake.ParseExpiryStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.parseExpiryReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeParser) ParseExpiryCallCount() int {
	fake.parseExpiryMutex.RLock()
	defer fake.parseExpiryMutex.RUnlock()
	return len(fake.parseExpiryArgsForCall)
}

func (fake *FakeParser) ParseExpiryCalls(stub func(string) (time.Time, error)) {
	fake.parseExpiryMutex.Lock()
	defer fake.parseExpiryMutex.Unlock()
	fake.ParseExpiryStub = stub
}

func (fake *FakeParser) ParseExpiryArgsForCall(i int) string {
	fake.parseExpiryMutex.RLock()
	defer fake.parseExpiryMutex.RUnlock()
	argsForCall := fake.parseExpiryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParser) ParseExpiryReturns(result1 time.Time, result2 error) {
	fake.parseExpiryMutex.Lock()
	defer fake.parseExpiryMutex.Unlock()
	fake.ParseExpiryStub = nil
	fake.parseExpiryReturns = struct {
		result1 time.Time
		result2 error
	}{result1, result2}
}

func (fake *FakeParser) ParseExpiryReturnsOnCall(i int, result1 time.Time, result2 error) {
	fake.parseExpiryMutex.Lock()
	defer fake.parseExpiryMutex.Unlock()
	fake.ParseExpiryStub = nil
	if fake.parseExpiryReturnsOnCall == nil {
		fake.parseExpiryReturnsOnCall = make(map[int]struct {
			result1 time.Time
			result2 error
		})
	}
	fake.parseExpiryReturnsOnCall[i] = struct {
		result1 time.Time
		result2 error
	}{result1, result2}
}

func (fake *FakeParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseExpiryMutex.RLock()
	defer fake.parseExpiryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeParser) recordInvocation(key string, args []interface{}) {
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

var _ token.Parser = new(FakeParser)
