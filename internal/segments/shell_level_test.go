package segments

import "testing"

func TestShellLevel_Output(t *testing.T) {
	tests := []struct {
		name string
		s    ShellLevel
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ShellLevel{}
			if got := s.Output(); got != tt.want {
				t.Errorf("ShellLevel.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShellLevel_Len(t *testing.T) {
	tests := []struct {
		name string
		s    ShellLevel
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ShellLevel{}
			if got := s.Len(); got != tt.want {
				t.Errorf("ShellLevel.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
