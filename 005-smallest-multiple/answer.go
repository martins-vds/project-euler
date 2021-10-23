package projecteuler

import "fmt"

func answer(limit int) int {
	smallestMultiple := 1 * 2 * 3

	for number := 4; number <= limit; number++ {
		remainder := smallestMultiple % number
		if remainder != 0 {
			fmt.Print(number)
			fmt.Print(" ")
			fmt.Print(smallestMultiple)
			fmt.Print(" ")
			fmt.Print(smallestMultiple / number)
			fmt.Print(" ")
			fmt.Println(remainder)

			if number%2 != 0 {
				smallestMultiple = (smallestMultiple * number)
			} else {
				smallestMultiple = (smallestMultiple * remainder)
			}
		}
	}

	return smallestMultiple
}
