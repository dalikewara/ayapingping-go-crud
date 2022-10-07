# errgo

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/errgo)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/errgo)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/errgo)
![GitHub license](https://img.shields.io/github/license/dalikewara/errgo)

**errgo** provides custom interface for error `errgo.ErrGo` with `code` and `message` as a standard.

## Getting started

### Installation

You can use the `go get` method:

```bash
go get github.com/dalikewara/errgo
```

### Usage

```go
type ErrGo interface {
    Error() error
    GetError() error
    GetCode() string
    GetMessage() string
    GetStatus() int
}
```

#### Generate new error

```go
err := errgo.New("01", "data not found")

fmt.Println(err.GetCode()) // 01
fmt.Println(err.GetMessage()) // data not found
fmt.Println(err.GetError().Error()) // 01||data not found
```

#### Generate new error with status

```go
err := errgo.NewWithStatus("01", "data not found", 200)

fmt.Println(err.GetCode()) // 01
fmt.Println(err.GetMessage()) // data not found
fmt.Println(err.GetStatus()) // 200
fmt.Println(err.GetError().Error()) // 01||data not found
```

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/errgo/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2022 [Dali Kewara](https://www.dalikewara.com)

### License

[MIT License](https://github.com/dalikewara/errgo/blob/master/LICENSE)
