package tokenizer

import (
	"errors"
	"fmt"
)

const (
	SEPARATOR TokenKey = 0
	NUMBER    TokenKey = 1
	WORD      TokenKey = 2
	SYMBOL    TokenKey = 3
)

var ErrNoProgress = errors.New("no progress")

type TokenKey int
type Token struct {
	Key   TokenKey
	Value string
}

type TokenizerContext struct {
	Source     string
	Index      int
	Separators []byte
	Tokens     []Token
}

type TokenizerOption func(*TokenizerContext)

func Tokenize(source string, opts ...TokenizerOption) ([]Token, error) {
	result := make([]Token, 0)
	context := &TokenizerContext{
		Source: source,
		Index:  0,
	}
	for _, opt := range opts {
		opt(context)
	}

	for {
		lastIndex := context.Index
		parseWhiteSpaces(context)
		if context.Index >= len(context.Source) {
			break
		}
		if separator, err := parseSeparator(context); err == nil {
			result = append(result, Token{
				Key:   SEPARATOR,
				Value: separator,
			})
			continue
		}
		if number, err := parseNumber(context); err == nil {
			result = append(result, Token{
				Key:   NUMBER,
				Value: number,
			})
			continue
		}
		if word, err := parseWord(context); err == nil {
			result = append(result, Token{
				Key:   WORD,
				Value: word,
			})
			continue
		}
		if lastIndex == context.Index {
			return nil, fmt.Errorf("%w", ErrNoProgress)
		}
	}

	return result, nil
}

func WithSeparators(separators []byte) TokenizerOption {
	return func(r *TokenizerContext) {
		r.Separators = separators
	}
}

func parseWhiteSpaces(context *TokenizerContext) {
	for context.Index < len(context.Source) {
		char := context.Source[context.Index]
		matched := false
		for _, ws := range []byte{' ', '\t'} {
			if char == ws {
				matched = true
				context.Index++
				break
			}
		}
		if !matched {
			return
		}
	}
}

func parseSeparator(context *TokenizerContext) (string, error) {
	for _, separator := range context.Separators {
		if context.Source[context.Index] == separator {
			context.Index += 1
			return string(separator), nil
		}
	}
	return "", errors.New("no separator")
}

func parseNumber(context *TokenizerContext) (string, error) {
	var acc []byte
	mustBeDigit := false
	hasDot := false
	pos := context.Index
	for ; pos < len(context.Source); pos++ {
		char := context.Source[pos]
		if char >= '0' && char <= '9' {
			mustBeDigit = false
			acc = append(acc, char)
		} else {
			if mustBeDigit {
				return "", errors.New("no number")
			}

			if char == '.' {
				if len(acc) == 0 || hasDot {
					return "", errors.New("no number")
				}
				mustBeDigit = true
				hasDot = true
				acc = append(acc, char)
				continue
			}

			break
		}
	}
	if len(acc) > 0 {
		context.Index = pos
		return string(acc), nil
	} else {
		return "", errors.New("no number")
	}
}

func parseWord(context *TokenizerContext) (string, error) {
	var acc []byte
	for context.Index < len(context.Source) {
		char := context.Source[context.Index]
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			acc = append(acc, char)
			context.Index++
		} else {
			break
		}
	}
	if len(acc) > 0 {
		return string(acc), nil
	} else {
		return "", errors.New("no word")
	}
}
