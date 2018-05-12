package assert

import (
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

func testAssert(t *testing.T, f func(bool, string, ...interface{})) {

	tests := []struct {
		format string
		args   []interface{}
	}{
		{format: "%v %v %v", args: []interface{}{1, "val", 3.14}},
		{format: "wrong num of placeholders ", args: []interface{}{1, "val", 3.14}},
		{format: "format", args: nil},
		{format: "", args: nil},
		{format: "", args: []interface{}{1, "val", 3.14}},
	}

	for _, tt := range tests {

		// run tests twice:
		// - once with the expected value (no assert)
		// - once with unexpected value (should assert if debug flag is on)
		for _, exp := range []bool{true, false} {

			got, msg := hasPanicked(func() { f(exp, tt.format, tt.args...) })
			// we want to panic if asserted value is not the one expected AND
			// the debug flag is set
			want := !exp && isDebug
			if got != want {
				if want {
					t.Error("should have panicked but didn't")
				} else {
					t.Errorf("should not have panicked but did, with:\n%v", msg)
				}
			}
		}
	}
}

func TestTrue(t *testing.T) {
	testAssert(t, func(exp bool, format string, a ...interface{}) { True(exp, a) })
}

func TestFalse(t *testing.T) {
	testAssert(t, func(exp bool, format string, a ...interface{}) { False(!exp, a) })
}

func TestTruef(t *testing.T) {
	testAssert(t, func(exp bool, format string, a ...interface{}) { Truef(exp, format, a) })
}

func TestFalsef(t *testing.T) {
	testAssert(t, func(exp bool, format string, a ...interface{}) { Falsef(!exp, format, a) })
}

func BenchmarkTrue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		True(true, "")
	}
}

func BenchmarkFalse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		False(false, "")
	}
}

func BenchmarkTruef(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Truef(true, "")
	}
}

func BenchmarkFalsef(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Falsef(false, "")
	}
}
