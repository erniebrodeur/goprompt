package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/erniebrodeur/goprompt/internal/builders"
	"github.com/erniebrodeur/goprompt/internal/segments"
	"github.com/mgutz/ansi"
)

var (
	git         = segments.NewGit()
	host        = segments.Host{}
	login       = segments.Login{}
	mid         = segments.Mid{}
	pwd         = segments.NewPwd()
	shell       = segments.ShellLevel{}
	currentTime = segments.CurrentTime{}
)

func main() {
	showVersion, showStatus, err := parseOptions(os.Args[1:])
	if err != nil {
		os.Exit(2)
	}

	if showVersion {
		fmt.Println("Version:", Version)
		os.Exit(0)
	}

	if showStatus {
		status()
		os.Exit(0)
	}

	mid.Count = builders.TerminalWidth() -
		git.Len() - pwd.Len() - host.Len() -
		login.Len() - shell.Len() - currentTime.Len() - 17 // special + spaces

	output()
}

func parseOptions(args []string) (showVersion, showStatus bool, err error) {
	flags := flag.NewFlagSet("goprompt", flag.ContinueOnError)
	flags.BoolVar(&showVersion, "v", false, "Show version")
	flags.BoolVar(&showVersion, "version", false, "Show version")
	flags.BoolVar(&showStatus, "s", false, "Show segment status")
	flags.BoolVar(&showStatus, "status", false, "Show segment status")
	err = flags.Parse(args)
	return
}

func output() {
	left := ansi.ColorFunc("blue+h")("─┤ ")
	right := ansi.ColorFunc("blue+h")(" ├─")
	fmt.Printf("%v%v%v%v%v%v%v%v%v%v%v%v\n%v ",
		left, pwd.ColoredOutput(), git.ColoredOutput(), right,
		left, login.ColoredOutput(), host.ColoredOutput(), right,
		mid.ColoredOutput(),
		left, currentTime.ColoredOutput(), right,
		shell.Output())
}

func status() {
	fmt.Printf("git:           %v\n", git.ColoredOutput())
	fmt.Printf("host:          %v\n", host.ColoredOutput())
	fmt.Printf("login:         %v\n", login.ColoredOutput())
	fmt.Printf("mid:           %v\n", mid.ColoredOutput())
	fmt.Printf("pwd:           %v\n", pwd.ColoredOutput())
	fmt.Printf("time:          %v\n", currentTime.ColoredOutput())
	fmt.Printf("shell:         %v\n", shell.Output())
}
