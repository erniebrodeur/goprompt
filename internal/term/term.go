package term

import (
	"os"
	"regexp"
	"unicode/utf8"

	"golang.org/x/sys/unix"
)

var ansiEscape = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func GetWidth() (int, error) {
    ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
    if err != nil {
        return 0, err
    }
    return int(ws.Col), nil
}

func DisplayWidth(s string) int {
    // Strip ANSI codes
    clean := ansiEscape.ReplaceAllString(s, "")
    // Then count runes
    return utf8.RuneCountInString(clean)
}
