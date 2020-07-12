package quicksort

import (
	"fmt"
)

func Do(list []int) []int {
	pi := (len(list) - 1) / 2
	p := list[pi]
	fmt.Printf("p: %v \n", p)
	l1 := list[:pi]
	l2 := list[pi+1:]
	fmt.Printf("l1: %v \n", l1)
	fmt.Printf("l2: %v \n", l2)

	sortPart(l1, p)
	sortPart(l2, p)
	result := append(append(l1, p), l2...)
	return result
}

func sortPart(iterList []int, p int) {
	for i, j := 0, len(iterList)-1; i < len(iterList); i, j = i+1, j-1 {
		if iterList[i] > p && iterList[j] < p {
			iterList[i], iterList[j] = iterList[j], iterList[i]
		}
	}
}
