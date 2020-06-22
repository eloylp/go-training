package main

import (
	"fmt"
)

func main() {
	m1 := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	m2 := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}

	result := make([][]int, len(m1))

	for m1y := 0; m1y < len(m1); m1y++ {
		result[m1y] = append(make([]int, len(m2)))
		for m2x := 0; m2x < len(m1[0]); m2x++ {
			for m2y := 0; m2y < len(m2); m2y++ {
				result[m1y][m2x] += m1[m1y][m2y] * m2[m2y][m2x]
			}
		}

	}

	//result := make([][]int, len(m1))
	//for m1y := 0; m1y < len(m1); m1y++ {
	//	result[m1y] = make([]int, len(m2[0]))
	//	for m2x := 0; m2x < len(m2[0]); m2x++ {
	//		for m2y := 0; m2y < len(m2); m2y++ {
	//			result[m1y][m2x] += m1[m1y][m2y] * m2[m2y][m2x]
	//		}
	//	}
	//}
	fmt.Println(result)
}
