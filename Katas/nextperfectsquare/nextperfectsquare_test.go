package nextperfectsquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int64
	Expected int64
}

var Results = []Result{
	{121, 144},
	{625, 676},
	{114, -1},
}

func TestNextSquare(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := FindNextSquare(res.Input)
		assert.Equal(want, got)
	}
}
