package porcupine

import (
	"sync"
)

type BPlusTree struct {
	root *BNode
	sync.RWMutex
}

type BNode struct {
	key   int
	value []byte

	left  *BPlusTree
	right *BPlusTree
}

func NewBPlusTree(filename string) (*BPlusTree, error) {
	return &BPlusTree{root: &BNode{}}, nil
}

func (t *BPlusTree) Insert(key int, value int) error {
	return nil
}

func (t *BPlusTree) Search(key int) (int, error) {
	return 0, nil
}
