package porcupine

import (
	"errors"
	"sync"

	constraints "golang.org/x/exp/constraints"
)

type BST[Key constraints.Ordered, Value any] struct {
	key   Key
	value any

	parent *BST[Key, Value]
	left   *BST[Key, Value]
	right  *BST[Key, Value]

	sync.RWMutex
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

// take the pointer to the root node
func (t *BST[K, V]) Get(key K) (int, error) {
	t.RLock()
	defer t.RUnlock()

	currentNode := t.parent

	for currentNode != nil {
		if currentNode.key == key {
			return currentNode.value.(int), nil
		} else if currentNode.key < key {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	return 0, errors.New("not found")
}

func (l *BST[string, any]) Put(key string, value any) {
	l.RWMutex.Lock()
	defer l.Unlock()
}

func (l *BST[string, any]) In(key string) bool {
	// O(n) in order traversal
	return false
}

// O log(N)
func (l *BST[string, any]) Del(key string) {
	l.RWMutex.Lock()
	defer l.Unlock()
}
