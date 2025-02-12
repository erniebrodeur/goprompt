package segments

import (
	"errors"

	gogit "github.com/go-git/go-git/v5"
)

type GitSegment struct {
}

func (g *GitSegment) Render(theme map[string]string) (string, error) {
	repo, err := gogit.PlainOpen(".")
	if err != nil {
		// Not a git repo => empty output
		return "", nil
	}
	head, err := repo.Head()
	if err != nil {
		return "[ERR]", err
	}

	branchName := head.Name().Short()
	if branchName == "" {
		return "[ERR]", errors.New("detached HEAD or unknown branch")
	}
	// Could check dirty state, skipping for brevity
	return branchName, nil
}
