package segments

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/builders"
	"github.com/mgutz/ansi"
)

// Git is for returning a specialized string representing the directories git status
type Git struct {
	branch, remoteBranch, direction, dirty, gitString string
	gitBuilder                                        func() string
}

// var gitHeaderRegexp = regexp.MustCompile(`## (?P<local_branch>(\S*|HEAD \(no branch\)))(\.\.\.|$)(?P<remote_branch>\S*)(..(?P<direction>ahead|behind) (?P<direction_count>\d)]|)`)
var gitHeaderRegexp = regexp.MustCompile(`## (\S*)(?:\.\.\.(\S*))( (?:\[(ahead|behind) (\d+))\]|)`)

// NewGit returns an instantiated Git Struct
func NewGit() *Git {
	g := Git{}
	g.gitBuilder = builders.Git
	return &g
}

// ColoredOutput returns Output wrapped in a color
func (g *Git) ColoredOutput() string {
	return ansi.ColorFunc("yellow+h")(g.Output())
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

func (g *Git) parseGit() Git {
	g.gitString = g.gitBuilder()

	if g.gitString == "" {
		return *g
	}

	if strings.Contains(g.gitString, "HEAD (no branch)") {
		g.branch = "HEAD (no branch)"
		return *g
	}

	if !strings.Contains(g.gitString, "...") {
		g.branch = strings.Split(g.gitString, "## ")[1]
		return *g
	}

	lines := strings.Split(g.gitString, "\n")

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
