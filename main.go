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

	defer l.current.RUnlock()
}

func main() {

	shareMe := LockingMap{
		RWMutex: sync.RWMutex{},
		fields:  make(map[string]int),
	}

	fmt.Println(shareMe)
}
