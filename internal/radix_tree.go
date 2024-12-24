package porcupine

/*
TODO:
most radix trees require to trade off tree height versus space efficiency
by setting a globally valid fanout parameter ??

TODO: implement the path collapsing/compression optimisation
 Morrison introduced path compression in order to store long strings
efficiently [16]. Knuth [17] analyzes these early trie variants
*/

type RadixTree struct {
	root *RadixNode
}

type RadixNode struct {
	children map[rune]*RadixNode
	isEnd    bool
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
