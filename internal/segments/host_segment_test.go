package segments

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Host{}", func() {
	host := Host{}

	Describe("Output()", func() {
		It("is expected to be blank", func() {
			Expect(host.Output()).To(BeEmpty())
		})
	})

	Describe("Len()", func() {
		It("is expected to be 0", func() {
			Expect(host.Len()).To(Equal(0))
		})
	})

	Context("When the environmental variable SSH_CLIENT is set", func() {
		originalClient := os.Getenv("SSH_CLIENT")
		hostname, _ := os.Hostname()
		wantedHostname := fmt.Sprintf("@%s", hostname)

		Describe("Output()", func() {
			It(fmt.Sprintf("is expected to be: @%s", hostname), func() {
				os.Setenv("SSH_CLIENT", "192.168.1.1")
				Expect(host.Output()).To(Equal(wantedHostname))
			})
		})

		Describe("Len()", func() {
			It(fmt.Sprintf("is expected to be: %v", len(wantedHostname)), func() {
				Expect(host.Len()).To(Equal(len(wantedHostname)))
			})
		})

		if originalClient != "" {
			os.Setenv("SSH_CLIENT", originalClient)
		} else {
			os.Unsetenv("SSH_CLIENT")
		}
	})
})
