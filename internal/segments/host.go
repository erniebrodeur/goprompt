package segments

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// Host is for returning the host if SSH'ed in
type Host struct{}

// ColoredOutput returns a color wrapped copy of Output
func (h Host) ColoredOutput() string {
	yellow := ansi.ColorFunc("yellow+h")
	return yellow(h.Output())
}

// Len return length of string without invisible characters counted
func (h Host) Len() int {
	return len(h.Output())
}

// Output the host if env SSH_CLIENT is set
func (h Host) Output() string {
	sshClient := os.Getenv("SSH_CLIENT")

	if sshClient != "" {
		hostnameValue, _ := os.Hostname()
		return fmt.Sprintf("@%v", hostnameValue)
	}

	return ""
}
