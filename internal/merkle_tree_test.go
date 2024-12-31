package porcupine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerkleTrieInsert(t *testing.T) {
	trie := NewMerkleTrie()
	words := []string{"hi", "hello", "hey"}

	for _, word := range words {
		trie.Insert(word)
	}

	for _, word := range words {
		proofs, exists := trie.GetProof(word)
		if !exists {
			t.Errorf("Expected word %s to exist in trie", word)
		}
		if len(proofs) == 0 && len(words) > 1 {
			t.Errorf("Expected proofs for word %s", word)
		}
	}

	if _, exists := trie.GetProof("none"); exists {
		t.Fail()
	}
}

func TestMerkleTrieRootHash(t *testing.T) {
	trie := NewMerkleTrie()
	trie2 := NewMerkleTrie()
	emptyHash := trie.GetSignedHead()

	trie.Insert("test")
	trie2.Insert("test")

	newHash := trie.GetSignedHead()

	assert.NotEqual(t, emptyHash, newHash)
	assert.Equal(t, trie2.GetSignedHead(), newHash)
}

// wtf?
func TestMerkleTrieProofVerification(t *testing.T) {
	trie := NewMerkleTrie()

	for _, word := range []string{"apple", "app", "banana", "band"} {
		trie.Insert(word)
	}

	tests := []struct {
		word    string
		exists  bool
		message string
	}{
		{"apple", true, "Should verify proof for existing word"},
		{"app", true, "Should verify proof for existing prefix"},
		{"ban", false, "Should not verify proof for partial word"},
		{"banana", true, "Should verify proof for existing word"},
		{"notexist", false, "Should not verify proof for non-existent word"},
	}

	for _, test := range tests {
		t.Run(test.word, func(t *testing.T) {
			proofs, exists := trie.GetProof(test.word)
			if exists != test.exists {
				t.Errorf("Expected exists=%v for word %s", test.exists, test.word)
				return
			}

			if !test.exists {
				return
			}

			// Verify the proof
			if !trie.VerifyProof(test.word, proofs) {
				t.Errorf("Proof verification failed for word %s", test.word)
			}

			// Verify proof fails for wrong word
			if trie.VerifyProof(test.word+"wrong", proofs) {
				t.Errorf("Proof verification should fail for wrong word")
			}
		})
	}
}

/*
func TestMerkleTrieConsistency(t *testing.T) {
	trie := NewMerkleTrie()
	words := []string{"a", "ab", "abc", "abcd"}

	var previousHash []byte

	// Insert words incrementally and verify consistency
	for i, word := range words {
		trie.Insert(word)
		currentHash := trie.GetSignedHead()

		if i > 0 && bytes.Equal(currentHash, previousHash) {
			t.Errorf("Hash should change after inserting %s", word)
		}

		// Verify all previously inserted words still have valid proofs
		for j := 0; j <= i; j++ {
			proofs, exists := trie.GetProof(words[j])
			if !exists {
				t.Errorf("Word %s should exist after inserting %s", words[j], word)
				continue
			}

			if !trie.VerifyProof(words[j], proofs) {
				t.Errorf("Proof verification failed for %s after inserting %s", words[j], word)
			}
		}

		previousHash = currentHash
	}
}

func TestMerkleTrieEmptyAndSingleNode(t *testing.T) {
	trie := NewMerkleTrie()

	if _, exists := trie.GetProof(""); exists {
		t.Error("Empty string should not exist in trie")
	}

	// Single node
	trie.Insert("single")
	proofs, exists := trie.GetProof("single")

	if !exists {
		t.Error("Single word should exist in trie")
	}

	if !trie.VerifyProof("single", proofs) {
		t.Error("Proof verification failed for single word")
	}

	// Empty string after insertion
	if _, exists := trie.GetProof(""); exists {
		t.Error("Empty string should not exist in trie after insertion")
	}
}

func TestMerkleTrieProofOrdering(t *testing.T) {
	trie := NewMerkleTrie()
	words := []string{"cab", "abc", "bac"}

	// Insert in specific order
	for _, word := range words {
		trie.Insert(word)
	}

	// Get proofs for each word
	proofs1, _ := trie.GetProof("abc")
	proofs2, _ := trie.GetProof("abc")

	// Verify proofs are deterministic
	if len(proofs1) != len(proofs2) {
		t.Error("Proof length should be consistent for same word")
	}

	for i := range proofs1 {
		if !bytes.Equal(proofs1[i].Digest, proofs2[i].Digest) {
			t.Error("Proofs should be deterministic for same word")
		}
	}
}
*/
