package theme_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/erniebrodeur/goprompt/internal/theme"
)

var _ = Describe("Theming", func() {
	Context("Hex color conversion", func() {
		It("parses #F92672 correctly", func() {
			got := theme.HexToANSI("#F92672")
			Expect(got).To(Equal("\033[38;2;249;38;114m"))
		})
	})

	Context("Missing theme keys", func() {
		It("falls back to default or prints no color", func() {
			// Implementation detail depends on how we handle missing keys
		})
	})
})

