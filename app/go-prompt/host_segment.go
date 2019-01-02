package main

import (
	"fmt"
	"os"
)

type hostSegment struct{}

func (s hostSegment) output() string {
	_, sshExists := os.LookupEnv("SSH_CLIENT")
	hostnameValue, _ := os.Hostname()

	if sshExists {
		return fmt.Sprintf("@%v", hostnameValue)
	}

	return ""
}

func (s hostSegment) len() int {
	return len(hostSegment.output(s))
}
