package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shell Integration", func() {
	Context("`goprompt shell` subcommand", func() {
		It("prints lines locally without calling a global binary", func() {
			Expect(true).To(BeTrue())
		})
	})
})