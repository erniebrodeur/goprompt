package segment

import (
	"path/filepath"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/model"
)

// DirSegment returns a shortened path (e.g., ".../Projects/goprompt").
type DirSegment struct{}

func NewDirSegment() *DirSegment {
	return &DirSegment{}
}

func (d *DirSegment) Name() string {
	return "directory"
}

// Enabled always returns true for the directory segment; you can customize if needed.
func (d *DirSegment) Enabled(ctx *Context) bool {
	return true
}

// Output shortens the path from ctx.Pwd (rather than calling os.Getwd() again).
func (d *DirSegment) Output(ctx *Context) (model.SegmentOutput, error) {
	shortened := shortenPath(ctx.Pwd)
	return model.SegmentOutput{
		Name: "directory",
		Text: shortened,
	}, nil
}

// shortenPath collapses the path to ".../<lastTwo>" if there are 3+ parts,
// or leaves it as-is if there's only 1-2 parts.
func shortenPath(full string) string {
	parts := strings.Split(full, string(filepath.Separator))
	if len(parts) < 3 {
		return full
	}

	lastTwo := filepath.Join(parts[len(parts)-2], parts[len(parts)-1])
	return ".../" + lastTwo
}
