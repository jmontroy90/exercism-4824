package collatzconjecture

import (
	"errors"
)

var (
	ErrMustBePositive = errors.New("input must be greater than zero")
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, ErrMustBePositive
	}
	var steps int
	for ; n != 1; steps++ { // this is a little cute
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return steps, nil
}
