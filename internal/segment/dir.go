package segment

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/model"
)

// DirSegment is responsible for returning a shortened current directory path.
type DirSegment struct{}

func NewDirSegment() *DirSegment {
	return &DirSegment{}
}

func (d *DirSegment) Name() string { 
	return "directory"
}

// Always enabled, since we can always get a directory (unless there's an error).
func (d *DirSegment) Enabled() bool {
	return true
}

// Output returns a short path. Example: 
//   /Users/you/Projects/goprompt -> ".../Projects/goprompt"
//   /home/username -> "~"
//   If there's only one or two segments, "..." won't be prepended.
func (d *DirSegment) Output() (model.SegmentOutput, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return model.SegmentOutput{}, err
	}

	shortened := shortenPath(cwd)

	return model.SegmentOutput{
		Name: "directory",
		Text: shortened,
	}, nil
}

// shortenPath converts an absolute path to a more compact version:
//   "/Users/you/Projects/goprompt" -> "~" or "~/<dir>" or ".../<last2segments>"
func shortenPath(full string) string {
	// Replace home dir with '~' if possible.
	usr, _ := user.Current()
	home := ""
	if usr != nil {
		home = usr.HomeDir
	}
	if home != "" && strings.HasPrefix(full, home) {
		full = "~" + strings.TrimPrefix(full, home)
	}

	// Split on the file separator.
	parts := strings.Split(full, string(os.PathSeparator))
	// If we have fewer than 3 parts (e.g., "~" or "~/<dir>"), just return as-is.
	if len(parts) < 3 {
		return full
	}

	// Rejoin the last two parts, prepend ".../" to indicate it's shortened.
	lastTwo := filepath.Join(parts[len(parts)-2], parts[len(parts)-1])
	return ".../" + lastTwo
}
