package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursion(t *testing.T) {
	result := Sum([]int{1, 2, 3, 4, 5})

	assert.Equal(t, 15, result)
}

func TestLength(t *testing.T) {
	lenght := CountLength([]int{1, 2, 3, 4, 5})

	assert.Equal(t, 5, lenght)
}

func TestFindMax(t *testing.T) {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	max := FindMax(numbers, numbers[0])

	assert.Equal(t, 9, max)
}
