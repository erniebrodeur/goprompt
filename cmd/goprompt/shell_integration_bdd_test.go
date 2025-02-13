package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shell Integration", func() {
	Context("`goprompt shell` subcommand", func() {
		It("prints lines for .bashrc/.zshrc to automatically run goprompt", func() {
			// We'll eventually check if the CLI outputs instructions
			// like: 'eval "$(goprompt shell)"'
		})
	})
})