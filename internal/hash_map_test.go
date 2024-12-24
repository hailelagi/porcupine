package porcupine

import (
	"fmt"
	"sync"
	"testing"
)

func TestLockingMappCounter(t *testing.T) {

	t.Run("increment", func(t *testing.T) {
		testMap := NewLockingMap()

		testMap.Put("test", 1)
		testMap.Put("test-x", 2)
		testMap.Put("test-y", 3)

		if res, err := testMap.Get("test"); err == nil {
			if res != 1 {
				t.Fail()
			}
		}

		if res, err := testMap.Get("test-x"); err == nil {
			if res != 2 {
				t.Fail()
			}
		}

		if res, err := testMap.Get("test-y"); err == nil {
			if res != 3 {
				t.Fail()
			}
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		testMap := NewLockingMap()

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

		if res, err := testMap.Get("test-998"); err == nil {
			if res != 998 {
				t.Fail()
			}
		}
	})
}
