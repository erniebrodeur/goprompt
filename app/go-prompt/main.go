package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// Goprompt comment
type Goprompt interface {
	TerminalWidth() int
}

type goPrompt struct{}

func main() {
	login := loginSegment{}
	host := hostSegment{}
	stretchy := stretchySegment{}

	leftHalf := fmt.Sprintf("%v%v%v%v", leftSegment{}.output(), pwdSegment{}.output(), gitSegment{}.output(), rightSegment{}.output())
	rightHalf := fmt.Sprintf("%v%v%v%v", leftSegment{}.output(), login.output(), host.output(), rightSegment{}.output())

	stretchy.lengthLeft = len(leftHalf)
	stretchy.lengthRight = len(rightHalf)

	mid := fmt.Sprintf("%s", stretchy.output())

	fmt.Print(leftHalf + mid + rightHalf + "\n%% ")
}

func (gp goPrompt) TerminalWidth() int {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 80 // this isn't right, but if term fails, always return 80.   even though we don't have a term (say for testing)
	}
	return int(ws.Col)
}
