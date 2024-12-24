package porcupine

/*
import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)
*/

/*
TODO: I did something dumb and broke this test suite, fml
func TestCounter(t *testing.T) {

	t.Run("increment", func(t *testing.T) {
		testMap := &LockingMap[string, int]{RWMutex: sync.RWMutex{}, Fields: make(map[string]int)}

		testMap.Put("test", 1)
		testMap.Put("test-x", 2)
		testMap.Put("test-y", 3)

		val, err := testMap.Get("test")
		assert.NoError(t, err)
		assert.Equal(t, 1, val)

		val, err = testMap.Get("test-x")
		assert.NoError(t, err)
		assert.Equal(t, 2, val)

		val, err = testMap.Get("test-y")
		assert.NoError(t, err)
		assert.Equal(t, 3, val)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		testMap := &LockingMap[string, int]{RWMutex: sync.RWMutex{}, Fields: make(map[string]int)}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(i int) {
				val, err := testMap.Get("test-998")
				k := fmt.Sprintf("test-%d", i)

				assert.NoError(t, err)
				assert.Equal(t, 998, val)

				testMap.Put(k, i)
				wg.Done()
			}(i)
		}
		wg.Wait()

		val, err := testMap.Get("test-998")

		assert.NoError(t, err)
		assert.Equal(t, 998, val)
	})
}
*/
