package main

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/erniebrodeur/goprompt/internal/theme"
)

func TestParseOptions(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantVersion bool
		wantStatus  bool
		wantTheme   string
	}{
		{name: "short version", args: []string{"-v"}, wantVersion: true},
		{name: "long version", args: []string{"--version"}, wantVersion: true},
		{name: "short status", args: []string{"-s"}, wantStatus: true},
		{name: "long status", args: []string{"--status"}, wantStatus: true},
		{name: "short theme", args: []string{"-t", "monokai"}, wantTheme: "monokai"},
		{name: "long theme", args: []string{"--theme", "monokai"}, wantTheme: "monokai"},
		{name: "long theme with equals", args: []string{"--theme=monokai"}, wantTheme: "monokai"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOptions(tt.args, io.Discard)
			if err != nil {
				t.Fatalf("parseOptions() error = %v", err)
			}
			if got.showVersion != tt.wantVersion {
				t.Errorf("parseOptions() version = %v, want %v", got.showVersion, tt.wantVersion)
			}
			if got.showStatus != tt.wantStatus {
				t.Errorf("parseOptions() status = %v, want %v", got.showStatus, tt.wantStatus)
			}
			if got.themeName != tt.wantTheme {
				t.Errorf("parseOptions() theme = %q, want %q", got.themeName, tt.wantTheme)
			}
		})
	}
}

func TestRenderPromptPreservesDefaultOutput(t *testing.T) {
	selected, err := theme.Lookup("default")
	if err != nil {
		t.Fatalf("theme.Lookup(default) error = %v", err)
	}

	parts := promptParts{
		pwd: "/tmp", git: ":main*", login: "dev", host: "@host", mid: "──",
		clock: "09:54am", date: "8/13", shell: "%%",
	}
	var output bytes.Buffer
	renderPrompt(&output, selected, parts)

	blue := func(value string) string { return "\x1b[0;94m" + value + "\x1b[0m" }
	green := func(value string) string { return "\x1b[0;92m" + value + "\x1b[0m" }
	yellow := func(value string) string { return "\x1b[0;93m" + value + "\x1b[0m" }
	cyan := func(value string) string { return "\x1b[0;96m" + value + "\x1b[0m" }
	want := blue("─┤ ") + green("/tmp") + yellow(":main*") + blue(" ├─") +
		blue("─┤ ") + green("dev") + yellow("@host") + blue(" ├─") + blue("──") +
		blue("─┤ ") + cyan("09:54am") + blue(" ─ ") + cyan("8/13") + blue(" ├─") + "\n%% "

	if got := output.String(); got != want {
		t.Fatalf("renderPrompt(default) = %q, want %q", got, want)
	}
}

func TestRenderPromptUsesMonokaiTrueColor(t *testing.T) {
	selected, err := theme.Lookup("monokai")
	if err != nil {
		t.Fatalf("theme.Lookup(monokai) error = %v", err)
	}

	parts := promptParts{
		pwd: "/tmp", git: ":main", login: "dev", host: "", mid: "─",
		clock: "09:54am", date: "8/13", shell: "%%",
	}
	var output bytes.Buffer
	renderPrompt(&output, selected, parts)

	for _, sequence := range []string{
		"\x1b[38;2;174;129;255m",
		"\x1b[38;2;166;226;46m",
		"\x1b[38;2;230;219;116m",
		"\x1b[38;2;102;217;239m",
		"\x1b[38;2;248;248;242m",
	} {
		if !strings.Contains(output.String(), sequence) {
			t.Errorf("renderPrompt(monokai) missing %q in %q", sequence, output.String())
		}
	}
	if strings.Contains(output.String(), "\x1b[48;2;") {
		t.Fatalf("renderPrompt(monokai) set a background color: %q", output.String())
	}
}

func TestRenderStatusUsesSelectedTheme(t *testing.T) {
	selected, err := theme.Lookup("monokai")
	if err != nil {
		t.Fatalf("theme.Lookup(monokai) error = %v", err)
	}

	parts := promptParts{
		git: ":main", login: "dev", pwd: "/tmp", clock: "09:54am", date: "8/13", shell: "%%",
	}
	var output bytes.Buffer
	renderStatus(&output, selected, parts)

	if !strings.Contains(output.String(), "git:           \x1b[38;2;230;219;116m:main\x1b[0m\n") {
		t.Fatalf("renderStatus(monokai) did not theme git output: %q", output.String())
	}
	if !strings.Contains(output.String(), "host:          \n") {
		t.Fatalf("renderStatus(monokai) added escapes to empty host output: %q", output.String())
	}
	if !strings.Contains(output.String(), "shell:         \x1b[38;2;248;248;242m%%\x1b[0m\n") {
		t.Fatalf("renderStatus(monokai) did not theme shell output: %q", output.String())
	}
}

func TestRenderStatusPreservesDefaultOutput(t *testing.T) {
	selected, err := theme.Lookup("default")
	if err != nil {
		t.Fatalf("theme.Lookup(default) error = %v", err)
	}

	parts := promptParts{
		git: ":main", login: "dev", pwd: "/tmp", clock: "09:54am", date: "8/13", shell: "%%",
	}
	var output bytes.Buffer
	renderStatus(&output, selected, parts)

	yellow := func(value string) string { return "\x1b[0;93m" + value + "\x1b[0m" }
	green := func(value string) string { return "\x1b[0;92m" + value + "\x1b[0m" }
	blue := func(value string) string { return "\x1b[0;94m" + value + "\x1b[0m" }
	cyan := func(value string) string { return "\x1b[0;96m" + value + "\x1b[0m" }
	want := "git:           " + yellow(":main") + "\n" +
		"host:          \n" +
		"login:         " + green("dev") + "\n" +
		"mid:           \n" +
		"pwd:           " + green("/tmp") + "\n" +
		"time:          " + cyan("09:54am") + blue(" ─ ") + cyan("8/13") + "\n" +
		"shell:         %%\n"

	if got := output.String(); got != want {
		t.Fatalf("renderStatus(default) = %q, want %q", got, want)
	}
}

func TestRunRejectsUnknownTheme(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		environmentTheme string
	}{
		{name: "environment", environmentTheme: "unknown"},
		{name: "cli", args: []string{"--theme", "unknown"}, environmentTheme: "monokai"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			exitCode := run(tt.args, tt.environmentTheme, &stdout, &stderr)

			if exitCode != 2 {
				t.Fatalf("run() exit code = %d, want 2", exitCode)
			}
			if stdout.Len() != 0 {
				t.Fatalf("run() stdout = %q, want empty", stdout.String())
			}
			want := "goprompt: unknown theme \"unknown\"; valid themes: default, monokai\n"
			if stderr.String() != want {
				t.Fatalf("run() stderr = %q, want %q", stderr.String(), want)
			}
		})
	}
}

func TestRunAppliesThemePrecedence(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		environmentTheme string
		wantSequence     string
		rejectSequence   string
	}{
		{
			name:             "environment selects monokai",
			args:             []string{"--status"},
			environmentTheme: "monokai",
			wantSequence:     "\x1b[38;2;",
			rejectSequence:   "\x1b[0;",
		},
		{
			name:             "cli overrides environment",
			args:             []string{"--status", "--theme", "default"},
			environmentTheme: "monokai",
			wantSequence:     "\x1b[0;",
			rejectSequence:   "\x1b[38;2;",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			exitCode := run(tt.args, tt.environmentTheme, &stdout, &stderr)

			if exitCode != 0 {
				t.Fatalf("run() exit code = %d, want 0; stderr = %q", exitCode, stderr.String())
			}
			if !strings.Contains(stdout.String(), tt.wantSequence) {
				t.Fatalf("run() stdout missing %q: %q", tt.wantSequence, stdout.String())
			}
			if strings.Contains(stdout.String(), tt.rejectSequence) {
				t.Fatalf("run() stdout contains %q: %q", tt.rejectSequence, stdout.String())
			}
			if stderr.Len() != 0 {
				t.Fatalf("run() stderr = %q, want empty", stderr.String())
			}
		})
	}
}

func TestRunAppliesThemeToNormalOutput(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		environmentTheme string
	}{
		{name: "cli", args: []string{"--theme", "monokai"}},
		{name: "environment", environmentTheme: "monokai"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			exitCode := run(tt.args, tt.environmentTheme, &stdout, &stderr)

			if exitCode != 0 {
				t.Fatalf("run() exit code = %d, want 0; stderr = %q", exitCode, stderr.String())
			}
			if !strings.Contains(stdout.String(), "\x1b[38;2;174;129;255m─┤ \x1b[0m") {
				t.Fatalf("run() normal output did not use Monokai frame color: %q", stdout.String())
			}
			if !strings.Contains(stdout.String(), "\n\x1b[38;2;248;248;242m%%\x1b[0m ") {
				t.Fatalf("run() normal output did not use Monokai prompt color: %q", stdout.String())
			}
			if stderr.Len() != 0 {
				t.Fatalf("run() stderr = %q, want empty", stderr.String())
			}
		})
	}
}

func TestRunPrintsFlagUsage(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"--help"}, "", &stdout, &stderr)

	if exitCode != 2 {
		t.Fatalf("run() exit code = %d, want 2", exitCode)
	}
	if stdout.Len() != 0 {
		t.Fatalf("run() stdout = %q, want empty", stdout.String())
	}
	for _, value := range []string{"Usage of goprompt:", "-theme string", "Select theme"} {
		if !strings.Contains(stderr.String(), value) {
			t.Errorf("run() help output missing %q: %q", value, stderr.String())
		}
	}
}

func TestRunValidatesThemeBeforeVersion(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		environmentTheme string
	}{
		{name: "cli", args: []string{"--version", "--theme", "unknown"}},
		{name: "environment", args: []string{"--version"}, environmentTheme: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			exitCode := run(tt.args, tt.environmentTheme, &stdout, &stderr)

			if exitCode != 2 {
				t.Fatalf("run() exit code = %d, want 2", exitCode)
			}
			if stdout.Len() != 0 {
				t.Fatalf("run() stdout = %q, want empty", stdout.String())
			}
			want := "goprompt: unknown theme \"unknown\"; valid themes: default, monokai\n"
			if stderr.String() != want {
				t.Fatalf("run() stderr = %q, want %q", stderr.String(), want)
			}
		})
	}
}
