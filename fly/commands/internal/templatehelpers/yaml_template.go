package templatehelpers

import (
	"fmt"
	"io/ioutil"

	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/fly/commands/internal/flaghelpers"
	"github.com/chenbh/concourse/vars"
	"sigs.k8s.io/yaml"
)

type YamlTemplateWithParams struct {
	filePath               atc.PathFlag
	templateVariablesFiles []atc.PathFlag
	templateVariables      []flaghelpers.VariablePairFlag
	yamlTemplateVariables  []flaghelpers.YAMLVariablePairFlag
}

func NewYamlTemplateWithParams(filePath atc.PathFlag, templateVariablesFiles []atc.PathFlag, templateVariables []flaghelpers.VariablePairFlag, yamlTemplateVariables []flaghelpers.YAMLVariablePairFlag) YamlTemplateWithParams {
	return YamlTemplateWithParams{
		filePath:               filePath,
		templateVariablesFiles: templateVariablesFiles,
		templateVariables:      templateVariables,
		yamlTemplateVariables:  yamlTemplateVariables,
	}
}

func (yamlTemplate YamlTemplateWithParams) Evaluate(
	allowEmpty bool,
	strict bool,
) ([]byte, error) {
	config, err := ioutil.ReadFile(string(yamlTemplate.filePath))
	if err != nil {
		return nil, fmt.Errorf("could not read file: %s", err.Error())
	}

	if strict {
		// We use a generic map here, since templates are not evaluated yet.
		// (else a template string may cause an error when a struct is expected)
		// If we don't check Strict now, then the subsequent steps will mask any
		// duplicate key errors.
		// We should consider being strict throughout the entire stack by default.
		err = yaml.UnmarshalStrict(config, make(map[string]interface{}))
		if err != nil {
			return nil, fmt.Errorf("error parsing yaml before applying templates: %s", err.Error())
		}
	}

	var params []vars.Variables

	// first, we take explicitly specified variables on the command line
	flagVars := vars.StaticVariables{}
	for _, f := range yamlTemplate.templateVariables {
		flagVars[f.Name] = f.Value
	}
	for _, f := range yamlTemplate.yamlTemplateVariables {
		flagVars[f.Name] = f.Value
	}
	params = append(params, flagVars)

	// second, we take all files. with values in the files specified later on command line taking precedence over the
	// same values in the files specified earlier on command line
	for i := len(yamlTemplate.templateVariablesFiles) - 1; i >= 0; i-- {
		path := yamlTemplate.templateVariablesFiles[i]
		templateVars, err := ioutil.ReadFile(string(path))
		if err != nil {
			return nil, fmt.Errorf("could not read template variables file (%s): %s", string(path), err.Error())
		}

		var staticVars vars.StaticVariables
		err = yaml.Unmarshal(templateVars, &staticVars)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal template variables (%s): %s", string(path), err.Error())
		}

		params = append(params, staticVars)
	}

	evaluatedConfig, err := vars.NewTemplateResolver(config, params).Resolve(false, allowEmpty)
	if err != nil {
		return nil, err
	}

	return evaluatedConfig, nil
}
