package segments

import "os"

// Login is for returning the current user logged in
type Login struct{}

// Len return length of string without invisible characters counted
func (l Login) Len() int {
	return len(l.Output())
}

// Output returns the currently signed in user
func (l Login) Output() string {
	return os.Getenv("USER")
}
