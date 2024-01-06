package porcupine

import (
	"fmt"
	"sync"
)

// FineGrainedCounter represents a simple concurrent counter with fine-grained locking.
type FineGrainedCounter struct {
	counts map[string]int
	locks  map[string]*sync.Mutex
}

// NewFineGrainedCounter creates a new FineGrainedCounter.
func NewFineGrainedCounter() *FineGrainedCounter {
	return &FineGrainedCounter{
		counts: make(map[string]int),
		locks:  make(map[string]*sync.Mutex),
	}
}

// Increment increments the counter associated with the given key.
func (c *FineGrainedCounter) Increment(key string) {
	// Lock the specific counter associated with the key
	c.getLock(key).Lock()
	defer c.getLock(key).Unlock()

	// Increment the counter
	c.counts[key]++
}

// GetValue returns the current value of the counter associated with the given key.
func (c *FineGrainedCounter) GetValue(key string) int {
	// Lock the specific counter associated with the key
	c.getLock(key).Lock()
	defer c.getLock(key).Unlock()

	// Return the current value of the counter
	return c.counts[key]
}

// getLock returns the lock associated with the given key.
func (c *FineGrainedCounter) getLock(key string) *sync.Mutex {
	// Create a lock if it doesn't exist for the given key
	if c.locks[key] == nil {
		c.locks[key] = &sync.Mutex{}
	}
	return c.locks[key]
}

func PrintFine() {
	counter := NewFineGrainedCounter()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := fmt.Sprintf("Key%d", index%3) // Using 3 different keys for demonstration
			counter.Increment(key)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final counter values
	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("Key%d", i)
		fmt.Printf("Counter for %s: %d\n", key, counter.GetValue(key))
	}
}
