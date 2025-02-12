package segments

import (
	"errors"

	gogit "github.com/go-git/go-git/v5"
)

type GitSegment struct {
	// Possibly store states like isDirty, branchName, etc.
}

func (g *GitSegment) Render(theme map[string]string) (string, error) {
	// Minimal example: try to open local git repo
	repo, err := gogit.PlainOpen(".")
	if err != nil {
		// Not a repo => return empty, not an error
		return "", nil
	}
	head, err := repo.Head()
	if err != nil {
		return "[ERR]", err
	}
	branchName := head.Name().Short()

	// Fake "dirty" check if needed (skipped for brevity)
	// If we can't detect properly, let's just say no error

	if branchName == "" {
		return "[ERR]", errors.New("detached or unknown branch")
	}
	return branchName, nil
}
