package builders

import (
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

// TerminalWidth is a build function for getting the current terminal width
func TerminalWidth() int {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)

	if err == nil {
		return int(ws.Col)
	}

	return 80
}

// Pwd is a build function for the current PWD
func Pwd() string {
	output, _ := os.Getwd()

	return output
}

// Git is a build function to return the appropriate git string
func Git() string {
	out, err := exec.Command("git", "status", "--porcelain", "--ahead-behind", "-b").Output()

	if err != nil {
		return ""
	}

	return string(out)
}
