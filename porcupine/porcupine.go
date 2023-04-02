package porcupine

type Store interface {
	// todo: make generic later
	// Write(string, interface{}) interface{} ?
	Put(string, int) int
	Del(string)

	// Read(string) interface{}
	Get(string) int
	In(string) bool
}

// Provides global access to a store and analytics
type Porcupine struct {
	store        Store
	env          string
	processCount int
}

func NewPorcupine(storeConfig string) *Porcupine {
	var store Store

	switch storeConfig {
	case "hashmap":
		store = &LockingMap{}
	// todo: add supported data structures
	default:
		store = &LockingMap{}
	}

	return &Porcupine{store: store, env: "dev", processCount: 1}
}
