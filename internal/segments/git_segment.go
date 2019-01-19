package segments

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type Git struct {
	branch, remoteBranch, direction, dirty, gitString string
	runner                                            func() string
}

var gitHeaderRegexp = regexp.MustCompile(`## (?P<local_branch>\w*)...(?P<remote_branch>\S*)(..(?P<direction>ahead|behind) (?P<direction_count>\d)]|)`)

// NewGit returns an instantiated Git Struct
func NewGit() *Git {
	g := Git{}
	g.runner = runGit
	return &g
}

// ColoredOutput returns Output wrapped in a color
func (g *Git) ColoredOutput() string {
	return g.Output()
}

// Len return length of string without invisible characters counted
func (g *Git) Len() int {
	return len(g.Output())
}

// Output returns a string of git info or blank
func (g *Git) Output() string {
	g.parseGit()

	if g.gitString == "" {
		return ""
	}

	return fmt.Sprintf(":%v%v%v", g.branch, g.dirty, g.direction)
}

func runGit() string {
	out, err := exec.Command("git", "status", "--porcelain", "--ahead-behind", "-b").Output()

	if err != nil {
		return ""
	}

	return string(out)
}

func (g *Git) parseGit() Git {
	g.gitString = g.runner()

	if g.gitString == "" {
		return *g
	}

	lines := strings.Split(string(g.gitString), "\n")

	gitHeaderRegexp.MatchString(lines[0])
	parts := gitHeaderRegexp.FindAllStringSubmatch(lines[0], -1)

	g.branch = parts[0][1]
	g.remoteBranch = parts[0][2]
	g.direction = directionOutput(parts[0][3])

	if len(lines) > 2 {
		g.dirty = "*"
	}

	return *g
}

func directionOutput(s string) string {
	if strings.Contains(s, "ahead") {
		return " (push)"
	}

	if strings.Contains(s, "behind") {
		return " (pull)"
	}

	return ""
}
