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

		It("uses environment values if no flags override them", func() {
			// We'll expect aggregator or main CLI to read these by default
		})
	})

	Context("Command line flags override environment", func() {
		It("applies the flag-specified theme/layout for a single run", func() {
			// e.g., `goprompt render --theme alt --layout "[$git]($dir)"`
			// Needs integration test in CLI
		})
	})
})

