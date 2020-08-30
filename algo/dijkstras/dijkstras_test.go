package dijkstras_test

import (
	"github.com/bmizerany/assert"
	"github.com/eloylp/go-training/algo/dijkstras"
	"math"
	"testing"
)

var (
	graph = dijkstras.Graph{
		"start": {"A": 5, "B": 2},
		"A":     {"C": 4, "D": 2},
		"B":     {"A": 8, "D": 7},
		"C":     {"D": 6, "fin": 3},
		"D":     {"fin": 1},
		"fin":   {},
	}
	// unknown
	costs = dijkstras.Costs{
		"A":   5,
		"B":   2,
		"C":   math.Inf(1),
		"D":   math.Inf(1),
		"fin": math.Inf(1),
	}
	expectedCosts = dijkstras.Costs{
		"A":   5,
		"B":   2,
		"C":   9,
		"D":   7,
		"fin": 8,
	}
	// unknown
	parents = dijkstras.Parents{
		"A":   "start",
		"B":   "start",
		"C":   "",
		"D":   "",
		"fin": "",
	}
	expectedParents = dijkstras.Parents{
		"A":   "start",
		"B":   "start",
		"C":   "A",
		"D":   "A",
		"fin": "D",
	}
)

func TestDijkstras(t *testing.T) {
	parents, costs := dijkstras.Dijkstras(graph, costs, parents)
	assert.Equal(t, expectedParents, parents)
	assert.Equal(t, expectedCosts, costs)
}

func TestLowestNode(t *testing.T) {
	processed := dijkstras.Processed{"A": true}
	costs := dijkstras.Costs{"A": 2, "B": 3, "C": 4}
	assert.Equal(t, "B", dijkstras.LowestNode(costs, processed))
}

func TestLowestNode_notfound(t *testing.T) {
	// setup all as processed
	processed := dijkstras.Processed{"A": true}
	processed["B"] = true
	processed["C"] = true
	costs := dijkstras.Costs{"A": 2, "B": 3, "C": 4}
	assert.Equal(t, "", dijkstras.LowestNode(costs, processed))
}
