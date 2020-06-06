package basetobase_test

import (
	"github.com/eloylp/go-training/algo/basetobase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseToBase(t *testing.T) {
	result := basetobase.BaseToBase("E", 16, 2)
	assert.Equal(t, "1110", result)
}
