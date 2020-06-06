package dectobase

import (
	"strconv"
)

func DecToBase(number, base int) string {
	return convertToBase(number, base, "")
}

func convertToBase(number, base int, reminder string) string {
	dec := int((float64(number / base)))
	mod := number % base
	if dec == 0 {
		return mapForBase(mod, base) + reminder
	}
	return convertToBase(dec, base, mapForBase(mod, base)+reminder)
}

func mapForBase(mod, base int) string {
	if base != 16 {
		return strconv.Itoa(mod)
	}
	res, ok := base16[mod]
	if !ok {
		return strconv.Itoa(mod)
	}
	return res
}

var base16 = map[int]string{
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
}
