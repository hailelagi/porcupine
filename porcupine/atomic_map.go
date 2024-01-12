package porcupine

import (
	"hash/fnv"
)

type Node struct {
	key   string
	value interface{}
	next  *Node
}

type HashTable struct {
	buckets []*Node
	size    int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func (ht *HashTable) Put(key string, value interface{}) {
	// todo
}

func (ht *HashTable) Get(key string) interface{} {
	// todo
	return nil
}

// hash the key
func hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32())
}
