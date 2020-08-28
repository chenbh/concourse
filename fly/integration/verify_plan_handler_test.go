package integration_test

import (
	"encoding/json"
	"net/http"

	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/atc/testhelpers"
	"github.com/onsi/gomega"
)

func VerifyPlan(expectedPlan atc.Plan) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var plan atc.Plan
		err := json.NewDecoder(r.Body).Decode(&plan)
		gomega.Expect(err).ToNot(gomega.HaveOccurred())

		gomega.Expect(plan).To(testhelpers.MatchPlan(expectedPlan))
	}
}
