package exec

import (
	"errors"
	"github.com/concourse/concourse/atc/types"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/runtime"
)

func NewVersionSourceFromPlan(getPlan *atc.GetPlan) VersionSource {
	if getPlan.Version != nil {
		return &StaticVersionSource{
			version: *getPlan.Version,
		}
	} else if getPlan.VersionFrom != nil {
		return &PutStepVersionSource{
			planID: *getPlan.VersionFrom,
		}
	} else {
		return &EmptyVersionSource{}
	}
}

type VersionSource interface {
	Version(RunState) (types.Version, error)
}

type StaticVersionSource struct {
	version types.Version
}

func (p *StaticVersionSource) Version(RunState) (types.Version, error) {
	return p.version, nil
}

var ErrPutStepVersionMissing = errors.New("version is missing from put step")

type PutStepVersionSource struct {
	planID atc.PlanID
}

func (p *PutStepVersionSource) Version(state RunState) (types.Version, error) {
	var info runtime.VersionResult
	if !state.Result(p.planID, &info) {
		return types.Version{}, ErrPutStepVersionMissing
	}

	return info.Version, nil
}

type EmptyVersionSource struct{}

func (p *EmptyVersionSource) Version(RunState) (types.Version, error) {
	return types.Version{}, nil
}
