package segments_test

import (
"os"
"path/filepath"
"time"

gogit "github.com/go-git/go-git/v5"
"github.com/go-git/go-git/v5/plumbing/object"
. "github.com/onsi/ginkgo/v2"
. "github.com/onsi/gomega"

"github.com/erniebrodeur/goprompt/internal/segments"
)

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
Author: &object.Signature{
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
It("shows the branch if HEAD is not detached and no changes exist", func() {
tmp := GinkgoT().TempDir()
Expect(os.Chdir(tmp)).To(Succeed())
_, err := gogit.PlainInit(tmp, false)
Expect(err).NotTo(HaveOccurred())
createCommitInRepo(tmp, "file1.txt")

gs := &segments.GitSegment{}
out, gErr := gs.Render(nil)
Expect(gErr).To(BeNil())
Expect(out).NotTo(Equal("[ERR]"))
})
})
})

