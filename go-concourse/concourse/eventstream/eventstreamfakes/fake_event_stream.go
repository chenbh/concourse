// Code generated by counterfeiter. DO NOT EDIT.
package eventstreamfakes

import (
	"sync"

	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/go-concourse/concourse/eventstream"
)

type FakeEventStream struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	NextEventStub        func() (atc.Event, error)
	nextEventMutex       sync.RWMutex
	nextEventArgsForCall []struct {
	}
	nextEventReturns struct {
		result1 atc.Event
		result2 error
	}
	nextEventReturnsOnCall map[int]struct {
		result1 atc.Event
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventStream) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeEventStream) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeEventStream) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeEventStream) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStream) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStream) NextEvent() (atc.Event, error) {
	fake.nextEventMutex.Lock()
	ret, specificReturn := fake.nextEventReturnsOnCall[len(fake.nextEventArgsForCall)]
	fake.nextEventArgsForCall = append(fake.nextEventArgsForCall, struct {
	}{})
	fake.recordInvocation("NextEvent", []interface{}{})
	fake.nextEventMutex.Unlock()
	if fake.NextEventStub != nil {
		return fake.NextEventStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.nextEventReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEventStream) NextEventCallCount() int {
	fake.nextEventMutex.RLock()
	defer fake.nextEventMutex.RUnlock()
	return len(fake.nextEventArgsForCall)
}

func (fake *FakeEventStream) NextEventCalls(stub func() (atc.Event, error)) {
	fake.nextEventMutex.Lock()
	defer fake.nextEventMutex.Unlock()
	fake.NextEventStub = stub
}

func (fake *FakeEventStream) NextEventReturns(result1 atc.Event, result2 error) {
	fake.nextEventMutex.Lock()
	defer fake.nextEventMutex.Unlock()
	fake.NextEventStub = nil
	fake.nextEventReturns = struct {
		result1 atc.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventStream) NextEventReturnsOnCall(i int, result1 atc.Event, result2 error) {
	fake.nextEventMutex.Lock()
	defer fake.nextEventMutex.Unlock()
	fake.NextEventStub = nil
	if fake.nextEventReturnsOnCall == nil {
		fake.nextEventReturnsOnCall = make(map[int]struct {
			result1 atc.Event
			result2 error
		})
	}
	fake.nextEventReturnsOnCall[i] = struct {
		result1 atc.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventStream) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.nextEventMutex.RLock()
	defer fake.nextEventMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventStream) recordInvocation(key string, args []interface{}) {
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

var _ eventstream.EventStream = new(FakeEventStream)
