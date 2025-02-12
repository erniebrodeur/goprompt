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
	// We'll do a quick partial check for each color
	s1 := theme.HexToANSI("#A6E22E")
	if !strings.Contains(s1, "166;226;46m") {
		t.Errorf("Expected #A6E22E => 166;226;46, got %q", s1)
	}
	s2 := theme.HexToANSI("#F92672")
	if !strings.Contains(s2, "249;38;114m") {
		t.Errorf("Expected #F92672 => 249;38;114, got %q", s2)
	}
}