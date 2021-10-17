# semerr

![Version](https://img.shields.io/github/v/tag/hedhyw/semerr)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedhyw/semerr)](https://goreportcard.com/report/github.com/hedhyw/semerr)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hedhyw/semerr)](https://pkg.go.dev/github.com/hedhyw/semerr?tab=doc)


Package `semerr` helps to work with errors in Golang.

<img alr="Go Bug" src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GO_BUG.png" width="100px">

## Multierror

If you need to handle more than one err, you can use `MultiError`:

```go
defer func() {
    // It will skip all nil errors.
    err = semerr.NewMultiError(err, f.Close())
}()
```

## Const error

An error that can be defined as `const`.

```go
var errMutable error = errors.New("mutable error") // Do not like this?
const errImmutable semerr.Error = "immutable error" // So use this.
```

## TemporaryError checking

The function `semerr.IsTemporaryError` checks that error has Temporary
method and it returns true.

```go
semerr.IsTemporaryError(context.DeadlineExceeded) // true
semerr.IsTemporaryError(context.Canceled) // false
```

## Errors with meaning

```go
errOriginal := errors.New("some error")
errWrapped := semerr.NewBadRequestError(errOriginal) // The text will be the same.

fmt.Println(errWrapped) // "some error"
fmt.Println(httperr.Code(errWrapped)) // http.StatusBadRequest
fmt.Println(grpcerr.Code(errWrapped)) // codes.InvalidArgument
fmt.Println(errors.Is(err, errOriginal)) // true
fmt.Println(semerr.NewBadRequestError(nil)) // nil
```

Also see:
```go
err := errors.New("some error")
err = semerr.NewBadRequestError(err)
err = semerr.NewForbiddenError(err)
err = semerr.NewInternalServerError(err)
err = semerr.NewNotFoundError(err)
err = semerr.NewServiceUnavailableError(err)
err = semerr.NewUnauthorizedError(err)
```

## Contributing

Pull requests are welcomed. If you want to add a new meaning error, then
edit the file
[internal/cmd/generator/errors.json](internal/cmd/generator/errors.json)
and generate a new code, for this run `make`.