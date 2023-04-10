package porcupine

import (
	"errors"
	"sync"
)

type LockingMap[K string, V any] struct {
	sync.RWMutex
	Fields map[K]V
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
	// keys := make([]string, 0, len(l.Fields))
	// for k := range l.Fields {
	// 	keys = append(keys, k)
	// }

	// for _, k := range keys {
	// 	if k == key {
	// 		return true
	// 	}
	// }
	_, found := l.Fields[key]

	return found
}

func (l *LockingMap[string, any]) Del(key string) {
	l.RWMutex.Lock()
	defer l.Unlock()

	delete(l.Fields, key)
}
