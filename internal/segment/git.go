package segment

import (
	"bytes"
	"os/exec"
	"strings"
)

type GitSegment struct{}

func NewGitSegment() *GitSegment {
    return &GitSegment{}
}

func (g *GitSegment) Name() string  { return "git" }
func (g *GitSegment) Enabled() bool {
    cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
    return cmd.Run() == nil
}

func (g *GitSegment) Output() (SegmentOutput, error) {
    branch, err := runCmd("git", "rev-parse", "--abbrev-ref", "HEAD")
    if err != nil {
        return SegmentOutput{}, err
    }
    branch = strings.TrimSpace(branch)

    status, err := runCmd("git", "status", "--porcelain")
    if err != nil {
        return SegmentOutput{}, err
    }
    dirty := ""
    if len(strings.TrimSpace(status)) > 0 {
        dirty = "*"
    }

    return SegmentOutput{
        Name: "git",
        Text: branch + dirty,
    }, nil
}

func runCmd(name string, args ...string) (string, error) {
    var out bytes.Buffer
    cmd := exec.Command(name, args...)
    cmd.Stdout = &out
    cmd.Stderr = &out
    err := cmd.Run()
    return out.String(), err
}
