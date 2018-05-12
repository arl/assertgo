// +build debug

package assert

import (
	"fmt"
)

// True panics if cond is false. Truef formats the panic message using the
// default formats for its operands.
//
// The behaviour displayed by True is enabled only if the 'debug' build tag has
// been provided during compilation, otherwise True is a noop.
func True(cond bool, a ...interface{}) {
	Truef(cond, fmt.Sprint(a...))
}

// False panics if cond is true. False formats the panic message using the
// default formats for its operands.
//
// The behaviour displayed by False is enabled only if the 'debug' build tag has
// been provided during compilation, otherwise False is a noop.
func False(cond bool, a ...interface{}) {
	Truef(!cond, fmt.Sprint(a...))
}

// Truef panics if cond is false. Truef formats the panic message according to a
// format specifier.
//
// The behaviour displayed by Truef is enabled only if the 'debug' build tag has
// been provided during compilation, otherwise Truef is a noop.
func Truef(cond bool, format string, a ...interface{}) {
	if !cond {
		fmt.Println("--- --- Debug Assertion Failed --- --- ---")
		if a == nil || len(a) == 0 {
			panic(format)
		} else {
			panic(fmt.Sprintf(format, a...))
		}
	}
}

// Falsef panics if cond is true. Falsef formats the panic message according to
// a format specifier.
//
// The behaviour displayed by Falsef is enabled only if the 'debug' build tag has
// been provided during compilation, otherwise Falsef is a noop.
func Falsef(cond bool, format string, a ...interface{}) {
	Truef(!cond, format, a...)
}
