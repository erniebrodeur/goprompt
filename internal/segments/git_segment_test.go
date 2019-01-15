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
		git := Git{parsed: true, gitString: "## master...origin/master [ahead 1]\n				M internal/segments/git_segment.go\n				M internal/segments/git_segment_test.go"}

		JustBeforeEach(func() {
			git.parseGit()
		})

		PDescribe(".Output()", func() {
			git := Git{}

			It("is expected to include ...", func() {
				Expect(git.Output()).NotTo(BeEmpty())
			})

			Context("and the local repo is ahead of the remote repo", func() {
				git := Git{}

				It("is expected to include (push)", func() {
					Expect(git.direction).To(ContainSubstring("(push)"))
				})
			})

			Context("and the local repo is behind of the remote repo", func() {
				git := Git{}

				It("is expected to include (pull)", func() {
					Expect(git.direction).To(ContainSubstring("(pull)"))
				})
			})
		})

		Describe(".parseGit()", func() {
			Context("when git.gitString is not set", func() {
				PIt("is expected to call .getGitString()", func() {})
			})

			It("is expected to set git.branch", func() {
				Expect(git.branch).To(Equal("master"))
			})
			It("is expected to set git.remoteBranch", func() {
				Expect(git.remoteBranch).To(Equal("origin/master"))
			})
			It("is expected to set git.direction", func() {
				Expect(git.direction).To(ContainSubstring("ahead"))
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

	Context("When the current directory is not inside of a git tree", func() {
		Describe(".Output()", func() {
			git := Git{gitString: "", parsed: true}

			It("is expected to be empty", func() {
				Expect(git.Output()).To(BeEmpty())
			})
		})
	})
})
