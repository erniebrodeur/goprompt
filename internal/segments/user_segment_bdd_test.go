package segments_test

import (
. "github.com/onsi/ginkgo/v2"
. "github.com/onsi/gomega"

"github.com/erniebrodeur/goprompt/internal/segments"
)

var _ = Describe("User Segment", func() {
Context("Normal user", func() {
It("shows the current username (placeholder)", func() {
us := &segments.UserSegment{}
out, err := us.Render(nil)
Expect(err).To(BeNil())
Expect(out).ToNot(BeEmpty())
// Could be 'USER_NOT_IMPL' right now
})
})
})