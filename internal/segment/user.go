package segment

import (
	"os"
	"os/user"

	"github.com/erniebrodeur/goprompt/internal/model"
)

type UserSegment struct {
    u *user.User
}

func NewUserSegment() *UserSegment {
    usr, _ := user.Current()
    return &UserSegment{u: usr}
}

func (u *UserSegment) Name() string  { return "user" }
func (u *UserSegment) Enabled() bool { return true }

func (u *UserSegment) Output() (model.SegmentOutput, error) {
    if u.u == nil {
        return model.SegmentOutput{}, nil
    }

    username := u.u.Username
    isRoot := (os.Geteuid() == 0)
    isSSH := false
    if os.Getenv("SSH_CONNECTION") != "" {
        isSSH = true
        host, _ := os.Hostname()
        username = username + "@" + host
    }

    return model.SegmentOutput{
        Name:   "user",
        Text:   username,
        IsRoot: isRoot,
        IsSSH:  isSSH,
    }, nil
}
