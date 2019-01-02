// UTF-8

package main

type leftSegment struct{}

func (s leftSegment) output() string {
	return "┤ "
}

func (s leftSegment) len() int {
	return 2
}
