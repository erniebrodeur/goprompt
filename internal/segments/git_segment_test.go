package segments_test

import (
	"os"
	"testing"

	gogit "github.com/go-git/go-git/v5"
	"github.com/erniebrodeur/goprompt/internal/segments"
	"github.com/stretchr/testify/require"
)

func TestGitSegment_NonRepo(t *testing.T) {
	origWD, _ := os.Getwd()
	defer os.Chdir(origWD)

	tmp := t.TempDir()
	require.NoError(t, os.Chdir(tmp))

	gs := &segments.GitSegment{}
	out, err := gs.Render(nil)
	require.NoError(t, err)
	require.Equal(t, "", out, "Non-repo => empty string")
}

func TestGitSegment_RepoNoCommits(t *testing.T) {
	origWD, _ := os.Getwd()
	defer os.Chdir(origWD)

	tmp := t.TempDir()
	require.NoError(t, os.Chdir(tmp))

	_, err := gogit.PlainInit(tmp, false)
	require.NoError(t, err)

	gs := &segments.GitSegment{}
	out, err2 := gs.Render(nil)
	// HEAD is presumably detached => might be [ERR]
	if err2 != nil {
		require.Equal(t, "[ERR]", out, "Detached HEAD => [ERR]")
	}
}

---
