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
	key     key
	value   value
	left    *AVLNode[key, value]
	right   *AVLNode[key, value]
	balance int
	height  int
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

	t.root = t.insert(t.root, key, value)
	return nil
}

func (t *AVL[K, V]) insert(node *AVLNode[K, V], key K, value V) *AVLNode[K, V] {
	if node == nil {
		return &AVLNode[K, V]{key: key, value: value, height: 1}
	}

	if key < node.key {
		node.left = t.insert(node.left, key, value)
	} else if key > node.key {
		node.right = t.insert(node.right, key, value)
	} else {
		node.value = value
		return node
	}

	node.height = 1 + max(maxHeight(node.left), maxHeight(node.right))
	node.balance = maxHeight(node.left) - maxHeight(node.right)

	return t.rebalance(node)
}

func (t *AVL[K, V]) rebalance(node *AVLNode[K, V]) *AVLNode[K, V] {
	if node.balance < -1 {
		if node.right != nil && node.right.balance > 0 {
			node.right = t.rotateRight(node.right)
		}
		return t.rotateLeft(node)
	} else if node.balance > 1 {
		if node.left != nil && node.left.balance < 0 {
			node.left = t.rotateLeft(node.left)
		}
		return t.rotateRight(node)
	}
	return node
}

func (t *AVL[K, V]) rotateLeft(node *AVLNode[K, V]) *AVLNode[K, V] {
	rightNode := node.right
	node.right = rightNode.left
	rightNode.left = node
	node.balance = maxHeight(node.left) - maxHeight(node.right)
	rightNode.balance = maxHeight(rightNode.left) - maxHeight(rightNode.right)
	return rightNode
}

func (t *AVL[K, V]) rotateRight(node *AVLNode[K, V]) *AVLNode[K, V] {
	leftNode := node.left
	node.left = leftNode.right
	leftNode.right = node
	node.balance = maxHeight(node.left) - maxHeight(node.right)
	leftNode.balance = maxHeight(leftNode.left) - maxHeight(leftNode.right)
	return leftNode
}

func maxHeight[K constraints.Ordered, V any](node *AVLNode[K, V]) int {
	if node == nil {
		return 0
	}
	leftHeight := maxHeight(node.left)
	rightHeight := maxHeight(node.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
