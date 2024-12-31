package porcupine

import (
	"bytes"
	"crypto/sha256"
	"sort"
)

// https://www.ralphmerkle.com/papers/Protocols.pdf
// https://git-scm.com/book/id/v2/Git-Internals-Git-Objects

// https://transparency.dev/verifiable-data-structures/
// https://github.com/google/trillian/blob/master/docs/papers/VerifiableDataStructures.pdf
// see also: https://transparency.dev/articles/tile-based-logs/

// https://ethereum.org/en/developers/docs/data-structures-and-encoding/patricia-merkle-trie/

/*
whoami: https://crt.sh/?q=google.com

proof of inclusion: position + H(neighbour)

properties:
- byzantine fault tolerance (we catch lying logs... sort of) (hard problem!)
- inclusion proof
- log consistency(prefix) proof
- fork consistency (gossip/cmpexchng tree head)
- imperfect design, sort of okay in practice
- audit is easier then true byzantine fault tolerance

[](1) -> [](2) -> [](3)
              \
			  -- > (lying is bad!)
*/

type MerkleTrie struct {
	root *MerkleTrieNode
}

type Proof struct {
	Digest []byte
	Key    string
}

type MerkleTrieNode struct {
	children map[rune]*MerkleTrieNode
	proof    Proof
	isEnd    bool
	value    string
}

func NewMerkleTrie() *MerkleTrie {
	return &MerkleTrie{
		root: &MerkleTrieNode{
			children: make(map[rune]*MerkleTrieNode),
		},
	}
}

func (t *MerkleTrie) GetSignedHead() []byte {
	return t.root.proof.Digest
}

func (t *MerkleTrie) Insert(word string) {
	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = &MerkleTrieNode{
				children: make(map[rune]*MerkleTrieNode),
			}
		}
		current = current.children[ch]
	}

	current.isEnd = true
	current.value = word
	t.updateHashes()
}

func (n *MerkleTrieNode) hash() []byte {
	h := sha256.New()

	if n.isEnd {
		h.Write([]byte(n.value))
	}

	// Hash children's digests in sorted order
	childKeys := make([]rune, 0, len(n.children))
	for k := range n.children {
		childKeys = append(childKeys, k)
	}
	sort.Slice(childKeys, func(i, j int) bool { return childKeys[i] <= childKeys[j] })

	for _, k := range childKeys {
		h.Write(n.children[k].proof.Digest)
	}

	return h.Sum(nil)
}

// Update hashes recursively from leaves to root
func (t *MerkleTrie) updateHashes() {
	var updateNode func(*MerkleTrieNode)
	updateNode = func(node *MerkleTrieNode) {
		for _, child := range node.children {
			updateNode(child)
		}
		node.proof.Digest = node.hash()
	}
	updateNode(t.root)
}

// GetProof returns a proof of inclusion for a given word
func (t *MerkleTrie) GetProof(word string) ([]Proof, bool) {
	proofs := []Proof{}
	current := t.root

	for _, ch := range word {
		node, exists := current.children[ch]
		if !exists {
			return nil, false
		}
		// Add sibling hashes to proof
		for k, sibling := range current.children {
			if k != ch {
				proofs = append(proofs, sibling.proof)
			}
		}
		current = node
	}

	if !current.isEnd || current.value != word {
		return nil, false
	}

	return proofs, true
}

// VerifyProof verifies a proof of inclusion
func (t *MerkleTrie) VerifyProof(word string, proofs []Proof) bool {
	h := sha256.New()
	h.Write([]byte(word))
	calculatedHash := h.Sum(nil)

	for _, proof := range proofs {
		h.Reset()
		h.Write(calculatedHash)
		h.Write(proof.Digest)
		calculatedHash = h.Sum(nil)
	}

	return bytes.Equal(calculatedHash, t.root.proof.Digest)
}
