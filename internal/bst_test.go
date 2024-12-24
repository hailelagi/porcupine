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

func BenchmarkBSTReadAndWrite(b *testing.B) {
	bst := NewBSTree()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(100_000)
			err := bst.Put(key, key+100)
			v, errGet := bst.Get(key)

			if v != key+100 || errGet != nil {
				b.Error("fail during r")
			}
			if err != nil {
				b.Errorf("fail during bench err: %v", err)
			}
		}
	})
}

// tree is unbalanced and therefore searches border on linear access.
func BenchmarkBSTMostlyReads(b *testing.B) {
	bst := NewBSTree()

	for i := 1; i <= 10_000; i++ {
		key := rand.Intn(10_000) + 1
		bst.Put(key, key*42)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(10_000) + 1
			v, err := bst.Get(key)

			if v != key*42 && err == nil {
				b.Errorf("fail during bench err")
			}
		}
	})
}

func BenchmarkBSTMostlyWrites(b *testing.B) {
	bst := NewBSTree()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(10_000) + 1
			err := bst.Put(key, key*100)

			if err != nil {
				b.Errorf("fail during bench err: %v, key: %v", err, key)
			}
		}
	})
}

func BenchmarkBSTUnbalancedRead(b *testing.B) {
	bst := NewBSTree()

	for i := 1; i <= 10_000; i++ {
		bst.Put(i, i*42)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(10_000) + 1
			value, err := bst.Get(key)

			if value != key*42 || err != nil {
				b.Errorf("fail during bench w")
			}
		}
	})
}

func BenchmarkBSTUnbalancedWrite(b *testing.B) {
	bst := NewBSTree()

	for i := 1; i <= 10_000; i++ {
		bst.Put(i, i*100)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(10_000) + 1
			err := bst.Put(key, key*42)

			if err != nil {
				b.Errorf("fail during bench w")
			}
		}
	})
}
