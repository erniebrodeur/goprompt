package segments_test

import (
	"os"
	"path/filepath"
	"time"

	gogit "github.com/go-git/go-git/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/erniebrodeur/goprompt/internal/segments"
)

// Common method to commit a file so we have a valid branch.
func createCommitInRepo(repoPath string, fileName string) {
	filePath := filepath.Join(repoPath, fileName)
	Expect(os.WriteFile(filePath, []byte("test content"), 0644)).To(Succeed())

	repo, err := gogit.PlainOpen(repoPath)
	Expect(err).NotTo(HaveOccurred())

	w, err := repo.Worktree()
	Expect(err).NotTo(HaveOccurred())

	_, err = w.Add(fileName)
	Expect(err).NotTo(HaveOccurred())

	_, err = w.Commit("Initial commit", &gogit.CommitOptions{
		Author: &gogit.Signature{
			Name:  "Test Author",
			Email: "test@example.com",
			When:  time.Now(),
		},
	})
	Expect(err).NotTo(HaveOccurred())
}

var _ = Describe("Git Segment States", func() {
	var origDir string

	BeforeEach(func() {
		origDir, _ = os.Getwd()
	})

	AfterEach(func() {
		os.Chdir(origDir)
	})

	Context("Clean repository", func() {
		It("shows the branch name if HEAD is not detached and no changes exist", func() {
			tmp := GinkgoT().TempDir()
			Expect(os.Chdir(tmp)).To(Succeed())

			_, err := gogit.PlainInit(tmp, false)
			Expect(err).NotTo(HaveOccurred())

			// Make an initial commit so HEAD is on 'master' or 'main'
			createCommitInRepo(tmp, "file1.txt")

			gs := &segments.GitSegment{}
			output, gErr := gs.Render(nil)

			// We'll expect something like "main" or "master"
			Expect(gErr).To(BeNil())
			Expect(output).NotTo(BeEmpty())
			Expect(output).NotTo(Equal("[ERR]"))
		})
	})

	Context("Dirty repository", func() {
		It("indicates a dirty branch (e.g., 'feature*') if uncommitted changes exist", func() {
			tmp := GinkgoT().TempDir()
			Expect(os.Chdir(tmp)).To(Succeed())

			_, err := gogit.PlainInit(tmp, false)
			Expect(err).NotTo(HaveOccurred())

			// Switch to a 'feature' branch for clarity
			repo, _ := gogit.PlainOpen(tmp)
			w, _ := repo.Worktree()
			Expect(w.Checkout(&gogit.CheckoutOptions{
				Branch: "refs/heads/feature",
				Create: true,
			})).To(Succeed())

			// Commit once so we have a baseline
			createCommitInRepo(tmp, "file1.txt")

			// Modify file but don't commit => dirty state
			Expect(os.WriteFile("file1.txt", []byte("more content"), 0644)).To(Succeed())

			gs := &segments.GitSegment{}
			output, gErr := gs.Render(nil)

			// We expect some indication of dirty. e.g., "feature*"
			Expect(gErr).To(BeNil())
			Expect(output).To(ContainSubstring("feature"))
			// We'll see how we represent dirty => star or something else
		})
	})

	Context("Detached HEAD", func() {
		It("returns [ERR] or 'DETACHED' if HEAD is truly detached", func() {
			tmp := GinkgoT().TempDir()
			Expect(os.Chdir(tmp)).To(Succeed())

			_, err := gogit.PlainInit(tmp, false)
			Expect(err).NotTo(HaveOccurred())

			// Commit on 'main'
			createCommitInRepo(tmp, "file1.txt")

			// Detach HEAD by checking out the commit
			repo, err := gogit.PlainOpen(tmp)
			Expect(err).NotTo(HaveOccurred())
			refs, err := repo.Head()
			Expect(err).NotTo(HaveOccurred())

			w, err := repo.Worktree()
			Expect(err).NotTo(HaveOccurred())
			Expect(w.Checkout(&gogit.CheckoutOptions{
				Hash:  refs.Hash(),
				Force: true,
			})).To(Succeed())

			// Now HEAD is detached
			gs := &segments.GitSegment{}
			output, gErr := gs.Render(nil)
			Expect(gErr).To(BeNil())
			Expect(output).To(SatisfyAny(
				Equal("[ERR]"),
				ContainSubstring("DETACHED"),
			))
		})
	})

	Context("Rebase in progress (not fully tested)", func() {
		It("detects rebase state and indicates it is dirty or rebasing", func() {
			// Implementation detail. We can simulate rebase by 
			// creating .git/REBASE_HEAD or .git/rebase-apply. 
			// Possibly advanced, left as a placeholder.
		})
	})
})