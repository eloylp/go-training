package quicksort

import (
	"fmt"
	"math/rand"
)

func Do(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pi := rand.Int() % len(a)
	fmt.Println(pi)
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
