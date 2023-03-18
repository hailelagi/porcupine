package main

import (
	"fmt"
	"sync"
)

type LockingMap struct {
	sync.RWMutex
	fields map[string]int
}

type Resource struct {
	current *LockingMap
}

func handle(l *Resource, k string, v int) {
	l.current.RWMutex.RLock()

	// mutate stuff
	l.current.fields[k] = v

	defer l.current.RUnlock()
}

func main() {
	shareMe := &LockingMap{
		RWMutex: sync.RWMutex{},
		fields:  make(map[string]int),
	}

	x := Resource{current: shareMe}
	ref := &x

	go func() {
		handle(ref, "test", 1)
		handle(ref, "test-x", 2)
		handle(ref, "test-y", 3)
		handle(ref, "test-z", 4)
	}()

	fmt.Println(shareMe.fields)
}
