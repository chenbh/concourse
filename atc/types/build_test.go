package types_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

var _ = Describe("Build", func() {
	Describe("OneOff", func() {
		It("returns true if there is no JobName", func() {
			build := types.Build{
				JobName: "",
			}
			Expect(build.OneOff()).To(BeTrue())
		})

		It("returns false if there is a JobName", func() {
			build := types.Build{
				JobName: "something",
			}
			Expect(build.OneOff()).To(BeFalse())
		})
	})

	Describe("IsRunning", func() {
		It("returns true if the build is pending", func() {
			build := types.Build{
				Status: string(types.StatusPending),
			}
			Expect(build.Abortable()).To(BeTrue())
		})

		It("returns true if the build is started", func() {
			build := types.Build{
				Status: string(types.StatusStarted),
			}
			Expect(build.Abortable()).To(BeTrue())
		})

		It("returns false if in any other state", func() {
			states := []types.BuildStatus{
				types.StatusAborted,
				types.StatusErrored,
				types.StatusFailed,
				types.StatusSucceeded,
			}

			for _, state := range states {
				build := types.Build{
					Status: string(state),
				}
				Expect(build.Abortable()).To(BeFalse())
			}
		})
	})

	Describe("Abortable", func() {
		It("returns true if the build is pending", func() {
			build := types.Build{
				Status: string(types.StatusPending),
			}
			Expect(build.Abortable()).To(BeTrue())
		})

		It("returns true if the build is started", func() {
			build := types.Build{
				Status: string(types.StatusStarted),
			}
			Expect(build.Abortable()).To(BeTrue())
		})

		It("returns false if in any other state", func() {
			states := []db.BuildStatus{
				db.BuildStatusAborted,
				db.BuildStatusErrored,
				db.BuildStatusFailed,
				db.BuildStatusSucceeded,
			}

			for _, state := range states {
				build := types.Build{
					Status: string(state),
				}
				Expect(build.Abortable()).To(BeFalse())
			}
		})
	})
})
