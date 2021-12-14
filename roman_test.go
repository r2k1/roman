package roman

import "testing"

func TestRomanToArabic(t *testing.T) {
	testCases := []struct {
		Roman    string
		Expected string
		Valid    bool
	}{
		// Base cases
		{"nulla", "0", true},
		{"I", "1", true},
		{"V", "5", true},
		{"X", "10", true},
		{"L", "50", true},
		{"C", "100", true},
		{"D", "500", true},
		{"M", "1000", true},
		//// Additive cases
		{"III", "3", true},
		{"IIIIIIIIII", "10", true},
		{"XXX", "30", true},
		{"MDCLXVI", "1666", true},
		//// Subtractive simple cases
		{"IV", "4", true},
		{"IX", "9", true},
		{"XL", "40", true},
		{"XC", "90", true},
		{"CD", "400", true},
		{"CM", "900", true},
		// More complex cases
		{"IIX", "10", true},
		{"IIXLL", "90", true}, // 1 - 1 - 10 + 50 + 50
		// Make sure we support output of Roman() function from Microsoft Excel
		{"CDXCIX", "499", true},
		{"LDVLIV", "499", true},
		{"XDIX", "499", true},
		{"VDIV", "499", true},
		{"ID", "499", true},
		{"CDXCV", "495", true},
		{"LDVL", "495", true},
		{"XDV", "495", true},
		{"VD", "495", true},
		// Unicode
		{"Ⅵ", "6", true},
		{"ⅤⅠ", "6", true},
		{"\u2164\u2160", "6", true}, // same as above
		// 1 + 1 + 1 - 1 + 5 + 5 + 1
		{"Ⅲ Ⅳ Ⅵ", "13", false},
		// Errors
		{" II ", "2", false},
		{"I I", "2", false},
		{"", "", false},
		{"V+V", "10", false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Roman, func(t *testing.T) {
			actual, err := ToArabic(testCase.Roman)
			// usually I would use testify/assert,
			// but for the sake of not introducing extra dependencies  use standard library
			if testCase.Valid && err != nil {
				t.Errorf("expected not error got %v", err)
			}
			if !testCase.Valid && err == nil {
				t.Errorf("expected error, got nothing")
			}
			if testCase.Expected != actual {
				t.Errorf("expected %v to be equal to %v", testCase.Expected, actual)
			}
		})
	}
}

func TestRomanToArabic_ErrorMessage(t *testing.T) {
	result, err := ToArabic("HELLO")
	if err.Error() != "unexpected symbol 'H' at position 0, unexpected symbol 'E'"+
		" at position 1, unexpected symbol 'O' at position 4" {
		t.Errorf("unexpected error %v", err)
	}
	if result != "100" {
		t.Errorf("expected 100, got: %v", result)
	}
}
