package highestscoreword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"man i need a taxi up to ubud", "taxi"},
	{"what time are we climbing up the volcano", "volcano"},
	{"take me to semynak", "semynak"},
	{"aa b", "aa"},
	{"b aa", "b"},
	{"bb d", "bb"},
	{"d bb", "d"},
	{"aaa b", "aaa"},
}

func TestHighestScoringWord(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := High(res.Input)
		assert.Equal(want, got)
	}
}
