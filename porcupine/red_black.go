package porcupine

import (
	"errors"
	"sync"

	constraints "golang.org/x/exp/constraints"
)

type RedBlack[key constraints.Ordered, value any] struct {
	root *RedBlackNode[key, value]
	sync.RWMutex
}

type RedBlackNode[key constraints.Ordered, value any] struct {
	key   key
	value value
	color bool

	left  *RedBlackNode[key, value]
	right *RedBlackNode[key, value]
}

func NewRedBlack() *RedBlack[int, int] {
	return &RedBlack[int, int]{root: &RedBlackNode[int, int]{}}
}

func (t *RedBlack[K, V]) Get(key K) (V, error) {
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

func (t *RedBlack[K, V]) Put(key K, value V) error {
	t.Lock()
	defer t.Unlock()

	currentNode := t.root

	if t.root == nil {
		t.root = &RedBlackNode[K, V]{key: key, value: value}

		return nil
	}

	for currentNode != nil {
		if currentNode.key == key {
			currentNode.value = value
			return nil
		} else if currentNode.key < key {
			next := currentNode.left
			if next == nil {
				currentNode.left = &RedBlackNode[K, V]{key: key, value: value}
				return nil
			} else {
				currentNode = next
			}
		} else {
			next := currentNode.right

			if next == nil {
				currentNode.right = &RedBlackNode[K, V]{key: key, value: value}
				return nil
			} else {
				currentNode = next
			}
		}
	}

	return nil
}
