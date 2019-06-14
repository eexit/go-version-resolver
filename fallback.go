package versionresolver

// Fallback is a resolver that returns a given fallback
// value when input is empty
func Fallback(fallback string) Resolver {
	return Resolver(func(input string) string {
		if input == "" {
			return fallback
		}
		return input
	})
}
