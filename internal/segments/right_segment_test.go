package segments

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RightSegment", func() {
	Describe("Output()", func() {
		It("is expected to be: ├", func() {
			Expect(rightSegment{}.Output()).To(Equal(" ├"))
		})
	})

	Describe("Len()", func() {
		It("is expected to be 2", func() {
			Expect(rightSegment{}.Len()).To(Equal(2))
		})
	})
})
