package theme_test

import (
	"strings"
	"testing"

	"github.com/erniebrodeur/goprompt/internal/theme"
)

func TestHexToANSI(t *testing.T) {
	tests := []struct {
		hex  string
		want string
	}{
		{"#FD971F", "\033[38;2;253;151;31m"},
		{"#F92672", "\033[38;2;249;38;114m"},
		{"#FFFFFF", "\033[38;2;255;255;255m"},
		{"#000000", "\033[38;2;0;0;0m"},
	}

	for _, tt := range tests {
		got := theme.HexToANSI(tt.hex)
		if got != tt.want {
			t.Errorf("HexToANSI(%q) = %q, want %q", tt.hex, got, tt.want)
		}
	}
}

func TestParseHexColor(t *testing.T) {
	r, g, b := theme.HexToANSI("#A6E22E"), theme.HexToANSI("#F92672"), theme.HexToANSI("#FD971F")
	// We won't do a direct eq because these are ANSI strings. We'll do partial checks.
	if !strings.Contains(r, "38;2;166;226;46m") {
		t.Errorf("Expected #A6E22E => 166;226;46, got %q", r)
	}
	if !strings.Contains(g, "249;38;114m") {
		t.Errorf("Expected #F92672 => 249;38;114, got %q", g)
	}
	if !strings.Contains(b, "253;151;31m") {
		t.Errorf("Expected #FD971F => 253;151;31, got %q", b)
	}
}
