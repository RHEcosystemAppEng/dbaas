package reconcilers

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var repoOrg = "quay.io/ecosystem-appeng/"
var _ = Describe("FetchImageAndVersion", func() {
	Context("Missing env var", func() {
		It("invalid, should return empty value", func() {
			os.Unsetenv("FOO")
			Expect(fetchEnvValue("FOO")).To(BeEmpty())
		})
		It("valid, should return default values from embedded file - config/default/manager-env-images.yaml", func() {
			os.Unsetenv(dbaasDynamicPluginVersion)
			os.Unsetenv(dbaasDynamicPluginImg)
			Expect(fetchEnvValue(dbaasDynamicPluginImg)).To(Equal(repoOrg + fetchEnvValue(dbaasDynamicPluginVersion)))
		})
	})

	Context("Existing env var", func() {
		It("should return set value", func() {
			imageTest := "test-image@sha256:fds45ds21kl"
			os.Setenv(dbaasDynamicPluginImg, imageTest)
			Expect(fetchEnvValue(dbaasDynamicPluginImg)).To(Equal(imageTest))
		})
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FetchEnvValue Suite")
}
