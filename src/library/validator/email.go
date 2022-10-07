package validator

import "github.com/dalikewara/strgo"

// ValidateEmail validates email format.
//	- `email` can only contain alphanumeric characters and these special characters: _.-@+
//	- its length must be greater than 3 and not more than 255
//	- special character must be followed by at least one alphanumeric character
//	- prefix and suffix cannot be a special character
//	- must contain char @ and must be appeared once in the string
func ValidateEmail(email string) error {
	return strgo.Byte(email, &strgo.ByteCondition{
		MinLength:        4,
		MaxLength:        255,
		OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.', '@', '-', '+'}...),
		MustBeFollowedBy: [2][]byte{{'_', '.', '@', '-', '+'}, strgo.AlphanumericByte},
		MustContainsOnce: []byte{'@'},
	})
}
