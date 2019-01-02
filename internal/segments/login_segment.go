package segments

import (
	"os"
)

type Login struct{}

func (s Login) Output() string {
	val, _ := os.LookupEnv("USER")

	return val
}

func (s Login) Len() int {
	return len(Login.Output(s))
}
