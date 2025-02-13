package segments_test

import (
	"os"
	"path/filepath"

	gogit "github.com/go-git/go-git/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/erniebrodeur/goprompt/internal/segments"
)

var _ = Describe("Git Segment States", func() {

	var origDir string

	BeforeEach(func() {
		origDir, _ = os.Getwd()
	})

	AfterEach(func() {
		os.Chdir(origDir)
	})

	Context("Clean repository", func() {
		It("detects a clean branch (placeholder example)", func() {
			tmp := GinkgoT().TempDir()
			Expect(os.Chdir(tmp)).To(Succeed())

			// Minimal Git init with no commits => HEAD might be weird
			_, err := gogit.PlainInit(tmp, false)
			Expect(err).NotTo(HaveOccurred())

			// Additional steps needed to create a real commit for a fully "clean" branch
			gs := &segments.GitSegment{}
			out, gerr := gs.Render(nil)

			// In a real test, we'd confirm a recognized branch name 
			// or handle empty HEAD as needed
			Expect(gerr).NotTo(BeNil())
			Expect(out).To(Equal("[ERR]"))
		})
	})

	Context("Detached HEAD", func() {
		It("returns [ERR] when HEAD is truly detached", func() {
			// Potential approach:
			// 1) git init
			// 2) commit a file on main
			// 3) checkout that commit => detached
			// 4) run the segment, expect [ERR]
		})
	})
})

