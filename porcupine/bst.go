package porcupine

import "sync"

type BST[Key comparable, Value any] struct {
	key   Key
	value any

	parent *BST[Key, Value]
	left   *BST[Key, Value]
	right  *BST[Key, Value]

	mux sync.RWMutex
}

// BST Property:
// all nodes left have key < x and right x > key.
func newSearchTree() BST[int, any] {
	return BST[int, any]{
		parent: nil,
		left:   nil,
		right:  nil,
	}
}
