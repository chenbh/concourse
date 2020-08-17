// Code generated by counterfeiter. DO NOT EDIT.
package transportfakes

import (
	"io"
	"net/http"
	"sync"

	"github.com/chenbh/concourse/v6/atc/worker/transport"
	"github.com/tedsuo/rata"
)

type FakeRequestGenerator struct {
	CreateRequestStub        func(string, rata.Params, io.Reader) (*http.Request, error)
	createRequestMutex       sync.RWMutex
	createRequestArgsForCall []struct {
		arg1 string
		arg2 rata.Params
		arg3 io.Reader
	}
	createRequestReturns struct {
		result1 *http.Request
		result2 error
	}
	createRequestReturnsOnCall map[int]struct {
		result1 *http.Request
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRequestGenerator) CreateRequest(arg1 string, arg2 rata.Params, arg3 io.Reader) (*http.Request, error) {
	fake.createRequestMutex.Lock()
	ret, specificReturn := fake.createRequestReturnsOnCall[len(fake.createRequestArgsForCall)]
	fake.createRequestArgsForCall = append(fake.createRequestArgsForCall, struct {
		arg1 string
		arg2 rata.Params
		arg3 io.Reader
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateRequest", []interface{}{arg1, arg2, arg3})
	fake.createRequestMutex.Unlock()
	if fake.CreateRequestStub != nil {
		return fake.CreateRequestStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createRequestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRequestGenerator) CreateRequestCallCount() int {
	fake.createRequestMutex.RLock()
	defer fake.createRequestMutex.RUnlock()
	return len(fake.createRequestArgsForCall)
}

func (fake *FakeRequestGenerator) CreateRequestCalls(stub func(string, rata.Params, io.Reader) (*http.Request, error)) {
	fake.createRequestMutex.Lock()
	defer fake.createRequestMutex.Unlock()
	fake.CreateRequestStub = stub
}

func (fake *FakeRequestGenerator) CreateRequestArgsForCall(i int) (string, rata.Params, io.Reader) {
	fake.createRequestMutex.RLock()
	defer fake.createRequestMutex.RUnlock()
	argsForCall := fake.createRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRequestGenerator) CreateRequestReturns(result1 *http.Request, result2 error) {
	fake.createRequestMutex.Lock()
	defer fake.createRequestMutex.Unlock()
	fake.CreateRequestStub = nil
	fake.createRequestReturns = struct {
		result1 *http.Request
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestGenerator) CreateRequestReturnsOnCall(i int, result1 *http.Request, result2 error) {
	fake.createRequestMutex.Lock()
	defer fake.createRequestMutex.Unlock()
	fake.CreateRequestStub = nil
	if fake.createRequestReturnsOnCall == nil {
		fake.createRequestReturnsOnCall = make(map[int]struct {
			result1 *http.Request
			result2 error
		})
	}
	fake.createRequestReturnsOnCall[i] = struct {
		result1 *http.Request
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRequestMutex.RLock()
	defer fake.createRequestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRequestGenerator) recordInvocation(key string, args []interface{}) {
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

var _ transport.RequestGenerator = new(FakeRequestGenerator)
