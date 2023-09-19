package segments

import (
	"os"

	"github.com/mgutz/ansi"
)

// Login is for returning the current user logged in
type Login struct{}

// ColoredOutput returns a color wrapped copy of Output
func (l Login) ColoredOutput() string {
	green := ansi.ColorFunc("green+h")

	return green(l.Output())
}

// Len return length of string without invisible characters counted
func (l Login) Len() int {
	return len(l.Output())
}

// Output returns the currently signed in user
func (l Login) Output() string {
	return os.Getenv("USER")
}
