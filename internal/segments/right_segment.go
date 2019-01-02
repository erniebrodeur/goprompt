// UTF-8

package segments

type Right struct{}

func (s Right) output() string {
	return " ├"
}

func (s Right) len() int {
	return 2
}
