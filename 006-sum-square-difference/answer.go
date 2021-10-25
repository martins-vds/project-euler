package projecteuler

func answer(limit int) int {
	return squareOfSums(limit) - sumOfSquares(limit)
}

func sumOfSquares(limit int) int {
	sum := 0

	for i := 1; i <= limit; i++ {
		sum += i * i
	}

	return sum
}

func squareOfSums(limit int) int {
	sum := 0

	for i := 1; i <= limit; i++ {
		sum += i
	}

	return sum * sum
}
