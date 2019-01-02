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
	host := segments.Host{}
	stretchy := segments.Stretchy{}

	leftHalf := fmt.Sprintf("%v%v%v%v", segments.Left{}.Output(), segments.Pwd{}.Output(), segments.Git{}.Output(), segments.Right{}.Output())
	rightHalf := fmt.Sprintf("%v%v%v%v", segments.Left{}.Output(), login.Output(), host.Output(), segments.Right{}.Output())

	stretchy.LengthLeft = len(leftHalf)
	stretchy.LengthRight = len(rightHalf)

	mid := "fmt.Sprintf(\"%s\", stretchy.Output())"

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
