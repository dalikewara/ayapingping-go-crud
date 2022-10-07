package validator

import "github.com/dalikewara/strgo"

// ValidateUsername validates username format.
//	- `username` can only contain alphanumeric characters, underscores and periods
//	- its length must be greater than 2 and not more than 20
//	- special character must be followed by at least one alphanumeric character
//	- prefix and suffix cannot be a special character
func ValidateUsername(username string) error {
	return strgo.Byte(username, &strgo.ByteCondition{
		MinLength:        3,
		MaxLength:        20,
		OnlyContains:     append(strgo.AlphanumericByte, []byte{'.', '_'}...),
		MustBeFollowedBy: [2][]byte{{'.', '_'}, strgo.AlphanumericByte},
	})
}
