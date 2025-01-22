package segment

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/colors"
)

type DirSegment struct{}

func NewDirSegment() *DirSegment {
    return &DirSegment{}
}

func (d *DirSegment) Name() string { return "directory" }

func (d *DirSegment) Enabled() bool { return true }

func (d *DirSegment) Render() (string, error) {
    path, err := os.Getwd()
    if err != nil {
        return "", err
    }

    home, err := user.Current()
    if err == nil && home != nil {
        homeDir := home.HomeDir
        if strings.HasPrefix(path, homeDir) {
            path = "~" + strings.TrimPrefix(path, homeDir)
        }
    }

    // Just coloring the directory for now
    return fmt.Sprintf("%s%s%s ", colors.Blue, path, colors.Reset), nil
}
