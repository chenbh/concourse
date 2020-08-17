package commands

import (
	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/fly/commands/internal/flaghelpers"
	"github.com/chenbh/concourse/v6/fly/commands/internal/templatehelpers"
	"github.com/chenbh/concourse/v6/fly/commands/internal/validatepipelinehelpers"

	// dynamically registered credential managers
	_ "github.com/chenbh/concourse/v6/atc/creds/conjur"
	_ "github.com/chenbh/concourse/v6/atc/creds/credhub"
	_ "github.com/chenbh/concourse/v6/atc/creds/dummy"
	_ "github.com/chenbh/concourse/v6/atc/creds/kubernetes"
	_ "github.com/chenbh/concourse/v6/atc/creds/secretsmanager"
	_ "github.com/chenbh/concourse/v6/atc/creds/ssm"
	_ "github.com/chenbh/concourse/v6/atc/creds/vault"
)

type ValidatePipelineCommand struct {
	Config atc.PathFlag `short:"c" long:"config" required:"true"        description:"Pipeline configuration file"`
	Strict bool         `short:"s" long:"strict"                        description:"Fail on warnings"`
	Output bool         `short:"o" long:"output"                        description:"Output templated pipeline to stdout"`

	Var     []flaghelpers.VariablePairFlag     `short:"v"  long:"var"       value-name:"[NAME=STRING]"  description:"Specify a string value to set for a variable in the pipeline"`
	YAMLVar []flaghelpers.YAMLVariablePairFlag `short:"y"  long:"yaml-var"  value-name:"[NAME=YAML]"    description:"Specify a YAML value to set for a variable in the pipeline"`

	VarsFrom []atc.PathFlag `short:"l"  long:"load-vars-from"  description:"Variable flag that can be used for filling in template values in configuration from a YAML file"`
}

func (command *ValidatePipelineCommand) Execute(args []string) error {
	yamlTemplate := templatehelpers.NewYamlTemplateWithParams(command.Config, command.VarsFrom, command.Var, command.YAMLVar)
	return validatepipelinehelpers.Validate(yamlTemplate, command.Strict, command.Output)
}
