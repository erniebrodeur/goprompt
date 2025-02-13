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
		// Stubbed approach
		rootCmd = GetRootCmdStub()
	})

	Context("No subcommand", func() {
		It("runs default aggregator logic (placeholder)", func() {
			out, err := executeCLI(rootCmd)
			Expect(err).To(BeNil())
			Expect(out).To(ContainSubstring("[rootCmd] Default aggregator-based render (placeholder)"))
		})
	})

	Context("`render` subcommand", func() {
		It("accepts flags for a one-off prompt generation", func() {
			out, err := executeCLI(rootCmd, "render", "--theme", "monokai_dark", "--layout", "$dir(2) $git")
			Expect(err).To(BeNil())
			Expect(out).To(Or(
				ContainSubstring("DIR_NOT_IMPL"),
				ContainSubstring("[ERR]"),
			))
		})
	})

	Context("`shell` subcommand", func() {
		It("prints lines for shell integration (placeholder)", func() {
			out, err := executeCLI(rootCmd, "shell")
			Expect(err).To(BeNil())
			Expect(out).To(ContainSubstring("eval"))
		})
	})

	Context("`theme` subcommand", func() {
		It("lists or applies themes (placeholder)", func() {
			out, err := executeCLI(rootCmd, "theme", "list")
			Expect(err).To(BeNil())
			Expect(out).To(Or(
				ContainSubstring("monokai_dark"),
				ContainSubstring("[ERR]"),
			))
		})
	})

	Context("`layout` subcommand", func() {
		It("lists or applies known layouts (placeholder)", func() {
			out, err := executeCLI(rootCmd, "layout", "list")
			Expect(err).To(BeNil())
			Expect(out).To(Or(
				ContainSubstring("single_line"),
				ContainSubstring("[ERR]"),
			))
		})
	})

	Context("Invalid flags or usage", func() {
		It("returns an error if user passes unknown flags (placeholder)", func() {
			out, err := executeCLI(rootCmd, "--bogus-flag")
			Expect(err).ToNot(BeNil())
			Expect(out).To(ContainSubstring("Error: unknown flag"))
		})
	})
})

// Minimal Stub
func GetRootCmdStub() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goprompt",
		Short: "GoPrompt is a customizable parallel prompt generator",
		Long:  `GoPrompt builds a prompt by running multiple segments in parallel, each with a short timeout.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("[rootCmd] Default aggregator-based render (placeholder)")
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
		Use:   "shell",
		Short: "Print shell integration lines",
		Run: func(c *cobra.Command, args []string) {
			c.Println("eval \"$(goprompt shell)\" # placeholder integration")
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

