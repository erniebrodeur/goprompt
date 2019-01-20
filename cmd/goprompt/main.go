package main // import "github.com/erniebrodeur/goprompt/cmd/goprompt"

import (
	"fmt"

	"github.com/erniebrodeur/goprompt/internal/builders"
	"github.com/erniebrodeur/goprompt/internal/segments"
)

var (
	git         = segments.NewGit()
	host        = segments.Host{}
	left        = segments.Bookend{Left: true}
	login       = segments.Login{}
	mid         = segments.Mid{}
	pwd         = segments.NewPwd()
	right       = segments.Bookend{}
	shell       = segments.ShellLevel{}
	currentTime = segments.CurrentTime{}
)

func main() {
	mid.Count = builders.TerminalWidth() -
		(left.Len() * 2) - (right.Len() * 2) -
		git.Len() - pwd.Len() - host.Len() -
		login.Len() - shell.Len() - currentTime.Len() - 7 // space count - 1(?)
	output()
}

func output() {
	l := left.Output()
	r := right.Output()
	fmt.Printf("%v %v%v %v%v%v %v%v %v%v %v %v\n%v ", l, pwd.Output(), git.Output(), r, mid.Output(), l, login.Output(), host.Output(), r, l, currentTime.Output(), r, shell.Output())
}

func status() {
	fmt.Printf("bookend left:  %v\n", left.Output())
	fmt.Printf("bookend right: %v\n", right.Output())
	fmt.Printf("git:           %v\n", git.Output())
	fmt.Printf("host:          %v\n", host.Output())
	fmt.Printf("login:         %v\n", login.Output())
	fmt.Printf("mid:           %v\n", mid.Output())
	fmt.Printf("pwd:           %v\n", pwd.Output())
	fmt.Printf("time:          %v\n", currentTime.Output())
	fmt.Printf("shell:         %v\n", shell.Output())
}
