package versionresolver

import "testing"

func TestFallback(t *testing.T) {
	tests := []struct {
		name     string
		fallback string
		input    string
		want     string
	}{
		{
			name:     "empty fallback and input return empty version",
			fallback: "",
			input:    "",
			want:     "",
		},
		{
			name:     "empty fallback and non-empty input pass input as version",
			fallback: "",
			input:    "foo",
			want:     "foo",
		},
		{
			name:     "non-empty fallback and non-empty input return input as version",
			fallback: "bar",
			input:    "foo",
			want:     "foo",
		},
		{
			name:     "non-empty fallback and empty input return fallback as version",
			fallback: "bar",
			input:    "",
			want:     "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fallback(tt.fallback)(tt.input); got != tt.want {
				t.Errorf("Fallback() = %v, want %v", got, tt.want)
			}
		})
	}
}
