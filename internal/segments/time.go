package segments

import (
	"fmt"
	"time"
)

type CurrentTime time.Time

// ColoredOutput the colorized version
func (t CurrentTime) ColoredOutput() string {
	return t.Output()
}

// Len return length of string without invisible characters counted
func (t CurrentTime) Len() int {
	return len(t.Output())
}

// Output returns the time in a nice format
func (t CurrentTime) Output() string {
	return fmt.Sprintf("%v", time.Now().Format(time.Stamp))
}
