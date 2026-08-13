package segments

import (
	"fmt"
	"os"
)

// Host is for returning the host if SSH'ed in
type Host struct{}

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
