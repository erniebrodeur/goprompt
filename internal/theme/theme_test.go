package theme

import (
	"strings"
	"testing"
)

func TestDefaultThemePreservesLegacyANSI(t *testing.T) {
	selected, err := Lookup("default")
	if err != nil {
		t.Fatalf("Lookup(default) error = %v", err)
	}

	tests := []struct {
		name  string
		style Style
		want  string
	}{
		{name: "frame", style: selected.Frame, want: "\x1b[0;94mframe\x1b[0m"},
		{name: "primary", style: selected.Primary, want: "\x1b[0;92mprimary\x1b[0m"},
		{name: "vcs", style: selected.VCS, want: "\x1b[0;93mvcs\x1b[0m"},
		{name: "clock", style: selected.Clock, want: "\x1b[0;96mclock\x1b[0m"},
		{name: "prompt", style: selected.Prompt, want: "prompt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.style.Apply(tt.name)
			if got != tt.want {
				t.Fatalf("Apply(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestMonokaiUsesTrueColorForegrounds(t *testing.T) {
	selected, err := Lookup("monokai")
	if err != nil {
		t.Fatalf("Lookup(monokai) error = %v", err)
	}

	tests := []struct {
		name  string
		style Style
		want  string
	}{
		{name: "frame", style: selected.Frame, want: "\x1b[38;2;174;129;255mframe\x1b[0m"},
		{name: "primary", style: selected.Primary, want: "\x1b[38;2;166;226;46mprimary\x1b[0m"},
		{name: "vcs", style: selected.VCS, want: "\x1b[38;2;230;219;116mvcs\x1b[0m"},
		{name: "clock", style: selected.Clock, want: "\x1b[38;2;102;217;239mclock\x1b[0m"},
		{name: "prompt", style: selected.Prompt, want: "\x1b[38;2;248;248;242mprompt\x1b[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.style.Apply(tt.name)
			if got != tt.want {
				t.Fatalf("Apply(%q) = %q, want %q", tt.name, got, tt.want)
			}
			if strings.Contains(got, "\x1b[48;2;") {
				t.Fatalf("Apply(%q) set a background color: %q", tt.name, got)
			}
		})
	}
}

func TestStyleLeavesEmptyValuesUnchanged(t *testing.T) {
	for _, themeName := range []string{"default", "monokai"} {
		selected, err := Lookup(themeName)
		if err != nil {
			t.Fatalf("Lookup(%q) error = %v", themeName, err)
		}

		for _, style := range []Style{selected.Frame, selected.Primary, selected.VCS, selected.Clock, selected.Prompt} {
			if got := style.Apply(""); got != "" {
				t.Fatalf("%s style applied escapes to an empty value: %q", themeName, got)
			}
		}
	}
}

func TestLookupNormalizesThemeNames(t *testing.T) {
	selected, err := Lookup("  MONOKAI  ")
	if err != nil {
		t.Fatalf("Lookup normalized name error = %v", err)
	}
	if selected.Name != "monokai" {
		t.Fatalf("Lookup normalized name = %q, want monokai", selected.Name)
	}
}

func TestResolveThemePrecedence(t *testing.T) {
	tests := []struct {
		name    string
		cliName string
		envName string
		want    string
	}{
		{name: "default", want: "default"},
		{name: "environment", envName: "monokai", want: "monokai"},
		{name: "cli", cliName: "default", envName: "monokai", want: "default"},
		{name: "valid cli overrides invalid environment", cliName: "monokai", envName: "unknown", want: "monokai"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selected, err := Resolve(tt.cliName, tt.envName)
			if err != nil {
				t.Fatalf("Resolve() error = %v", err)
			}
			if selected.Name != tt.want {
				t.Fatalf("Resolve() = %q, want %q", selected.Name, tt.want)
			}
		})
	}
}

func TestResolveRejectsUnknownTheme(t *testing.T) {
	_, err := Resolve("unknown", "monokai")
	if err == nil {
		t.Fatal("Resolve() error = nil, want unknown theme error")
	}

	want := `unknown theme "unknown"; valid themes: default, monokai`
	if err.Error() != want {
		t.Fatalf("Resolve() error = %q, want %q", err, want)
	}
}
