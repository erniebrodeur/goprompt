package segments

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Git struct {
	branch       string
	remoteBranch string
	direction    string
	directionBy  int
	dirty        bool
}

var gitHeaderRegexp = regexp.MustCompile(`## (?P<local_branch>\w*)...(?P<remote_branch>\S*)(..(?P<direction>ahead|behind) (?P<direction_count>\d)]|)`)

func (g Git) ColoredOutput() string {
	return g.Output()
}

// Len return length of string without invisible characters counted
func (g Git) Len() int {
	return len(g.Output())
}

func (g Git) Output() string {
	g = Git.parse(g)
	output := fmt.Sprintf(":%v", g.branch)

	if g.dirty {
		output += "*"
	}

	if len(g.direction) != 0 {
		if g.direction == "ahead" {
			output += " (push)"
		} else {
			output += " (pull)"
		}
	}

	return output
}

func (g Git) parse() Git {
	out, err := exec.Command("git", "status", "--porcelain", "--ahead-behind", "-b").Output()

	if err != nil {
		return g
	}

	lines := strings.Split(string(out), "\n")

	g.dirty = false
	gitHeaderRegexp.MatchString(lines[0])

	parts := gitHeaderRegexp.FindAllSubmatch([]byte(lines[0]), -1)
	g.branch = string(parts[0][1])
	g.remoteBranch = string(parts[0][2])
	g.direction = string(parts[0][4])
	g.directionBy, _ = strconv.Atoi(string(parts[0][5]))

	if len(lines) > 2 {
		g.dirty = true
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	scanner.Split(bufio.ScanLines)

	return g
}
