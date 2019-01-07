package segments

import "testing"

func TestGit_ColoredOutput(t *testing.T) {
	tests := []struct {
		name string
		g    Git
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.ColoredOutput(); got != tt.want {
				t.Errorf("Git.ColoredOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGit_Len(t *testing.T) {
	tests := []struct {
		name string
		g    Git
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Len(); got != tt.want {
				t.Errorf("Git.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGit_Output(t *testing.T) {
	tests := []struct {
		name string
		g    Git
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Output(); got != tt.want {
				t.Errorf("Git.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}
