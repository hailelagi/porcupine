package porcupine

// Store is an interface for a key-value store.
type Store interface {
	// todo: make generic later
	// Write(string, interface{}) interface{}
	Put(string, int) int
	Del(string)

	// Read(string) interface{}
	Get(string) int
	In(string) bool
}

// Porcupine is a global in-memory read/write store.
type Porcupine struct {
	Store        Store
	Name         string
	env          string
	processCount int
}

// New `Porcupine` instance.
func NewPorcupine(storeConfig string) *Porcupine {
	var store Store

	switch storeConfig {
	case "hashmap":
		store = &LockingMap{}
	// todo: add supported data structures
	default:
		store = &LockingMap{}
	}

	return &Porcupine{Store: store, env: "dev", Name: "hashMap", processCount: 1}
}
