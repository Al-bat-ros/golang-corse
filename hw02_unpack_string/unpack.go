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

	for i := 0; i < len(nStr); i++ {
		if unicode.IsDigit(nStr[0]) {
			return "", ErrInvalidString
		}

		if i < len(nStr)-1 && unicode.IsDigit(nStr[i]) && unicode.IsDigit(nStr[i+1]) {
			return "", ErrInvalidString
		}

		if i < len(nStr)-1 && unicode.IsDigit(nStr[i+1]) {
			num, _ := strconv.Atoi(string(nStr[i+1]))
			count = num
			if string(nStr[i]) != "\n" {
				mulStr := strings.Repeat(string(nStr[i]), count)
				resultStr.WriteString(mulStr)
			} else {
				mulStr := strings.Repeat("\\n", count)
				resultStr.WriteString(mulStr)
			}
		} else if !unicode.IsDigit(nStr[i]) && string(nStr[i]) != "\n" {
			resultStr.WriteString(string(nStr[i]))
		}
	}
	return resultStr.String(), nil
}
