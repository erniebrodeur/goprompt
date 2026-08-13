package segments

// Segment is an interface for a prompt segment
type Segment interface {
	Len() int
	Output() string
}
