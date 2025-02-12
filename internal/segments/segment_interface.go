package segments

type Segment interface {
	Render(theme map[string]string) (string, error)
}

---
