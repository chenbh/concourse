package integration_test

import (
	"github.com/concourse/concourse/atc/types"
	"net/http"

	"github.com/concourse/concourse/atc/wrappa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Concurrent request limits", func() {
	BeforeEach(func() {
		cmd.ConcurrentRequestLimits = map[wrappa.LimitedRoute]int{
			wrappa.LimitedRoute(types.ListAllJobs): 0,
		}
	})

	It("disables ListAllJobs requests", func() {
		client := login(atcURL, "test", "test")
		httpClient := client.HTTPClient()
		request, _ := http.NewRequest(
			"GET",
			client.URL()+"/api/v1/jobs",
			nil,
		)

		response, _ := httpClient.Do(request)

		Expect(response.StatusCode).To(Equal(http.StatusNotImplemented))
	})
})
