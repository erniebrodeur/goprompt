package segments

// Segment is an interface for a prompt segment
type Segment interface {
	ColoredOutput() string
	Len() int
	Output() string
}
