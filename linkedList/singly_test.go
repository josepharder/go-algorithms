package linkedList

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppeding(t *testing.T) {
	tests := []struct {
		nodes []int
	}{
		{
			[]int{0, 17},
		}, {
			[]int{},
		}, {
			[]int{-1},
		},
	}

	for _, test := range tests {
		var list LinkedList

		for _, value := range test.nodes {
			list.Append(value)
		}

		assert.Equal(t, len(test.nodes), list.Size)
	}
}

func TestRemovingNodeByValue(t *testing.T) {
	tests := []struct {
		nodes         []int
		nodesToRemove []int
		expected      int
	}{
		{
			[]int{0, 17},
			[]int{0},
			1},
		{
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{-3, -1, 0, 1},
			[]int{-3, 0},
			2,
		},
		{
			[]int{-3, -1, 0, 0, 1},
			[]int{-3, -1, 0, 0, 1},
			0,
		},
	}

	for _, test := range tests {
		var list LinkedList

		for _, value := range test.nodes {
			list.Append(value)
		}

		for _, value := range test.nodesToRemove {
			if err := list.RemoveNodeByValue(value); err != nil {
				log.Println(err)
			}
		}

		assert.Equal(t, test.expected, list.Size)
	}
}

func TestRemovingNodeByPosition(t *testing.T) {
	tests := []struct {
		nodes             []int
		positionsToRemove []int
		expected          int
	}{
		{
			[]int{0, 17},
			[]int{0, 0},
			0,
		},
		{
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{-3, -1, 0, 1},
			[]int{0, 2},
			2,
		},
		{
			[]int{-3, -1, 0, 0, 1},
			[]int{4, 3, 2, 1, 0},
			0,
		},
	}

	for _, test := range tests {
		var list LinkedList

		for _, value := range test.nodes {
			list.Append(value)
		}

		for _, position := range test.positionsToRemove {
			if err := list.RemoveNodeByPosition(position); err != nil {
				log.Println(err)
			}
		}

		assert.Equal(t, test.expected, list.Size)
	}
}
