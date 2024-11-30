// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate is a lil baby state machine.
func Abbreviate(s string) string {
	var sb strings.Builder
	newWord := true
	for _, r := range s {
		if newWord && isLetter(r) {
			sb.WriteRune(toUpper(r))
			newWord = false
		}
		if r == ' ' || r == '-' {
			newWord = true
		}
	}
	return sb.String()
}

func toUpper(r rune) rune {
	if isUpper(r) {
		return r
	}
	// assume lowercase
	return r - 32 // 32 code points between an upper and lower
}

func isUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func isLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func isLetter(r rune) bool {
	return isUpper(r) || isLower(r)
}
