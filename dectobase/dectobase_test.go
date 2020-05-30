package dectobase_test

import (
	"github.com/eloylp/go-training/dectobase"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Sample struct {
	Description string
	Decimal     int
	ToBase      int
	Result      string
}

func TestDecToBase(t *testing.T) {
	for _, c := range cases() {
		t.Run(c.Description, func(t *testing.T) {
			result := dectobase.DecToBase(c.Decimal, c.ToBase)
			assert.Equal(t, c.Result, result)
		})
	}
}

func cases() []Sample {
	return []Sample{
		{"Decimal 5 to base 2", 5, 2, "101"},
		{"Decimal 10 to base 2", 10, 2, "1010"},
		{"Decimal 10 to base 8", 10, 8, "12"},
		{"Decimal 10 to base 16", 10, 16, "A"},
		{"Decimal 11 to base 16", 11, 16, "B"},
		{"Decimal 12 to base 16", 12, 16, "C"},
		{"Decimal 13 to base 16", 13, 16, "D"},
		{"Decimal 14 to base 16", 14, 16, "E"},
		{"Decimal 15 to base 16", 15, 16, "F"},
		{"Decimal 175 to base 16", 175, 16, "AF"},
	}
}
