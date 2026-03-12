package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var result strings.Builder
	stringUnicode := []rune(s)

	for i, r := range stringUnicode {
		if i == 0 && unicode.IsDigit(r) {
			return "", ErrInvalidString
		}
		if i > 0 && unicode.IsDigit(r) && unicode.IsDigit(stringUnicode[i-1]) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(r) {
			repeatCount, _ := strconv.Atoi(string(r))

			if repeatCount == 0 {
				currentStr := result.String()
				if len(currentStr) > 0 {
					tmpChar := currentStr[:len(currentStr)-len(string(stringUnicode[i-1]))]
					result.Reset()
					result.WriteString(tmpChar)
				}
				continue
			}

			subStringRepeat := strings.Repeat(string(stringUnicode[i-1]), repeatCount-1)
			result.WriteString(subStringRepeat)
		} else {
			result.WriteRune(r)
		}
	}

	return result.String(), nil
}
