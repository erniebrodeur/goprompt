package main_test

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	// Import the main package path but rename to avoid direct conflicts
	goprompt "github.com/erniebrodeur/goprompt/cmd/goprompt"
)

func executeCommand(cmd *cobra.Command, args ...string) (string, error) {
	buf := new(strings.Builder)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return buf.String(), err
}

func TestRootNoArgs(t *testing.T) {
	// We call a function from the 'main' package that returns the rootCmd
	rootCmd := goprompt.GetRootCmd()

	out, err := executeCommand(rootCmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "[rootCmd]") {
		t.Errorf("expected root cmd placeholder, got %q", out)
	}
}

// If you want to test 'render' subcommand:
func TestRenderCmd(t *testing.T) {
	rootCmd := goprompt.GetRootCmd()
	out, err := executeCommand(rootCmd, "render")
	if err != nil {
		t.Fatalf("render cmd error: %v", err)
	}
	// The aggregator returns something like a dir + possibly empty Git
	// We can do a minimal check:
	if out == "" {
		t.Errorf("render output empty, expected dir or partial prompt")
	}
}

