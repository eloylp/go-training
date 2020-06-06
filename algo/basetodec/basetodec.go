package basetodec

import (
	"github.com/eloylp/go-training/algo/reverse"
	"math"
)

var numbers = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
}

func BaseToDec(number string, base int) int {
	var total int
	number2 := reverse.ReverseStringRune(number)
	for i, n := range number2 {
		number := numbers[string(n)]
		total += number * int(math.Pow(float64(base), float64(i)))
	}
	return total
}
