package creds

import (
	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/vars"
)

type Source struct {
	variablesResolver vars.Variables
	rawSource         types.Source
}

func NewSource(variables vars.Variables, source types.Source) Source {
	return Source{
		variablesResolver: variables,
		rawSource:         source,
	}
}

func (s Source) Evaluate() (types.Source, error) {
	var source types.Source
	err := evaluate(s.variablesResolver, s.rawSource, &source)
	if err != nil {
		return nil, err
	}

	return source, nil
}
