package projecteuler

func answer(nthTerm int) int {
	sum := 0

	x, y, z := 1, 2, 0

	for x <= nthTerm {
		if x%2 == 0 {
			sum += x
		}

		z = x + y
		x = y
		y = z
	}

	return sum
}
