package porcupine

import (
	"math/rand"
	"testing"
)

func TestBSTEmpty(t *testing.T) {
	bst := NewBSTree()

	_, err := bst.Get(10)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestBSTPutAndGet(t *testing.T) {
	bst := NewBSTree()
	table := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for _, i := range table {
		err := bst.Put(i, i*100)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	for _, i := range table {
		value, err := bst.Get(i)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if value != i*100 {
			t.Errorf("Expected value, got %v", value)
		}
	}

}

func BenchmarkBSTGet(b *testing.B) {
	bst := NewBSTree()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(7)
			err := bst.Put(key, key+100)

			if err != nil {
				b.Errorf("fail during get bench err: %v", err)
			}
		}
	})
}

func BenchmarkBSTMostlyReads(b *testing.B) {
	bst := NewBSTree()

	// Initialize the map with some data
	for i := 1; i <= 1000; i++ {
		bst.Put(i, i*100)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(10) + 1
			_, err := bst.Get(key)

			if err != nil {
				b.Errorf("fail during get bench err: %v, key: %v", err, key)
			}
		}
	})
}

func BenchmarkBSTMostlyWrites(b *testing.B) {
	bst := NewBSTree()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(10)
			err := bst.Put(key, key*100)

			if err != nil {
				b.Errorf("fail during get bench err: %v, key: %v", err, key)
			}
		}
	})
}
