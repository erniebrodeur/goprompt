package segments // import "github.com/erniebrodeur/goprompt/internal/segments"

type Segment interface {
	ColoredOutput() string
	Len() int
	Output() string
}
