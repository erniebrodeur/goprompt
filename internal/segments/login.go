package segments

import (
	"os"

	"github.com/mgutz/ansi"
)

type Login struct{}

func (l Login) ColoredOutput() string {
	phosphorize := ansi.ColorFunc("green+h:black")

	return phosphorize(l.Output())
}

// Len return length of string without invisible characters counted
func (l Login) Len() int {
	return len(l.Output())
}

func (l Login) Output() string {
	return os.Getenv("USER")
}
