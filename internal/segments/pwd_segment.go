package segments

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

type Pwd struct{}

func (p Pwd) ColoredOutput() string {
	return p.Output()
}

// Len return length of string without invisible characters counted
func (p Pwd) Len() int {
	return len(p.Output())
}

func (p Pwd) Output() string {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	var terminalWidth int

	if err == nil {
		terminalWidth = int(ws.Col)
	}

	fmt.Println(terminalWidth)
	output, _ := os.Getwd()
	return output
}
