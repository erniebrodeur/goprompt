package main

import "testing"

func TestParseOptions(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantVersion bool
		wantStatus  bool
	}{
		{name: "short version", args: []string{"-v"}, wantVersion: true},
		{name: "long version", args: []string{"--version"}, wantVersion: true},
		{name: "short status", args: []string{"-s"}, wantStatus: true},
		{name: "long status", args: []string{"--status"}, wantStatus: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVersion, gotStatus, err := parseOptions(tt.args)
			if err != nil {
				t.Fatalf("parseOptions() error = %v", err)
			}
			if gotVersion != tt.wantVersion {
				t.Errorf("parseOptions() version = %v, want %v", gotVersion, tt.wantVersion)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("parseOptions() status = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
