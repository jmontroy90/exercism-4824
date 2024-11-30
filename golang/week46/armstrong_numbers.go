package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	ds := strconv.Itoa(n)
	l := len(ds)

	var sum int
	for _, d := range ds {
		sum += int(math.Pow(float64(d-'0'), float64(l)))
	}

	return sum == n
}
