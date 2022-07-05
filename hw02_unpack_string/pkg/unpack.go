package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// Lazy check for multiple numbers in raw.
func isInvalid(input string) bool {
	invalid, _ := regexp.MatchString(`\d{2}|^[0-9]`, input)
	return invalid
}

func processInputRunes(runes []rune) string {
	unpackedString := strings.Builder{}
	unpackedString.Grow(len(runes))

	for i := 0; i < len(runes); i++ {
		// Don't append numbers to strings builder.
		if unicode.IsDigit(runes[i]) {
			continue
		}
		repeats := 1
		repeatedSymbol := string(runes[i])
		// Check if we are at the end of the string.
		if len(runes)-i > 1 {
			// Check if the next rune is a number and get it's value.
			if unicode.IsDigit(runes[i+1]) {
				repeats, _ = strconv.Atoi(string(runes[i+1]))
			}
			// Check if we are almost at the end of the string.
			if len(runes)-i > 2 {
				// Trying to find symbols like `\n`.
				if unicode.IsDigit(runes[i+2]) && (string(runes[i]) == `\`) {
					repeats, _ = strconv.Atoi(string(runes[i+2]))
					repeatedSymbol = `\` + string(runes[i+1])
					i++
				}
			}
		}
		unpackedString.WriteString(strings.Repeat(repeatedSymbol, repeats))
	}
	return unpackedString.String()
}

func Unpack(input string) (string, error) {
	if isInvalid(input) {
		return "", ErrInvalidString
	}
	return processInputRunes([]rune(input)), nil
}
