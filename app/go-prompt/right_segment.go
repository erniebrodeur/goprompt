// UTF-8

package main

type rightSegment struct{}

func (s rightSegment) output() string {
	return " ├"
}

func (s rightSegment) len() int {
	return 2
}
