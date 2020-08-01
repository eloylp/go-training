package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

type Number struct {
	ID, Date                           string
	N1, N2, N3, N4, N5, N6, Comp, Rein int8
	Joker                              string
}

func main() {
	numbers := Numbers()
	var data []float64
	for _, n := range numbers {
		data = append(data, float64(n.N1))
	}
	outliers := outliers(data, 0.5)

	m := map[int]int{}
	for _, out := range outliers {
		m[int(out)]++
	}
	fmt.Println(m)
}

func Numbers() []*Number {
	file, err := os.Open("./numbers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	var count int
	var n []*Number
	for {
		rec, err := r.Read()
		if count == 0 {
			count++
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		n = append(n, &Number{
			ID:    rec[0],
			Date:  rec[1],
			N1:    MustParseInt(rec[2]),
			N2:    MustParseInt(rec[3]),
			N3:    MustParseInt(rec[4]),
			N4:    MustParseInt(rec[5]),
			N5:    MustParseInt(rec[6]),
			N6:    MustParseInt(rec[7]),
			Comp:  MustParseInt(rec[8]),
			Rein:  MustParseInt(rec[9]),
			Joker: rec[10],
		})
		count++
	}
	return n
}

func MustParseInt(s string) int8 {
	if s == "" {
		return -1
	}
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		panic(s)
	}
	return int8(i)
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
