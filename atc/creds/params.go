package creds

import (
	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/vars"
)

type Params struct {
	variablesResolver vars.Variables
	rawParams         types.Params
}

func NewParams(variables vars.Variables, params types.Params) Params {
	return Params{
		variablesResolver: variables,
		rawParams:         params,
	}
}

func (p Params) Evaluate() (types.Params, error) {
	var params types.Params
	err := evaluate(p.variablesResolver, p.rawParams, &params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
