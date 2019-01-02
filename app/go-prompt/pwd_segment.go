package main

import "os"

type pwdSegment struct{}

// pwd git login host left right left right
func (s pwdSegment) maxLen(terminalWidth int) int {
	return terminalWidth -
		pwdSegment{}.len() -
		gitSegment{}.len() -
		loginSegment{}.len() -
		hostSegment{}.len() -
		(leftSegment{}.len() * 2) -
		(rightSegment{}.len() * 2)
}
func (s pwdSegment) len() int {
	return len(pwdSegment.output(s))
}

func (s pwdSegment) output() string {
	out, _ := os.Getwd()

	return string(out)
}
