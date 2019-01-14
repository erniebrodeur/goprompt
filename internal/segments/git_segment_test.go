package segments

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// type Git struct {
// 	branch, remoteBranch, direction, dirty string
// }

var _ = Describe("Git{}", func() {
	Context("When the current directory is inside of a git tree", func() {
		Context("and the local repo is ahead of the remote repo", func() {
			Describe("git{}.direction", func() {
				git := Git{}

				PIt("is expected to include (push)", func() {
					Expect(git.direction).To(ContainSubstring("(push)"))
				})
			})
		})

		Context("and the local repo is behind of the remote repo", func() {
			Describe("git{}.direction", func() {
				git := Git{}

				PIt("is expected to include (pull)", func() {
					Expect(git.direction).To(ContainSubstring("(pull)"))
				})
			})
		})

		Describe("git{}.direction", func() {
			PIt("it is expected to be blank")
		})

		Describe("git{}.branch", func() {
			PIt("is expected to be the current branch", func() {})
		})

		Describe("git{}.remoteBranch", func() {
			PIt("is expected to be the current branch", func() {})
		})

		Describe("git{}.dirty", func() {
			PIt("is expected to be a *")
		})

		Describe(".Output()", func() {
			git := Git{}

			It("is expected to include ...", func() {
				Expect(git.Output()).To(BeEmpty())
			})
		})

		Describe(".parse()", func() {
			PIt("is expected to set git{}.branch", func() {})
			PIt("is expected to set git{}.remoteBranch", func() {})
			PIt("is expected to set git{}.direction", func() {})
			PIt("is expected to set git{}.dirty", func() {})
		})
	})
})
