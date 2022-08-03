package wheremyanagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Target   []string
	Expected []string
}

var Results = []Result{
	{"abba", []string{"aabb", "abcd", "bbaa", "dada"}, []string{"aabb", "bbaa"}},
	{"laser", []string{"lazing", "lazy", "lacer"}, []string{}},
}

func TestAnagrams(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := Anagrams(res.Input, res.Target)
		assert.Equal(want, got)
	}
}
