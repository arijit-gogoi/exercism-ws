package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood divides the food among the cows.
func DivideFood(fc FodderCalculator, cows int) (foodPerCow float64, err error) {
	amount, err := fc.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}

	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	foodPerCow = amount * factor / float64(cows)
	return foodPerCow, err
}

// ValidateInputAndDivideFood validates input and call DivideFood.
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {

	if cows > 0 {
		return DivideFood(fc, cows)
	}

	return 0.0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// ValidateNumberOfCows validates the number of cows.
func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	case cows == 0:
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	default:
		return nil
	}
}
