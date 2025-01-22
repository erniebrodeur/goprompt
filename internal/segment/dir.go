package segment

import (
	"os"
	"os/user"
	"strings"
)

type DirSegment struct{}

func NewDirSegment() *DirSegment {
    return &DirSegment{}
}

func (d *DirSegment) Name() string   { return "directory" }
func (d *DirSegment) Enabled() bool  { return true }

func (d *DirSegment) Output() (SegmentOutput, error) {
    cwd, err := os.Getwd()
    if err != nil {
        return SegmentOutput{}, err
    }

    usr, _ := user.Current()
    if usr != nil {
        if strings.HasPrefix(cwd, usr.HomeDir) {
            cwd = "~" + strings.TrimPrefix(cwd, usr.HomeDir)
        }
    }

    return SegmentOutput{
        Name: "directory",
        Text: cwd,
    }, nil
}
