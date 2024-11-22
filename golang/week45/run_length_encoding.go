package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return input
	}
	curr := rune(input[0])
	count := 1
	var res string
	for _, r := range []rune(input[1:]) {
		if curr == r {
			count++
			continue
		}
		res += addForCount(curr, count)
		count = 1
		curr = r
	}
	// encode last bit, since we were just counting
	return res + addForCount(curr, count)
}

func addForCount(c rune, n int) string {
	if n == 1 {
		return fmt.Sprintf("%v", string(c))
	} else {
		return fmt.Sprintf("%v%v", n, string(c))
	}
}

func RunLengthDecode(input string) string {
	if len(input) == 0 {
		return input
	}
	var startSlice, currNum int
	var buf string
	for ix, r := range []rune(input) {
		if unicode.IsDigit(r) {
			continue // still reading numbers
		}
		if startSlice == ix { // no digit was read
			currNum = 1
		} else {
			currNum, _ = strconv.Atoi(input[startSlice:ix])
		}
		buf += strings.Repeat(string(r), currNum)
		startSlice = ix + 1
	}
	return buf
}
