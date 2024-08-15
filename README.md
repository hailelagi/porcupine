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
- A ring buffer ala [postgres style](https://www.interdb.jp/pg/pgsql08/05.html)
- a [CTrie](https://en.wikipedia.org/wiki/Ctrie)
- [a hash array mapped tree](https://en.wikipedia.org/wiki/Hash_array_mapped_trie)
- a merkle tree

```zsh
porcupine get <key>              # get a value by key
porcupine set <key> <value>      # set key to a value
porcupine ls                     # list stored keys
porcupine rm  <key>              # remove an entry by key
porcupine config --list          # list data-structures
porcupine config <datastore>     # choose data-structure 
porcupine bench
```
