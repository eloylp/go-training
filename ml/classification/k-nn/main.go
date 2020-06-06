package main

import (
	"fmt"
	"log"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	rawData, err := base.ParseCSVToInstances("data.csv", false)
	if err != nil {
		log.Fatal(err)
	}
	k := 2
	train, test := base.InstancesTrainTestSplit(rawData, 0.70)
	cls := knn.NewKnnClassifier("euclidean", "linear", k)
	if err := cls.Fit(train); err != nil {
		log.Fatal(err)
	}
	p, err := cls.Predict(test)
	if err != nil {
		log.Fatal(err)
	}
	confusionMat, err := evaluation.GetConfusionMatrix(test, p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}
