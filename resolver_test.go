package versionresolver

import (
	"os"
	"testing"
)

func TestCompose(t *testing.T) {
	type args struct {
		resolvers []Resolver
	}
	tests := []struct {
		name  string
		input string
		args  args
		want  string
	}{
		{
			name:  "no resolver returns input value",
			input: "foo",
			args:  args{resolvers: []Resolver{}},
			want:  "foo",
		},
		{
			name:  "single resolver that forwards input value",
			input: "foo",
			args: args{resolvers: []Resolver{func(arg string) string {
				return arg
			}}},
			want: "foo",
		},
		{
			name:  "resolvers that mutate input value",
			input: "foo",
			args: args{resolvers: []Resolver{
				func(input string) string {
					return input
				},
				func(input string) string {
					return input + "bar"
				},
				func(input string) string {
					return input + "baz"
				},
			}},
			want: "foobarbaz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compose(tt.args.resolvers...)(tt.input); got != tt.want {
				t.Errorf("Compose() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("non-empty env resolver with fallback resolver", func(t *testing.T) {
		envvar := "foo"
		os.Setenv(envvar, "bar")
		defer os.Clearenv()

		c := Compose(Env, Fallback("test"))

		if got := c(envvar); got != "bar" {
			t.Errorf("Composer(Env, Fallback) = %v, want %v", got, "bar")
		}
	})

	t.Run("empty env resolver with fallback resolver", func(t *testing.T) {
		c := Compose(Env, Fallback("test"))

		if got := c("ghost"); got != "test" {
			t.Errorf("Composer(Env, Fallback) = %v, want %v", got, "test")
		}
	})

	t.Run("empty env, fallback to ecs no-tag metadata file, fallback to ecs metadata file", func(t *testing.T) {
		c := Compose(
			Env,
			Fallback("./fixtures/ecs-metadata-no-tag.json"),
			ECS,
			Fallback("./fixtures/ecs-metadata.json"),
			ECS,
		)

		if got := c("null"); got != "2.4" {
			t.Errorf("Composer(Env, Fallback, ECS, Fallback, ECS)= %v, want %v", got, "2.4")
		}
	})
}
