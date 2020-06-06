package gdc

import (
	"sort"
)

func GDC(numbers []int) int {
	sort.Ints(numbers)
	divisors := [][]int{}
	for _, n := range numbers {
		d2 := []int{}
		for d := 1; d <= n; d++ {
			if n%d == 0 {
				d2 = append(d2, d)
			}
		}
		divisors = append(divisors, d2)
	}
	dvs := []int{}
	for i := 0; i <= numbers[len(numbers)-1]; i++ {
		present := true
		for _, dr := range divisors {
			if !isPresent(i, dr) {
				present = false
			}
		}
		if present {
			dvs = append(dvs, i)
		}
	}
	if len(dvs) == 0 {
		return 1
	}
	sort.Ints(dvs)
	return dvs[len(dvs)-1]
}

func isPresent(i int, dr1 []int) bool {
	for _, v := range dr1 {
		if i == v {
			return true
		}
	}
	return false
}

func EuclideanGDC(a int, b int) int {
	if b == 0 {
		return a
	}
	a, b = b, a%b
	return EuclideanGDC(a, b)
}
