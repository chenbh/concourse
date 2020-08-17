// Code generated by counterfeiter. DO NOT EDIT.
package containerserverfakes

import (
	"sync"
	"time"

	"github.com/chenbh/concourse/v6/atc/api/containerserver"
)

type FakeInterceptTimeout struct {
	ChannelStub        func() <-chan time.Time
	channelMutex       sync.RWMutex
	channelArgsForCall []struct {
	}
	channelReturns struct {
		result1 <-chan time.Time
	}
	channelReturnsOnCall map[int]struct {
		result1 <-chan time.Time
	}
	ErrorStub        func() error
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
	}
	errorReturns struct {
		result1 error
	}
	errorReturnsOnCall map[int]struct {
		result1 error
	}
	ResetStub        func()
	resetMutex       sync.RWMutex
	resetArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterceptTimeout) Channel() <-chan time.Time {
	fake.channelMutex.Lock()
	ret, specificReturn := fake.channelReturnsOnCall[len(fake.channelArgsForCall)]
	fake.channelArgsForCall = append(fake.channelArgsForCall, struct {
	}{})
	fake.recordInvocation("Channel", []interface{}{})
	fake.channelMutex.Unlock()
	if fake.ChannelStub != nil {
		return fake.ChannelStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.channelReturns
	return fakeReturns.result1
}

func (fake *FakeInterceptTimeout) ChannelCallCount() int {
	fake.channelMutex.RLock()
	defer fake.channelMutex.RUnlock()
	return len(fake.channelArgsForCall)
}

func (fake *FakeInterceptTimeout) ChannelCalls(stub func() <-chan time.Time) {
	fake.channelMutex.Lock()
	defer fake.channelMutex.Unlock()
	fake.ChannelStub = stub
}

func (fake *FakeInterceptTimeout) ChannelReturns(result1 <-chan time.Time) {
	fake.channelMutex.Lock()
	defer fake.channelMutex.Unlock()
	fake.ChannelStub = nil
	fake.channelReturns = struct {
		result1 <-chan time.Time
	}{result1}
}

func (fake *FakeInterceptTimeout) ChannelReturnsOnCall(i int, result1 <-chan time.Time) {
	fake.channelMutex.Lock()
	defer fake.channelMutex.Unlock()
	fake.ChannelStub = nil
	if fake.channelReturnsOnCall == nil {
		fake.channelReturnsOnCall = make(map[int]struct {
			result1 <-chan time.Time
		})
	}
	fake.channelReturnsOnCall[i] = struct {
		result1 <-chan time.Time
	}{result1}
}

func (fake *FakeInterceptTimeout) Error() error {
	fake.errorMutex.Lock()
	ret, specificReturn := fake.errorReturnsOnCall[len(fake.errorArgsForCall)]
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
	}{})
	fake.recordInvocation("Error", []interface{}{})
	fake.errorMutex.Unlock()
	if fake.ErrorStub != nil {
		return fake.ErrorStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.errorReturns
	return fakeReturns.result1
}

func (fake *FakeInterceptTimeout) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeInterceptTimeout) ErrorCalls(stub func() error) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = stub
}

func (fake *FakeInterceptTimeout) ErrorReturns(result1 error) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = nil
	fake.errorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInterceptTimeout) ErrorReturnsOnCall(i int, result1 error) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = nil
	if fake.errorReturnsOnCall == nil {
		fake.errorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.errorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInterceptTimeout) Reset() {
	fake.resetMutex.Lock()
	fake.resetArgsForCall = append(fake.resetArgsForCall, struct {
	}{})
	fake.recordInvocation("Reset", []interface{}{})
	fake.resetMutex.Unlock()
	if fake.ResetStub != nil {
		fake.ResetStub()
	}
}

func (fake *FakeInterceptTimeout) ResetCallCount() int {
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	return len(fake.resetArgsForCall)
}

func (fake *FakeInterceptTimeout) ResetCalls(stub func()) {
	fake.resetMutex.Lock()
	defer fake.resetMutex.Unlock()
	fake.ResetStub = stub
}

func (fake *FakeInterceptTimeout) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.channelMutex.RLock()
	defer fake.channelMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterceptTimeout) recordInvocation(key string, args []interface{}) {
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

var _ containerserver.InterceptTimeout = new(FakeInterceptTimeout)
