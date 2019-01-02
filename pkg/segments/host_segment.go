package segments

import (
	"fmt"
	"os"
)

type Host struct{}

func (s Host) output() string {
	_, sshExists := os.LookupEnv("SSH_CLIENT")
	hostnameValue, _ := os.Hostname()

	if sshExists {
		return fmt.Sprintf("@%v", hostnameValue)
	}

	return ""
}

func (s Host) len() int {
	return len(Host.output(s))
}
