package porcupine

type RadixNode struct {
	children map[rune]*RadixNode
	isEnd    bool
}

type RadixTree struct {
	root *RadixNode
}

func NewRadixTree() *RadixTree {
	return &RadixTree{
		root: &RadixNode{children: make(map[rune]*RadixNode)},
	}
}

func (t *RadixTree) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &RadixNode{children: make(map[rune]*RadixNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *RadixTree) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func (t *RadixTree) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return true
}
