// UTF-8

package segments

import "strings"

var TerminalWidth = 200

type Stretchy struct {
	LengthLeft, LengthRight int
}

func (s Stretchy) Output() string {
	// the meta chars I'm using have a len of 3.  We have five chars, so we do 5*3-3 (since they already have len 1)
	return strings.Repeat("─", TerminalWidth-s.LengthLeft-s.LengthRight)
}

func (s Stretchy) Len() int {
	return 1
}
