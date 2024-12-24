package porcupine

import "testing"

func TestRadixTree_InsertAndSearch(t *testing.T) {
	tests := []struct {
		insertions []string
		search     string
		expected   bool
	}{
		{[]string{"hello", "world"}, "hello", true},
		{[]string{"hello", "world"}, "world", true},
		{[]string{"hello", "world"}, "hell", false},
		{[]string{"hello", "world"}, "worlds", false},
	}

	for _, tt := range tests {
		tree := NewRadixTree()
		for _, word := range tt.insertions {
			tree.Insert(word)
		}

		if result := tree.Search(tt.search); result != tt.expected {
			t.Errorf("Search(%q) = %v; want %v", tt.search, result, tt.expected)
		}
	}
}

func TestRadixTree_StartsWith(t *testing.T) {
	tests := []struct {
		insertions []string
		prefix     string
		expected   bool
	}{
		{[]string{"hello", "world"}, "hell", true},
		{[]string{"hello", "world"}, "wor", true},
		{[]string{"hello", "world"}, "heaven", false},
		{[]string{"hello", "world"}, "worlz", false},
	}

	for _, tt := range tests {
		tree := NewRadixTree()
		for _, word := range tt.insertions {
			tree.Insert(word)
		}

		if result := tree.StartsWith(tt.prefix); result != tt.expected {
			t.Errorf("StartsWith(%q) = %v; want %v", tt.prefix, result, tt.expected)
		}
	}
}
