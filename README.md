# aerrors

[![Go Reference](https://pkg.go.dev/badge/github.com/kamiaka/aerrors.svg)](https://pkg.go.dev/github.com/kamiaka/aerrors)

Aerrors is package for Golang augmented errors.

## Usage

Can be used like existing package.

Like an `errors.New`

```go
err := aerrors.New("new error")

fmt.Println(err)
// Output:
// new error
```

Like a `fmt.Errorf`

```go
err := aerrors.Errorf("error: %d", 42)

fmt.Println(err)
// Output:
// error: 42
```

Aerror's error has verbose information in addition to default errors.

```go
err := aerrors.New("new error")

fmt.Printf("%+v", err)
// Output:
// new error:
//     priority: Error
//     callers: main.main:<GOPATH>example/main.go:10
```

Can also add more information.

```go
err := aerrors.New("new error").WithValue(
  aerrors.String("foo", "Foo"),
  aerrors.Int("number", 42)
)

fmt.Printf("%+v", err)
// Output:
// new error:
//     priority: Error
//     callers: main.main:example/main.go:10
//     foo: Foo
//     number: 42
```

Can also create extend error.

```go
appError := aerrors.New("app error")
err := appError.New("oops")

fmt.Println(err)
fmt.Printf("is parent: %v", errors.Is(err, appError))
fmt.Printf("parent: %v", err.Parent)
// Output:
// oops
// is parent: true
// parent: app error
```

Can also create wrapped error.

```go
appError := appError := aerrors.New("application error")
err := appError.Errorf("error: %w", errors.New("oops"))

fmt.Println(err)
// Output:
// error: oops:
//     priority: Error
//     parent: application error
//     origin: oops
//     callers: main.main:example/main.go:12
//   - oops
```

### Trim GOPATH from callers and stack traces

Use `--trimpath` option.

```sh
go build --trimpath
```

## License

Aerrors is licensed under the [MIT](./LICENSE) license.
