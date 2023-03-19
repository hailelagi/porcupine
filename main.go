package main

import (
	"fmt"
	"sync"

	"github.com/hailelagi/porcupine-go/porcupine"
)

func main() {
	ref := &porcupine.LockingMap{
		RWMutex: sync.RWMutex{},
		Fields:  make(map[string]int),
	}

	go porcupine.Handle(ref, "test", 1)
	go porcupine.Handle(ref, "test-x", 2)
	go porcupine.Handle(ref, "test-y", 3)
	go porcupine.Handle(ref, "test-z", 4)
	go porcupine.Handle(ref, "test", 69)

	fmt.Println(ref.Fields)
}
