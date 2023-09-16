package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {

	switch fnb.(type) {
	case FancyNumber:
		v := fnb.Value()
		i, _ := strconv.Atoi(v)
		return i
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	num := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %d.0", num)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) (ans string) {
	switch v := i.(type) {
	case int:
		ans = DescribeNumber(float64(v))
	case float64:
		ans = DescribeNumber(v)
	case NumberBox:
		ans = DescribeNumberBox(v)
	case FancyNumberBox:
		ans = DescribeFancyNumberBox(v)
	default:
		ans = "Return to sender"
	}

	return ans
}
