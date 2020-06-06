package main

import (
	"fmt"

	"github.com/mash/gokmeans"
)

func main() {
	observations := []gokmeans.Node{
		{4}, {5}, {6}, {8}, {10}, {12}, {15}, {0}, {-1},
	}
	k := 2 // number of clusters
	if success, centroids := gokmeans.Train(observations, k, 50); success {
		fmt.Println("Centroids are:")
		for _, centroid := range centroids {
			fmt.Println(centroid)
		}
		fmt.Println("Clusters are:")
		for _, observation := range observations {
			index := gokmeans.Nearest(observation, centroids)
			fmt.Println("Observation", observation, "belongs in cluster", index+1)
		}
	}
}
