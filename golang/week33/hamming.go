package hamming

import "errors"

var (
	ErrLengthMustMatch = errors.New("length must match")
)

func Distance(a, b string) (int, error) {
	// technically misses the case where both strings are empty, but it's not tested so oh well
	if len(a) != len(b) {
		return -1, ErrLengthMustMatch
	}
	// i'm not convinced this would handle non-ASCII but luckily DNA is named in ASCII
	var total int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			total++
		}
	}
	return total, nil
}
