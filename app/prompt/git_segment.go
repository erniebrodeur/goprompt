package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var gitHeaderRegexp = regexp.MustCompile(`## (?P<local_branch>\w*)...(?P<remote_branch>\S*)(..(?P<direction>ahead|behind) (?P<direction_count>\d)]|)`)

type gitSegment struct {
	branch       string
	remoteBranch string
	direction    string
	directionBy  int
	dirty        bool
}

func (s gitSegment) output() string {
	s = gitSegment.parse(s)
	output := fmt.Sprintf(":%s", s.branch)

	if s.dirty {
		output += "*"
	}

	if len(s.direction) != 0 {
		if s.direction == "ahead" {
			output += " (push)"
		} else {
			output += " (pull)"
		}
	}

	return output
}

func (s gitSegment) len() int {
	return len(gitSegment.output(s))
}

func (s gitSegment) parse() gitSegment {
	out, err := exec.Command("git", "status", "--porcelain", "--ahead-behind", "-b").Output()

	if err != nil {
		return s
	}

	lines := strings.Split(string(out), "\n")

	s.dirty = false
	gitHeaderRegexp.MatchString(lines[0])

	parts := gitHeaderRegexp.FindAllSubmatch([]byte(lines[0]), -1)
	s.branch = string(parts[0][1])
	s.remoteBranch = string(parts[0][2])
	s.direction = string(parts[0][4])
	s.directionBy, _ = strconv.Atoi(string(parts[0][5]))

	if len(lines) > 2 {
		s.dirty = true
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	scanner.Split(bufio.ScanLines)

	return s
}
