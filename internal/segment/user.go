package segment

import (
	"os"
	"os/user"
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

func (u *UserSegment) Output() (SegmentOutput, error) {
    if u.u == nil {
        return SegmentOutput{}, nil
    }

    username := u.u.Username
    root := (os.Geteuid() == 0)
    ssh := false
    if os.Getenv("SSH_CONNECTION") != "" {
        ssh = true
        host, _ := os.Hostname()
        username = username + "@" + host
    }

    return SegmentOutput{
        Name:   "user",
        Text:   username,
        IsRoot: root,
        IsSSH:  ssh,
    }, nil
}
