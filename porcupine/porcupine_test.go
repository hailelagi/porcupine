package porcupine

import (
	"fmt"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("increment", func(t *testing.T) {
		testMap := &LockingMap[string, int]{RWMutex: sync.RWMutex{}, Fields: make(map[string]int)}

		testMap.Put("test", 1)
		testMap.Put("test-x", 2)
		testMap.Put("test-y", 3)

		assert(t, testMap, "test", 1)
		assert(t, testMap, "test-x", 2)
		assert(t, testMap, "test-y", 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		testMap := &LockingMap[string, int]{RWMutex: sync.RWMutex{}, Fields: make(map[string]int)}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(i int) {
				k := fmt.Sprintf("test-%d", i)

				testMap.Put(k, i)
				wg.Done()
			}(i)
		}
		wg.Wait()

		assert(t, testMap, "test-998", 998)
	})
}

func assert(t testing.TB, result *LockingMap[string, int], key string, want int) {
	t.Helper()
	if result.Fields[key] != want {
		t.Fail()
	}
}
