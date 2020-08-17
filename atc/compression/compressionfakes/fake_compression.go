// Code generated by counterfeiter. DO NOT EDIT.
package compressionfakes

import (
	"io"
	"sync"

	"github.com/concourse/baggageclaim"
	"github.com/chenbh/concourse/v6/atc/compression"
)

type FakeCompression struct {
	EncodingStub        func() baggageclaim.Encoding
	encodingMutex       sync.RWMutex
	encodingArgsForCall []struct {
	}
	encodingReturns struct {
		result1 baggageclaim.Encoding
	}
	encodingReturnsOnCall map[int]struct {
		result1 baggageclaim.Encoding
	}
	NewReaderStub        func(io.ReadCloser) (io.ReadCloser, error)
	newReaderMutex       sync.RWMutex
	newReaderArgsForCall []struct {
		arg1 io.ReadCloser
	}
	newReaderReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	newReaderReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCompression) Encoding() baggageclaim.Encoding {
	fake.encodingMutex.Lock()
	ret, specificReturn := fake.encodingReturnsOnCall[len(fake.encodingArgsForCall)]
	fake.encodingArgsForCall = append(fake.encodingArgsForCall, struct {
	}{})
	fake.recordInvocation("Encoding", []interface{}{})
	fake.encodingMutex.Unlock()
	if fake.EncodingStub != nil {
		return fake.EncodingStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.encodingReturns
	return fakeReturns.result1
}

func (fake *FakeCompression) EncodingCallCount() int {
	fake.encodingMutex.RLock()
	defer fake.encodingMutex.RUnlock()
	return len(fake.encodingArgsForCall)
}

func (fake *FakeCompression) EncodingCalls(stub func() baggageclaim.Encoding) {
	fake.encodingMutex.Lock()
	defer fake.encodingMutex.Unlock()
	fake.EncodingStub = stub
}

func (fake *FakeCompression) EncodingReturns(result1 baggageclaim.Encoding) {
	fake.encodingMutex.Lock()
	defer fake.encodingMutex.Unlock()
	fake.EncodingStub = nil
	fake.encodingReturns = struct {
		result1 baggageclaim.Encoding
	}{result1}
}

func (fake *FakeCompression) EncodingReturnsOnCall(i int, result1 baggageclaim.Encoding) {
	fake.encodingMutex.Lock()
	defer fake.encodingMutex.Unlock()
	fake.EncodingStub = nil
	if fake.encodingReturnsOnCall == nil {
		fake.encodingReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.Encoding
		})
	}
	fake.encodingReturnsOnCall[i] = struct {
		result1 baggageclaim.Encoding
	}{result1}
}

func (fake *FakeCompression) NewReader(arg1 io.ReadCloser) (io.ReadCloser, error) {
	fake.newReaderMutex.Lock()
	ret, specificReturn := fake.newReaderReturnsOnCall[len(fake.newReaderArgsForCall)]
	fake.newReaderArgsForCall = append(fake.newReaderArgsForCall, struct {
		arg1 io.ReadCloser
	}{arg1})
	fake.recordInvocation("NewReader", []interface{}{arg1})
	fake.newReaderMutex.Unlock()
	if fake.NewReaderStub != nil {
		return fake.NewReaderStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newReaderReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCompression) NewReaderCallCount() int {
	fake.newReaderMutex.RLock()
	defer fake.newReaderMutex.RUnlock()
	return len(fake.newReaderArgsForCall)
}

func (fake *FakeCompression) NewReaderCalls(stub func(io.ReadCloser) (io.ReadCloser, error)) {
	fake.newReaderMutex.Lock()
	defer fake.newReaderMutex.Unlock()
	fake.NewReaderStub = stub
}

func (fake *FakeCompression) NewReaderArgsForCall(i int) io.ReadCloser {
	fake.newReaderMutex.RLock()
	defer fake.newReaderMutex.RUnlock()
	argsForCall := fake.newReaderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCompression) NewReaderReturns(result1 io.ReadCloser, result2 error) {
	fake.newReaderMutex.Lock()
	defer fake.newReaderMutex.Unlock()
	fake.NewReaderStub = nil
	fake.newReaderReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCompression) NewReaderReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.newReaderMutex.Lock()
	defer fake.newReaderMutex.Unlock()
	fake.NewReaderStub = nil
	if fake.newReaderReturnsOnCall == nil {
		fake.newReaderReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.newReaderReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCompression) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.encodingMutex.RLock()
	defer fake.encodingMutex.RUnlock()
	fake.newReaderMutex.RLock()
	defer fake.newReaderMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCompression) recordInvocation(key string, args []interface{}) {
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

var _ compression.Compression = new(FakeCompression)
