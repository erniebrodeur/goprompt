package segments // import "github.com/erniebrodeur/prompt/internal/segments"

type Segment interface {
	Output() string
	Len() int
	ColoredOutput() string
}
