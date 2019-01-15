package segments

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type Git struct {
	branch, remoteBranch, direction, dirty, gitString string
	parsed                                            bool
}

var gitHeaderRegexp = regexp.MustCompile(`## (?P<local_branch>\w*)...(?P<remote_branch>\S*)(..(?P<direction>ahead|behind) (?P<direction_count>\d)]|)`)

func (g Git) ColoredOutput() string {
	return g.Output()
}

// Len return length of string without invisible characters counted
func (g Git) Len() int {
	return len(g.Output())
}

func (g *Git) Output() string {
	g.parseGit()

	if g.gitString == "" {
		return ""
	}

	output := fmt.Sprintf(":%v%v %v", g.branch, g.dirty, g.direction)

	return output
}

func (g *Git) runGit() {
	out, err := exec.Command("git", "status", "--porcelain", "--ahead-behind", "-b").Output()

	if err != nil {
		return
	}

	g.gitString = string(out)
}

func (g *Git) parseGit() Git {
	if g.parsed == false {
		g.runGit()
	}

	g.parsed = true

	lines := strings.Split(string(g.gitString), "\n")

	if g.gitString == "" {
		return *g
	}

	gitHeaderRegexp.MatchString(lines[0])
	parts := gitHeaderRegexp.FindAllStringSubmatch(lines[0], -1)

	g.branch = parts[0][1]
	g.remoteBranch = parts[0][2]
	g.direction = parts[0][3]

	if len(lines) > 2 {
		g.dirty = "*"
	}

	scanner := bufio.NewScanner(strings.NewReader(string(g.gitString)))
	scanner.Split(bufio.ScanLines)

	return *g
}
