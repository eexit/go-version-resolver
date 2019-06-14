package versionresolver

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	tests := []struct {
		name   string
		envvar string
		envval string
		want   string
	}{
		{
			name:   "non-existing envvar returns empty version",
			envvar: "",
			envval: "",
			want:   "",
		},
		{
			name:   "empty envvar value returns empty version",
			envvar: "foo",
			envval: "",
			want:   "",
		},
		{
			name:   "envvar value is returned as version",
			envvar: "foo",
			envval: "1.0.0",
			want:   "1.0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.envvar, tt.envval)
			defer os.Clearenv()

			if got := Env(tt.envvar); got != tt.want {
				t.Errorf("Env() = %v, want %v", got, tt.want)
			}
		})
	}
}
