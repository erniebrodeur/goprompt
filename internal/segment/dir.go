package segment

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/model"
)

type DirSegment struct{}

func NewDirSegment() *DirSegment {
    return &DirSegment{}
}

func (d *DirSegment) Name() string   { return "directory" }
func (d *DirSegment) Enabled() bool  { return true }

// Output returns a short path plus the current git branch (if any).
// Example: ".../goprompt:main*" (if the repo is dirty)
func (d *DirSegment) Output() (model.SegmentOutput, error) {
    cwd, err := os.Getwd()
    if err != nil {
        return model.SegmentOutput{}, err
    }

    // 1) Shorten path to ".../<lastDir>"
    shortened := shortenPath(cwd)

    // 2) If in a git repo, append ":branch*"
    if isInGitRepo() {
        branch, dirty := currentBranchAndDirty()
        if branch != "" {
            shortened += ":" + branch
            if dirty {
                shortened += "*"
            }
        }
    }

    return model.SegmentOutput{
        Name: "directory",
        Text: shortened,
    }, nil
}

// shortenPath converts /Users/you/Projects/goprompt -> .../Projects/goprompt
// If there's only one or two segments, it won't prepend "..."
func shortenPath(full string) string {
    // Use user's home as ~ if relevant
    usr, _ := user.Current()
    home := ""
    if usr != nil {
        home = usr.HomeDir
    }
    if home != "" && strings.HasPrefix(full, home) {
        full = "~" + strings.TrimPrefix(full, home)
    }

    // Split path, ignoring ~ as part
    parts := strings.Split(full, string(os.PathSeparator))
    if len(parts) < 3 {
        // e.g. "~" or "~/<dir>"
        return full
    }

    // Rejoin the last two parts, prepend ".../"
    lastTwo := filepath.Join(parts[len(parts)-2], parts[len(parts)-1])
    return ".../" + lastTwo
}

// isInGitRepo checks if we're inside a Git work tree
func isInGitRepo() bool {
    cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
    return cmd.Run() == nil
}

// currentBranchAndDirty returns (branchName, isDirty)
func currentBranchAndDirty() (string, bool) {
    branch, err := runCmd("git", "rev-parse", "--abbrev-ref", "HEAD")
    if err != nil {
        return "", false
    }
    branch = strings.TrimSpace(branch)

    status, err := runCmd("git", "status", "--porcelain")
    if err != nil {
        return branch, false
    }
    dirty := (len(strings.TrimSpace(status)) > 0)
    return branch, dirty
}
