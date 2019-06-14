package vresolver

import "testing"

func TestECS(t *testing.T) {
	tests := []struct {
		name string
		file string
		want string
	}{
		{
			name: "empty file name file",
			file: "",
			want: "",
		},
		{
			name: "non-existing file",
			file: "ghost.json",
			want: "",
		},
		{
			name: "non-JSON file",
			file: "./fixtures/ecs-non-json.html",
			want: "",
		},
		{
			name: "valid metadata file",
			file: "./fixtures/ecs-metadata.json",
			want: "2.4",
		},
		{
			name: "valid metadata file but no tag",
			file: "./fixtures/ecs-metadata-no-tag.json",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ECS(tt.file); got != tt.want {
				t.Errorf("ECS() = %v, want %v", got, tt.want)
			}
		})
	}
}
