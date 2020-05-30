package basetodec_test

import (
	"github.com/eloylp/go-training/basetodec"
	"github.com/stretchr/testify/assert"
	"testing"
)

type sample struct {
	Description string
	Number      string
	Base        int
	Expected    int
}

func TestBaseToDec(t *testing.T) {
	for _, c := range cases() {
		t.Run(c.Description, func(t *testing.T) {
			result := basetodec.BaseToDec(c.Number, c.Base)
			assert.Equal(t, c.Expected, result)
		})
	}
}

func cases() []sample {

	return []sample{
		{"Base 2 value 5", "101", 2, 5},
		{"Base 2 value 11", "1011", 2, 11},
		{"Base 16 value 172", "AF", 16, 175},
		{"Base 16 value 126", "7E", 16, 126},
	}
}
