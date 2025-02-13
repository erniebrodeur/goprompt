package layout_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Layout Placeholders", func() {
	Context("Invalid placeholders", func() {
		It("either leaves them literal or handles them as errors", func() {
			// Implementation detail for how we treat unknown placeholders
		})
	})

	Context("Multiple placeholders with partial error", func() {
		It("replaces only the failing segment with [ERR]", func() {
			// Possibly tested by a mock aggregator if needed
		})
	})
})