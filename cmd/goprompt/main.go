package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/erniebrodeur/goprompt/internal/builders"
	"github.com/erniebrodeur/goprompt/internal/segments"
	prompttheme "github.com/erniebrodeur/goprompt/internal/theme"
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
	os.Exit(run(os.Args[1:], os.Getenv("GOPROMPT_THEME"), os.Stdout, os.Stderr))
}

type options struct {
	showVersion bool
	showStatus  bool
	themeName   string
}

func run(args []string, environmentTheme string, stdout, stderr io.Writer) int {
	parsedOptions, err := parseOptions(args, stderr)
	if err != nil {
		return 2
	}

	selectedTheme, err := prompttheme.Resolve(parsedOptions.themeName, environmentTheme)
	if err != nil {
		fmt.Fprintf(stderr, "goprompt: %v\n", err)
		return 2
	}

	if parsedOptions.showVersion {
		fmt.Fprintln(stdout, "Version:", Version)
		return 0
	}

	if parsedOptions.showStatus {
		status(stdout, selectedTheme)
		return 0
	}

	mid.Count = builders.TerminalWidth() -
		git.Len() - pwd.Len() - host.Len() -
		login.Len() - shell.Len() - currentTime.Len() - 17 // special + spaces

	output(stdout, selectedTheme)
	return 0
}

func parseOptions(args []string, output io.Writer) (options, error) {
	var parsedOptions options
	flags := flag.NewFlagSet("goprompt", flag.ContinueOnError)
	flags.SetOutput(output)
	flags.BoolVar(&parsedOptions.showVersion, "v", false, "Show version")
	flags.BoolVar(&parsedOptions.showVersion, "version", false, "Show version")
	flags.BoolVar(&parsedOptions.showStatus, "s", false, "Show segment status")
	flags.BoolVar(&parsedOptions.showStatus, "status", false, "Show segment status")
	flags.StringVar(&parsedOptions.themeName, "t", "", "Select theme")
	flags.StringVar(&parsedOptions.themeName, "theme", "", "Select theme")

	return parsedOptions, flags.Parse(args)
}

type promptParts struct {
	pwd   string
	git   string
	login string
	host  string
	mid   string
	clock string
	date  string
	shell string
}

func output(writer io.Writer, selectedTheme prompttheme.Theme) {
	renderPrompt(writer, selectedTheme, collectPromptParts())
}

func status(writer io.Writer, selectedTheme prompttheme.Theme) {
	renderStatus(writer, selectedTheme, collectPromptParts())
}

func collectPromptParts() promptParts {
	clock, date := currentTime.Parts()
	return promptParts{
		pwd:   pwd.Output(),
		git:   git.Output(),
		login: login.Output(),
		host:  host.Output(),
		mid:   mid.Output(),
		clock: clock,
		date:  date,
		shell: shell.Output(),
	}
}

func renderPrompt(writer io.Writer, selectedTheme prompttheme.Theme, parts promptParts) {
	left := selectedTheme.Frame.Apply("─┤ ")
	right := selectedTheme.Frame.Apply(" ├─")
	fmt.Fprintf(writer, "%v%v%v%v%v%v%v%v%v%v%v%v%v%v\n%v ",
		left,
		selectedTheme.Primary.Apply(parts.pwd),
		selectedTheme.VCS.Apply(parts.git),
		right,
		left,
		selectedTheme.Primary.Apply(parts.login),
		selectedTheme.VCS.Apply(parts.host),
		right,
		selectedTheme.Frame.Apply(parts.mid),
		left,
		selectedTheme.Clock.Apply(parts.clock),
		selectedTheme.Frame.Apply(" ─ "),
		selectedTheme.Clock.Apply(parts.date),
		right,
		selectedTheme.Prompt.Apply(parts.shell),
	)
}

func renderStatus(writer io.Writer, selectedTheme prompttheme.Theme, parts promptParts) {
	fmt.Fprintf(writer, "git:           %v\n", selectedTheme.VCS.Apply(parts.git))
	fmt.Fprintf(writer, "host:          %v\n", selectedTheme.VCS.Apply(parts.host))
	fmt.Fprintf(writer, "login:         %v\n", selectedTheme.Primary.Apply(parts.login))
	fmt.Fprintf(writer, "mid:           %v\n", selectedTheme.Frame.Apply(parts.mid))
	fmt.Fprintf(writer, "pwd:           %v\n", selectedTheme.Primary.Apply(parts.pwd))
	fmt.Fprintf(writer, "time:          %v%v%v\n",
		selectedTheme.Clock.Apply(parts.clock),
		selectedTheme.Frame.Apply(" ─ "),
		selectedTheme.Clock.Apply(parts.date),
	)
	fmt.Fprintf(writer, "shell:         %v\n", selectedTheme.Prompt.Apply(parts.shell))
}
