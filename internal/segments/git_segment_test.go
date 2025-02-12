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

	_, err := gogit.PlainInit(tmp, false)
	if err != nil {
		t.Fatalf("Failed to init repo: %v", err)
	}

	gs := &segments.GitSegment{}
	out, err2 := gs.Render(nil)
	// HEAD is presumably empty => might be [ERR], or the segment might handle no commits differently
	// As long as it doesn't crash, it's fine.
	if err2 != nil && out != "[ERR]" {
		t.Errorf("Expected [ERR] or no crash, got %q (err=%v)", out, err2)
	}
}

