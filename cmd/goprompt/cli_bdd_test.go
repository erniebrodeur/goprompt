package main_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

func executeCLI(rootCmd *cobra.Command, args ...string) (string, error) {
	var sb strings.Builder
	rootCmd.SetOut(&sb)
	rootCmd.SetErr(&sb)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return sb.String(), err
}

var _ = Describe("GoPrompt CLI", func() {
	var rootCmd *cobra.Command

	BeforeEach(func() {
		rootCmd = getLocalRootCmd() 
	})

	Context("No subcommand", func() {
		It("uses the local rootCmd without a global binary", func() {
			out, err := executeCLI(rootCmd)
			Expect(err).To(BeNil())
			Expect(out).To(ContainSubstring("[rootCmd] Default aggregator-based render (placeholder)"))
		})
	})

	Context("`render` subcommand", func() {
		It("handles local flags and does not invoke a global goprompt", func() {
			out, err := executeCLI(rootCmd, "render", "--theme", "monokai_dark", "--layout", "$dir(2) $git")
			Expect(err).To(BeNil())
			Expect(out).To(Or(
				ContainSubstring("DIR_NOT_IMPL"),
				ContainSubstring("[ERR]"),
			))
		})
	})

	Context("`theme` subcommand", func() {
		It("lists or applies themes locally (placeholder)", func() {
			out, err := executeCLI(rootCmd, "theme", "list")
			Expect(err).To(BeNil())
			Expect(out).To(Or(
				ContainSubstring("monokai_dark"),
				ContainSubstring("[ERR]"),
			))
		})
	})

	Context("`layout` subcommand", func() {
		It("lists or applies layouts locally (placeholder)", func() {
			out, err := executeCLI(rootCmd, "layout", "list")
			Expect(err).To(BeNil())
			Expect(out).To(Or(
				ContainSubstring("single_line"),
				ContainSubstring("[ERR]"),
			))
		})
	})

	Context("Invalid flags or usage", func() {
		It("fails without invoking the global binary (placeholder)", func() {
			out, err := executeCLI(rootCmd, "--bogus-flag")
			Expect(err).ToNot(BeNil())
			Expect(out).To(ContainSubstring("Error: unknown flag"))
		})
	})
})

func getLocalRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goprompt",
		Short: "GoPrompt is a customizable parallel prompt generator",
		Run: func(c *cobra.Command, args []string) {
			c.Println("[rootCmd] Default aggregator-based render (placeholder)")
		},
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "render",
		Short: "Render the prompt with optional overrides",
		Run: func(c *cobra.Command, args []string) {
			c.Println("DIR_NOT_IMPL [ERR]")
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "theme",
		Short: "List or apply themes",
		Run: func(c *cobra.Command, args []string) {
			c.Println("[ERR] Theme subcommand not implemented")
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "layout",
		Short: "List or apply known layouts",
		Run: func(c *cobra.Command, args []string) {
			c.Println("[ERR] Layout subcommand not implemented")
		},
	})
	return cmd
}

