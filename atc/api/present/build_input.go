package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func BuildInput(input db.BuildInput, config types.JobInputParams, resource db.Resource) types.BuildInput {
	return types.BuildInput{
		Name:     input.Name,
		Resource: resource.Name(),
		Type:     resource.Type(),
		Source:   resource.Source(),
		Params:   config.Params,
		Version:  types.Version(input.Version),
		Tags:     config.Tags,
	}
}
