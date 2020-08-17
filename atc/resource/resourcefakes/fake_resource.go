// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"context"
	"sync"

	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/atc/resource"
	"github.com/chenbh/concourse/v6/atc/runtime"
)

type FakeResource struct {
	CheckStub        func(context.Context, runtime.ProcessSpec, runtime.Runner) ([]atc.Version, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 context.Context
		arg2 runtime.ProcessSpec
		arg3 runtime.Runner
	}
	checkReturns struct {
		result1 []atc.Version
		result2 error
	}
	checkReturnsOnCall map[int]struct {
		result1 []atc.Version
		result2 error
	}
	GetStub        func(context.Context, runtime.ProcessSpec, runtime.Runner) (runtime.VersionResult, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 runtime.ProcessSpec
		arg3 runtime.Runner
	}
	getReturns struct {
		result1 runtime.VersionResult
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 runtime.VersionResult
		result2 error
	}
	PutStub        func(context.Context, runtime.ProcessSpec, runtime.Runner) (runtime.VersionResult, error)
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 context.Context
		arg2 runtime.ProcessSpec
		arg3 runtime.Runner
	}
	putReturns struct {
		result1 runtime.VersionResult
		result2 error
	}
	putReturnsOnCall map[int]struct {
		result1 runtime.VersionResult
		result2 error
	}
	SignatureStub        func() ([]byte, error)
	signatureMutex       sync.RWMutex
	signatureArgsForCall []struct {
	}
	signatureReturns struct {
		result1 []byte
		result2 error
	}
	signatureReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResource) Check(arg1 context.Context, arg2 runtime.ProcessSpec, arg3 runtime.Runner) ([]atc.Version, error) {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 context.Context
		arg2 runtime.ProcessSpec
		arg3 runtime.Runner
	}{arg1, arg2, arg3})
	fake.recordInvocation("Check", []interface{}{arg1, arg2, arg3})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.checkReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResource) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeResource) CheckCalls(stub func(context.Context, runtime.ProcessSpec, runtime.Runner) ([]atc.Version, error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeResource) CheckArgsForCall(i int) (context.Context, runtime.ProcessSpec, runtime.Runner) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeResource) CheckReturns(result1 []atc.Version, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 []atc.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) CheckReturnsOnCall(i int, result1 []atc.Version, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 []atc.Version
			result2 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 []atc.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Get(arg1 context.Context, arg2 runtime.ProcessSpec, arg3 runtime.Runner) (runtime.VersionResult, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 runtime.ProcessSpec
		arg3 runtime.Runner
	}{arg1, arg2, arg3})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResource) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeResource) GetCalls(stub func(context.Context, runtime.ProcessSpec, runtime.Runner) (runtime.VersionResult, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeResource) GetArgsForCall(i int) (context.Context, runtime.ProcessSpec, runtime.Runner) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeResource) GetReturns(result1 runtime.VersionResult, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 runtime.VersionResult
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) GetReturnsOnCall(i int, result1 runtime.VersionResult, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 runtime.VersionResult
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 runtime.VersionResult
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Put(arg1 context.Context, arg2 runtime.ProcessSpec, arg3 runtime.Runner) (runtime.VersionResult, error) {
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 context.Context
		arg2 runtime.ProcessSpec
		arg3 runtime.Runner
	}{arg1, arg2, arg3})
	fake.recordInvocation("Put", []interface{}{arg1, arg2, arg3})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.putReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResource) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeResource) PutCalls(stub func(context.Context, runtime.ProcessSpec, runtime.Runner) (runtime.VersionResult, error)) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = stub
}

func (fake *FakeResource) PutArgsForCall(i int) (context.Context, runtime.ProcessSpec, runtime.Runner) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	argsForCall := fake.putArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeResource) PutReturns(result1 runtime.VersionResult, result2 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 runtime.VersionResult
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) PutReturnsOnCall(i int, result1 runtime.VersionResult, result2 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 runtime.VersionResult
			result2 error
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 runtime.VersionResult
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Signature() ([]byte, error) {
	fake.signatureMutex.Lock()
	ret, specificReturn := fake.signatureReturnsOnCall[len(fake.signatureArgsForCall)]
	fake.signatureArgsForCall = append(fake.signatureArgsForCall, struct {
	}{})
	fake.recordInvocation("Signature", []interface{}{})
	fake.signatureMutex.Unlock()
	if fake.SignatureStub != nil {
		return fake.SignatureStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.signatureReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResource) SignatureCallCount() int {
	fake.signatureMutex.RLock()
	defer fake.signatureMutex.RUnlock()
	return len(fake.signatureArgsForCall)
}

func (fake *FakeResource) SignatureCalls(stub func() ([]byte, error)) {
	fake.signatureMutex.Lock()
	defer fake.signatureMutex.Unlock()
	fake.SignatureStub = stub
}

func (fake *FakeResource) SignatureReturns(result1 []byte, result2 error) {
	fake.signatureMutex.Lock()
	defer fake.signatureMutex.Unlock()
	fake.SignatureStub = nil
	fake.signatureReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) SignatureReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.signatureMutex.Lock()
	defer fake.signatureMutex.Unlock()
	fake.SignatureStub = nil
	if fake.signatureReturnsOnCall == nil {
		fake.signatureReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.signatureReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	fake.signatureMutex.RLock()
	defer fake.signatureMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResource) recordInvocation(key string, args []interface{}) {
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

var _ resource.Resource = new(FakeResource)
