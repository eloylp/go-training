package reverse

import (
	"strings"
)

func Sum(input []int) int {
	var res int
	for _, v := range input {
		res += v
	}
	return res
}

func StringPreparedSlice(input string) string {
	length := len(input)
	orderedParts := make([]string, length)
	for i, j := length-1, 0; i >= 0; i, j = i-1, j+1 {
		orderedParts[j] = string(input[i])
	}
	return strings.Join(orderedParts, "")
}

func StringRune(input string) string {
	var res string
	for _, v := range input {
		res = string(v) + res
	}
	return res
}
