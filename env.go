package versionresolver

import "os"

// Env returns an environment variable resolver.
// Argument is the environment variable name
func Env(envvar string) string {
	return os.Getenv(envvar)
}
