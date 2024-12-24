package porcupine

// import (
// 	"math/rand"
// 	"testing"
// )

// func TestAVLEmpty(t *testing.T) {
// 	avl := NewAVL()

// 	_, err := avl.Get(10)
// 	if err == nil {
// 		t.Errorf("Expected error, got nil")
// 	}
// }

// func TestAVLPutAndGet(t *testing.T) {
// 	avl := NewAVL()
// 	table := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

// 	for _, i := range table {
// 		err := avl.Put(i, i*100)

// 		if err != nil {
// 			t.Errorf("Unexpected error: %v", err)
// 		}
// 	}

// 	for _, i := range table {
// 		value, err := avl.Get(i)
// 		if err != nil {
// 			t.Errorf("Unexpected error: %v", err)
// 		}
// 		if value != i*100 {
// 			t.Errorf("Expected value, got %v", value)
// 		}
// 	}

// }

// func BenchmarkAVLReadAndWrite(b *testing.B) {
// 	avl := NewAVL()
// 	b.ResetTimer()

// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			key := rand.Intn(100_000)
// 			err := avl.Put(key, key+100)
// 			v, errGet := avl.Get(key)

// 			if v != key+100 || errGet != nil {
// 				b.Error("fail during r")
// 			}
// 			if err != nil {
// 				b.Errorf("fail during bench err: %v", err)
// 			}
// 		}
// 	})
// }

// func BenchmarkAVLMostlyReads(b *testing.B) {
// 	avl := NewAVL()

// 	for i := 1; i <= 10_000; i++ {
// 		key := rand.Intn(10_000) + 1
// 		avl.Put(key, key*42)
// 	}

// 	b.ResetTimer()

// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			key := rand.Intn(10_000) + 1
// 			v, err := avl.Get(key)

// 			if v != key*42 && err == nil {
// 				b.Errorf("fail during bench err")
// 			}
// 		}
// 	})
// }

// func BenchmarkAVLMostlyWrites(b *testing.B) {
// 	AVL := NewAVL()
// 	b.ResetTimer()

// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			key := rand.Intn(10_000) + 1
// 			err := AVL.Put(key, key*100)

// 			if err != nil {
// 				b.Errorf("fail during bench err: %v, key: %v", err, key)
// 			}
// 		}
// 	})
// }

// func BenchmarkAVLUnbalancedRead(b *testing.B) {
// 	avl := NewAVL()

// 	for i := 1; i <= 10_000; i++ {
// 		avl.Put(i, i*42)
// 	}

// 	b.ResetTimer()

// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			key := rand.Intn(10_000) + 1
// 			value, err := avl.Get(key)

// 			if value != key*42 || err != nil {
// 				b.Errorf("fail during bench w")
// 			}
// 		}
// 	})
// }

// func BenchmarkAVLUnbalancedWrite(b *testing.B) {
// 	avl := NewAVL()

// 	for i := 1; i <= 10_000; i++ {
// 		avl.Put(i, i*100)
// 	}

// 	b.ResetTimer()

// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			key := rand.Intn(10_000) + 1
// 			err := avl.Put(key, key*42)

// 			if err != nil {
// 				b.Errorf("fail during bench w")
// 			}
// 		}
// 	})
// }
