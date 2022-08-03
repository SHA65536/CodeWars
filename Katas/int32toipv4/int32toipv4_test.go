package int32toipv4

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    uint32
	Expected string
}

var Results = []Result{
	{2154959208, "128.114.17.104"},
	{2149583361, "128.32.10.1"},
	{0, "0.0.0.0"},
	{math.MaxUint32, "255.255.255.255"},
}

func TestIntToIPv4(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := Int32ToIp(res.Input)
		assert.Equal(want, got)
	}
}
