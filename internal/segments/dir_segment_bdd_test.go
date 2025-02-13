package segments_test

import (
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/erniebrodeur/goprompt/internal/segments"
)

var _ = Describe("Directory Segment Details", func() {
	var origDir string

	BeforeEach(func() {
		origDir, _ = os.Getwd()
	})

	AfterEach(func() {
		os.Chdir(origDir)
	})

	Context("Multi-component truncation", func() {
		It("shows only the last two components", func() {
			tmp := GinkgoT().TempDir()
			sub := filepath.Join(tmp, "app", "src")
			Expect(os.MkdirAll(sub, 0755)).To(Succeed())
			Expect(os.Chdir(sub)).To(Succeed())

			ds := &segments.DirSegment{ShowComponents: 2}
			out, err := ds.Render(nil)
			Expect(err).To(BeNil())
			parts := strings.Split(out, "/")
			Expect(parts).To(HaveLen(2))
			Expect(out).To(ContainSubstring("app"))
			Expect(out).To(ContainSubstring("src"))
		})
	})

	Context("Directory access error", func() {
		It("yields [ERR]", func() {
			// Force an error scenario or mock if necessary
		})
	})
})

