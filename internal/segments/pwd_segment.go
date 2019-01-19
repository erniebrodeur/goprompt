package segments

import (
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

// Pwd is for returning the current directory
type Pwd struct {
	terminalWidthBuilder func() int
	pwdBuilder           func() string
}

// NewPwd will generate a new copy of pwd{} with default builders
func NewPwd() *Pwd {
	p := Pwd{}
	p.terminalWidthBuilder = buildTerminalWidth
	p.pwdBuilder = buildPwd
	return &p
}

func buildTerminalWidth() int {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)

	if err == nil {
		return int(ws.Col)
	}

	return 80
}

func buildPwd() string {
	output, _ := os.Getwd()

	return output
}

// ColoredOutput returns a color wrapped copy
func (p Pwd) ColoredOutput() string {
	return p.Output()
}

// Len return length of string without invisible characters counted
func (p Pwd) Len() int {
	return len(p.Output())
}

// Output returns a specially modified pwd for space constraints
func (p Pwd) Output() string {
	parts := strings.Split(buildPwd(), "/")
	outputLen := 0
	end := 0

	for i := len(parts) - 1; i >= 0; i-- {
		outputLen += len(parts[i]) + 1 // for the / char

		if outputLen > buildTerminalWidth()/4 {
			end = i
			return ".../" + strings.Join(parts[end:len(parts)], "/")
		}
	}
	return buildPwd()
}
