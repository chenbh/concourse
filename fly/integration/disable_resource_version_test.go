package integration_test

import (
	"fmt"
	"github.com/concourse/concourse/atc/types"
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/ghttp"
	"github.com/tedsuo/rata"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Fly CLI", func() {
	Describe("disable-resource-version", func() {
		var (
			expectedGetStatus    int
			expectedPutStatus    int
			disablePath, getPath string
			err                  error
			teamName             = "main"
			pipelineName         = "pipeline"
			resourceName         = "resource"
			resourceVersionID    = "42"
			disableVersion       = "some:value"
			pipelineResource     = fmt.Sprintf("%s/%s", pipelineName, resourceName)
			expectedVersion      = types.ResourceVersion{
				ID:      42,
				Version: types.Version{"some": "value"},
				Enabled: true,
			}
		)

		Context("make sure the command exists", func() {
			It("calls the disable-resource-version command", func() {
				flyCmd := exec.Command(flyPath, "disable-resource-version")
				sess, err := gexec.Start(flyCmd, GinkgoWriter, GinkgoWriter)

				Expect(err).ToNot(HaveOccurred())
				Consistently(sess.Err).ShouldNot(gbytes.Say("error: Unknown command"))

				<-sess.Exited
			})
		})

		Context("when the resource is specified", func() {
			Context("when the resource version json string is specified", func() {
				BeforeEach(func() {
					getPath, err = types.Routes.CreatePathForRoute(types.ListResourceVersions, rata.Params{
						"pipeline_name": pipelineName,
						"team_name":     teamName,
						"resource_name": resourceName,
					})
					Expect(err).NotTo(HaveOccurred())

					disablePath, err = types.Routes.CreatePathForRoute(types.DisableResourceVersion, rata.Params{
						"pipeline_name":              pipelineName,
						"team_name":                  teamName,
						"resource_name":              resourceName,
						"resource_config_version_id": resourceVersionID,
					})
					Expect(err).NotTo(HaveOccurred())

				})

				JustBeforeEach(func() {
					atcServer.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("GET", getPath, "filter=some:value"),
							ghttp.RespondWithJSONEncoded(expectedGetStatus, []types.ResourceVersion{expectedVersion}),
						),
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("PUT", disablePath),
							ghttp.RespondWith(expectedPutStatus, nil),
						),
					)
				})

				Context("when the resource and version exists", func() {
					BeforeEach(func() {
						expectedGetStatus = http.StatusOK
						expectedPutStatus = http.StatusOK

						expectedVersion.Enabled = true
					})

					It("disables the resource version", func() {
						Expect(func() {
							flyCmd := exec.Command(flyPath, "-t", targetName, "disable-resource-version", "-r", pipelineResource, "-v", disableVersion)

							sess, err := gexec.Start(flyCmd, GinkgoWriter, GinkgoWriter)
							Expect(err).NotTo(HaveOccurred())

							Eventually(sess.Out).Should(gbytes.Say(fmt.Sprintf("disabled '%s' with version {\"some\":\"value\"}\n", pipelineResource)))

							<-sess.Exited
							Expect(sess.ExitCode()).To(Equal(0))
						}).To(Change(func() int {
							return len(atcServer.ReceivedRequests())
						}).By(3))
					})
				})

				Context("when the resource does not exist", func() {
					BeforeEach(func() {
						expectedGetStatus = http.StatusNotFound
					})

					It("errors", func() {
						Expect(func() {
							flyCmd := exec.Command(flyPath, "-t", targetName, "disable-resource-version", "-r", pipelineResource, "-v", disableVersion)

							sess, err := gexec.Start(flyCmd, GinkgoWriter, GinkgoWriter)
							Expect(err).NotTo(HaveOccurred())

							Eventually(sess.Err).Should(gbytes.Say(fmt.Sprintf("could not find version matching {\"some\":\"value\"}\n")))

							<-sess.Exited
							Expect(sess.ExitCode()).To(Equal(1))
						}).To(Change(func() int {
							return len(atcServer.ReceivedRequests())
						}).By(2))
					})
				})

				Context("when the resource version does not exist", func() {
					BeforeEach(func() {
						expectedPutStatus = http.StatusNotFound
						expectedGetStatus = http.StatusOK

						expectedVersion.Enabled = true
					})

					It("fails to disable", func() {
						Expect(func() {
							flyCmd := exec.Command(flyPath, "-t", targetName, "disable-resource-version", "-r", pipelineResource, "-v", disableVersion)

							sess, err := gexec.Start(flyCmd, GinkgoWriter, GinkgoWriter)
							Expect(err).NotTo(HaveOccurred())

							Eventually(sess.Err).Should(gbytes.Say(fmt.Sprintf("could not disable '%s', make sure the resource version exists", pipelineResource)))

							<-sess.Exited
							Expect(sess.ExitCode()).To(Equal(1))
						}).To(Change(func() int {
							return len(atcServer.ReceivedRequests())
						}).By(3))
					})
				})
				Context("when the resource version is already disabled", func() {
					BeforeEach(func() {
						expectedGetStatus = http.StatusOK
						expectedVersion.Enabled = false
					})

					It("returns successfully without calling api", func() {
						Expect(func() {
							flyCmd := exec.Command(flyPath, "-t", targetName, "disable-resource-version", "-r", pipelineResource, "-v", disableVersion)

							sess, err := gexec.Start(flyCmd, GinkgoWriter, GinkgoWriter)
							Expect(err).NotTo(HaveOccurred())

							Eventually(sess.Out).Should(gbytes.Say(fmt.Sprintf("disabled '%s' with version {\"some\":\"value\"}\n", pipelineResource)))

							<-sess.Exited
							Expect(sess.ExitCode()).To(Equal(0))

							for _, request := range atcServer.ReceivedRequests() {
								Expect(request.RequestURI).NotTo(Equal(disablePath))
							}
						}).To(Change(func() int {
							return len(atcServer.ReceivedRequests())
						}).By(2))
					})
				})
			})
		})
	})
})
