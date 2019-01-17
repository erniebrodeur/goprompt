package segments

type Segment interface {
	ColoredOutput() string
	Len() int
	Output() string
}
