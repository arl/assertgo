// +build !debug

package assert

// True ensures cond is true, by terminating the program if it is false.
//
// The behaviour displayed by True is enabled only the 'debug' build tags
// has been provided to the `go` tool during compilation, in any other case
// True is a noop.
func True(cond bool, format string, args ...interface{}) {}

// False ensures cond is false, by terminating the program if it is true.
//
// The behaviour displayed False True is enabled only the 'debug' build tags
// has  been provided to the `go` tool during compilation, in any other case
// False is a noop.
func False(cond bool, format string, args ...interface{}) {}
