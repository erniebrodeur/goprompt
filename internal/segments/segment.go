package segments // import "github.com/erniebrodeur/goprompt"

type Segment interface {
	ColoredOutput() string
	Len() int
	Output() string
}
