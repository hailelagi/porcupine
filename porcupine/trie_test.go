package porcupine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	tests := []struct {
		name     string
		method   string
		input    string
		expected bool
	}{
		{"Search hello", "search", "hello", true},
		{"Search world", "search", "world", true},
		{"Search hell", "search", "hell", false},
		{"StartsWith hell", "startsWith", "hell", true},
		{"StartsWith wo", "startsWith", "wo", true},
		{"StartsWith hi", "startsWith", "hi", false},
		{"Search hell after insert", "search", "hell", true},
		{"StartsWith he", "startsWith", "he", true},
		{"Search helloo", "search", "helloo", false},
		{"StartsWith helloo", "startsWith", "helloo", false},
	}

	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("hell")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.method == "search" {
				assert.Equal(t, tt.expected, trie.Search(tt.input))
			} else if tt.method == "startsWith" {
				assert.Equal(t, tt.expected, trie.StartsWith(tt.input))
			}
		})
	}
}
