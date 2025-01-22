package segment

import (
	"os"
	"os/user"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/model"
)

// UserSegment shows the current username, and marks root differently if needed.
type UserSegment struct{}

func NewUserSegment() *UserSegment {
	return &UserSegment{}
}

func (u *UserSegment) Name() string {
	return "user"
}

// Enabled always returns true here. If you only want to show the user segment
// in certain contexts, you could check ctx.IsSSH, etc.
func (u *UserSegment) Enabled(ctx *Context) bool {
	return true
}

// Output returns the username; sets IsRoot if necessary.
func (u *UserSegment) Output(ctx *Context) (model.SegmentOutput, error) {
	username := "unknown"
	if usr, err := user.Current(); err == nil {
		// Some shells pass "USER", fallback to that if needed
		username = usr.Username
	} else if envUser := os.Getenv("USER"); envUser != "" {
		username = envUser
	}

	// Clean up domain parts if on Windows or domain\user format
	if idx := strings.IndexRune(username, '\\'); idx != -1 {
		username = username[idx+1:]
	}

	return model.SegmentOutput{
		Name:   "user",
		Text:   username,
		IsRoot: ctx.IsRoot,
	}, nil
}
