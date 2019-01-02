// UTF-8

package segments

type Right struct{}

func (s Right) Output() string {
	return " ├"
}

func (s Right) Len() int {
	return 2
}
