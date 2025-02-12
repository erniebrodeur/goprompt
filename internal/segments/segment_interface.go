package segments

// Segment is the interface all segments must implement.
type Segment interface {
	Render(theme map[string]string) (string, error)
}

