package segments

import "os"

type Pwd struct{}

// pwd git login host left right left right
func (s Pwd) maxLen(terminalWidth int) int {
	return 2000
	// return terminalWidth -
	// 	Pwd{}.len() -
	// 	gitSegment{}.len() -
	// 	loginSegment{}.len() -
	// 	hostSegment{}.len() -
	// 	(leftSegment{}.len() * 2) -
	// 	(rightSegment{}.len() * 2)
}
func (s Pwd) len() int {
	return len(Pwd.output(s))
}

func (s Pwd) output() string {
	out, _ := os.Getwd()

	return string(out)
}
