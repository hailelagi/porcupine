package porcupine

import (
	"log"
	"math/rand"
	"testing"
)

func FuzzInsertKeys(f *testing.F) {
	var tree BTree

	for key := 1; key < 10_000; key++ {
		f.Add(key)
	}

	f.Fuzz(func(t *testing.T, key int) {
		tree.Insert(key)
		found := keyExists(&tree, key)

		if !found {
			t.Errorf("not found %v", key)
		}
	})
}

func FuzzSearchKeys(f *testing.F) {
	var tree BTree

	for key := 1; key < 10_000; key++ {
		f.Add(key)
	}

	f.Fuzz(func(t *testing.T, key int) {
		var found bool
		tree.Insert(key)

		data, _, err := tree.Search(key)

		if err != nil {
			t.Errorf("could not search tree %v", err)
		}

		for _, d := range data {
			if d == key {
				found = true
			}
		}

		if !found {
			t.Errorf("did not find key inserted")
		}
	})
}

func FuzzDeleteKeys(f *testing.F) {
	var tree BTree

	for key := 1; key < 10_000; key++ {
		f.Add(key)
	}

	f.Fuzz(func(t *testing.T, key int) {
		tree.Insert(key)
		err := tree.Delete(key)

		if err != nil {
			t.Errorf("deletion errored %v", err)
		}

		v, _, _ := tree.root.Search(key)

		for _, d := range v.data {
			if d == key {
				t.Errorf("found deleted key/value %v", v)
			}
		}

	})
}

func keyExists(t *BTree, key int) bool {
	n, _, _ := t.root.Search(key)

	for _, v := range n.data {
		if v == key {
			return true
		}
	}

	return false
}

func BenchmarkBTree(b *testing.B) {
	var tree BTree

	for i := 0; i <= 100_000; i++ {
		key := i
		// value := i * 10
		err := tree.Insert(key)
		if err != nil {
			b.Errorf("Error inserting key %d: %v", key, err)
		}
	}

	b.ResetTimer()

	b.Run("write", func(pb *testing.B) {
		for i := 10_000; i < pb.N; i++ {
			// value := i * 10
			key := rand.Intn(100_000)
			err := tree.Insert(key)

			if err == ErrDuplicateKey {
				log.Print("warn duplicate insert")
			} else if err != nil {
				b.Errorf("Error inserting key %d: %v", i, err)
			}
		}
	})

	log.Printf("current node count: %v", tree.nodeCount)

	b.Run("access", func(pb *testing.B) {
		for i := 0; i < pb.N; i++ {
			key := rand.Intn(100_000)

			n, idx, err := tree.Search(key)
			if err != nil {
				b.Errorf("Error searching key %d: %v node: %v at: %v", key, err, n, idx)
			}
		}
	})

	log.Printf("current node count: %v", tree.nodeCount)

	b.Run("read/write", func(pb *testing.B) {
		for i := 0; i <= pb.N; i++ {
			key := rand.Intn(100_000)
			err := tree.Insert(key)
			n, idx, searchErr := tree.Search(key)

			if searchErr != nil {
				b.Errorf("Error searching key %d: %v node: %v at: %v", i, err, n, idx)
			}

			if err != nil {
				b.Logf("warning inserting %d: %v", i, err)
			}
		}
	})

	log.Printf("current node count: %v", tree.nodeCount)
}

func BenchmarkBTreeConcurrentAccess(b *testing.B) {
	var tree BTree

	for i := 0; i <= 100_000; i++ {
		key := i
		// value := i * 10
		err := tree.Insert(key)
		if err != nil {
			b.Errorf("Error inserting key %d: %v", key, err)
		}
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(100_000)
			n, idx, err := tree.Search(key)

			if err != nil {
				b.Errorf("Error searching key %d: %v node: %v at: %v", key, err, n, idx)
			}
		}
	})
}

func BenchmarkBTreeConcurrentWriter(b *testing.B) {
	var tree BTree

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Intn(100_000)
			err := tree.Insert(key)
			if err != nil {
				b.Errorf("Error inserting key %d: %v", key, err)
			}
		}
	})
}
