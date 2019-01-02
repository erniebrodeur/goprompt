package segments

import (
	"fmt"
	"os"
)

type ShellLevel struct{}

func (s ShellLevel) Output() string {
	if os.Getenv("USER") == "root" {
		return "#"
	}
	fmt.Println(s.Len())
	return "%"
}

func (s ShellLevel) Len() int {
	return 1
}
