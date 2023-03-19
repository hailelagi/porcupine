package porcupine

import "sync"

type LockingMap struct {
	sync.RWMutex
	Fields map[string]int
}

func Handle(l *LockingMap, k string, v int) {
	l.RWMutex.Lock()
	defer l.Unlock()

	l.Fields[k] = v
}
