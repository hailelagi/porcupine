package porcupine

type RedBlackTree struct {
	root *RBNode
}

type RBNode struct {
	color bool
	node  *RedBlackTree
}
