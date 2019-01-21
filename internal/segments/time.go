package segments

import (
	"fmt"
	"time"

	"github.com/mgutz/ansi"
)

type CurrentTime time.Time

// ColoredOutput the colorized version
func (t CurrentTime) ColoredOutput() string {
	yellow := ansi.ColorFunc("cyan+h:black")
	return yellow(t.Output())
}

// Len return length of string without invisible characters counted
func (t CurrentTime) Len() int {
	return len(t.Output()) - 2 // special char consideration
}

// Output returns the time in a nice format
func (t CurrentTime) Output() string {
	return fmt.Sprintf("%v", time.Now().Format("03:04pm ─ 1/2"))
}
