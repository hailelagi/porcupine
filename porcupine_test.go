package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("increment", func(t *testing.T) {
		testMap := &LockingMap{RWMutex: sync.RWMutex{}, fields: make(map[string]int)}

		Handle(testMap, "test", 1)
		Handle(testMap, "test-x", 2)
		Handle(testMap, "test-y", 3)

		assert(t, testMap, "test-x", 2)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		testMap := &LockingMap{RWMutex: sync.RWMutex{}, fields: make(map[string]int)}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(i int) {
				k := fmt.Sprintf("test-%d", i)

				Handle(testMap, k, i)
				wg.Done()
			}(i)
		}
		wg.Wait()

		assert(t, testMap, "test-998", 998)
	})
}

func assert(t testing.TB, result *LockingMap, key string, want int) {
	t.Helper()
	if result.fields[key] != want {
		t.Fail()
	}
}
