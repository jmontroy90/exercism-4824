package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var res strings.Builder
	if number%3 == 0 {
		res.WriteString("Pling")
	}
	if number%5 == 0 {
		res.WriteString("Plang")
	}
	if number%7 == 0 {
		res.WriteString("Plong")
	}
	if res.Len() == 0 {
		res.WriteString(strconv.Itoa(number))
	}
	return res.String()
}
