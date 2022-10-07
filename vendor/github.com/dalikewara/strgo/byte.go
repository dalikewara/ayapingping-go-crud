package strgo

import (
	"errors"
	"strconv"
)

const asciiMaxLen = 128

const asciiMaxDec = 127

type asciis [asciiMaxLen]byte

type ByteCondition struct {
	MinLength                   int
	MaxLength                   int
	OnlyContains                []byte
	OnlyContainsPrefix          []byte
	OnlyContainsSuffix          []byte
	MustContains                []byte
	MustContainsOnce            []byte
	MustNotContains             []byte
	MustNotContainsPrefix       []byte
	MustNotContainsSuffix       []byte
	MustBeFollowedBy            [2][]byte
	MayContainsOnce             []byte
	AtLeastHaveUpperLetterCount int
	AtLeastHaveLowerLetterCount int
	AtLeastHaveNumberCount      int
	AtLeastHaveSpecialCharCount int
}

// Byte matches the string based on the ByteCondition.
// If one doesn't match, it will return an error.
// This function can only validate ASCII characters (0-127).
// Ref: https://en.wikipedia.org/wiki/ASCII
func Byte(text string, cond *ByteCondition) error {
	if text == "" {
		return errors.New("the string is empty")
	}

	textLen := len(text)

	if cond.MinLength > 0 && textLen < cond.MinLength {
		return errors.New("the string length cannot be less than " + strconv.Itoa(cond.MinLength))
	}
	if cond.MaxLength > 0 && textLen > cond.MaxLength {
		return errors.New("the string length cannot be more than " + strconv.Itoa(cond.MaxLength))
	}

	var (
		onlyContains,
		onlyContainsPrefix,
		onlyContainsSuffix,
		mustContains,
		mustNotContains,
		mustNotContainsPrefix,
		mustNotContainsSuffix,
		mustBeFollowedBy,
		mustBeFollowedByPairs,
		mayContainsOnce asciis
		atLeastHaveUpperLetterCount,
		atLeastHaveLowerLetterCount,
		atLeastHaveNumberCount,
		atLeastHaveSpecialCharCount int
	)

	if cond.OnlyContains != nil {
		setASCIICond(&onlyContains, &cond.OnlyContains)
	}
	if cond.OnlyContainsPrefix != nil {
		setASCIICond(&onlyContainsPrefix, &cond.OnlyContainsPrefix)
	}
	if cond.OnlyContainsSuffix != nil {
		setASCIICond(&onlyContainsSuffix, &cond.OnlyContainsSuffix)
	}
	if cond.MustContains != nil {
		setASCIICond(&mustContains, &cond.MustContains)
	}
	if cond.MustContainsOnce != nil {
		setASCIICondDouble(&mustContains, &mayContainsOnce, &cond.MustContainsOnce)
	}
	if cond.MustNotContains != nil {
		setASCIICond(&mustNotContains, &cond.MustNotContains)
	}
	if cond.MustNotContainsPrefix != nil {
		setASCIICond(&mustNotContainsPrefix, &cond.MustNotContainsPrefix)
	}
	if cond.MustNotContainsSuffix != nil {
		setASCIICond(&mustNotContainsSuffix, &cond.MustNotContainsSuffix)
	}
	if cond.MayContainsOnce != nil {
		setASCIICond(&mayContainsOnce, &cond.MayContainsOnce)
	}
	if cond.MustBeFollowedBy[0] != nil && cond.MustBeFollowedBy[1] != nil {
		setASCIICond(&mustBeFollowedBy, &cond.MustBeFollowedBy[0])
		setASCIICond(&mustBeFollowedByPairs, &cond.MustBeFollowedBy[1])
	}
	atLeastHaveUpperLetterCount = cond.AtLeastHaveUpperLetterCount
	atLeastHaveLowerLetterCount = cond.AtLeastHaveLowerLetterCount
	atLeastHaveNumberCount = cond.AtLeastHaveNumberCount
	atLeastHaveSpecialCharCount = cond.AtLeastHaveSpecialCharCount

	textLenMaxIndex := textLen - 1

	for i, v := range text {
		if v > asciiMaxDec {
			return errors.New("the char: " + string(v) + ", is not a valid ascii format")
		}
		if i == 0 {
			if cond.OnlyContainsPrefix != nil && onlyContainsPrefix[v] < 1 {
				return errors.New("the string cannot contain prefix char: " + string(v))
			}
			if cond.MustNotContainsPrefix != nil && mustNotContainsPrefix[v] > 0 {
				return errors.New("the string must not contain prefix: " + string(v))
			}
		}
		if i == textLenMaxIndex {
			if cond.OnlyContainsSuffix != nil && onlyContainsSuffix[v] < 1 {
				return errors.New("the string cannot contain suffix char: " + string(v))
			}
			if cond.MustNotContainsSuffix != nil && mustNotContainsSuffix[v] > 0 {
				return errors.New("the string must not contain suffix: " + string(v))
			}
		}
		if cond.OnlyContains != nil && onlyContains[v] < 1 {
			return errors.New("the string cannot contain char: " + string(v))
		}
		if cond.MustNotContains != nil && mustNotContains[v] > 0 {
			return errors.New("the string must not contain char: " + string(v))
		}
		if (cond.MustContains != nil || cond.MustContainsOnce != nil) && mustContains[v] > 0 {
			mustContains[v] = 0
		}
		if (cond.MayContainsOnce != nil || cond.MustContainsOnce != nil) && mayContainsOnce[v] > 0 {
			if mayContainsOnce[v] > 1 {
				return errors.New("the char: " + string(v) + ", must be appeared once in the string")
			}
			mayContainsOnce[v] += 1
		}
		if cond.MustBeFollowedBy[0] != nil && cond.MustBeFollowedBy[1] != nil && mustBeFollowedBy[v] > 0 {
			if i == 0 || (i+1) == textLen {
				return errors.New("the char: " + string(v) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
			}
			if i > 0 && i < textLen && mustBeFollowedByPairs[text[i-1]] < 1 {
				return errors.New("the char: " + string(v) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
			}
			if (i+1) < textLen && mustBeFollowedByPairs[text[i+1]] < 1 {
				return errors.New("the char: " + string(v) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
			}
		}
		if atLeastHaveUpperLetterCount > 0 && (v >= 'A' && v <= 'Z') {
			atLeastHaveUpperLetterCount -= 1
		}
		if atLeastHaveLowerLetterCount > 0 && (v >= 'a' && v <= 'z') {
			atLeastHaveLowerLetterCount -= 1
		}
		if atLeastHaveNumberCount > 0 && (v >= '0' && v <= '9') {
			atLeastHaveNumberCount -= 1
		}
		if atLeastHaveSpecialCharCount > 0 && (v < 'A' || v > 'Z') && (v < 'a' || v > 'z') && (v < '0' || v > '9') {
			atLeastHaveSpecialCharCount -= 1
		}
	}
	if cond.MustContains != nil || cond.MustContainsOnce != nil {
		for b, v := range mustContains {
			if v > 0 {
				return errors.New("the string must contain char: " + string(rune(b)))
			}
		}
	}
	if atLeastHaveUpperLetterCount > 0 {
		return errors.New("the string must have at least " + strconv.Itoa(cond.AtLeastHaveUpperLetterCount) + " upper case letter(s)")
	}
	if atLeastHaveLowerLetterCount > 0 {
		return errors.New("the string must have at least " + strconv.Itoa(cond.AtLeastHaveLowerLetterCount) + " lower case letter(s)")
	}
	if atLeastHaveNumberCount > 0 {
		return errors.New("the string must have at least " + strconv.Itoa(cond.AtLeastHaveNumberCount) + " number(s)")
	}
	if atLeastHaveSpecialCharCount > 0 {
		return errors.New("the string must have at least " + strconv.Itoa(cond.AtLeastHaveSpecialCharCount) + " special char(s)")
	}

	return nil
}

func setASCIICond(c *asciis, b *[]byte) {
	for _, v := range *b {
		c[v] = 1
	}
}

func setASCIICondDouble(c, c2 *asciis, b *[]byte) {
	for _, v := range *b {
		c[v] = 1
		c2[v] = 1
	}
}
