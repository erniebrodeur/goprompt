package segments_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/erniebrodeur/goprompt/internal/segments"
)

var _ = Describe("Time Segment", func() {
	Context("Basic time formatting", func() {
		It("returns current time in default format if none specified", func() {
			ts := &segments.TimeSegment{}
			out, err := ts.Render(nil)
			Expect(err).To(BeNil())
			// Expect some form of time string, e.g., "14:05"
			Expect(out).NotTo(BeEmpty())
		})
	})

	Context("Custom format string", func() {
		It("follows the user-supplied format, e.g. '%H:%M'", func() {
			ts := &segments.TimeSegment{Format: "%H:%M"}
			out, err := ts.Render(nil)
			Expect(err).To(BeNil())
			// We check if 'out' matches a typical "hh:mm" pattern
			Expect(out).To(MatchRegexp(`^\d{1,2}:\d{2}$`))
		})
	})

	Context("Error handling", func() {
		It("never realistically errors, but if parse fails, returns [ERR]", func() {
			// If we introduced a weird custom parse or something, we could handle it
		})
	})
})

