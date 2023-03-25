package porcupine

import "sync"

type LockingMap struct {
	sync.RWMutex
	Fields map[string]int
}

func (l *LockingMap) Get(key string) int {
	return 0
}

func (l *LockingMap) Put(key string, value int) int {
	l.RWMutex.Lock()
	defer l.Unlock()

	l.Fields[key] = value
	return value
}

func (l *LockingMap) In(key string) bool {
	return false
}

func (l *LockingMap) Del(key string) {
}
