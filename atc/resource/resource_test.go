package resource_test

import (
	"github.com/concourse/concourse/atc/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/concourse/concourse/atc/resource"
)

var _ = Describe("Resource", func() {
	Describe("ResourcesDir", func() {
		It("returns a file path with a prefix", func() {
			Expect(ResourcesDir("some-prefix")).To(ContainSubstring("some-prefix"))
		})
	})

	Describe("Signature", func() {
		var (
			resource Resource
			source   types.Source
			params   types.Params
			version  types.Version
		)

		BeforeEach(func() {
			source = types.Source{"some-source-key": "some-source-value"}
			params = types.Params{"some-params-key": "some-params-value"}
			version = types.Version{"some-version-key": "some-version-value"}

			resource = NewResourceFactory().NewResource(source, params, version)
		})

		It("marshals the source, params and version", func() {
			actualSignature, err := resource.Signature()
			Expect(err).ToNot(HaveOccurred())
			Expect(actualSignature).To(MatchJSON(`{
			  "source": {
				"some-source-key": "some-source-value"
			  },
			  "params": {
				"some-params-key": "some-params-value"
			  },
			  "version": {
				"some-version-key": "some-version-value"
			  }
			}`))
		})
	})
})
