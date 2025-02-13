package segments_test

import (
. "github.com/onsi/ginkgo/v2"
. "github.com/onsi/gomega"

"github.com/erniebrodeur/goprompt/internal/segments"
"regexp"
)

var _ = Describe("Time Segment", func() {
Context("Basic time formatting", func() {
It("returns current time in default format if none specified (placeholder)", func() {
ts := &segments.TimeSegment{}
out, err := ts.Render(nil)
Expect(err).To(BeNil())
Expect(out).To(ContainSubstring("TIME("))
})
})

Context("Custom format string", func() {
It("uses '%H:%M' or similar (placeholder)", func() {
ts := &segments.TimeSegment{Format: "%H:%M"}
out, err := ts.Render(nil)
Expect(err).To(BeNil())

// We'll just test we got something like TIME(%H:%M)
Expect(regexp.MustCompile(`TIME$begin:math:text$%H:%M$end:math:text$`)).To(MatchString(out))
})
})
})

