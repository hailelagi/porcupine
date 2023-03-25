package porcupine

type AVLTree struct {
	name string
	node *AVLTree
}

func (*AVLTree) Get(key string) int {
	return 0
}

func (a *AVLTree) Put(key string, value int) int {
	return 0
}

func (a *AVLTree) In(key string) bool {
	return false
}

func (a *AVLTree) Del(key string) {
}

func rebalance(*AVLTree) {
}
