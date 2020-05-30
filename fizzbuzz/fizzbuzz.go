package fizzbuzz

import (
	"strconv"
	"strings"
)

func FizzBuzz(n int) string {
	result := []string{}
	for i := 1; i <= n; i++ {
		isFizz := i%3 == 0
		isBuzz := i%5 == 0

		if isFizz && isBuzz {
			result = append(result, "FizzBuzz")
			continue
		}
		if isFizz {
			result = append(result, "Fizz")
			continue
		}
		if isBuzz {
			result = append(result, "Buzz")
			continue
		}
		result = append(result, strconv.Itoa(i))
	}
	return strings.Join(result, ", ")

}
