package segments

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bookend{}", func() {
	Context("When Bookend{}", func() {
		Describe("Output()", func() {
			It("is expected to be: ├", func() {
				Expect(Bookend{}.Output()).To(Equal("├"))
			})
		})
	})

	Context("When Bookend{Left:true} is true", func() {
		Describe("Output()", func() {
			It("is expected to be: ┤", func() {
				Expect(Bookend{Left: true}.Output()).To(Equal("┤"))
			})
		})
	})

	Describe("Len()", func() {
		It("is expected to be 2", func() {
			Expect(Bookend{}.Len()).To(Equal(1))
		})
	})
})
