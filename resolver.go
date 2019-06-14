package vresolver

// Resolver is the type used for all resolvers
type Resolver func(input string) string

// Compose allows to compose resolvers chained in a FILO fashion
func Compose(resolvers ...Resolver) Resolver {
	return func(input string) string {
		if len(resolvers) == 0 {
			return input
		}

		head := resolvers[:len(resolvers)-1]
		last := resolvers[len(resolvers)-1]

		if len(head) == 0 {
			return last(input)
		}

		return last(Compose(head...)(input))
	}
}
