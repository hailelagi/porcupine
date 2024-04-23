package porcupine

/*
func BenchmarkInsertBTree(b *testing.B) {
	tree, err := NewBPlusTree("datafile.csv")
	if err != nil {
		b.Fatalf("Error opening tree file: %v", err)
	}

	for i := 0; i < 100_000; i++ {
		key := i
		value := i * 10
		err := tree.Insert(key, value)
		if err != nil {
			b.Errorf("Error inserting key %d: %v", key, err)
		}
	}
}

func BenchmarkAccessBTree(b *testing.B) {
	tree, err := NewBPlusTree("datafile.csv")

	if err != nil {
		b.Fatalf("Error opening tree file: %v", err)
	}

	for i := 0; i < 100_000; i++ {
		key := i
		value := i * 10
		err := tree.Insert(key, value)
		if err != nil {
			b.Errorf("Error inserting key %d: %v", key, err)
		}
	}

	for i := 0; i < 100_000; i++ {
		key := i
		value := i * 10
		err := tree.GetValue(key, value)
		if err != nil {
			b.Errorf("Error inserting key %d: %v", key, err)
		}
	}

}
*/
