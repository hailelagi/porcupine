package porcupine

import (
	"hash/fnv"
)

type Node struct {
	key   string
	value interface{}
	next  *Node
}

type Table struct {
	buckets []*Node
	size    int
}

func NewHashTable(size int) *Table {
	return &Table{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func (ht *Table) Put(key string, value interface{}) {
	// todo
}

func (ht *Table) Get(key string) interface{} {
	// todo
	return nil
}

// hash the key
func hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32())
}
