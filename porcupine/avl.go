package porcupine

import (
	"errors"
	"sync"

	constraints "golang.org/x/exp/constraints"
)

type AVL[key constraints.Ordered, value any] struct {
	root *AVLNode[key, value]
	sync.RWMutex
}

type AVLNode[key constraints.Ordered, value any] struct {
	key   key
	value value
	left  *AVLNode[key, value]
	right *AVLNode[key, value]

	// balance = h(left) - h(right)
	balance int
}

func NewAVL() *AVL[int, int] {
	return &AVL[int, int]{root: &AVLNode[int, int]{}}
}

func (t *AVL[K, V]) Get(key K) (V, error) {
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

func (t *AVL[K, V]) Put(key K, value V) error {
	t.Lock()
	defer t.Unlock()

	currentNode := t.root

	if t.root == nil {
		t.root = &AVLNode[K, V]{key: key, value: value}

		return nil
	}

	for currentNode != nil {
		if currentNode.key == key {
			currentNode.value = value
			return nil
		} else if currentNode.key < key {
			next := currentNode.left
			if next == nil {
				currentNode.left = &AVLNode[K, V]{key: key, value: value}
				return nil
			} else {
				currentNode = next
			}
		} else {
			next := currentNode.right

			if next == nil {
				currentNode.right = &AVLNode[K, V]{key: key, value: value}
				return nil
			} else {
				currentNode = next
			}
		}
	}

	return nil
}

func rebalance() {
	// rebalance -1, 0, 1
}
