# assertgo : Conditionally compiled assertions in Go

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/arl/assertgo) 

## Example

**`main.go`**:

```go
package main

import (
	"fmt"

	"github.com/arl/assertgo"
)

func main() {
	assert.True(true, "never printed")
	assert.Truef(false, "panic and printed if -debug tag on")
	fmt.Println("program end")
}
```


### Normal run

    $ go run main.go
    program end

Nothing happened because `assert.True()` is conditionnaly compiled to a noop.


### Debug run

    $ go run -tags debug main.go
    2016/11/02 22:43:59 --- --- Debug Assertion Failed --- --- ---
    panic: sent false to assert.True
    
    goroutine 1 [running]:
    panic(0x496be0, 0xc420062200)
    	/usr/local/go/src/runtime/panic.go:500 +0x1a1
    github.com/arl/assertgo.True(0xc42003ff00, 0x4b5298, 0x19, 0x0, 0x0, 0x0)
    	/home/panty/godev/src/github.com/arl/assertgo/assert.go:19 +0x13b
    main.main()
    	/home/panty/godev/src/github.com/arl/tagtest/main.go:10 +0x5a
    exit status 2

By providing the `debug` tag to `go build` or any other `go` tool that accepts
[build tags](https://golang.org/pkg/go/build/), `assert.True` is conditionnaly
compiled to a function that calls `panic` if the assertion is `false`.


### Benchmarks

With debug flag:

    $ go test -bench '.*' -run '^$' -tags debug
    goos: linux
    goarch: amd64
    pkg: github.com/arl/assertgo
    BenchmarkTrue-4         30000000                40.5 ns/op
    BenchmarkFalse-4        30000000                38.9 ns/op
    BenchmarkTruef-4        1000000000               2.90 ns/op
    BenchmarkFalsef-4       300000000                5.91 ns/op
    PASS
    ok      github.com/arl/assertgo    8.039s

Without debug flag, noop (or close):

    $ go test -bench '.*' -run '^$'
    goos: linux
    goarch: amd64
    pkg: github.com/arl/assertgo
    BenchmarkTrue-4         2000000000               1.82 ns/op
    BenchmarkFalse-4        2000000000               1.82 ns/op
    BenchmarkTruef-4        1000000000               2.53 ns/op
    BenchmarkFalsef-4       1000000000               2.55 ns/op
    PASS
    ok      github.com/arl/assertgo    13.260s

