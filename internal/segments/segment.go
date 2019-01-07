package segments // import "github.com/erniebrodeur/prompt/internal/segments"

type Segment interface {
	ColoredOutput() string
	Len() int
	Output() string
}
