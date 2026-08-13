package theme

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mgutz/ansi"
)

const (
	DefaultName = "default"
	reset       = "\x1b[0m"
)

// Style applies one foreground color without changing the terminal background.
type Style struct {
	start string
}

// Apply wraps a non-empty value in this style's foreground color.
func (s Style) Apply(value string) string {
	if value == "" || s.start == "" {
		return value
	}

	return s.start + value + reset
}

// Theme maps prompt content to semantic color roles.
type Theme struct {
	Name    string
	Frame   Style
	Primary Style
	VCS     Style
	Clock   Style
	Prompt  Style
}

var themes = map[string]Theme{
	DefaultName: {
		Name:    DefaultName,
		Frame:   legacyStyle("blue+h"),
		Primary: legacyStyle("green+h"),
		VCS:     legacyStyle("yellow+h"),
		Clock:   legacyStyle("cyan+h"),
	},
	"monokai": {
		Name:    "monokai",
		Frame:   trueColor(0xAE, 0x81, 0xFF),
		Primary: trueColor(0xA6, 0xE2, 0x2E),
		VCS:     trueColor(0xE6, 0xDB, 0x74),
		Clock:   trueColor(0x66, 0xD9, 0xEF),
		Prompt:  trueColor(0xF8, 0xF8, 0xF2),
	},
}

// Lookup returns a theme by its stable, case-insensitive name.
func Lookup(name string) (Theme, error) {
	normalized := strings.ToLower(strings.TrimSpace(name))
	selected, ok := themes[normalized]
	if !ok {
		return Theme{}, fmt.Errorf(
			"unknown theme %q; valid themes: %s",
			normalized,
			strings.Join(Names(), ", "),
		)
	}

	return selected, nil
}

// Resolve selects the CLI theme, then the environment theme, then the default.
func Resolve(cliName, envName string) (Theme, error) {
	name := strings.TrimSpace(cliName)
	if name == "" {
		name = strings.TrimSpace(envName)
	}
	if name == "" {
		name = DefaultName
	}

	return Lookup(name)
}

// Names returns all stable theme identifiers in sorted order.
func Names() []string {
	names := make([]string, 0, len(themes))
	for name := range themes {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func legacyStyle(name string) Style {
	return Style{start: ansi.ColorCode(name)}
}

func trueColor(red, green, blue uint8) Style {
	return Style{start: fmt.Sprintf("\x1b[38;2;%d;%d;%dm", red, green, blue)}
}
