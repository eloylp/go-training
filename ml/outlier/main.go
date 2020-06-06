package main

import (
	"fmt"
	"math"
)

func main() {
	data := []float64{5, 5.5, 6, 5.3, 7, 5.1, -2, 45}
	limit := float64(0.7)
	result := outliers(data, limit)
	fmt.Println("outliers:", result)
}

func outliers(x []float64, limit float64) []float64 {
	deviation := math.Sqrt(variance(x))
	mean := mean(x)
	anomaly := deviation * limit
	lowerLimit := mean - anomaly
	upperLimit := mean + anomaly
	fmt.Println(lowerLimit, upperLimit)

	outliers := []float64{}
	for _, v := range x {
		if v < lowerLimit || v > upperLimit {
			outliers = append(outliers, v)
		}
	}
	return outliers
}

func variance(x []float64) float64 {
	mean := mean(x)
	sum := float64(0)
	for _, v := range x {
		sum += (v - mean) * (v - mean)
	}
	return sum / float64(len(x))
}

func mean(x []float64) float64 {
	var sum float64
	for _, v := range x {
		sum += v
	}
	return sum / float64(len(x))
}
