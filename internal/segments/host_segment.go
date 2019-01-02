package segments

import (
	"fmt"
	"os"
)

type Host struct{}

func (s Host) Output() string {
	_, sshExists := os.LookupEnv("SSH_CLIENT")
	hostnameValue, _ := os.Hostname()

	if sshExists {
		return fmt.Sprintf("@%v", hostnameValue)
	}

	return ""
}

func (s Host) Len() int {
	return len(Host.Output(s))
}
