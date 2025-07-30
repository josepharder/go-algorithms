package trie

type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
	Value    string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
		IsEnd:    false,
	}
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	current := t.Root

	for _, ch := range word {
		if _, exists := current.Children[ch]; !exists {
			current.Children[ch] = NewTrieNode()
		}

		current = current.Children[ch]
	}

	current.IsEnd = true
	current.Value = word
}

func (t *Trie) Search(word string) bool {
	current := t.Root

	for _, ch := range word {
		if _, exists := current.Children[ch]; !exists {
			return false
		}

		current = current.Children[ch]
	}

	return current.IsEnd
}

func findLongestCommonPrefix(words []string) string {
	trie := NewTrie()

	for _, word := range words {
		trie.Insert(word)
	}

	var commonPrefix func(node *TrieNode) []rune

	commonPrefix = func(node *TrieNode) []rune {
		if node == nil || len(node.Children) > 1 || node.IsEnd {
			return nil
		}

		for ch, child := range node.Children {
			return append([]rune{ch}, commonPrefix(child)...)
		}

		return nil
	}

	return string(commonPrefix(trie.Root))
}
