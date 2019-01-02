package main

import (
	"os"
)

type loginSegment struct{}

func (s loginSegment) output() string {
	val, _ := os.LookupEnv("USER")

	return val
}

func (s loginSegment) len() int {
	return len(loginSegment.output(s))
}
