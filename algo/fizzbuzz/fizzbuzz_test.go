package fizzbuzz_test

import (
	"github.com/eloylp/go-training/algo/fizzbuzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	input := 15
	expected := "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz"
	result := fizzbuzz.FizzBuzz(input)
	assert.Equal(t, expected, result)
}
