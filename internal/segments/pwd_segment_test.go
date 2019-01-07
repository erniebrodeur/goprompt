package segments

import "testing"

func TestPwd_ColoredOutput(t *testing.T) {
	tests := []struct {
		name string
		p    Pwd
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ColoredOutput(); got != tt.want {
				t.Errorf("Pwd.ColoredOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPwd_Len(t *testing.T) {
	tests := []struct {
		name string
		p    Pwd
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("Pwd.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPwd_Output(t *testing.T) {
	tests := []struct {
		name string
		p    Pwd
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Output(); got != tt.want {
				t.Errorf("Pwd.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}
