package storage

type Storage[K any, V any] interface {
	Add(K, V) error

	Remove(K) error

	Get(K) (V, error)
}
