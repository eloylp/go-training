package factor

func Factor(primes []int, number int) []int {
	var result []int
	for _, p := range primes {
		for number%p == 0 {
			number = number / p
			result = append(result, p)
		}
	}
	if number > 1 {
		result = append(result, number)
	}
	return result
}
