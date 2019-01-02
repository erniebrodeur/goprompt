// UTF-8

package main

import "strings"

type stretchySegment struct {
	lengthLeft, lengthRight int
}

func (s stretchySegment) output() string {
	// the meta chars I'm using have a len of 3.  We have five chars, so we do 5*3-3 (since they already have len 1)
	return strings.Repeat("─", TerminalWidth-s.lengthLeft-s.lengthRight)
}

func (s stretchySegment) len() int {
	return 1
}
