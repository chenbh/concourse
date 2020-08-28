// Code generated by counterfeiter. DO NOT EDIT.
package accessorfakes

import (
	"sync"

	"github.com/chenbh/concourse/atc/api/accessor"
)

type FakeNotifications struct {
	ListenStub        func(string) (chan bool, error)
	listenMutex       sync.RWMutex
	listenArgsForCall []struct {
		arg1 string
	}
	listenReturns struct {
		result1 chan bool
		result2 error
	}
	listenReturnsOnCall map[int]struct {
		result1 chan bool
		result2 error
	}
	UnlistenStub        func(string, chan bool) error
	unlistenMutex       sync.RWMutex
	unlistenArgsForCall []struct {
		arg1 string
		arg2 chan bool
	}
	unlistenReturns struct {
		result1 error
	}
	unlistenReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNotifications) Listen(arg1 string) (chan bool, error) {
	fake.listenMutex.Lock()
	ret, specificReturn := fake.listenReturnsOnCall[len(fake.listenArgsForCall)]
	fake.listenArgsForCall = append(fake.listenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Listen", []interface{}{arg1})
	fake.listenMutex.Unlock()
	if fake.ListenStub != nil {
		return fake.ListenStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNotifications) ListenCallCount() int {
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	return len(fake.listenArgsForCall)
}

func (fake *FakeNotifications) ListenCalls(stub func(string) (chan bool, error)) {
	fake.listenMutex.Lock()
	defer fake.listenMutex.Unlock()
	fake.ListenStub = stub
}

func (fake *FakeNotifications) ListenArgsForCall(i int) string {
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	argsForCall := fake.listenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNotifications) ListenReturns(result1 chan bool, result2 error) {
	fake.listenMutex.Lock()
	defer fake.listenMutex.Unlock()
	fake.ListenStub = nil
	fake.listenReturns = struct {
		result1 chan bool
		result2 error
	}{result1, result2}
}

func (fake *FakeNotifications) ListenReturnsOnCall(i int, result1 chan bool, result2 error) {
	fake.listenMutex.Lock()
	defer fake.listenMutex.Unlock()
	fake.ListenStub = nil
	if fake.listenReturnsOnCall == nil {
		fake.listenReturnsOnCall = make(map[int]struct {
			result1 chan bool
			result2 error
		})
	}
	fake.listenReturnsOnCall[i] = struct {
		result1 chan bool
		result2 error
	}{result1, result2}
}

func (fake *FakeNotifications) Unlisten(arg1 string, arg2 chan bool) error {
	fake.unlistenMutex.Lock()
	ret, specificReturn := fake.unlistenReturnsOnCall[len(fake.unlistenArgsForCall)]
	fake.unlistenArgsForCall = append(fake.unlistenArgsForCall, struct {
		arg1 string
		arg2 chan bool
	}{arg1, arg2})
	fake.recordInvocation("Unlisten", []interface{}{arg1, arg2})
	fake.unlistenMutex.Unlock()
	if fake.UnlistenStub != nil {
		return fake.UnlistenStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.unlistenReturns
	return fakeReturns.result1
}

func (fake *FakeNotifications) UnlistenCallCount() int {
	fake.unlistenMutex.RLock()
	defer fake.unlistenMutex.RUnlock()
	return len(fake.unlistenArgsForCall)
}

func (fake *FakeNotifications) UnlistenCalls(stub func(string, chan bool) error) {
	fake.unlistenMutex.Lock()
	defer fake.unlistenMutex.Unlock()
	fake.UnlistenStub = stub
}

func (fake *FakeNotifications) UnlistenArgsForCall(i int) (string, chan bool) {
	fake.unlistenMutex.RLock()
	defer fake.unlistenMutex.RUnlock()
	argsForCall := fake.unlistenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNotifications) UnlistenReturns(result1 error) {
	fake.unlistenMutex.Lock()
	defer fake.unlistenMutex.Unlock()
	fake.UnlistenStub = nil
	fake.unlistenReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNotifications) UnlistenReturnsOnCall(i int, result1 error) {
	fake.unlistenMutex.Lock()
	defer fake.unlistenMutex.Unlock()
	fake.UnlistenStub = nil
	if fake.unlistenReturnsOnCall == nil {
		fake.unlistenReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unlistenReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNotifications) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	fake.unlistenMutex.RLock()
	defer fake.unlistenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNotifications) recordInvocation(key string, args []interface{}) {
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

var _ accessor.Notifications = new(FakeNotifications)
