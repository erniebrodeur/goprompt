package segments

import "testing"

func TestHost_ColoredOutput(t *testing.T) {
	tests := []struct {
		name string
		h    Host
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.ColoredOutput(); got != tt.want {
				t.Errorf("Host.ColoredOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHost_Len(t *testing.T) {
	tests := []struct {
		name string
		h    Host
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("Host.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHost_Output(t *testing.T) {
	tests := []struct {
		name string
		h    Host
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Output(); got != tt.want {
				t.Errorf("Host.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}
