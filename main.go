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
	// l.current.RWMutex.RLock()
	// mutate stuff
	// this can and does panic due to "test"=1 and "test"= 69
	// accessing the same memory space must use a "write" lock

	l.current.RWMutex.Lock()
	l.current.fields[k] = v

	defer l.current.Unlock()
}

func main() {
	shareMe := &LockingMap{
		RWMutex: sync.RWMutex{},
		fields:  make(map[string]int),
	}

	x := Resource{current: shareMe}
	ref := &x

	go handle(ref, "test", 1)
	go handle(ref, "test-x", 2)
	go handle(ref, "test-y", 3)
	go handle(ref, "test-z", 4)
	go handle(ref, "test", 69)

	fmt.Println(shareMe.fields)
}
