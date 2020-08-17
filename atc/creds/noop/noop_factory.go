package noop

import "github.com/chenbh/concourse/v6/atc/creds"

type noopFactory struct{}

func NewNoopFactory() *noopFactory {
	return &noopFactory{}
}

func (*noopFactory) NewSecrets() creds.Secrets {
	return &Noop{}
}
