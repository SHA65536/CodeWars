package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Target   int
	Expected [2]int
}

var Results = []Result{
	{[]int{2, 7, 11, 15}, 9, [2]int{0, 1}},
	{[]int{3, 2, 4}, 6, [2]int{1, 2}},
	{[]int{3, 3}, 6, [2]int{0, 1}},
	{[]int{1, 2, 3, 4, 5, 6, 7}, 11, [2]int{4, 5}},
	{[]int{7, 6, 5, 4, 3, 2, 1}, 11, [2]int{1, 2}},
}

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := TwoSum(res.Input, res.Target)
		assert.Equal(want, got)
	}
}
