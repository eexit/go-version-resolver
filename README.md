# V(ersion) Resolver [![Build Status](https://travis-ci.org/eexit/vresolver.svg?branch=master)](https://travis-ci.org/eexit/vresolver) [![codecov](https://codecov.io/gh/eexit/vresolver/branch/master/graph/badge.svg)](https://codecov.io/gh/eexit/vresolver)

Simple and flexible way to fetch your app version.

Wether your app is a CLI or runs in background and generate logs, it's never been as easy to fetch your app version with this lib.

Supported resolvers:

- Environment variable
- [AWS ECS Metadata File](https://blog.eexit.net/aws-ecs-seamlessly-handle-app-versioning/)
- Fallback
- Composite

### Example

```go
import "github.com/eexit/vresolver"

// Tries to fetch the version from the runtime ECS container
// and fallbacks to "bulk-version" otherwise
version := vresolver.Compose(
	vresolver r.Env,
	vresolver.ECS,
	vresolver.Fallback("bulk-version"),
)("ECS_CONTAINER_METADATA_FILE")

```

## Installation

```bash
$ go get -u github.com/eexit/vresolver
```

## Test

```bash
$ go test -v -race ./...
```

## Custom resolver

You can build a custom resolver as long as it matches the following type:

```go
type Resolver func(input string) string
```

For instance, let's create a `Panic` resolver, that panics if no version is resolved:

```go
// Panic panics when app has no version
func Panic(input string) string {
	if input == "" {
		panic("no app version was resolved")
	}
	return input
}

version := vresolver.Compose(vresolver.Env, Panic)("MY_APP_VERSION")
```


