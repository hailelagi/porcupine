package porcupine

import (
	"errors"
	"sync"
)

type LockingMap[K string, V any] struct {
	sync.RWMutex
	Fields map[K]V
}

func NewLockingMap() *LockingMap[string, int] {
	return &LockingMap[string, int]{Fields: make(map[string]int)}
}

func (l *LockingMap[string, any]) Get(key string) (any, error) {
	l.RLock()
	defer l.RUnlock()

	value, found := l.Fields[key]

	if !found {
		return value, errors.New("not found")
	}

	return value, nil
}

func (l *LockingMap[string, any]) Put(key string, value any) {
	l.RWMutex.Lock()
	defer l.Unlock()

	l.Fields[key] = value
}

func (l *LockingMap[string, any]) In(key string) bool {
	_, found := l.Fields[key]

	return found
}

func (l *LockingMap[string, any]) Del(key string) {
	l.RWMutex.Lock()
	defer l.Unlock()

	delete(l.Fields, key)
}
