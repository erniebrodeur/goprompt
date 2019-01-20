package segments

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ahead  = "## master...origin/master [ahead 1]\n				M internal/segments/git_segment.go\n				M internal/segments/git_segment_test.go"
	behind = "## master...origin/master [behind 1]\n				M internal/segments/git_segment.go\n				M internal/segments/git_segment_test.go"
)

var _ = Describe("Git{}", func() {
	var git = NewGit()

	JustBeforeEach(func() {
		git = NewGit()
	})

	Context("When the current directory is inside of a git tree", func() {
		Describe(".Output()", func() {
			Context("and the local repo is ahead of the remote repo", func() {
				It("is expected to include (push)", func() {
					git.gitBuilder = func() string { return ahead }
					git.parseGit()
					Expect(git.direction).To(ContainSubstring("(push)"))
				})
			})

			Context("and the local repo is behind the remote repo", func() {
				It("is expected to include (pull)", func() {
					git.gitBuilder = func() string { return behind }
					git.parseGit()
					Expect(git.direction).To(ContainSubstring("(pull)"))
				})
			})
		})

		Describe(".parseGit()", func() {
			JustBeforeEach(func() {
				git.gitBuilder = func() string { return ahead }
				git.parseGit()
			})

			It("is expected to set git.branch", func() {
				Expect(git.branch).To(Equal("master"))
			})
			It("is expected to set git.remoteBranch", func() {
				Expect(git.remoteBranch).To(Equal("origin/master"))
			})
			It("is expected to set git.direction", func() {
				Expect(git.direction).To(ContainSubstring("(push)"))
			})
			It("is expected to set git.dirty", func() {
				Expect(git.dirty).To(Equal("*"))
			})
		})
	})

	Context("When the current directory is not inside of a git tree", func() {
		Describe(".Output()", func() {
			git := NewGit()
			git.gitBuilder = func() string { return "" }

			It("is expected to be empty", func() {
				Expect(git.Output()).To(BeEmpty())
			})
		})
	})
})
