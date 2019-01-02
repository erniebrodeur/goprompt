// UTF-8

package segments

type Left struct{}

func (s Left) Output() string {
	return "┤ "
}

func (s Left) Len() int {
	return 2
}
