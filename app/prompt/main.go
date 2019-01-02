package main // github.com/erniebrodeur/prompt

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// TerminalWidth comment
var TerminalWidth = 0

// Prompt comment
type Prompt interface {
	Output()
}

type PromptData struct{}

func (p PromptData) Output() string {
	login := loginSegment{}
	host := hostSegment{}
	stretchy := stretchySegment{}

	leftHalf := fmt.Sprintf("%v%v%v%v", leftSegment{}.output(), pwdSegment{}.output(), gitSegment{}.output(), rightSegment{}.output())
	rightHalf := fmt.Sprintf("%v%v%v%v", leftSegment{}.output(), login.output(), host.output(), rightSegment{}.output())

	stretchy.lengthLeft = len(leftHalf)
	stretchy.lengthRight = len(rightHalf)

	mid := "fmt.Sprintf(\"%s\", stretchy.output())"

	return fmt.Sprint(leftHalf + mid + rightHalf + "\n%% ")
}

func initialize() {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		TerminalWidth = 80 // this isn't right, but if term fails, always return 80.   even though we don't have a term (say for testing)
	}

	TerminalWidth = int(ws.Col)
}

func main() {
	fmt.Print(PromptData{}.Output())
}
