package main // import "github.com/erniebrodeur/prompt"

import (
	"fmt"

	"github.com/erniebrodeur/prompt/internal/segments"
	"golang.org/x/sys/unix"
)

var git = segments.Git{}
var host = segments.Host{}
var left = segments.Bookend{Left: true}
var login = segments.Login{}
var mid = segments.Mid{}
var pwd = segments.Pwd{}
var right = segments.Bookend{}
var shell = segments.ShellLevel{}

func main() {
	var terminalWidth = 40

	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)

	if err == nil {
		terminalWidth = int(ws.Col)
	}

	mid.Count = terminalWidth
	status()
}

func status() {
	fmt.Printf("bookend left:  %v\n", left.Output())
	fmt.Printf("bookend right: %v\n", right.Output())
	fmt.Printf("git:           %v\n", git.Output())
	fmt.Printf("host:          %v\n", host.Output())
	fmt.Printf("login:         %v\n", login.Output())
	fmt.Printf("mid:           %v\n", mid.Output())
	fmt.Printf("pwd:           %v\n", pwd.Output())
	fmt.Printf("shell:         %v\n", shell.Output())
}
