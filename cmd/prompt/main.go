package main // import "github.com/erniebrodeur/prompt"

import (
	"fmt"

	"github.com/erniebrodeur/prompt/internal/segments"
)

var login = segments.Login{}
var pwd = segments.Pwd{}
var host = segments.Host{}
var git = segments.Git{}
var mid = segments.Mid{}
var left = segments.Bookend{Left: true}
var right = segments.Bookend{}

func main() {
	status()
}

func status() {
	fmt.Printf("login:         %v\n", login.Output())
	fmt.Printf("pwd:           %v\n", pwd.Output())
	fmt.Printf("host:          %v\n", host.Output())
	fmt.Printf("git:           %v\n", git.Output())
	fmt.Printf("mid:           %v\n", mid.Output())
	fmt.Printf("bookend left:  %v\n", left.Output())
	fmt.Printf("bookend right: %v\n", right.Output())
}
