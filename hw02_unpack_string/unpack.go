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
		if unicode.IsDigit(r) {
			if i == 0 || unicode.IsDigit(stringUnicode[i-1]) {
				return "", ErrInvalidString
			}

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

			// Добавляем символ repeatCount - 1 раз (один раз он уже добавился на итерации с буквой)
			for j := 0; j < repeatCount-1; j++ {
				result.WriteRune(stringUnicode[i-1])
			}
		} else {
			// Если это не цифра, просто записываем символ
			result.WriteRune(r)
		}
	}

	return result.String(), nil
}
