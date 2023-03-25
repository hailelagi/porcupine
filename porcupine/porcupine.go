package porcupine

type Store interface {
	// todo: make generic later
	Put(string) int
	Get(string) int
	In(string) bool
	Del(string) bool
}
