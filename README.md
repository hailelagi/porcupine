# Porcupine (go)

A tiny distributed key-value store server for localhost. See also: [porcupine-rs](https://github.com/hailelagi/porcupine-rs)

## Getting started

Start a single node porcupine, which default to `localhost:8080`:

```zsh
porcupine server
```

and issue it commands in a new terminal session:

```zsh
porcupine ls
```

## Forming A key-value cluster

Spin up as many porcupine servers as you'd like, for e.g:

```zsh
porcupine server --port-range=8080..8084 --cluster=5
```

Or

```zsh
porcupine server --port=8080 --cluster="myClusterName"
# in new shell
porcupine server --port=3000 --cluster="myClusterName"
```

## Commands

```zsh
porcupine get <key>              # get a value by key
porcupine set <key> <value>      # set key to a value
porcupine ls                     # list stored keys
porcupine rm  <key>              # remove an entry by key
porcupine config --list          # list data-structures
porcupine config <datastore>     # choose data-structure 
```
