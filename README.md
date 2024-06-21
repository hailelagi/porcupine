# Porcupine

Key value store experiments, data structures and exploratory concept examples/snippets.

## what kinda works

```zsh
cd /porcupine
go test .
go test -bench=.
```

and

```zsh
porcupine server
porcupine graph
```

## what would be nice to do
- a vanilla AVL Tree
- a vanilla red black tree
- add an ART index
- a CTrie  https://en.wikipedia.org/wiki/Ctrie
- a hash array mapped tree
// https://www.cs.umd.edu/class/fall2017/cmsc430/slides/HAMT.pdf
// https://worace.works/2016/05/24/hash-array-mapped-tries/
// https://en.wikipedia.org/wiki/Hash_array_mapped_trie
- a merkle tree
- a simple LSM
  
```zsh
porcupine get <key>              # get a value by key
porcupine set <key> <value>      # set key to a value
porcupine ls                     # list stored keys
porcupine rm  <key>              # remove an entry by key
porcupine config --list          # list data-structures
porcupine config <datastore>     # choose data-structure 
porcupine bench
```
