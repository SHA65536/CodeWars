package sumofnumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	A, B     int
	Expected int
}

var Results = []Result{
	{0, 1, 1},
	{1, 2, 3},
	{5, -1, 14},
	{505, 4, 127759},
	{321, 123, 44178},
	{0, -1, -1},
	{-50, 0, -1275},
	{-1, -5, -15},
	{-5, -5, -5},
	{-505, 4, -127755},
	{-321, 123, -44055},
	{0, 0, 0},
	{-5, -1, -15},
	{5, 1, 15},
	{-17, -17, -17},
	{17, 17, 17},
}

func TestSumOfNumbers(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := GetSum(res.A, res.B)
		assert.Equal(want, got)
	}
}
