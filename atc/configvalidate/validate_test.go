package configvalidate_test

import (
	"encoding/json"
	"github.com/concourse/concourse/atc/types"
	"strings"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/configvalidate"

	// load dummy credential manager
	_ "github.com/concourse/concourse/atc/creds/dummy"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ValidateConfig", func() {
	var (
		config        types.Config
		warnings      []types.ConfigWarning
		errorMessages []string
	)

	BeforeEach(func() {
		config = types.Config{
			Groups: types.GroupConfigs{
				{
					Name:      "some-group",
					Jobs:      []string{"some-job"},
					Resources: []string{"some-resource"},
				},
				{
					Name: "some-other-group",
					Jobs: []string{"some-empty-job"},
				},
			},

			VarSources: types.VarSourceConfigs{},

			Resources: types.ResourceConfigs{
				{
					Name: "some-resource",
					Type: "some-type",
					Source: types.Source{
						"source-config": "some-value",
					},
				},
			},

			ResourceTypes: types.ResourceTypes{
				{
					Name: "some-resource-type",
					Type: "some-type",
					Source: types.Source{
						"source-config": "some-value",
					},
				},
			},

			Jobs: types.JobConfigs{
				{
					Name:   "some-job",
					Public: true,
					Serial: true,
					PlanSequence: []types.Step{
						{
							Config: &types.GetStep{
								Name:     "some-input",
								Resource: "some-resource",
								Params: types.Params{
									"some-param": "some-value",
								},
							},
						},
						{
							Config: &types.LoadVarStep{
								Name: "some-var",
								File: "some-input/some-file.json",
							},
						},
						{
							Config: &types.TaskStep{
								Name:       "some-task",
								Privileged: true,
								ConfigPath: "some/config/path.yml",
							},
						},
						{
							Config: &types.PutStep{
								Name: "some-resource",
								Params: types.Params{
									"some-param": "some-value",
								},
							},
						},
						{
							Config: &types.SetPipelineStep{
								Name: "some-pipeline",
								File: "some-file",
							},
						},
					},
				},
				{
					Name: "some-empty-job",
				},
			},
		}

		atc.EnableAcrossStep = true
	})

	JustBeforeEach(func() {
		warnings, errorMessages = configvalidate.Validate(config)
	})

	Context("when the config is valid", func() {
		It("returns no error", func() {
			Expect(errorMessages).To(HaveLen(0))
		})
	})

	Describe("invalid identifiers", func() {

		Context("when a group has an invalid identifier", func() {
			BeforeEach(func() {
				config.Groups = append(config.Groups, types.GroupConfig{
					Name: "_some-group",
					Jobs: []string{"some-job"},
				})
			})

			It("returns a warning", func() {
				Expect(warnings).To(HaveLen(1))
				Expect(warnings[0].Message).To(ContainSubstring("'_some-group' is not a valid identifier"))
			})
		})

		Context("when a resource has an invalid identifier", func() {
			BeforeEach(func() {
				config.Resources = append(config.Resources, types.ResourceConfig{
					Name: "some_resource",
					Type: "some-type",
					Source: types.Source{
						"source-config": "some-value",
					},
				})
			})

			It("returns a warning", func() {
				Expect(warnings).To(HaveLen(1))
				Expect(warnings[0].Message).To(ContainSubstring("'some_resource' is not a valid identifier"))
			})
		})

		Context("when a resource type has an invalid identifier", func() {
			BeforeEach(func() {
				config.ResourceTypes = append(config.ResourceTypes, types.ResourceType{
					Name: "_some-resource-type",
					Type: "some-type",
					Source: types.Source{
						"source-config": "some-value",
					},
				})
			})

			It("returns a warning", func() {
				Expect(warnings).To(HaveLen(1))
				Expect(warnings[0].Message).To(ContainSubstring("'_some-resource-type' is not a valid identifier"))
			})
		})

		Context("when a var source has an invalid identifier", func() {
			BeforeEach(func() {
				config.VarSources = append(config.VarSources, types.VarSourceConfig{
					Name:   "_some-var-source",
					Type:   "dummy",
					Config: "",
				})
			})

			It("returns a warning", func() {
				Expect(warnings).To(HaveLen(1))
				Expect(warnings[0].Message).To(ContainSubstring("'_some-var-source' is not a valid identifier"))
			})
		})

		Context("when a job has an invalid identifier", func() {
			BeforeEach(func() {
				config.Jobs = append(config.Jobs, types.JobConfig{
					Name: "_some-job",
				})
			})

			It("returns a warning", func() {
				Expect(warnings).To(HaveLen(1))
				Expect(warnings[0].Message).To(ContainSubstring("'_some-job' is not a valid identifier"))
			})
		})

		Context("when a step has an invalid identifier", func() {
			var job types.JobConfig

			BeforeEach(func() {
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.GetStep{
						Name: "_get-step",
					},
				})
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.TaskStep{
						Name: "_task-step",
					},
				})
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.PutStep{
						Name: "_put-step",
					},
				})
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.SetPipelineStep{
						Name: "_set-pipeline-step",
					},
				})
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.LoadVarStep{
						Name: "_load-var-step",
					},
				})

				config.Jobs = append(config.Jobs, job)
			})

			It("returns a warning", func() {
				Expect(warnings).To(HaveLen(5))
				Expect(warnings[0].Message).To(ContainSubstring("'_get-step' is not a valid identifier"))
				Expect(warnings[1].Message).To(ContainSubstring("'_task-step' is not a valid identifier"))
				Expect(warnings[2].Message).To(ContainSubstring("'_put-step' is not a valid identifier"))
				Expect(warnings[3].Message).To(ContainSubstring("'_set-pipeline-step' is not a valid identifier"))
				Expect(warnings[4].Message).To(ContainSubstring("'_load-var-step' is not a valid identifier"))
			})
		})
	})

	Describe("invalid groups", func() {
		Context("when the groups reference a bogus resource", func() {
			BeforeEach(func() {
				config.Groups = append(config.Groups, types.GroupConfig{
					Name:      "bogus",
					Resources: []string{"bogus-resource"},
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid groups:"))
				Expect(errorMessages[0]).To(ContainSubstring("unknown resource 'bogus-resource'"))
			})
		})

		Context("when the groups reference a bogus job", func() {
			BeforeEach(func() {
				config.Groups = append(config.Groups, types.GroupConfig{
					Name: "bogus",
					Jobs: []string{"bogus-job"},
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid groups:"))
				Expect(errorMessages[0]).To(ContainSubstring("unknown job 'bogus-job'"))
			})
		})

		Context("when there are jobs excluded from groups", func() {
			BeforeEach(func() {
				config.Jobs = append(config.Jobs, types.JobConfig{
					Name: "stand-alone-job",
				})
				config.Jobs = append(config.Jobs, types.JobConfig{
					Name: "other-stand-alone-job",
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid groups:"))
				Expect(errorMessages[0]).To(ContainSubstring("job 'stand-alone-job' belongs to no group"))
				Expect(errorMessages[0]).To(ContainSubstring("job 'other-stand-alone-job' belongs to no group"))
			})

		})

		Context("when the groups have two duplicate names", func() {
			BeforeEach(func() {
				config.Groups = append(config.Groups, types.GroupConfig{
					Name: "some-group",
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid groups:"))
				Expect(errorMessages[0]).To(ContainSubstring("'some-group' appears 2 times. Duplicate names are not allowed."))
			})
		})

		Context("when the groups have four duplicate names", func() {
			BeforeEach(func() {
				config.Groups = append(config.Groups, types.GroupConfig{
					Name: "some-group",
				}, types.GroupConfig{
					Name: "some-group",
				}, types.GroupConfig{
					Name: "some-group",
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				errorMessage := strings.Trim(errorMessages[0], "\n")
				errorLines := strings.Split(errorMessage, "\n")
				Expect(errorLines).To(HaveLen(2))
				Expect(errorLines[0]).To(ContainSubstring("invalid groups:"))
				Expect(errorLines[1]).To(ContainSubstring("group 'some-group' appears 4 times. Duplicate names are not allowed."))
			})
		})
	})

	Describe("invalid var sources", func() {
		Context("when a var source type is invalid", func() {
			BeforeEach(func() {
				config.VarSources = append(config.VarSources, types.VarSourceConfig{
					Name:   "some",
					Type:   "some",
					Config: "",
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("unknown credential manager type: some"))
			})
		})

		Context("when config is invalid", func() {
			BeforeEach(func() {
				config.VarSources = append(config.VarSources, types.VarSourceConfig{
					Name:   "some",
					Type:   "dummy",
					Config: "",
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("failed to create credential manager some: invalid dummy credential manager config"))
			})
		})

		Context("when duplicate var source names", func() {
			BeforeEach(func() {
				config.VarSources = append(config.VarSources,
					types.VarSourceConfig{
						Name: "some",
						Type: "dummy",
						Config: map[string]interface{}{
							"vars": map[string]interface{}{"k2": "v2"},
						},
					},
					types.VarSourceConfig{
						Name: "some",
						Type: "dummy",
						Config: map[string]interface{}{
							"vars": map[string]interface{}{"k2": "v2"},
						},
					},
				)
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("duplicate var_source name: some"))
			})
		})

		Context("when var source's dependency cannot be resolved", func() {
			BeforeEach(func() {
				config.VarSources = append(config.VarSources,
					types.VarSourceConfig{
						Name: "s1",
						Type: "dummy",
						Config: map[string]interface{}{
							"vars": map[string]interface{}{"k": "v"},
						},
					},
					types.VarSourceConfig{
						Name: "s2",
						Type: "dummy",
						Config: map[string]interface{}{
							"vars": map[string]interface{}{"k": "((s1:k))"},
						},
					},
					types.VarSourceConfig{
						Name: "s3",
						Type: "dummy",
						Config: map[string]interface{}{
							"vars": map[string]interface{}{"k": "((none:k))"},
						},
					},
				)
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("could not resolve inter-dependent var sources: s3"))
			})
		})

		Context("when var source names are circular", func() {
			BeforeEach(func() {
				config.VarSources = append(config.VarSources,
					types.VarSourceConfig{
						Name: "s1",
						Type: "dummy",
						Config: map[string]interface{}{
							"vars": map[string]interface{}{"k": "((s3:v))"},
						},
					},
					types.VarSourceConfig{
						Name: "s2",
						Type: "dummy",
						Config: map[string]interface{}{
							"vars": map[string]interface{}{"k": "((s1:k))"},
						},
					},
					types.VarSourceConfig{
						Name: "s3",
						Type: "dummy",
						Config: map[string]interface{}{
							"vars": map[string]interface{}{"k": "((s2:k))"},
						},
					},
				)
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("could not resolve inter-dependent var sources: s1, s2, s3"))
			})
		})
	})

	Describe("invalid resources", func() {
		Context("when a resource has no name", func() {
			BeforeEach(func() {
				config.Resources = append(config.Resources, types.ResourceConfig{
					Name: "",
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid resources:"))
				Expect(errorMessages[0]).To(ContainSubstring("resources[1] has no name"))
			})
		})

		Context("when a resource has no type", func() {
			BeforeEach(func() {
				config.Resources = append(config.Resources, types.ResourceConfig{
					Name: "bogus-resource",
					Type: "",
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid resources:"))
				Expect(errorMessages[0]).To(ContainSubstring("resources.bogus-resource has no type"))
			})
		})

		Context("when a resource has no name or type", func() {
			BeforeEach(func() {
				config.Resources = append(config.Resources, types.ResourceConfig{
					Name: "",
					Type: "",
				})
			})

			It("returns an error describing both errors", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid resources:"))
				Expect(errorMessages[0]).To(ContainSubstring("resources[1] has no name"))
				Expect(errorMessages[0]).To(ContainSubstring("resources[1] has no type"))
			})
		})

		Context("when two resources have the same name", func() {
			BeforeEach(func() {
				config.Resources = append(config.Resources, config.Resources...)
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid resources:"))
				Expect(errorMessages[0]).To(ContainSubstring(
					"resources[0] and resources[1] have the same name ('some-resource')",
				))
			})
		})
	})

	Describe("unused resources", func() {
		BeforeEach(func() {
			config = types.Config{
				Resources: types.ResourceConfigs{
					{
						Name: "unused-resource",
						Type: "some-type",
					},
					{
						Name: "get",
						Type: "some-type",
					},
					{
						Name: "get-alias",
						Type: "some-type",
					},
					{
						Name: "resource",
						Type: "some-type",
					},
					{
						Name: "put",
						Type: "some-type",
					},
					{
						Name: "put-alias",
						Type: "some-type",
					},
					{
						Name: "do",
						Type: "some-type",
					},
					{
						Name: "aggregate",
						Type: "some-type",
					},
					{
						Name: "parallel",
						Type: "some-type",
					},
					{
						Name: "abort",
						Type: "some-type",
					},
					{
						Name: "error",
						Type: "some-type",
					},
					{
						Name: "failure",
						Type: "some-type",
					},
					{
						Name: "ensure",
						Type: "some-type",
					},
					{
						Name: "success",
						Type: "some-type",
					},
					{
						Name: "try",
						Type: "some-type",
					},
					{
						Name: "another-job",
						Type: "some-type",
					},
				},

				Jobs: types.JobConfigs{
					{
						Name: "some-job",
						PlanSequence: []types.Step{
							{
								Config: &types.GetStep{
									Name: "get",
								},
							},
							{
								Config: &types.GetStep{
									Name:     "get-alias",
									Resource: "resource",
								},
							},
							{
								Config: &types.PutStep{
									Name: "put",
								},
							},
							{
								Config: &types.PutStep{
									Name:     "put-alias",
									Resource: "resource",
								},
							},
							{
								Config: &types.DoStep{
									Steps: []types.Step{
										{
											Config: &types.GetStep{
												Name: "do",
											},
										},
									},
								},
							},
							{
								Config: &types.AggregateStep{
									Steps: []types.Step{
										{
											Config: &types.GetStep{
												Name: "aggregate",
											},
										},
									},
								},
							},
							{
								Config: &types.InParallelStep{
									Config: types.InParallelConfig{
										Steps: []types.Step{
											{
												Config: &types.GetStep{
													Name: "parallel",
												},
											},
										},
										Limit:    1,
										FailFast: true,
									},
								},
							},
							{
								Config: &types.OnAbortStep{
									Step: &types.TaskStep{
										Name:       "some-task",
										ConfigPath: "some/config/path.yml",
									},
									Hook: types.Step{
										Config: &types.GetStep{
											Name: "abort",
										},
									},
								},
							},
							{
								Config: &types.OnErrorStep{
									Step: &types.TaskStep{
										Name:       "some-task",
										ConfigPath: "some/config/path.yml",
									},
									Hook: types.Step{
										Config: &types.GetStep{
											Name: "error",
										},
									},
								},
							},
							{
								Config: &types.OnFailureStep{
									Step: &types.TaskStep{
										Name:       "some-task",
										ConfigPath: "some/config/path.yml",
									},
									Hook: types.Step{
										Config: &types.GetStep{
											Name: "failure",
										},
									},
								},
							},
							{
								Config: &types.OnSuccessStep{
									Step: &types.TaskStep{
										Name:       "some-task",
										ConfigPath: "some/config/path.yml",
									},
									Hook: types.Step{
										Config: &types.GetStep{
											Name: "success",
										},
									},
								},
							},
							{
								Config: &types.EnsureStep{
									Step: &types.TaskStep{
										Name:       "some-task",
										ConfigPath: "some/config/path.yml",
									},
									Hook: types.Step{
										Config: &types.GetStep{
											Name: "ensure",
										},
									},
								},
							},
							{
								Config: &types.TryStep{
									Step: types.Step{
										Config: &types.GetStep{
											Name: "try",
										},
									},
								},
							},
							{
								Config: &types.TaskStep{
									Name:       "some-task",
									ConfigPath: "some/config/path.yml",
								},
							},
						},
					},
					{
						Name: "another-job",
						PlanSequence: []types.Step{
							{
								Config: &types.GetStep{
									Name: "another-job",
								},
							},
							{
								Config: &types.TaskStep{
									Name:       "some-task",
									ConfigPath: "some/config/path.yml",
								},
							},
						},
					},
				},
			}
		})

		Context("when a resource is not used in any jobs", func() {
			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("resource 'unused-resource' is not used"))
				Expect(errorMessages[0]).To(ContainSubstring("resource 'get-alias' is not used"))
				Expect(errorMessages[0]).To(ContainSubstring("resource 'put-alias' is not used"))
			})
		})
	})

	Describe("invalid resource types", func() {
		Context("when a resource type has no name", func() {
			BeforeEach(func() {
				config.ResourceTypes = append(config.ResourceTypes, types.ResourceType{
					Name: "",
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid resource types:"))
				Expect(errorMessages[0]).To(ContainSubstring("resource_types[1] has no name"))
			})
		})

		Context("when a resource has no type", func() {
			BeforeEach(func() {
				config.ResourceTypes = append(config.ResourceTypes, types.ResourceType{
					Name: "bogus-resource-type",
					Type: "",
				})
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid resource types:"))
				Expect(errorMessages[0]).To(ContainSubstring("resource_types.bogus-resource-type has no type"))
			})
		})

		Context("when a resource has no name or type", func() {
			BeforeEach(func() {
				config.ResourceTypes = append(config.ResourceTypes, types.ResourceType{
					Name: "",
					Type: "",
				})
			})

			It("returns an error describing both errors", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid resource types:"))
				Expect(errorMessages[0]).To(ContainSubstring("resource_types[1] has no name"))
				Expect(errorMessages[0]).To(ContainSubstring("resource_types[1] has no type"))
			})
		})

		Context("when two resource types have the same name", func() {
			BeforeEach(func() {
				config.ResourceTypes = append(config.ResourceTypes, config.ResourceTypes...)
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid resource types:"))
				Expect(errorMessages[0]).To(ContainSubstring("resource_types[0] and resource_types[1] have the same name ('some-resource-type')"))
			})
		})
	})

	Describe("validating a job", func() {
		var job types.JobConfig

		BeforeEach(func() {
			job = types.JobConfig{
				Name: "some-other-job",
			}
			config.Groups = []types.GroupConfig{}
		})

		Context("when a job has no name", func() {
			BeforeEach(func() {
				job.Name = ""
				config.Jobs = append(config.Jobs, job)
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
				Expect(errorMessages[0]).To(ContainSubstring("jobs[2] has no name"))
			})
		})

		Context("when a job has a negative build_logs_to_retain", func() {
			BeforeEach(func() {
				job.BuildLogsToRetain = -1
				config.Jobs = append(config.Jobs, job)
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
				Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job has negative build_logs_to_retain: -1"))
			})
		})

		Context("when a job has duplicate inputs", func() {
			BeforeEach(func() {
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.GetStep{
						Name: "some-resource",
					},
				})
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.GetStep{
						Name: "some-resource",
					},
				})
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.GetStep{
						Name: "some-resource",
					},
				})

				config.Jobs = append(config.Jobs, job)
			})

			It("returns an error for each step", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
				Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[1].get(some-resource): repeated name"))
				Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[2].get(some-resource): repeated name"))
			})
		})

		Context("when a job has duplicate inputs with different resources", func() {
			BeforeEach(func() {
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.GetStep{
						Name:     "some-resource",
						Resource: "a",
					},
				})
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.GetStep{
						Name:     "some-resource",
						Resource: "b",
					},
				})
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.GetStep{
						Name:     "some-resource",
						Resource: "c",
					},
				})

				config.Jobs = append(config.Jobs, job)
			})

			It("returns an error for each step", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
				Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[1].get(some-resource): repeated name"))
				Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[2].get(some-resource): repeated name"))
			})
		})

		Context("when a job gets the same resource multiple times but with different names", func() {
			BeforeEach(func() {
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.GetStep{
						Name:     "a",
						Resource: "some-resource",
					},
				})
				job.PlanSequence = append(job.PlanSequence, types.Step{
					Config: &types.GetStep{
						Name:     "b",
						Resource: "some-resource",
					},
				})

				config.Jobs = append(config.Jobs, job)
			})

			It("returns no errors", func() {
				Expect(errorMessages).To(HaveLen(0))
			})
		})

		Describe("plans", func() {
			Context("when a task plan has neither a config or a path set", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.TaskStep{
							Name: "lol",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].task(lol): must specify either `file:` or `config:`"))
				})
			})

			Context("when a task plan has config path and config specified", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.TaskStep{
							Name:       "lol",
							ConfigPath: "task.yml",
							Config: &types.TaskConfig{
								Params: types.TaskEnv{
									"param1": "value1",
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].task(lol): must specify one of `file:` or `config:`, not both"))
				})
			})

			Context("when a task plan is invalid", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.TaskStep{
							Name: "some-resource",
							Config: &types.TaskConfig{
								Params: types.TaskEnv{
									"param1": "value1",
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].task(some-resource).config: missing 'platform'"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].task(some-resource).config: missing path to executable to run"))
				})
			})

			Context("when a put plan has refers to a resource that does exist", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.PutStep{
							Name: "some-resource",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a get plan has refers to a resource that does not exist", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name: "some-nonexistent-resource",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].get(some-nonexistent-resource): unknown resource 'some-nonexistent-resource'"))
				})
			})

			Context("when a put plan has refers to a resource that does not exist", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.PutStep{
							Name: "some-nonexistent-resource",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].put(some-nonexistent-resource): unknown resource 'some-nonexistent-resource'"))
				})
			})

			Context("when a get plan has a custom name but refers to a resource that does exist", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:     "custom-name",
							Resource: "some-resource",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a get plan has a custom name but refers to a resource that does not exist", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:     "custom-name",
							Resource: "some-missing-resource",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].get(custom-name): unknown resource 'some-missing-resource'"))
				})
			})

			Context("when a put plan has a custom name but refers to a resource that does exist", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.PutStep{
							Name:     "custom-name",
							Resource: "some-resource",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a put plan has a custom name but refers to a resource that does not exist", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.PutStep{
							Name:     "custom-name",
							Resource: "some-missing-resource",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("does return an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].put(custom-name): unknown resource 'some-missing-resource'"))
				})
			})

			Context("when a job success hook refers to a resource that does exist", func() {
				BeforeEach(func() {
					job.OnSuccess = &types.Step{
						Config: &types.GetStep{
							Name: "some-resource",
						},
					}

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a job success hook refers to a resource that does not exist", func() {
				BeforeEach(func() {
					job.OnSuccess = &types.Step{
						Config: &types.GetStep{
							Name: "some-nonexistent-resource",
						},
					}

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.on_success.get(some-nonexistent-resource): unknown resource 'some-nonexistent-resource'"))
				})
			})

			Context("when a job failure hook refers to a resource that does exist", func() {
				BeforeEach(func() {
					job.OnFailure = &types.Step{
						Config: &types.GetStep{
							Name: "some-resource",
						},
					}

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a job failure hook refers to a resource that does not exist", func() {
				BeforeEach(func() {
					job.OnFailure = &types.Step{
						Config: &types.GetStep{
							Name: "some-nonexistent-resource",
						},
					}

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.on_failure.get(some-nonexistent-resource): unknown resource 'some-nonexistent-resource'"))
				})
			})

			Context("when a job error hook refers to a resource that does exist", func() {
				BeforeEach(func() {
					job.OnError = &types.Step{
						Config: &types.GetStep{
							Name: "some-resource",
						},
					}

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a job ensure hook refers to a resource that does not exist", func() {
				BeforeEach(func() {
					job.OnError = &types.Step{
						Config: &types.GetStep{
							Name: "some-nonexistent-resource",
						},
					}

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.on_error.get(some-nonexistent-resource): unknown resource 'some-nonexistent-resource'"))
				})
			})

			Context("when a job abort hook refers to a resource that does exist", func() {
				BeforeEach(func() {
					job.OnAbort = &types.Step{
						Config: &types.GetStep{
							Name: "some-resource",
						},
					}

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a job abort hook refers to a resource that does not exist", func() {
				BeforeEach(func() {
					job.OnAbort = &types.Step{
						Config: &types.GetStep{
							Name: "some-nonexistent-resource",
						},
					}

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.on_abort.get(some-nonexistent-resource): unknown resource 'some-nonexistent-resource'"))
				})
			})

			Context("when a job ensure hook refers to a resource that does exist", func() {
				BeforeEach(func() {
					job.Ensure = &types.Step{
						Config: &types.GetStep{
							Name: "some-resource",
						},
					}

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a job ensure hook refers to a resource that does not exist", func() {
				BeforeEach(func() {
					job.Ensure = &types.Step{
						Config: &types.GetStep{
							Name: "some-nonexistent-resource",
						},
					}

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("invalid jobs:"))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.ensure.get(some-nonexistent-resource): unknown resource 'some-nonexistent-resource'"))
				})
			})

			Context("when a get plan refers to a 'put' resource that exists in another job's hook", func() {
				var (
					job1 types.JobConfig
					job2 types.JobConfig
				)
				BeforeEach(func() {
					job1 = types.JobConfig{
						Name: "job-one",
					}
					job2 = types.JobConfig{
						Name: "job-two",
					}

					job1.PlanSequence = append(job1.PlanSequence, types.Step{
						Config: &types.OnSuccessStep{
							Step: &types.TaskStep{
								Name:       "job-one",
								ConfigPath: "job-one-config-path",
							},
							Hook: types.Step{
								Config: &types.PutStep{
									Name: "some-resource",
								},
							},
						},
					})

					job2.PlanSequence = append(job2.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:   "some-resource",
							Passed: []string{"job-one"},
						},
					})
					config.Jobs = append(config.Jobs, job1, job2)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a get plan refers to a 'get' resource that exists in another job's hook", func() {
				var (
					job1 types.JobConfig
					job2 types.JobConfig
				)
				BeforeEach(func() {
					job1 = types.JobConfig{
						Name: "job-one",
					}
					job2 = types.JobConfig{
						Name: "job-two",
					}

					job1.PlanSequence = append(job1.PlanSequence, types.Step{
						Config: &types.OnSuccessStep{
							Step: &types.TaskStep{
								Name:       "job-one",
								ConfigPath: "job-one-config-path",
							},
							Hook: types.Step{
								Config: &types.GetStep{
									Name: "some-resource",
								},
							},
						},
					})

					job2.PlanSequence = append(job2.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:   "some-resource",
							Passed: []string{"job-one"},
						},
					})
					config.Jobs = append(config.Jobs, job1, job2)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a get plan refers to a 'put' resource that exists in another job's try-step", func() {
				var (
					job1 types.JobConfig
					job2 types.JobConfig
				)
				BeforeEach(func() {
					job1 = types.JobConfig{
						Name: "job-one",
					}
					job2 = types.JobConfig{
						Name: "job-two",
					}

					job1.PlanSequence = append(job1.PlanSequence, types.Step{
						Config: &types.TryStep{
							Step: types.Step{
								Config: &types.PutStep{
									Name: "some-resource",
								},
							},
						},
					})

					job2.PlanSequence = append(job2.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:   "some-resource",
							Passed: []string{"job-one"},
						},
					})
					config.Jobs = append(config.Jobs, job1, job2)

				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a get plan refers to a 'get' resource that exists in another job's try-step", func() {
				var (
					job1 types.JobConfig
					job2 types.JobConfig
				)
				BeforeEach(func() {
					job1 = types.JobConfig{
						Name: "job-one",
					}
					job2 = types.JobConfig{
						Name: "job-two",
					}

					job1.PlanSequence = append(job1.PlanSequence, types.Step{
						Config: &types.TryStep{
							Step: types.Step{
								Config: &types.GetStep{
									Name: "some-resource",
								},
							},
						},
					})

					job2.PlanSequence = append(job2.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:   "some-resource",
							Passed: []string{"job-one"},
						},
					})
					config.Jobs = append(config.Jobs, job1, job2)

				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a plan has an invalid step within an abort", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.OnAbortStep{
							Step: &types.GetStep{
								Name: "some-resource",
							},
							Hook: types.Step{
								Config: &types.PutStep{
									Name:     "custom-name",
									Resource: "some-missing-resource",
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("throws a validation error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].on_abort.put(custom-name): unknown resource 'some-missing-resource'"))
				})
			})

			Context("when a plan has an invalid step within an error", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.OnErrorStep{
							Step: &types.GetStep{
								Name: "some-resource",
							},
							Hook: types.Step{
								Config: &types.PutStep{
									Name:     "custom-name",
									Resource: "some-missing-resource",
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("throws a validation error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].on_error.put(custom-name): unknown resource 'some-missing-resource'"))
				})
			})

			Context("when a plan has an invalid step within an ensure", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.EnsureStep{
							Step: &types.GetStep{
								Name: "some-resource",
							},
							Hook: types.Step{
								Config: &types.PutStep{
									Name:     "custom-name",
									Resource: "some-missing-resource",
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("throws a validation error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].ensure.put(custom-name): unknown resource 'some-missing-resource'"))
				})
			})

			Context("when a plan has an invalid step within a success", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.OnSuccessStep{
							Step: &types.GetStep{
								Name: "some-resource",
							},
							Hook: types.Step{
								Config: &types.PutStep{
									Name:     "custom-name",
									Resource: "some-missing-resource",
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("throws a validation error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].on_success.put(custom-name): unknown resource 'some-missing-resource'"))
				})
			})

			Context("when a plan has an invalid step within a failure", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.OnFailureStep{
							Step: &types.GetStep{
								Name: "some-resource",
							},
							Hook: types.Step{
								Config: &types.PutStep{
									Name:     "custom-name",
									Resource: "some-missing-resource",
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("throws a validation error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].on_failure.put(custom-name): unknown resource 'some-missing-resource'"))
				})
			})

			Context("when a plan has an invalid step within a try", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.TryStep{
							Step: types.Step{
								Config: &types.PutStep{
									Name:     "custom-name",
									Resource: "some-missing-resource",
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("throws a validation error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].try.put(custom-name): unknown resource 'some-missing-resource'"))
				})
			})

			Context("when a plan has an invalid timeout in a step", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.TimeoutStep{
							Step: &types.GetStep{
								Name: "some-resource",
							},
							Duration: "nope",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("throws a validation error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].timeout: invalid duration 'nope'"))
				})
			})

			Context("when a retry plan has a negative attempts number", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.RetryStep{
							Step: &types.PutStep{
								Name: "some-resource",
							},
							Attempts: 0,
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("does return an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].attempts: must be greater than 0"))
				})
			})

			Context("when a set_pipeline step has no file configured", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.SetPipelineStep{
							Name: "other-pipeline",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("does return an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].set_pipeline(other-pipeline): no file specified"))
				})
			})

			Context("when a job's input's passed constraints reference a bogus job", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:   "lol",
							Passed: []string{"bogus-job"},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].get(lol).passed: unknown job 'bogus-job'"))
				})
			})

			Context("when a job's input's passed constraints references a valid job that has the resource as an output", func() {
				BeforeEach(func() {
					config.Jobs[0].PlanSequence = append(config.Jobs[0].PlanSequence, types.Step{
						Config: &types.PutStep{
							Name:     "custom-name",
							Resource: "some-resource",
						},
					})

					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:   "some-resource",
							Passed: []string{"some-job"},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a job's input's passed constraints references a valid job that has the resource as an input", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:   "some-resource",
							Passed: []string{"some-job"},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a job's input's passed constraints references a valid job that has the resource (with a custom name) as an input", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:     "custom-name",
							Resource: "some-resource",
							Passed:   []string{"some-job"},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("does not return an error", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when a job's input's passed constraints references a valid job that does not have the resource as an input or output", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.GetStep{
							Name:   "some-resource",
							Passed: []string{"some-empty-job"},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].get(some-resource).passed: job 'some-empty-job' does not interact with resource 'some-resource'"))
				})
			})

			Context("when a load_var has not defined 'File'", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.LoadVarStep{
							Name: "a-var",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].load_var(a-var): no file specified"))
				})
			})

			Context("when two load_var steps have same name", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.LoadVarStep{
							Name: "a-var",
							File: "file1",
						},
					}, types.Step{
						Config: &types.LoadVarStep{
							Name: "a-var",
							File: "file1",
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[1].load_var(a-var): repeated var name"))
				})
			})

			Context("when a step has unknown fields", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.TaskStep{
							Name:       "task",
							ConfigPath: "some-file",
						},
						UnknownFields: map[string]*json.RawMessage{"bogus": nil},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring(`jobs.some-other-job.plan.do[0]: unknown fields ["bogus"]`))
				})
			})

			Context("when an across step is valid", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.AcrossStep{
							Step: &types.PutStep{
								Name: "some-resource",
							},
							Vars: []types.AcrossVarConfig{
								{
									Var:    "var1",
									Values: []interface{}{"v1", "v2"},
								},
								{
									Var:         "var2",
									MaxInFlight: &types.MaxInFlightConfig{Limit: 2},
									Values:      []interface{}{"v1", "v2"},
								},
								{
									Var:         "var3",
									MaxInFlight: &types.MaxInFlightConfig{All: true},
									Values:      []interface{}{"v1", "v2"},
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("succeeds", func() {
					Expect(errorMessages).To(HaveLen(0))
				})
			})

			Context("when an across step has no vars", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.AcrossStep{
							Step: &types.PutStep{
								Name: "some-resource",
							},
							Vars: []types.AcrossVarConfig{},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].across: no vars specified"))
				})
			})

			Context("when an across step repeats a var name", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.AcrossStep{
							Step: &types.PutStep{
								Name: "some-resource",
							},
							Vars: []types.AcrossVarConfig{
								{
									Var: "var1",
								},
								{
									Var: "var1",
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].across[1]: repeated var name"))
				})
			})

			Context("when an across step shadows a var name from a parent scope", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence,
						types.Step{Config: &types.LoadVarStep{
							Name: "var1",
							File: "unused",
						}},
						types.Step{
							Config: &types.AcrossStep{
								Step: &types.PutStep{
									Name: "some-resource",
								},
								Vars: []types.AcrossVarConfig{
									{
										Var: "var1",
									},
								},
							},
						})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns a warning", func() {
					Expect(errorMessages).To(BeEmpty())
					Expect(warnings).To(HaveLen(1))
					Expect(warnings[0].Message).To(ContainSubstring("jobs.some-other-job.plan.do[1].across[0]: shadows local var 'var1'"))
				})
			})

			Context("when a substep of the across step shadows a var name from a parent scope", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence,
						types.Step{Config: &types.LoadVarStep{
							Name: "a",
							File: "unused",
						}},
						types.Step{
							Config: &types.AcrossStep{
								Step: &types.LoadVarStep{
									Name: "a",
									File: "unused",
								},
								Vars: []types.AcrossVarConfig{
									{
										Var: "b",
									},
								},
							},
						})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns a warning", func() {
					Expect(errorMessages).To(BeEmpty())
					Expect(warnings).To(HaveLen(1))
					Expect(warnings[0].Message).To(ContainSubstring("jobs.some-other-job.plan.do[1].across.load_var(a): shadows local var 'a'"))
				})
			})

			Context("when an across step has a non-positive limit", func() {
				BeforeEach(func() {
					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.AcrossStep{
							Step: &types.PutStep{
								Name: "some-resource",
							},
							Vars: []types.AcrossVarConfig{
								{
									Var:         "var",
									MaxInFlight: &types.MaxInFlightConfig{Limit: 0},
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].across[0].max_in_flight: must be greater than 0"))
				})
			})

			Context("when the across step is not enabled", func() {
				BeforeEach(func() {
					atc.EnableAcrossStep = false

					job.PlanSequence = append(job.PlanSequence, types.Step{
						Config: &types.AcrossStep{
							Step: &types.PutStep{
								Name: "some-resource",
							},
							Vars: []types.AcrossVarConfig{
								{
									Var: "var",
								},
							},
						},
					})

					config.Jobs = append(config.Jobs, job)
				})

				It("returns an error", func() {
					Expect(errorMessages).To(HaveLen(1))
					Expect(errorMessages[0]).To(ContainSubstring("jobs.some-other-job.plan.do[0].across: the across step must be explicitly opted-in to using the `--enable-across-step` flag"))
				})
			})
		})

		Context("when two jobs have the same name", func() {
			BeforeEach(func() {
				config.Jobs = append(config.Jobs, config.Jobs...)
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("jobs[0] and jobs[2] have the same name ('some-job')"))
			})
		})

		Context("when a job has build_log_retention and deprecated build_logs_to_retain", func() {
			BeforeEach(func() {
				config.Jobs[0].BuildLogRetention = &types.BuildLogRetention{
					Builds: 1,
					Days:   1,
				}
				config.Jobs[0].BuildLogsToRetain = 1
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("jobs.some-job can't use both build_log_retention and build_logs_to_retain"))
			})
		})

		Context("when a job has negative build_logs_to_retain", func() {
			BeforeEach(func() {
				config.Jobs[0].BuildLogsToRetain = -1
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("jobs.some-job has negative build_logs_to_retain: -1"))
			})
		})

		Context("when a job has negative build_log_retention values", func() {
			BeforeEach(func() {
				config.Jobs[0].BuildLogRetention = &types.BuildLogRetention{
					Builds: -1,
					Days:   -1,
				}
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("jobs.some-job has negative build_log_retention.builds: -1"))
				Expect(errorMessages[0]).To(ContainSubstring("jobs.some-job has negative build_log_retention.days: -1"))
			})
		})
	})

	Describe("validating display config", func() {
		Context("when the background image is a valid http URL", func() {
			BeforeEach(func() {
				config.Display = &types.DisplayConfig{
					BackgroundImage: "http://example.com/image.jpg",
				}
			})

			It("does not return an error", func() {
				Expect(errorMessages).To(HaveLen(0))
			})
		})

		Context("when the background image is a valid relative URL", func() {
			BeforeEach(func() {
				config.Display = &types.DisplayConfig{
					BackgroundImage: "public/images/image.jpg",
				}
			})

			It("does not return an error", func() {
				Expect(errorMessages).To(HaveLen(0))
			})
		})

		Context("when the background image uses an unsupported scheme", func() {
			BeforeEach(func() {
				config.Display = &types.DisplayConfig{
					BackgroundImage: "data:image/png;base64, iVBORw0KGgoA",
				}
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid display config:"))
				Expect(errorMessages[0]).To(ContainSubstring("background_image scheme must be either http, https or relative"))
			})
		})

		Context("when the background image is an invalid URL", func() {
			BeforeEach(func() {
				config.Display = &types.DisplayConfig{
					BackgroundImage: "://example.com",
				}
			})

			It("returns an error", func() {
				Expect(errorMessages).To(HaveLen(1))
				Expect(errorMessages[0]).To(ContainSubstring("invalid display config:"))
				Expect(errorMessages[0]).To(ContainSubstring("background_image is not a valid URL: ://example.com"))
			})
		})
	})
})
