package porcupine

import "sync"

type LockingMap struct {
	sync.RWMutex
	fields map[string]int
}

func Handle(l *LockingMap, k string, v int) {
	l.RWMutex.Lock()
	defer l.Unlock()

	l.fields[k] = v
}
