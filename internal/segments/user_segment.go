package segments

type UserSegment struct{}

func (u *UserSegment) Render(theme map[string]string) (string, error) {
// Minimal placeholder
// Possibly check if root user => "root", else normal user
return "USER_NOT_IMPL", nil
}

