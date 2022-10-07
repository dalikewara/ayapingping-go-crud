package strgo

import (
	"errors"
	"strconv"
	"strings"
)

type StringCondition struct {
	MinLength                 int
	MaxLength                 int
	OnlyContainsPrefixWord    []string
	OnlyContainsSuffixWord    []string
	MustContainsWord          []string
	MustContainsWordOnce      []string
	MustNotContainsWord       []string
	MustNotContainsPrefixWord []string
	MustNotContainsSuffixWord []string
	MayContainsWordOnce       []string
}

// String matches the string based on the StringCondition.
// If one doesn't match, it will return an error.
func String(text string, cond *StringCondition) error {
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

	if cond.OnlyContainsPrefixWord != nil {
		for _, v := range cond.OnlyContainsPrefixWord {
			if v != "" && text[:len(v)] == v {
				cond.OnlyContainsPrefixWord = nil
				break
			}
		}
		if cond.OnlyContainsPrefixWord != nil {
			return errors.New("the string prefix doesn't match with the given prefix words")
		}
	}
	if cond.OnlyContainsSuffixWord != nil {
		for _, v := range cond.OnlyContainsSuffixWord {
			if v != "" && text[textLen-len(v):] == v {
				cond.OnlyContainsSuffixWord = nil
				break
			}
		}
		if cond.OnlyContainsSuffixWord != nil {
			return errors.New("the string suffix doesn't match with the given suffix words")
		}
	}
	if cond.MustNotContainsPrefixWord != nil {
		for _, v := range cond.MustNotContainsPrefixWord {
			if v != "" && text[:len(v)] == v {
				return errors.New("the string must not contain prefix word: " + v)
			}
		}
	}
	if cond.MustNotContainsSuffixWord != nil {
		for _, v := range cond.MustNotContainsSuffixWord {
			if v != "" && text[textLen-len(v):] == v {
				return errors.New("the string must not contain suffix word: " + v)
			}
		}
	}

	if cond.MustContainsWord != nil {
		for _, v := range cond.MustContainsWord {
			if v != "" && !strings.Contains(text, v) {
				return errors.New("the string must contain word: " + v)
			}
		}
	}
	if cond.MustContainsWordOnce != nil {
		for _, v := range cond.MustContainsWordOnce {
			if v != "" && strings.Count(text, v) != 1 {
				return errors.New("the string must contain word: " + v + ", and it must be appeared once in the string")
			}
		}
	}
	if cond.MustNotContainsWord != nil {
		for _, v := range cond.MustNotContainsWord {
			if v != "" && strings.Contains(text, v) {
				return errors.New("the string must not contain word: " + v)
			}
		}
	}
	if cond.MayContainsWordOnce != nil {
		for _, v := range cond.MayContainsWordOnce {
			if v != "" && strings.Count(text, v) > 1 {
				return errors.New("the word: " + v + ", must be appeared once in the string")
			}
		}
	}

	return nil
}
