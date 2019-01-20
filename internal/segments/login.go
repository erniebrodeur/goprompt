package segments

import "os"

type Login struct{}

func (l Login) ColoredOutput() string {
	return l.Output()
}

// Len return length of string without invisible characters counted
func (l Login) Len() int {
	return len(l.Output())
}

func (l Login) Output() string {
	return os.Getenv("USER")
}
