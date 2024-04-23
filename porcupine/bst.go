package porcupine

import (
	"errors"
	"sync"

	constraints "golang.org/x/exp/constraints"
)

type BST[Key constraints.Ordered, Value any] struct {
	root *BSTNode[Key, Value]
	sync.RWMutex
}

type BSTNode[Key constraints.Ordered, Value any] struct {
	key   Key
	value Value

	left  *BSTNode[Key, Value]
	right *BSTNode[Key, Value]
}

// BST Property:
// all nodes left have key < x and right x > key.
func NewBSTree() BST[int, int] {
	return BST[int, int]{
		root: &BSTNode[int, int]{key: 0, value: 0},
	}
}

// take the pointer to the root node
func (t *BST[K, V]) Get(key K) (V, error) {
	t.RLock()
	defer t.RUnlock()

	var empty V
	currentNode := t.root

	for currentNode != nil {
		if currentNode.key == key {
			return currentNode.value, nil
		} else if currentNode.key < key {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	return empty, errors.New("not found")
}

// devolves into a linkedlist
func (t *BST[K, V]) Put(key K, value V) error {
	t.Lock()
	defer t.Unlock()
	var emptyKey K

	currentNode := t.root

	if t.root.key == emptyKey {
		currentNode.key = key
		currentNode.value = value

		return nil
	}

	for currentNode != nil {
		if currentNode.key == key {
			currentNode.value = value
			return nil
		} else if currentNode.key <= key {
			next := currentNode.left
			if next == nil {
				currentNode.left = &BSTNode[K, V]{key: key, value: value}
				return nil
			} else {
				currentNode = next
			}
		} else {
			next := currentNode.right

			if next == nil {
				currentNode.right = &BSTNode[K, V]{key: key, value: value}
				return nil
			} else {
				currentNode = next
			}
		}
	}

	return nil
}
