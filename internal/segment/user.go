package segment

import (
	"fmt"
	"os"
	"os/user"

	"github.com/erniebrodeur/goprompt/internal/colors"
)

type UserSegment struct {
    currentUser *user.User
}

func NewUserSegment() *UserSegment {
    u, _ := user.Current()
    return &UserSegment{currentUser: u}
}

func (u *UserSegment) Name() string { return "user" }

func (u *UserSegment) Enabled() bool { return true }

func (u *UserSegment) Render() (string, error) {
    if u.currentUser == nil {
        // If something weird happens, skip
        return "", nil
    }

    uid := os.Geteuid()
    var color string
    if uid == 0 {
        color = colors.Red
    } else {
        color = colors.Green
    }

    userStr := u.currentUser.Username

    // Check SSH env
    if os.Getenv("SSH_CONNECTION") != "" {
        hostname, _ := os.Hostname()
        return fmt.Sprintf("%s%s@%s%s ", color, userStr, hostname, colors.Reset), nil
    }

    // local user
    return fmt.Sprintf("%s%s%s ", color, userStr, colors.Reset), nil
}
