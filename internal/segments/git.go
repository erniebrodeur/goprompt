package segments

import (
	"fmt"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/builders"
)

// Git is for returning a specialized string representing the directories git status
type Git struct {
	branch, remoteBranch, direction, dirty, gitString string
	gitBuilder                                        func() string
}

// NewGit returns an instantiated Git Struct
func NewGit() *Git {
	g := Git{}
	g.gitBuilder = builders.Git
	return &g
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
	g.branch = ""
	g.remoteBranch = ""
	g.direction = ""
	g.dirty = ""
	g.gitString = g.gitBuilder()

	if g.gitString == "" {
		return *g
	}

	lines := strings.Split(strings.TrimRight(g.gitString, "\n"), "\n")
	header := strings.TrimPrefix(lines[0], "## ")

	if len(lines) > 1 {
		g.dirty = "*"
	}

	if strings.HasPrefix(header, "No commits yet on ") {
		g.branch = "No commits yet"
		return *g
	}

	if header == "HEAD (no branch)" {
		g.branch = "HEAD (no branch)"
		return *g
	}

	separator := strings.Index(header, "...")
	if separator == -1 {
		g.branch = header
		return *g
	}

	g.branch = header[:separator]
	remoteStatus := header[separator+3:]
	if statusSeparator := strings.IndexByte(remoteStatus, ' '); statusSeparator >= 0 {
		g.remoteBranch = remoteStatus[:statusSeparator]
		g.direction = directionOutput(remoteStatus[statusSeparator+1:])
	} else {
		g.remoteBranch = remoteStatus
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
