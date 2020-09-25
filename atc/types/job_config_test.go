package types_test

import (
	. "github.com/concourse/concourse/atc/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JobConfig", func() {
	Describe("MaxInFlight", func() {
		It("returns the raw MaxInFlight if set", func() {
			jobConfig := JobConfig{
				RawMaxInFlight: 42,
			}

			Expect(jobConfig.MaxInFlight()).To(Equal(42))
		})

		It("returns 1 if Serial is true or SerialGroups has items in it", func() {
			jobConfig := JobConfig{
				Serial:       true,
				SerialGroups: []string{},
			}

			Expect(jobConfig.MaxInFlight()).To(Equal(1))

			jobConfig.SerialGroups = []string{
				"one",
			}
			Expect(jobConfig.MaxInFlight()).To(Equal(1))

			jobConfig.Serial = false
			Expect(jobConfig.MaxInFlight()).To(Equal(1))
		})

		It("returns 1 if Serial is true or SerialGroups has items in it, even if raw MaxInFlight is set", func() {
			jobConfig := JobConfig{
				Serial:         true,
				SerialGroups:   []string{},
				RawMaxInFlight: 3,
			}

			Expect(jobConfig.MaxInFlight()).To(Equal(1))

			jobConfig.SerialGroups = []string{
				"one",
			}
			Expect(jobConfig.MaxInFlight()).To(Equal(1))

			jobConfig.Serial = false
			Expect(jobConfig.MaxInFlight()).To(Equal(1))
		})

		It("returns 0 if MaxInFlight is not set, Serial is false, and SerialGroups is empty", func() {
			jobConfig := JobConfig{
				Serial:       false,
				SerialGroups: []string{},
			}

			Expect(jobConfig.MaxInFlight()).To(Equal(0))
		})
	})

	Describe("Inputs", func() {
		var (
			jobConfig JobConfig

			inputs []JobInputParams
		)

		BeforeEach(func() {
			jobConfig = JobConfig{}
		})

		JustBeforeEach(func() {
			inputs = jobConfig.Inputs()
		})

		Context("with a build plan", func() {
			Context("with an empty plan", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{}
				})

				It("returns an empty set of inputs", func() {
					Expect(inputs).To(BeEmpty())
				})
			})

			Context("with two serial gets", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &GetStep{
								Name:    "some-get-plan",
								Passed:  []string{"a", "b"},
								Trigger: true,
							},
						},
						{
							Config: &GetStep{
								Name: "some-other-get-plan",
							},
						},
					}
				})

				It("uses both for inputs", func() {
					Expect(inputs).To(Equal([]JobInputParams{
						{
							JobInput: JobInput{
								Name:     "some-get-plan",
								Resource: "some-get-plan",
								Passed:   []string{"a", "b"},
								Trigger:  true,
							},
						},
						{
							JobInput: JobInput{
								Name:     "some-other-get-plan",
								Resource: "some-other-get-plan",
								Trigger:  false,
							},
						},
					}))

				})
			})

			Context("when a plan has a version on a get", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &GetStep{
								Name: "a",
								Version: &VersionConfig{
									Every: true,
								},
							},
						},
					}
				})

				It("returns an input config with the version", func() {
					Expect(inputs).To(Equal(
						[]JobInputParams{
							{
								JobInput: JobInput{
									Name:     "a",
									Resource: "a",
									Version: &VersionConfig{
										Every: true,
									},
								},
							},
						},
					))
				})
			})

			Context("when a job has an ensure hook", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &GetStep{
								Name: "a",
							},
						},
					}

					jobConfig.Ensure = &Step{
						Config: &GetStep{
							Name: "b",
						},
					}
				})

				It("returns an input config for all get plans", func() {
					Expect(inputs).To(ConsistOf(
						JobInputParams{
							JobInput: JobInput{
								Name:     "a",
								Resource: "a",
							},
						},
						JobInputParams{
							JobInput: JobInput{
								Name:     "b",
								Resource: "b",
							},
						},
					))
				})
			})

			Context("when a job has a success hook", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &GetStep{
								Name: "a",
							},
						},
					}

					jobConfig.OnSuccess = &Step{
						Config: &GetStep{
							Name: "b",
						},
					}
				})

				It("returns an input config for all get plans", func() {
					Expect(inputs).To(ConsistOf(
						JobInputParams{
							JobInput: JobInput{
								Name:     "a",
								Resource: "a",
							},
						},
						JobInputParams{
							JobInput: JobInput{
								Name:     "b",
								Resource: "b",
							},
						},
					))

				})
			})

			Context("when a job has a failure hook", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &GetStep{
								Name: "a",
							},
						},
					}

					jobConfig.OnFailure = &Step{
						Config: &GetStep{
							Name: "b",
						},
					}
				})

				It("returns an input config for all get plans", func() {
					Expect(inputs).To(ConsistOf(
						JobInputParams{
							JobInput: JobInput{
								Name:     "a",
								Resource: "a",
							},
						},
						JobInputParams{
							JobInput: JobInput{
								Name:     "b",
								Resource: "b",
							},
						},
					))

				})
			})

			Context("when a job has an abort hook", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &GetStep{
								Name: "a",
							},
						},
					}

					jobConfig.OnAbort = &Step{
						Config: &GetStep{
							Name: "b",
						},
					}
				})

				It("returns an input config for all get plans", func() {
					Expect(inputs).To(ConsistOf(
						JobInputParams{
							JobInput: JobInput{
								Name:     "a",
								Resource: "a",
							},
						},
						JobInputParams{
							JobInput: JobInput{
								Name:     "b",
								Resource: "b",
							},
						},
					))

				})
			})

			Context("when a job has an error hook", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &GetStep{
								Name: "a",
							},
						},
					}

					jobConfig.OnError = &Step{
						Config: &GetStep{
							Name: "b",
						},
					}
				})

				It("returns an input config for all get plans", func() {
					Expect(inputs).To(ConsistOf(
						JobInputParams{
							JobInput: JobInput{
								Name:     "a",
								Resource: "a",
							},
						},
						JobInputParams{
							JobInput: JobInput{
								Name:     "b",
								Resource: "b",
							},
						},
					))

				})
			})

			Context("when a resource is specified", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &GetStep{
								Name:     "some-get-plan",
								Resource: "some-get-resource",
							},
						},
					}
				})

				It("uses it as resource in the input config", func() {
					Expect(inputs).To(Equal([]JobInputParams{
						{
							JobInput: JobInput{
								Name:     "some-get-plan",
								Resource: "some-get-resource",
								Trigger:  false,
							},
						},
					}))

				})
			})

			Context("when a simple aggregate plan is the first step", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &AggregateStep{
								Steps: []Step{
									{
										Config: &GetStep{
											Name: "a",
										},
									},
									{
										Config: &PutStep{
											Name: "y",
										},
									},
									{
										Config: &GetStep{
											Name:     "b",
											Resource: "some-resource", Passed: []string{"x"},
										},
									},
									{
										Config: &GetStep{
											Name: "c", Trigger: true,
										},
									},
								},
							},
						},
					}
				})

				It("returns an input config for all get plans", func() {
					Expect(inputs).To(Equal([]JobInputParams{
						{
							JobInput: JobInput{
								Name:     "a",
								Resource: "a",
								Trigger:  false,
							},
						},
						{
							JobInput: JobInput{
								Name:     "b",
								Resource: "some-resource",
								Passed:   []string{"x"},
								Trigger:  false,
							},
						},
						{
							JobInput: JobInput{
								Name:     "c",
								Resource: "c",
								Trigger:  true,
							},
						},
					}))

				})
			})

			Context("when there are no gets in the plan", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &PutStep{
								Name: "some-put-plan",
							},
						},
					}
				})

				It("returns an empty set of inputs", func() {
					Expect(inputs).To(BeEmpty())
				})
			})
		})
	})

	Describe("Outputs", func() {
		var (
			jobConfig JobConfig

			outputs []JobOutput
		)

		BeforeEach(func() {
			jobConfig = JobConfig{}
		})

		JustBeforeEach(func() {
			outputs = jobConfig.Outputs()
		})

		Context("with a build plan", func() {
			Context("with an empty plan", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{}
				})

				It("returns an empty set of outputs", func() {
					Expect(outputs).To(BeEmpty())
				})
			})

			Context("when a simple plan is configured", func() {
				BeforeEach(func() {
					jobConfig.PlanSequence = []Step{
						{
							Config: &PutStep{
								Name:     "some-name",
								Resource: "some-resource",
							},
						},
					}
				})

				It("returns an output for all of the put plans present", func() {
					Expect(outputs).To(Equal([]JobOutput{
						{
							Name:     "some-name",
							Resource: "some-resource",
						},
					}))

				})
			})
		})

		Context("when a job has an ensure hook", func() {
			BeforeEach(func() {
				jobConfig.PlanSequence = []Step{
					{
						Config: &PutStep{
							Name: "a",
						},
					},
				}

				jobConfig.Ensure = &Step{
					Config: &PutStep{
						Name: "b",
					},
				}
			})

			It("returns an input config for all get plans", func() {
				Expect(outputs).To(ConsistOf(
					JobOutput{
						Name:     "a",
						Resource: "a",
					},
					JobOutput{
						Name:     "b",
						Resource: "b",
					},
				))
			})
		})

		Context("when a job has a success hook", func() {
			BeforeEach(func() {
				jobConfig.PlanSequence = []Step{
					{
						Config: &PutStep{
							Name: "a",
						},
					},
				}

				jobConfig.OnSuccess = &Step{
					Config: &PutStep{
						Name: "b",
					},
				}
			})

			It("returns an input config for all get plans", func() {
				Expect(outputs).To(ConsistOf(
					JobOutput{
						Name:     "a",
						Resource: "a",
					},
					JobOutput{
						Name:     "b",
						Resource: "b",
					},
				))

			})
		})

		Context("when a job has a failure hook", func() {
			BeforeEach(func() {
				jobConfig.PlanSequence = []Step{
					{
						Config: &PutStep{
							Name: "a",
						},
					},
				}

				jobConfig.OnFailure = &Step{
					Config: &PutStep{
						Name: "b",
					},
				}
			})

			It("returns an input config for all get plans", func() {
				Expect(outputs).To(ConsistOf(
					JobOutput{
						Name:     "a",
						Resource: "a",
					},
					JobOutput{
						Name:     "b",
						Resource: "b",
					},
				))

			})
		})

		Context("when a job has an abort hook", func() {
			BeforeEach(func() {
				jobConfig.PlanSequence = []Step{
					{
						Config: &PutStep{
							Name: "a",
						},
					},
				}

				jobConfig.OnAbort = &Step{
					Config: &PutStep{
						Name: "b",
					},
				}
			})

			It("returns an input config for all get plans", func() {
				Expect(outputs).To(ConsistOf(
					JobOutput{
						Name:     "a",
						Resource: "a",
					},
					JobOutput{
						Name:     "b",
						Resource: "b",
					},
				))

			})
		})

		Context("when a job has an error hook", func() {
			BeforeEach(func() {
				jobConfig.PlanSequence = []Step{
					{
						Config: &PutStep{
							Name: "a",
						},
					},
				}

				jobConfig.OnError = &Step{
					Config: &PutStep{
						Name: "b",
					},
				}
			})

			It("returns an input config for all get plans", func() {
				Expect(outputs).To(ConsistOf(
					JobOutput{
						Name:     "a",
						Resource: "a",
					},
					JobOutput{
						Name:     "b",
						Resource: "b",
					},
				))

			})
		})

		Context("when the plan contains no puts steps", func() {
			BeforeEach(func() {
				jobConfig.PlanSequence = []Step{
					{
						Config: &GetStep{
							Name: "some-put-plan",
						},
					},
				}
			})

			It("returns an empty set of outputs", func() {
				Expect(outputs).To(BeEmpty())
			})
		})
	})
})
