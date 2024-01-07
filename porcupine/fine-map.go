package porcupine

import (
	"sync"
)

// MapSingleMutex uses a single global mutex to protect all operations on the map.
type MapSingleMutex struct {
	Data   map[int]int
	global sync.RWMutex
}

func NewMapSingleMutex() *MapSingleMutex {
	return &MapSingleMutex{
		Data:   make(map[int]int),
		global: sync.RWMutex{},
	}
}

func (m *MapSingleMutex) Increment(key int) {
	m.global.Lock()
	defer m.global.Unlock()

	m.Data[key]++
}

func (m *MapSingleMutex) GetValue(key int) int {
	m.global.RLock()
	defer m.global.RUnlock()

	return m.Data[key]
}

// ConcurrentMap uses multiple local mutexes partitioned by the hash of the key to allow concurrent writes.
type ConcurrentMap struct {
	Data   map[int]int
	locks  []*sync.RWMutex
	global sync.RWMutex
}

// the relationship between the number of locks
// and the size of the underlying
// `bucket index modulo lock array size`
func NewMap(numLocks int) *ConcurrentMap {
	locks := make([]*sync.RWMutex, numLocks)
	for i := range locks {
		locks[i] = &sync.RWMutex{}
	}

	return &ConcurrentMap{
		Data:  make(map[int]int),
		locks: locks,
	}
}

// Writes are much quicker as they can be parallelised as long as the
// hashing function behaves properly and segments the keys
func (m *ConcurrentMap) Increment(key int) {
	lock := m.locks[key%len(m.locks)]
	lock.Lock()
	defer lock.Unlock()

	m.global.Lock()
	defer m.global.Unlock()

	m.Data[key]++
}

// This seems to make reads slightly more expensive
// as you're acquiring/releasing multiple locks
// one for the global RWMutex
// and a second access to the partioned lock
func (m *ConcurrentMap) GetValue(key int) int {
	lock := m.locks[key%len(m.locks)]
	lock.RLock()
	defer lock.RUnlock()

	m.global.RLock()
	defer m.global.RUnlock()

	return m.Data[key]
}
