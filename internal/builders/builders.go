package builders

import (
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func TerminalWidth() int {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)

	if err == nil {
		return int(ws.Col)
	}

	return 80
}

func Pwd() string {
	output, _ := os.Getwd()

	return output
}

func Git() string {
	out, err := exec.Command("git", "status", "--porcelain", "--ahead-behind", "-b").Output()

	if err != nil {
		return ""
	}

	return string(out)
}
