package aerrors

import (
	"errors"
	"fmt"
)

func ExampleNew() {
	err := New("new error")

	fmt.Println(err)
	// Output:
	// new error
}

func ExampleNew_verbose() {
	err := New("new error")

	fmt.Printf("%+v", err)
	// Output:
	// new error:
	//     priority: Error
	//     callers: aerrors.ExampleNew_verbose:github.com/kamiaka/aerrors/error_test.go:17
}

func ExampleNew_with_options() {
	err := New("new error", Priority(Emergency))

	fmt.Println(err.Priority())
	// Output:
	// Emergency
}

func ExampleErrorf() {
	err := Errorf("error: %d", 42)

	fmt.Println(err)
	// Output:
	// error: 42
}

func ExampleErrorf_with_wrapped_error() {
	origin := errors.New("oops")
	err := Errorf("wrap error: %w", origin)

	fmt.Println(err)
	// Output:
	// wrap error: oops
}

func ExampleErr_New() {
	apiError := New("api error")
	err := apiError.New("oops")

	fmt.Println(err)
	fmt.Println(err.Parent())
	// Output:
	// oops
	// api error
}

func ExampleErr_Wrap() {
	apiError := New("api error")
	err := apiError.Wrap(errors.New("oops"))

	fmt.Println(err)
	// Output:
	// oops
}

func ExampleErr_Is() {
	parent := New("parent error")
	err := parent.New("child error")
	other := errors.New("other error")
	otherChild := parent.New("other child error")

	fmt.Printf("err: %v\n", errors.Is(err, err))
	fmt.Printf("parent: %v\n", errors.Is(err, parent))
	fmt.Printf("friped: %v\n", errors.Is(parent, err))
	fmt.Printf("other: %v\n", errors.Is(err, other))
	fmt.Printf("other child: %v\n", errors.Is(err, otherChild))
	fmt.Printf("clone: %v\n", errors.Is(err, err.Clone()))
	// Output:
	// err: true
	// parent: true
	// friped: false
	// other: false
	// other child: false
	// clone: false
}
