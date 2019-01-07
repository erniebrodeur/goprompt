package segments

import (
	"fmt"
	"os"
)

type Host struct{}

func (h Host) ColoredOutput() string {
	return h.Output()
}

// Len return length of string without invisible characters counted
func (h Host) Len() int {
	return len(h.Output())
}

func (h Host) Output() string {
	_, sshExists := os.LookupEnv("SSH_CLIENT")
	hostnameValue, _ := os.Hostname()

	if sshExists {
		return fmt.Sprintf("@%v", hostnameValue)
	}

	return ""
}
