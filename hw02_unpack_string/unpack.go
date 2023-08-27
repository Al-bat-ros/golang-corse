package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	nStr := []rune(str)
	var resultStr strings.Builder
	var count int
	if len(nStr) != 0 {
		if unicode.IsDigit(nStr[0]) {
			return "", ErrInvalidString
		}
	}
	for i := range nStr {
		if i < len(nStr)-1 && unicode.IsDigit(nStr[i]) && unicode.IsDigit(nStr[i+1]) {
			return "", ErrInvalidString
		}

		if i < len(nStr)-1 && unicode.IsDigit(nStr[i+1]) {
			if num, err := strconv.Atoi(string(nStr[i+1])); err == nil {
				count = num
			}
			if count > 0 {
				mulStr := strings.Repeat(string(nStr[i]), count)
				resultStr.WriteString(mulStr)
			}
		} else if !unicode.IsDigit(nStr[i]) {
			resultStr.WriteString(string(nStr[i]))
		}
	}
	return resultStr.String(), nil
}
