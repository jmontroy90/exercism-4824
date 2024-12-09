package diamond

import (
	"errors"
	"strings"
)

var (
	ErrMustBeCapitalLetter = errors.New("ErrMustBeCapitalLetter")
)

func Gen(char byte) (string, error) {
	r := rune(char)
	if r < 65 || r > 90 {
		return "", ErrMustBeCapitalLetter
	}
	if r == 'A' {
		return "A", nil
	}
	// Ok, I've possibly coded the weirdest, worst version of this imaginable.
	// I insisted on determining the placement positions from the midpoint, instead of from the beginning and end.
	// I did this with some crazy-ass modulo, which I actually needed Python's modulo (per below), not Golang's.
	// I need to spend time with this and think about why I borked it so bad on this.
	// First impression - you can either calculate indexes from some arbitrary point, OR from head and tail. Maybe think on both before acting.
	// And people didn't worry about decrementing the rune either - you just ALSO start placing rows from the end up.
	// So take advantage of symmetry!
	diff := int(r - 'A')
	dim := ((diff + 1) * 2) - 1
	mid := ((dim + 1) / 2) - 1
	var sb strings.Builder
	for row := 0; row < dim; row++ {
		for col := 0; col < dim; col++ {
			p1 := mod(mid-row, dim-1)
			p2 := mod(mid+row, dim-1)
			if row == mid { // special case for center where mod fails us
				p1, p2 = 0, dim-1
			}
			if col == abs(p1) || col == abs(p2) {
				sb.WriteRune(rune('A' + abs(abs(row-mid)-mid)))
			} else {
				sb.WriteRune(' ')
			}
		}
		if row+1 == dim {
			break
		}
		sb.WriteString("\n")
	}
	return sb.String(), nil
}

// See: https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go
func mod(a, b int) int {
	return (a%b + b) % b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
