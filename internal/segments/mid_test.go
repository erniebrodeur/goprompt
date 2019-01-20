package segments

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mid{}", func() {
	mid := Mid{Count: 1}

	Context("When Mid{Count: 1}", func() {
		Describe("Output()", func() {
			It("is expected to be: ─", func() {
				Expect(mid.Output()).To(Equal("─"))
			})
		})
	})

	Describe("Len()", func() {
		It("is expected to be m.Count", func() {
			Expect(mid.Len()).To(Equal(mid.Count))
		})
	})
})
