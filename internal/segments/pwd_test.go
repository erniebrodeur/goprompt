package segments

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pwd{}", func() {
	path := "/path"
	longPath := "/a/really/long/path/somewhere"
	var pwd = &Pwd{}

	JustBeforeEach(func() {
		pwd = NewPwd()
		pwd.pwdBuilder = func() string { return path }
	})

	Describe("Output()", func() {
		It("is expected to be the current working directory", func() {
			Expect(pwd.Output()).To(Equal(path))
		})

		Context("When the length of PWD is greater than 1/4th of terminalWidth", func() {
			It("is expected to replace segments on the left with ...", func() {
				pwd.pwdBuilder = func() string { return longPath }
				Expect(strings.HasPrefix(pwd.Output(), "...")).To(Equal(true))
			})
		})
	})

	Describe("Len()", func() {
		It("is expected to be the length of the current working directory.", func() {
			Expect(pwd.Len()).To(Equal(len(path)))
		})
	})
})
