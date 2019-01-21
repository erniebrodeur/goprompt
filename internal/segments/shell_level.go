package segments

import (
	"os"
)

// ShellLevel returns % for user and # for root
type ShellLevel struct{}

// Output returns % for user and # for root
func (s ShellLevel) Output() string {
	if os.Getenv("USER") == "root" {
		return "#"
	}

	return "%%"
}

// Len returns 1
func (s ShellLevel) Len() int {
	return 1
}
