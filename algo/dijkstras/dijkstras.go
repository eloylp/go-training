package dijkstras

import (
	"math"
)

type Graph map[string]map[string]float64
type Costs map[string]float64
type Parents map[string]string
type Processed map[string]bool

func Dijkstras(graph Graph, costs Costs, parents Parents) (Parents, Costs) {
	processed := Processed{}
	node := LowestNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbours := graph[node]
		for k, _ := range neighbours {
			newCost := cost + neighbours[k]
			if costs[k] > newCost {
				costs[k] = newCost
				parents[k] = node

			}
		}
		processed[node] = true
		node = LowestNode(costs, processed)
	}
	return parents, costs
}

func LowestNode(costs Costs, processed Processed) string {
	var lowestCostNode string
	lowestCost := math.Inf(1)
	for node, cost := range costs {
		_, ok := processed[node]
		if cost < lowestCost && !ok {
			lowestCostNode = node
			lowestCost = cost
		}
	}
	return lowestCostNode
}
