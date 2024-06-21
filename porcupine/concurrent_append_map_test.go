package porcupine

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentAppendMapCounter(t *testing.T) {

	t.Run("increment", func(t *testing.T) {
		testMap := NewConcurrentAppendMap()

		testMap.Put("test", 1)
		testMap.Put("test-x", 2)
		testMap.Put("test-y", 3)

		if testMap.Get("test") != 1 {
			t.Fail()
		}

		if testMap.Get("test-x") != 2 {
			t.Fail()
		}

		if testMap.Get("test-y") != 3 {
			t.Fail()
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		testMap := NewConcurrentAppendMap()

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

		if testMap.Get("test-998") != 998 {
			t.Fail()
		}
	})
}
