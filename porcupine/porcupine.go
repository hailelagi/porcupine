package porcupine

import constraints "golang.org/x/exp/constraints"

// Table is an interface for an unordered key-value data structure
type Table[Key comparable, Value any] interface {
	Get(Key) (Value, error)
	Put(Key, Value)
	Del(Key)
	In(Key) bool
}

// OrderTable is an interface for an unordered key-value data structure
type OrderTable[Key constraints.Ordered, Value any] interface {
	Get(Key) (Value, error)
	Put(Key, Value)
	Del(Key)
	In(Key) bool
}

// Store is an interface for a key-value store.
type Store[Key any, Value any] interface {
	Get(Key) (Value, error)
	Put(Key, Value)
	Del(Key)
	In(Key) bool
}

// StoreCluster defines the strategy and api implementation of a multi-node Store.
// a `Cluster` must specify a strategy and a mode:
// strategies:
// 1. single-writer // mpsc // single write + replicas
// 2. multi-writer - defaults to raft for consensus

// mode:
// 1. available
// 2. consistent
type StoreCluster[Key comparable, Value any] interface {
	Store[Key, Value]
	Strategy() string
	Mode() string
}

// todo constrain input on ["single-writer-replica", "multi-writer-replic"]
type mode = string

// todo constrain input on ["available", "consistent"]
type strategy = string

type ClusterConfig struct {
	strategy  strategy
	mode      mode
	instances int32
	ports     []int32
}

// Porcupine is a global in-memory read/write store.
type Porcupine struct {
	Store        Store[string, any]
	Name         string
	env          string
	processCount int
}

// New `Porcupine` instance. This should not be copied after instantiation.
// todo: hold a *Porcupine on init
func NewPorcupine(storeConfig string) Porcupine {
	var store Store[string, any]

	switch storeConfig {
	case "hashmap":
		store = &LockingMap[string, any]{}
	// todo: add supported data structures
	default:
		panic("todo")
	}

	return Porcupine{Store: store, env: "dev", Name: "hashMap", processCount: 1}
}

func SpawnPorcupines(storeConfig string, clusterConfig string) []Porcupine {
	var cluster ClusterConfig
	var nodes []Porcupine

	cluster = ClusterConfig{
		strategy:  "single-writer-replica",
		mode:      "available",
		instances: 2,
		ports:     []int32{8080, 8081},
	}

	for i := 0; i >= int(cluster.instances); i++ {
		node := NewPorcupine(storeConfig)
		nodes = append(nodes, node)
	}

	return nodes
}
