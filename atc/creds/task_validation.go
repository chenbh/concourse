package creds

import (
	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/vars"
)

type TaskEnvValidator struct {
	variablesResolver vars.Variables
	rawTaskEnv        types.TaskEnv
}

func NewTaskEnvValidator(variables vars.Variables, params types.TaskEnv) TaskEnvValidator {
	return TaskEnvValidator{
		variablesResolver: variables,
		rawTaskEnv:        params,
	}
}

func (s TaskEnvValidator) Validate() error {
	var params types.TaskEnv
	return evaluate(s.variablesResolver, s.rawTaskEnv, &params)
}

type TaskVarsValidator struct {
	variablesResolver vars.Variables
	rawTaskVars       types.Params
}

func NewTaskVarsValidator(variables vars.Variables, taskVars types.Params) TaskVarsValidator {
	return TaskVarsValidator{
		variablesResolver: variables,
		rawTaskVars:       taskVars,
	}
}

func (s TaskVarsValidator) Validate() error {
	var params types.Params
	return evaluate(s.variablesResolver, s.rawTaskVars, &params)
}
