package porcupine

import (
	"math/rand"
	"testing"
)

func TestRedBlackEmpty(t *testing.T) {
	rb := NewRedBlack()

	_, err := rb.Get(10)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestRedBlackPutAndGet(t *testing.T) {
	rb := NewRedBlack()
	table := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for _, i := range table {
		err := rb.Put(i, i*100)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	for _, i := range table {
		value, err := rb.Get(i)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if value != i*100 {
			t.Errorf("Expected value, got %v", value)
		}
	}

}

func BenchmarkRedBlackReadAndWrite(b *testing.B) {
	rb := NewRedBlack()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(100_000)
			err := rb.Put(key, key+100)
			v, errGet := rb.Get(key)

			if v != key+100 || errGet != nil {
				b.Error("fail during r")
			}
			if err != nil {
				b.Errorf("fail during bench err: %v", err)
			}
		}
	})
}

func BenchmarkRedBlackMostlyReads(b *testing.B) {
	rb := NewRedBlack()

	for i := 1; i <= 10_000; i++ {
		key := rand.Intn(10_000) + 1
		rb.Put(key, key*42)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(10_000) + 1
			v, err := rb.Get(key)

			if v != key*42 && err == nil {
				b.Errorf("fail during bench err")
			}
		}
	})
}

func BenchmarkRedBlackMostlyWrites(b *testing.B) {
	rb := NewRedBlack()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(10_000) + 1
			err := rb.Put(key, key*100)

			if err != nil {
				b.Errorf("fail during bench err: %v, key: %v", err, key)
			}
		}
	})
}

func BenchmarkRedBlackUnbalancedRead(b *testing.B) {
	rb := NewRedBlack()

	for i := 1; i <= 10_000; i++ {
		rb.Put(i, i*42)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(10_000) + 1
			value, err := rb.Get(key)

			if value != key*42 || err != nil {
				b.Errorf("fail during bench w")
			}
		}
	})
}

func BenchmarkRedBlackUnbalancedWrite(b *testing.B) {
	rb := NewRedBlack()

	for i := 1; i <= 10_000; i++ {
		rb.Put(i, i*100)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(10_000) + 1
			err := rb.Put(key, key*42)

			if err != nil {
				b.Errorf("fail during bench w")
			}
		}
	})
}