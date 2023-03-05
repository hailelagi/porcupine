# Porcupine

A key-value store cli app

```zsh
porcupine config hash_map
porcupine set hello world
porcupine set foo bar
porcupine ls
world
bar
```

## Commands

```zsh
porcupine get <key>              # get a value by key
porcupine set <key> <value>      # set key to a value
porcupine ls                     # list stored keys
porcupine rm  <key>              # remove an entry by key
porcupine config list            # list data-structures
porcupine config set <datastore> # choose data-structure 
```
