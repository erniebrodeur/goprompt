package theme_test

import (
	"strings"
	"testing"

	"github.com/erniebrodeur/goprompt/internal/theme"
	"github.com/stretchr/testify/require"
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
		tt := tt
		t.Run(tt.hex, func(t *testing.T) {
			got := theme.HexToANSI(tt.hex)
			require.Equal(t, tt.want, got, "Mismatch for hex %s", tt.hex)
		})
	}
}

func TestParseHexColor(t *testing.T) {
	s1 := theme.HexToANSI("#A6E22E")
	require.Contains(t, s1, "166;226;46m", "Expected #A6E22E => 166;226;46")

	s2 := theme.HexToANSI("#F92672")
	require.Contains(t, s2, "249;38;114m", "Expected #F92672 => 249;38;114")
}