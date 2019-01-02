package segments

import (
	"fmt"
	"os"
)

type shellLevelSegment struct{}

func (s shellLevelSegment) output() string {
	if os.Getenv("USER") == "root" {
		return "#"
	}
	fmt.Println(s.len())
	return "%"
}

func (s shellLevelSegment) len() int {
	return 1
}
