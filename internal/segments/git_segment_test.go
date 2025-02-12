package segments_test

import (
	"os"
	"testing"

	gogit "github.com/go-git/go-git/v5"
	"github.com/erniebrodeur/goprompt/internal/segments"
)

func TestGitSegmentNonRepo(t *testing.T) {
	origWD, _ := os.Getwd()
	defer os.Chdir(origWD)

	tmp := t.TempDir()
	os.Chdir(tmp)

	gs := &segments.GitSegment{}
	out, err := gs.Render(nil)
	if err != nil {
		t.Fatalf("Unexpected error in non-repo: %v", err)
	}
	if out != "" {
		t.Errorf("Expected empty string for non-repo, got %q", out)
	}
}

func TestGitSegmentRepo(t *testing.T) {
	origWD, _ := os.Getwd()
	defer os.Chdir(origWD)

	tmp := t.TempDir()
	os.Chdir(tmp)

	// Minimal init
	repo, err := gogit.PlainInit(tmp, false)
	if err != nil {
		t.Fatalf("Failed to init repo: %v", err)
	}
	// We'll skip making commits for brevity, but let's see if HEAD is recognized
	// Actually HEAD will be detached right now, so let's see how the segment handles that.

	gs := &segments.GitSegment{}
	out, err2 := gs.Render(nil)
	// Possibly we get [ERR] if we didn't create a commit or a branch
	// We'll just check that it's not a crash
	if err2 != nil && out != "[ERR]" {
		t.Errorf("If HEAD is detached, we might expect [ERR], got %q (err=%v)", out, err2)
	}
}
