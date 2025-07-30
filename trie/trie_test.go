package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieStructure(t *testing.T) {
	prefix := findLongestCommonPrefix([]string{"fa", "fe"})

	assert.Equal(t, "f", prefix)
}
