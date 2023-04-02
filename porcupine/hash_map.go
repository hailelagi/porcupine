package porcupine

import "sync"

type LockingMap struct {
	sync.RWMutex
	Fields map[string]int
}

func (l *LockingMap) Get(key string) int {
	return l.Fields[key]
}

func (l *LockingMap) Put(key string, value int) int {
	l.RWMutex.Lock()
	defer l.Unlock()

	l.Fields[key] = value
	return value
}

func (l *LockingMap) In(key string) bool {
	keys := make([]string, 0, len(l.Fields))
	for k := range l.Fields {
		keys = append(keys, k)
	}

	for _, k := range keys {
		if k == key {
			return true
		}
	}

	return false
}

func (l *LockingMap) Del(key string) {
	l.RWMutex.Lock()
	defer l.Unlock()

	delete(l.Fields, key)
}
