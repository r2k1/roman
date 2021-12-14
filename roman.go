package roman

import (
	"fmt"
	"strconv"
)

var romanToInt = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var unicodeToASCIIRomanMapping = map[rune]string{
	'I': "I",
	'V': "V",
	'X': "X",
	'L': "L",
	'C': "C",
	'D': "D",
	'M': "M",
	'i': "I",
	'v': "V",
	'x': "X",
	'l': "L",
	'c': "C",
	'd': "D",
	'm': "M",
	// NOTE: next characters are not ASCII
	// It's unicode runes https://www.unicode.org/charts/PDF/U2150.pdf not strings
	'Ⅰ': "I",
	'Ⅱ': "II",
	'Ⅲ': "III",
	'Ⅳ': "IV",
	'Ⅴ': "V",
	'Ⅵ': "VI",
	'Ⅶ': "VII",
	'Ⅷ': "VIII",
	'Ⅸ': "IX",
	'Ⅹ': "X",
	'Ⅺ': "XI",
	'Ⅻ': "XII",
	'Ⅼ': "L",
	'Ⅽ': "C",
	'Ⅾ': "D",
	'Ⅿ': "M",
	// TODO: add other unicode characters
}

// ToArabic converts a Roman number to its equivalent Arabic representation
// It makes a best guess when parsing an incorrect input.
// In this case an error will be returned along with a best guess.
// The input is assumed to be in an "Irregular subtractive notation".
// It may be not what Roman used, but it's something widely used in a modern world.
// It seems to cover main cases written in "Additive notation"
// https://en.wikipedia.org/wiki/Roman_numerals
// TODO: add support for Fractions, using string as an output leaves our some room for choosing an output format
// TODO: should it return an error for an incorrect input, (such as "XXXX"?)
func ToArabic(roman string) (string, error) {
	// https://en.wikipedia.org/wiki/Roman_numerals#Zero
	if roman == "nulla" {
		return "0", nil
	}
	var err error
	// attempt to parse incorrect input
	letters, err := transformInput(roman)
	if len(roman) == 0 {
		//return "", mergeErrors(err, errors.New("input is empty"))
	}
	result := 0
	for i := range letters {
		current, ok := romanToInt[letters[i]]
		if !ok {
			return "", fmt.Errorf("internal error, unexpected character %s", strconv.QuoteRune(letters[i]))
		}
		next := 0
		if i+1 < len(letters) {
			next = romanToInt[letters[i+1]]
		}
		if next != 0 && current < next {
			result -= current
		} else {
			result += current
		}
	}
	return strconv.Itoa(result), err
}

// transformInput transforms unicode string to a fixed list of characters represented in romanToInt.
// When it encounters an unknown characters it removes it from the input and returns an error.
func transformInput(input string) ([]rune, error) {
	output := ""
	var err error
	for pos, letter := range input {
		// a single unicode rune may be replaced with multiple ASCII runes
		roman, ok := unicodeToASCIIRomanMapping[letter]
		if !ok {
			err = mergeErrors(err, fmt.Errorf("unexpected symbol %s at position %v", strconv.QuoteRune(letter), pos))
			continue
		}
		output += roman
	}
	return []rune(output), err
}

func mergeErrors(err1, err2 error) error {
	if err1 == nil {
		return err2
	}
	if err2 == nil {
		return err1
	}
	return fmt.Errorf("%v, %v", err1, err2)
}
