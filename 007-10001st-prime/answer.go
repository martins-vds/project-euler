package projecteuler

func answer(nthTerm int) int {

	numbers := primeNumbers(nthTerm * nthTerm)
	nthPrimeCount := 0
	nthPrime := 0

	for i := 2; i < len(numbers) && nthPrimeCount < nthTerm; i++ {
		if numbers[i] {
			nthPrimeCount++
			nthPrime = i
		}
	}

	return nthPrime
}

func primeNumbers(limit int) []bool {
	numbers := make([]bool, limit)

	for k := 2; k < limit; k++ {
		numbers[k] = true
	}

	for k := 2; k*k < limit; k++ {
		if numbers[k] {
			for l := k * k; l < limit; l += k {
				numbers[l] = false
			}
		}
	}

	return numbers
}
