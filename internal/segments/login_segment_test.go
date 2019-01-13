package segments

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Login{}", func() {
	login := Login{}
	user := os.Getenv("USER")

	Describe("Output()", func() {
		It("is expected to be the logged in user", func() {
			Expect(login.Output()).To(Equal(user))
		})
	})

	Describe("Len()", func() {
		It("is expected to be the length of user", func() {
			Expect(login.Len()).To(Equal(len(user)))
		})
	})
})
