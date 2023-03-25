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

	go ref.Put("test", 1)
	go ref.Put("test-x", 2)
	go ref.Put("test-y", 3)
	go ref.Put("test-z", 4)
	go ref.Put("test", 69)

	fmt.Println(ref.Fields)
}
