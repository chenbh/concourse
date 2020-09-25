package resource

import (
	"context"
	"encoding/json"
	"github.com/concourse/concourse/atc/types"
	"path/filepath"

	"github.com/concourse/concourse/atc/runtime"
)

//go:generate counterfeiter . ResourceFactory
type ResourceFactory interface {
	NewResource(source types.Source, params types.Params, version types.Version) Resource
}

type resourceFactory struct {
}

func NewResourceFactory() ResourceFactory {
	return resourceFactory{}

}
func (rf resourceFactory) NewResource(source types.Source, params types.Params, version types.Version) Resource {
	return &resource{
		Source:  source,
		Params:  params,
		Version: version,
	}
}

//go:generate counterfeiter . Resource

type Resource interface {
	Get(context.Context, runtime.ProcessSpec, runtime.Runner) (runtime.VersionResult, error)
	Put(context.Context, runtime.ProcessSpec, runtime.Runner) (runtime.VersionResult, error)
	Check(context.Context, runtime.ProcessSpec, runtime.Runner) ([]types.Version, error)
	Signature() ([]byte, error)
}

type ResourceType string

type Metadata interface {
	Env() []string
}

func ResourcesDir(suffix string) string {
	return filepath.Join("/tmp", "build", suffix)
}

type resource struct {
	Source  types.Source  `json:"source"`
	Params  types.Params  `json:"params,omitempty"`
	Version types.Version `json:"version,omitempty"`
}

func (resource *resource) Signature() ([]byte, error) {
	return json.Marshal(resource)
}
