package theme_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Theming", func() {
	Context("Hex color conversion", func() {
		It("handles valid #RRGGBB codes", func() {
			// e.g. #F92672 => [249,38,114] => \033[38;2;249;38;114m
		})
	})

	Context("Missing theme keys", func() {
		It("returns default or no color if a segment's state key doesn't exist", func() {
			// e.g. 'git.dirty' missing => fallback
		})
	})

	Context("Segment-based states", func() {
		It("applies segment-specific coloring, e.g. git.clean, git.dirty, user.root, etc.", func() {
			// We'll see how each segment picks its subkey
		})
	})
})

