package segments

import (
	"os"
)

type ShellLevel struct{}

func (s ShellLevel) Output() string {
	if os.Getenv("USER") == "root" {
		return "#"
	}

	return "%"
}

func (s ShellLevel) Len() int {
	return 1
}
