package isbn10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected bool
}

var Results = []Result{
	/*{"1112223339", true},
	{"048665088X", true},
	{"1293000000", true},
	{"1234554321", true},
	{"1234512345", false},
	{"1293", false},*/
	{"X123456788", false},
	{"ABCDEFGHIJ", false},
	{"XXXXXXXXXX", false},
}

func TestISBN10Validation(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := ValidISBN10(res.Input)
		assert.Equal(want, got, res)
	}
}
