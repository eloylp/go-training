package quicksort

import (
	"math/rand"
)

func Do(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pi := rand.Int() % len(a)
	a[pi], a[right] = a[right], a[pi]
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	Do(a[:left])
	Do(a[left+1:])
	return a
}

func DoSimple(a []int) []int {
	if len(a) < 2 {
		return a
	}
	pi := rand.Int() % len(a)
	vpi := a[pi]
	var left []int
	var right []int
	for _, v := range a {
		if v < vpi {
			left = append(left, v)
		} else if v > vpi {
			right = append(right, v)
		}
	}
	return append(append(DoSimple(left), vpi), DoSimple(right)...)
}
