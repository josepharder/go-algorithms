package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {

	items := []int{1, 3, 7, 9, 15, 20, 25, 30, 32}

	tests := []struct {
		item     int
		expected int
		error    string
	}{
		{9, 3, ""},
		{32, 8, ""},
		{1, 0, ""},
		{100, 0, "not found"},
	}

	for _, test := range tests {
		index, err := BinarySearch(items, test.item, 0, len(items)-1)

		if err != nil {
			assert.Error(t, err)
			return
		}

		assert.Equal(t, test.expected, *index)
	}
}
