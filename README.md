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

```zsh
porcupine get <key>              # get a value by key
porcupine set <key> <value>      # set key to a value
porcupine ls                     # list stored keys
porcupine rm  <key>              # remove an entry by key
porcupine config --list          # list data-structures
porcupine config <datastore>     # choose data-structure 
porcupine bench
```
