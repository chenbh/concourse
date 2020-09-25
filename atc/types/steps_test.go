package types_test

import (
	"encoding/json"
	"strings"

	"github.com/concourse/concourse/atc/types"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"sigs.k8s.io/yaml"
)

type StepsSuite struct {
	suite.Suite
	*require.Assertions
}

type StepTest struct {
	Title string

	ConfigYAML string
	StepConfig types.StepConfig

	UnknownFields map[string]*json.RawMessage
	Err           string
}

var factoryTests = []StepTest{
	{
		Title: "get step",
		ConfigYAML: `
			get: some-name
			resource: some-resource
			params: {some: params}
			version: {some: version}
			tags: [tag-1, tag-2]
		`,
		StepConfig: &types.GetStep{
			Name:     "some-name",
			Resource: "some-resource",
			Params:   types.Params{"some": "params"},
			Version:  &types.VersionConfig{Pinned: types.Version{"some": "version"}},
			Tags:     []string{"tag-1", "tag-2"},
		},
	},
	{
		Title: "put step",

		ConfigYAML: `
			put: some-name
			resource: some-resource
			params: {some: params}
			tags: [tag-1, tag-2]
			inputs: all
			get_params: {some: get-params}
		`,
		StepConfig: &types.PutStep{
			Name:      "some-name",
			Resource:  "some-resource",
			Params:    types.Params{"some": "params"},
			Tags:      []string{"tag-1", "tag-2"},
			Inputs:    &types.InputsConfig{All: true},
			GetParams: types.Params{"some": "get-params"},
		},
	},
	{
		Title: "task step",

		ConfigYAML: `
			task: some-task
			privileged: true
			config:
			  platform: linux
			  run: {path: hello}
			file: some-task-file
			vars: {some: vars}
			params: {SOME: PARAMS}
			tags: [tag-1, tag-2]
			input_mapping: {generic: specific}
			output_mapping: {specific: generic}
			image: some-image
		`,

		StepConfig: &types.TaskStep{
			Name:       "some-task",
			Privileged: true,
			Config: &types.TaskConfig{
				Platform: "linux",
				Run:      types.TaskRunConfig{Path: "hello"},
			},
			ConfigPath:        "some-task-file",
			Vars:              types.Params{"some": "vars"},
			Params:            types.TaskEnv{"SOME": "PARAMS"},
			Tags:              []string{"tag-1", "tag-2"},
			InputMapping:      map[string]string{"generic": "specific"},
			OutputMapping:     map[string]string{"specific": "generic"},
			ImageArtifactName: "some-image",
		},
	},
	{
		Title: "task step with non-string params",

		ConfigYAML: `
			task: some-task
			file: some-task-file
			params:
			  NUMBER: 42
			  FLOAT: 1.5
			  BOOL: yes
			  OBJECT: {foo: bar}
		`,

		StepConfig: &types.TaskStep{
			Name:       "some-task",
			ConfigPath: "some-task-file",
			Params: types.TaskEnv{
				"NUMBER": "42",
				"FLOAT":  "1.5",
				"BOOL":   "true",
				"OBJECT": `{"foo":"bar"}`,
			},
		},
	},
	{
		Title: "set_pipeline step",

		ConfigYAML: `
			set_pipeline: some-pipeline
			file: some-pipeline-file
			vars: {some: vars}
			var_files: [file-1, file-2]
		`,

		StepConfig: &types.SetPipelineStep{
			Name:     "some-pipeline",
			File:     "some-pipeline-file",
			Vars:     types.Params{"some": "vars"},
			VarFiles: []string{"file-1", "file-2"},
		},
	},
	{
		Title: "load_var step",

		ConfigYAML: `
			load_var: some-var
			file: some-var-file
			format: raw
			reveal: true
		`,

		StepConfig: &types.LoadVarStep{
			Name:   "some-var",
			File:   "some-var-file",
			Format: "raw",
			Reveal: true,
		},
	},
	{
		Title: "try step",

		ConfigYAML: `
			try:
			  load_var: some-var
			  file: some-file
		`,

		StepConfig: &types.TryStep{
			Step: types.Step{
				Config: &types.LoadVarStep{
					Name: "some-var",
					File: "some-file",
				},
			},
		},
	},
	{
		Title: "do step",

		ConfigYAML: `
			do:
			- load_var: some-var
			  file: some-file
			- load_var: some-other-var
			  file: some-other-file
		`,

		StepConfig: &types.DoStep{
			Steps: []types.Step{
				{
					Config: &types.LoadVarStep{
						Name: "some-var",
						File: "some-file",
					},
				},
				{
					Config: &types.LoadVarStep{
						Name: "some-other-var",
						File: "some-other-file",
					},
				},
			},
		},
	},
	{
		Title: "in_parallel step with simple list",

		ConfigYAML: `
			in_parallel:
			- load_var: some-var
			  file: some-file
			- load_var: some-other-var
			  file: some-other-file
		`,

		StepConfig: &types.InParallelStep{
			Config: types.InParallelConfig{
				Steps: []types.Step{
					{
						Config: &types.LoadVarStep{
							Name: "some-var",
							File: "some-file",
						},
					},
					{
						Config: &types.LoadVarStep{
							Name: "some-other-var",
							File: "some-other-file",
						},
					},
				},
			},
		},
	},
	{
		Title: "in_parallel step with config",

		ConfigYAML: `
			in_parallel:
			  steps:
			  - load_var: some-var
			    file: some-file
			  - load_var: some-other-var
			    file: some-other-file
			  limit: 3
			  fail_fast: true
		`,

		StepConfig: &types.InParallelStep{
			Config: types.InParallelConfig{
				Steps: []types.Step{
					{
						Config: &types.LoadVarStep{
							Name: "some-var",
							File: "some-file",
						},
					},
					{
						Config: &types.LoadVarStep{
							Name: "some-other-var",
							File: "some-other-file",
						},
					},
				},
				Limit:    3,
				FailFast: true,
			},
		},
	},
	{
		Title: "aggregate step",

		ConfigYAML: `
			aggregate:
			- load_var: some-var
			  file: some-file
			- load_var: some-other-var
			  file: some-other-file
		`,

		StepConfig: &types.AggregateStep{
			Steps: []types.Step{
				{
					Config: &types.LoadVarStep{
						Name: "some-var",
						File: "some-file",
					},
				},
				{
					Config: &types.LoadVarStep{
						Name: "some-other-var",
						File: "some-other-file",
					},
				},
			},
		},
	},
	{
		Title: "across step",

		ConfigYAML: `
			load_var: some-var
			file: some-file
			across:
			- var: var1
			  values: [1, 2, 3]
			  max_in_flight: 3
			- var: var2
			  values: ["a", "b"]
			  max_in_flight: all
			- var: var3
			  values: [{a: "a", b: "b"}]
			fail_fast: true
		`,

		StepConfig: &types.AcrossStep{
			Step: &types.LoadVarStep{
				Name: "some-var",
				File: "some-file",
			},
			Vars: []types.AcrossVarConfig{
				{
					Var:         "var1",
					Values:      []interface{}{float64(1), float64(2), float64(3)},
					MaxInFlight: &types.MaxInFlightConfig{Limit: 3},
				},
				{
					Var:         "var2",
					Values:      []interface{}{"a", "b"},
					MaxInFlight: &types.MaxInFlightConfig{All: true},
				},
				{
					Var:    "var3",
					Values: []interface{}{map[string]interface{}{"a": "a", "b": "b"}},
				},
			},
			FailFast: true,
		},
	},
	{
		Title: "across step with invalid field",

		ConfigYAML: `
			load_var: some-var
			file: some-file
			across:
			- var: var1
			  values: [1, 2, 3]
			  bogus_field: lol what ru gonna do about it 
		`,

		Err: `error unmarshaling JSON: while decoding JSON: malformed across step: json: unknown field "bogus_field"`,
	},
	{
		Title: "across step with invalid max_in_flight",

		ConfigYAML: `
			load_var: some-var
			file: some-file
			across:
			- var: var1
			  values: [1, 2, 3]
			  max_in_flight: some
		`,

		Err: `error unmarshaling JSON: while decoding JSON: malformed across step: invalid max_in_flight "some"`,
	},
	{
		Title: "timeout modifier",

		ConfigYAML: `
			load_var: some-var
			file: some-file
			timeout: 1h
		`,

		StepConfig: &types.TimeoutStep{
			Step: &types.LoadVarStep{
				Name: "some-var",
				File: "some-file",
			},
			Duration: "1h",
		},
	},
	{
		Title: "attempts modifier",

		ConfigYAML: `
			load_var: some-var
			file: some-file
			attempts: 3
		`,

		StepConfig: &types.RetryStep{
			Step: &types.LoadVarStep{
				Name: "some-var",
				File: "some-file",
			},
			Attempts: 3,
		},
	},
	{
		Title: "precedence of all hooks and modifiers",

		ConfigYAML: `
			load_var: some-var
			file: some-file
			timeout: 1h
			attempts: 3
			across:
			- var: version
			  values: [v1, v2, v3]
			on_success:
			  load_var: success-var
			  file: success-file
			on_failure:
			  load_var: failure-var
			  file: failure-file
			on_abort:
			  load_var: abort-var
			  file: abort-file
			on_error:
			  load_var: error-var
			  file: error-file
			ensure:
			  load_var: ensure-var
			  file: ensure-file
		`,

		StepConfig: &types.EnsureStep{
			Step: &types.OnErrorStep{
				Step: &types.OnAbortStep{
					Step: &types.OnFailureStep{
						Step: &types.OnSuccessStep{
							Step: &types.AcrossStep{
								Step: &types.RetryStep{
									Step: &types.TimeoutStep{
										Step: &types.LoadVarStep{
											Name: "some-var",
											File: "some-file",
										},
										Duration: "1h",
									},
									Attempts: 3,
								},
								Vars: []types.AcrossVarConfig{
									{
										Var:    "version",
										Values: []interface{}{"v1", "v2", "v3"},
									},
								},
							},
							Hook: types.Step{
								Config: &types.LoadVarStep{
									Name: "success-var",
									File: "success-file",
								},
							},
						},
						Hook: types.Step{
							Config: &types.LoadVarStep{
								Name: "failure-var",
								File: "failure-file",
							},
						},
					},
					Hook: types.Step{
						Config: &types.LoadVarStep{
							Name: "abort-var",
							File: "abort-file",
						},
					},
				},
				Hook: types.Step{
					Config: &types.LoadVarStep{
						Name: "error-var",
						File: "error-file",
					},
				},
			},
			Hook: types.Step{
				Config: &types.LoadVarStep{
					Name: "ensure-var",
					File: "ensure-file",
				},
			},
		},
	},
	{
		Title: "unknown field with get step",

		ConfigYAML: `
			get: some-name
			bogus: foo
		`,

		StepConfig: &types.GetStep{
			Name: "some-name",
		},

		UnknownFields: map[string]*json.RawMessage{"bogus": rawMessage(`"foo"`)},
	},
	{
		Title: "multiple steps defined",

		ConfigYAML: `
			put: some-name
			get: some-other-name
		`,

		StepConfig: &types.PutStep{
			Name: "some-name",
		},

		UnknownFields: map[string]*json.RawMessage{"get": rawMessage(`"some-other-name"`)},
	},
	{
		Title: "step cannot contain only modifiers",

		ConfigYAML: `
			attempts: 2
		`,

		StepConfig: &types.RetryStep{
			Attempts: 2,
		},

		Err: "no core step type declared (e.g. get, put, task, etc.)",
	},
}

func (test StepTest) Run(s *StepsSuite) {
	cleanIndents := strings.ReplaceAll(test.ConfigYAML, "\t", "")

	var step types.Step
	actualErr := yaml.Unmarshal([]byte(cleanIndents), &step)
	if test.Err != "" {
		s.Contains(actualErr.Error(), test.Err)
		return
	} else {
		s.NoError(actualErr)
	}

	s.Equal(test.StepConfig, step.Config)
	s.Equal(test.UnknownFields, step.UnknownFields)

	remarshalled, err := json.Marshal(step)
	s.NoError(err)

	var reStep types.Step
	err = yaml.Unmarshal(remarshalled, &reStep)
	s.NoError(err)

	s.Equal(test.StepConfig, reStep.Config)
}

func (s *StepsSuite) TestFactory() {
	for _, test := range factoryTests {
		s.Run(test.Title, func() {
			test.Run(s)
		})
	}
}

func rawMessage(s string) *json.RawMessage {
	raw := json.RawMessage(s)
	return &raw
}
