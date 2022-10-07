# strgo

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/strgo)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/strgo)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/strgo)
![GitHub license](https://img.shields.io/github/license/dalikewara/strgo)

**strgo** is a helper to validate a string based on a format specification that you set before. You can use this package to validate
some common string case like `username`, `name`, `password`, `email`, etc.

## Getting started

### Installation

You can use the `go get` method:

```bash
go get github.com/dalikewara/strgo
```

### Usage

This is an example to validate a `username` string:

- `username` can only contain alphanumeric characters, underscores and periods
- its length must be greater than 2 and not more than 20
- special character must be followed by at least one alphanumeric character
- prefix and suffix cannot be a special character
- allowed special characters must be appeared once in the string

```go
func validate(username string) error {
    return strgo.Byte(username, &strgo.ByteCondition{
        MinLength:        3,
        MaxLength:        20,
        OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.'}...),
        MustBeFollowedBy: [2][]byte{{'_', '.'}, strgo.AlphanumericByte},
        MayContainsOnce:  []byte{'_', '.'},
    })
}
validate("johndoe") // valid
validate("john_doe") // valid
validate("john.doe") // valid
validate("johndoe123") // valid
validate("john.doe123") // valid
validate("john_doe123") // valid
validate("john_doe.123") // valid
validate("_johndoe") // not valid
validate("__johndoe") // not valid
validate(".johndoe") // not valid
validate("..johndoe") // not valid
validate("johndoe_") // not valid
validate("johndoe__") // not valid
validate("johndoe.") // not valid
validate("johndoe..") // not valid
validate("john__doe") // not valid
validate("john_.doe") // not valid
validate("john..doe") // not valid
validate("john._doe") // not valid
validate("john@doe") // not valid
validate("@johndoe") // not valid
validate("johndoe@") // not valid
validate("joh_nd_oe") // not valid
validate("joh.nd.oe") // not valid
```

...this is an example to validate an `email` string:

- `email` can only contain alphanumeric characters and these special characters: `_.-@+`
- its length must be greater than 3 and not more than 255
- special character must be followed by at least one alphanumeric character
- prefix and suffix cannot be a special character
- must contain char `@` and must be appeared once in the string

```go
func validate(email string) error {
    return strgo.Byte(email, &strgo.ByteCondition{
        MinLength:        4,
        MaxLength:        255,
        OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.', '@', '-', '+'}...),
        MustBeFollowedBy: [2][]byte{{'_', '.', '@', '-', '+'}, strgo.AlphanumericByte},
        MustContainsOnce: []byte{'@'},
    })
}
validate("johndoe@email.com") // valid
validate("john_doe@email.com") // valid
validate("john_do.e@email.com") // valid
validate("john-doe@email.com") // valid
validate("johndoe@email") // valid
validate("johndoe123@email") // valid
validate("johndoe123@email") // valid
validate("john+doe123@email") // valid
validate("johndoe123email") // not valid
validate("johndoe123.email") // not valid
validate("john@doe123@email") // not valid
validate(".johndoe123@email") // not valid
validate("johndoe123@email.") // not valid
validate("johndoe123@") // not valid
validate("john_.doe123@email") // not valid
validate("johndoe123.@email") // not valid
```

...and this is an example to validate a `password` string:

- `password` can only contain alphanumeric characters and special characters: !"#$% &'()*+,-./:;<=>?@[\]^_`{|}~
- its length must be greater than 5 and not more than 32
- at least contain one lower and upper case letter, one number and one special character

```go
func validate(password string) error {
    return strgo.Byte(password, &strgo.ByteCondition{
        MinLength:                   6,
        MaxLength:                   32,
        OnlyContains:                strgo.CharsByte,
        AtLeastHaveUpperLetterCount: 1,
        AtLeastHaveLowerLetterCount: 1,
        AtLeastHaveNumberCount:      1,
        AtLeastHaveSpecialCharCount: 1,
    })
}
validate("J()hndoe123") // valid
validate("John_doe123") // valid
validate("johndoe") // not valid
validate("johndoe123") // not valid
validate("Johndoe123") // not valid
```

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/strgo/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2022 [Dali Kewara](https://www.dalikewara.com)

### License

[MIT License](https://github.com/dalikewara/strgo/blob/master/LICENSE)
