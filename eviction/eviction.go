package eviction

type Eviction[K any] interface {
	KeyAccessed(K)

	Evict() K
}
