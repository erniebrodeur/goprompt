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
		PContext("and the local repo is ahead of the remote repo", func() {
			Describe("git{}.direction", func() {
				git := Git{}

				It("is expected to include (push)", func() {
					Expect(git.direction).To(ContainSubstring("(push)"))
				})
			})
		})

		PContext("and the local repo is behind of the remote repo", func() {
			Describe("git{}.direction", func() {
				git := Git{}

				It("is expected to include (pull)", func() {
					Expect(git.direction).To(ContainSubstring("(pull)"))
				})
			})
		})

		Describe("git.direction", func() {
			PIt("it is expected to be blank")
		})

		Describe("git.branch", func() {
			PIt("is expected to be the current branch", func() {})
		})

		Describe("git.remoteBranch", func() {
			PIt("is expected to be the current branch", func() {})
		})

		Describe("git.dirty", func() {
			PIt("is expected to be a *")
		})

		Describe(".Output()", func() {
			git := Git{}

			It("is expected to include ...", func() {
				Expect(git.Output()).To(BeEmpty())
			})
		})

		Describe(".parse()", func() {
			Context("when git.gitString is not set", func() {
				PIt("is expected to call .getGitString()", func() {})
			})
			var git = Git{}

			JustBeforeEach(func() {
				git = Git{gitString: "## master...origin/master\n				M internal/segments/git_segment.go\n				M internal/segments/git_segment_test.go"}
				git.parse()
			})

			It("is expected to set git.branch", func() {
				Expect(git.branch).To(Equal("master"))
			})
			It("is expected to set git.remoteBranch", func() {
				Expect(git.remoteBranch).To(Equal(""))
			})
			It("is expected to set git.direction", func() {
				Expect(git.direction).To(Equal("master"))
			})
			It("is expected to set git.dirty", func() {
				Expect(git.dirty).To(Equal("*"))
			})

		})
		PDescribe(".getGitString()", func() {
			PIt("is expected to set git.gitString", func() {})
			PIt("is expected to call exec.Command() ...", func() {})
		})
	})

	PContext("When the current directory is not inside of a git tree", func() {
		It("is expected to be blank", func() {})
	})
})
