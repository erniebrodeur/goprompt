package main_test

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	goprompt "github.com/erniebrodeur/goprompt/cmd/goprompt"
	"github.com/stretchr/testify/require"
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
	rootCmd := goprompt.GetRootCmd() // We'll define a getter in render.go or similar
	out, err := executeCommand(rootCmd)
	require.NoError(t, err)
	require.Contains(t, out, "[rootCmd]", "Expected placeholder mention in root command output")
}

func TestRenderCmd(t *testing.T) {
	rootCmd := goprompt.GetRootCmd()
	out, err := executeCommand(rootCmd, "render")
	require.NoError(t, err)
	require.NotEmpty(t, out, "render should produce some aggregator output (dir or empty git).")
}

---
