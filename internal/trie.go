package porcupine

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// insert into the tree
func (t *Trie) Insert(word string) {
	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		current = current.children[ch]
	}
	current.isEnd = true
}

// exact substring match
func (t *Trie) Search(word string) bool {
	current := t.root
	for _, ch := range word {
		node, exists := current.children[ch]
		if !exists {
			return false
		}
		current = node
	}
	return current.isEnd
}

// This is the prefix search
func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for _, ch := range prefix {
		node, exists := current.children[ch]
		if !exists {
			return false
		}
		current = node
	}
	return true
}
