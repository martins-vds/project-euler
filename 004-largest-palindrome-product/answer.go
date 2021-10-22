package projecteuler

import (
	"math"
	"strconv"
)

func answer(digits int) int {
	upperBound := (int)(math.Pow(10, float64(digits))) - 1
	lowerBound := (int)(math.Pow(10, float64(digits-1))) - 1

	maxPalindrome := 0

	for x := upperBound; x > 0 && x > lowerBound; x-- {
		for y := upperBound; y > 0 && y > lowerBound; y-- {
			prod := x * y
			if isPalindrome(prod) {
				if prod > maxPalindrome {
					maxPalindrome = prod
				}
			}
		}
	}

	return maxPalindrome
}

func isPalindrome(number int) bool {
	t := strconv.Itoa(number)

	return t == reverse(t)
}

func reverse(number string) string {
	var reversed string

	for _, s := range number {
		reversed = string(s) + reversed
	}

	return reversed
}
