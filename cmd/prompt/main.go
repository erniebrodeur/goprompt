package main // import "github.com/erniebrodeur/prompt"

import (
	"fmt"

	"github.com/erniebrodeur/prompt/internal/segments"

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
	login := segments.Login{}
	host := segments.host{}
	stretchy := segments.Stretchy{}

	leftHalf := fmt.Sprintf("%v%v%v%v", segments.Left{}.output(), segments.Pwd{}.output(), segments.Git{}.output(), segments.Right{}.output())
	rightHalf := fmt.Sprintf("%v%v%v%v", segments.Left{}.output(), login.output(), host.output(), segments.Right{}.output())

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
