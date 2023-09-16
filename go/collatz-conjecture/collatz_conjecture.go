package collatzconjecture

import (
	"errors"
)

// CollatzConjecture returns the number of steps the Collatz sequence needs to reach 1.
// An error is returned if the input is not a positive integer.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("Invalid input: input is negative.")
	}
	return collatz(n), nil
}

func collatz(n int) (steps int) {
	if n == 1 {
		return 0
	}

	if isEven(n) {
		return 1 + collatz(n/2)
	}

	return 1 + collatz(3*n+1)

}

func isEven(n int) bool {
	return n%2 == 0
}
