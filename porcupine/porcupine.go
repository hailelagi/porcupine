package porcupine

// Store is an interface for a key-value store.
type Store[Key string, Value any] interface {
	Get(Key) (Value, error)
	Put(Key, Value)
	Del(Key)
	In(Key) bool
}

// Porcupine is a global in-memory read/write store.
type Porcupine struct {
	Store        Store[string, any]
	Name         string
	env          string
	processCount int
}

// New `Porcupine` instance.
func NewPorcupine(storeConfig string) *Porcupine {
	var store Store[string, any]

	switch storeConfig {
	case "hashmap":
		store = &LockingMap[string, any]{}
	// todo: add supported data structures
	default:
		panic("todo")
	}

	return &Porcupine{Store: store, env: "dev", Name: "hashMap", processCount: 1}
}
