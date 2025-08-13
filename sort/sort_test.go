package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Arr      []int
	Expected []int
}

func TestBubbleSort(t *testing.T) {
	tests := GenerateTests()

	for _, test := range tests {
		assert.Equal(t, test.Expected, BubbleSort(test.Arr))

	}
}

func TestMergeSort(t *testing.T) {
	tests := GenerateTests()

	for _, test := range tests {
		assert.Equal(t, test.Expected, MergeSort(test.Arr))

	}
}

func GenerateTests() []Test {
	return []Test{
		{
			Arr:      []int{9, 8, 3, 4, 0, 1, 4, 13},
			Expected: []int{0, 1, 3, 4, 4, 8, 9, 13},
		},
		{
			Arr:      []int{-4, -1, -3, 0, 0, -99, 99},
			Expected: []int{-99, -4, -3, -1, 0, 0, 99},
		},
	}
}
