// Code generated by counterfeiter. DO NOT EDIT.
package tokenfakes

import (
	"sync"

	"github.com/chenbh/concourse/v6/atc/db"
	"github.com/chenbh/concourse/v6/skymarshal/token"
)

type FakeClaimsParser struct {
	ParseClaimsStub        func(string) (db.Claims, error)
	parseClaimsMutex       sync.RWMutex
	parseClaimsArgsForCall []struct {
		arg1 string
	}
	parseClaimsReturns struct {
		result1 db.Claims
		result2 error
	}
	parseClaimsReturnsOnCall map[int]struct {
		result1 db.Claims
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClaimsParser) ParseClaims(arg1 string) (db.Claims, error) {
	fake.parseClaimsMutex.Lock()
	ret, specificReturn := fake.parseClaimsReturnsOnCall[len(fake.parseClaimsArgsForCall)]
	fake.parseClaimsArgsForCall = append(fake.parseClaimsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ParseClaims", []interface{}{arg1})
	fake.parseClaimsMutex.Unlock()
	if fake.ParseClaimsStub != nil {
		return fake.ParseClaimsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.parseClaimsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClaimsParser) ParseClaimsCallCount() int {
	fake.parseClaimsMutex.RLock()
	defer fake.parseClaimsMutex.RUnlock()
	return len(fake.parseClaimsArgsForCall)
}

func (fake *FakeClaimsParser) ParseClaimsCalls(stub func(string) (db.Claims, error)) {
	fake.parseClaimsMutex.Lock()
	defer fake.parseClaimsMutex.Unlock()
	fake.ParseClaimsStub = stub
}

func (fake *FakeClaimsParser) ParseClaimsArgsForCall(i int) string {
	fake.parseClaimsMutex.RLock()
	defer fake.parseClaimsMutex.RUnlock()
	argsForCall := fake.parseClaimsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClaimsParser) ParseClaimsReturns(result1 db.Claims, result2 error) {
	fake.parseClaimsMutex.Lock()
	defer fake.parseClaimsMutex.Unlock()
	fake.ParseClaimsStub = nil
	fake.parseClaimsReturns = struct {
		result1 db.Claims
		result2 error
	}{result1, result2}
}

func (fake *FakeClaimsParser) ParseClaimsReturnsOnCall(i int, result1 db.Claims, result2 error) {
	fake.parseClaimsMutex.Lock()
	defer fake.parseClaimsMutex.Unlock()
	fake.ParseClaimsStub = nil
	if fake.parseClaimsReturnsOnCall == nil {
		fake.parseClaimsReturnsOnCall = make(map[int]struct {
			result1 db.Claims
			result2 error
		})
	}
	fake.parseClaimsReturnsOnCall[i] = struct {
		result1 db.Claims
		result2 error
	}{result1, result2}
}

func (fake *FakeClaimsParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseClaimsMutex.RLock()
	defer fake.parseClaimsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClaimsParser) recordInvocation(key string, args []interface{}) {
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

var _ token.ClaimsParser = new(FakeClaimsParser)
