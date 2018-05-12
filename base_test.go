package assert

import (
	"fmt"
	"testing"
)

func hasPanicked(f func()) (bool, interface{}) {

	didPanic := false
	var msg interface{}
	func() {

		defer func() {
			if msg = recover(); msg != nil {
				didPanic = true
			}
		}()

		// call the target function
		f()

	}()

	return didPanic, msg
}

type assertFunc func(bool, string, ...interface{})

func testAssert(t *testing.T, f assertFunc, correct bool) {
	tests := []struct {
		correct     bool
		format      string
		args        []interface{}
		shouldPanic bool
	}{
		// correct assertion: should never panic
		{correct: correct, format: "%v %v %v", args: []interface{}{1, "val", 3.14}, shouldPanic: false},
		{correct: correct, format: "wrong num of placeholders ", args: []interface{}{1, "val", 3.14}, shouldPanic: false},
		{correct: correct, format: "format", args: nil, shouldPanic: false},
		{correct: correct, format: "", args: nil, shouldPanic: false},
		{correct: correct, format: "", args: []interface{}{1, "val", 3.14}, shouldPanic: false},

		// uncorrect assertion: should panic if debug flag is set
		{correct: !correct, format: "%v %v %v", args: []interface{}{1, "val", 3.14}, shouldPanic: isDebug},
		{correct: !correct, format: "wrong num of placeholders ", args: []interface{}{1, "val", 3.14}, shouldPanic: isDebug},
		{correct: !correct, format: "format", args: nil, shouldPanic: isDebug},
		{correct: !correct, format: "", args: nil, shouldPanic: isDebug},
		{correct: !correct, format: "", args: []interface{}{1, "val", 3.14}, shouldPanic: isDebug},
	}

	fmt.Println("debug flag is set?", isDebug)

	for _, tt := range tests {
		got, msg := hasPanicked(func() {
			f(tt.correct, tt.format, tt.args...)
		})
		if got != tt.shouldPanic {
			if tt.shouldPanic {
				t.Error("assert.True should have panicked but didn't")
			} else {
				t.Errorf("assert.True should not have panicked but did, with:\n%v", msg)
			}
		}
	}
}

func TestTrue(t *testing.T)  { testAssert(t, True, true) }
func TestFalse(t *testing.T) { testAssert(t, False, false) }

/*
func TestFalse(t *testing.T) {
	tests := []struct {
		cond   bool
		format string
		args   []interface{}
	}{}
	for _, tt := range tests {
		True(tt.cond, tt.format, tt.args...)
	}
}
*/
