package segments

type Mid struct {
	Left bool
}

func (m Mid) ColoredOutput() string {
	return m.Output()
}

// Len return length of string without invisible characters counted
func (m Mid) Len() int {
	return len(m.Output())
}

func (m Mid) Output() string {
	return "─"
}
