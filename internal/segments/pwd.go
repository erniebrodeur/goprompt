package segments

import (
	"os"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/builders"
	"github.com/mgutz/ansi"
)

// Pwd is for returning the current directory
type Pwd struct {
	terminalWidthBuilder func() int
	pwdBuilder           func() string
}

// NewPwd will generate a new copy of pwd{} with default builders
func NewPwd() *Pwd {
	p := Pwd{}
	p.terminalWidthBuilder = builders.TerminalWidth
	p.pwdBuilder = builders.Pwd
	return &p
}

// ColoredOutput returns a color wrapped copy of Output
func (p Pwd) ColoredOutput() string {
	return ansi.ColorFunc("green+h:black")(p.Output())
}

// Len return length of string without invisible characters counted
func (p Pwd) Len() int {
	return len(p.Output())
}

// Output returns a specially modified pwd for space constraints
func (p Pwd) Output() string {
	pwd := p.pwdBuilder()
	homeDir := os.Getenv("HOME")

	fixedPwd := strings.Replace(pwd, homeDir, "~", 1)
	parts := strings.Split(fixedPwd, "/")
	outputLen := 0
	end := 0

	for i := len(parts) - 1; i >= 0; i-- {
		outputLen += len(parts[i]) + 1 // for the / char

		if outputLen > p.terminalWidthBuilder()/8 {
			end = i
			return ".../" + strings.Join(parts[end:], "/")
		}
	}

	return fixedPwd
}
