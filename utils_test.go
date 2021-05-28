package aerrors

import (
	"errors"
	"fmt"
)

func ExampleAsErr() {
	parent := New("parent error")
	child := parent.New("child error")
	other := errors.New("other error")
	wrapped := fmt.Errorf("wrapped error: %w", child)

	cases := []error{
		parent,
		child,
		other,
		wrapped,
	}

	for i, err := range cases {
		e, ok := AsErr(err)
		fmt.Printf("#%d: %v, %v\n", i, ok, e)
	}

	// Output:
	// #0: true, parent error
	// #1: true, child error
	// #2: false, <nil>
	// #3: true, child error
}
