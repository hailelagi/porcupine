package porcupine

import (
	"math/rand"
	"testing"
)

/*
50/50 read and write balance
*/
func BenchmarkConcurrentMap(b *testing.B) {
	myMap := NewMap(8)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(7)
			myMap.Increment(key)
			_ = myMap.GetValue(key)
		}
	})
}

func BenchmarkMapSingleMutex(b *testing.B) {
	myMap := NewMapSingleMutex()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(58)
			myMap.Increment(key)
			_ = myMap.GetValue(key)
		}
	})
}

/*
Workload skews towards writes
*/

func BenchmarkMapSingleMutexMostlyWrites(b *testing.B) {
	myMap := NewMapSingleMutex()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(7)
			myMap.Increment(key)
		}
	})
}

func BenchmarkConcurrentMapMostlyWrites(b *testing.B) {
	myMap := NewMap(8)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(7)
			myMap.Increment(key)
		}
	})
}

/*
Workload skews towards reads
*/

func BenchmarkMapSingleMutexMostlyReads(b *testing.B) {
	myMap := NewMapSingleMutex()

	// Initialize the map with some data
	for i := 0; i < 1000; i++ {
		myMap.Increment(i)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(7)
			_ = myMap.GetValue(key)
		}
	})
}

func BenchmarkConcurrentMapMostlyReads(b *testing.B) {
	myMap := NewMap(8)

	// Initialize the map with some data
	for i := 0; i < 1000; i++ {
		myMap.Increment(i)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(7)
			_ = myMap.GetValue(key)
		}
	})
}
