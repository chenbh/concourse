// Code generated by counterfeiter. DO NOT EDIT.
package accessorfakes

import (
	"sync"

	"github.com/chenbh/concourse/v6/atc/api/accessor"
	"github.com/chenbh/concourse/v6/atc/db"
)

type FakeTeamFetcher struct {
	GetTeamsStub        func() ([]db.Team, error)
	getTeamsMutex       sync.RWMutex
	getTeamsArgsForCall []struct {
	}
	getTeamsReturns struct {
		result1 []db.Team
		result2 error
	}
	getTeamsReturnsOnCall map[int]struct {
		result1 []db.Team
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTeamFetcher) GetTeams() ([]db.Team, error) {
	fake.getTeamsMutex.Lock()
	ret, specificReturn := fake.getTeamsReturnsOnCall[len(fake.getTeamsArgsForCall)]
	fake.getTeamsArgsForCall = append(fake.getTeamsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetTeams", []interface{}{})
	fake.getTeamsMutex.Unlock()
	if fake.GetTeamsStub != nil {
		return fake.GetTeamsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTeamsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTeamFetcher) GetTeamsCallCount() int {
	fake.getTeamsMutex.RLock()
	defer fake.getTeamsMutex.RUnlock()
	return len(fake.getTeamsArgsForCall)
}

func (fake *FakeTeamFetcher) GetTeamsCalls(stub func() ([]db.Team, error)) {
	fake.getTeamsMutex.Lock()
	defer fake.getTeamsMutex.Unlock()
	fake.GetTeamsStub = stub
}

func (fake *FakeTeamFetcher) GetTeamsReturns(result1 []db.Team, result2 error) {
	fake.getTeamsMutex.Lock()
	defer fake.getTeamsMutex.Unlock()
	fake.GetTeamsStub = nil
	fake.getTeamsReturns = struct {
		result1 []db.Team
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamFetcher) GetTeamsReturnsOnCall(i int, result1 []db.Team, result2 error) {
	fake.getTeamsMutex.Lock()
	defer fake.getTeamsMutex.Unlock()
	fake.GetTeamsStub = nil
	if fake.getTeamsReturnsOnCall == nil {
		fake.getTeamsReturnsOnCall = make(map[int]struct {
			result1 []db.Team
			result2 error
		})
	}
	fake.getTeamsReturnsOnCall[i] = struct {
		result1 []db.Team
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTeamsMutex.RLock()
	defer fake.getTeamsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTeamFetcher) recordInvocation(key string, args []interface{}) {
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

var _ accessor.TeamFetcher = new(FakeTeamFetcher)
