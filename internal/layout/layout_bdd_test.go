package layout_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Layout Placeholders", func() {
	Context("Parsing and replacing placeholders", func() {
		It("replaces $dir(2), $git, $user, $time, etc. with real segment outputs", func() {
			// We'll eventually have a layout parser that calls aggregator with a list of segments
			// This test will fail until we implement that logic
		})
	})

	Context("Invalid placeholders", func() {
		It("leaves them literal or yields an error, based on design choice", func() {
			// e.g., $whatever(??)
		})
	})

	Context("Partial errors", func() {
		It("only inserts [ERR] for the failing segment, leaves others intact", func() {
			// aggregator concurrency logic merges results back into placeholders
		})
	})
})

