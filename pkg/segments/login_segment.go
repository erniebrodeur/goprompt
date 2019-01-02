package segments

import (
	"os"
)

type Login struct{}

func (s Login) output() string {
	val, _ := os.LookupEnv("USER")

	return val
}

func (s Login) len() int {
	return len(Login.output(s))
}
