package porcupine

import (
	"github.com/lrita/cmap"
)

// Cmap is a drop-ish in replacement for sync.Map
type ConcurrentMap struct {
	Fields *cmap.Cmap
}

func (c ConcurrentMap) Get(key string) int {
	value, found := c.Fields.Load(key)

	if found && value != nil {
		return value.(int)
	} else {
		return 0
	}
}

func (c ConcurrentMap) Put(key string, value int) int {
	existing, _ := c.Fields.LoadOrStore(key, value)

	return existing.(int)
}

func (c ConcurrentMap) In(key string) bool {
	_, found := c.Fields.Load(key)

	return found
}

func (c ConcurrentMap) Del(key string) bool {
	_, found := c.Fields.LoadAndDelete(key)

	return found
}
