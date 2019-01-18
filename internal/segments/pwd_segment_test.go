package segments

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pwd{}", func() {
	path := "/a/path/somewhere"
	long_path := "/a/really/long/path/somewhere"

	Describe("Output()", func() {
		It("is expected to be the current working directory", func() {
			Expect(Pwd{TerminalWidth: 80, Path: path}.Output()).To(Equal(path))
		})

		Context("When the length of PWD is greater than 1/4th of terminalWidth", func() {
			It("is expected to ...", func() {
				Expect(Pwd{TerminalWidth: 80, Path: long_path}.Output()).To(Equal(path))
			})
		})
	})

	Describe("Len()", func() {
		It("is expected to be the length of the current working directory.", func() {
			Expect(Pwd{TerminalWidth: 80, Path: path}.Len()).To(Equal(len(path)))
		})
	})
})
