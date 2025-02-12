package goprompt_test

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/erniebrodeur/goprompt/cmd/goprompt"
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
	rootCmd := goprompt.GetRootCmd()
	output, err := executeCommand(rootCmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(output, "[rootCmd]") {
		t.Errorf("expected root cmd placeholder mention, got %q", output)
	}
}

