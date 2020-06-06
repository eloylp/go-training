package twothatsum

import (
	"math"
)

func TwoThatSumFirstOccurrence(list []int, sum int) (ia, ib int) {
	for k1, v1 := range list {
		for k2, v2 := range list {
			if v1+v2 == sum {
				return k1, k2
			}
		}
	}
	return -1, -1
}

func TwoThatSumLessDifference(list []int, sum int) (ia, ib int) {
	selA := -1
	selB := -1

	for k1, v1 := range list {
		for k2, v2 := range list {
			if v1+v2 == sum && selA == -1 && selB == -1 {
				selA = k1
				selB = k2
				continue
			}
			currentDelta := math.Abs(float64(list[v2]) - float64(list[v1]))
			selectedDelta := math.Abs(float64(list[selB]) - float64(list[selA]))
			if v1+v2 == sum && (currentDelta < selectedDelta) {
				selA = k1
				selB = k2
			}
		}
	}
	return selA, selB
}
