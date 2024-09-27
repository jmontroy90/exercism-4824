package week2

import (
	"unicode/utf8"
)

func Reverse(input string) string {
	inLen := utf8.RuneCountInString(input) // just using len will count bytes - bad for UTF-8 multibyte runes.
	rev := make([]rune, inLen)
	// converting our input explicitly to runes means both our element AND index iterator tracks runes, not bytes.
	// iterating over the string yields a rune element BUT a byte-counting index.
	for ix, l := range []rune(input) {
		rev[inLen-ix-1] = l
	}
	return string(rev)
}

func ReverseAwkward(input string) string {
	//func Reverse(input string) string {
	inLen := int32(utf8.RuneCountInString(input))
	rev := make([]int32, inLen)
	//fmt.Println(inLen, input)
	var runeIx int32 // this is annoying, but the range index counts by bytes, not runes
	for _, l := range input {
		//fmt.Printf("%d, %v\n", runeIx, l)
		rev[inLen-runeIx-1] = l
		runeIx++
	}
	return string(rev)
}
