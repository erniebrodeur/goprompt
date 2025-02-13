package segments_test

import (
	"os"
	"os/user"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/erniebrodeur/goprompt/internal/segments"
)

var _ = Describe("User Segment", func() {
	Context("Normal user", func() {
		It("shows the current username if it’s not root", func() {
			// We look up the actual user from OS
			// In real tests, we might mock or just verify we get a non-root name
			currUser, err := user.Current()
			Expect(err).NotTo(HaveOccurred())

			uid, _ := strconv.Atoi(currUser.Uid)
			if uid == 0 {
				Skip("Running as root, skipping normal user test")
			}

			us := &segments.UserSegment{}
			out, segErr := us.Render(nil)
			Expect(segErr).To(BeNil())
			Expect(out).To(Equal(currUser.Username))
		})
	})

	Context("Root user", func() {
		It("shows an indicator (like 'root') if UID=0", func() {
			// Hard to simulate unless we run tests as root or mock user.Current
			Skip("Would require root or mocking user info")
		})
	})

	Context("Custom theming for user role", func() {
		It("applies 'user.normal' color if not root, 'user.root' if root", func() {
			// We rely on a theme map with keys like user.normal, user.root
			// For now, it's unimplemented
		})
	})
})

