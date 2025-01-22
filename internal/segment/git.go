package segment

import (
	"github.com/erniebrodeur/goprompt/internal/model"
	"github.com/go-git/go-git/v5"
)

// GitSegment is a segment for displaying Git info.
type GitSegment struct{}

// NewGitSegment returns a new GitSegment.
func NewGitSegment() *GitSegment {
	return &GitSegment{}
}

// Name identifies this segment by name.
func (g *GitSegment) Name() string {
	return "git"
}

// Enabled determines if we're in a Git repo.
func (g *GitSegment) Enabled() bool {
	_, err := git.PlainOpen(".")
	return err == nil
}

// Output fetches the current branch and dirty status.
func (g *GitSegment) Output() (model.SegmentOutput, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return model.SegmentOutput{}, err
	}
	headRef, err := repo.Head()
	if err != nil {
		return model.SegmentOutput{}, err
	}

	branch := headRef.Name().Short() // e.g. "main" or "HEAD"
	if !headRef.Name().IsBranch() {
		branch = "(detached)"
	}

	wt, err := repo.Worktree()
	if err != nil {
		return model.SegmentOutput{}, err
	}
	status, err := wt.Status()
	if err != nil {
		return model.SegmentOutput{}, err
	}

	dirty := ""
	if !status.IsClean() {
		dirty = "*"
	}

	return model.SegmentOutput{
		Name: "git",
		Text: branch + dirty,
	}, nil
}
