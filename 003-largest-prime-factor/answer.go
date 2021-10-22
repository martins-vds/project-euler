package projecteuler

var factors map[int]int

func answer(number int) int {
	factors = make(map[int]int)

	factor := 2
	temp := number
	for temp > 1 {
		if temp%factor == 0 {
			temp = temp / factor
			factors[factor]++
		} else {
			factor++
		}
	}

	return maxFactor()
}

func shouldIgnore(number int) bool {
	for factor := range factors {
		if number%factor == 0 {
			return true
		}
	}

	return false
}

func maxFactor() int {
	max := 0

	for factor := range factors {
		if factor > max {
			max = factor
		}
	}

	return max
}
