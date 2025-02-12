package segments

// Segment is the interface all segments must implement.
// They receive a theme map, returning final text or "[ERR]" plus an error if needed.
type Segment interface {
	Render(theme map[string]string) (string, error)
}
