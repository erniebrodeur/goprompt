package env_test

import (
	"os"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Environment Overrides", func() {
	Context("GOTHEME and GOLAYOUT", func() {
		BeforeEach(func() {
			os.Setenv("GOTHEME", "monokai_dark")
			os.Setenv("GOLAYOUT", "$dir(1) $git")
		})

		AfterEach(func() {
			os.Unsetenv("GOTHEME")
			os.Unsetenv("GOLAYOUT")
		})

		It("uses environment values if no flags override them (placeholder test)", func() {
			// Minimal usage of Gomega:
			Expect(true).To(BeTrue())
		})
	})

	Context("Command line flags override environment", func() {
		It("applies the flag-specified theme/layout for a single run (placeholder test)", func() {
			// Minimal usage of Gomega:
			Expect(true).To(BeTrue())
		})
	})
})

