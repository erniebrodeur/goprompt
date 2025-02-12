package segments_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/erniebrodeur/goprompt/internal/segments"
)

func TestDirSegment(t *testing.T) {
	origWD, _ := os.Getwd()
	defer os.Chdir(origWD)

	tempDir := t.TempDir()
	os.Chdir(tempDir)

	d := &segments.DirSegment{ShowComponents: 1}
	out, err := d.Render(map[string]string{"dir.normal": "#FD971F"})
	if err != nil {
		t.Fatalf("Render returned error: %v", err)
	}
	want := filepath.Base(tempDir)
	if !strings.Contains(out, want) {
		t.Errorf("Expected DirSegment output to contain %q, got %q", want, out)
	}
}

func TestDirSegmentError(t *testing.T) {
	// Hard to force an os.Getwd error in normal usage, so we won't do a full injection approach here.
	// We'll do a quick check: if normal usage is run, we don't get [ERR].
	d := &segments.DirSegment{}
	out, err := d.Render(nil)
	if out == "[ERR]" || err != nil {
		t.Errorf("Expected normal directory output, got %q (err=%v)", out, err)
	}
}

