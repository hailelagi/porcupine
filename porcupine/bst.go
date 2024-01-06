package porcupine

type BST[K string, V any] struct {
	Node map[K]V
}
