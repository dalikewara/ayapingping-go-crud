package validator

import "github.com/dalikewara/strgo"

// ValidatePassword validates password format.
//	- `password` can only contain alphanumeric characters and special characters: !"#$% &'()*+,-./:;<=>?@[\]^_`{|}~
//	- its length must be greater than 5 and not more than 32
//	- at least contain one lower and upper case letter, one number and one special character
func ValidatePassword(password string) error {
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
