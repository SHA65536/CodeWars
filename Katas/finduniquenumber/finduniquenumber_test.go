package finduniquenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []float32
	Expected float32
}

var Results = []Result{
	{[]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}, 2.0},
	{[]float32{0, 0, 0.55, 0, 0}, 0.55},
	{[]float32{0, 0, 1}, 1},
	{[]float32{0, 1, 0}, 1},
	{[]float32{1, 0, 0}, 1},
}

func TestFindUniqueNumber(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := FindUniq(res.Input)
		assert.Equal(want, got)
	}
}
