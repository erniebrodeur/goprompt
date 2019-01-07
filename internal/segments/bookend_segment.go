package segments

type Bookend struct {
	Left bool
}

func (b Bookend) ColoredOutput() string {
	return b.Output()
}

// Len return length of string without invisible characters counted
func (b Bookend) Len() int {
	return 1
}

func (b Bookend) Output() string {
	if b.Left {
		return "┤"
	}
	return "├"
}
