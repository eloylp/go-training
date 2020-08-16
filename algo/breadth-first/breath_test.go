package breadth_test

import (
	"github.com/eloylp/go-training/algo/breadth-first"
	"testing"
)

var graph = breadth.Graph{
	"pepe":     []string{"nano", "juana", "paca"},
	"nano":     []string{},
	"juana":    []string{"rakshita"},
	"paca":     []string{"pakis"},
	"rakshita": []string{"pakis"},
	"pakis":    []string{"rakshita"},
}

func TestBreath(t *testing.T) {
	type args struct {
		start  string
		search string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Look for pakis from pepe", args{
				start:  "pepe",
				search: "pakis",
			},
			"paca",
		},
		{
			"Look for pakis from juana", args{
				start:  "juana",
				search: "pakis",
			},
			"rakshita",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breadth.Breath(graph, tt.args.start, tt.args.search); got != tt.want {
				t.Errorf("Breath() = %v, want %v", got, tt.want)
			}
		})
	}
}
