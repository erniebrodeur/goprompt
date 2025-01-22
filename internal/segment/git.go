package segment

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/colors"
)

type GitSegment struct{}

func NewGitSegment() *GitSegment {
    return &GitSegment{}
}

func (g *GitSegment) Name() string { return "git" }

// If we're in a git repo
func (g *GitSegment) Enabled() bool {
    cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
    return cmd.Run() == nil
}

func (g *GitSegment) Render() (string, error) {
    branch, err := runCmd("git", "rev-parse", "--abbrev-ref", "HEAD")
    if err != nil {
        return "", err
    }
    branch = strings.TrimSpace(branch)

    status, err := runCmd("git", "status", "--porcelain")
    if err != nil {
        return "", err
    }

    dirty := ""
    if len(strings.TrimSpace(status)) > 0 {
        dirty = "*"
    }

    return fmt.Sprintf("%s%s%s%s ", colors.Yellow, branch, dirty, colors.Reset), nil
}

func runCmd(name string, args ...string) (string, error) {
    var out bytes.Buffer
    cmd := exec.Command(name, args...)
    cmd.Stdout = &out
    cmd.Stderr = &out
    err := cmd.Run()
    return out.String(), err
}
