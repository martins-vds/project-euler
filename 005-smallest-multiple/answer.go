package projecteuler

func answer(limit int) int {
	smallestMultiple := 1

	for number := 2; number <= limit; number++ {
		remainder := smallestMultiple % number

		if remainder != 0 {
			if isPrime(number) {
				smallestMultiple *= number
			} else {
				smallestMultiple *= number / remainder
			}
		}
	}

	return smallestMultiple
}

func isPrime(number int) bool {
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
