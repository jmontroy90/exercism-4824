package romannumerals

import (
	"errors"
	"fmt"
	"strings"
)

var (
	roman  = []rune{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
	arabic = []int{1000, 500, 100, 50, 10, 5, 1}
)

var (
	ErrorOutOfRange = errors.New("out of range, must be 0 > n < 4000")
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", ErrorOutOfRange
	}
	var res, stg strings.Builder
	for i, m := range arabic {
		// base conversion
		for input >= m {
			stg.WriteRune(roman[i])
			input -= m // so we work with the proper remainder next time
		}
		// special case: we wrote 4 of the same numeral consecutively
		if s := stg.String(); len(s) == 4 {
			res.WriteString(fmt.Sprintf("%c%c", roman[i], roman[i-1]))
		} else {
			res.WriteString(s)
		}
		// special case: YZY => ZX
		// this method of examining your results post-pass seems dubious
		if r := res.String(); hasNumeralInversion(r) {
			res.Reset()
			res.WriteString(fmt.Sprintf("%s%c%c", r[:len(r)-3], roman[i], roman[i-2]))
		}
		stg.Reset()
	}
	return res.String(), nil
}

// a "numeral inversion" is of the form YZY, needs to be ZX;
// e.g. VIV => IX, LXL => XC, DCD => CM
func hasNumeralInversion(r string) bool {
	return len(r) >= 3 &&
		r[len(r)-3] == r[len(r)-1] && // check third-to-last == last
		r[len(r)-2] != r[len(r)-1] // check second-to-last != last
}
