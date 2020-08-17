package dummy

import (
	"github.com/chenbh/concourse/v6/atc/creds"
	"github.com/chenbh/concourse/v6/vars"
)

type SecretsFactory struct {
	vars vars.StaticVariables
}

func NewSecretsFactory(flags []VarFlag) *SecretsFactory {
	vars := vars.StaticVariables{}
	for _, flag := range flags {
		vars[flag.Name] = flag.Value
	}

	return &SecretsFactory{
		vars: vars,
	}
}

func (factory *SecretsFactory) NewSecrets() creds.Secrets {
	return &Secrets{
		StaticVariables: factory.vars,
	}
}
