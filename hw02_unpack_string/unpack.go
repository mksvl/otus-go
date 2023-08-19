package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/example/hello/reverse"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if len(input) != 0 {
		if unicode.IsDigit(rune(input[0])) {
			return "", ErrInvalidString
		}
	}

	var result strings.Builder

	// Переворачиваем строку, чтобы было удобно считать повторения
	input = reverse.String(input)
	var repeatCount int
	var prevIsDigit bool

	for _, char := range input {
		if unicode.IsDigit(char) {
			if prevIsDigit {
				return "", ErrInvalidString
			}
			// Флаг что предыдущий символ это число
			prevIsDigit = true
			// Преобразование руны-цифры в число
			repeatCount, _ = strconv.Atoi(string(char))
		} else {
			if prevIsDigit {
				result.WriteString(strings.Repeat(string(char), repeatCount))
				prevIsDigit = false
			} else {
				result.WriteRune(char)
			}
		}
	}

	finalResult := reverse.String(result.String())
	return finalResult, nil
}
