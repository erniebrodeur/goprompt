package segments

import (
	"fmt"
	"time"
)

// CurrentTime is a formated representation of the current time
type CurrentTime time.Time

// Len return length of string without invisible characters counted
func (t CurrentTime) Len() int {
	return len(t.Output()) - 2 // special char consideration
}

// Parts returns the clock and date portions of the time segment.
func (t CurrentTime) Parts() (string, string) {
	now := time.Now()
	return now.Format("03:04pm"), now.Format("1/2")
}

// Output returns the time in a nice format
func (t CurrentTime) Output() string {
	clock, date := t.Parts()
	return fmt.Sprintf("%v ─ %v", clock, date)
}
