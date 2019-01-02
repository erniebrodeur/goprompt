// UTF-8

package segments

type Left struct{}

func (s Left) output() string {
	return "┤ "
}

func (s Left) len() int {
	return 2
}
