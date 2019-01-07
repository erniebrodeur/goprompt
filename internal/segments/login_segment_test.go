package segments

import "testing"

func TestLogin_ColoredOutput(t *testing.T) {
	tests := []struct {
		name string
		l    Login
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.ColoredOutput(); got != tt.want {
				t.Errorf("Login.ColoredOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogin_Len(t *testing.T) {
	tests := []struct {
		name string
		l    Login
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); got != tt.want {
				t.Errorf("Login.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogin_Output(t *testing.T) {
	tests := []struct {
		name string
		l    Login
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Output(); got != tt.want {
				t.Errorf("Login.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}
