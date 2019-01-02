package segments

import "os"

type Pwd struct{}

// pwd git login host left right left right
func (s Pwd) MaxLen(terminalWidth int) int {
	return 2000
	// return terminalWidth -
	// 	Pwd{}.Len() -
	// 	gitSegment{}.Len() -
	// 	loginSegment{}.Len() -
	// 	hostSegment{}.Len() -
	// 	(leftSegment{}.Len() * 2) -
	// 	(rightSegment{}.Len() * 2)
}
func (s Pwd) Len() int {
	return len(Pwd.Output(s))
}

func (s Pwd) Output() string {
	out, _ := os.Getwd()

	return string(out)
}
