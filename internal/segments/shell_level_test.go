package segments

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShellLevel{}", func() {
	Describe("Output()", func() {
		It("is expected to be %", func() {
			Expect(ShellLevel{}.Output()).To(Equal("%"))
		})
	})

	Describe("Len()", func() {
		It("is expected to be the length of 1", func() {
			Expect(ShellLevel{}.Len()).To(Equal(1))
		})
	})

	Context("When the user is root", func() {
		username := os.Getenv("USER")

		Describe("Output()", func() {
			It("is expected to be #", func() {
				os.Setenv("USER", "root")
				Expect(ShellLevel{}.Output()).To(Equal("#"))
				os.Setenv("USER", username)
			})
		})

	})
})
