package wrappa_test

import (
	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/atc/wrappa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Concurrent Request Policy", func() {
	Describe("LimitedRoute#UnmarshalFlag", func() {
		It("unmarshals ListAllJobs", func() {
			var flagValue wrappa.LimitedRoute
			flagValue.UnmarshalFlag(types.ListAllJobs)
			expected := wrappa.LimitedRoute(types.ListAllJobs)
			Expect(flagValue).To(Equal(expected))
		})

		It("raises an error when the action is not supported", func() {
			var flagValue wrappa.LimitedRoute
			err := flagValue.UnmarshalFlag(types.CreateJobBuild)

			expected := "action 'CreateJobBuild' is not supported"
			Expect(err.Error()).To(ContainSubstring(expected))
		})

		It("error message describes supported actions", func() {
			var flagValue wrappa.LimitedRoute
			err := flagValue.UnmarshalFlag(types.CreateJobBuild)

			expected := "Supported actions are: "
			Expect(err.Error()).To(ContainSubstring(expected))
		})
	})

	Describe("ConcurrentRequestPolicy#HandlerPool", func() {
		It("returns true when an action is limited", func() {
			policy := wrappa.NewConcurrentRequestPolicy(
				map[wrappa.LimitedRoute]int{
					wrappa.LimitedRoute(types.CreateJobBuild): 0,
				},
			)

			_, found := policy.HandlerPool(types.CreateJobBuild)
			Expect(found).To(BeTrue())
		})

		It("returns false when an action is not limited", func() {
			policy := wrappa.NewConcurrentRequestPolicy(
				map[wrappa.LimitedRoute]int{
					wrappa.LimitedRoute(types.CreateJobBuild): 0,
				},
			)

			_, found := policy.HandlerPool(types.ListAllPipelines)
			Expect(found).To(BeFalse())
		})

		It("holds a reference to its pool", func() {
			policy := wrappa.NewConcurrentRequestPolicy(
				map[wrappa.LimitedRoute]int{
					wrappa.LimitedRoute(types.CreateJobBuild): 1,
				},
			)
			pool1, _ := policy.HandlerPool(types.CreateJobBuild)
			pool1.TryAcquire()
			pool2, _ := policy.HandlerPool(types.CreateJobBuild)
			Expect(pool2.TryAcquire()).To(BeFalse())
		})
	})
})
