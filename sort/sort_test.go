package sort

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

func TestInsertionSort(t *testing.T) {
	tests := GenerateTests()

	for _, test := range tests {
		InsertionSort(test.Arr, 0, len(test.Arr)-1)
		assert.Equal(t, test.Expected, test.Arr)
	}
}

func TestQuickSort(t *testing.T) {
	tests := GenerateTests()

	for _, test := range tests {
		QuickSort(test.Arr, 0, len(test.Arr)-1)
		assert.Equal(t, test.Expected, test.Arr)
	}
}

func GenerateTests() []Test {
	return []Test{
		{
			Arr:      []int{2, 14, 11, 7, 3, 8, 6, 1, 9},
			Expected: []int{1, 2, 3, 6, 7, 8, 9, 11, 14},
		},
		{
			Arr:      []int{9, 5, 3, 7},
			Expected: []int{3, 5, 7, 9},
		},
		{
			Arr:      []int{5, 4, 6, 3, 7, 1},
			Expected: []int{1, 3, 4, 5, 6, 7},
		},
		{
			Arr:      []int{9, 8, 3, 4, 0, 1, 4, 13},
			Expected: []int{0, 1, 3, 4, 4, 8, 9, 13},
		},
		{
			Arr:      []int{-4, -1, -3, 0, 0, -99, 99},
			Expected: []int{-99, -4, -3, -1, 0, 0, 99},
		},
		{
			Arr:      []int{1, 1, 1, 1, 1, 1, 0},
			Expected: []int{0, 1, 1, 1, 1, 1, 1},
		},
	}
}
