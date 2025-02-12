package segments_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/erniebrodeur/goprompt/internal/segments"
	"github.com/stretchr/testify/require"
)

func TestDirSegment_Normal(t *testing.T) {
	origWD, _ := os.Getwd()
	defer os.Chdir(origWD)

	tmp := t.TempDir()
	require.NoError(t, os.Chdir(tmp))

	d := &segments.DirSegment{ShowComponents: 1}
	out, err := d.Render(map[string]string{"dir.normal": "#FD971F"})
	require.NoError(t, err)
	require.NotEmpty(t, out)

	want := filepath.Base(tmp)
	require.Contains(t, out, want, "Expected last path component in dir output")
}

func TestDirSegment_Error(t *testing.T) {
	d := &segments.DirSegment{}
	out, err := d.Render(nil)
	require.NoError(t, err)
	require.NotEqual(t, "[ERR]", out, "Should not produce [ERR] in a normal scenario")
}

---
