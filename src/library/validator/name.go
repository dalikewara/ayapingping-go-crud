package validator

import (
	"github.com/dalikewara/strgo"
)

// ValidateName validates name format.
//	- `name` can only contain alphanumeric characters and these special characters: . '-
//	- its length must be greater than 1 and not more than 255
func ValidateName(name string) error {
	return strgo.Byte(name, &strgo.ByteCondition{
		MinLength:    2,
		MaxLength:    255,
		OnlyContains: append(strgo.AlphanumericByte, []byte{'.', ' ', '\'', '-'}...),
	})
}
