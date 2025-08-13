package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursion(t *testing.T) {
	assert.Equal(t, 15, Sum([]int{1, 2, 3, 4, 5}))
}

func TestLength(t *testing.T) {
	assert.Equal(t, 5, CountLength([]int{1, 2, 3, 4, 5}))
}

func TestFindMax(t *testing.T) {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	assert.Equal(t, 9, FindMax(numbers, numbers[0]))
}
