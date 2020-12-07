package aerrors

import (
	"errors"
	"fmt"
)

func ExampleNew() {
	err := New("new error")

	fmt.Println(err.Error())
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
	fmt.Println(err.Unwrap())
	// Output:
	// wrap error: oops
	// oops
}

func ExampleErr_New() {
	appError := New("app error")
	err := appError.New("oops", Priority(Info))

	fmt.Println(err)
	fmt.Println(err.Parent())
	// Output:
	// oops
	// app error
}

func ExampleErr_New_withOption() {
	appError := New("app error")
	err := appError.New("oops", Priority(Info))

	fmt.Println(err.Priority())
	fmt.Println(err.Parent().Priority())
	// Output:
	// Info
	// Error
}

func ExampleErr_Errorf() {
	appError := New("app error")
	err := appError.Errorf("error: %d", 42)

	fmt.Println(err)
	// Output:
	// error: 42
}

func ExampleErr_Wrap() {
	appError := New("app error")
	err := appError.Wrap(errors.New("oops"))

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
	// Output:
	// err: true
	// parent: true
	// friped: false
	// other: false
	// other child: false
}

func ExampleErr_WithValue() {
	err := New("new error").WithValue(String("str", "Foo"), Bool("bool", true))

	for _, v := range err.Values() {
		fmt.Printf("%s: %s\n", v.Label, v.Value)
	}
	// Output:
	// str: Foo
	// bool: true
}

func ExampleErr_Callers() {
	err := New("new error")
	f := err.Callers()
	for {
		frame, more := f.Next()
		fmt.Println(frame.Function)
		if !more {
			break
		}
	}

	// Output:
	// github.com/kamiaka/aerrors.ExampleErr_Callers
}

func ExampleErr_WithPriority() {
	err := New("new error")
	fmt.Println(err.Priority())

	err.WithPriority(Info)
	fmt.Println(err.Priority())

	// Output:
	// Error
	// Info
}

func ExampleErr_Config() {
	err := New("new error")

	fmt.Println(err.Config().CallerDepth)
	// Output:
	// 1
}
