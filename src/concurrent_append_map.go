package porcupine

import (
	"sync"
)

// see: https://github.com/golang/go/issues/21035
// see: https://github.com/golang/go/issues/28938
// see: https://github.com/golang/go/issues/47643
type ConcurrentAppendMap struct {
	Fields *sync.Map
}

func (c *ConcurrentAppendMap) Get(key string) int {
	value, found := c.Fields.Load(key)

	if found && value != nil {
		return value.(int)
	} else {
		return 0
	}
}

func (c *ConcurrentAppendMap) Put(key string, value int) int {
	existing, _ := c.Fields.LoadOrStore(key, value)

	return existing.(int)
}

func (c *ConcurrentAppendMap) In(key string) bool {
	_, found := c.Fields.Load(key)

	return found
}

func (c *ConcurrentAppendMap) Del(key string) bool {
	_, found := c.Fields.LoadAndDelete(key)

	return found
}
