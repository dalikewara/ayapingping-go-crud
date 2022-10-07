package entity

import (
	"github.com/dalikewara/ayapingping-go-crud/src/library/validator"
	"github.com/dalikewara/errgo"
	"time"
)

// ID

type ID uint32

func (i *ID) Primitive() uint32 {
	return uint32(*i)
}

func (i *ID) IsValid() bool {
	return *i > 0
}

// Username

type Username string

func (u *Username) Primitive() string {
	return string(*u)
}

func (u *Username) Validate() error {
	return validator.ValidateUsername(u.Primitive())
}

// Email

type Email string

func (e *Email) Primitive() string {
	return string(*e)
}

func (e *Email) Validate() error {
	return validator.ValidateEmail(e.Primitive())
}

// Password

type Password string

func (p *Password) Primitive() string {
	return string(*p)
}

func (p *Password) Validate() error {
	return validator.ValidatePassword(p.Primitive())
}

// Active Status

type ActiveStatus uint8

func (a *ActiveStatus) Primitive() uint8 {
	return uint8(*a)
}

func (a *ActiveStatus) IsActive() bool {
	return *a == 1
}

// Time

type Time string

type Timezone struct {
	Name   string
	Offset int
}

func (t *Time) Primitive() string {
	return string(*t)
}

func (t *Time) SetFromTime(tm time.Time, tz Timezone) {
	*t = Time(tm.In(time.FixedZone(tz.Name, tz.Offset)).Format(time.RFC3339))
}

func (t *Time) ToTime() (time.Time, error) {
	return time.Parse(time.RFC3339, t.Primitive())
}

func (t *Time) IsZero() bool {
	parsedTime, err := t.ToTime()
	if err != nil {
		return true
	}
	return parsedTime.IsZero()
}

// FirstName

type FirstName string

func (f *FirstName) Primitive() string {
	return string(*f)
}

func (f *FirstName) Validate() error {
	return validator.ValidateName(f.Primitive())
}

func (f *FirstName) IsEmpty() bool {
	return *f == ""
}

// LastName

type LastName string

func (l *LastName) Primitive() string {
	return string(*l)
}

func (l *LastName) Validate() error {
	return validator.ValidateName(l.Primitive())
}

func (l *LastName) IsEmpty() bool {
	return *l == ""
}

// Image

type Image string

func (i *Image) Primitive() string {
	return string(*i)
}

// Gender

type Gender uint8

func (g *Gender) Primitive() uint8 {
	return uint8(*g)
}

func (g *Gender) IsEmpty() bool {
	return *g == 0
}

func (g *Gender) IsValid() bool {
	return *g > 0 && *g <= 2
}

func (g *Gender) IsMale() bool {
	return *g == 1
}

func (g *Gender) IsFemale() bool {
	return *g == 2
}

// Username or email

type UsernameOrEmail string

func (u *UsernameOrEmail) Primitive() string {
	return string(*u)
}

func (u *UsernameOrEmail) Validate() error {
	if err := validator.ValidateEmail(u.Primitive()); err != nil {
		if err = validator.ValidateUsername(u.Primitive()); err != nil {
			return err
		}
	}
	return nil
}

// Error standard

type StdError interface {
	GetError() error
	GetCode() string
	GetMessage() string
	GetStatus() int
}

// NewStdError generates new StdError.
func NewStdError(code, message string, status int) StdError {
	return errgo.NewWithStatus(code, message, status)
}
