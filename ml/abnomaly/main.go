package main

import (
	"fmt"
	"github.com/lytics/anomalyzer"
)

func main() {
	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: 0.1,
		UpperBound:  5,
		LowerBound:  anomalyzer.NA,
		ActiveSize:  1,
		NSeasons:    4,
		Methods:     []string{"diff", "fence", "magnitude", "ks"},
	}
	data := []float64{0.3, 0.2, 4, 5, 3, 5, 10}
	anom, _ := anomalyzer.NewAnomalyzer(conf, data)
	prob := anom.Push(3)
	fmt.Println("Anomalous probability:", prob)
}
