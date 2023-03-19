package main

import (
	"fmt"
	"sync"
)

type LockingMap struct {
	sync.RWMutex
	fields map[string]int
}

func Handle(l *LockingMap, k string, v int) {
	l.RWMutex.Lock()
	defer l.Unlock()

	l.fields[k] = v
}

func main() {
	ref := &LockingMap{
		RWMutex: sync.RWMutex{},
		fields:  make(map[string]int),
	}

	go Handle(ref, "test", 1)
	go Handle(ref, "test-x", 2)
	go Handle(ref, "test-y", 3)
	go Handle(ref, "test-z", 4)
	go Handle(ref, "test", 69)

	fmt.Println(ref.fields)
}
