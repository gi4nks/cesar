package internal

import (
	"errors"
)

// Converter represents a converter for integer ranges to Roman numerals.
// It defines the start and end values of the range to be converted.
type Converter struct {
	start int // Starting value of the range
	end   int // Ending value of the range
}

// NewConverter creates a new Converter instance with the specified range.
// It returns an error if the range is invalid.
func NewConverter(s int, e int) (*Converter, error) {
	// Validate the range.
	if s < 1 || s > 3999 || e < 1 || e > 3999 {
		return nil, errors.New("start and end values must be between 1 and 3999")
	}

	if s > e {
		return nil, errors.New("start value must be smaller than or equal to end value")
	}

	// Create and return a new Converter instance.
	return &Converter{start: s, end: e}, nil
}

// Convert converts the range from start to end (inclusive) to Roman numerals.
// It iterates over each integer value within the range and converts it to its
// equivalent Roman numeral representation using the toRoman method.
func (c *Converter) Convert() map[int]string {
	result := make(map[int]string)
	for i := c.start; i <= c.end; i++ {
		result[i] = c.toRoman(i)
	}
	return result
}

// toRoman converts the given integer to its equivalent Roman numeral representation.
// It uses a greedy algorithm to subtract the largest possible value from the input number
// and appends the corresponding Roman numeral symbol to the result string.
func (c *Converter) toRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	letters := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; num > 0 && i < len(nums); i++ {
		for num >= nums[i] {
			result += letters[i]
			num -= nums[i]
		}
	}

	return result
}
