package segments

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bookend", func() {
	Describe("Output() is expected to be: ", func() {
		It("┤", func() {
			Expect(Bookend{Left: true}.Output()).To(Equal("┤"))
		})
	})

	Describe("Len()", func() {
		It("is expected to be 2", func() {
			Expect(Bookend{}.Len()).To(Equal(1))
		})
	})

	Context("When {Left} is true", func() {
		Describe("Output()", func() {
			It("is expected to be: ┤", func() {
				Expect(Bookend{Left: true}.Output()).To(Equal("┤"))
			})
		})
	})
})
